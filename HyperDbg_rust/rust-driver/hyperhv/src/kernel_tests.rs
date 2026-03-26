use alloc::collections::BTreeMap;
use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;
use crate::*;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TestError {
    TestFailed,
    TestSkipped,
    TestTimeout,
    NotInitialized,
    InvalidParameter,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TestStatus {
    NotRun,
    Running,
    Passed,
    Failed,
    Skipped,
    Timeout,
}

#[derive(Debug, Clone)]
pub struct TestResult {
    pub test_name: String,
    pub status: TestStatus,
    pub duration_ms: u64,
    pub error_message: Option<String>,
    pub assertions_passed: u64,
    pub assertions_failed: u64,
}

impl TestResult {
    pub fn new(test_name: String) -> Self {
        Self {
            test_name,
            status: TestStatus::NotRun,
            duration_ms: 0,
            error_message: None,
            assertions_passed: 0,
            assertions_failed: 0,
        }
    }

    pub fn passed(&mut self, duration_ms: u64) {
        self.status = TestStatus::Passed;
        self.duration_ms = duration_ms;
    }

    pub fn failed(&mut self, duration_ms: u64, error_message: String) {
        self.status = TestStatus::Failed;
        self.duration_ms = duration_ms;
        self.error_message = Some(error_message);
    }

    pub fn skipped(&mut self, duration_ms: u64, reason: String) {
        self.status = TestStatus::Skipped;
        self.duration_ms = duration_ms;
        self.error_message = Some(reason);
    }

    pub fn timeout(&mut self, duration_ms: u64) {
        self.status = TestStatus::Timeout;
        self.duration_ms = duration_ms;
        self.error_message = Some("Test timed out".to_string());
    }
}

#[derive(Debug, Clone)]
pub struct TestSuite {
    pub name: String,
    pub tests: Vec<TestResult>,
    pub total_tests: u64,
    pub passed_tests: u64,
    pub failed_tests: u64,
    pub skipped_tests: u64,
    pub total_duration_ms: u64,
}

impl TestSuite {
    pub fn new(name: String) -> Self {
        Self {
            name,
            tests: Vec::new(),
            total_tests: 0,
            passed_tests: 0,
            failed_tests: 0,
            skipped_tests: 0,
            total_duration_ms: 0,
        }
    }

    pub fn add_test(&mut self, test: TestResult) {
        self.total_tests += 1;
        match test.status {
            TestStatus::Passed => self.passed_tests += 1,
            TestStatus::Failed => self.failed_tests += 1,
            TestStatus::Skipped => self.skipped_tests += 1,
            _ => {}
        }
        self.total_duration_ms += test.duration_ms;
        self.tests.push(test);
    }

    pub fn get_success_rate(&self) -> f64 {
        if self.total_tests == 0 {
            return 0.0;
        }
        (self.passed_tests as f64 / self.total_tests as f64) * 100.0
    }
}

pub struct TestRunner {
    initialized: AtomicBool,
    suites: Mutex<BTreeMap<String, TestSuite>>,
    total_suites: AtomicU64,
    total_tests: AtomicU64,
    total_passed: AtomicU64,
    total_failed: AtomicU64,
    total_skipped: AtomicU64,
    total_duration_ms: AtomicU64,
}

impl TestRunner {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            suites: Mutex::new(BTreeMap::new()),
            total_suites: AtomicU64::new(0),
            total_tests: AtomicU64::new(0),
            total_passed: AtomicU64::new(0),
            total_failed: AtomicU64::new(0),
            total_skipped: AtomicU64::new(0),
            total_duration_ms: AtomicU64::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), TestError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(TestError::NotInitialized);
        }

        self.suites.lock().clear();
        self.total_suites.store(0, Ordering::Release);
        self.total_tests.store(0, Ordering::Release);
        self.total_passed.store(0, Ordering::Release);
        self.total_failed.store(0, Ordering::Release);
        self.total_skipped.store(0, Ordering::Release);
        self.total_duration_ms.store(0, Ordering::Release);
        self.initialized.store(true, Ordering::Release);

        Ok(())
    }

    pub fn deinitialize(&mut self) {
        self.suites.lock().clear();
        self.total_suites.store(0, Ordering::Release);
        self.total_tests.store(0, Ordering::Release);
        self.total_passed.store(0, Ordering::Release);
        self.total_failed.store(0, Ordering::Release);
        self.total_skipped.store(0, Ordering::Release);
        self.total_duration_ms.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn create_suite(&self, name: String) -> Result<(), TestError> {
        if !self.is_initialized() {
            return Err(TestError::NotInitialized);
        }

        let mut suites = self.suites.lock();
        if suites.contains_key(&name) {
            return Err(TestError::InvalidParameter);
        }

        suites.insert(name.clone(), TestSuite::new(name));
        self.total_suites.fetch_add(1, Ordering::AcqRel);

        Ok(())
    }

    pub fn run_test<F>(&self, suite_name: &str, test_name: String, test_func: F) -> Result<(), TestError>
    where
        F: FnOnce() -> Result<(), String>,
    {
        if !self.is_initialized() {
            return Err(TestError::NotInitialized);
        }

        let start_time = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        let mut result = TestResult::new(test_name.clone());
        result.status = TestStatus::Running;

        match test_func() {
            Ok(()) => {
                let end_time = unsafe {
                    let mut rax: u64;
                    core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                    rax
                };
                let duration_ms = (end_time - start_time) / 1_000_000;
                result.passed(duration_ms);
                self.total_passed.fetch_add(1, Ordering::AcqRel);
            }
            Err(error) => {
                let end_time = unsafe {
                    let mut rax: u64;
                    core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
                    rax
                };
                let duration_ms = (end_time - start_time) / 1_000_000;
                result.failed(duration_ms, error);
                self.total_failed.fetch_add(1, Ordering::AcqRel);
            }
        }

        self.total_tests.fetch_add(1, Ordering::AcqRel);
        self.total_duration_ms.fetch_add(result.duration_ms, Ordering::AcqRel);

        let mut suites = self.suites.lock();
        if let Some(suite) = suites.get_mut(suite_name) {
            suite.add_test(result);
        }

        Ok(())
    }

    pub fn skip_test(&self, suite_name: &str, test_name: String, reason: String) -> Result<(), TestError> {
        if !self.is_initialized() {
            return Err(TestError::NotInitialized);
        }

        let start_time = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };

        let mut result = TestResult::new(test_name.clone());
        let end_time = unsafe {
            let mut rax: u64;
            core::arch::asm!("rdtsc", out("rax") rax, out("rdx") _);
            rax
        };
        let duration_ms = (end_time - start_time) / 1_000_000;
        result.skipped(duration_ms, reason);

        self.total_tests.fetch_add(1, Ordering::AcqRel);
        self.total_skipped.fetch_add(1, Ordering::AcqRel);
        self.total_duration_ms.fetch_add(duration_ms, Ordering::AcqRel);

        let mut suites = self.suites.lock();
        if let Some(suite) = suites.get_mut(suite_name) {
            suite.add_test(result);
        }

        Ok(())
    }

    pub fn get_suite(&self, name: &str) -> Option<TestSuite> {
        let suites = self.suites.lock();
        suites.get(name).cloned()
    }

    pub fn get_all_suites(&self) -> Vec<TestSuite> {
        let suites = self.suites.lock();
        suites.values().cloned().collect()
    }

    pub fn get_total_stats(&self) -> (u64, u64, u64, u64, u64) {
        (
            self.total_suites.load(Ordering::Acquire),
            self.total_tests.load(Ordering::Acquire),
            self.total_passed.load(Ordering::Acquire),
            self.total_failed.load(Ordering::Acquire),
            self.total_skipped.load(Ordering::Acquire),
        )
    }

    pub fn get_total_duration_ms(&self) -> u64 {
        self.total_duration_ms.load(Ordering::Acquire)
    }

    pub fn get_overall_success_rate(&self) -> f64 {
        let total = self.total_tests.load(Ordering::Acquire);
        let passed = self.total_passed.load(Ordering::Acquire);

        if total == 0 {
            return 0.0;
        }

        (passed as f64 / total as f64) * 100.0
    }
}

