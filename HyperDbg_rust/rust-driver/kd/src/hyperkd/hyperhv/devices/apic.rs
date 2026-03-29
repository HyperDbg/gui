use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, AtomicU64, Ordering};
use spin::Mutex;

use wdk_sys::ntddk::{
    KeGetCurrentProcessorNumberEx,
    KeQueryActiveProcessors,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ApicError {
    NotInitialized,
    AlreadyInitialized,
    InvalidInterrupt,
    InvalidVector,
    NotSupported,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ApicMode {
    Xapic,
    X2apic,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DeliveryMode {
    Fixed = 0,
    LowestPriority = 1,
    SMI = 2,
    NMI = 4,
    Init = 5,
    ExtInt = 7,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DestinationMode {
    Physical = 0,
    Logical = 1,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TriggerMode {
    Edge = 0,
    Level = 1,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DeliveryStatus {
    Idle = 0,
    Pending = 1,
}

#[derive(Debug, Clone, Copy)]
pub struct ApicId {
    pub apic_id: u8,
    pub x2apic_id: u32,
}

impl ApicId {
    pub const fn new(apic_id: u8) -> Self {
        Self {
            apic_id,
            x2apic_id: apic_id as u32,
        }
    }

    pub fn from_x2apic(x2apic_id: u32) -> Self {
        Self {
            apic_id: (x2apic_id & 0xFF) as u8,
            x2apic_id,
        }
    }
}

#[derive(Debug, Clone, Copy)]
pub struct Icr {
    pub vector: u8,
    pub delivery_mode: DeliveryMode,
    pub destination_mode: DestinationMode,
    pub delivery_status: DeliveryStatus,
    pub level: bool,
    pub trigger_mode: TriggerMode,
    pub shorthand: u8,
    pub destination: u32,
}

impl Icr {
    pub fn new() -> Self {
        Self {
            vector: 0,
            delivery_mode: DeliveryMode::Fixed,
            destination_mode: DestinationMode::Physical,
            delivery_status: DeliveryStatus::Idle,
            level: false,
            trigger_mode: TriggerMode::Edge,
            shorthand: 0,
            destination: 0,
        }
    }

    pub fn to_low_dword(&self) -> u32 {
        let mut value: u32 = 0;
        value |= self.vector as u32;
        value |= (self.delivery_mode as u32) << 8;
        value |= (self.destination_mode as u32) << 11;
        value |= (self.delivery_status as u32) << 12;
        if self.level {
            value |= 1 << 14;
        }
        value |= (self.trigger_mode as u32) << 15;
        value |= (self.shorthand as u32) << 18;
        value
    }

    pub fn to_high_dword(&self) -> u32 {
        self.destination
    }
}

#[derive(Debug, Clone, Copy)]
pub struct ApicRegister {
    pub offset: u16,
    pub value: u32,
}

impl ApicRegister {
    pub fn new(offset: u16) -> Self {
        Self {
            offset,
            value: 0,
        }
    }
}

#[derive(Debug)]
pub struct InterruptHandler {
    pub vector: u8,
    pub handler: u64,
    pub installed: bool,
    pub call_count: AtomicU32,
}

impl Clone for InterruptHandler {
    fn clone(&self) -> Self {
        Self {
            vector: self.vector,
            handler: self.handler,
            installed: self.installed,
            call_count: AtomicU32::new(self.call_count.load(Ordering::Acquire)),
        }
    }
}

impl InterruptHandler {
    pub fn new(vector: u8) -> Self {
        Self {
            vector,
            handler: 0,
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

pub struct ApicManager {
    initialized: AtomicBool,
    apic_mode: AtomicU32,
    apic_base: AtomicU64,
    apic_id: Mutex<ApicId>,
    handlers: alloc::collections::BTreeMap<u8, InterruptHandler>,
    processor_count: AtomicU32,
    interrupt_count: AtomicU32,
}

impl ApicManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            apic_mode: AtomicU32::new(ApicMode::Xapic as u32),
            apic_base: AtomicU64::new(0),
            apic_id: Mutex::new(ApicId::new(0)),
            handlers: alloc::collections::BTreeMap::new(),
            processor_count: AtomicU32::new(0),
            interrupt_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), ApicError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(ApicError::AlreadyInitialized);
        }

        unsafe {
            self.detect_apic_mode();
            self.read_apic_base();
            self.read_apic_id();
        }

        self.processor_count.store(unsafe { KeQueryActiveProcessors() } as u32, Ordering::Release);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.handlers.clear();
        self.interrupt_count.store(0, Ordering::Release);
        self.processor_count.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_apic_mode(&self) -> ApicMode {
        match self.apic_mode.load(Ordering::Acquire) {
            0 => ApicMode::Xapic,
            _ => ApicMode::X2apic,
        }
    }

    pub fn get_apic_base(&self) -> u64 {
        self.apic_base.load(Ordering::Acquire)
    }

    pub fn get_apic_id(&self) -> ApicId {
        *self.apic_id.lock()
    }

    pub fn get_processor_count(&self) -> u32 {
        self.processor_count.load(Ordering::Acquire)
    }

    pub fn read_apic_register(&self, offset: u16) -> u32 {
        unsafe {
            self.read_register(offset)
        }
    }

    pub fn write_apic_register(&self, offset: u16, value: u32) {
        unsafe {
            self.write_register(offset, value);
        }
    }

    pub fn send_ipi(&self, icr: &Icr) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        unsafe {
            self.send_interrupt(icr);
        }

        Ok(())
    }

    pub fn send_ipi_to_processor(&self, processor_id: u32, vector: u8) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        let mut icr = Icr::new();
        icr.vector = vector;
        icr.delivery_mode = DeliveryMode::Fixed;
        icr.destination_mode = DestinationMode::Physical;
        icr.destination = processor_id;

        self.send_ipi(&icr)
    }

    pub fn send_ipi_to_all(&self, vector: u8) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        let mut icr = Icr::new();
        icr.vector = vector;
        icr.delivery_mode = DeliveryMode::Fixed;
        icr.destination_mode = DestinationMode::Physical;
        icr.shorthand = 2;

        self.send_ipi(&icr)
    }

    pub fn send_ipi_to_all_excluding_self(&self, vector: u8) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        let mut icr = Icr::new();
        icr.vector = vector;
        icr.delivery_mode = DeliveryMode::Fixed;
        icr.destination_mode = DestinationMode::Physical;
        icr.shorthand = 3;

        self.send_ipi(&icr)
    }

    pub fn install_interrupt_handler(&mut self, vector: u8, handler: u64) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        if vector < 0x20 || vector > 0xFF {
            return Err(ApicError::InvalidVector);
        }

        let mut int_handler = InterruptHandler::new(vector);
        int_handler.handler = handler;
        int_handler.installed = true;

        self.handlers.insert(vector, int_handler);

        Ok(())
    }

    pub fn uninstall_interrupt_handler(&mut self, vector: u8) -> Result<(), ApicError> {
        if !self.is_initialized() {
            return Err(ApicError::NotInitialized);
        }

        self.handlers.remove(&vector);

        Ok(())
    }

    pub fn get_interrupt_handler(&self, vector: u8) -> Option<&InterruptHandler> {
        self.handlers.get(&vector)
    }

    pub fn get_handler_count(&self) -> usize {
        self.handlers.len()
    }

    pub fn get_all_handlers(&self) -> alloc::vec::Vec<&InterruptHandler> {
        self.handlers.values().collect()
    }

    pub fn get_interrupt_count(&self) -> u32 {
        self.interrupt_count.load(Ordering::Acquire)
    }

    pub fn increment_interrupt_count(&self) {
        self.interrupt_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn reset_interrupt_count(&self) {
        self.interrupt_count.store(0, Ordering::Release);
    }

    unsafe fn detect_apic_mode(&mut self) {
        let mut eax: u32 = 1;
        let mut ebx: u32;
        let mut ecx: u32;
        let mut edx: u32;

        core::arch::asm!(
            "push rbx",
            "cpuid",
            "mov {0}, ebx",
            "pop rbx",
            out(reg) ebx,
            inout("eax") eax,
            out("ecx") ecx,
            out("edx") edx,
            options(nomem, nostack),
        );

        let x2apic_supported = (ecx & (1 << 21)) != 0;

        if x2apic_supported {
            let apic_base_msr: u64;
            let apic_base_low: u32;
            let apic_base_high: u32;
            core::arch::asm!(
                "rdmsr",
                inlateout("ecx") 0x1Bu32 => _,
                out("eax") apic_base_low,
                out("edx") apic_base_high,
                options(nomem, nostack),
            );
            apic_base_msr = ((apic_base_high as u64) << 32) | (apic_base_low as u64);

            let x2apic_enabled = (apic_base_msr & (1 << 10)) != 0;

            if x2apic_enabled {
                self.apic_mode.store(ApicMode::X2apic as u32, Ordering::Release);
            } else {
                self.apic_mode.store(ApicMode::Xapic as u32, Ordering::Release);
            }
        } else {
            self.apic_mode.store(ApicMode::Xapic as u32, Ordering::Release);
        }
    }

    unsafe fn read_apic_base(&mut self) {
        let apic_base_msr: u64;
        let apic_base_low: u32;
        let apic_base_high: u32;
        core::arch::asm!(
            "rdmsr",
            inlateout("ecx") 0x1Bu32 => _,
            out("eax") apic_base_low,
            out("edx") apic_base_high,
            options(nomem, nostack),
        );
        apic_base_msr = ((apic_base_high as u64) << 32) | (apic_base_low as u64);

        self.apic_base.store(apic_base_msr & 0xFFFFF000, Ordering::Release);
    }

    unsafe fn read_apic_id(&mut self) {
        let mode = self.get_apic_mode();

        if mode == ApicMode::X2apic {
            let apic_id: u32;
            core::arch::asm!(
                "rdmsr",
                inlateout("ecx") 0x802u32 => _,
                out("eax") apic_id,
                out("edx") _,
                options(nomem, nostack),
            );

            *self.apic_id.lock() = ApicId::from_x2apic(apic_id);
        } else {
            let apic_base = self.get_apic_base();
            let apic_id_ptr = (apic_base + 0x20) as *const u32;
            let apic_id = *apic_id_ptr;

            *self.apic_id.lock() = ApicId::new(((apic_id >> 24) & 0xFF) as u8);
        }
    }

    unsafe fn read_register(&self, offset: u16) -> u32 {
        let mode = self.get_apic_mode();

        if mode == ApicMode::X2apic {
            let value: u32;
            core::arch::asm!(
                "rdmsr",
                inlateout("ecx") (0x800u32 + offset as u32) => _,
                out("eax") value,
                out("edx") _,
                options(nomem, nostack),
            );
            value
        } else {
            let apic_base = self.get_apic_base();
            let register_ptr = (apic_base + offset as u64) as *const u32;
            *register_ptr
        }
    }

    unsafe fn write_register(&self, offset: u16, value: u32) {
        let mode = self.get_apic_mode();

        if mode == ApicMode::X2apic {
            core::arch::asm!(
                "wrmsr",
                in("ecx") 0x800u32 + offset as u32,
                in("eax") value,
                in("edx") 0u32,
                options(nomem, nostack),
            );
        } else {
            let apic_base = self.get_apic_base();
            let register_ptr = (apic_base + offset as u64) as *mut u32;
            *register_ptr = value;
        }
    }

    unsafe fn send_interrupt(&self, icr: &Icr) {
        let mode = self.get_apic_mode();

        if mode == ApicMode::X2apic {
            let icr_low = icr.to_low_dword();
            let icr_high = icr.to_high_dword();

            core::arch::asm!(
                "wrmsr",
                in("ecx") 0x830u32,
                in("eax") icr_low,
                in("edx") icr_high,
                options(nomem, nostack),
            );
        } else {
            let apic_base = self.get_apic_base();

            let icr_low_ptr = (apic_base + 0x300) as *mut u32;
            let icr_high_ptr = (apic_base + 0x310) as *mut u32;

            *icr_high_ptr = icr.to_high_dword();
            *icr_low_ptr = icr.to_low_dword();
        }
    }
}

pub static APIC_MANAGER: Mutex<ApicManager> = Mutex::new(ApicManager::new());

pub fn initialize_apic() -> Result<(), ApicError> {
    let mut manager = APIC_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_apic() {
    let mut manager = APIC_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_apic_initialized() -> bool {
    let manager = APIC_MANAGER.lock();
    manager.is_initialized()
}

pub fn get_apic_mode() -> ApicMode {
    let manager = APIC_MANAGER.lock();
    manager.get_apic_mode()
}

pub fn get_apic_base() -> u64 {
    let manager = APIC_MANAGER.lock();
    manager.get_apic_base()
}

pub fn get_apic_id() -> ApicId {
    let manager = APIC_MANAGER.lock();
    manager.get_apic_id()
}

pub fn get_processor_count() -> u32 {
    let manager = APIC_MANAGER.lock();
    manager.get_processor_count()
}

pub fn read_apic_register(offset: u16) -> u32 {
    let manager = APIC_MANAGER.lock();
    manager.read_apic_register(offset)
}

pub fn write_apic_register(offset: u16, value: u32) {
    let manager = APIC_MANAGER.lock();
    manager.write_apic_register(offset, value)
}

pub fn send_ipi(icr: &Icr) -> Result<(), ApicError> {
    let manager = APIC_MANAGER.lock();
    manager.send_ipi(icr)
}

pub fn send_ipi_to_processor(processor_id: u32, vector: u8) -> Result<(), ApicError> {
    let manager = APIC_MANAGER.lock();
    manager.send_ipi_to_processor(processor_id, vector)
}

pub fn send_ipi_to_all(vector: u8) -> Result<(), ApicError> {
    let manager = APIC_MANAGER.lock();
    manager.send_ipi_to_all(vector)
}

pub fn send_ipi_to_all_excluding_self(vector: u8) -> Result<(), ApicError> {
    let manager = APIC_MANAGER.lock();
    manager.send_ipi_to_all_excluding_self(vector)
}

pub fn install_interrupt_handler(vector: u8, handler: u64) -> Result<(), ApicError> {
    let mut manager = APIC_MANAGER.lock();
    manager.install_interrupt_handler(vector, handler)
}

pub fn uninstall_interrupt_handler(vector: u8) -> Result<(), ApicError> {
    let mut manager = APIC_MANAGER.lock();
    manager.uninstall_interrupt_handler(vector)
}

pub fn get_interrupt_handler(vector: u8) -> Option<u64> {
    let manager = APIC_MANAGER.lock();
    manager.get_interrupt_handler(vector).map(|h| h.handler)
}

pub fn get_handler_count() -> usize {
    let manager = APIC_MANAGER.lock();
    manager.get_handler_count()
}

pub fn get_all_handlers() -> alloc::vec::Vec<u64> {
    let manager = APIC_MANAGER.lock();
    manager.get_all_handlers().iter().map(|h| h.handler).collect()
}

pub fn get_interrupt_count() -> u32 {
    let manager = APIC_MANAGER.lock();
    manager.get_interrupt_count()
}

pub fn increment_interrupt_count() {
    let manager = APIC_MANAGER.lock();
    manager.increment_interrupt_count()
}

pub fn reset_interrupt_count() {
    let manager = APIC_MANAGER.lock();
    manager.reset_interrupt_count()
}

pub fn get_interrupt_call_count(vector: u8) -> Option<u32> {
    let manager = APIC_MANAGER.lock();
    manager.get_interrupt_handler(vector).map(|h| h.get_call_count())
}

pub fn reset_interrupt_call_count(vector: u8) -> Option<()> {
    let manager = APIC_MANAGER.lock();
    manager.get_interrupt_handler(vector).map(|h| h.reset_call_count())
}
