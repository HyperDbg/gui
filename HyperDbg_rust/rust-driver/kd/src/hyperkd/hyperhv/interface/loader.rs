use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn MmGetSystemRoutineAddress(routine_name: *const u8) -> *mut u8;
    fn ExAllocatePoolWithTag(pool_type: u32, number_of_bytes: usize, tag: u32) -> *mut u8;
    fn ExFreePool(pool: *mut u8);
    fn RtlInitUnicodeString(destination_string: *mut u16, source_string: *const u16);
    fn RtlUnicodeStringToAnsiString(destination_string: *mut u8, source_string: *const u16) -> u32;
    fn IoAllocateMdl(virtual_address: *mut u8, length: u32, secondary_buffer: bool, charge_quota: bool, chain_flags: u32) -> *mut u8;
    fn IoFreeMdl(mdl: *mut u8);
    fn MmGetMdlVirtualAddress(mdl: *const u8) -> *mut u8;
    fn MmGetMdlPfnArray(mdl: *const u8) -> *mut u64;
    fn MmMapLockedPages(mdl: *mut u8, access_mode: u32, cache_type: u32) -> *mut u8;
    fn MmUnmapLockedPages(base_address: *mut u8, mdl: *mut u8);
    fn KeInitializeSpinLock(spin_lock: *mut u8);
    fn KeAcquireSpinLock(spin_lock: *mut u8);
    fn KeReleaseSpinLock(spin_lock: *mut u8);
    fn KeInitializeEvent(event: *mut u8, event_type: u32, initial_state: bool);
    fn KeSetEvent(event: *mut u8, increment: i32, wait: bool) -> i32;
    fn KeWaitForSingleObject(object: *mut u8, wait_reason: u32, wait_mode: u32, alertable: bool, timeout: *const i64) -> u32;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LoaderError {
    NotInitialized,
    AlreadyInitialized,
    InitializationFailed,
    LoadFailed,
    UnloadFailed,
    InvalidParameter,
    SymbolNotFound,
    MemoryAllocationFailed,
    InvalidAddress,
    AccessDenied,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LoaderState {
    Uninitialized,
    Loading,
    Loaded,
    Unloading,
    Error,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PoolType {
    NonPagedPool,
    PagedPool,
    NonPagedPoolCacheAligned,
    PagedPoolCacheAligned,
}

#[derive(Debug, Clone)]
pub struct ModuleInfo {
    pub base_address: u64,
    pub size: u64,
    pub name: alloc::string::String,
    pub path: alloc::string::String,
    pub checksum: u32,
    pub timestamp: u32,
}

impl ModuleInfo {
    pub fn new() -> Self {
        Self {
            base_address: 0,
            size: 0,
            name: alloc::string::String::new(),
            path: alloc::string::String::new(),
            checksum: 0,
            timestamp: 0,
        }
    }
}

#[derive(Debug)]
pub struct LoadContext {
    pub module_info: ModuleInfo,
    pub loaded: AtomicBool,
    pub load_time: u64,
}

impl Clone for LoadContext {
    fn clone(&self) -> Self {
        Self {
            module_info: self.module_info.clone(),
            loaded: AtomicBool::new(self.loaded.load(Ordering::Acquire)),
            load_time: self.load_time,
        }
    }
}

impl LoadContext {
    pub fn new() -> Self {
        Self {
            module_info: ModuleInfo::new(),
            loaded: AtomicBool::new(false),
            load_time: 0,
        }
    }
}

pub struct LoaderManager {
    initialized: AtomicBool,
    state: AtomicU32,
    loaded_modules: alloc::collections::BTreeMap<u64, ModuleInfo>,
    current_context: Mutex<Option<Box<LoadContext>>>,
    total_modules: AtomicU32,
    load_count: AtomicU32,
    unload_count: AtomicU32,
}

impl LoaderManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            state: AtomicU32::new(LoaderState::Uninitialized as u32),
            loaded_modules: alloc::collections::BTreeMap::new(),
            current_context: Mutex::new(None),
            total_modules: AtomicU32::new(0),
            load_count: AtomicU32::new(0),
            unload_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), LoaderError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(LoaderError::AlreadyInitialized);
        }

        self.loaded_modules.clear();
        self.total_modules.store(0, Ordering::Release);
        self.load_count.store(0, Ordering::Release);
        self.unload_count.store(0, Ordering::Release);
        self.state.store(LoaderState::Loaded as u32, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.loaded_modules.clear();
        *self.current_context.lock() = None;
        self.state.store(LoaderState::Uninitialized as u32, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_state(&self) -> LoaderState {
        match self.state.load(Ordering::Acquire) {
            0 => LoaderState::Uninitialized,
            1 => LoaderState::Loading,
            2 => LoaderState::Loaded,
            3 => LoaderState::Unloading,
            _ => LoaderState::Error,
        }
    }

    pub fn set_state(&self, state: LoaderState) {
        self.state.store(state as u32, Ordering::Release);
    }

    pub fn load_module(&mut self, module_info: ModuleInfo) -> Result<(), LoaderError> {
        if !self.is_initialized() {
            return Err(LoaderError::NotInitialized);
        }

        if module_info.base_address == 0 || module_info.size == 0 {
            return Err(LoaderError::InvalidParameter);
        }

        self.set_state(LoaderState::Loading);

        let mut context = Box::new(LoadContext::new());
        context.module_info = module_info.clone();
        context.load_time = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        self.loaded_modules.insert(module_info.base_address, module_info);
        self.total_modules.fetch_add(1, Ordering::AcqRel);
        self.load_count.fetch_add(1, Ordering::AcqRel);

        context.loaded.store(true, Ordering::Release);
        *self.current_context.lock() = Some(context);

        self.set_state(LoaderState::Loaded);

        Ok(())
    }

    pub fn unload_module(&mut self, base_address: u64) -> Result<(), LoaderError> {
        if !self.is_initialized() {
            return Err(LoaderError::NotInitialized);
        }

        self.set_state(LoaderState::Unloading);

        if self.loaded_modules.remove(&base_address).is_none() {
            self.set_state(LoaderState::Error);
            return Err(LoaderError::InvalidAddress);
        }

        self.total_modules.fetch_sub(1, Ordering::AcqRel);
        self.unload_count.fetch_add(1, Ordering::AcqRel);

        *self.current_context.lock() = None;

        self.set_state(LoaderState::Loaded);

        Ok(())
    }

    pub fn find_module_by_address(&self, address: u64) -> Option<ModuleInfo> {
        self.loaded_modules
            .iter()
            .find(|(&base, module)| address >= base && address < base + module.size)
            .map(|(_, module)| module.clone())
    }

    pub fn find_module_by_name(&self, name: &str) -> Option<ModuleInfo> {
        self.loaded_modules
            .values()
            .find(|module| module.name == name)
            .cloned()
    }

    pub fn get_all_modules(&self) -> alloc::vec::Vec<ModuleInfo> {
        self.loaded_modules.values().cloned().collect()
    }

    pub fn get_module_count(&self) -> u32 {
        self.total_modules.load(Ordering::Acquire)
    }

    pub fn get_load_count(&self) -> u32 {
        self.load_count.load(Ordering::Acquire)
    }

    pub fn get_unload_count(&self) -> u32 {
        self.unload_count.load(Ordering::Acquire)
    }

    pub fn get_current_context(&self) -> Option<LoadContext> {
        self.current_context.lock().as_ref().map(|c| (**c).clone())
    }

    pub fn allocate_pool(&self, size: usize, pool_type: PoolType) -> Result<*mut u8, LoaderError> {
        if !self.is_initialized() {
            return Err(LoaderError::NotInitialized);
        }

        let tag = 0x4C6F6444u32;
        let pool_type_value = match pool_type {
            PoolType::NonPagedPool => 0,
            PoolType::PagedPool => 1,
            PoolType::NonPagedPoolCacheAligned => 2,
            PoolType::PagedPoolCacheAligned => 3,
        };

        let memory = unsafe { ExAllocatePoolWithTag(pool_type_value, size, tag) };

        if memory.is_null() {
            return Err(LoaderError::MemoryAllocationFailed);
        }

        Ok(memory)
    }

    pub fn free_pool(&self, pool: *mut u8) {
        if !pool.is_null() {
            unsafe { ExFreePool(pool) };
        }
    }

    pub fn get_system_routine(&self, routine_name: &str) -> Result<*mut u8, LoaderError> {
        if !self.is_initialized() {
            return Err(LoaderError::NotInitialized);
        }

        let name_bytes = routine_name.as_bytes();
        let name_ptr = name_bytes.as_ptr() as *const u8;

        let routine = unsafe { MmGetSystemRoutineAddress(name_ptr) };

        if routine.is_null() {
            return Err(LoaderError::SymbolNotFound);
        }

        Ok(routine)
    }

    pub fn validate_module(&self, module_info: &ModuleInfo) -> bool {
        module_info.base_address != 0
            && module_info.size != 0
            && !module_info.name.is_empty()
    }

    pub fn get_module_statistics(&self) -> (u32, u32, u32, u32) {
        (
            self.get_module_count(),
            self.get_load_count(),
            self.get_unload_count(),
            self.loaded_modules.len() as u32,
        )
    }

    pub fn clear_all_modules(&mut self) {
        self.loaded_modules.clear();
        self.total_modules.store(0, Ordering::Release);
    }
}

pub static LOADER_MANAGER: Mutex<LoaderManager> = Mutex::new(LoaderManager::new());

pub fn initialize_loader() -> Result<(), LoaderError> {
    let mut manager = LOADER_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_loader() {
    let mut manager = LOADER_MANAGER.lock();
    manager.deinitialize();
}

pub fn load_module(module_info: ModuleInfo) -> Result<(), LoaderError> {
    let mut manager = LOADER_MANAGER.lock();
    manager.load_module(module_info)
}

pub fn unload_module(base_address: u64) -> Result<(), LoaderError> {
    let mut manager = LOADER_MANAGER.lock();
    manager.unload_module(base_address)
}

pub fn find_module_by_address(address: u64) -> Option<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.find_module_by_address(address)
}

pub fn find_module_by_name(name: &str) -> Option<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.find_module_by_name(name)
}

