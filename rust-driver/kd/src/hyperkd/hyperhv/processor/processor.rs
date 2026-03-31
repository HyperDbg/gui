use crate::generated::*;

use wdk_sys::{
    KIRQL,
    KDPC,
    GROUP_AFFINITY,
    PGROUP_AFFINITY,
    PROCESSOR_NUMBER,
    PRKDPC,
    PVOID,
    USHORT,
};

pub const ALL_PROCESSOR_GROUPS: USHORT = 0xFFFF;
pub const IRQL_PASSIVE: KIRQL = 0;
pub const IRQL_APC: KIRQL = 1;
pub const IRQL_DISPATCH: KIRQL = 2;
pub const IRQL_CLOCK: KIRQL = 28;
pub const IRQL_DEVICE: KIRQL = 3;
pub const IRQL_DPC: KIRQL = 2;
pub const IRQL_HIGH: KIRQL = 31;

#[inline(always)]
pub fn get_current_processor_number() -> u32 {
    unsafe { 
        let mut proc_number: PROCESSOR_NUMBER = core::mem::zeroed();
        KeGetCurrentProcessorNumberEx(&mut proc_number);
        proc_number.Number as u32
    }
}

#[inline(always)]
pub fn get_processor_count() -> u32 {
    unsafe { KeQueryActiveProcessorCountEx(ALL_PROCESSOR_GROUPS) }
}

#[inline(always)]
pub fn set_processor_affinity(affinity: usize) {
    unsafe { 
        let mut group_affinity = GROUP_AFFINITY {
            Mask: affinity as u64,
            Group: 0,
            Reserved: [0; 3],
        };
        KeSetSystemGroupAffinityThread(&mut group_affinity as PGROUP_AFFINITY, core::ptr::null_mut());
    }
}

#[inline(always)]
pub fn revert_to_user_affinity() {
    unsafe { KeRevertToUserAffinityThread() }
}

#[inline(always)]
pub fn get_current_irql() -> KIRQL {
    unsafe { KeGetCurrentIrql() }
}

#[inline(always)]
pub fn raise_irql(new_irql: KIRQL) -> KIRQL {
    unsafe { KfRaiseIrql(new_irql) }
}

#[inline(always)]
pub fn lower_irql(new_irql: KIRQL) {
    unsafe { KeLowerIrql(new_irql) }
}

pub struct ProcessorAffinityGuard {
    old_affinity: GROUP_AFFINITY,
}

impl ProcessorAffinityGuard {
    pub fn new(affinity: usize) -> Self {
        unsafe {
            let mut old_affinity: GROUP_AFFINITY = core::mem::zeroed();
            let mut group_affinity = GROUP_AFFINITY {
                Mask: affinity as u64,
                Group: 0,
                Reserved: [0; 3],
            };
            KeSetSystemGroupAffinityThread(&mut group_affinity as PGROUP_AFFINITY, &mut old_affinity as PGROUP_AFFINITY);
            Self { old_affinity }
        }
    }
}

impl Drop for ProcessorAffinityGuard {
    fn drop(&mut self) {
        unsafe { 
            KeSetSystemGroupAffinityThread(&mut self.old_affinity as PGROUP_AFFINITY, core::ptr::null_mut());
        }
    }
}

pub struct IrqlGuard {
    old_irql: KIRQL,
}

impl IrqlGuard {
    pub fn new(new_irql: KIRQL) -> Self {
        let old_irql = raise_irql(new_irql);
        Self { old_irql }
    }

    pub fn raise_to_dispatch() -> Self {
        Self::new(IRQL_DISPATCH)
    }

    pub fn raise_to_high() -> Self {
        Self::new(IRQL_HIGH)
    }
}

impl Drop for IrqlGuard {
    fn drop(&mut self) {
        lower_irql(self.old_irql);
    }
}

pub type DpcRoutine = unsafe extern "C" fn(Dpc: PRKDPC, DeferredContext: PVOID, SystemArgument1: PVOID, SystemArgument2: PVOID);

pub struct Dpc {
    dpc: KDPC,
}

impl Dpc {
    pub fn new(routine: DpcRoutine, context: PVOID) -> Self {
        let mut dpc = Self {
            dpc: unsafe { core::mem::zeroed() },
        };

        unsafe {
            KeInitializeDpc(
                &mut dpc.dpc,
                Some(routine),
                context,
            );
        }

        dpc
    }

    pub fn insert(&mut self, arg1: PVOID, arg2: PVOID) -> bool {
        unsafe { KeInsertQueueDpc(&mut self.dpc, arg1, arg2) != 0 }
    }
}

pub fn query_active_processors() -> u64 {
    unsafe { KeQueryActiveProcessors() }
}

pub fn run_on_all_processors<F>(mut _callback: F) 
where 
    F: FnMut(u32) + Send + Sync,
{
    let processor_count = get_processor_count();
    for i in 0..processor_count {
        let _guard = ProcessorAffinityGuard::new(1 << i);
        _callback(i);
    }
}

pub fn run_on_processor<F>(processor_id: u32, mut callback: F)
where
    F: FnMut() + Send + Sync,
{
    let _guard = ProcessorAffinityGuard::new(1 << processor_id);
    callback();
}
