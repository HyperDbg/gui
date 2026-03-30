use alloc::collections::BTreeMap;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

use wdk_sys::{
    HANDLE,
    PEPROCESS,
    PVOID,
    STATUS_SUCCESS,
};

use wdk_sys::ntddk::{
    PsLookupProcessByProcessId,
    ObfDereferenceObject,
    IoGetCurrentProcess,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ProcessError {
    NotInitialized,
    InvalidProcessId,
    ProcessNotFound,
    AccessDenied,
    OperationFailed,
}

#[derive(Debug, Clone, Copy)]
pub struct ProcessInfo {
    pub process_id: u32,
    pub parent_process_id: u32,
    pub session_id: u32,
    pub thread_count: u32,
    pub handle_count: u32,
    pub peb_address: u64,
    pub base_address: u64,
    pub image_base: u64,
    pub entry_point: u64,
}

#[derive(Debug, Clone)]
pub struct ProcessObject {
    pub process_id: u32,
    pub process_handle: PEPROCESS,
    pub image_name: alloc::string::String,
    pub image_path: alloc::string::String,
    pub command_line: alloc::string::String,
    pub thread_ids: alloc::vec::Vec<u32>,
    pub is_debugged: bool,
    pub is_suspended: bool,
    pub creation_time: u64,
    pub exit_time: u64,
    pub exit_code: u32,
}

unsafe impl Send for ProcessObject {}

impl ProcessObject {
    pub fn new(process_id: u32) -> Self {
        Self {
            process_id,
            process_handle: core::ptr::null_mut(),
            image_name: alloc::string::String::new(),
            image_path: alloc::string::String::new(),
            command_line: alloc::string::String::new(),
            thread_ids: alloc::vec::Vec::new(),
            is_debugged: false,
            is_suspended: false,
            creation_time: 0,
            exit_time: 0,
            exit_code: 0,
        }
    }
}

pub struct ProcessManager {
    initialized: AtomicBool,
    processes: Mutex<BTreeMap<u32, ProcessObject>>,
    current_process_id: AtomicU32,
    total_processes: AtomicU32,
    debugged_processes: AtomicU32,
}

unsafe impl Send for ProcessManager {}
unsafe impl Sync for ProcessManager {}

impl ProcessManager {
    pub const fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            processes: Mutex::new(BTreeMap::new()),
            current_process_id: AtomicU32::new(0),
            total_processes: AtomicU32::new(0),
            debugged_processes: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), ProcessError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(ProcessError::NotInitialized);
        }

        let current_process = unsafe { IoGetCurrentProcess() };
        if current_process.is_null() {
            return Err(ProcessError::ProcessNotFound);
        }

        self.current_process_id.store(self.extract_process_id(current_process), Ordering::Release);
        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        let mut processes = self.processes.lock();
        processes.clear();
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_current_process_id(&self) -> u32 {
        self.current_process_id.load(Ordering::Acquire)
    }

    pub fn lookup_process(&self, process_id: u32) -> Result<ProcessObject, ProcessError> {
        if !self.is_initialized() {
            return Err(ProcessError::NotInitialized);
        }

        let processes = self.processes.lock();
        processes.get(&process_id).cloned().ok_or(ProcessError::ProcessNotFound)
    }

    pub fn lookup_process_by_handle(&self, process_handle: PEPROCESS) -> Result<ProcessObject, ProcessError> {
        if !self.is_initialized() {
            return Err(ProcessError::NotInitialized);
        }

        let process_id = self.extract_process_id(process_handle);
        self.lookup_process(process_id)
    }

    pub fn add_process(&self, process: ProcessObject) -> Result<(), ProcessError> {
        if !self.is_initialized() {
            return Err(ProcessError::NotInitialized);
        }

        let mut processes = self.processes.lock();
        processes.insert(process.process_id, process);
        self.total_processes.fetch_add(1, Ordering::AcqRel);

        Ok(())
    }

    pub fn remove_process(&self, process_id: u32) -> Result<(), ProcessError> {
        if !self.is_initialized() {
            return Err(ProcessError::NotInitialized);
        }

        let mut processes = self.processes.lock();
        if processes.remove(&process_id).is_some() {
            Ok(())
        } else {
            Err(ProcessError::ProcessNotFound)
        }
    }

    pub fn get_all_processes(&self) -> alloc::vec::Vec<ProcessObject> {
        let processes = self.processes.lock();
        processes.values().cloned().collect()
    }

    pub fn get_process_count(&self) -> u32 {
        let processes = self.processes.lock();
        processes.len() as u32
    }

    pub fn get_total_processes(&self) -> u32 {
        self.total_processes.load(Ordering::Acquire)
    }

    pub fn set_debugged(&self, process_id: u32, debugged: bool) -> Result<(), ProcessError> {
        let mut processes = self.processes.lock();
        if let Some(process) = processes.get_mut(&process_id) {
            if process.is_debugged != debugged {
                process.is_debugged = debugged;
                if debugged {
                    self.debugged_processes.fetch_add(1, Ordering::AcqRel);
                } else {
                    self.debugged_processes.fetch_sub(1, Ordering::AcqRel);
                }
            }
            Ok(())
        } else {
            Err(ProcessError::ProcessNotFound)
        }
    }

    pub fn is_debugged(&self, process_id: u32) -> bool {
        let processes = self.processes.lock();
        processes.get(&process_id).map_or(false, |p| p.is_debugged)
    }

    pub fn get_debugged_process_count(&self) -> u32 {
        self.debugged_processes.load(Ordering::Acquire)
    }

    pub fn add_thread_to_process(&self, process_id: u32, thread_id: u32) -> Result<(), ProcessError> {
        let mut processes = self.processes.lock();
        if let Some(process) = processes.get_mut(&process_id) {
            if !process.thread_ids.contains(&thread_id) {
                process.thread_ids.push(thread_id);
            }
            Ok(())
        } else {
            Err(ProcessError::ProcessNotFound)
        }
    }

    pub fn remove_thread_from_process(&self, process_id: u32, thread_id: u32) -> Result<(), ProcessError> {
        let mut processes = self.processes.lock();
        if let Some(process) = processes.get_mut(&process_id) {
            process.thread_ids.retain(|&tid| tid != thread_id);
            Ok(())
        } else {
            Err(ProcessError::ProcessNotFound)
        }
    }

    pub fn get_process_threads(&self, process_id: u32) -> Result<alloc::vec::Vec<u32>, ProcessError> {
        let processes = self.processes.lock();
        processes.get(&process_id)
            .map(|p| p.thread_ids.clone())
            .ok_or(ProcessError::ProcessNotFound)
    }

    fn extract_process_id(&self, _process_handle: PEPROCESS) -> u32 {
        0
    }

    pub fn lookup_process_object(&self, process_id: u32) -> Result<PEPROCESS, ProcessError> {
        if !self.is_initialized() {
            return Err(ProcessError::NotInitialized);
        }

        let mut process: PEPROCESS = core::ptr::null_mut();
        let status = unsafe {
            PsLookupProcessByProcessId(process_id as HANDLE, &mut process)
        };

        if status != STATUS_SUCCESS {
            return Err(ProcessError::ProcessNotFound);
        }

        Ok(process)
    }

    pub fn dereference_process(&self, process: PEPROCESS) {
        if !process.is_null() {
            unsafe {
                ObfDereferenceObject(process as PVOID);
            }
        }
    }
}

