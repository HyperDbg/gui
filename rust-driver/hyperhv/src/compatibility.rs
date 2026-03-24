use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn ExAllocatePoolWithTag(pool_type: u32, number_of_bytes: usize, tag: u32) -> *mut u8;
    fn ExFreePool(pool: *mut u8);
    fn KeGetCurrentIrql() -> u8;
    fn KeGetCurrentProcessorNumber() -> u32;
    fn KeGetActiveProcessors() -> u32;
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum CompatibilityError {
    NotInitialized,
    VmxNotSupported,
    VmxBugPresent,
    EptNotSupported,
    VpidNotSupported,
    InsufficientProcessors,
    InvalidProcessorCount,
    CheckFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum FeatureStatus {
    NotSupported,
    Supported,
    Unknown,
}

#[derive(Debug, Clone, Copy)]
pub struct CpuFeatures {
    pub vmx_supported: bool,
    pub ept_supported: bool,
    pub vpid_supported: bool,
    pub invept_supported: bool,
    pub invvpid_supported: bool,
    pub vmx_bug_present: bool,
    pub processor_count: u32,
    pub max_physical_address_bits: u8,
    pub max_linear_address_bits: u8,
}

impl CpuFeatures {
    pub fn new() -> Self {
        Self {
            vmx_supported: false,
            ept_supported: false,
            vpid_supported: false,
            invept_supported: false,
            invvpid_supported: false,
            vmx_bug_present: false,
            processor_count: 0,
            max_physical_address_bits: 0,
            max_linear_address_bits: 0,
        }
    }
}

#[derive(Debug, Clone)]
pub struct CompatibilityCheckResult {
    pub vmx_supported: bool,
    pub ept_supported: bool,
    pub vpid_supported: bool,
    pub processor_count_sufficient: bool,
    pub vmx_bug_check_passed: bool,
    pub overall_compatible: bool,
    pub errors: alloc::vec::Vec<CompatibilityError>,
    pub warnings: alloc::vec::Vec<alloc::string::String>,
}

impl CompatibilityCheckResult {
    pub fn new() -> Self {
        Self {
            vmx_supported: false,
            ept_supported: false,
            vpid_supported: false,
            processor_count_sufficient: false,
            vmx_bug_check_passed: false,
            overall_compatible: false,
            errors: alloc::vec::Vec::new(),
            warnings: alloc::vec::Vec::new(),
        }
    }

    pub fn add_error(&mut self, error: CompatibilityError) {
        self.errors.push(error);
    }

    pub fn add_warning(&mut self, warning: alloc::string::String) {
        self.warnings.push(warning);
    }

    pub fn is_compatible(&self) -> bool {
        self.overall_compatible && self.errors.is_empty()
    }
}

pub struct CompatibilityChecker {
    initialized: AtomicBool,
    cpu_features: Mutex<Box<CpuFeatures>>,
    check_result: Mutex<Box<CompatibilityCheckResult>>,
}

impl CompatibilityChecker {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            cpu_features: Mutex::new(Box::new(CpuFeatures::new())),
            check_result: Mutex::new(Box::new(CompatibilityCheckResult::new())),
        }
    }

    pub fn initialize(&mut self) -> Result<(), CompatibilityError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(CompatibilityError::NotInitialized);
        }

        let mut features = self.cpu_features.lock();
        self.detect_cpu_features(&mut features);

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn check_compatibility(&self) -> CompatibilityCheckResult {
        if !self.is_initialized() {
            let mut result = CompatibilityCheckResult::new();
            result.add_error(CompatibilityError::NotInitialized);
            return result;
        }

        let features = self.cpu_features.lock();
        let mut result = CompatibilityCheckResult::new();

        result.vmx_supported = features.vmx_supported;
        result.ept_supported = features.ept_supported;
        result.vpid_supported = features.vpid_supported;
        result.processor_count_sufficient = features.processor_count >= 2;
        result.vmx_bug_check_passed = !features.vmx_bug_present;

        if !features.vmx_supported {
            result.add_error(CompatibilityError::VmxNotSupported);
        }

        if !features.ept_supported {
            result.add_error(CompatibilityError::EptNotSupported);
        }

        if !features.vpid_supported {
            result.add_warning(alloc::string::String::from("VPID not supported, performance may be degraded"));
        }

        if features.processor_count < 2 {
            result.add_error(CompatibilityError::InsufficientProcessors);
        }

        if features.vmx_bug_present {
            result.add_error(CompatibilityError::VmxBugPresent);
        }

        result.overall_compatible = result.vmx_supported
            && result.ept_supported
            && result.processor_count_sufficient
            && result.vmx_bug_check_passed;

        *self.check_result.lock() = Box::new(result.clone());
        result
    }

    pub fn get_cpu_features(&self) -> CpuFeatures {
        *self.cpu_features.lock()
    }

    pub fn get_last_check_result(&self) -> CompatibilityCheckResult {
        (*self.check_result.lock()).clone()
    }

    pub fn is_vmx_supported(&self) -> bool {
        self.cpu_features.lock().vmx_supported
    }

    pub fn is_ept_supported(&self) -> bool {
        self.cpu_features.lock().ept_supported
    }

    pub fn is_vpid_supported(&self) -> bool {
        self.cpu_features.lock().vpid_supported
    }

    pub fn get_processor_count(&self) -> u32 {
        self.cpu_features.lock().processor_count
    }

    unsafe fn detect_cpu_features(&self, features: &mut CpuFeatures) {
        features.processor_count = KeGetActiveProcessors();

        let mut eax: u32;
        let mut ebx: u32;
        let mut ecx: u32;
        let mut edx: u32;

        core::arch::asm!(
            "cpuid",
            inlateout("eax") 1 => eax,
            out("ebx") ebx,
            out("ecx") ecx,
            out("edx") edx,
            options(nomem, nostack),
        );

        features.vmx_supported = (ecx & (1 << 5)) != 0;

        if features.vmx_supported {
            let vmx_basic = self.read_msr(0x480);

            let vmx_ept_vpid_cap = self.read_msr(0x48C);

            features.ept_supported = (vmx_ept_vpid_cap & (1 << 1)) != 0;
            features.vpid_supported = (vmx_ept_vpid_cap & (1 << 0)) != 0;
            features.invept_supported = (vmx_ept_vpid_cap & (1 << 20)) != 0;
            features.invvpid_supported = (vmx_ept_vpid_cap & (1 << 21)) != 0;

            let vmx_cr0_fixed0 = self.read_msr(0x486);
            let vmx_cr0_fixed1 = self.read_msr(0x487);

            features.vmx_bug_present = (vmx_cr0_fixed0 & vmx_cr0_fixed1) == 0;

            let addr_size = (vmx_basic >> 48) & 0xFF;
            features.max_physical_address_bits = (addr_size & 0x7F) as u8;
            features.max_linear_address_bits = ((addr_size >> 8) & 0x7F) as u8;
        }
    }

    unsafe fn read_msr(&self, msr: u32) -> u64 {
        let low: u32;
        let high: u32;
        core::arch::asm!(
            "rdmsr",
            inlateout("ecx") msr => _,
            out("eax") low,
            out("edx") high,
            options(nomem, nostack, pure),
        );
        ((high as u64) << 32) | (low as u64)
    }
}

pub static COMPATIBILITY_CHECKER: Mutex<CompatibilityChecker> = Mutex::new(CompatibilityChecker::new());

pub fn initialize_compatibility_checker() -> Result<(), CompatibilityError> {
    let mut checker = COMPATIBILITY_CHECKER.lock();
    checker.initialize()
}

pub fn deinitialize_compatibility_checker() {
    let mut checker = COMPATIBILITY_CHECKER.lock();
    checker.deinitialize();
}

pub fn check_compatibility() -> CompatibilityCheckResult {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.check_compatibility()
}

pub fn get_cpu_features() -> CpuFeatures {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.get_cpu_features()
}

pub fn is_vmx_supported() -> bool {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.is_vmx_supported()
}

pub fn is_ept_supported() -> bool {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.is_ept_supported()
}

pub fn is_vpid_supported() -> bool {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.is_vpid_supported()
}

pub fn get_processor_count() -> u32 {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.get_processor_count()
}

pub fn get_last_check_result() -> CompatibilityCheckResult {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.get_last_check_result()
}

pub fn is_compatibility_checker_initialized() -> bool {
    let checker = COMPATIBILITY_CHECKER.lock();
    checker.is_initialized()
}
