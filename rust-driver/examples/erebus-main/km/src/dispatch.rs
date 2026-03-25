use alloc::format;
use core::{ffi::c_void, ptr::{null_mut, addr_of_mut, addr_of}};

use crate::{
    logger::LogLevel,
    memory::{is_valid_user_memory, ke_read_virtual_memory, ke_write_virtual_memory},
    println,
    process::Process,
};

use shared::{
    ioctl::{EREBUS_IOCTL_READ, EREBUS_IOCTL_WRITE, IOCTL_SEND_DATA, IOCTL_RECEIVE_DATA},
    ipc::Request,
};
use wdk::nt_success;
use wdk_sys::{
    _IO_STACK_LOCATION, DEVICE_OBJECT, IO_NO_INCREMENT, NTSTATUS, PIRP, STATUS_ACCESS_VIOLATION,
    STATUS_BUFFER_ALL_ZEROS, STATUS_BUFFER_TOO_SMALL, STATUS_INVALID_DEVICE_REQUEST,
    STATUS_SUCCESS, STATUS_UNSUCCESSFUL, ntddk::IofCompleteRequest,
    ntddk::RtlCopyMemoryNonTemporal,
};

// =========================================================================
// IoctlBuffer Helper
// =========================================================================

struct IoctlBuffer {
    len: u32,
    buf: *mut c_void,
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
}

impl IoctlBuffer {
    fn new(p_stack_location: *mut _IO_STACK_LOCATION, p_irp: PIRP) -> Self {
        IoctlBuffer {
            len: 0,
            buf: null_mut(),
            p_stack_location,
            p_irp,
        }
    }

    fn get_buf_to_req(&mut self) -> Result<Request, NTSTATUS> {
        self.receive()?;

        if self.len < core::mem::size_of::<Request>() as u32 {
            println!(
                LogLevel::Error,
                "Buffer too small for Request. Expected {} bytes, got {} bytes.",
                core::mem::size_of::<Request>(),
                self.len
            );
            return Err(STATUS_BUFFER_TOO_SMALL);
        }

        let input_buffer =
            unsafe { core::slice::from_raw_parts(self.buf as *const u8, self.len as usize) };
        if input_buffer.is_empty() {
            println!(LogLevel::Error, "Error reading string passed to IOCTL");
            return Err(STATUS_UNSUCCESSFUL);
        }

        let input_buffer_ptr: *const [u8; size_of::<Request>()] =
            input_buffer.as_ptr() as *const [u8; size_of::<Request>()];

        let request = unsafe {
            core::mem::transmute::<[u8; size_of::<Request>()], Request>(*input_buffer_ptr)
        };

        Ok(request)
    }

    fn receive(&mut self) -> Result<(), NTSTATUS> {
        let input_len: u32 = unsafe {
            (*self.p_stack_location)
                .Parameters
                .DeviceIoControl
                .InputBufferLength
        };

        let input_buffer: *mut c_void = unsafe { (*self.p_irp).AssociatedIrp.SystemBuffer };
        if input_buffer.is_null() {
            println!("Input buffer is null.");
            return Err(STATUS_BUFFER_ALL_ZEROS);
        };

        self.len = input_len;
        self.buf = input_buffer;

        Ok(())
    }

    fn send_str(&self, input_str: &str) -> Result<(), NTSTATUS> {
        unsafe { (*self.p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS };

        let response = input_str.as_bytes();
        let response_len = response.len();
        unsafe { (*self.p_irp).IoStatus.Information = response_len as u64 };

        println!(
            LogLevel::Info,
            "Sending a message back to user-land {:?}",
            core::str::from_utf8(response).unwrap()
        );

        unsafe {
            if !(*self.p_irp).AssociatedIrp.SystemBuffer.is_null() {
                RtlCopyMemoryNonTemporal(
                    (*self.p_irp).AssociatedIrp.SystemBuffer,
                    response as *const _ as *mut c_void,
                    response_len as u64,
                );
            } else {
                println!(
                    LogLevel::Error,
                    "Error handling IOCTL, SystemBuffer was null."
                );
                return Err(STATUS_UNSUCCESSFUL);
            }
        }

        Ok(())
    }
}

// =========================================================================
// Dispatch Handlers
// =========================================================================

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn create_close(_device: *mut DEVICE_OBJECT, p_irp: PIRP) -> NTSTATUS {
    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = 0;

        IofCompleteRequest(p_irp, IO_NO_INCREMENT as i8);
    }