impl Default for ProcessManager {
    fn default() -> Self {
        Self::new()
    }
}

pub static PROCESS_MANAGER: Mutex<ProcessManager> = Mutex::new(ProcessManager::new());

pub fn initialize_process_manager() -> Result<(), ProcessError> {
    let mut manager = PROCESS_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_process_manager() {
    let mut manager = PROCESS_MANAGER.lock();
    manager.deinitialize();
}

pub fn get_current_process_id() -> u32 {
    let manager = PROCESS_MANAGER.lock();
    manager.get_current_process_id()
}

pub fn lookup_process(process_id: u32) -> Result<ProcessObject, ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.lookup_process(process_id)
}

pub fn lookup_process_by_handle(process_handle: PEPROCESS) -> Result<ProcessObject, ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.lookup_process_by_handle(process_handle)
}

pub fn add_process(process: ProcessObject) -> Result<(), ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.add_process(process)
}

pub fn remove_process(process_id: u32) -> Result<(), ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.remove_process(process_id)
}

pub fn get_all_processes() -> alloc::vec::Vec<ProcessObject> {
    let manager = PROCESS_MANAGER.lock();
    manager.get_all_processes()
}

pub fn get_process_count() -> u32 {
    let manager = PROCESS_MANAGER.lock();
    manager.get_process_count()
}

pub fn set_process_debugged(process_id: u32, debugged: bool) -> Result<(), ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.set_debugged(process_id, debugged)
}

pub fn is_process_debugged(process_id: u32) -> bool {
    let manager = PROCESS_MANAGER.lock();
    manager.is_debugged(process_id)
}

pub fn add_thread_to_process(process_id: u32, thread_id: u32) -> Result<(), ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.add_thread_to_process(process_id, thread_id)
}

pub fn remove_thread_from_process(process_id: u32, thread_id: u32) -> Result<(), ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.remove_thread_from_process(process_id, thread_id)
}

pub fn get_process_threads(process_id: u32) -> Result<alloc::vec::Vec<u32>, ProcessError> {
    let manager = PROCESS_MANAGER.lock();
    manager.get_process_threads(process_id)
}
