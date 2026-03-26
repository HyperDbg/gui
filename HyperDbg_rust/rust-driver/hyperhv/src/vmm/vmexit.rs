use core::arch::asm;
use core::arch::naked_asm;
use crate::*;

#[no_mangle]
pub unsafe extern "C" fn AsmVmexitHandler() -> ! {
    naked_asm!(
        "push r15",
        "push r14",
        "push r13",
        "push r12",
        "push r11",
        "push r10",
        "push r9",
        "push r8",
        "push rdi",
        "push rsi",
        "push rbp",
        "push rbx",
        "push rdx",
        "push rcx",
        "push rax",
        
        "mov rax, cr2",
        "push rax",
        
        "mov rax, [rsp + 0x80]",
        "push rax",
        
        "lea rcx, [rsp + 0x10]",
        "lea rdx, [rsp + 0x88]",
        "lea r8, [rsp + 0x90]",
        
        "sub rsp, 0x28",
        "call {vmexit_handler}",
        "add rsp, 0x28",
        
        "add rsp, 8",
        "pop rax",
        "mov cr2, rax",
        
        "test al, al",
        "jz 2f",
        
        "pop rax",
        "pop rcx",
        "pop rdx",
        "pop rbx",
        "pop rbp",
        "pop rsi",
        "pop rdi",
        "pop r8",
        "pop r9",
        "pop r10",
        "pop r11",
        "pop r12",
        "pop r13",
        "pop r14",
        "pop r15",
        
        "vmresume",
        "jmp 3f",
        
        "2:",
        "pop rax",
        "pop rcx",
        "pop rdx",
        "pop rbx",
        "pop rbp",
        "pop rsi",
        "pop rdi",
        "pop r8",
        "pop r9",
        "pop r10",
        "pop r11",
        "pop r12",
        "pop r13",
        "pop r14",
        "pop r15",
        
        "add rsp, 16",
        "xor rax, rax",
        "ret",
        
        "3:",
        "int3",
        vmexit_handler = sym VmxVmexitHandlerEntry,
    );
}

#[no_mangle]
pub unsafe extern "C" fn VmxVmexitHandlerEntry(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let exit_reason = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_EXIT_REASON) as u32 & 0xFFFF;
    let exit_qualification = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_EXIT_QUALIFICATION);
    let instruction_length = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_EXIT_INSTRUCTION_LEN) as u32;

    let core_id = crate::processor::get_current_processor_number();

    let mut increment_rip = true;

    {
        let mut context = VMX_CONTEXT.lock();
        if let Some(vcpu) = context.get_vcpu_mut(core_id) {
            vcpu.exit_reason = exit_reason;
            vcpu.exit_qualification = exit_qualification;
            vcpu.guest_rip = *guest_rip as u64;
            vcpu.guest_rsp = *guest_rsp as u64;
            vcpu.guest_rflags = *guest_rflags as u64;
            vcpu.increment_rip = true;
        }
    }

    let should_continue = handle_vmexit(exit_reason, exit_qualification, guest_rsp, guest_rip, guest_rflags);

    {
        let mut context = VMX_CONTEXT.lock();
        if let Some(vcpu) = context.get_vcpu_mut(core_id) {
            increment_rip = vcpu.increment_rip;
        }
    }

    if increment_rip && should_continue {
        *guest_rip = *guest_rip.wrapping_add(instruction_length as u64);
    }

    should_continue
}

