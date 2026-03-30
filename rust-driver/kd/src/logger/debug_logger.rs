use alloc::vec;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU32, AtomicU64, Ordering};
use spin::Mutex;

use wdk_sys::{
    NTSTATUS,
    LARGE_INTEGER,
    PSTRING,
    ULONG,
    HANDLE,
    PVOID,
    IO_STATUS_BLOCK,
    STATUS_SUCCESS,
    BOOLEAN,
};
use wdk_sys::ntddk::{
    DbgSetDebugPrintCallback,
    KeQuerySystemTimePrecise,
    ExSystemTimeToLocalTime,
    RtlTimeToTimeFields,
    ZwWriteFile,
    ZwFlushBuffersFile,
};

const DEBUG_LOG_BUFFER_SIZE: usize = 0x1000 * 8;
const MAX_DBG_PRINT_LOG_LENGTH: usize = 512;

#[repr(C)]
pub struct DebugLogEntry {
    pub timestamp: LARGE_INTEGER,
    pub log_line_length: u16,
    pub log_line: [u8; 0],
}

impl DebugLogEntry {
    pub const fn size_with_line(line_len: usize) -> usize {
        core::mem::size_of::<LARGE_INTEGER>() + core::mem::size_of::<u16>() + line_len
    }
}

pub struct DebugLogBuffer {
    pub log_entries: Vec<u8>,
    pub next_log_offset: AtomicU32,
    pub overflowed_log_size: AtomicU32,
}

impl DebugLogBuffer {
    pub fn new() -> Self {
        Self {
            log_entries: vec![0u8; DEBUG_LOG_BUFFER_SIZE],
            next_log_offset: AtomicU32::new(0),
            overflowed_log_size: AtomicU32::new(0),
        }
    }

    pub fn clear(&mut self) {
        self.next_log_offset.store(0, Ordering::SeqCst);
        self.overflowed_log_size.store(0, Ordering::SeqCst);
        self.log_entries.fill(0);
    }
}

impl Default for DebugLogBuffer {
    fn default() -> Self {
        Self::new()
    }
}

pub struct PairedDebugLogBuffer {
    pub buffer_valid: AtomicBool,
    pub active_log_buffer: Mutex<usize>,
    pub buffers: [Mutex<DebugLogBuffer>; 2],
}

impl PairedDebugLogBuffer {
    pub const fn new() -> Self {
        Self {
            buffer_valid: AtomicBool::new(false),
            active_log_buffer: Mutex::new(0),
            buffers: [
                Mutex::new(DebugLogBuffer {
                    log_entries: Vec::new(),
                    next_log_offset: AtomicU32::new(0),
                    overflowed_log_size: AtomicU32::new(0),
                }),
                Mutex::new(DebugLogBuffer {
                    log_entries: Vec::new(),
                    next_log_offset: AtomicU32::new(0),
                    overflowed_log_size: AtomicU32::new(0),
                }),
            ],
        }
    }

    pub fn initialize(&self) -> bool {
        let mut buf0 = self.buffers[0].lock();
        buf0.log_entries = vec![0u8; DEBUG_LOG_BUFFER_SIZE];
        buf0.next_log_offset.store(0, Ordering::SeqCst);
        buf0.overflowed_log_size.store(0, Ordering::SeqCst);

        let mut buf1 = self.buffers[1].lock();
        buf1.log_entries = vec![0u8; DEBUG_LOG_BUFFER_SIZE];
        buf1.next_log_offset.store(0, Ordering::SeqCst);
        buf1.overflowed_log_size.store(0, Ordering::SeqCst);

        *self.active_log_buffer.lock() = 0;
        self.buffer_valid.store(true, Ordering::SeqCst);
        true
    }
}

static G_PAIRED_LOG_BUFFER: PairedDebugLogBuffer = PairedDebugLogBuffer::new();

