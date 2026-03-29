#![allow(non_snake_case)]
#![allow(dead_code)]

extern crate alloc;

pub mod exported {
    pub use wdk_sys::ntddk::{
        ExFreePool,
        KeGetCurrentProcessorNumberEx,
        KeQueryActiveProcessorCount,
        KeSetTargetProcessorDpc,
        KeStackAttachProcess,
        KeUnstackDetachProcess,
        MmGetPhysicalMemoryRanges,
        MmGetVirtualForPhysical,
        MmIsAddressValid,
        ObfDereferenceObject,
        PsGetCurrentProcessId,
        PsGetCurrentThreadId,
        PsLookupProcessByProcessId,
        PsLookupThreadByThreadId,
        PsTerminateSystemThread,
    };

    pub use wdk_sys::{
        BOOLEAN,
        CCHAR,
        HANDLE,
        KAPC_STATE,
        KPRIORITY,
        NTSTATUS,
        PEPROCESS,
        PETHREAD,
        PKAFFINITY,
        PKTHREAD,
        PRKAPC_STATE,
        PRKPROCESS,
        PVOID,
        ULONG,
        PRKDPC,
        PPROCESSOR_NUMBER,
        PHYSICAL_ADDRESS,
        PPHYSICAL_MEMORY_RANGE,
    };
}

pub mod not_exported {
    use super::*;

    extern "C" {
        pub fn KeGenericCallDpc(
            DpcRoutine: unsafe extern "C" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
            Context: *mut core::ffi::c_void,
        );
        pub fn KeInsertQueueDpc(Dpc: PRKDPC, SystemArgument1: PVOID, SystemArgument2: PVOID) -> BOOLEAN;
        pub fn KeSignalCallDpcDone();
        pub fn KeSignalCallDpcSynchronize();

        pub fn ObDereferenceObject(Object: PVOID);
        pub fn ObOpenObjectByPointer(
            Object: PVOID,
            DesiredAccess: u32,
            PassedAccessState: PVOID,
            AccessMode: i32,
            ObjectType: PVOID,
            HandleAttributes: u32,
            ProcessHandle: HANDLE,
        ) -> NTSTATUS;

        pub fn PsGetCurrentProcess() -> PEPROCESS;
        pub fn PsGetCurrentThread() -> PETHREAD;
        pub fn PsGetNextProcess(Process: PEPROCESS) -> PEPROCESS;
        pub fn PsGetNextProcessThread(Thread: PETHREAD, Process: PEPROCESS) -> PETHREAD;
        pub fn PsGetProcessImageFileName(Process: PEPROCESS) -> *mut u8;
        pub fn PsGetProcessPeb(Process: PEPROCESS) -> PVOID;
        pub fn PsGetProcessSectionBaseAddress(Process: PEPROCESS) -> PVOID;
        pub fn PsGetProcessWow64Process(Process: PEPROCESS) -> PVOID;

        pub fn PsGetContextThread(Thread: PETHREAD, Context: *mut core::ffi::c_void) -> NTSTATUS;
        pub fn PsSetContextThread(Thread: PETHREAD, Context: *mut core::ffi::c_void) -> NTSTATUS;

        pub fn PsSuspendProcess(Process: PEPROCESS) -> NTSTATUS;
        pub fn PsResumeProcess(Process: PEPROCESS) -> NTSTATUS;
        pub fn PsSuspendThread(Thread: PETHREAD, PreviousSuspendCount: *mut ULONG) -> NTSTATUS;
        pub fn PsResumeThread(Thread: PETHREAD, PreviousSuspendCount: *mut ULONG) -> NTSTATUS;

        pub fn RtlCopyUnicodeString(DestinationString: PVOID, SourceString: PVOID);
        pub fn RtlPcToFileHeader(PcValue: PVOID, BaseOfImage: *mut PVOID) -> PVOID;
    }
}

pub use exported::*;
pub use not_exported::*;
