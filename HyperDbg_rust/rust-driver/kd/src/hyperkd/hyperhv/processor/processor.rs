use wdk_sys::{
    KIRQL,
    KDPC,
    GROUP_AFFINITY,
    PROCESSOR_NUMBER,
    PKDPC,
};

extern "system" {
    fn KeGetCurrentProcessorNumberEx(proc_number: *mut PROCESSOR_NUMBER) -> u32;
    fn KeQueryActiveProcessorCountEx(relationship_type: u16) -> u32;
    fn KeSetSystemGroupAffinityThread(affinity: *const GROUP_AFFINITY, previous_affinity: *mut GROUP_AFFINITY);
    fn KeRevertToUserAffinityThread();
    fn KeGetCurrentIrql() -> KIRQL;
    fn KeRaiseIrql(new_irql: KIRQL, old_irql: *mut KIRQL);
    fn KeLowerIrql(new_irql: KIRQL);
    fn KeInitializeDpc(dpc: PKDPC, routine: Option<unsafe extern "system" fn(PKDPC, *mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void)>, context: *mut core::ffi::c_void);
    fn KeInsertQueueDpc(dpc: PKDPC, arg1: *mut core::ffi::c_void, arg2: *mut core::ffi::c_void) -> i32;
    fn KeQueryActiveProcessors() -> u64;
}

pub const ALL_PROCESSOR_GROUPS: u16 = 0xFFFF;
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
        let group_affinity = GROUP_AFFINITY {
            Group: 0,
            Mask: affinity as u64,
            Reserved: [0; 3],
        };
        KeSetSystemGroupAffinityThread(&group_affinity, core::ptr::null_mut());
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
    let mut old_irql: KIRQL = 0;
    unsafe { KeRaiseIrql(new_irql, &mut old_irql) }
    old_irql
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
            let group_affinity = GROUP_AFFINITY {
                Group: 0,
                Mask: affinity as u64,
                Reserved: [0; 3],
            };
            KeSetSystemGroupAffinityThread(&group_affinity, &mut old_affinity);
            Self { old_affinity }
        }
    }
}

impl Drop for ProcessorAffinityGuard {
    fn drop(&mut self) {
        unsafe { 
            KeSetSystemGroupAffinityThread(&self.old_affinity, core::ptr::null_mut());
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

pub type DpcRoutine = extern "system" fn(*mut u8, *mut u8, *mut u8, *mut u8);

pub struct Dpc {
    dpc: KDPC,
}

impl Dpc {
    pub fn new(routine: DpcRoutine, context: *mut u8) -> Self {
        let mut dpc = Self {
            dpc: unsafe { core::mem::zeroed() },
        };

        unsafe {
            KeInitializeDpc(
                &mut dpc.dpc,
                Some(core::mem::transmute(routine)),
                context as *mut core::ffi::c_void,
            );
        }

        dpc
    }

    pub fn insert(&mut self, arg1: *mut u8, arg2: *mut u8) -> bool {
        unsafe { 
            KeInsertQueueDpc(
                &mut self.dpc, 
                arg1 as *mut core::ffi::c_void, 
                arg2 as *mut core::ffi::c_void
            ) != 0 
        }
    }
}

pub fn run_on_all_processors<F>(mut f: F)
where
    F: FnMut(u32) + Send + Sync,
{
    let processor_count = get_processor_count();

    for i in 0..processor_count {
        let affinity = 1usize << i;
        let _guard = ProcessorAffinityGuard::new(affinity);
        f(i);
    }
}

pub fn run_on_processor<F>(processor_id: u32, f: F)
where
    F: FnOnce(),
{
    let affinity = 1usize << processor_id;
    let _guard = ProcessorAffinityGuard::new(affinity);
    f();
}

pub fn get_processor_index_from_affinity(affinity: usize) -> u32 {
    let mut index = 0;
    let mut mask = affinity;

    while mask != 0 {
        if mask & 1 != 0 {
            return index;
        }
        mask >>= 1;
        index += 1;
    }

    0
}

pub fn get_affinity_from_processor_index(index: u32) -> usize {
    1usize << index
}

pub fn is_processor_active(processor_id: u32) -> bool {
    let affinity = get_affinity_from_processor_index(processor_id);
    let active_mask = unsafe { KeQueryActiveProcessors() };
    (active_mask & affinity as u64) != 0
}

pub fn for_each_processor<F>(mut f: F)
where
    F: FnMut(u32),
{
    let processor_count = get_processor_count();

    for i in 0..processor_count {
        if is_processor_active(i) {
            f(i);
        }
    }
}

pub struct ProcessorInfo {
    pub id: u32,
    pub apic_id: u32,
    pub is_bootstrap: bool,
}

pub fn get_processor_info(processor_id: u32) -> Option<ProcessorInfo> {
    if processor_id >= get_processor_count() {
        return None;
    }

    Some(ProcessorInfo {
        id: processor_id,
        apic_id: processor_id,
        is_bootstrap: processor_id == 0,
    })
}

pub fn get_bootstrap_processor_id() -> u32 {
    0
}

pub fn is_bootstrap_processor() -> bool {
    get_current_processor_number() == 0
}