unsafe fn handle_vmexit(
    exit_reason: u32,
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    match exit_reason {
        0 => handle_exception_nmi(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        1 => handle_external_interrupt(),
        3 => handle_init_signal(),
        4 => handle_startup_ipi(),
        7 => handle_interrupt_window(),
        8 => handle_nmi_window(),
        10 => handle_cpuid(guest_rsp, guest_rip),
        12 => handle_hlt(),
        18 => handle_vmcall(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        28 => handle_cr_access(exit_qualification, guest_rsp, guest_rip),
        29 => handle_dr_access(exit_qualification, guest_rsp, guest_rip),
        30 => handle_io_instruction(exit_qualification, guest_rsp, guest_rip),
        31 => handle_rdmsr(exit_qualification, guest_rsp),
        32 => handle_wrmsr(exit_qualification, guest_rsp),
        48 => handle_ept_violation(exit_qualification, guest_rip),
        49 => handle_ept_misconfig(exit_qualification, guest_rip),
        52 => handle_invpcid(exit_qualification),
        _ => true,
    }
}

unsafe fn handle_exception_nmi(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let vector = exit_qualification & 0xFF;
    let error_code_valid = (exit_qualification & (1 << 11)) != 0;
    let nmi = (exit_qualification & (1 << 12)) != 0;
    let deliver_error_code = (exit_qualification & (1 << 13)) != 0;
    let instruction_length = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_EXIT_INSTRUCTION_LEN) as u32;

    match vector {
        0 => handle_divide_error(guest_rsp, guest_rip, guest_rflags),
        1 => handle_debug_exception(guest_rsp, guest_rip, guest_rflags),
        2 => handle_nmi(guest_rsp, guest_rip, guest_rflags),
        3 => handle_breakpoint(guest_rsp, guest_rip, guest_rflags),
        4 => handle_overflow(guest_rsp, guest_rip, guest_rflags),
        5 => handle_bound_range(guest_rsp, guest_rip, guest_rflags),
        6 => handle_invalid_opcode(guest_rsp, guest_rip, guest_rflags),
        7 => handle_device_not_available(guest_rsp, guest_rip, guest_rflags),
        8 => handle_double_fault(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        10 => handle_invalid_tss(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        11 => handle_segment_not_present(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        12 => handle_stack_segment_fault(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        13 => handle_general_protection(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        14 => handle_page_fault(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        16 => handle_x87_fpu_error(guest_rsp, guest_rip, guest_rflags),
        17 => handle_alignment_check(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        18 => handle_machine_check(guest_rsp, guest_rip, guest_rflags),
        19 => handle_simd_exception(guest_rsp, guest_rip, guest_rflags),
        _ => true,
    }
}

unsafe fn handle_divide_error(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_debug_exception(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let rflags = *guest_rflags;
    *guest_rflags = rflags & !(1 << 8);

    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    true
}

unsafe fn handle_nmi(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let mut ctrl = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_PIN_BASED) as u32;
    ctrl |= crate::vmm::vmx::PIN_BASED_EXEC_CTRL_NMI_EXITING;
    crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_PIN_BASED, ctrl as u64);
    true
}

unsafe fn handle_breakpoint(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    true
}

unsafe fn handle_overflow(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_bound_range(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_invalid_opcode(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_device_not_available(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let cr0 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR0);
    if (cr0 & crate::vmm::vmx::CR0_TS) != 0 {
        let core_id = crate::processor::get_current_processor_number();
        let mut context = VMX_CONTEXT.lock();
        if let Some(vcpu) = context.get_vcpu_mut(core_id) {
            vcpu.increment_rip = false;
        }
    }
    true
}

unsafe fn handle_double_fault(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_invalid_tss(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_segment_not_present(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_stack_segment_fault(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_general_protection(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_page_fault(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let cr2 = crate::vmm::vmx::read_cr2();
    
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    true
}

unsafe fn handle_x87_fpu_error(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_alignment_check(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_machine_check(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    false
}

unsafe fn handle_simd_exception(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    true
}

unsafe fn handle_external_interrupt() -> bool {
    let mut ctrl = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_PIN_BASED) as u32;
    
    if (ctrl & crate::vmm::vmx::PIN_BASED_EXEC_CTRL_EXTINT_EXITING) != 0 {
        ctrl &= !crate::vmm::vmx::PIN_BASED_EXEC_CTRL_EXTINT_EXITING;
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_PIN_BASED, ctrl as u64);
    }

    true
}

unsafe fn handle_init_signal() -> bool {
    true
}

unsafe fn handle_startup_ipi() -> bool {
    true
}

unsafe fn handle_interrupt_window() -> bool {
    let mut ctrl = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_PRIMARY_PROC_BASED) as u32;
    ctrl &= !crate::vmm::vmx::PROC_BASED_EXEC_CTRL_INT_WINDOW_EXITING;
    crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_PRIMARY_PROC_BASED, ctrl as u64);
    true
}

unsafe fn handle_nmi_window() -> bool {
    let mut ctrl = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_PRIMARY_PROC_BASED) as u32;
    ctrl &= !crate::vmm::vmx::PROC_BASED_EXEC_CTRL_NMI_WINDOW_EXITING;
    crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_PRIMARY_PROC_BASED, ctrl as u64);
    true
}

unsafe fn handle_cpuid(guest_rsp: *mut u64, guest_rip: *mut u64) -> bool {
    let leaf: u32;
    let sub_leaf: u32;

    asm!(
        "mov {0}, eax",
        "mov {1}, ecx",
        out(reg) leaf,
        out(reg) sub_leaf,
    );

    let (eax, ebx, ecx, edx) = match leaf {
        0x40000000 => {
            (0x40000001, 0x52455059, 0x47424448, 0x00000000)
        }
        0x40000001 => {
            (HYPERDBG_SIGNATURE as u32, (HYPERDBG_SIGNATURE >> 32) as u32, 0x00000001, 0x00000000)
        }
        _ => {
            let mut eax_out: u32;
            let mut ebx_out: u32;
            let mut ecx_out: u32;
            let mut edx_out: u32;

            asm!(
                "cpuid",
                inout("eax") leaf => eax_out,
                inout("ecx") sub_leaf => ecx_out,
                out("ebx") ebx_out,
                out("edx") edx_out,
            );

            (eax_out, ebx_out, ecx_out, edx_out)
        }
    };

    asm!(
        "mov eax, {0}",
        "mov ebx, {1}",
        "mov ecx, {2}",
        "mov edx, {3}",
        in(reg) eax,
        in(reg) ebx,
        in(reg) ecx,
        in(reg) edx,
    );

    true
}

unsafe fn handle_hlt() -> bool {
    true
}

unsafe fn handle_vmcall(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let rax: u64;
    let rbx: u64;
    let rcx: u64;
    let rdx: u64;

    asm!(
        "mov {0}, rax",
        "mov {1}, rbx",
        "mov {2}, rcx",
        "mov {3}, rdx",
        out(reg) rax,
        out(reg) rbx,
        out(reg) rcx,
        out(reg) rdx,
    );

    let vmcall_number = rax as u32;

    match vmcall_number {
        0x1 => handle_vmxoff(guest_rsp, guest_rip, guest_rflags),
        0x2 => handle_test(guest_rsp, guest_rip, guest_rflags),
        0x10 => handle_set_breakpoint(rbx, rcx),
        0x11 => handle_clear_breakpoint(rbx),
        0x12 => handle_read_memory(rbx, rcx, rdx, guest_rsp),
        0x13 => handle_write_memory(rbx, rcx, rdx),
        0x14 => handle_get_context(guest_rsp),
        0x15 => handle_set_context(rbx, guest_rsp),
        _ => true,
    }
}

unsafe fn handle_vmxoff(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::processor::get_current_processor_number();

    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.has_launched = false;
        vcpu.state = VmxState::Terminated;
    }

    false
}

unsafe fn handle_test(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    asm!(
        "mov rax, 0x12345678",
    );
    true
}

unsafe fn handle_set_breakpoint(address: u64, process_id: u64) -> bool {
    let mut events = crate::globals::EVENTS.lock();
    if let Some(ref mut event_manager) = *events {
        let id = address ^ process_id ^ (crate::processor::get_current_processor_number() as u64);
        
        let config = crate::globals::BreakpointEventConfig {
            id,
            address,
            process_id: process_id as u32,
            enabled: true,
            hit_count: 0,
        };
        
        event_manager.breakpoint_events.insert(address, config);
    }

    let guest_cr3 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR3);
    let physical_address = translate_guest_virtual_to_physical(address, guest_cr3);
    
    if physical_address != 0 {
        let mut hook_ctx = crate::hooks::HOOK_CONTEXT.lock();
        let _ = hook_ctx.set_ept_hook(address, 0, process_id as u32);
    }

    asm!(
        "mov rax, 1",
    );
    true
}

unsafe fn handle_clear_breakpoint(address: u64) -> bool {
    let mut events = crate::globals::EVENTS.lock();
    if let Some(ref mut event_manager) = *events {
        event_manager.breakpoint_events.remove(&address);
    }

    let mut hook_ctx = crate::hooks::HOOK_CONTEXT.lock();
    let _ = hook_ctx.remove_ept_hook(address);

    asm!(
        "mov rax, 1",
    );
    true
}

unsafe fn handle_read_memory(address: u64, size: u64, buffer: u64, guest_rsp: *mut u64) -> bool {
    if size == 0 || buffer == 0 {
        asm!("mov rax, 0");
        return true;
    }

    let guest_cr3 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR3);
    
    let mut bytes_read: usize = 0;
    let mut current_va = address;
    let buffer_ptr = buffer as *mut u8;
    let mut buffer_offset: usize = 0;

    while bytes_read < size as usize {
        let page_offset = current_va & 0xFFF;
        let remaining_in_page = 0x1000 - page_offset;
        let remaining_to_read = size as usize - bytes_read;
        let chunk_size = core::cmp::min(remaining_in_page as usize, remaining_to_read);

        let physical_address = translate_guest_virtual_to_physical(current_va, guest_cr3);
        
        if physical_address == 0 {
            break;
        }

        if let Ok(mapped_va) = crate::memory::MemoryManager::map_physical_memory(
            physical_address,
            chunk_size,
            crate::memory::CACHE_TYPE_CACHED,
        ) {
            core::ptr::copy_nonoverlapping(
                mapped_va as *const u8,
                buffer_ptr.add(buffer_offset),
                chunk_size,
            );
            
            crate::memory::MemoryManager::unmap_physical_memory(mapped_va, chunk_size);
            bytes_read += chunk_size;
            buffer_offset += chunk_size;
            current_va += chunk_size as u64;
        } else {
            break;
        }
    }

    asm!(
        "mov rax, {0}",
        in(reg) bytes_read as u64,
    );
    true
}

unsafe fn handle_write_memory(address: u64, size: u64, buffer: u64) -> bool {
    if size == 0 || buffer == 0 {
        asm!("mov rax, 0");
        return true;
    }

    let guest_cr3 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR3);
    
    let mut bytes_written: usize = 0;
    let mut current_va = address;
    let buffer_ptr = buffer as *const u8;
    let mut buffer_offset: usize = 0;

    while bytes_written < size as usize {
        let page_offset = current_va & 0xFFF;
        let remaining_in_page = 0x1000 - page_offset;
        let remaining_to_write = size as usize - bytes_written;
        let chunk_size = core::cmp::min(remaining_in_page as usize, remaining_to_write);

        let physical_address = translate_guest_virtual_to_physical(current_va, guest_cr3);
        
        if physical_address == 0 {
            break;
        }

        if let Ok(mapped_va) = crate::memory::MemoryManager::map_physical_memory(
            physical_address,
            chunk_size,
            crate::memory::CACHE_TYPE_CACHED,
        ) {
            core::ptr::copy_nonoverlapping(
                buffer_ptr.add(buffer_offset),
                mapped_va as *mut u8,
                chunk_size,
            );
            
            crate::memory::MemoryManager::unmap_physical_memory(mapped_va, chunk_size);
            bytes_written += chunk_size;
            buffer_offset += chunk_size;
            current_va += chunk_size as u64;
        } else {
            break;
        }
    }

    asm!(
        "mov rax, {0}",
        in(reg) bytes_written as u64,
    );
    true
}

#[repr(C)]
pub struct GuestContext {
    pub rip: u64,
    pub rsp: u64,
    pub rflags: u64,
    pub rax: u64,
    pub rbx: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub rbp: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
    pub cr0: u64,
    pub cr2: u64,
    pub cr3: u64,
    pub cr4: u64,
    pub cr8: u64,
    pub dr0: u64,
    pub dr1: u64,
    pub dr2: u64,
    pub dr3: u64,
    pub dr6: u64,
    pub dr7: u64,
    pub cs: u16,
    pub ds: u16,
    pub es: u16,
    pub fs: u16,
    pub gs: u16,
    pub ss: u16,
    pub gdtr_base: u64,
    pub gdtr_limit: u16,
    pub idtr_base: u64,
    pub idtr_limit: u16,
    pub ldtr: u16,
    pub tr: u16,
    pub fs_base: u64,
    pub gs_base: u64,
    pub kernel_gs_base: u64,
    pub efer: u64,
    pub sysenter_cs: u64,
    pub sysenter_esp: u64,
    pub sysenter_eip: u64,
    pub star: u64,
    pub lstar: u64,
    pub cstar: u64,
    pub sfmask: u64,
    pub debugctl: u64,
    pub last_branch_from: u64,
    pub last_branch_to: u64,
    pub last_exception_from: u64,
    pub last_exception_to: u64,
    pub tsc: u64,
    pub tsc_aux: u64,
}

unsafe fn handle_get_context(guest_rsp: *mut u64) -> bool {
    let context_size = core::mem::size_of::<GuestContext>();
    let context_ptr = guest_rsp as *mut GuestContext;
    
    if context_ptr.is_null() {
        asm!("mov rax, 0");
        return true;
    }

    let rip = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_RIP);
    let rsp = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_RSP);
    let rflags = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_RFLAGS);
    let cr0 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR0);
    let cr3 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR3);
    let cr4 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR4);
    let dr7 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_DR7);
    
    let cs = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CS_SELECTOR) as u16;
    let ds = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_DS_SELECTOR) as u16;
    let es = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_ES_SELECTOR) as u16;
    let fs = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_FS_SELECTOR) as u16;
    let gs = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_GS_SELECTOR) as u16;
    let ss = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_SS_SELECTOR) as u16;
    
    let fs_base = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_FS_BASE);
    let gs_base = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_GS_BASE);
    let gdtr_base = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_GDTR_BASE);
    let gdtr_limit = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_GDTR_LIMIT) as u16;
    let idtr_base = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_IDTR_BASE);
    let idtr_limit = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_IDTR_LIMIT) as u16;
    let ldtr = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_LDTR_SELECTOR) as u16;
    let tr = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_TR_SELECTOR) as u16;

    let core_id = crate::processor::get_current_processor_number();
    let (kernel_gs_base, dr0, dr1, dr2, dr3, dr6, debugctl, last_branch_from, last_branch_to, last_exception_from, last_exception_to) = {
        let context = VMX_CONTEXT.lock();
        if let Some(vcpus) = context.vcpus.as_ref() {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &vcpus[core_id as usize];
                (
                    vcpu.kernel_gs_base,
                    vcpu.dr0,
                    vcpu.dr1,
                    vcpu.dr2,
                    vcpu.dr3,
                    vcpu.dr6,
                    vcpu.debugctl,
                    vcpu.last_branch_from,
                    vcpu.last_branch_to,
                    vcpu.last_exception_from,
                    vcpu.last_exception_to,
                )
            } else {
                (0, 0, 0, 0, 0, 0xFFFF0FF0, 0, 0, 0, 0, 0)
            }
        } else {
            (0, 0, 0, 0, 0, 0xFFFF0FF0, 0, 0, 0, 0, 0)
        }
    };

    let efer = crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_EFER);
    let star = crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_STAR);
    let lstar = crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_LSTAR);
    let cstar = crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_CSTAR);
    let sfmask = crate::vmm::vmx::read_msr(crate::vmm::vmx::IA32_SFMASK);
    let sysenter_cs = crate::vmm::vmx::read_msr(0x174);
    let sysenter_esp = crate::vmm::vmx::read_msr(0x175);
    let sysenter_eip = crate::vmm::vmx::read_msr(0x176);

    let tsc: u64;
    asm!("rdtsc", out("edx") _, out("eax") tsc);

    let context = GuestContext {
        rip,
        rsp,
        rflags,
        rax: *guest_rsp.add(0),
        rbx: *guest_rsp.add(1),
        rcx: *guest_rsp.add(2),
        rdx: *guest_rsp.add(3),
        rsi: *guest_rsp.add(4),
        rdi: *guest_rsp.add(5),
        rbp: *guest_rsp.add(6),
        r8: *guest_rsp.add(7),
        r9: *guest_rsp.add(8),
        r10: *guest_rsp.add(9),
        r11: *guest_rsp.add(10),
        r12: *guest_rsp.add(11),
        r13: *guest_rsp.add(12),
        r14: *guest_rsp.add(13),
        r15: *guest_rsp.add(14),
        cr0,
        cr2: crate::vmm::vmx::read_cr2(),
        cr3,
        cr4,
        cr8: 0,
        dr0,
        dr1,
        dr2,
        dr3,
        dr6,
        dr7,
        cs,
        ds,
        es,
        fs,
        gs,
        ss,
        gdtr_base,
        gdtr_limit,
        idtr_base,
        idtr_limit,
        ldtr,
        tr,
        fs_base,
        gs_base,
        kernel_gs_base,
        efer,
        sysenter_cs,
        sysenter_esp,
        sysenter_eip,
        star,
        lstar,
        cstar,
        sfmask,
        debugctl,
        last_branch_from,
        last_branch_to,
        last_exception_from,
        last_exception_to,
        tsc,
        tsc_aux: 0,
    };

    core::ptr::write(context_ptr, context);

    asm!("mov rax, 1");
    true
}

