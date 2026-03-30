use alloc::collections::BTreeMap;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LbrError {
    NotInitialized,
    NotSupported,
    InvalidParameter,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LbrFormat {
    Legacy,
    Architectural,
}

#[derive(Debug, Clone, Copy)]
pub struct LbrEntry {
    pub from: u64,
    pub to: u64,
    pub timestamp: u64,
    pub mispredicted: bool,
    pub cycles: u64,
}

impl LbrEntry {
    pub fn new() -> Self {
        Self {
            from: 0,
            to: 0,
            timestamp: 0,
            mispredicted: false,
            cycles: 0,
        }
    }
}

#[derive(Debug, Clone, Copy)]
pub struct LbrStack {
    pub entries: [LbrEntry; 32],
    pub top: usize,
    pub depth: usize,
}

impl LbrStack {
    pub fn new() -> Self {
        Self {
            entries: [LbrEntry::new(); 32],
            top: 0,
            depth: 0,
        }
    }

    pub fn push(&mut self, entry: LbrEntry) {
        if self.depth < 32 {
            self.entries[self.top] = entry;
            self.top = (self.top + 1) % 32;
            self.depth = self.depth.min(31) + 1;
        }
    }

    pub fn pop(&mut self) -> Option<LbrEntry> {
        if self.depth == 0 {
            return None;
        }

        self.top = if self.top == 0 { 31 } else { self.top - 1 };
        self.depth -= 1;
        Some(self.entries[self.top])
    }

    pub fn clear(&mut self) {
        self.top = 0;
        self.depth = 0;
        for i in 0..32 {
            self.entries[i] = LbrEntry::new();
        }
    }

    pub fn get_entries(&self) -> &[LbrEntry] {
        &self.entries[..self.depth]
    }
}

pub struct LbrManager {
    enabled: AtomicBool,
    initialized: AtomicBool,
    format: Mutex<LbrFormat>,
    stacks: Mutex<BTreeMap<u32, LbrStack>>,
    max_depth: AtomicU64,
    total_branches: AtomicU64,
    total_mispredictions: AtomicU64,
}

impl LbrManager {
    pub const fn new() -> Self {
        Self {
            enabled: AtomicBool::new(false),
            initialized: AtomicBool::new(false),
            format: Mutex::new(LbrFormat::Architectural),
            stacks: Mutex::new(BTreeMap::new()),
            max_depth: AtomicU64::new(32),
            total_branches: AtomicU64::new(0),
            total_mispredictions: AtomicU64::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), LbrError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(LbrError::NotInitialized);
        }

        if !self.check_lbr_support() {
            return Err(LbrError::NotSupported);
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.enabled.store(false, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_enabled(&self) -> bool {
        self.enabled.load(Ordering::Acquire)
    }

    pub fn enable(&self) -> Result<(), LbrError> {
        if !self.is_initialized() {
            return Err(LbrError::NotInitialized);
        }

        self.enable_lbr_hardware();
        self.enabled.store(true, Ordering::Release);
        Ok(())
    }

    pub fn disable(&self) -> Result<(), LbrError> {
        if !self.is_initialized() {
            return Err(LbrError::NotInitialized);
        }

        self.disable_lbr_hardware();
        self.enabled.store(false, Ordering::Release);
        Ok(())
    }

    pub fn set_format(&self, format: LbrFormat) -> Result<(), LbrError> {
        if !self.is_initialized() {
            return Err(LbrError::NotInitialized);
        }

        let mut fmt = self.format.lock();
        *fmt = format;
        Ok(())
    }

    pub fn get_format(&self) -> LbrFormat {
        let fmt = self.format.lock();
        *fmt
    }

    pub fn set_max_depth(&self, depth: u64) -> Result<(), LbrError> {
        if depth == 0 || depth > 32 {
            return Err(LbrError::InvalidParameter);
        }

        self.max_depth.store(depth, Ordering::Release);
        Ok(())
    }

    pub fn get_max_depth(&self) -> u64 {
        self.max_depth.load(Ordering::Acquire)
    }

    pub fn record_branch(&self, core_id: u32, from: u64, to: u64, mispredicted: bool, cycles: u64) {
        if !self.is_enabled() {
            return;
        }

        let timestamp = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        let entry = LbrEntry {
            from,
            to,
            timestamp,
            mispredicted,
            cycles,
        };

        let mut stacks = self.stacks.lock();
        let stack = stacks.entry(core_id).or_insert_with(LbrStack::new);
        stack.push(entry);

        self.total_branches.fetch_add(1, Ordering::AcqRel);
        if mispredicted {
            self.total_mispredictions.fetch_add(1, Ordering::AcqRel);
        }
    }

    pub fn get_stack(&self, core_id: u32) -> Option<LbrStack> {
        let stacks = self.stacks.lock();
        stacks.get(&core_id).cloned()
    }

    pub fn clear_stack(&self, core_id: u32) {
        let mut stacks = self.stacks.lock();
        if let Some(stack) = stacks.get_mut(&core_id) {
            stack.clear();
        }
    }

    pub fn clear_all_stacks(&self) {
        let mut stacks = self.stacks.lock();
        for (_, stack) in stacks.iter_mut() {
            stack.clear();
        }
    }

    pub fn get_total_branches(&self) -> u64 {
        self.total_branches.load(Ordering::Acquire)
    }

    pub fn get_total_mispredictions(&self) -> u64 {
        self.total_mispredictions.load(Ordering::Acquire)
    }

    pub fn get_misprediction_rate(&self) -> f64 {
        let total = self.get_total_branches();
        if total == 0 {
            return 0.0;
        }

        let mispred = self.get_total_mispredictions();
        (mispred as f64) / (total as f64)
    }

    fn check_lbr_support(&self) -> bool {
        unsafe {
            let mut eax: u32 = 0x80000001;
            let mut ebx: u32;
            let mut ecx: u32;
            let mut edx: u32;

            core::arch::asm!(
                "push rbx",
                "cpuid",
                "mov {0}, ebx",
                "pop rbx",
                out(reg) ebx,
                inout("eax") eax,
                out("ecx") ecx,
                out("edx") edx,
            );

            (edx & (1 << 0)) != 0
        }
    }

    fn enable_lbr_hardware(&self) {
        unsafe {
            let mut cr4: u64;
            core::arch::asm!("mov {}, cr4", out(reg) cr4);
            cr4 |= 1 << 8 ;
            core::arch::asm!("mov cr4, {}", in(reg) cr4);

            let mut debugctl: u64;
            core::arch::asm!("mov {}, msr 0x1d9", out(reg) debugctl, options(nomem, nostack));
            debugctl |= 1 << 0 ;
            core::arch::asm!("mov msr 0x1d9, {}", in(reg) debugctl, options(nomem, nostack));
        }
    }

    fn disable_lbr_hardware(&self) {
        unsafe {
            let mut debugctl: u64;
            core::arch::asm!("mov {}, msr 0x1d9", out(reg) debugctl, options(nomem, nostack));
            debugctl &= !(1 << 0);
            core::arch::asm!("mov msr 0x1d9, {}", in(reg) debugctl, options(nomem, nostack));
        }
    }
}

static LBR_MANAGER: Mutex<LbrManager> = Mutex::new(LbrManager::new());

pub fn initialize_lbr() -> Result<(), LbrError> {
    let mut manager = LBR_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_lbr() {
    let mut manager = LBR_MANAGER.lock();
    manager.deinitialize();
}

pub fn enable_lbr() -> Result<(), LbrError> {
    let manager = LBR_MANAGER.lock();
    manager.enable()
}

pub fn disable_lbr() -> Result<(), LbrError> {
    let manager = LBR_MANAGER.lock();
    manager.disable()
}

pub fn record_branch(core_id: u32, from: u64, to: u64, mispredicted: bool, cycles: u64) {
    let manager = LBR_MANAGER.lock();
    manager.record_branch(core_id, from, to, mispredicted, cycles);
}

pub fn get_lbr_stack(core_id: u32) -> Option<LbrStack> {
    let manager = LBR_MANAGER.lock();
    manager.get_stack(core_id)
}

pub fn clear_lbr_stack(core_id: u32) {
    let manager = LBR_MANAGER.lock();
    manager.clear_stack(core_id);
}

pub fn clear_all_lbr_stacks() {
    let manager = LBR_MANAGER.lock();
    manager.clear_all_stacks();
}

pub fn get_total_branches() -> u64 {
    let manager = LBR_MANAGER.lock();
    manager.get_total_branches()
}

pub fn get_total_mispredictions() -> u64 {
    let manager = LBR_MANAGER.lock();
    manager.get_total_mispredictions()
}

pub fn get_misprediction_rate() -> f64 {
    let manager = LBR_MANAGER.lock();
    manager.get_misprediction_rate()
}