    println!("IRP received...");

    STATUS_SUCCESS
}

macro_rules! handle_ioctl_fn {
    ($fn_name:ident, $p_stack_location:expr, $p_irp:expr) => {{
        if let Err(err) = $fn_name($p_stack_location, $p_irp) {
            println!(LogLevel::Error, "Error: {:#x}", err);
            err
        } else {
            STATUS_SUCCESS
        }
    }};
}

#[allow(clippy::cast_possible_truncation)]
pub unsafe extern "C" fn handle_ioctl(_device: *mut DEVICE_OBJECT, pirp: PIRP) -> NTSTATUS {
    unsafe {
        let p_stack_location: *mut _IO_STACK_LOCATION =
            crate::ffi::IoGetCurrentIrpStackLocation(pirp);

        if p_stack_location.is_null() {
            println!("Unable to get stack location for IRP.");
            return STATUS_UNSUCCESSFUL;
        }

        let control_code = (*p_stack_location).Parameters.DeviceIoControl.IoControlCode;
        println!(
            LogLevel::Info,
            "Received an IOCTL code: {:#x}", control_code
        );

        let status: NTSTATUS = match control_code {
            EREBUS_IOCTL_READ => {
                handle_ioctl_fn!(ioctl_handler_read, p_stack_location, pirp)
            }
            EREBUS_IOCTL_WRITE => {
                handle_ioctl_fn!(ioctl_handler_write, p_stack_location, pirp)
            }
            IOCTL_SEND_DATA => {
                handle_ioctl_fn!(ioctl_handler_send_data, p_stack_location, pirp)
            }
            IOCTL_RECEIVE_DATA => {
                handle_ioctl_fn!(ioctl_handler_receive_data, p_stack_location, pirp)
            }
            _ => {
                println!(
                    LogLevel::Error,
                    "Unhandled IOCTL control code: {:#x}", control_code
                );
                STATUS_INVALID_DEVICE_REQUEST
            }
        };

        IofCompleteRequest(pirp, IO_NO_INCREMENT as i8);

        status
    }
}

// =========================================================================
// Specific Handlers
// =========================================================================