unsafe fn handle_set_context(flags: u64, guest_rsp: *mut u64) -> bool {
    let context_ptr = guest_rsp as *const GuestContext;
    
    if context_ptr.is_null() {
        asm!("mov rax, 0");
        return true;
    }

    let context = core::ptr::read(context_ptr);
    let core_id = crate::processor::get_current_processor_number();

    if flags & 0x01 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_RIP, context.rip);
    }
    
    if flags & 0x02 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_RSP, context.rsp);
    }
    
    if flags & 0x04 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_RFLAGS, context.rflags);
    }
    
    if flags & 0x08 != 0 {
        *guest_rsp.add(0) = context.rax;
        *guest_rsp.add(1) = context.rbx;
        *guest_rsp.add(2) = context.rcx;
        *guest_rsp.add(3) = context.rdx;
        *guest_rsp.add(4) = context.rsi;
        *guest_rsp.add(5) = context.rdi;
        *guest_rsp.add(6) = context.rbp;
        *guest_rsp.add(7) = context.r8;
        *guest_rsp.add(8) = context.r9;
        *guest_rsp.add(9) = context.r10;
        *guest_rsp.add(10) = context.r11;
        *guest_rsp.add(11) = context.r12;
        *guest_rsp.add(12) = context.r13;
        *guest_rsp.add(13) = context.r14;
        *guest_rsp.add(14) = context.r15;
    }
    
    if flags & 0x10 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_CR0, context.cr0);
    }
    
    if flags & 0x20 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_CR3, context.cr3);
    }
    
    if flags & 0x40 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_CR4, context.cr4);
    }
    
    if flags & 0x80 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_DR7, context.dr7);
    }
    
    if flags & 0x100 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_FS_BASE, context.fs_base);
    }
    
    if flags & 0x200 != 0 {
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_GS_BASE, context.gs_base);
    }

    if flags & 0x400 != 0 {
        let mut vmx_context = VMX_CONTEXT.lock();
        if let Some(vcpus) = vmx_context.vcpus.as_mut() {
            if (core_id as usize) < vcpus.len() {
                vcpus[core_id as usize].kernel_gs_base = context.kernel_gs_base;
            }
        }
    }

    if flags & 0x800 != 0 {
        let mut vmx_context = VMX_CONTEXT.lock();
        if let Some(vcpus) = vmx_context.vcpus.as_mut() {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &mut vcpus[core_id as usize];
                vcpu.dr0 = context.dr0;
                vcpu.dr1 = context.dr1;
                vcpu.dr2 = context.dr2;
                vcpu.dr3 = context.dr3;
                vcpu.dr6 = context.dr6;
            }
        }
    }

    if flags & 0x1000 != 0 {
        let mut vmx_context = VMX_CONTEXT.lock();
        if let Some(vcpus) = vmx_context.vcpus.as_mut() {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &mut vcpus[core_id as usize];
                vcpu.debugctl = context.debugctl;
                vcpu.last_branch_from = context.last_branch_from;
                vcpu.last_branch_to = context.last_branch_to;
                vcpu.last_exception_from = context.last_exception_from;
                vcpu.last_exception_to = context.last_exception_to;
            }
        }
    }

    if flags & 0x2000 != 0 {
        crate::vmm::vmx::write_msr(crate::vmm::vmx::IA32_EFER, context.efer);
    }

    if flags & 0x4000 != 0 {
        crate::vmm::vmx::write_msr(crate::vmm::vmx::IA32_STAR, context.star);
        crate::vmm::vmx::write_msr(crate::vmm::vmx::IA32_LSTAR, context.lstar);
        crate::vmm::vmx::write_msr(crate::vmm::vmx::IA32_CSTAR, context.cstar);
        crate::vmm::vmx::write_msr(crate::vmm::vmx::IA32_SFMASK, context.sfmask);
    }

    if flags & 0x8000 != 0 {
        crate::vmm::vmx::write_msr(0x174, context.sysenter_cs);
        crate::vmm::vmx::write_msr(0x175, context.sysenter_esp);
        crate::vmm::vmx::write_msr(0x176, context.sysenter_eip);
    }

    asm!("mov rax, 1");
    true
}