pub static TEST_RUNNER: Mutex<TestRunner> = Mutex::new(TestRunner::new());

pub fn initialize_test_runner() -> Result<(), TestError> {
    let mut runner = TEST_RUNNER.lock();
    runner.initialize()
}

pub fn deinitialize_test_runner() {
    let mut runner = TEST_RUNNER.lock();
    runner.deinitialize();
}

pub fn create_test_suite(name: String) -> Result<(), TestError> {
    let runner = TEST_RUNNER.lock();
    runner.create_suite(name)
}

pub fn run_test<F>(suite_name: &str, test_name: String, test_func: F) -> Result<(), TestError>
where
    F: FnOnce() -> Result<(), String>,
{
    let runner = TEST_RUNNER.lock();
    runner.run_test(suite_name, test_name, test_func)
}

pub fn skip_test(suite_name: &str, test_name: String, reason: String) -> Result<(), TestError> {
    let runner = TEST_RUNNER.lock();
    runner.skip_test(suite_name, test_name, reason)
}

pub fn get_test_suite(name: &str) -> Option<TestSuite> {
    let runner = TEST_RUNNER.lock();
    runner.get_suite(name)
}

pub fn get_all_test_suites() -> Vec<TestSuite> {
    let runner = TEST_RUNNER.lock();
    runner.get_all_suites()
}

