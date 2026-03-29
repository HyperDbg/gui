use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use wdk_sys::ntddk::{
    KeGetCurrentIrql,
    KeGetCurrentProcessorNumberEx,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum IdtError {
    NotInitialized,
    AlreadyInitialized,
    InvalidInterrupt,
    InvalidHandler,
    AllocationFailed,
    InvalidParameter,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum InterruptType {
    InterruptGate = 0xE,
    TrapGate = 0xF,
    TaskGate = 0x5,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DescriptorPrivilege {
    Ring0 = 0,
    Ring1 = 1,
    Ring2 = 2,
    Ring3 = 3,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum GateSize {
    Size16 = 0,
    Size32 = 1,
}

#[repr(C, packed)]
#[derive(Debug, Clone, Copy)]
pub struct IdtEntry {
    pub offset_low: u16,
    pub selector: u16,
    pub ist: u8,
    pub type_attr: u8,
    pub offset_mid: u16,
    pub offset_high: u32,
    pub reserved: u32,
}

impl IdtEntry {
    pub fn new() -> Self {
        Self {
            offset_low: 0,
            selector: 0,
            ist: 0,
            type_attr: 0,
            offset_mid: 0,
            offset_high: 0,
            reserved: 0,
        }
    }

    pub fn set_handler(&mut self, handler: u64, selector: u16, ist: u8, gate_type: InterruptType, privilege: DescriptorPrivilege, size: GateSize) {
        self.offset_low = (handler & 0xFFFF) as u16;
        self.offset_mid = ((handler >> 16) & 0xFFFF) as u16;
        self.offset_high = ((handler >> 32) & 0xFFFFFFFF) as u32;
        self.selector = selector;
        self.ist = ist;
        self.type_attr = ((gate_type as u8) & 0x0F)
            | ((privilege as u8) << 5)
            | ((size as u8) << 7)
            | 0x80;
    }

    pub fn get_handler(&self) -> u64 {
        ((self.offset_high as u64) << 32)
            | ((self.offset_mid as u64) << 16)
            | (self.offset_low as u64)
    }

    pub fn is_present(&self) -> bool {
        (self.type_attr & 0x80) != 0
    }

    pub fn set_present(&mut self, present: bool) {
        if present {
            self.type_attr |= 0x80;
        } else {
            self.type_attr &= !0x80;
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct IdtPointer {
    pub limit: u16,
    pub base: u64,
}

impl IdtPointer {
    pub fn new(base: u64, limit: u16) -> Self {
        Self { limit, base }
    }
}

#[derive(Debug)]
pub struct InterruptHandler {
    pub interrupt_number: u8,
    pub handler: u64,
    pub original_handler: u64,
    pub installed: bool,
    pub call_count: AtomicU32,
}

impl Clone for InterruptHandler {
    fn clone(&self) -> Self {
        Self {
            interrupt_number: self.interrupt_number,
            handler: self.handler,
            original_handler: self.original_handler,
            installed: self.installed,
            call_count: AtomicU32::new(self.call_count.load(Ordering::Acquire)),
        }
    }
}

impl InterruptHandler {
    pub fn new(interrupt_number: u8) -> Self {
        Self {
            interrupt_number,
            handler: 0,
            original_handler: 0,
            installed: false,
            call_count: AtomicU32::new(0),
        }
    }

    pub fn increment_call_count(&self) {
        self.call_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn get_call_count(&self) -> u32 {
        self.call_count.load(Ordering::Acquire)
    }

    pub fn reset_call_count(&self) {
        self.call_count.store(0, Ordering::Release);
    }
}

pub struct IdtManager {
    initialized: AtomicBool,
    idt_base: u64,
    idt_entries: alloc::vec::Vec<IdtEntry>,
    handlers: alloc::collections::BTreeMap<u8, InterruptHandler>,
    original_idt: alloc::vec::Vec<IdtEntry>,
    total_handlers: AtomicU32,
    active_handlers: AtomicU32,
}

impl IdtManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            idt_base: 0,
            idt_entries: alloc::vec::Vec::new(),
            handlers: alloc::collections::BTreeMap::new(),
            original_idt: alloc::vec::Vec::new(),
            total_handlers: AtomicU32::new(0),
            active_handlers: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), IdtError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(IdtError::AlreadyInitialized);
        }

        self.idt_entries.resize(256, IdtEntry::new());
        self.original_idt.resize(256, IdtEntry::new());

        unsafe {
            self.load_current_idt();
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        unsafe {
            self.restore_original_idt();
        }
        self.handlers.clear();
        self.idt_entries.clear();
        self.original_idt.clear();
        self.total_handlers.store(0, Ordering::Release);
        self.active_handlers.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn install_handler(
        &mut self,
        interrupt_number: u8,
        handler: u64,
        selector: u16,
        ist: u8,
        gate_type: InterruptType,
        privilege: DescriptorPrivilege,
        size: GateSize,
    ) -> Result<(), IdtError> {
        if !self.is_initialized() {
            return Err(IdtError::NotInitialized);
        }

        if handler == 0 {
            return Err(IdtError::InvalidHandler);
        }

        let interrupt_num = interrupt_number as usize;
        let original = self.original_idt[interrupt_num];
        let entry = &mut self.idt_entries[interrupt_num];
        self.original_idt[interrupt_num] = *entry;

        entry.set_handler(handler, selector, ist, gate_type, privilege, size);

        let updated_entry = *entry;

        let mut int_handler = InterruptHandler::new(interrupt_number);
        int_handler.handler = handler;
        int_handler.original_handler = original.get_handler();
        int_handler.installed = true;

        self.handlers.insert(interrupt_number, int_handler);
        self.total_handlers.fetch_add(1, Ordering::AcqRel);
        self.active_handlers.fetch_add(1, Ordering::AcqRel);

        unsafe {
            self.update_idt_entry(interrupt_number, &updated_entry);
        }

        Ok(())
    }

    pub fn uninstall_handler(&mut self, interrupt_number: u8) -> Result<(), IdtError> {
        if !self.is_initialized() {
            return Err(IdtError::NotInitialized);
        }

        if let Some(handler) = self.handlers.remove(&interrupt_number) {
            if handler.installed {
                let original_entry = self.original_idt[interrupt_number as usize];
                self.idt_entries[interrupt_number as usize] = original_entry;

                unsafe {
                    self.update_idt_entry(interrupt_number, &original_entry);
                }

                self.active_handlers.fetch_sub(1, Ordering::AcqRel);
            }
        }

        Ok(())
    }

    pub fn get_handler(&self, interrupt_number: u8) -> Option<&InterruptHandler> {
        self.handlers.get(&interrupt_number)
    }

    pub fn get_handler_count(&self) -> u32 {
        self.total_handlers.load(Ordering::Acquire)
    }

    pub fn get_active_handler_count(&self) -> u32 {
        self.active_handlers.load(Ordering::Acquire)
    }

    pub fn get_idt_entry(&self, interrupt_number: u8) -> Option<IdtEntry> {
        Some(self.idt_entries[interrupt_number as usize])
    }

    pub fn get_all_handlers(&self) -> alloc::vec::Vec<&InterruptHandler> {
        self.handlers.values().collect()
    }

    pub fn get_idt_base(&self) -> u64 {
        self.idt_base
    }

    unsafe fn load_current_idt(&mut self) {
        let mut idt_ptr: u64 = 0;
        core::arch::asm!(
            "sidt [{}]",
            in(reg) &mut idt_ptr,
            options(nomem, nostack),
        );

        self.idt_base = idt_ptr & 0xFFFFFFFFFFFFF000;

        let idt_entries_ptr = self.idt_base as *const IdtEntry;
        for i in 0..256 {
            self.idt_entries[i] = *idt_entries_ptr.add(i);
            self.original_idt[i] = *idt_entries_ptr.add(i);
        }
    }

    unsafe fn update_idt_entry(&self, interrupt_number: u8, entry: &IdtEntry) {
        let idt_entries_ptr = self.idt_base as *mut IdtEntry;
        *idt_entries_ptr.add(interrupt_number as usize) = *entry;
    }

    unsafe fn restore_original_idt(&mut self) {
        for i in 0..256 {
            let original_entry = self.original_idt[i];
            self.idt_entries[i] = original_entry;

            let idt_entries_ptr = self.idt_base as *mut IdtEntry;
            *idt_entries_ptr.add(i) = original_entry;
        }
    }
}

pub static IDT_MANAGER: Mutex<IdtManager> = Mutex::new(IdtManager::new());

pub fn initialize_idt() -> Result<(), IdtError> {
    let mut manager = IDT_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_idt() {
    let mut manager = IDT_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_idt_initialized() -> bool {
    let manager = IDT_MANAGER.lock();
    manager.is_initialized()
}

pub fn install_interrupt_handler(
    interrupt_number: u8,
    handler: u64,
    selector: u16,
    ist: u8,
    gate_type: InterruptType,
    privilege: DescriptorPrivilege,
    size: GateSize,
) -> Result<(), IdtError> {
    let mut manager = IDT_MANAGER.lock();
    manager.install_handler(interrupt_number, handler, selector, ist, gate_type, privilege, size)
}

pub fn uninstall_interrupt_handler(interrupt_number: u8) -> Result<(), IdtError> {
    let mut manager = IDT_MANAGER.lock();
    manager.uninstall_handler(interrupt_number)
}

pub fn get_interrupt_handler(interrupt_number: u8) -> Option<u64> {
    let manager = IDT_MANAGER.lock();
    manager.get_handler(interrupt_number).map(|h| h.handler)
}

pub fn get_handler_count() -> u32 {
    let manager = IDT_MANAGER.lock();
    manager.get_handler_count()
}

pub fn get_active_handler_count() -> u32 {
    let manager = IDT_MANAGER.lock();
    manager.get_active_handler_count()
}

pub fn get_idt_entry(interrupt_number: u8) -> Option<IdtEntry> {
    let manager = IDT_MANAGER.lock();
    manager.get_idt_entry(interrupt_number)
}

pub fn get_all_handlers() -> alloc::vec::Vec<u64> {
    let manager = IDT_MANAGER.lock();
    manager.get_all_handlers().iter().map(|h| h.handler).collect()
}

pub fn get_idt_base() -> u64 {
    let manager = IDT_MANAGER.lock();
    manager.get_idt_base()
}

pub fn get_interrupt_call_count(interrupt_number: u8) -> Option<u32> {
    let manager = IDT_MANAGER.lock();
    manager.get_handler(interrupt_number).map(|h| h.get_call_count())
}

pub fn reset_interrupt_call_count(interrupt_number: u8) -> Option<()> {
    let manager = IDT_MANAGER.lock();
    manager.get_handler(interrupt_number).map(|h| h.reset_call_count())
}