unsafe fn translate_guest_virtual_to_physical(virtual_address: u64, cr3: u64) -> u64 {
    let translator = crate::memory::VirtualAddressTranslator::new(cr3);
    translator.translate(virtual_address).unwrap_or(0)
}

unsafe fn handle_cr_access(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let cr_num = (exit_qualification >> 4) & 0xF;
    let access_type = exit_qualification & 0x3;
    let lmsw_source = (exit_qualification >> 6) & 0x1;
    let gpr = (exit_qualification >> 8) & 0xF;

    match cr_num {
        0 => handle_cr0_access(access_type, gpr, guest_rsp),
        3 => handle_cr3_access(access_type, gpr, guest_rsp),
        4 => handle_cr4_access(access_type, gpr, guest_rsp),
        8 => handle_cr8_access(access_type, gpr, guest_rsp),
        _ => true,
    }
}

unsafe fn handle_cr0_access(access_type: u64, gpr: u64, guest_rsp: *mut u64) -> bool {
    let guest_cr0 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR0);
    let cr0_shadow = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_CR0_SHADOW);
    
    if access_type == 0 {
        let gpr_offset = (gpr as usize) * 8;
        let value = *guest_rsp.add(gpr_offset / 8);
        
        let new_shadow = value;
        let new_cr0 = (guest_cr0 & !crate::vmm::vmx::CR0_PE) | (value & crate::vmm::vmx::CR0_PE);
        
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_CR0_SHADOW, new_shadow);
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_CR0, new_cr0);
    } else {
        let shadow_value = cr0_shadow;
        let gpr_offset = (gpr as usize) * 8;
        *guest_rsp.add(gpr_offset / 8) = shadow_value;
    }

    true
}