pub fn get_test_stats() -> (u64, u64, u64, u64, u64) {
    let runner = TEST_RUNNER.lock();
    runner.get_total_stats()
}

pub fn get_test_duration_ms() -> u64 {
    let runner = TEST_RUNNER.lock();
    runner.get_total_duration_ms()
}

pub fn get_overall_success_rate() -> f64 {
    let runner = TEST_RUNNER.lock();
    runner.get_overall_success_rate()
}

#[macro_export]
macro_rules! assert_eq {
    ($left:expr, $right:expr) => {
        if $left != $right {
            return Err(format!("assertion failed: {} == {}", stringify!($left), stringify!($right)));
        }
    };
}

#[macro_export]
macro_rules! assert_ne {
    ($left:expr, $right:expr) => {
        if $left == $right {
            return Err(format!("assertion failed: {} != {}", stringify!($left), stringify!($right)));
        }
    };
}

#[macro_export]
macro_rules! assert_true {
    ($cond:expr) => {
        if !$cond {
            return Err(format!("assertion failed: {} is true", stringify!($cond)));
        }
    };
}

#[macro_export]
macro_rules! assert_false {
    ($cond:expr) => {
        if $cond {
            return Err(format!("assertion failed: {} is false", stringify!($cond)));
        }
    };
}

#[macro_export]
macro_rules! assert_ok {
    ($expr:expr) => {
        match $expr {
            Ok(_) => {},
            Err(e) => return Err(format!("assertion failed: {} is Ok, got Err: {:?}", stringify!($expr), e)),
        }
    };
}

#[macro_export]
macro_rules! assert_err {
    ($expr:expr) => {
        match $expr {
            Ok(v) => return Err(format!("assertion failed: {} is Err, got Ok: {:?}", stringify!($expr), v)),
            Err(_) => {},
        }
    };
}

pub mod vmx_tests {
    use super::*;

    pub fn test_vmx_support() -> Result<(), String> {
        let cr4 = unsafe { crate::vmm::vmx::read_cr4() };
        assert_true!(cr4 & crate::vmm::vmx::CR4_VMXE != 0);
        Ok(())
    }

