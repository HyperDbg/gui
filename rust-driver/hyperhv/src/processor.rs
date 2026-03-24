extern "system" {
    fn KeGetCurrentProcessorNumberEx(proc_number: *mut u32) -> u32;

    fn KeQueryActiveProcessorCountEx(relationship_type: u32) -> u32;

    fn KeSetSystemAffinityThread(affinity: usize);

    fn KeRevertToUserAffinityThread();

    fn KeGetCurrentIrql() -> u8;

    fn KeRaiseIrql(new_irql: u8, old_irql: *mut u8);

    fn KeLowerIrql(new_irql: u8);

    fn KeInitializeDpc(
        dpc: *mut u8,
        routine: extern "system" fn(*mut u8, *mut u8, *mut u8, *mut u8),
        context: *mut u8,
    );

    fn KeInsertQueueDpc(dpc: *mut u8, arg1: *mut u8, arg2: *mut u8) -> bool;

    fn KeGenericCallDpc(
        routine: extern "system" fn(*mut u8, *mut u8, *mut u8, *mut u8),
        context: *mut u8,
    );

    fn KeSignalCallDpcDone(context: *mut u8);
}

pub const IRQL_PASSIVE: u8 = 0;
pub const IRQL_APC: u8 = 1;
pub const IRQL_DISPATCH: u8 = 2;
pub const IRQL_CLOCK: u8 = 28;
pub const IRQL_DEVICE: u8 = 3;
pub const IRQL_DPC: u8 = 2;
pub const IRQL_HIGH: u8 = 31;

#[inline(always)]
pub fn get_current_processor_number() -> u32 {
    unsafe { KeGetCurrentProcessorNumberEx(core::ptr::null_mut()) }
}

#[inline(always)]
pub fn get_processor_count() -> u32 {
    unsafe { KeQueryActiveProcessorCountEx(0) }
}

#[inline(always)]
pub fn set_processor_affinity(affinity: usize) {
    unsafe { KeSetSystemAffinityThread(affinity) }
}

#[inline(always)]
pub fn revert_to_user_affinity() {
    unsafe { KeRevertToUserAffinityThread() }
}

#[inline(always)]
pub fn get_current_irql() -> u8 {
    unsafe { KeGetCurrentIrql() }
}

#[inline(always)]
pub fn raise_irql(new_irql: u8) -> u8 {
    let mut old_irql: u8 = 0;
    unsafe { KeRaiseIrql(new_irql, &mut old_irql) }
    old_irql
}

#[inline(always)]
pub fn lower_irql(new_irql: u8) {
    unsafe { KeLowerIrql(new_irql) }
}

pub struct ProcessorAffinityGuard {
    old_affinity: usize,
}

impl ProcessorAffinityGuard {
    pub fn new(affinity: usize) -> Self {
        set_processor_affinity(affinity);
        Self { old_affinity: 0 }
    }
}

impl Drop for ProcessorAffinityGuard {
    fn drop(&mut self) {
        revert_to_user_affinity();
    }
}

pub struct IrqlGuard {
    old_irql: u8,
}

impl IrqlGuard {
    pub fn new(new_irql: u8) -> Self {
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
    dpc: [u8; 64],
}

impl Dpc {
    pub fn new(routine: DpcRoutine, context: *mut u8) -> Self {
        let mut dpc = Self {
            dpc: [0; 64],
        };

        unsafe {
            KeInitializeDpc(
                dpc.dpc.as_mut_ptr(),
                routine,
                context,
            );
        }

        dpc
    }

    pub fn insert(&mut self, arg1: *mut u8, arg2: *mut u8) -> bool {
        unsafe { KeInsertQueueDpc(self.dpc.as_mut_ptr(), arg1, arg2) }
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
    let active_mask = unsafe { core::ptr::read_volatile(&(*get_active_processor_mask())) };
    (active_mask & affinity) != 0
}

extern "system" {
    static KeActiveProcessors: usize;
}

pub fn get_active_processor_mask() -> *const usize {
    unsafe { &KeActiveProcessors as *const usize }
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