unsafe fn handle_cr3_access(access_type: u64, gpr: u64, guest_rsp: *mut u64) -> bool {
    if access_type == 0 {
        let gpr_offset = (gpr as usize) * 8;
        let value = *guest_rsp.add(gpr_offset / 8);
        
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_CR3, value);
        
        if crate::vmm::ept::EPT_CONTEXT.lock().is_some() {
            crate::vmm::vmx::invept(1, 0);
        }
    } else {
        let guest_cr3 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR3);
        let gpr_offset = (gpr as usize) * 8;
        *guest_rsp.add(gpr_offset / 8) = guest_cr3;
    }

    true
}

unsafe fn handle_cr4_access(access_type: u64, gpr: u64, guest_rsp: *mut u64) -> bool {
    let guest_cr4 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_CR4);
    let cr4_shadow = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_CR4_SHADOW);
    
    if access_type == 0 {
        let gpr_offset = (gpr as usize) * 8;
        let value = *guest_rsp.add(gpr_offset / 8);
        
        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_CTRL_CR4_SHADOW, value);
    } else {
        let gpr_offset = (gpr as usize) * 8;
        *guest_rsp.add(gpr_offset / 8) = cr4_shadow;
    }

    true
}

unsafe fn handle_cr8_access(access_type: u64, gpr: u64, guest_rsp: *mut u64) -> bool {
    if access_type == 0 {
        let gpr_offset = (gpr as usize) * 8;
        let value = *guest_rsp.add(gpr_offset / 8);
        
        let tpr = (value & 0xF) << 4;
        crate::vmm::vmx::vmwrite(0x2012, tpr as u64);
    } else {
        let tpr = crate::vmm::vmx::vmread(0x2012);
        let cr8 = (tpr >> 4) & 0xF;
        let gpr_offset = (gpr as usize) * 8;
        *guest_rsp.add(gpr_offset / 8) = cr8;
    }

    true
}

