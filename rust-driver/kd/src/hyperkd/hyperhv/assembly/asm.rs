use core::arch::asm;

#[inline(always)]
pub unsafe fn read_msr(msr: u32) -> u64 {
    let low: u32;
    let high: u32;
    asm!(
        "rdmsr",
        inlateout("ecx") msr => _,
        out("eax") low,
        out("edx") high,
        options(nomem, nostack, pure)
    );
    ((high as u64) << 32) | (low as u64)
}

#[inline(always)]
pub unsafe fn write_msr(msr: u32, value: u64) {
    let low = value as u32;
    let high = (value >> 32) as u32;
    asm!(
        "wrmsr",
        inlateout("ecx") msr => _,
        in("eax") low,
        in("edx") high,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr0() -> u64 {
    let value: u64;
    asm!(
        "mov {}, cr0",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr0(value: u64) {
    asm!(
        "mov cr0, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr2() -> u64 {
    let value: u64;
    asm!(
        "mov {}, cr2",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr2(value: u64) {
    asm!(
        "mov cr2, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr3() -> u64 {
    let value: u64;
    asm!(
        "mov {}, cr3",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr3(value: u64) {
    asm!(
        "mov cr3, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr4() -> u64 {
    let value: u64;
    asm!(
        "mov {}, cr4",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr4(value: u64) {
    asm!(
        "mov cr4, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_cr8() -> u64 {
    let value: u64;
    asm!(
        "mov {}, cr8",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_cr8(value: u64) {
    asm!(
        "mov cr8, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr0() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr0",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr0(value: u64) {
    asm!(
        "mov dr0, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr1() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr1",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr1(value: u64) {
    asm!(
        "mov dr1, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr2() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr2",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr2(value: u64) {
    asm!(
        "mov dr2, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr3() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr3",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr3(value: u64) {
    asm!(
        "mov dr3, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr6() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr6",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr6(value: u64) {
    asm!(
        "mov dr6, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr7() -> u64 {
    let value: u64;
    asm!(
        "mov {}, dr7",
        out(reg) value,
        options(nomem, nostack, pure)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr7(value: u64) {
    asm!(
        "mov dr7, {}",
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn read_tsc() -> u64 {
    let high: u32;
    let low: u32;
    asm!(
        "rdtsc",
        out("eax") low,
        out("edx") high,
        options(nomem, nostack, pure)
    );
    ((high as u64) << 32) | (low as u64)
}

#[inline(always)]
pub unsafe fn read_tscp(aux: &mut u32) -> u64 {
    let high: u32;
    let low: u32;
    asm!(
        "rdtscp",
        out("eax") low,
        out("edx") high,
        out("ecx") *aux,
        options(nomem, nostack, pure)
    );
    ((high as u64) << 32) | (low as u64)
}

#[inline(always)]
pub unsafe fn cpuid(leaf: u32, sub_leaf: u32) -> (u32, u32, u32, u32) {
    let mut eax: u32 = leaf;
    let mut ebx: u32;
    let mut ecx: u32 = sub_leaf;
    let mut edx: u32;
    asm!(
        "push rbx",
        "cpuid",
        "mov {0}, ebx",
        "pop rbx",
        out(reg) ebx,
        inout("eax") eax,
        inout("ecx") ecx,
        out("edx") edx,
        options(nomem, nostack)
    );
    (eax, ebx, ecx, edx)
}

#[inline(always)]
pub unsafe fn invlpg(address: u64) {
    asm!(
        "invlpg [{}]",
        in(reg) address,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn invpcid(type_: u32, descriptor: u64) {
    asm!(
        "invpcid {}, [{}]",
        in(reg) type_,
        in(reg) descriptor,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn sfence() {
    asm!(
        "sfence",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn lfence() {
    asm!(
        "lfence",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn mfence() {
    asm!(
        "mfence",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn pause() {
    asm!(
        "pause",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn cli() {
    asm!(
        "cli",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn sti() {
    asm!(
        "sti",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn hlt() {
    asm!(
        "hlt",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn nop() {
    asm!(
        "nop",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn int3() {
    asm!(
        "int3",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn ud2() {
    asm!(
        "ud2",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn rdtsc() -> u64 {
    read_tsc()
}

#[inline(always)]
pub unsafe fn rdtscp(aux: &mut u32) -> u64 {
    read_tscp(aux)
}

#[inline(always)]
pub unsafe fn wbinvd() {
    asm!(
        "wbinvd",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn clflush(address: u64) {
    asm!(
        "clflush [{}]",
        in(reg) address,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn clwb(address: u64) {
    asm!(
        "clwb [{}]",
        in(reg) address,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn clflushopt(address: u64) {
    asm!(
        "clflushopt [{}]",
        in(reg) address,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn xsaveopt(ptr: *mut u8, mask: u64) {
    asm!(
        "xsaveopt [{}], {}",
        in(reg) ptr,
        in(reg) mask,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn xrstor(ptr: *const u8, mask: u64) {
    asm!(
        "xrstor [{}], {}",
        in(reg) ptr,
        in(reg) mask,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn xgetbv(index: u32) -> u64 {
    let low: u32;
    let high: u32;
    asm!(
        "xgetbv",
        inlateout("ecx") index => _,
        out("eax") low,
        out("edx") high,
        options(nomem, nostack, pure)
    );
    ((high as u64) << 32) | (low as u64)
}

#[inline(always)]
pub unsafe fn xsetbv(index: u32, value: u64) {
    let low = value as u32;
    let high = (value >> 32) as u32;
    asm!(
        "xsetbv",
        inlateout("ecx") index => _,
        in("eax") low,
        in("edx") high,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn vmxon(region: u64) -> bool {
    let success: u8;
    asm!(
        "vmxon [{}]",
        "setne al",
        in(reg) region,
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmxoff() {
    asm!(
        "vmxoff",
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn vmclear(region: u64) -> bool {
    let success: u8;
    asm!(
        "vmclear [{}]",
        "setne al",
        in(reg) region,
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmptrld(region: u64) -> bool {
    let success: u8;
    asm!(
        "vmptrld [{}]",
        "setne al",
        in(reg) region,
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmptrst(region: *mut u64) -> bool {
    let success: u8;
    asm!(
        "vmptrst [{}]",
        "setne al",
        in(reg) region,
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmlaunch() -> bool {
    let success: u8;
    asm!(
        "vmlaunch",
        "setne al",
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmresume() -> bool {
    let success: u8;
    asm!(
        "vmresume",
        "setne al",
        out("al") success,
        options(nostack)
    );
    success != 0
}

#[inline(always)]
pub unsafe fn vmread(field: u32) -> u64 {
    let mut value: u64;
    asm!(
        "vmread {}, {}",
        in(reg) field,
        out(reg) value,
        options(nomem, nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn vmwrite(field: u32, value: u64) {
    asm!(
        "vmwrite {}, {}",
        in(reg) field,
        in(reg) value,
        options(nomem, nostack)
    );
}

#[inline(always)]
pub unsafe fn invept(type_: u32, descriptor: u64) {
    asm!(
        "invept {}, [{}]",
        in(reg) type_,
        in(reg) descriptor,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn invvpid(type_: u32, descriptor: u64) {
    asm!(
        "invvpid {}, [{}]",
        in(reg) type_,
        in(reg) descriptor,
        options(nostack)
    );
}
