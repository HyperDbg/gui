use core::sync::atomic::{AtomicBool, Ordering};
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TracingError {
    NotInitialized,
    InvalidParameter,
    OperationFailed,
    InvalidState,
    AccessDenied,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum TracingMode {
    Disabled,
    RegularStepIn,
    InstrumentationStepIn,
    StepOver,
    StepOut,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct TracingContext {
    pub tracing_mode: TracingMode,
    pub is_tracing_active: bool,
    pub current_rip: u64,
    pub step_count: u64,
    pub max_steps: u64,
    pub break_on_return: bool,
    pub return_address: u64,
    pub trap_flag_masked: bool,
    pub interruptibility_state: u64,
}

impl TracingContext {
    pub fn new() -> Self {
        Self {
            tracing_mode: TracingMode::Disabled,
            is_tracing_active: false,
            current_rip: 0,
            step_count: 0,
            max_steps: 0,
            break_on_return: false,
            return_address: 0,
            trap_flag_masked: false,
            interruptibility_state: 0,
        }
    }

    pub fn is_active(&self) -> bool {
        self.is_tracing_active && self.tracing_mode != TracingMode::Disabled
    }

    pub fn should_continue(&self) -> bool {
        if self.max_steps == 0 {
            return true;
        }

        self.step_count < self.max_steps
    }

    pub fn should_break_on_return(&self) -> bool {
        self.break_on_return && self.current_rip == self.return_address
    }
}

pub struct TracingManager {
    contexts: alloc::vec::Vec<TracingContext>,
    processor_count: u32,
    initialized: AtomicBool,
}

impl TracingManager {
    pub const fn new(processor_count: u32) -> Self {
        Self {
            contexts: alloc::vec::Vec::new(),
            processor_count,
            initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&mut self) -> Result<(), TracingError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(TracingError::NotInitialized);
        }

        for _ in 0..self.processor_count {
            self.contexts.push(TracingContext::new());
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        for context in &mut self.contexts {
            context.tracing_mode = TracingMode::Disabled;
            context.is_tracing_active = false;
            context.step_count = 0;
        }
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn get_context(&self, core_id: u32) -> Result<&TracingContext, TracingError> {
        if !self.is_initialized() {
            return Err(TracingError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(TracingError::InvalidParameter);
        }

        Ok(&self.contexts[core_id as usize])
    }

    pub fn get_context_mut(&mut self, core_id: u32) -> Result<&mut TracingContext, TracingError> {
        if !self.is_initialized() {
            return Err(TracingError::NotInitialized);
        }

        if core_id >= self.processor_count {
            return Err(TracingError::InvalidParameter);
        }

        Ok(&mut self.contexts[core_id as usize])
    }

    pub fn start_tracing(
        &mut self,
        core_id: u32,
        mode: TracingMode,
        max_steps: u64,
    ) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;

        context.tracing_mode = mode;
        context.is_tracing_active = true;
        context.step_count = 0;
        context.max_steps = max_steps;
        context.break_on_return = false;
        context.return_address = 0;

        unsafe {
            match mode {
                TracingMode::InstrumentationStepIn => {
                    self.perform_instrumentation_step_in(core_id)?;
                }
                TracingMode::RegularStepIn => {
                    self.perform_regular_step_in(core_id)?;
                }
                TracingMode::StepOver => {
                    self.perform_step_over(core_id)?;
                }
                TracingMode::StepOut => {
                    self.perform_step_out(core_id)?;
                }
                TracingMode::Disabled => {
                    return Err(TracingError::InvalidState);
                }
            }
        }

        Ok(())
    }

    pub fn stop_tracing(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;

        context.tracing_mode = TracingMode::Disabled;
        context.is_tracing_active = false;
        context.step_count = 0;

        unsafe { self.restore_system_state(core_id)?; }

        Ok(())
    }

    pub fn handle_step(&mut self, core_id: u32, rip: u64) -> Result<bool, TracingError> {
        let context = self.get_context_mut(core_id)?;

        if !context.is_active() {
            return Ok(false);
        }

        context.current_rip = rip;
        context.step_count += 1;

        if context.should_break_on_return() {
            self.stop_tracing(core_id)?;
            return Ok(true);
        }

        if !context.should_continue() {
            self.stop_tracing(core_id)?;
            return Ok(true);
        }

        unsafe {
            match context.tracing_mode {
                TracingMode::InstrumentationStepIn => {
                    self.handle_mtf(core_id)?;
                }
                TracingMode::RegularStepIn => {
                    self.check_for_continuing_steps(core_id)?;
                }
                TracingMode::StepOver => {
                    self.handle_step_over(core_id)?;
                }
                TracingMode::StepOut => {
                    self.handle_step_out(core_id)?;
                }
                TracingMode::Disabled => {
                    return Ok(false);
                }
            }
        }

        Ok(false)
    }

    unsafe fn perform_instrumentation_step_in(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;
        
        let last_vmexit_rip = vm_func_get_last_vmexit_rip(core_id);
        disassembler_show_one_instruction_in_vmx_root_mode(last_vmexit_rip, false);

        context.tracing_mode = TracingMode::InstrumentationStepIn;

        vm_func_register_mtf_break(core_id);
        vm_func_enable_mtf_and_change_external_interrupt_state(core_id);

        Ok(())
    }

    unsafe fn perform_regular_step_in(&mut self, core_id: u32) -> Result<(), TracingError> {
        let mut interruptibility: u64 = 0;
        let mut interruptibility_old: u64 = 0;

        vm_func_set_rflag_trap_flag(true);

        interruptibility_old = vm_func_get_interruptibility_state();
        interruptibility = interruptibility_old;

        interruptibility = vm_func_clear_stepping_bits(interruptibility);

        if interruptibility != interruptibility_old {
            vm_func_set_interruptibility_state(interruptibility);
        }

        let context = self.get_context_mut(core_id)?;
        context.interruptibility_state = interruptibility_old;

        Ok(())
    }

    unsafe fn perform_step_over(&mut self, core_id: u32) -> Result<(), TracingError> {
        let last_vmexit_rip = vm_func_get_last_vmexit_rip(core_id);
        let instruction_bytes = self.read_instruction_bytes(last_vmexit_rip)?;

        let is_call = self.is_call_instruction(&instruction_bytes);
        let next_rip = if is_call {
            Some(self.get_next_instruction_address(last_vmexit_rip, &instruction_bytes)?)
        } else {
            None
        };

        let context = self.get_context_mut(core_id)?;

        if let Some(rip) = next_rip {
            context.break_on_return = true;
            context.return_address = rip;
            self.set_temporary_breakpoint(core_id, rip)?;
        } else {
            self.perform_regular_step_in(core_id)?;
        }

        Ok(())
    }

    unsafe fn perform_step_out(&mut self, core_id: u32) -> Result<(), TracingError> {
        let last_vmexit_rip = vm_func_get_last_vmexit_rip(core_id);
        let return_address = self.get_return_address(last_vmexit_rip)?;

        let context = self.get_context_mut(core_id)?;
        context.break_on_return = true;
        context.return_address = return_address;

        self.set_temporary_breakpoint(core_id, return_address)?;

        Ok(())
    }

    unsafe fn handle_mtf(&mut self, core_id: u32) -> Result<(), TracingError> {
        self.restore_system_state(core_id)?;

        meta_dispatch_event_instrumentation_trace(core_id)?;

        self.check_for_continuing_steps(core_id)?;

        Ok(())
    }

    unsafe fn restore_system_state(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;

        context.tracing_mode = TracingMode::Disabled;

        vm_func_unregister_mtf_break(core_id);
        vm_func_enable_and_check_for_previous_external_interrupts(core_id);

        Ok(())
    }

    unsafe fn check_for_continuing_steps(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context(core_id)?;

        if context.is_active() {
            vm_func_change_mtf_unsetting_state(core_id, true);
        }

        Ok(())
    }

    unsafe fn handle_step_over(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;

        if context.break_on_return && context.current_rip == context.return_address {
            context.break_on_return = false;
            context.return_address = 0;
            self.remove_temporary_breakpoint(core_id)?;
        }

        Ok(())
    }

    unsafe fn handle_step_out(&mut self, core_id: u32) -> Result<(), TracingError> {
        let context = self.get_context_mut(core_id)?;

        if context.break_on_return && context.current_rip == context.return_address {
            context.break_on_return = false;
            context.return_address = 0;
            self.remove_temporary_breakpoint(core_id)?;
            self.stop_tracing(core_id)?;
        }

        Ok(())
    }

    unsafe fn read_instruction_bytes(&self, address: u64) -> Result<[u8; 15], TracingError> {
        extern "C" {
            fn mmlayout_read_memory_safe(address: u64, buffer: *mut u8, size: usize) -> bool;
        }

        let mut bytes = [0u8; 15];
        let success = mmlayout_read_memory_safe(address, bytes.as_mut_ptr(), 15);

        if success {
            Ok(bytes)
        } else {
            Err(TracingError::AccessDenied)
        }
    }

    fn is_call_instruction(&self, bytes: &[u8; 15]) -> bool {
        bytes[0] == 0xE8 || 
        (bytes[0] == 0xFF && bytes[1] & 0x38 == 0x10) ||
        (bytes[0] == 0xFF && bytes[1] & 0x38 == 0x18)
    }

    fn get_next_instruction_address(&self, rip: u64, bytes: &[u8; 15]) -> Result<u64, TracingError> {
        if bytes[0] == 0xE8 {
            let offset = i32::from_le_bytes([bytes[1], bytes[2], bytes[3], bytes[4]]);
            Ok(rip + offset as u64 + 5)
        } else if bytes[0] == 0xFF && bytes[1] & 0x38 == 0x10 {
            Ok(rip + 2)
        } else if bytes[0] == 0xFF && bytes[1] & 0x38 == 0x18 {
            Ok(rip + 3)
        } else {
            Ok(rip + 1)
        }
    }

    unsafe fn get_return_address(&self, rip: u64) -> Result<u64, TracingError> {
        extern "C" {
            fn mmlayout_read_memory_safe(address: u64, buffer: *mut u8, size: usize) -> bool;
        }

        let mut rsp: u64 = 0;
        let success = mmlayout_read_memory_safe(rip + 8, &mut rsp as *mut u64 as *mut u8, 8);

        if success {
            Ok(rsp)
        } else {
            Err(TracingError::AccessDenied)
        }
    }

    unsafe fn set_temporary_breakpoint(&self, core_id: u32, address: u64) -> Result<(), TracingError> {
        extern "C" {
            fn vm_func_set_software_breakpoint(core_id: u32, address: u64) -> bool;
        }

        let success = vm_func_set_software_breakpoint(core_id, address);

        if success {
            Ok(())
        } else {
            Err(TracingError::OperationFailed)
        }
    }

    unsafe fn remove_temporary_breakpoint(&self, core_id: u32) -> Result<(), TracingError> {
        extern "C" {
            fn vm_func_remove_software_breakpoint(core_id: u32) -> bool;
        }

        let success = vm_func_remove_software_breakpoint(core_id);

        if success {
            Ok(())
        } else {
            Err(TracingError::OperationFailed)
        }
    }
}

unsafe fn vm_func_get_last_vmexit_rip(core_id: u32) -> u64 {
    extern "C" {
        fn vm_func_get_last_vmexit_rip(core_id: u32) -> u64;
    }

    vm_func_get_last_vmexit_rip(core_id)
}

unsafe fn vm_func_register_mtf_break(core_id: u32) {
    extern "C" {
        fn vm_func_register_mtf_break(core_id: u32);
    }

    vm_func_register_mtf_break(core_id);
}

unsafe fn vm_func_unregister_mtf_break(core_id: u32) {
    extern "C" {
        fn vm_func_unregister_mtf_break(core_id: u32);
    }

    vm_func_unregister_mtf_break(core_id);
}

unsafe fn vm_func_enable_mtf_and_change_external_interrupt_state(core_id: u32) {
    extern "C" {
        fn vm_func_enable_mtf_and_change_external_interrupt_state(core_id: u32);
    }

    vm_func_enable_mtf_and_change_external_interrupt_state(core_id);
}

unsafe fn vm_func_enable_and_check_for_previous_external_interrupts(core_id: u32) {
    extern "C" {
        fn vm_func_enable_and_check_for_previous_external_interrupts(core_id: u32);
    }

    vm_func_enable_and_check_for_previous_external_interrupts(core_id);
}

unsafe fn vm_func_change_mtf_unsetting_state(core_id: u32, enable: bool) {
    extern "C" {
        fn vm_func_change_mtf_unsetting_state(core_id: u32, enable: bool);
    }

    vm_func_change_mtf_unsetting_state(core_id, enable);
}

unsafe fn vm_func_set_rflag_trap_flag(enable: bool) {
    extern "C" {
        fn vm_func_set_rflag_trap_flag(enable: bool);
    }

    vm_func_set_rflag_trap_flag(enable);
}

unsafe fn vm_func_get_interruptibility_state() -> u64 {
    extern "C" {
        fn vm_func_get_interruptibility_state() -> u64;
    }

    vm_func_get_interruptibility_state()
}

unsafe fn vm_func_set_interruptibility_state(state: u64) {
    extern "C" {
        fn vm_func_set_interruptibility_state(state: u64);
    }

    vm_func_set_interruptibility_state(state);
}

unsafe fn vm_func_clear_stepping_bits(interruptibility: u64) -> u64 {
    extern "C" {
        fn vm_func_clear_stepping_bits(interruptibility: u64) -> u64;
    }

    vm_func_clear_stepping_bits(interruptibility)
}

unsafe fn disassembler_show_one_instruction_in_vmx_root_mode(rip: u64, show_hex: bool) {
    extern "C" {
        fn disassembler_show_one_instruction_in_vmx_root_mode(rip: u64, show_hex: bool);
    }

    disassembler_show_one_instruction_in_vmx_root_mode(rip, show_hex);
}

unsafe fn meta_dispatch_event_instrumentation_trace(core_id: u32) -> Result<(), TracingError> {
    extern "C" {
        fn meta_dispatch_event_instrumentation_trace(core_id: u32) -> i32;
    }

    let result = meta_dispatch_event_instrumentation_trace(core_id);

    if result == 0 {
        Ok(())
    } else {
        Err(TracingError::OperationFailed)
    }
}

pub static TRACING_MANAGER: Mutex<TracingManager> = Mutex::new(TracingManager::new(0));

pub fn initialize_tracing_manager(processor_count: u32) -> Result<(), TracingError> {
    let mut manager = TRACING_MANAGER.lock();
    *manager = TracingManager::new(processor_count);
    manager.initialize()
}

pub fn deinitialize_tracing_manager() {
    let mut manager = TRACING_MANAGER.lock();
    manager.deinitialize();
}

pub unsafe fn start_tracing(
    core_id: u32,
    mode: TracingMode,
    max_steps: u64,
) -> Result<(), TracingError> {
    let mut manager = TRACING_MANAGER.lock();
    manager.start_tracing(core_id, mode, max_steps)
}

pub unsafe fn stop_tracing(core_id: u32) -> Result<(), TracingError> {
    let mut manager = TRACING_MANAGER.lock();
    manager.stop_tracing(core_id)
}

pub unsafe fn handle_step(core_id: u32, rip: u64) -> Result<bool, TracingError> {
    let mut manager = TRACING_MANAGER.lock();
    manager.handle_step(core_id, rip)
}

pub unsafe fn perform_regular_step_in_instruction() -> Result<(), TracingError> {
    let mut manager = TRACING_MANAGER.lock();
    let core_id = 0;
    
    manager.perform_regular_step_in(core_id)?;

    extern "C" {
        fn ps_get_current_process_id() -> u32;
        fn ps_get_current_thread_id() -> u32;
        fn breakpoint_restore_the_trap_flag_once_triggered(process_id: u32, thread_id: u32) -> bool;
    }

    let process_id = ps_get_current_process_id();
    let thread_id = ps_get_current_thread_id();

    if !breakpoint_restore_the_trap_flag_once_triggered(process_id, thread_id) {
        // Log warning
    }

    Ok(())
}