unsafe fn handle_dr_access(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let dr_num = (exit_qualification >> 8) & 0x7;
    let is_write = (exit_qualification & 1) == 0;
    let gpr = (exit_qualification >> 16) & 0xF;

    let core_id = crate::processor::get_current_processor_number();
    let guest_dr7 = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_DR7);
    let gpr_offset = (gpr as usize) * 8;

    if is_write {
        let value = *guest_rsp.add(gpr_offset / 8);
        
        let mut context = VMX_CONTEXT.lock();
        if let Some(vcpus) = context.vcpus.as_mut() {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &mut vcpus[core_id as usize];
                
                match dr_num {
                    0 => vcpu.dr0 = value,
                    1 => vcpu.dr1 = value,
                    2 => vcpu.dr2 = value,
                    3 => vcpu.dr3 = value,
                    6 => vcpu.dr6 = value,
                    7 => {
                        crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_DR7, value);
                    }
                    _ => {}
                }
            }
        }
    } else {
        let mut context = VMX_CONTEXT.lock();
        if let Some(vcpus) = context.vcpus.as_ref() {
            if (core_id as usize) < vcpus.len() {
                let vcpu = &vcpus[core_id as usize];
                
                match dr_num {
                    0 => *guest_rsp.add(gpr_offset / 8) = vcpu.dr0,
                    1 => *guest_rsp.add(gpr_offset / 8) = vcpu.dr1,
                    2 => *guest_rsp.add(gpr_offset / 8) = vcpu.dr2,
                    3 => *guest_rsp.add(gpr_offset / 8) = vcpu.dr3,
                    6 => *guest_rsp.add(gpr_offset / 8) = vcpu.dr6,
                    7 => *guest_rsp.add(gpr_offset / 8) = guest_dr7,
                    _ => *guest_rsp.add(gpr_offset / 8) = 0,
                }
            } else {
                *guest_rsp.add(gpr_offset / 8) = 0;
            }
        } else {
            match dr_num {
                6 => *guest_rsp.add(gpr_offset / 8) = 0xFFFF0FF0,
                7 => *guest_rsp.add(gpr_offset / 8) = guest_dr7,
                _ => *guest_rsp.add(gpr_offset / 8) = 0,
            }
        }
    }

    true
}

