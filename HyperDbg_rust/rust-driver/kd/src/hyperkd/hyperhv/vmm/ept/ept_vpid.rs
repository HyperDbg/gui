use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn ExAllocatePoolWithTag(pool_type: u32, number_of_bytes: usize, tag: u32) -> *mut u8;
    fn ExFreePool(pool: *mut u8);
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EptError {
    NotInitialized,
    AlreadyInitialized,
    InvalidEptPointer,
    InvalidContext,
    NotSupported,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VpidError {
    NotInitialized,
    AlreadyInitialized,
    InvalidVpid,
    InvalidContext,
    NotSupported,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum InveptType {
    SingleContext = 1,
    GlobalContext = 2,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum InvvpidType {
    IndividualAddress = 0,
    SingleContext = 1,
    AllContexts = 2,
    SingleContextRetainingGlobals = 3,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct InveptDescriptor {
    pub ept_pointer: u64,
    pub reserved: u64,
}

impl InveptDescriptor {
    pub fn new(ept_pointer: u64) -> Self {
        Self {
            ept_pointer,
            reserved: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct InvvpidDescriptor {
    pub vpid: u16,
    pub padding: u16,
    pub linear_address: u64,
}

impl InvvpidDescriptor {
    pub fn new(vpid: u16, linear_address: u64) -> Self {
        Self {
            vpid,
            padding: 0,
            linear_address,
        }
    }
}

#[derive(Debug, Clone, Copy)]
pub struct EptPointer {
    pub eptp: u64,
    pub memory_type: u8,
    pub page_walk_length: u8,
    pub dirty_flag_enabled: bool,
    pub accessed_and_dirty_flag_enabled: bool,
}

impl EptPointer {
    pub fn new(eptp: u64) -> Self {
        Self {
            eptp,
            memory_type: (eptp & 0x38) as u8,
            page_walk_length: ((eptp >> 3) & 0x7) as u8,
            dirty_flag_enabled: (eptp & (1 << 6)) != 0,
            accessed_and_dirty_flag_enabled: (eptp & (1 << 21)) != 0,
        }
    }

    pub fn create(memory_type: u8, page_walk_length: u8, dirty_flag: bool, accessed_dirty_flag: bool) -> u64 {
        let mut eptp: u64 = 0;
        eptp |= (memory_type as u64) & 0x38;
        eptp |= ((page_walk_length as u64) & 0x7) << 3;
        if dirty_flag {
            eptp |= 1 << 6;
        }
        if accessed_dirty_flag {
            eptp |= 1 << 21;
        }
        eptp
    }
}

#[derive(Debug)]
pub struct VpidContext {
    pub vpid: u16,
    pub allocated: bool,
    pub invvpid_count: AtomicU32,
}

impl Clone for VpidContext {
    fn clone(&self) -> Self {
        Self {
            vpid: self.vpid,
            allocated: self.allocated,
            invvpid_count: AtomicU32::new(self.invvpid_count.load(Ordering::Acquire)),
        }
    }
}

impl VpidContext {
    pub fn new(vpid: u16) -> Self {
        Self {
            vpid,
            allocated: false,
            invvpid_count: AtomicU32::new(0),
        }
    }

    pub fn increment_invvpid_count(&self) {
        self.invvpid_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn get_invvpid_count(&self) -> u32 {
        self.invvpid_count.load(Ordering::Acquire)
    }

    pub fn reset_invvpid_count(&self) {
        self.invvpid_count.store(0, Ordering::Release);
    }
}

pub struct EptManager {
    initialized: AtomicBool,
    invept_supported: AtomicBool,
    ept_pointers: alloc::collections::BTreeMap<u64, EptPointer>,
    invept_count: AtomicU32,
}

impl EptManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            invept_supported: AtomicBool::new(false),
            ept_pointers: alloc::collections::BTreeMap::new(),
            invept_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), EptError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(EptError::AlreadyInitialized);
        }

        let supported = unsafe { self.check_invept_support() };
        self.invept_supported.store(supported, Ordering::Release);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.ept_pointers.clear();
        self.invept_count.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_invept_supported(&self) -> bool {
        self.invept_supported.load(Ordering::Acquire)
    }

    pub fn register_ept_pointer(&mut self, eptp: u64) -> Result<(), EptError> {
        if !self.is_initialized() {
            return Err(EptError::NotInitialized);
        }

        if eptp == 0 {
            return Err(EptError::InvalidEptPointer);
        }

        let ept_pointer = EptPointer::new(eptp);
        self.ept_pointers.insert(eptp, ept_pointer);

        Ok(())
    }

    pub fn unregister_ept_pointer(&mut self, eptp: u64) -> Result<(), EptError> {
        if !self.is_initialized() {
            return Err(EptError::NotInitialized);
        }

        self.ept_pointers.remove(&eptp);

        Ok(())
    }

    pub fn invept(&self, invept_type: InveptType, descriptor: &InveptDescriptor) -> Result<(), EptError> {
        if !self.is_initialized() {
            return Err(EptError::NotInitialized);
        }

        if !self.is_invept_supported() {
            return Err(EptError::NotSupported);
        }

        unsafe {
            self.execute_invept(invept_type, descriptor);
        }

        self.invept_count.fetch_add(1, Ordering::AcqRel);
        Ok(())
    }

    pub fn invept_single_context(&self, eptp: u64) -> Result<(), EptError> {
        if !self.is_initialized() {
            return Err(EptError::NotInitialized);
        }

        let descriptor = InveptDescriptor::new(eptp);
        self.invept(InveptType::SingleContext, &descriptor)
    }

    pub fn invept_global(&self) -> Result<(), EptError> {
        if !self.is_initialized() {
            return Err(EptError::NotInitialized);
        }

        let descriptor = InveptDescriptor::new(0);
        self.invept(InveptType::GlobalContext, &descriptor)
    }

    pub fn get_ept_pointer(&self, eptp: u64) -> Option<EptPointer> {
        self.ept_pointers.get(&eptp).copied()
    }

    pub fn get_all_ept_pointers(&self) -> alloc::vec::Vec<EptPointer> {
        self.ept_pointers.values().copied().collect()
    }

    pub fn get_ept_pointer_count(&self) -> usize {
        self.ept_pointers.len()
    }

    pub fn get_invept_count(&self) -> u32 {
        self.invept_count.load(Ordering::Acquire)
    }

    unsafe fn check_invept_support(&self) -> bool {
        let vmx_ept_vpid_cap: u64;
        let vmx_ept_cap_low: u32;
        let vmx_ept_cap_high: u32;
        core::arch::asm!(
            "rdmsr",
            inlateout("ecx") 0x48Cu32 => _,
            out("eax") vmx_ept_cap_low,
            out("edx") vmx_ept_cap_high,
            options(nomem, nostack),
        );
        vmx_ept_vpid_cap = ((vmx_ept_cap_high as u64) << 32) | (vmx_ept_cap_low as u64);

        (vmx_ept_vpid_cap & (1 << 20)) != 0
    }

    unsafe fn execute_invept(&self, invept_type: InveptType, descriptor: &InveptDescriptor) {
        let mut descriptor_value = [0u64; 2];
        descriptor_value[0] = descriptor.ept_pointer;
        descriptor_value[1] = descriptor.reserved;

        core::arch::asm!(
            "invept [{0}], {1}",
            in(reg) descriptor_value.as_ptr(),
            in(reg) invept_type as u64,
            options(nomem, nostack),
        );
    }
}

pub struct VpidManager {
    initialized: AtomicBool,
    invvpid_supported: AtomicBool,
    vpid_contexts: alloc::collections::BTreeMap<u16, VpidContext>,
    next_vpid: AtomicU32,
    invvpid_count: AtomicU32,
}

impl VpidManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            invvpid_supported: AtomicBool::new(false),
            vpid_contexts: alloc::collections::BTreeMap::new(),
            next_vpid: AtomicU32::new(1),
            invvpid_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), VpidError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(VpidError::AlreadyInitialized);
        }

        let supported = unsafe { self.check_invvpid_support() };
        self.invvpid_supported.store(supported, Ordering::Release);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.vpid_contexts.clear();
        self.next_vpid.store(1, Ordering::Release);
        self.invvpid_count.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_invvpid_supported(&self) -> bool {
        self.invvpid_supported.load(Ordering::Acquire)
    }

    pub fn allocate_vpid(&mut self) -> Result<u16, VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        let vpid = self.next_vpid.fetch_add(1, Ordering::AcqRel) as u16;
        let mut context = VpidContext::new(vpid);
        context.allocated = true;

        self.vpid_contexts.insert(vpid, context);

        Ok(vpid)
    }

    pub fn free_vpid(&mut self, vpid: u16) -> Result<(), VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        self.vpid_contexts.remove(&vpid);

        Ok(())
    }

    pub fn invvpid(&self, invvpid_type: InvvpidType, descriptor: &InvvpidDescriptor) -> Result<(), VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        if !self.is_invvpid_supported() {
            return Err(VpidError::NotSupported);
        }

        unsafe {
            self.execute_invvpid(invvpid_type, descriptor);
        }

        self.invvpid_count.fetch_add(1, Ordering::AcqRel);

        if let Some(context) = self.vpid_contexts.get(&descriptor.vpid) {
            context.increment_invvpid_count();
        }

        Ok(())
    }

    pub fn invvpid_individual_address(&self, vpid: u16, linear_address: u64) -> Result<(), VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        let descriptor = InvvpidDescriptor::new(vpid, linear_address);
        self.invvpid(InvvpidType::IndividualAddress, &descriptor)
    }

    pub fn invvpid_single_context(&self, vpid: u16) -> Result<(), VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        let descriptor = InvvpidDescriptor::new(vpid, 0);
        self.invvpid(InvvpidType::SingleContext, &descriptor)
    }

    pub fn invvpid_all_contexts(&self) -> Result<(), VpidError> {
        if !self.is_initialized() {
            return Err(VpidError::NotInitialized);
        }

        let descriptor = InvvpidDescriptor::new(0, 0);
        self.invvpid(InvvpidType::AllContexts, &descriptor)
    }

    pub fn get_vpid_context(&self, vpid: u16) -> Option<&VpidContext> {
        self.vpid_contexts.get(&vpid)
    }

    pub fn get_all_vpid_contexts(&self) -> alloc::vec::Vec<&VpidContext> {
        self.vpid_contexts.values().collect()
    }

    pub fn get_vpid_count(&self) -> usize {
        self.vpid_contexts.len()
    }

    pub fn get_invvpid_count(&self) -> u32 {
        self.invvpid_count.load(Ordering::Acquire)
    }

    unsafe fn check_invvpid_support(&self) -> bool {
        let vmx_ept_vpid_cap_low: u32;
        let vmx_ept_vpid_cap_high: u32;
        core::arch::asm!(
            "rdmsr",
            inlateout("ecx") 0x48Cu32 => _,
            out("eax") vmx_ept_vpid_cap_low,
            out("edx") vmx_ept_vpid_cap_high,
            options(nomem, nostack),
        );

        let vmx_ept_vpid_cap = ((vmx_ept_vpid_cap_high as u64) << 32) | (vmx_ept_vpid_cap_low as u64);

        (vmx_ept_vpid_cap & (1 << 21)) != 0
    }

    unsafe fn execute_invvpid(&self, invvpid_type: InvvpidType, descriptor: &InvvpidDescriptor) {
        let mut descriptor_value = [0u64; 2];
        descriptor_value[0] = descriptor.vpid as u64;
        descriptor_value[1] = descriptor.linear_address;

        core::arch::asm!(
            "invvpid [{0}], {1}",
            in(reg) descriptor_value.as_ptr(),
            in(reg) invvpid_type as u64,
            options(nomem, nostack),
        );
    }
}