    pub fn test_vmx_msr_access() -> Result<(), String> {
        let vmx_basic = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_BASIC) };
        assert_true!(vmx_basic != 0);
        Ok(())
    }

    pub fn test_vmx_cr0_cr4_fixed() -> Result<(), String> {
        let cr0_fixed0 = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_CR0_FIXED0) };
        let cr0_fixed1 = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_CR0_FIXED1) };
        let cr4_fixed0 = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_CR4_FIXED0) };
        let cr4_fixed1 = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_CR4_FIXED1) };

        assert_true!(cr0_fixed0 != 0);
        assert_true!(cr0_fixed1 != 0);
        assert_true!(cr4_fixed0 != 0);
        assert_true!(cr4_fixed1 != 0);
        Ok(())
    }
}

pub mod ept_tests {
    use super::*;

    pub fn test_ept_support() -> Result<(), String> {
        let vmx_ept_vpid_cap = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_EPT_VPID_CAP) };
        assert_true!(vmx_ept_vpid_cap != 0);
        Ok(())
    }

    pub fn test_ept_invept_support() -> Result<(), String> {
        let vmx_ept_vpid_cap = unsafe { crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_VMX_EPT_VPID_CAP) };
        assert_true!(vmx_ept_vpid_cap & (1 << 20) != 0);
        Ok(())
    }
}

pub mod memory_tests {
    use super::*;

    pub fn test_memory_allocation() -> Result<(), String> {
        let size = 0x1000;
        let ptr = unsafe { crate::memory::allocate_non_paged_pool(size) };
        assert_true!(!ptr.is_null());
        unsafe { crate::memory::free_pool(ptr) };
        Ok(())
    }

    pub fn test_physical_address_translation() -> Result<(), String> {
        let va = 0x1000u64;
        let pa = unsafe { crate::memory::virtual_to_physical(va) };
        assert_true!(pa != 0);
        Ok(())
    }
}

pub mod hook_tests {
    use super::*;

    pub fn test_hook_manager_initialization() -> Result<(), String> {
        let _hook_manager = crate::hooks::HookContext {
            ept_hooks: alloc::collections::BTreeMap::new(),
            syscall_hooks: alloc::collections::BTreeMap::new(),
            inline_hooks: alloc::collections::BTreeMap::new(),
            ept_hook2_detours: alloc::collections::BTreeMap::new(),
            ept_hook2_initialized: false,
        };
        Ok(())
    }
}

pub fn run_all_kernel_tests() -> Result<(), TestError> {
    initialize_test_runner()?;
    create_test_suite("VMX Tests".to_string())?;
    create_test_suite("EPT Tests".to_string())?;
    create_test_suite("Memory Tests".to_string())?;
    create_test_suite("Hook Tests".to_string())?;

    run_test("VMX Tests", "test_vmx_support".to_string(), vmx_tests::test_vmx_support)?;
    run_test("VMX Tests", "test_vmx_msr_access".to_string(), vmx_tests::test_vmx_msr_access)?;
    run_test("VMX Tests", "test_vmx_cr0_cr4_fixed".to_string(), vmx_tests::test_vmx_cr0_cr4_fixed)?;

    run_test("EPT Tests", "test_ept_support".to_string(), ept_tests::test_ept_support)?;
    run_test("EPT Tests", "test_ept_invept_support".to_string(), ept_tests::test_ept_invept_support)?;

    run_test("Memory Tests", "test_memory_allocation".to_string(), memory_tests::test_memory_allocation)?;
    run_test("Memory Tests", "test_physical_address_translation".to_string(), memory_tests::test_physical_address_translation)?;

    run_test("Hook Tests", "test_hook_manager_initialization".to_string(), hook_tests::test_hook_manager_initialization)?;

    Ok(())
}