pub unsafe fn save_debug_output_line(
    timestamp: &LARGE_INTEGER,
    log_line: &[u8],
    paired_log_buffer: &PairedDebugLogBuffer,
) {
    if log_line.is_empty() || log_line.len() > MAX_DBG_PRINT_LOG_LENGTH {
        return;
    }

    let line_len = if log_line.last() == Some(&b'\r') {
        log_line.len() - 1
    } else {
        log_line.len()
    };

    if line_len == 0 {
        return;
    }

    let log_entry_size = DebugLogEntry::size_with_line(line_len);

    let active_idx = *paired_log_buffer.active_log_buffer.lock();
    let mut buffer = paired_log_buffer.buffers[active_idx].lock();

    let current_offset = buffer.next_log_offset.load(Ordering::SeqCst) as usize;
    
    if current_offset + log_entry_size > DEBUG_LOG_BUFFER_SIZE {
        buffer.overflowed_log_size.fetch_add(log_entry_size as u32, Ordering::SeqCst);
        return;
    }

    let entries = buffer.log_entries.as_mut_ptr();
    let entry_ptr = entries.add(current_offset) as *mut u8;

    core::ptr::copy_nonoverlapping(
        timestamp as *const LARGE_INTEGER as *const u8,
        entry_ptr,
        core::mem::size_of::<LARGE_INTEGER>(),
    );

    let len_ptr = entry_ptr.add(core::mem::size_of::<LARGE_INTEGER>()) as *mut u16;
    core::ptr::write_volatile(len_ptr, line_len as u16);
    
    let log_line_ptr = entry_ptr.add(core::mem::size_of::<LARGE_INTEGER>() + core::mem::size_of::<u16>());
    core::ptr::copy_nonoverlapping(log_line.as_ptr(), log_line_ptr, line_len);

    buffer.next_log_offset.fetch_add(log_entry_size as u32, Ordering::SeqCst);
}

pub unsafe fn save_debug_output(
    output: &[u8],
    paired_log_buffer: &PairedDebugLogBuffer,
) {
    if output.is_empty() || output.len() > MAX_DBG_PRINT_LOG_LENGTH {
        return;
    }

    let mut timestamp: LARGE_INTEGER = core::mem::zeroed();
    KeQuerySystemTimePrecise(&mut timestamp);

    let mut start = 0;
    for (i, &byte) in output.iter().enumerate() {
        if byte == b'\n' {
            if start < i {
                save_debug_output_line(&timestamp, &output[start..i], paired_log_buffer);
            }
            start = i + 1;
        }
    }

    if start < output.len() {
        save_debug_output_line(&timestamp, &output[start..], paired_log_buffer);
    }
}

pub unsafe extern "C" fn debug_print_callback(
    output: PSTRING,
    component_id: ULONG,
    level: ULONG,
) {
    let _ = component_id;
    let _ = level;

    if output.is_null() {
        return;
    }

    let output_ref = &*output;
    if output_ref.Length == 0 || output_ref.Buffer.is_null() {
        return;
    }

    let output_slice = core::slice::from_raw_parts(
        output_ref.Buffer as *const u8,
        output_ref.Length as usize,
    );

    save_debug_output(output_slice, &G_PAIRED_LOG_BUFFER);
}

pub struct FlushBufferThreadContext {
    pub thread_exit_event: AtomicBool,
    pub paired_log_buffer: &'static PairedDebugLogBuffer,
    pub log_file_handle: AtomicU64,
    pub max_overflowed_log_size: AtomicU32,
}

impl FlushBufferThreadContext {
    pub const fn new() -> Self {
        Self {
            thread_exit_event: AtomicBool::new(false),
            paired_log_buffer: &G_PAIRED_LOG_BUFFER,
            log_file_handle: AtomicU64::new(0),
            max_overflowed_log_size: AtomicU32::new(0),
        }
    }
}

static G_THREAD_CONTEXT: FlushBufferThreadContext = FlushBufferThreadContext::new();

