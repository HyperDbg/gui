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
