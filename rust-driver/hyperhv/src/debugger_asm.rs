use core::arch::asm;

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

extern "C" {
    fn vmx_vmexit_handler(context: *mut DebuggerContext) -> bool;
    fn vmx_vmxoff_handler();
}

#[no_mangle]
pub unsafe extern "C" fn asm_vmexit_handler() -> ! {
    let mut context: DebuggerContext = DebuggerContext::new();

    asm!(
        "sub rsp, 0x110",
        
        "movaps xmmword ptr [rsp + 0x000], xmm0",
        "movaps xmmword ptr [rsp + 0x010], xmm1",
        "movaps xmmword ptr [rsp + 0x020], xmm2",
        "movaps xmmword ptr [rsp + 0x030], xmm3",
        "movaps xmmword ptr [rsp + 0x040], xmm4",
        "movaps xmmword ptr [rsp + 0x050], xmm5",
        
        "stmxcsr dword ptr [rsp + 0x100]",
        
        "pushfq",
        
        "mov [rsp + 0x108], rax",
        "mov [rsp + 0x110], rcx",
        "mov [rsp + 0x118], rdx",
        "mov [rsp + 0x120], rbx",
        "mov [rsp + 0x128], rbp",
        "mov [rsp + 0x130], rsi",
        "mov [rsp + 0x138], rdi",
        "mov [rsp + 0x140], r8",
        "mov [rsp + 0x148], r9",
        "mov [rsp + 0x150], r10",
        "mov [rsp + 0x158], r11",
        "mov [rsp + 0x160], r12",
        "mov [rsp + 0x168], r13",
        "mov [rsp + 0x170], r14",
        "mov [rsp + 0x178], r15",
        
        inlateout("rsp") rsp => _,
        
        clobber_abi("system"),
    );

    context.xmm_regs.xmm0 = *(0 as *const u128).add(0);
    context.xmm_regs.xmm1 = *(0 as *const u128).add(1);
    context.xmm_regs.xmm2 = *(0 as *const u128).add(2);
    context.xmm_regs.xmm3 = *(0 as *const u128).add(3);
    context.xmm_regs.xmm4 = *(0 as *const u128).add(4);
    context.xmm_regs.xmm5 = *(0 as *const u128).add(5);
    context.mxcsr = *(0 as *const u32).add(64);
    context.rflags = *(0 as *const u64).add(34);
    context.guest_context.rax = *(0 as *const u64).add(36);
    context.guest_context.rcx = *(0 as *const u64).add(37);
    context.guest_context.rdx = *(0 as *const u64).add(38);
    context.guest_context.rbx = *(0 as *const u64).add(39);
    context.guest_context.rbp = *(0 as *const u64).add(40);
    context.guest_context.rsi = *(0 as *const u64).add(41);
    context.guest_context.rdi = *(0 as *const u64).add(42);
    context.guest_context.r8 = *(0 as *const u64).add(43);
    context.guest_context.r9 = *(0 as *const u64).add(44);
    context.guest_context.r10 = *(0 as *const u64).add(45);
    context.guest_context.r11 = *(0 as *const u64).add(46);
    context.guest_context.r12 = *(0 as *const u64).add(47);
    context.guest_context.r13 = *(0 as *const u64).add(48);
    context.guest_context.r14 = *(0 as *const u64).add(49);
    context.guest_context.r15 = *(0 as *const u64).add(50);

    let should_vmxoff = vmx_vmexit_handler(&mut context);

    if should_vmxoff {
        asm_vmrestore_and_vmxoff();
    } else {
        asm_vmrestore_and_vmresume(&context);
    }

    core::arch::asm!("ud2", options(noreturn));
}

