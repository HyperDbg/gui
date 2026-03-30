use core::arch::asm;
use core::arch::global_asm;
use core::arch::naked_asm;
use crate::hyperkd::VMX_CONTEXT;
use crate::hyperkd::VmxState;
use super::vmx::*;
use crate::hyperkd::hyperhv::HYPERDBG_SIGNATURE;

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct GuestRegisters {
    pub rax: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rbx: u64,
    pub rbp: u64,
    pub rsi: u64,
    pub rdi: u64,
    pub r8: u64,
    pub r9: u64,
    pub r10: u64,
    pub r11: u64,
    pub r12: u64,
    pub r13: u64,
    pub r14: u64,
    pub r15: u64,
    pub cr2: u64,
}

global_asm! {
    r#"
.global asm_set_cpuid_result
asm_set_cpuid_result:
    push rbx
    mov rax, rcx
    mov rbx, rdx
    mov rcx, r8
    mov rdx, r9
    pop rbx
    ret

.global asm_cpuid
asm_cpuid:
    push rbx
    mov eax, ecx
    mov ecx, edx
    cpuid
    mov [r8], eax
    mov [r9], ebx
    mov rax, [rsp + 0x30]
    mov [rax], ecx
    mov rax, [rsp + 0x38]
    mov [rax], edx
    pop rbx
    ret
"#,
}

extern "C" {
    fn asm_set_cpuid_result(eax: u64, ebx: u64, ecx: u64, edx: u64);
    fn asm_cpuid(leaf: u32, sub_leaf: u32, eax_out: *mut u32, ebx_out: *mut u32, ecx_out: *mut u32, edx_out: *mut u32);
}

#[no_mangle]
#[unsafe(naked)]
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
    guest_regs: *const GuestRegisters,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let exit_reason = vmread(VMCS_EXIT_REASON) as u32 & 0xFFFF;
    let exit_qualification = vmread(VMCS_EXIT_QUALIFICATION);
    let instruction_length = vmread(VMCS_EXIT_INSTRUCTION_LEN) as u32;

    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();

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
            
            if !guest_regs.is_null() {
                let regs = &*guest_regs;
                vcpu.guest_rax = regs.rax;
                vcpu.guest_rbx = regs.rbx;
                vcpu.guest_rcx = regs.rcx;
                vcpu.guest_rdx = regs.rdx;
                vcpu.guest_rsi = regs.rsi;
                vcpu.guest_rdi = regs.rdi;
                vcpu.guest_rbp = regs.rbp;
                vcpu.guest_r8 = regs.r8;
                vcpu.guest_r9 = regs.r9;
                vcpu.guest_r10 = regs.r10;
                vcpu.guest_r11 = regs.r11;
                vcpu.guest_r12 = regs.r12;
                vcpu.guest_r13 = regs.r13;
                vcpu.guest_r14 = regs.r14;
                vcpu.guest_r15 = regs.r15;
                vcpu.guest_cr2 = regs.cr2;
            }
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
        let current_rip = *guest_rip;
        *guest_rip = current_rip.wrapping_add(instruction_length as u64);
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
        37 => handle_mtf(guest_rsp, guest_rip, guest_rflags),
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

    match vector {
        0 => handle_divide_error(guest_rsp, guest_rip, guest_rflags),
        1 => handle_debug_exception(guest_rsp, guest_rip, guest_rflags),
        2 => handle_nmi(guest_rsp, guest_rip, guest_rflags),
        3 => handle_breakpoint(guest_rsp, guest_rip, guest_rflags),
        14 => handle_page_fault(exit_qualification, guest_rsp, guest_rip, guest_rflags),
        _ => true,
    }
}