pub static EPT_MANAGER: Mutex<EptManager> = Mutex::new(EptManager::new());
pub static VPID_MANAGER: Mutex<VpidManager> = Mutex::new(VpidManager::new());

pub fn initialize_ept() -> Result<(), EptError> {
    let mut manager = EPT_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_ept() {
    let mut manager = EPT_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_ept_initialized() -> bool {
    let manager = EPT_MANAGER.lock();
    manager.is_initialized()
}

pub fn is_invept_supported() -> bool {
    let manager = EPT_MANAGER.lock();
    manager.is_invept_supported()
}

pub fn register_ept_pointer(eptp: u64) -> Result<(), EptError> {
    let mut manager = EPT_MANAGER.lock();
    manager.register_ept_pointer(eptp)
}

pub fn unregister_ept_pointer(eptp: u64) -> Result<(), EptError> {
    let mut manager = EPT_MANAGER.lock();
    manager.unregister_ept_pointer(eptp)
}

pub fn invept(invept_type: InveptType, descriptor: &InveptDescriptor) -> Result<(), EptError> {
    let manager = EPT_MANAGER.lock();
    manager.invept(invept_type, descriptor)
}

pub fn invept_single_context(eptp: u64) -> Result<(), EptError> {
    let manager = EPT_MANAGER.lock();
    manager.invept_single_context(eptp)
}

pub fn invept_global() -> Result<(), EptError> {
    let manager = EPT_MANAGER.lock();
    manager.invept_global()
}

pub fn initialize_vpid() -> Result<(), VpidError> {
    let mut manager = VPID_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_vpid() {
    let mut manager = VPID_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_vpid_initialized() -> bool {
    let manager = VPID_MANAGER.lock();
    manager.is_initialized()
}

pub fn is_invvpid_supported() -> bool {
    let manager = VPID_MANAGER.lock();
    manager.is_invvpid_supported()
}

pub fn allocate_vpid() -> Result<u16, VpidError> {
    let mut manager = VPID_MANAGER.lock();
    manager.allocate_vpid()
}

pub fn free_vpid(vpid: u16) -> Result<(), VpidError> {
    let mut manager = VPID_MANAGER.lock();
    manager.free_vpid(vpid)
}

pub fn invvpid(invvpid_type: InvvpidType, descriptor: &InvvpidDescriptor) -> Result<(), VpidError> {
    let manager = VPID_MANAGER.lock();
    manager.invvpid(invvpid_type, descriptor)
}

pub fn invvpid_individual_address(vpid: u16, linear_address: u64) -> Result<(), VpidError> {
    let manager = VPID_MANAGER.lock();
    manager.invvpid_individual_address(vpid, linear_address)
}

pub fn invvpid_single_context(vpid: u16) -> Result<(), VpidError> {
    let manager = VPID_MANAGER.lock();
    manager.invvpid_single_context(vpid)
}

pub fn invvpid_all_contexts() -> Result<(), VpidError> {
    let manager = VPID_MANAGER.lock();
    manager.invvpid_all_contexts()
}

pub fn get_vpid_context(vpid: u16) -> Option<VpidContext> {
    let manager = VPID_MANAGER.lock();
    manager.get_vpid_context(vpid).cloned()
}

pub fn get_ept_pointer(eptp: u64) -> Option<EptPointer> {
    let manager = EPT_MANAGER.lock();
    manager.get_ept_pointer(eptp)
}

pub fn get_ept_pointer_count() -> usize {
    let manager = EPT_MANAGER.lock();
    manager.get_ept_pointer_count()
}

pub fn get_vpid_count() -> usize {
    let manager = VPID_MANAGER.lock();
    manager.get_vpid_count()
}

pub fn get_invept_count() -> u32 {
    let manager = EPT_MANAGER.lock();
    manager.get_invept_count()
}

pub fn get_invvpid_count() -> u32 {
    let manager = VPID_MANAGER.lock();
    manager.get_invvpid_count()
}
