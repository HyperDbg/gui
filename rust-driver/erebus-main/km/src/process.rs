use wdk::nt_success;
use wdk_sys::{
    HANDLE, NTSTATUS, PEPROCESS,
    ntddk::{ObfDereferenceObject, PsLookupProcessByProcessId},
};

pub(crate) struct Process {
    pub process: PEPROCESS,
}

impl Process {
    pub fn by_id(process_id: u32) -> Result<Self, NTSTATUS> {
        let mut process = core::ptr::null_mut();

        let status = unsafe { PsLookupProcessByProcessId(process_id as HANDLE, &mut process) };

        if nt_success(status) {
            Ok(Self { process })
        } else {
            Err(status)
        }
    }
}

impl Drop for Process {
    fn drop(&mut self) {
        if !self.process.is_null() {
            unsafe {
                ObfDereferenceObject(self.process as _);
            }
        }
    }
}