unsafe fn handle_divide_error(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }
    let _ = (guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_debug_exception(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let rflags = *guest_rflags;
    *guest_rflags = rflags & !(1 << 8);

    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    let _ = (guest_rsp, guest_rip);
    true
}

unsafe fn handle_nmi(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let mut ctrl = vmread(VMCS_CTRL_PIN_BASED) as u32;
    ctrl |= PIN_BASED_EXEC_CTRL_NMI_EXITING as u32;
    vmwrite(VMCS_CTRL_PIN_BASED, ctrl as u64);
    let _ = (guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_breakpoint(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    let _ = (guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_page_fault(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let _cr2 = read_cr2();
    
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = false;
    }

    let _ = (exit_qualification, guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_external_interrupt() -> bool {
    let mut ctrl = vmread(VMCS_CTRL_PIN_BASED) as u32;
    
    if (ctrl & PIN_BASED_EXT_INT_EXITING as u32) != 0 {
        ctrl &= !(PIN_BASED_EXT_INT_EXITING as u32);
        vmwrite(VMCS_CTRL_PIN_BASED, ctrl as u64);
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
    let mut ctrl = vmread(VMCS_CTRL_PROC_BASED) as u32;
    ctrl &= !(PROC_BASED_INT_WINDOW_EXITING as u32);
    vmwrite(VMCS_CTRL_PROC_BASED, ctrl as u64);
    true
}

unsafe fn handle_nmi_window() -> bool {
    let mut ctrl = vmread(VMCS_CTRL_PROC_BASED) as u32;
    ctrl &= !(PROC_BASED_NMI_WINDOW_EXITING as u32);
    vmwrite(VMCS_CTRL_PROC_BASED, ctrl as u64);
    true
}

unsafe fn handle_cpuid(guest_rsp: *mut u64, guest_rip: *mut u64) -> bool {
    let leaf: u32;
    let sub_leaf: u32;

    asm!(
        "",
        out("eax") leaf,
        out("ecx") sub_leaf,
        options(nomem, preserves_flags),
    );

    let (eax, ebx, ecx, edx) = match leaf {
        0x40000000 => {
            (0x40000001, 0x52455059, 0x47424448, 0x00000000)
        }
        0x40000001 => {
            (HYPERDBG_SIGNATURE as u32, (HYPERDBG_SIGNATURE >> 32) as u32, 0x00000001, 0x00000000)
        }
        _ => {
            let mut eax_out: u32 = 0;
            let mut ebx_out: u32 = 0;
            let mut ecx_out: u32 = 0;
            let mut edx_out: u32 = 0;

            asm_cpuid(leaf, sub_leaf, &mut eax_out, &mut ebx_out, &mut ecx_out, &mut edx_out);

            (eax_out, ebx_out, ecx_out, edx_out)
        }
    };

    let eax_u64 = eax as u64;
    let ebx_u64 = ebx as u64;
    let ecx_u64 = ecx as u64;
    let edx_u64 = edx as u64;

    asm_set_cpuid_result(eax_u64, ebx_u64, ecx_u64, edx_u64);

    let _ = (guest_rsp, guest_rip);
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

    let result = match vmcall_number {
        0x1 => handle_vmxoff(guest_rsp, guest_rip, guest_rflags),
        0x2 => handle_test(guest_rsp, guest_rip, guest_rflags),
        _ => true,
    };
    let _ = (exit_qualification, rbx, rcx, rdx);
    result
}

unsafe fn handle_vmxoff(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();

    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.has_launched = false;
        vcpu.state = VmxState::Terminated;
    }

    let _ = (guest_rsp, guest_rip, guest_rflags);
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
    let _ = (guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_mtf(
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
    guest_rflags: *mut u64,
) -> bool {
    let core_id = crate::hyperkd::hyperhv::processor::get_current_processor_number();
    let mut context = VMX_CONTEXT.lock();
    if let Some(vcpu) = context.get_vcpu_mut(core_id) {
        vcpu.increment_rip = true;
    }
    let _ = (guest_rsp, guest_rip, guest_rflags);
    true
}

unsafe fn handle_ept_violation(exit_qualification: u64, guest_rip: *mut u64) -> bool {
    let _ = (exit_qualification, guest_rip);
    true
}

unsafe fn handle_ept_misconfig(exit_qualification: u64, guest_rip: *mut u64) -> bool {
    let _ = (exit_qualification, guest_rip);
    true
}

unsafe fn handle_invpcid(exit_qualification: u64) -> bool {
    let _ = exit_qualification;
    true
}

unsafe fn handle_cr_access(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let _ = (exit_qualification, guest_rsp, guest_rip);
    true
}

unsafe fn handle_dr_access(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let _ = (exit_qualification, guest_rsp, guest_rip);
    true
}

unsafe fn handle_io_instruction(
    exit_qualification: u64,
    guest_rsp: *mut u64,
    guest_rip: *mut u64,
) -> bool {
    let _ = (exit_qualification, guest_rsp, guest_rip);
    true
}

unsafe fn handle_rdmsr(exit_qualification: u64, guest_rsp: *mut u64) -> bool {
    let _ = (exit_qualification, guest_rsp);
    true
}

unsafe fn handle_wrmsr(exit_qualification: u64, guest_rsp: *mut u64) -> bool {
    let _ = (exit_qualification, guest_rsp);
    true
}