#[inline(never)]
unsafe fn asm_vmrestore_and_vmxoff() -> ! {
    asm!(
        "ldmxcsr dword ptr [rsp + 0x100]",
        
        "movaps xmm0, xmmword ptr [rsp + 0x000]",
        "movaps xmm1, xmmword ptr [rsp + 0x010]",
        "movaps xmm2, xmmword ptr [rsp + 0x020]",
        "movaps xmm3, xmmword ptr [rsp + 0x030]",
        "movaps xmm4, xmmword ptr [rsp + 0x040]",
        "movaps xmm5, xmmword ptr [rsp + 0x050]",
        
        "add rsp, 0x110",
        
        "popfq",
        
        clobber_abi("system"),
    );

    vmx_vmxoff_handler();

    core::arch::asm!("ud2", options(noreturn));
}

#[inline(never)]
unsafe fn asm_vmrestore_and_vmresume(context: *const DebuggerContext) -> ! {
    asm!(
        "ldmxcsr dword ptr [rsp + 0x100]",
        
        "movaps xmm0, xmmword ptr [rsp + 0x000]",
        "movaps xmm1, xmmword ptr [rsp + 0x010]",
        "movaps xmm2, xmmword ptr [rsp + 0x020]",
        "movaps xmm3, xmmword ptr [rsp + 0x030]",
        "movaps xmm4, xmmword ptr [rsp + 0x040]",
        "movaps xmm5, xmmword ptr [rsp + 0x050]",
        
        "mov rax, [rsp + 0x108]",
        "mov rcx, [rsp + 0x110]",
        "mov rdx, [rsp + 0x118]",
        "mov rbx, [rsp + 0x120]",
        "mov rbp, [rsp + 0x128]",
        "mov rsi, [rsp + 0x130]",
        "mov rdi, [rsp + 0x138]",
        "mov r8, [rsp + 0x140]",
        "mov r9, [rsp + 0x148]",
        "mov r10, [rsp + 0x150]",
        "mov r11, [rsp + 0x158]",
        "mov r12, [rsp + 0x160]",
        "mov r13, [rsp + 0x168]",
        "mov r14, [rsp + 0x170]",
        "mov r15, [rsp + 0x178]",
        
        "add rsp, 0x110",
        
        "popfq",
        
        "vmresume",
        
        clobber_abi("system"),
    );

    core::arch::asm!("ud2", options(noreturn));
}