unsafe fn handle_io_instruction(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let port = (exit_qualification >> 16) as u16;
    let is_string = (exit_qualification & (1 << 4)) != 0;
    let is_write = (exit_qualification & (1 << 3)) == 0;
    let size = ((exit_qualification >> 7) & 0x7) as u32 + 1;
    let rep = (exit_qualification & (1 << 5)) != 0;
    let operand_size = (exit_qualification & (1 << 6)) != 0;

    let core_id = crate::processor::get_current_processor_number();
    
    if is_string {
        let rsi: u64;
        let rdi: u64;
        let rcx: u64;
        
        asm!(
            "mov {0}, rsi",
            "mov {1}, rdi",
            "mov {2}, rcx",
            out(reg) rsi,
            out(reg) rdi,
            out(reg) rcx,
        );

        let count = if rep { rcx as usize } else { 1 };
        let direction = (crate::vmm::vmx::read_rflags() & 0x400) != 0;
        
        let mut current_addr = if is_write { rsi } else { rdi };
        
        for _ in 0..count {
            if is_write {
                let data: u32 = match size {
                    1 => core::ptr::read_volatile(current_addr as *const u8) as u32,
                    2 => core::ptr::read_volatile(current_addr as *const u16) as u32,
                    4 => core::ptr::read_volatile(current_addr as *const u32),
                    _ => 0,
                };
                
                match size {
                    1 => asm!("out dx, al", in("dx") port, in("al") data as u8),
                    2 => asm!("out dx, ax", in("dx") port, in("ax") data as u16),
                    4 => asm!("out dx, eax", in("dx") port, in("eax") data),
                    _ => {}
                }
            } else {
                let mut data: u32 = 0;
                
                match size {
                    1 => {
                        let mut al: u8;
                        asm!("in al, dx", out("al") al, in("dx") port);
                        data = al as u32;
                    }
                    2 => {
                        let mut ax: u16;
                        asm!("in ax, dx", out("ax") ax, in("dx") port);
                        data = ax as u32;
                    }
                    4 => {
                        asm!("in eax, dx", out("eax") data, in("dx") port);
                    }
                    _ => {}
                }
                
                match size {
                    1 => core::ptr::write_volatile(current_addr as *mut u8, data as u8),
                    2 => core::ptr::write_volatile(current_addr as *mut u16, data as u16),
                    4 => core::ptr::write_volatile(current_addr as *mut u32, data),
                    _ => {}
                }
            }
            
            if direction {
                current_addr = current_addr.wrapping_sub(size as u64);
            } else {
                current_addr = current_addr.wrapping_add(size as u64);
            }
        }
        
        if rep {
            asm!("xor ecx, ecx");
        }
        
        if is_write {
            asm!("mov rsi, {0}", in(reg) current_addr);
        } else {
            asm!("mov rdi, {0}", in(reg) current_addr);
        }
    } else {
        if is_write {
            let eax: u32;
            asm!("mov {0}, eax", out(reg) eax);
            
            match size {
                1 => asm!("out dx, al", in("dx") port, in("al") eax as u8),
                2 => asm!("out dx, ax", in("dx") port, in("ax") eax as u16),
                4 => asm!("out dx, eax", in("dx") port, in("eax") eax),
                _ => {}
            }
        } else {
            let mut data: u32 = 0;
            
            match size {
                1 => {
                    let mut al: u8;
                    asm!("in al, dx", out("al") al, in("dx") port);
                    data = al as u32;
                }
                2 => {
                    let mut ax: u16;
                    asm!("in ax, dx", out("ax") ax, in("dx") port);
                    data = ax as u32;
                }
                4 => {
                    asm!("in eax, dx", out("eax") data, in("dx") port);
                }
                _ => {}
            }
            
            asm!("mov eax, {0}", in(reg) data);
        }
    }

    true
}

unsafe fn handle_rdmsr(
    exit_qualification: u64,
    guest_rsp: *mut u64,
) -> bool {
    let ecx: u32;
    asm!("mov {0}, ecx", out(reg) ecx);

    let value = match ecx {
        crate::vmm::vmx::IA32_FS_BASE => crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_FS_BASE),
        crate::vmm::vmx::IA32_GS_BASE => crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_GUEST_GS_BASE),
        crate::vmm::vmx::IA32_KERNEL_GS_BASE => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].kernel_gs_base
                } else {
                    0
                }
            } else {
                0
            }
        }
        crate::vmm::vmx::IA32_DEBUGCTL => {
            let core_id = crate::processor::get_current_processor_number();
            let context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].debugctl
                } else {
                    0
                }
            } else {
                0
            }
        }
        crate::vmm::vmx::IA32_LASTBRANCHFROMIP => {
            let core_id = crate::processor::get_current_processor_number();
            let context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_branch_from
                } else {
                    0
                }
            } else {
                0
            }
        }
        crate::vmm::vmx::IA32_LASTBRANCHTOIP => {
            let core_id = crate::processor::get_current_processor_number();
            let context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_branch_to
                } else {
                    0
                }
            } else {
                0
            }
        }
        crate::vmm::vmx::IA32_LASTINTOIP => {
            let core_id = crate::processor::get_current_processor_number();
            let context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_exception_from
                } else {
                    0
                }
            } else {
                0
            }
        }
        crate::vmm::vmx::IA32_LASTINTFROMIP => {
            let core_id = crate::processor::get_current_processor_number();
            let context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_ref() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_exception_to
                } else {
                    0
                }
            } else {
                0
            }
        }
        _ => crate::vmm::vmx::read_msr(ecx),
    };

    let eax = value as u32;
    let edx = (value >> 32) as u32;

    asm!(
        "mov eax, {0}",
        "mov edx, {1}",
        in(reg) eax,
        in(reg) edx,
    );

    true
}