fn ioctl_handler_read(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> Result<(), NTSTATUS> {
    let mut ioctl_buffer = IoctlBuffer::new(p_stack_location, p_irp);

    let request = ioctl_buffer.get_buf_to_req()?;
    println!(LogLevel::Info, "Received Request: {:?}", request);

    let Request {
        process_id,
        address,
        buffer,
        size,
    } = request;

    if size == 0 {
        println!(
            LogLevel::Error,
            "Invalid size specified in IOCTL request: {}", size
        );
        return Err(STATUS_UNSUCCESSFUL);
    }

    let Process { process } = Process::by_id(process_id)?;

    println!(
        LogLevel::Success,
        "Resolved process with PID {} and _EPROCESS at {:?}", process_id, process
    );

    if !is_valid_user_memory(address as _, size as _) {
        println!(
            LogLevel::Error,
            "Invalid memory range: {:p}+{:#x}", address, size
        );
        return Err(STATUS_ACCESS_VIOLATION);
    }

    if !is_valid_user_memory(buffer as _, size as _) {
        println!(
            LogLevel::Error,
            "Invalid buffer range: {:p}+{:#x}", buffer, size
        );
        return Err(STATUS_ACCESS_VIOLATION);
    }

    let mut bytes_read = 0;
    let status = unsafe { ke_read_virtual_memory(process, address, buffer, size, &mut bytes_read) };

    if !nt_success(status) {
        println!(
            LogLevel::Error,
            "Error copying VirtualMemory! Error: {:#x}", status
        );
        return Err(STATUS_UNSUCCESSFUL);
    }

    println!(
        LogLevel::Success,
        "Read {} bytes from {:p}", bytes_read, address
    );

    ioctl_buffer.send_str(&format!("Copied {} bytes from {:p}!", bytes_read, address))?;

    Ok(())
}

fn ioctl_handler_write(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> Result<(), NTSTATUS> {
    let mut ioctl_buffer = IoctlBuffer::new(p_stack_location, p_irp);

    let request = ioctl_buffer.get_buf_to_req()?;
    println!(LogLevel::Info, "Received Request: {:?}", request);

    let Request {
        process_id,
        address,
        buffer,
        size,
    } = request;

    if size == 0 {
        println!(
            LogLevel::Error,
            "Invalid size specified in IOCTL request: {}", size
        );
        return Err(STATUS_UNSUCCESSFUL);
    }

    let Process { process } = Process::by_id(process_id)?;

    println!(
        LogLevel::Success,
        "Resolved process with PID {} and _EPROCESS at {:?}", process_id, process
    );

    if !is_valid_user_memory(address as _, size as _) {
        println!(
            LogLevel::Error,
            "Invalid memory range: {:p}+{:#x}", address, size
        );
        return Err(STATUS_ACCESS_VIOLATION);
    }

    if !is_valid_user_memory(buffer as _, size as _) {
        println!(
            LogLevel::Error,
            "Invalid buffer range: {:p}+{:#x}", buffer, size
        );
        return Err(STATUS_ACCESS_VIOLATION);
    }

    let mut bytes_written = 0;
    let status =
        unsafe { ke_write_virtual_memory(process, buffer, address, size, &mut bytes_written) };

    if !nt_success(status) {
        println!(
            LogLevel::Error,
            "Error copying VirtualMemory! Error: {:#x}", status
        );
        return Err(STATUS_UNSUCCESSFUL);
    }

    println!(
        LogLevel::Success,
        "Wrote {} bytes to {:p}", bytes_written, address
    );

    ioctl_buffer.send_str(&format!("Copied {} bytes to {:p}!", bytes_written, address))?;

    Ok(())
}

static mut SYSDATA_BUFFER: [u8; 4096] = [0; 4096];
static mut SYSDATA_LENGTH: u32 = 0;

fn ioctl_handler_send_data(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> Result<(), NTSTATUS> {
    let input_len: u32 = unsafe {
        (*p_stack_location)
            .Parameters
            .DeviceIoControl
            .InputBufferLength
    };

    let input_buffer: *mut c_void = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };
    if input_buffer.is_null() || input_len == 0 {
        println!(LogLevel::Error, "Input buffer is null or empty");
        return Err(STATUS_UNSUCCESSFUL);
    }

    let copy_length = input_len.min(4096);
    unsafe {
        let buffer_ptr = addr_of_mut!(SYSDATA_BUFFER) as *mut u8;
        core::ptr::copy_nonoverlapping(
            input_buffer as *const u8,
            buffer_ptr,
            copy_length as usize,
        );
        *addr_of_mut!(SYSDATA_LENGTH) = copy_length;
    }

    println!(
        LogLevel::Success,
        "IOCTL_SEND_DATA: Received {} bytes from user",
        copy_length
    );

    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = copy_length as u64;
    }

    Ok(())
}

fn ioctl_handler_receive_data(
    p_stack_location: *mut _IO_STACK_LOCATION,
    p_irp: PIRP,
) -> Result<(), NTSTATUS> {
    let output_len: u32 = unsafe {
        (*p_stack_location)
            .Parameters
            .DeviceIoControl
            .OutputBufferLength
    };

    let output_buffer: *mut c_void = unsafe { (*p_irp).AssociatedIrp.SystemBuffer };
    if output_buffer.is_null() || output_len == 0 {
        println!(LogLevel::Error, "Output buffer is null or empty");
        return Err(STATUS_UNSUCCESSFUL);
    }

    let prefix = b"received data by user ";
    let prefix_len = prefix.len();
    let data_len = unsafe { *addr_of!(SYSDATA_LENGTH) as usize };
    let total_len = prefix_len + data_len;

    let bytes_to_copy = total_len.min(output_len as usize);

    unsafe {
        if bytes_to_copy > 0 {
            let prefix_copy_len = prefix_len.min(bytes_to_copy);
            core::ptr::copy_nonoverlapping(
                prefix.as_ptr(),
                output_buffer as *mut u8,
                prefix_copy_len,
            );

            if bytes_to_copy > prefix_len {
                let data_copy_len = bytes_to_copy - prefix_len;
                let buffer_ptr = addr_of!(SYSDATA_BUFFER) as *const u8;
                core::ptr::copy_nonoverlapping(
                    buffer_ptr,
                    (output_buffer as *mut u8).add(prefix_len),
                    data_copy_len,
                );
            }
        }
    }

    println!(
        LogLevel::Success,
        "IOCTL_RECEIVE_DATA: Returning {} bytes to user",
        bytes_to_copy
    );

    unsafe {
        (*p_irp).IoStatus.__bindgen_anon_1.Status = STATUS_SUCCESS;
        (*p_irp).IoStatus.Information = bytes_to_copy as u64;
    }

    Ok(())
}