#[no_mangle]
pub unsafe extern "C" fn asm_vmx_vmcall(vmcall_number: u64, param1: u64, param2: u64, param3: u64) -> u64 {
    let result: u64;

    asm!(
        "push r10",
        "push r11",
        "push r12",
        "mov r10, 0x48564653",
        "mov r11, 0x564d43616c6c",
        "mov r12, 0x4e4f485950455256",
        "vmcall",
        "pop r12",
        "pop r11",
        "pop r10",
        in("rax") vmcall_number,
        in("rcx") param1,
        in("rdx") param2,
        in("r8") param3,
        out("rax") result,
        clobber_abi("system"),
    );

    result
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
pub unsafe fn save_guest_registers(context: &mut GuestContext) {
    asm!(
        "mov {0}, rax",
        "mov {1}, rcx",
        "mov {2}, rdx",
        "mov {3}, rbx",
        "mov {4}, rsp",
        "mov {5}, rbp",
        "mov {6}, rsi",
        "mov {7}, rdi",
        "mov {8}, r8",
        "mov {9}, r9",
        "mov {10}, r10",
        "mov {11}, r11",
        "mov {12}, r12",
        "mov {13}, r13",
        "mov {14}, r14",
        "mov {15}, r15",
        out(reg) context.rax,
        out(reg) context.rcx,
        out(reg) context.rdx,
        out(reg) context.rbx,
        out(reg) context.rsp,
        out(reg) context.rbp,
        out(reg) context.rsi,
        out(reg) context.rdi,
        out(reg) context.r8,
        out(reg) context.r9,
        out(reg) context.r10,
        out(reg) context.r11,
        out(reg) context.r12,
        out(reg) context.r13,
        out(reg) context.r14,
        out(reg) context.r15,
        clobber_abi("system"),
    );

    let rip: u64;
    let rflags: u64;
    asm!(
        "lea {0}, [rip + 0]",
        "pushfq",
        "pop {1}",
        out(reg) rip,
        out(reg) rflags,
        options(nomem, nostack),
    );

    context.rip = rip;
    context.rflags = rflags;
}

#[inline(always)]
pub unsafe fn restore_guest_registers(context: &GuestContext) {
    asm!(
        "mov rax, {0}",
        "mov rcx, {1}",
        "mov rdx, {2}",
        "mov rbx, {3}",
        "mov rsp, {4}",
        "mov rbp, {5}",
        "mov rsi, {6}",
        "mov rdi, {7}",
        "mov r8, {8}",
        "mov r9, {9}",
        "mov r10, {10}",
        "mov r11, {11}",
        "mov r12, {12}",
        "mov r13, {13}",
        "mov r14, {14}",
        "mov r15, {15}",
        in(reg) context.rax,
        in(reg) context.rcx,
        in(reg) context.rdx,
        in(reg) context.rbx,
        in(reg) context.rsp,
        in(reg) context.rbp,
        in(reg) context.rsi,
        in(reg) context.rdi,
        in(reg) context.r8,
        in(reg) context.r9,
        in(reg) context.r10,
        in(reg) context.r11,
        in(reg) context.r12,
        in(reg) context.r13,
        in(reg) context.r14,
        in(reg) context.r15,
        clobber_abi("system"),
    );

    asm!(
        "push {0}",
        "popfq",
        in(reg) context.rflags,
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn save_xmm_registers(xmm_regs: &mut XmmRegisters) {
    asm!(
        "movaps xmm0, xmmword ptr [rsp + 0x000]",
        "movaps xmm1, xmmword ptr [rsp + 0x010]",
        "movaps xmm2, xmmword ptr [rsp + 0x020]",
        "movaps xmm3, xmmword ptr [rsp + 0x030]",
        "movaps xmm4, xmmword ptr [rsp + 0x040]",
        "movaps xmm5, xmmword ptr [rsp + 0x050]",
        in(reg) xmm_regs,
        clobber_abi("system"),
    );
}

#[inline(always)]
pub unsafe fn restore_xmm_registers(xmm_regs: &XmmRegisters) {
    asm!(
        "movaps xmmword ptr [rsp + 0x000], xmm0",
        "movaps xmmword ptr [rsp + 0x010], xmm1",
        "movaps xmmword ptr [rsp + 0x020], xmm2",
        "movaps xmmword ptr [rsp + 0x030], xmm3",
        "movaps xmmword ptr [rsp + 0x040], xmm4",
        "movaps xmmword ptr [rsp + 0x050], xmm5",
        in(reg) xmm_regs,
        clobber_abi("system"),
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
        clobber_abi("system"),
    );
    rflags
}

#[inline(always)]
pub unsafe fn restore_rflags(rflags: u64) {
    asm!(
        "push {0}",
        "popfq",
        in(reg) rflags,
        clobber_abi("system"),
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
        clobber_abi("system"),
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
        "mov {0}, cs",
        out(reg) cs,
        options(nomem, nostack),
    );
    cs
}

#[inline(always)]
pub unsafe fn read_segment_ds() -> u16 {
    let ds: u16;
    asm!(
        "mov {0}, ds",
        out(reg) ds,
        options(nomem, nostack),
    );
    ds
}

#[inline(always)]
pub unsafe fn read_segment_es() -> u16 {
    let es: u16;
    asm!(
        "mov {0}, es",
        out(reg) es,
        options(nomem, nostack),
    );
    es
}

#[inline(always)]
pub unsafe fn read_segment_fs() -> u16 {
    let fs: u16;
    asm!(
        "mov {0}, fs",
        out(reg) fs,
        options(nomem, nostack),
    );
    fs
}

#[inline(always)]
pub unsafe fn read_segment_gs() -> u16 {
    let gs: u16;
    asm!(
        "mov {0}, gs",
        out(reg) gs,
        options(nomem, nostack),
    );
    gs
}

#[inline(always)]
pub unsafe fn read_segment_ss() -> u16 {
    let ss: u16;
    asm!(
        "mov {0}, ss",
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