pub fn get_all_modules() -> alloc::vec::Vec<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.get_all_modules()
}

pub fn get_module_count() -> u32 {
    let manager = LOADER_MANAGER.lock();
    manager.get_module_count()
}

pub fn get_loader_state() -> LoaderState {
    let manager = LOADER_MANAGER.lock();
    manager.get_state()
}

pub fn allocate_pool(size: usize, pool_type: PoolType) -> Result<*mut u8, LoaderError> {
    let manager = LOADER_MANAGER.lock();
    manager.allocate_pool(size, pool_type)
}

pub fn free_pool(pool: *mut u8) {
    let manager = LOADER_MANAGER.lock();
    manager.free_pool(pool)
}

pub fn get_system_routine(routine_name: &str) -> Result<*mut u8, LoaderError> {
    let manager = LOADER_MANAGER.lock();
    manager.get_system_routine(routine_name)
}

pub fn get_loader_statistics() -> (u32, u32, u32, u32) {
    let manager = LOADER_MANAGER.lock();
    manager.get_module_statistics()
}

pub fn clear_all_modules() {
    let mut manager = LOADER_MANAGER.lock();
    manager.clear_all_modules()
}

pub fn is_loader_initialized() -> bool {
    let manager = LOADER_MANAGER.lock();
    manager.is_initialized()
}
