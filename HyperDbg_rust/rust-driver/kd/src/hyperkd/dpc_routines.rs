#![allow(dead_code)]

use core::sync::atomic::{AtomicI32, AtomicU32, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::common::msr::{read_msr, write_msr};

static ONE_CORE_LOCK: AtomicI32 = AtomicI32::new(0);

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DpcError {
    InvalidCoreNumber,
    InsufficientResources,
    Timeout,
    Failed,
}

pub struct DpcContext {
    pub core_id: u32,
    pub routine: Option<unsafe extern "C" fn(*mut core::ffi::c_void)>,
    pub context: *mut core::ffi::c_void,
    pub arg1: u64,
    pub arg2: u64,
    pub arg3: u64,
    pub arg4: u64,
}

impl Default for DpcContext {
    fn default() -> Self {
        Self {
            core_id: 0,
            routine: None,
            context: core::ptr::null_mut(),
            arg1: 0,
            arg2: 0,
            arg3: 0,
            arg4: 0,
        }
    }
}

pub struct DpcRoutine {
    pub routine: Option<unsafe extern "C" fn(*mut DpcContext)>,
    pub context: *mut DpcContext,
}

impl Default for DpcRoutine {
    fn default() -> Self {
        Self {
            routine: None,
            context: core::ptr::null_mut(),
        }
    }
}

impl DpcRoutine {
    pub fn new(routine: unsafe extern "C" fn(*mut DpcContext), context: *mut DpcContext) -> Self {
        Self {
            routine: Some(routine),
            context,
        }
    }
}

#[repr(C)]
pub struct KDPC {
    pub type_: u16,
    pub importance: u8,
    pub number: u8,
    pub list_entry: [u64; 2],
    pub deferred_routine: unsafe extern "C" fn(*mut KDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
    pub deferred_context: *mut core::ffi::c_void,
    pub system_argument1: *mut core::ffi::c_void,
    pub system_argument2: *mut core::ffi::c_void,
}

pub unsafe fn dpc_routine_run_task_on_single_core(
    core_number: u32,
    routine: unsafe extern "C" fn(*mut core::ffi::c_void),
    context: *mut core::ffi::c_void,
) -> Result<(), DpcError> {
    let processor_count = get_processor_count();

    if core_number >= processor_count {
        return Err(DpcError::InvalidCoreNumber);
    }

    let dpc = allocate_dpc()?;
    if dpc.is_null() {
        return Err(DpcError::InsufficientResources);
    }

    initialize_dpc(dpc, dpc_routine_wrapper, context as *mut core::ffi::c_void);
    set_target_processor_dpc(dpc, core_number as i8);

    if !spinlock_try_lock(&ONE_CORE_LOCK) {
        free_dpc(dpc);
        return Err(DpcError::Failed);
    }

    insert_queue_dpc(dpc, core::ptr::null_mut(), core::ptr::null_mut());

    spinlock_lock(&ONE_CORE_LOCK);
    spinlock_unlock(&ONE_CORE_LOCK);

    free_dpc(dpc);

    Ok(())
}

pub unsafe fn dpc_routine_run_on_all_cores(
    routine: unsafe extern "C" fn(*mut KDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
    context: *mut core::ffi::c_void,
) -> Result<(), DpcError> {
    extern "C" {
        fn KeGenericCallDpc(
            routine: unsafe extern "C" fn(*mut KDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
            context: *mut core::ffi::c_void,
        );
    }

    KeGenericCallDpc(routine, context);
    Ok(())
}

unsafe extern "C" fn dpc_routine_wrapper(
    _dpc: *mut KDPC,
    context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    if let Some(routine) = core::mem::transmute::<*mut core::ffi::c_void, Option<unsafe extern "C" fn(*mut core::ffi::c_void)>>(context) {
        routine(context);
    }

    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_write_msr(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    let current_core = get_current_processor_number();

    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if g_DbgState.is_null() {
        spinlock_unlock(&ONE_CORE_LOCK);
        return;
    }

    let dbg_state = g_DbgState.add(current_core as usize * core::mem::size_of::<ProcessorDebuggingState>());

    let msr = (*dbg_state).msr_state.msr;
    let value = (*dbg_state).msr_state.value;

    write_msr(msr, value);

    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_read_msr(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    let current_core = get_current_processor_number();

    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if g_DbgState.is_null() {
        spinlock_unlock(&ONE_CORE_LOCK);
        return;
    }

    let dbg_state = g_DbgState.add(current_core as usize * core::mem::size_of::<ProcessorDebuggingState>());

    let msr = (*dbg_state).msr_state.msr;
    let value = read_msr(msr);

    (*dbg_state).msr_state.value = value;

    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_write_msr_to_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let current_core = get_current_processor_number();

    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if !g_DbgState.is_null() {
        let dbg_state = g_DbgState.add(current_core as usize * core::mem::size_of::<ProcessorDebuggingState>());
        let msr = (*dbg_state).msr_state.msr;
        let value = (*dbg_state).msr_state.value;
        write_msr(msr, value);
    }

    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_read_msr_to_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let current_core = get_current_processor_number();

    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if !g_DbgState.is_null() {
        let dbg_state = g_DbgState.add(current_core as usize * core::mem::size_of::<ProcessorDebuggingState>());
        let msr = (*dbg_state).msr_state.msr;
        let value = read_msr(msr);
        (*dbg_state).msr_state.value = value;
    }

    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_vmexit_and_halt_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    use crate::hyperkd::hyperhv::vmm::vmx::vmcall::VmcallNumber;

    extern "C" {
        fn asm_vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64;
    }

    asm_vmx_vmcall(VmcallNumber::Test as u64, 0, 0, 0);

    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

#[repr(C)]
pub struct ProcessorDebuggingState {
    pub core_id: u32,
    pub msr_state: MsrState,
}

#[repr(C)]
pub struct MsrState {
    pub msr: u32,
    pub value: u64,
}

#[inline(always)]
fn spinlock_try_lock(lock: &AtomicI32) -> bool {
    lock.compare_exchange(0, 1, Ordering::Acquire, Ordering::Relaxed).is_ok()
}

#[inline(always)]
fn spinlock_lock(lock: &AtomicI32) {
    while lock.compare_exchange(0, 1, Ordering::Acquire, Ordering::Relaxed).is_err() {
        core::hint::spin_loop();
    }
}

#[inline(always)]
fn spinlock_unlock(lock: &AtomicI32) {
    lock.store(0, Ordering::Release);
}

unsafe fn allocate_dpc() -> Result<*mut KDPC, DpcError> {
    extern "C" {
        fn ExAllocatePoolWithTag(pool_type: u32, size: usize, tag: u32) -> *mut core::ffi::c_void;
    }

    const NON_PAGED_POOL: u32 = 0;
    const POOL_TAG: u32 = 0x6370644B;

    let dpc = ExAllocatePoolWithTag(NON_PAGED_POOL, core::mem::size_of::<KDPC>(), POOL_TAG);

    if dpc.is_null() {
        Err(DpcError::InsufficientResources)
    } else {
        Ok(dpc as *mut KDPC)
    }
}

unsafe fn free_dpc(dpc: *mut KDPC) {
    extern "C" {
        fn ExFreePool(address: *mut core::ffi::c_void);
    }

    if !dpc.is_null() {
        ExFreePool(dpc as *mut core::ffi::c_void);
    }
}

unsafe fn initialize_dpc(
    dpc: *mut KDPC,
    routine: unsafe extern "C" fn(*mut KDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
    context: *mut core::ffi::c_void,
) {
    extern "C" {
        fn KeInitializeDpc(
            dpc: *mut KDPC,
            routine: unsafe extern "C" fn(*mut KDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void),
            context: *mut core::ffi::c_void,
        );
    }

    KeInitializeDpc(dpc, routine, context);
}

unsafe fn set_target_processor_dpc(dpc: *mut KDPC, processor_number: i8) {
    extern "C" {
        fn KeSetTargetProcessorDpc(dpc: *mut KDPC, processor_number: i8);
    }

    KeSetTargetProcessorDpc(dpc, processor_number);
}

unsafe fn insert_queue_dpc(dpc: *mut KDPC, system_argument1: *mut core::ffi::c_void, system_argument2: *mut core::ffi::c_void) {
    extern "C" {
        fn KeInsertQueueDpc(dpc: *mut KDPC, system_argument1: *mut core::ffi::c_void, system_argument2: *mut core::ffi::c_void) -> bool;
    }

    KeInsertQueueDpc(dpc, system_argument1, system_argument2);
}

unsafe fn signal_call_dpc_done(system_argument1: *mut core::ffi::c_void) {
    extern "C" {
        fn KeSignalCallDpcDone(system_argument1: *mut core::ffi::c_void);
    }

    KeSignalCallDpcDone(system_argument1);
}

unsafe fn signal_call_dpc_synchronize(system_argument2: *mut core::ffi::c_void) {
    extern "C" {
        fn KeSignalCallDpcSynchronize(system_argument2: *mut core::ffi::c_void);
    }

    KeSignalCallDpcSynchronize(system_argument2);
}

pub fn get_processor_count() -> u32 {
    extern "C" {
        fn KeQueryActiveProcessorCount(relationship_info: *mut core::ffi::c_void) -> u32;
    }

    unsafe { KeQueryActiveProcessorCount(core::ptr::null_mut()) }
}

pub fn get_current_processor_number() -> u32 {
    extern "C" {
        fn KeGetCurrentProcessorNumberEx(info: *mut core::ffi::c_void) -> u32;
    }

    unsafe { KeGetCurrentProcessorNumberEx(core::ptr::null_mut()) }
}

pub unsafe fn broadcast_msr_write(msr: u32, value: u64) -> Result<(), DpcError> {
    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if g_DbgState.is_null() {
        return Err(DpcError::Failed);
    }

    let processor_count = get_processor_count();
    for i in 0..processor_count {
        let dbg_state = g_DbgState.add(i as usize * core::mem::size_of::<ProcessorDebuggingState>());
        (*dbg_state).msr_state.msr = msr;
        (*dbg_state).msr_state.value = value;
    }

    dpc_routine_run_on_all_cores(dpc_routine_write_msr_to_all_cores, core::ptr::null_mut())
}

pub unsafe fn broadcast_msr_read(msr: u32) -> Result<(), DpcError> {
    extern "C" {
        static mut g_DbgState: *mut u8;
    }

    if g_DbgState.is_null() {
        return Err(DpcError::Failed);
    }

    let processor_count = get_processor_count();
    for i in 0..processor_count {
        let dbg_state = g_DbgState.add(i as usize * core::mem::size_of::<ProcessorDebuggingState>());
        (*dbg_state).msr_state.msr = msr;
    }

    dpc_routine_run_on_all_cores(dpc_routine_read_msr_to_all_cores, core::ptr::null_mut())
}

static DPC_QUEUE_HEAD: AtomicU64 = AtomicU64::new(0);
static DPC_QUEUE_TAIL: AtomicU64 = AtomicU64::new(0);
static DPC_ROUND_ROBIN_INDEX: AtomicU32 = AtomicU32::new(0);

pub fn enqueue_dpc_round_robin(routine: DpcRoutine) -> Result<(), DpcError> {
    if routine.routine.is_none() {
        return Err(DpcError::Failed);
    }

    let processor_count = get_processor_count();
    if processor_count == 0 {
        return Err(DpcError::Failed);
    }

    let target_core = DPC_ROUND_ROBIN_INDEX.fetch_add(1, Ordering::AcqRel) % processor_count;

    unsafe {
        let context = routine.context;
        if !context.is_null() {
            (*context).core_id = target_core;
        }

        if let Some(r) = routine.routine {
            r(context);
        }
    }

    Ok(())
}

pub fn broadcast_to_all_cores_dpc(target_core: u32) -> Result<(), DpcError> {
    let processor_count = get_processor_count();

    for i in 0..processor_count {
        if target_core == 0xFFFFFFFF || i == target_core {
            let mut context = DpcContext::default();
            context.core_id = i;
            
            unsafe {
                if let Some(r) = context.routine {
                    r(context.context);
                }
            }
        }
    }

    Ok(())
}

pub unsafe fn dpc_perform_task_on_core(
    core_id: u32,
    task_type: u32,
    param1: u64,
    param2: u64,
    param3: u64,
) -> Result<u64, DpcError> {
    let processor_count = get_processor_count();
    if core_id >= processor_count {
        return Err(DpcError::InvalidCoreNumber);
    }

    let mut context = DpcContext::default();
    context.core_id = core_id;
    context.arg1 = task_type as u64;
    context.arg2 = param1;
    context.arg3 = param2;
    context.arg4 = param3;

    Ok(0)
}

pub fn dpc_synchronize_all_cores() -> Result<(), DpcError> {
    extern "C" {
        fn KeFlushQueuedDpcs();
    }

    unsafe {
        KeFlushQueuedDpcs();
    }

    Ok(())
}

#[derive(Debug, Clone, Copy)]
#[repr(C)]
pub struct DebuggerEventOptions {
    pub optional_param1: u64,
    pub optional_param2: u64,
    pub optional_param3: u64,
    pub optional_param4: u64,
}

#[derive(Debug, Clone, Copy)]
#[repr(C)]
pub struct EptSingleHookUnhookingDetails {
    pub physical_address: u64,
    pub original_entry: u64,
}

pub const EXCEPTION_VECTOR_BREAKPOINT: u64 = 3;
pub const EXCEPTION_VECTOR_DEBUG_BREAKPOINT: u64 = 1;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u64)]
pub enum VmcallNumber {
    Test = 0x0,
    Vmxoff = 0x1,
    SetHiddenCcBreakpoint = 0x2,
    RemoveAndInvalidateAllEptHooks = 0x3,
    RemoveAndInvalidateSingleEptHook = 0x4,
    ChangeMsrBitmapRead = 0x5,
    ChangeMsrBitmapWrite = 0x6,
    ResetMsrBitmapRead = 0x7,
    ResetMsrBitmapWrite = 0x8,
    SetRdtscExiting = 0x9,
    UnsetRdtscExiting = 0xA,
    SetRdpmcExiting = 0xB,
    UnsetRdpmcExiting = 0xC,
    SetExceptionBitmap = 0xD,
    UnsetExceptionBitmap = 0xE,
    EnableMovToDebugRegsExiting = 0xF,
    DisableMovToDebugRegsExiting = 0x10,
    EnableMovToControlRegsExiting = 0x11,
    DisableMovToControlRegsExiting = 0x12,
    EnableExternalInterruptExiting = 0x13,
    DisableExternalInterruptExiting = 0x14,
    EnableSyscallHookEfer = 0x15,
    DisableSyscallHookEfer = 0x16,
    ChangeIoBitmap = 0x17,
    ResetIoBitmap = 0x18,
    EnableMovToCr3Exiting = 0x19,
    DisableMovToCr3Exiting = 0x1A,
    ChangeToMbecSupportedEptp = 0x1B,
    RestoreToNormalEptp = 0x1C,
    DisableOrEnableMbec = 0x1D,
    EnableDirtyLoggingMechanism = 0x1E,
    DisableDirtyLoggingMechanism = 0x1F,
    DisableRdtscExitingOnlyForTscEvents = 0x20,
    DisableMovToHwDrExitingOnlyForDrEvents = 0x21,
    DisableMovToCrExitingOnlyForCrEvents = 0x22,
    ResetExceptionBitmapOnlyOnClearingExceptionEvents = 0x23,
    DisableExternalInterruptExitingOnlyToClearInterruptCommands = 0x24,
    SetVmExitOnNmis = 0x25,
    UnsetVmExitOnNmis = 0x26,
    UnhookAllPages = 0x27,
    UnhookSinglePage = 0x28,
    InveptAllContexts = 0x29,
    InveptSingleContext = 0x2A,
}

#[inline(always)]
unsafe fn asm_vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64 {
    extern "C" {
        fn __vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64;
    }
    __vmx_vmcall(vmcall_number, param1, param2, param3)
}

pub unsafe fn dpc_routine_perform_virtualization(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    extern "C" {
        fn VmxPerformVirtualizationOnSpecificCore();
    }
    VmxPerformVirtualizationOnSpecificCore();
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_perform_change_msr_bitmap_read_on_single_core(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeMsrBitmapRead as u64, deferred_context as u64, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_change_msr_bitmap_write_on_single_core(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeMsrBitmapWrite as u64, deferred_context as u64, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_enable_rdtsc_exiting_on_single_core(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetRdtscExiting as u64, 0, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_enable_rdpmc_exiting_on_single_core(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetRdpmcExiting as u64, 0, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_set_exception_bitmap_on_single_core(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetExceptionBitmap as u64, deferred_context as u64, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_enable_mov_to_debug_registers_exiting(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableMovToDebugRegsExiting as u64, 0, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_enable_mov_to_control_register_exiting(
    _dpc: *mut KDPC,
    event_options: *mut DebuggerEventOptions,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    let opt1 = if event_options.is_null() { 0 } else { (*event_options).optional_param1 };
    let opt2 = if event_options.is_null() { 0 } else { (*event_options).optional_param2 };
    asm_vmx_vmcall(VmcallNumber::EnableMovToControlRegsExiting as u64, opt1, opt2, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_set_external_interrupt_exiting_on_single_core(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableExternalInterruptExiting as u64, 0, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_enable_efer_syscall_hook_on_single_core(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableSyscallHookEfer as u64, 0, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_perform_change_io_bitmap_on_single_core(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    _system_argument1: *mut core::ffi::c_void,
    _system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeIoBitmap as u64, deferred_context as u64, 0, 0);
    spinlock_unlock(&ONE_CORE_LOCK);
}

pub unsafe fn dpc_routine_enable_mov_to_cr3_exiting(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableMovToCr3Exiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_change_to_mbec_supported_eptp(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeToMbecSupportedEptp as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_restore_to_normal_eptp(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::RestoreToNormalEptp as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_or_disable_mbec(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableOrEnableMbec as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_mov_to_cr3_exiting(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableMovToCr3Exiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_efer_syscall_events(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableSyscallHookEfer as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_efer_syscall_events(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableSyscallHookEfer as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_pml(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableDirtyLoggingMechanism as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_pml(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableDirtyLoggingMechanism as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_change_msr_bitmap_read_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeMsrBitmapRead as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_reset_msr_bitmap_read_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ResetMsrBitmapRead as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_change_msr_bitmap_write_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeMsrBitmapWrite as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_reset_msr_bitmap_write_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ResetMsrBitmapWrite as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_rdtsc_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetRdtscExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_rdtsc_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetRdtscExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_rdtsc_exiting_for_clearing_tsc_events_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableRdtscExitingOnlyForTscEvents as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_mov2_dr_exiting_for_clearing_dr_events_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableMovToHwDrExitingOnlyForDrEvents as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_mov2_cr_exiting_for_clearing_cr_events_all_cores(
    _dpc: *mut KDPC,
    event_options: *mut DebuggerEventOptions,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let opt1 = if event_options.is_null() { 0 } else { (*event_options).optional_param1 };
    let opt2 = if event_options.is_null() { 0 } else { (*event_options).optional_param2 };
    asm_vmx_vmcall(VmcallNumber::DisableMovToCrExitingOnlyForCrEvents as u64, opt1, opt2, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_rdpmc_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetRdpmcExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_rdpmc_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetRdpmcExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_set_exception_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetExceptionBitmap as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_unset_exception_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetExceptionBitmap as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_reset_exception_bitmap_only_on_clearing_exception_events_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ResetExceptionBitmapOnlyOnClearingExceptionEvents as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_mov_debug_register_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableMovToDebugRegsExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_mov_control_register_exiting_all_cores(
    _dpc: *mut KDPC,
    event_options: *mut DebuggerEventOptions,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let opt1 = if event_options.is_null() { 0 } else { (*event_options).optional_param1 };
    let opt2 = if event_options.is_null() { 0 } else { (*event_options).optional_param2 };
    asm_vmx_vmcall(VmcallNumber::EnableMovToControlRegsExiting as u64, opt1, opt2, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_mov_control_register_exiting_all_cores(
    _dpc: *mut KDPC,
    event_options: *mut DebuggerEventOptions,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let opt1 = if event_options.is_null() { 0 } else { (*event_options).optional_param1 };
    let opt2 = if event_options.is_null() { 0 } else { (*event_options).optional_param2 };
    asm_vmx_vmcall(VmcallNumber::DisableMovToControlRegsExiting as u64, opt1, opt2, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_mov_debug_register_exiting_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableMovToDebugRegsExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_set_enable_external_interrupt_exiting_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::EnableExternalInterruptExiting as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_set_disable_external_interrupt_exiting_only_on_clearing_interrupt_events_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::DisableExternalInterruptExitingOnlyToClearInterruptCommands as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_change_io_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ChangeIoBitmap as u64, deferred_context as u64, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_reset_io_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::ResetIoBitmap as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_breakpoint_on_exception_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetExceptionBitmap as u64, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_breakpoint_on_exception_bitmap_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetExceptionBitmap as u64, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_nmi_vmexit_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetVmExitOnNmis as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_nmi_vmexit_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetVmExitOnNmis as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_enable_db_and_bp_exiting_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::SetExceptionBitmap as u64, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    asm_vmx_vmcall(VmcallNumber::SetExceptionBitmap as u64, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_disable_db_and_bp_exiting_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnsetExceptionBitmap as u64, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    asm_vmx_vmcall(VmcallNumber::UnsetExceptionBitmap as u64, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_remove_hook_and_invalidate_all_entries_on_all_cores(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    asm_vmx_vmcall(VmcallNumber::UnhookAllPages as u64, 0, 0, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_remove_hook_and_invalidate_single_entry_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    let unhooking_detail = deferred_context as *mut EptSingleHookUnhookingDetails;
    let physical_address = if unhooking_detail.is_null() { 0 } else { (*unhooking_detail).physical_address };
    let original_entry = if unhooking_detail.is_null() { 0 } else { (*unhooking_detail).original_entry };
    asm_vmx_vmcall(VmcallNumber::UnhookSinglePage as u64, physical_address, original_entry, 0);
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_invalidate_ept_on_all_cores(
    _dpc: *mut KDPC,
    deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    if deferred_context.is_null() {
        asm_vmx_vmcall(VmcallNumber::InveptAllContexts as u64, 0, 0, 0);
    } else {
        extern "C" {
            static mut g_GuestState: *mut u8;
        }
        let current_core = get_current_processor_number();
        if !g_GuestState.is_null() {
            let guest_state = g_GuestState.add(current_core as usize * core::mem::size_of::<VIRTUAL_MACHINE_STATE>());
            let ept_pointer = (*guest_state).ept_pointer.value;
            asm_vmx_vmcall(VmcallNumber::InveptSingleContext as u64, ept_pointer, 0, 0);
        }
    }
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_initialize_guest(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    extern "C" {
        fn HvInitializeGuest();
    }
    HvInitializeGuest();
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

pub unsafe fn dpc_routine_terminate_guest(
    _dpc: *mut KDPC,
    _deferred_context: *mut core::ffi::c_void,
    system_argument1: *mut core::ffi::c_void,
    system_argument2: *mut core::ffi::c_void,
) {
    extern "C" {
        fn HvTerminateGuest();
    }
    HvTerminateGuest();
    signal_call_dpc_synchronize(system_argument2);
    signal_call_dpc_done(system_argument1);
}

use crate::hyperkd::hyperhv::state::VIRTUAL_MACHINE_STATE;
