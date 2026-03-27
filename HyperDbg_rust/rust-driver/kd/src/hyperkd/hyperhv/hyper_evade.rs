use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;
use super::*;
use super::vmm::vmx::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum HyperEvadeError {
    NotInitialized,
    AlreadyEnabled,
    AlreadyDisabled,
    InvalidParameter,
    NotSupported,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum EvadeFeature {
    VmxBasic,
    VmxInstructions,
    Timing,
    CpuId,
    MsrAccess,
    MemoryAccess,
    Interrupts,
    Exceptions,
    All,
}

#[derive(Debug, Clone)]
pub struct SyscallFootprintEntry {
    pub syscall_number: u32,
    pub call_count: u64,
    pub last_call_timestamp: u64,
    pub min_time: u64,
    pub max_time: u64,
    pub total_time: u64,
}

#[derive(Debug, Clone)]
pub struct VmxFootprintEntry {
    pub vmx_instruction: VmxInstruction,
    pub execute_count: u64,
    pub last_execute_timestamp: u64,
    pub min_time: u64,
    pub max_time: u64,
    pub total_time: u64,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum VmxInstruction {
    Vmcall,
    Vmlaunch,
    Vmresume,
    Vmxoff,
    Vmclear,
    Vmptrld,
    Vmptrst,
    Vmread,
    Vmwrite,
    Invept,
    Invvpid,
    Vmfunc,
}

pub struct HyperEvadeManager {
    enabled: AtomicBool,
    initialized: AtomicBool,
    features_enabled: Mutex<BTreeMap<EvadeFeature, bool>>,
    syscall_footprints: Mutex<BTreeMap<u32, SyscallFootprintEntry>>,
    vmx_footprints: Mutex<BTreeMap<VmxInstruction, VmxFootprintEntry>>,
    total_syscalls: AtomicU64,
    total_vmx_instructions: AtomicU64,
}

impl HyperEvadeManager {
    pub fn new() -> Self {
        let mut features_enabled = BTreeMap::new();
        features_enabled.insert(EvadeFeature::VmxBasic, false);
        features_enabled.insert(EvadeFeature::VmxInstructions, false);
        features_enabled.insert(EvadeFeature::Timing, false);
        features_enabled.insert(EvadeFeature::CpuId, false);
        features_enabled.insert(EvadeFeature::MsrAccess, false);
        features_enabled.insert(EvadeFeature::MemoryAccess, false);
        features_enabled.insert(EvadeFeature::Interrupts, false);
        features_enabled.insert(EvadeFeature::Exceptions, false);
        features_enabled.insert(EvadeFeature::All, false);

        Self {
            enabled: AtomicBool::new(false),
            initialized: AtomicBool::new(false),
            features_enabled: Mutex::new(features_enabled),
            syscall_footprints: Mutex::new(BTreeMap::new()),
            vmx_footprints: Mutex::new(BTreeMap::new()),
            total_syscalls: AtomicU64::new(0),
            total_vmx_instructions: AtomicU64::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), HyperEvadeError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(HyperEvadeError::AlreadyEnabled);
        }

        self.syscall_footprints.lock().clear();
        self.vmx_footprints.lock().clear();
        self.total_syscalls.store(0, Ordering::Release);
        self.total_vmx_instructions.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.enabled.store(false, Ordering::Release);
        self.syscall_footprints.lock().clear();
        self.vmx_footprints.lock().clear();
        self.total_syscalls.store(0, Ordering::Release);
        self.total_vmx_instructions.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_enabled(&self) -> bool {
        self.enabled.load(Ordering::Acquire)
    }

    pub fn enable(&self, feature: EvadeFeature) -> Result<(), HyperEvadeError> {
        if !self.is_initialized() {
            return Err(HyperEvadeError::NotInitialized);
        }

        let mut features = self.features_enabled.lock();
        features.insert(feature, true);

        if feature == EvadeFeature::All {
            for (_, enabled) in features.iter_mut() {
                *enabled = true;
            }
        }

        self.enabled.store(true, Ordering::Release);
        Ok(())
    }

    pub fn disable(&self, feature: EvadeFeature) -> Result<(), HyperEvadeError> {
        if !self.is_initialized() {
            return Err(HyperEvadeError::NotInitialized);
        }

        let mut features = self.features_enabled.lock();
        features.insert(feature, false);

        if feature == EvadeFeature::All {
            for (_, enabled) in features.iter_mut() {
                *enabled = false;
            }
        }

        let any_enabled = features.values().any(|&e| e);
        self.enabled.store(any_enabled, Ordering::Release);
        Ok(())
    }

    pub fn is_feature_enabled(&self, feature: EvadeFeature) -> bool {
        let features = self.features_enabled.lock();
        *features.get(&feature).unwrap_or(&false)
    }

    pub fn record_syscall(&self, syscall_number: u32, execution_time: u64) {
        let timestamp = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        let mut footprints = self.syscall_footprints.lock();
        let entry = footprints.entry(syscall_number).or_insert(SyscallFootprintEntry {
            syscall_number,
            call_count: 0,
            last_call_timestamp: 0,
            min_time: u64::MAX,
            max_time: 0,
            total_time: 0,
        });

        entry.call_count += 1;
        entry.last_call_timestamp = timestamp;
        entry.min_time = entry.min_time.min(execution_time);
        entry.max_time = entry.max_time.max(execution_time);
        entry.total_time += execution_time;

        self.total_syscalls.fetch_add(1, Ordering::AcqRel);
    }

    pub fn record_vmx_instruction(&self, instruction: VmxInstruction, execution_time: u64) {
        let timestamp = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        let mut footprints = self.vmx_footprints.lock();
        let entry = footprints.entry(instruction).or_insert(VmxFootprintEntry {
            vmx_instruction: instruction,
            execute_count: 0,
            last_execute_timestamp: 0,
            min_time: u64::MAX,
            max_time: 0,
            total_time: 0,
        });

        entry.execute_count += 1;
        entry.last_execute_timestamp = timestamp;
        entry.min_time = entry.min_time.min(execution_time);
        entry.max_time = entry.max_time.max(execution_time);
        entry.total_time += execution_time;

        self.total_vmx_instructions.fetch_add(1, Ordering::AcqRel);
    }

    pub fn get_syscall_footprint(&self, syscall_number: u32) -> Option<SyscallFootprintEntry> {
        let footprints = self.syscall_footprints.lock();
        footprints.get(&syscall_number).cloned()
    }

    pub fn get_all_syscall_footprints(&self) -> Vec<SyscallFootprintEntry> {
        let footprints = self.syscall_footprints.lock();
        footprints.values().cloned().collect()
    }

    pub fn get_vmx_footprint(&self, instruction: VmxInstruction) -> Option<VmxFootprintEntry> {
        let footprints = self.vmx_footprints.lock();
        footprints.get(&instruction).cloned()
    }

    pub fn get_all_vmx_footprints(&self) -> Vec<VmxFootprintEntry> {
        let footprints = self.vmx_footprints.lock();
        footprints.values().cloned().collect()
    }

    pub fn get_total_syscalls(&self) -> u64 {
        self.total_syscalls.load(Ordering::Acquire)
    }

    pub fn get_total_vmx_instructions(&self) -> u64 {
        self.total_vmx_instructions.load(Ordering::Acquire)
    }

    pub fn clear_footprints(&self) {
        self.syscall_footprints.lock().clear();
        self.vmx_footprints.lock().clear();
        self.total_syscalls.store(0, Ordering::Release);
        self.total_vmx_instructions.store(0, Ordering::Release);
    }
}

pub static HYPER_EVADE_MANAGER: Mutex<HyperEvadeManager> = Mutex::new(HyperEvadeManager::new());

pub fn initialize_hyper_evade() -> Result<(), HyperEvadeError> {
    let mut manager = HYPER_EVADE_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_hyper_evade() {
    let mut manager = HYPER_EVADE_MANAGER.lock();
    manager.deinitialize();
}

pub fn enable_evade_feature(feature: EvadeFeature) -> Result<(), HyperEvadeError> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.enable(feature)
}

pub fn disable_evade_feature(feature: EvadeFeature) -> Result<(), HyperEvadeError> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.disable(feature)
}

pub fn is_evade_feature_enabled(feature: EvadeFeature) -> bool {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.is_feature_enabled(feature)
}

pub fn record_syscall_execution(syscall_number: u32, execution_time: u64) {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.record_syscall(syscall_number, execution_time);
}

pub fn record_vmx_execution(instruction: VmxInstruction, execution_time: u64) {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.record_vmx_instruction(instruction, execution_time);
}

pub fn get_syscall_footprint(syscall_number: u32) -> Option<SyscallFootprintEntry> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.get_syscall_footprint(syscall_number)
}

pub fn get_all_syscall_footprints() -> Vec<SyscallFootprintEntry> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.get_all_syscall_footprints()
}

pub fn get_vmx_footprint(instruction: VmxInstruction) -> Option<VmxFootprintEntry> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.get_vmx_footprint(instruction)
}

pub fn get_all_vmx_footprints() -> Vec<VmxFootprintEntry> {
    let manager = HYPER_EVADE_MANAGER.lock();
    manager.get_all_vmx_footprints()
}
