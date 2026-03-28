use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::vmm::vmx::*;
use crate::hyperkd::hyperhv::common::msr::read_msr;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TransparencyError {
    NotInitialized,
    AlreadyEnabled,
    AlreadyDisabled,
    InvalidParameter,
    NotSupported,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
pub enum TransparencyFeature {
    HideVmxFromCpuid,
    HideVmxFromMsr,
    HideVmxFromRdtsc,
    HideVmxFromInterrupts,
    HideVmxFromExceptions,
    HideVmxFromMemory,
    HideVmxFromTiming,
    HideVmxFromDebugRegisters,
    All,
}

#[derive(Debug, Clone)]
pub struct TimingAdjustment {
    pub base_overhead: u64,
    pub adjustment_factor: f64,
    pub enabled: bool,
}

#[derive(Debug, Clone)]
pub struct CpuidMask {
    pub leaf: u32,
    pub sub_leaf: u32,
    pub register_mask: u32,
    pub masked_value: u32,
    pub original_value: u32,
}

#[derive(Debug, Clone)]
pub struct MsrMask {
    pub msr: u32,
    pub mask: u64,
    pub masked_value: u64,
    pub original_value: u64,
}

pub struct TransparencyManager {
    enabled: AtomicBool,
    initialized: AtomicBool,
    features_enabled: Mutex<BTreeMap<TransparencyFeature, bool>>,
    cpuid_masks: Mutex<BTreeMap<(u32, u32), CpuidMask>>,
    msr_masks: Mutex<BTreeMap<u32, MsrMask>>,
    timing_adjustment: Mutex<TimingAdjustment>,
    total_adjustments: AtomicU64,
}

impl TransparencyManager {
    pub const fn new() -> Self {
        Self {
            enabled: AtomicBool::new(false),
            initialized: AtomicBool::new(false),
            features_enabled: Mutex::new(BTreeMap::new()),
            cpuid_masks: Mutex::new(BTreeMap::new()),
            msr_masks: Mutex::new(BTreeMap::new()),
            timing_adjustment: Mutex::new(TimingAdjustment {
                base_overhead: 0,
                adjustment_factor: 1.0,
                enabled: false,
            }),
            total_adjustments: AtomicU64::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), TransparencyError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(TransparencyError::AlreadyEnabled);
        }

        {
            let mut features_enabled = self.features_enabled.lock();
            features_enabled.insert(TransparencyFeature::HideVmxFromCpuid, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromMsr, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromRdtsc, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromInterrupts, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromExceptions, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromMemory, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromTiming, false);
            features_enabled.insert(TransparencyFeature::HideVmxFromDebugRegisters, false);
            features_enabled.insert(TransparencyFeature::All, false);
        }

        self.cpuid_masks.lock().clear();
        self.msr_masks.lock().clear();
        self.total_adjustments.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.enabled.store(false, Ordering::Release);
        self.cpuid_masks.lock().clear();
        self.msr_masks.lock().clear();
        self.total_adjustments.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn is_enabled(&self) -> bool {
        self.enabled.load(Ordering::Acquire)
    }

    pub fn enable(&self, feature: TransparencyFeature) -> Result<(), TransparencyError> {
        if !self.is_initialized() {
            return Err(TransparencyError::NotInitialized);
        }

        let mut features = self.features_enabled.lock();
        features.insert(feature, true);

        if feature == TransparencyFeature::All {
            for (_, enabled) in features.iter_mut() {
                *enabled = true;
            }
        }

        self.enabled.store(true, Ordering::Release);
        Ok(())
    }

    pub fn disable(&self, feature: TransparencyFeature) -> Result<(), TransparencyError> {
        if !self.is_initialized() {
            return Err(TransparencyError::NotInitialized);
        }

        let mut features = self.features_enabled.lock();
        features.insert(feature, false);

        if feature == TransparencyFeature::All {
            for (_, enabled) in features.iter_mut() {
                *enabled = false;
            }
        }

        let any_enabled = features.values().any(|&e| e);
        self.enabled.store(any_enabled, Ordering::Release);
        Ok(())
    }

    pub fn is_feature_enabled(&self, feature: TransparencyFeature) -> bool {
        let features = self.features_enabled.lock();
        *features.get(&feature).unwrap_or(&false)
    }

    pub unsafe fn add_cpuid_mask(&self, leaf: u32, sub_leaf: u32, register_mask: u32, masked_value: u32) -> Result<(), TransparencyError> {
        if !self.is_initialized() {
            return Err(TransparencyError::NotInitialized);
        }

        let mut cpuid_masks = self.cpuid_masks.lock();
        
        let original_value = self.get_cpuid_value(leaf, sub_leaf, register_mask);
        
        cpuid_masks.insert((leaf, sub_leaf), CpuidMask {
            leaf,
            sub_leaf,
            register_mask,
            masked_value,
            original_value,
        });

        self.total_adjustments.fetch_add(1, Ordering::AcqRel);
        Ok(())
    }

    pub unsafe fn remove_cpuid_mask(&self, leaf: u32, sub_leaf: u32) -> Option<CpuidMask> {
        let mut cpuid_masks = self.cpuid_masks.lock();
        cpuid_masks.remove(&(leaf, sub_leaf))
    }

    pub unsafe fn get_masked_cpuid_value(&self, leaf: u32, sub_leaf: u32, register_mask: u32) -> u32 {
        let cpuid_masks = self.cpuid_masks.lock();
        
        if let Some(mask) = cpuid_masks.get(&(leaf, sub_leaf)) {
            if mask.register_mask == register_mask {
                return mask.masked_value;
            }
        }

        self.get_cpuid_value(leaf, sub_leaf, register_mask)
    }

    unsafe fn get_cpuid_value(&self, leaf: u32, sub_leaf: u32, register_mask: u32) -> u32 {
        let mut eax: u32 = 0;
        let mut ebx: u32 = 0;
        let mut ecx: u32 = sub_leaf;
        let mut edx: u32 = 0;

        core::arch::asm!(
            "push rbx",
            "cpuid",
            "mov {0}, ebx",
            "pop rbx",
            out(reg) ebx,
            inout("eax") leaf => eax,
            inout("ecx") ecx,
            inlateout("edx") 0 => edx,
        );

        match register_mask {
            0 => eax,
            1 => ebx,
            2 => ecx,
            3 => edx,
            _ => eax,
        }
    }

    pub unsafe fn add_msr_mask(&self, msr: u32, mask: u64, masked_value: u64) -> Result<(), TransparencyError> {
        if !self.is_initialized() {
            return Err(TransparencyError::NotInitialized);
        }

        let mut msr_masks = self.msr_masks.lock();
        
        let original_value = read_msr(msr);
        
        msr_masks.insert(msr, MsrMask {
            msr,
            mask,
            masked_value,
            original_value,
        });

        self.total_adjustments.fetch_add(1, Ordering::AcqRel);
        Ok(())
    }

    pub unsafe fn remove_msr_mask(&self, msr: u32) -> Option<MsrMask> {
        let mut msr_masks = self.msr_masks.lock();
        msr_masks.remove(&msr)
    }

    pub unsafe fn get_masked_msr_value(&self, msr: u32) -> u64 {
        let msr_masks = self.msr_masks.lock();
        
        if let Some(mask) = msr_masks.get(&msr) {
            return mask.masked_value;
        }

        read_msr(msr)
    }

    pub fn set_timing_adjustment(&self, base_overhead: u64, adjustment_factor: f64) -> Result<(), TransparencyError> {
        if !self.is_initialized() {
            return Err(TransparencyError::NotInitialized);
        }

        let mut timing = self.timing_adjustment.lock();
        timing.base_overhead = base_overhead;
        timing.adjustment_factor = adjustment_factor;
        timing.enabled = true;

        Ok(())
    }

    pub fn get_adjusted_tsc(&self) -> u64 {
        let timing = self.timing_adjustment.lock();
        
        if !timing.enabled {
            return unsafe {
                let mut rax: u64;
                let mut rdx: u64;
                core::arch::asm!("rdtsc", out("rax") rax, out("rdx") rdx);
                (rdx << 32) | rax
            };
        }

        let raw_tsc = unsafe {
            let mut rax: u64;
            let mut rdx: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") rdx);
            (rdx << 32) | rax
        };

        let adjusted = (raw_tsc as f64 * timing.adjustment_factor) as u64;
        adjusted.saturating_sub(timing.base_overhead)
    }

    pub fn get_all_cpuid_masks(&self) -> Vec<CpuidMask> {
        let cpuid_masks = self.cpuid_masks.lock();
        cpuid_masks.values().cloned().collect()
    }

    pub fn get_all_msr_masks(&self) -> Vec<MsrMask> {
        let msr_masks = self.msr_masks.lock();
        msr_masks.values().cloned().collect()
    }

    pub fn get_total_adjustments(&self) -> u64 {
        self.total_adjustments.load(Ordering::Acquire)
    }

    pub fn clear_all_masks(&self) {
        self.cpuid_masks.lock().clear();
        self.msr_masks.lock().clear();
        self.total_adjustments.store(0, Ordering::Release);
    }
}

pub static TRANSPARENCY_MANAGER: Mutex<TransparencyManager> = Mutex::new(TransparencyManager::new());

pub fn initialize_transparency() -> Result<(), TransparencyError> {
    let mut manager = TRANSPARENCY_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_transparency() {
    let mut manager = TRANSPARENCY_MANAGER.lock();
    manager.deinitialize();
}

pub fn enable_transparency_feature(feature: TransparencyFeature) -> Result<(), TransparencyError> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.enable(feature)
}

pub fn disable_transparency_feature(feature: TransparencyFeature) -> Result<(), TransparencyError> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.disable(feature)
}