pub unsafe fn flush_debug_log_entries(thread_context: &FlushBufferThreadContext) -> NTSTATUS {
    let paired_log_buffer = thread_context.paired_log_buffer;

    let old_active_idx = *paired_log_buffer.active_log_buffer.lock();
    let new_active_idx = 1 - old_active_idx;
    
    {
        let mut active = paired_log_buffer.active_log_buffer.lock();
        *active = new_active_idx;
    }

    let old_buffer = paired_log_buffer.buffers[old_active_idx].lock();
    let next_offset = old_buffer.next_log_offset.load(Ordering::SeqCst) as usize;

    let file_handle = thread_context.log_file_handle.load(Ordering::SeqCst) as HANDLE;
    if file_handle.is_null() {
        return STATUS_SUCCESS;
    }

    let mut offset = 0;
    while offset < next_offset {
        let entry_ptr = old_buffer.log_entries.as_ptr().add(offset);
        
        let timestamp = core::ptr::read(entry_ptr as *const LARGE_INTEGER);
        let log_line_length = core::ptr::read(entry_ptr.add(core::mem::size_of::<LARGE_INTEGER>()) as *const u16) as usize;
        let log_line = core::slice::from_raw_parts(
            entry_ptr.add(core::mem::size_of::<LARGE_INTEGER>() + core::mem::size_of::<u16>()),
            log_line_length,
        );

        let mut write_buffer: [u8; MAX_DBG_PRINT_LOG_LENGTH + 50] = [0; MAX_DBG_PRINT_LOG_LENGTH + 50];
        
        let mut local_time: LARGE_INTEGER = core::mem::zeroed();
        let mut time_fields: wdk_sys::TIME_FIELDS = core::mem::zeroed();

        ExSystemTimeToLocalTime(&timestamp as *const _ as *mut _, &mut local_time);
        RtlTimeToTimeFields(&mut local_time as *mut _, &mut time_fields);

        let written = core::fmt::write(
            &mut ArrayWriter(&mut write_buffer),
            format_args!(
                "{:02}-{:02} {:02}:{:02}:{:02}.{:03} {}\r\n",
                time_fields.Month,
                time_fields.Day,
                time_fields.Hour,
                time_fields.Minute,
                time_fields.Second,
                time_fields.Milliseconds,
                core::str::from_utf8_unchecked(log_line),
            ),
        );

        if written.is_err() {
            break;
        }

        let len = write_buffer.iter().position(|&b| b == 0).unwrap_or(write_buffer.len());
        
        let mut io_status: IO_STATUS_BLOCK = core::mem::zeroed();

        let status = ZwWriteFile(
            file_handle,
            core::ptr::null_mut(),
            None,
            core::ptr::null_mut(),
            &mut io_status,
            write_buffer.as_mut_ptr() as PVOID,
            len as ULONG,
            core::ptr::null_mut(),
            core::ptr::null_mut(),
        );

        if status != STATUS_SUCCESS {
            break;
        }

        offset += DebugLogEntry::size_with_line(log_line_length);
    }

    if next_offset > 0 {
        let mut io_status: IO_STATUS_BLOCK = core::mem::zeroed();
        let _ = ZwFlushBuffersFile(file_handle, &mut io_status);
    }

    let overflow = old_buffer.overflowed_log_size.load(Ordering::SeqCst);
    if overflow > thread_context.max_overflowed_log_size.load(Ordering::SeqCst) {
        thread_context.max_overflowed_log_size.store(overflow, Ordering::SeqCst);
    }

    drop(old_buffer);
    
    let mut old_buffer = paired_log_buffer.buffers[old_active_idx].lock();
    old_buffer.clear();

    STATUS_SUCCESS
}

pub unsafe fn start_debug_print_callback() -> NTSTATUS {
    if !G_PAIRED_LOG_BUFFER.initialize() {
        return 0xC0000001u32 as i32; 
    }

    let status = DbgSetDebugPrintCallback(Some(debug_print_callback), 1 as BOOLEAN);
    if status != STATUS_SUCCESS {
        G_PAIRED_LOG_BUFFER.buffer_valid.store(false, Ordering::SeqCst);
    }

    status
}

pub unsafe fn stop_debug_print_callback() {
    let _ = DbgSetDebugPrintCallback(None, 0 as BOOLEAN);
    G_PAIRED_LOG_BUFFER.buffer_valid.store(false, Ordering::SeqCst);
}

pub unsafe fn start_debug_logging() -> NTSTATUS {
    start_debug_print_callback()
}

pub unsafe fn stop_debug_logging() {
    stop_debug_print_callback();
}

struct ArrayWriter<'a>(&'a mut [u8]);

impl<'a> core::fmt::Write for ArrayWriter<'a> {
    fn write_str(&mut self, s: &str) -> core::fmt::Result {
        let bytes = s.as_bytes();
        let len = bytes.len().min(self.0.len());
        self.0[..len].copy_from_slice(&bytes[..len]);
        Ok(())
    }
}

pub fn get_max_overflowed_log_size() -> u32 {
    G_THREAD_CONTEXT.max_overflowed_log_size.load(Ordering::SeqCst)
}