unsafe fn handle_wrmsr(
    exit_qualification: u64,
    guest_rsp: *mut u64,
) -> bool {
    let ecx: u32;
    let eax: u32;
    let edx: u32;

    asm!(
        "mov {0}, ecx",
        "mov {1}, eax",
        "mov {2}, edx",
        out(reg) ecx,
        out(reg) eax,
        out(reg) edx,
    );

    let value = ((edx as u64) << 32) | (eax as u64);

    match ecx {
        crate::vmm::vmx::IA32_FS_BASE => {
            crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_FS_BASE, value);
        }
        crate::vmm::vmx::IA32_GS_BASE => {
            crate::vmm::vmx::vmwrite(crate::vmm::vmx::VMCS_GUEST_GS_BASE, value);
        }
        crate::vmm::vmx::IA32_KERNEL_GS_BASE => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].kernel_gs_base = value;
                }
            }
        }
        crate::vmm::vmx::IA32_DEBUGCTL => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].debugctl = value;
                }
            }
        }
        crate::vmm::vmx::IA32_LASTBRANCHFROMIP => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_branch_from = value;
                }
            }
        }
        crate::vmm::vmx::IA32_LASTBRANCHTOIP => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_branch_to = value;
                }
            }
        }
        crate::vmm::vmx::IA32_LASTINTOIP => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_exception_from = value;
                }
            }
        }
        crate::vmm::vmx::IA32_LASTINTFROMIP => {
            let core_id = crate::processor::get_current_processor_number();
            let mut context = VMX_CONTEXT.lock();
            if let Some(vcpus) = context.vcpus.as_mut() {
                if (core_id as usize) < vcpus.len() {
                    vcpus[core_id as usize].last_exception_to = value;
                }
            }
        }
        _ => {
            crate::vmm::vmx::write_msr(ecx, value);
        }
    }

    true
}

unsafe fn handle_ept_violation(
    exit_qualification: u64,
    guest_rip: *mut u64,
) -> bool {
    let guest_physical_address = crate::vmm::vmx::vmread(0x2400);
    let guest_linear_address = crate::vmm::vmx::vmread(0x2402);

    let is_read = (exit_qualification & 1) != 0;
    let is_write = (exit_qualification & 2) != 0;
    let is_execute = (exit_qualification & 4) != 0;
    let readable = (exit_qualification & 8) != 0;
    let writable = (exit_qualification & 16) != 0;
    let executable = (exit_qualification & 32) != 0;
    let valid_guest_state = (exit_qualification & 128) != 0;
    let is_user_mode_access = (exit_qualification & 256) != 0;
    let caused_by_translation = (exit_qualification & 0x10000000) != 0;
    let is_guest_linear_address_valid = (exit_qualification & 0x80000000) != 0;
    let nmi_unblocking_due_to_iret = (exit_qualification & 0x20000000) != 0;

    let core_id = crate::processor::get_current_processor_number();
    let page_aligned_pa = guest_physical_address & !0xFFF;

    let mut ept_ctx = EPT_CONTEXT.lock();
    if let Some(ref mut ept) = *ept_ctx {
        if ept.handle_ept_violation(guest_physical_address, is_read, is_write, is_execute) {
            crate::vmm::vmx::invept(1, ept.get_eptp());
            return true;
        }
        
        if !readable && !writable && !executable && valid_guest_state {
            if ept.ept_context.translate(page_aligned_pa).is_none() {
                let _ = ept.ept_context.map_page(
                    page_aligned_pa,
                    page_aligned_pa,
                    true,
                    true,
                    true,
                );
                crate::vmm::vmx::invept(1, ept.get_eptp());
                return true;
            }
        }
        
        if is_execute && !executable {
            let events = crate::globals::EVENTS.lock();
            if let Some(ref event_manager) = *events {
                if let Some(_bp_event) = event_manager.breakpoint_events.get(&guest_linear_address) {
                    drop(events);
                    
                    let mut context = VMX_CONTEXT.lock();
                    if let Some(vcpus) = context.vcpus.as_mut() {
                        if (core_id as usize) < vcpus.len() {
                            vcpus[core_id as usize].increment_rip = false;
                        }
                    }
                    
                    return true;
                }
            }
        }
        
        if is_write && !writable {
            if ept.is_hooked(page_aligned_pa) {
                if ept.restore_original_page(page_aligned_pa) {
                    crate::vmm::vmx::invept(1, ept.get_eptp());
                    return true;
                }
            }
        }
        
        if is_read && !readable {
            if ept.is_hooked(page_aligned_pa) {
                if ept.switch_to_shadow_page(page_aligned_pa) {
                    crate::vmm::vmx::invept(1, ept.get_eptp());
                    return true;
                }
            }
        }
    }

    true
}

unsafe fn handle_ept_misconfig(
    exit_qualification: u64,
    guest_rip: *mut u64,
) -> bool {
    let guest_physical_address = crate::vmm::vmx::vmread(0x2400);
    let guest_linear_address = crate::vmm::vmx::vmread(0x2402);

    let mut ept_ctx = EPT_CONTEXT.lock();
    if let Some(ref mut ept) = *ept_ctx {
        if ept.handle_ept_misconfig(guest_physical_address) {
            crate::vmm::vmx::invept(1, ept.get_eptp());
            return true;
        }
    }

    true
}

unsafe fn handle_invpcid(exit_qualification: u64) -> bool {
    let invpcid_type = exit_qualification & 0x3;
    let descriptor_gpa = crate::vmm::vmx::vmread(0x2400);

    match invpcid_type {
        0 => {
            crate::vmm::vmx::invept(1, 0);
        }
        1 => {
            crate::vmm::vmx::invept(1, 0);
        }
        2 => {
            let eptp = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_EPTP);
            crate::vmm::vmx::invept(2, eptp);
        }
        3 => {
            let eptp = crate::vmm::vmx::vmread(crate::vmm::vmx::VMCS_CTRL_EPTP);
            crate::vmm::vmx::invept(2, eptp);
        }
        _ => {}
    }

    true
}