pub fn is_transparency_feature_enabled(feature: TransparencyFeature) -> bool {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.is_feature_enabled(feature)
}

pub unsafe fn add_cpuid_mask(leaf: u32, sub_leaf: u32, register_mask: u32, masked_value: u32) -> Result<(), TransparencyError> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.add_cpuid_mask(leaf, sub_leaf, register_mask, masked_value)
}

pub unsafe fn remove_cpuid_mask(leaf: u32, sub_leaf: u32) -> Option<CpuidMask> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.remove_cpuid_mask(leaf, sub_leaf)
}

pub unsafe fn get_masked_cpuid_value(leaf: u32, sub_leaf: u32, register_mask: u32) -> u32 {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.get_masked_cpuid_value(leaf, sub_leaf, register_mask)
}

pub unsafe fn add_msr_mask(msr: u32, mask: u64, masked_value: u64) -> Result<(), TransparencyError> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.add_msr_mask(msr, mask, masked_value)
}

pub unsafe fn remove_msr_mask(msr: u32) -> Option<MsrMask> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.remove_msr_mask(msr)
}

pub unsafe fn get_masked_msr_value(msr: u32) -> u64 {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.get_masked_msr_value(msr)
}

pub fn set_timing_adjustment(base_overhead: u64, adjustment_factor: f64) -> Result<(), TransparencyError> {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.set_timing_adjustment(base_overhead, adjustment_factor)
}

pub fn get_adjusted_tsc() -> u64 {
    let manager = TRANSPARENCY_MANAGER.lock();
    manager.get_adjusted_tsc()
}
