use alloc::boxed::Box;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;




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

        if base_address == 0 {
            return Err(LoaderError::InvalidParameter);
        }

        self.set_state(LoaderState::Unloading);

        if self.loaded_modules.remove(&base_address).is_some() {
            self.total_modules.fetch_sub(1, Ordering::AcqRel);
            self.unload_count.fetch_add(1, Ordering::AcqRel);
        }

        self.set_state(LoaderState::Loaded);

        Ok(())
    }

    pub fn get_module(&self, base_address: u64) -> Option<ModuleInfo> {
        self.loaded_modules.get(&base_address).cloned()
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

    pub fn enumerate_modules(&self) -> alloc::vec::Vec<ModuleInfo> {
        self.loaded_modules.values().cloned().collect()
    }

    pub fn find_module_by_name(&self, name: &str) -> Option<ModuleInfo> {
        self.loaded_modules
            .values()
            .find(|m| m.name == name)
            .cloned()
    }

    pub fn find_module_by_address(&self, address: u64) -> Option<ModuleInfo> {
        self.loaded_modules
            .values()
            .find(|m| address >= m.base_address && address < m.base_address + m.size)
            .cloned()
    }
}

impl Default for LoaderManager {
    fn default() -> Self {
        Self::new()
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

pub fn get_module(base_address: u64) -> Option<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.get_module(base_address)
}

pub fn get_module_count() -> u32 {
    let manager = LOADER_MANAGER.lock();
    manager.get_module_count()
}

pub fn enumerate_modules() -> alloc::vec::Vec<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.enumerate_modules()
}

pub fn find_module_by_name(name: &str) -> Option<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.find_module_by_name(name)
}

pub fn find_module_by_address(address: u64) -> Option<ModuleInfo> {
    let manager = LOADER_MANAGER.lock();
    manager.find_module_by_address(address)
}
