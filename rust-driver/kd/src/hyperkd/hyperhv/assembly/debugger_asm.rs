use core::arch::{asm, global_asm};

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct GuestContext {
    pub rax: u64,
    pub rcx: u64,
    pub rdx: u64,
    pub rbx: u64,
    pub rsp: u64,
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
    pub rip: u64,
    pub rflags: u64,
}

impl GuestContext {
    pub fn new() -> Self {
        Self {
            rax: 0,
            rcx: 0,
            rdx: 0,
            rbx: 0,
            rsp: 0,
            rbp: 0,
            rsi: 0,
            rdi: 0,
            r8: 0,
            r9: 0,
            r10: 0,
            r11: 0,
            r12: 0,
            r13: 0,
            r14: 0,
            r15: 0,
            rip: 0,
            rflags: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct XmmRegisters {
    pub xmm0: u128,
    pub xmm1: u128,
    pub xmm2: u128,
    pub xmm3: u128,
    pub xmm4: u128,
    pub xmm5: u128,
}

impl XmmRegisters {
    pub fn new() -> Self {
        Self {
            xmm0: 0,
            xmm1: 0,
            xmm2: 0,
            xmm3: 0,
            xmm4: 0,
            xmm5: 0,
        }
    }
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct DebuggerContext {
    pub xmm_regs: XmmRegisters,
    pub mxcsr: u32,
    pub guest_context: GuestContext,
    pub rflags: u64,
}

impl DebuggerContext {
    pub fn new() -> Self {
        Self {
            xmm_regs: XmmRegisters::new(),
            mxcsr: 0,
            guest_context: GuestContext::new(),
            rflags: 0,
        }
    }
}

global_asm!(
    r#"
.set context_rax, 0x00
.set context_rcx, 0x08
.set context_rdx, 0x10
.set context_rbx, 0x18
.set context_rsp, 0x20
.set context_rbp, 0x28
.set context_rsi, 0x30
.set context_rdi, 0x38
.set context_r8,  0x40
.set context_r9,  0x48
.set context_r10, 0x50
.set context_r11, 0x58
.set context_r12, 0x60
.set context_r13, 0x68
.set context_r14, 0x70
.set context_r15, 0x78
.set context_rip, 0x80
.set context_rflags, 0x88

.global asm_save_guest_registers
asm_save_guest_registers:
    mov [rcx + context_rax], rax
    mov [rcx + context_rcx], rcx
    mov [rcx + context_rdx], rdx
    mov [rcx + context_rbx], rbx
    mov [rcx + context_rsp], rsp
    mov [rcx + context_rbp], rbp
    mov [rcx + context_rsi], rsi
    mov [rcx + context_rdi], rdi
    mov [rcx + context_r8],  r8
    mov [rcx + context_r9],  r9
    mov [rcx + context_r10], r10
    mov [rcx + context_r11], r11
    mov [rcx + context_r12], r12
    mov [rcx + context_r13], r13
    mov [rcx + context_r14], r14
    mov [rcx + context_r15], r15
    lea rax, [rip + 0]
    mov [rcx + context_rip], rax
    pushfq
    pop rax
    mov [rcx + context_rflags], rax
    ret

.global asm_restore_guest_registers
asm_restore_guest_registers:
    mov rax, [rcx + context_rax]
    mov rdx, [rcx + context_rdx]
    mov rbx, [rcx + context_rbx]
    mov rbp, [rcx + context_rbp]
    mov rsi, [rcx + context_rsi]
    mov rdi, [rcx + context_rdi]
    mov r8,  [rcx + context_r8]
    mov r9,  [rcx + context_r9]
    mov r10, [rcx + context_r10]
    mov r11, [rcx + context_r11]
    mov r12, [rcx + context_r12]
    mov r13, [rcx + context_r13]
    mov r14, [rcx + context_r14]
    mov r15, [rcx + context_r15]
    push qword ptr [rcx + context_rflags]
    popfq
    mov rsp, [rcx + context_rsp]
    mov rcx, [rcx + context_rcx]
    ret
"#
);

extern "C" {
    pub fn asm_save_guest_registers(context: *mut GuestContext);
    pub fn asm_restore_guest_registers(context: *const GuestContext);
}

#[inline(always)]
pub unsafe fn save_guest_registers(context: &mut GuestContext) {
    asm_save_guest_registers(context);
}

#[inline(always)]
pub unsafe fn restore_guest_registers(context: &GuestContext) {
    asm_restore_guest_registers(context);
}

global_asm!(
    r#"
.global asm_vmx_vmcall_impl
asm_vmx_vmcall_impl:
    push r10
    push r11
    push r12
    mov r10, 0x48564653
    mov r11, 0x564d43616c6c
    mov r12, 0x4e4f485950455256
    vmcall
    pop r12
    pop r11
    pop r10
    ret
"#,
);

extern "C" {
    fn asm_vmx_vmcall_impl(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64;
}

#[no_mangle]
pub unsafe extern "C" fn asm_vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64 {
    asm_vmx_vmcall_impl(vmcall_number, param1, param2, param3)
}

#[no_mangle]
pub unsafe extern "C" fn asm_vmxoff_handler() -> ! {
    asm!(
        "vmxoff",
        clobber_abi("system"),
    );

    core::arch::asm!("iretq", options(noreturn));
}

#[inline(always)]
pub unsafe fn save_xmm_registers(xmm_regs: &mut XmmRegisters) {
    asm!(
        "movaps xmm0, xmmword ptr [{0} + 0x000]",
        "movaps xmm1, xmmword ptr [{0} + 0x010]",
        "movaps xmm2, xmmword ptr [{0} + 0x020]",
        "movaps xmm3, xmmword ptr [{0} + 0x030]",
        "movaps xmm4, xmmword ptr [{0} + 0x040]",
        "movaps xmm5, xmmword ptr [{0} + 0x050]",
        in(reg) xmm_regs as *mut XmmRegisters,
    );
}

#[inline(always)]
pub unsafe fn restore_xmm_registers(xmm_regs: &XmmRegisters) {
    asm!(
        "movaps xmmword ptr [{0} + 0x000], xmm0",
        "movaps xmmword ptr [{0} + 0x010], xmm1",
        "movaps xmmword ptr [{0} + 0x020], xmm2",
        "movaps xmmword ptr [{0} + 0x030], xmm3",
        "movaps xmmword ptr [{0} + 0x040], xmm4",
        "movaps xmmword ptr [{0} + 0x050], xmm5",
        in(reg) xmm_regs as *const XmmRegisters,
    );
}

#[inline(always)]
pub unsafe fn read_mxcsr() -> u32 {
    let mxcsr: u32;
    asm!(
        "stmxcsr {0}",
        out(reg) mxcsr,
        options(nomem, nostack),
    );
    mxcsr
}

#[inline(always)]
pub unsafe fn write_mxcsr(mxcsr: u32) {
    asm!(
        "ldmxcsr {0}",
        in(reg) mxcsr,
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn save_rflags() -> u64 {
    let rflags: u64;
    asm!(
        "pushfq",
        "pop {0}",
        out(reg) rflags,
    );
    rflags
}

#[inline(always)]
pub unsafe fn restore_rflags(rflags: u64) {
    asm!(
        "push {0}",
        "popfq",
        in(reg) rflags,
    );
}

#[inline(always)]
pub unsafe fn enable_interrupts() {
    asm!("sti", options(nomem, nostack));
}

#[inline(always)]
pub unsafe fn disable_interrupts() {
    asm!("cli", options(nomem, nostack));
}

#[inline(always)]
pub unsafe fn save_interrupt_state() -> bool {
    let rflags: u64;
    asm!(
        "pushfq",
        "pop {0}",
        out(reg) rflags,
    );
    (rflags & 0x200) != 0
}

#[inline(always)]
pub unsafe fn restore_interrupt_state(enabled: bool) {
    if enabled {
        enable_interrupts();
    } else {
        disable_interrupts();
    }
}

#[inline(always)]
pub unsafe fn get_current_rip() -> u64 {
    let rip: u64;
    asm!(
        "lea {0}, [rip + 0]",
        out(reg) rip,
        options(nomem, nostack),
    );
    rip
}

#[inline(always)]
pub unsafe fn get_current_rsp() -> u64 {
    let rsp: u64;
    asm!(
        "mov {0}, rsp",
        out(reg) rsp,
        options(nomem, nostack),
    );
    rsp
}

#[inline(always)]
pub unsafe fn set_current_rsp(rsp: u64) {
    asm!(
        "mov rsp, {0}",
        in(reg) rsp,
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn read_segment_cs() -> u16 {
    let cs: u16;
    asm!(
        "mov {0:x}, cs",
        out(reg) cs,
        options(nomem, nostack),
    );
    cs
}

#[inline(always)]
pub unsafe fn read_segment_ds() -> u16 {
    let ds: u16;
    asm!(
        "mov {0:x}, ds",
        out(reg) ds,
        options(nomem, nostack),
    );
    ds
}

#[inline(always)]
pub unsafe fn read_segment_es() -> u16 {
    let es: u16;
    asm!(
        "mov {0:x}, es",
        out(reg) es,
        options(nomem, nostack),
    );
    es
}

#[inline(always)]
pub unsafe fn read_segment_fs() -> u16 {
    let fs: u16;
    asm!(
        "mov {0:x}, fs",
        out(reg) fs,
        options(nomem, nostack),
    );
    fs
}

#[inline(always)]
pub unsafe fn read_segment_gs() -> u16 {
    let gs: u16;
    asm!(
        "mov {0:x}, gs",
        out(reg) gs,
        options(nomem, nostack),
    );
    gs
}

#[inline(always)]
pub unsafe fn read_segment_ss() -> u16 {
    let ss: u16;
    asm!(
        "mov {0:x}, ss",
        out(reg) ss,
        options(nomem, nostack),
    );
    ss
}

#[inline(always)]
pub unsafe fn read_fs_base() -> u64 {
    let base: u64;
    asm!(
        "rdfsbase {0}",
        out(reg) base,
        options(nomem, nostack),
    );
    base
}

#[inline(always)]
pub unsafe fn write_fs_base(base: u64) {
    asm!(
        "wrfsbase {0}",
        in(reg) base,
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn read_gs_base() -> u64 {
    let base: u64;
    asm!(
        "rdgsbase {0}",
        out(reg) base,
        options(nomem, nostack),
    );
    base
}

#[inline(always)]
pub unsafe fn write_gs_base(base: u64) {
    asm!(
        "wrgsbase {0}",
        in(reg) base,
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn swapgs() {
    asm!(
        "swapgs",
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn sysret() {
    asm!(
        "sysret",
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn sysexit() {
    asm!(
        "sysexit",
        options(nomem, nostack),
    );
}

#[inline(always)]
pub unsafe fn iret() {
    asm!(
        "iretq",
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn call_far(selector: u16, offset: u64) {
    asm!(
        "push {0}",
        "push {1}",
        "lretq",
        in(reg) selector as u64,
        in(reg) offset,
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn jmp_far(selector: u16, offset: u64) -> ! {
    asm!(
        "push {0}",
        "push {1}",
        "retfq",
        in(reg) selector as u64,
        in(reg) offset,
        options(noreturn),
    );
}
