#![allow(dead_code)]


#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebugRegisterType {
    BreakOnInstructionFetch = 0,
    BreakOnWriteOnly = 1,
    BreakOnIoReadWrite = 2,
    BreakOnReadAndWrite = 3,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DebugRegister {
    Dr0 = 0,
    Dr1 = 1,
    Dr2 = 2,
    Dr3 = 3,
}

#[repr(C)]
#[derive(Clone, Copy, Default)]
pub struct Dr6 {
    pub value: u64,
}

impl Dr6 {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn is_breakpoint_triggered(&self, dr: DebugRegister) -> bool {
        let mask = 1 << (dr as u64);
        (self.value & mask) != 0
    }

    pub fn clear_breakpoint(&mut self, dr: DebugRegister) {
        let mask = !(1 << (dr as u64));
        self.value &= mask;
    }

    pub fn is_single_step(&self) -> bool {
        (self.value & (1 << 14)) != 0
    }

    pub fn clear_single_step(&mut self) {
        self.value &= !(1 << 14);
    }

    pub fn is_task_switch(&self) -> bool {
        (self.value & (1 << 15)) != 0
    }
}

#[repr(C)]
#[derive(Clone, Copy)]
pub struct Dr7 {
    pub value: u64,
}

impl Dr7 {
    pub fn new() -> Self {
        Self { value: 0 }
    }

    pub fn set_local_breakpoint(&mut self, dr: DebugRegister, enabled: bool) {
        let bit = dr as usize * 2;
        if enabled {
            self.value |= 1 << bit;
        } else {
            self.value &= !(1 << bit);
        }
    }

    pub fn set_global_breakpoint(&mut self, dr: DebugRegister, enabled: bool) {
        let bit = dr as usize * 2 + 1;
        if enabled {
            self.value |= 1 << bit;
        } else {
            self.value &= !(1 << bit);
        }
    }

    pub fn set_read_write(&mut self, dr: DebugRegister, rw: DebugRegisterType) {
        let shift = 16 + (dr as usize * 4);
        let mask = 0x3u64 << shift;
        self.value = (self.value & !mask) | ((rw as u64) << shift);
    }

    pub fn set_length(&mut self, dr: DebugRegister, len: u8) {
        let shift = 18 + (dr as usize * 4);
        let mask = 0x3u64 << shift;
        let len_encoded = match len {
            1 => 0b00,
            2 => 0b01,
            4 => 0b10,
            8 => 0b11,
            _ => 0b00,
        };
        self.value = (self.value & !mask) | ((len_encoded as u64) << shift);
    }

    pub fn set_local_exact_breakpoint(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1 << 8;
        } else {
            self.value &= !(1 << 8);
        }
    }

    pub fn set_global_exact_breakpoint(&mut self, enabled: bool) {
        if enabled {
            self.value |= 1 << 9;
        } else {
            self.value &= !(1 << 9);
        }
    }

    pub fn is_breakpoint_enabled(&self, dr: DebugRegister) -> bool {
        let mask = 0x3u64 << (dr as usize * 2);
        (self.value & mask) != 0
    }

    pub fn get_read_write(&self, dr: DebugRegister) -> DebugRegisterType {
        let shift = 16 + (dr as usize * 4);
        let rw = (self.value >> shift) & 0x3;
        match rw {
            0 => DebugRegisterType::BreakOnInstructionFetch,
            1 => DebugRegisterType::BreakOnWriteOnly,
            2 => DebugRegisterType::BreakOnIoReadWrite,
            3 => DebugRegisterType::BreakOnReadAndWrite,
            _ => DebugRegisterType::BreakOnInstructionFetch,
        }
    }

    pub fn get_length(&self, dr: DebugRegister) -> u8 {
        let shift = 18 + (dr as usize * 4);
        let len = (self.value >> shift) & 0x3;
        match len {
            0b00 => 1,
            0b01 => 2,
            0b10 => 4,
            0b11 => 8,
            _ => 1,
        }
    }
}

impl Default for Dr7 {
    fn default() -> Self {
        Self::new()
    }
}

#[inline(always)]
pub unsafe fn read_dr0() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr0",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr0(value: u64) {
    core::arch::asm!(
        "mov dr0, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr1() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr1",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr1(value: u64) {
    core::arch::asm!(
        "mov dr1, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr2() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr2",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr2(value: u64) {
    core::arch::asm!(
        "mov dr2, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr3() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr3",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr3(value: u64) {
    core::arch::asm!(
        "mov dr3, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr6() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr6",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr6(value: u64) {
    core::arch::asm!(
        "mov dr6, {0}",
        in(reg) value,
        options(nostack)
    );
}

#[inline(always)]
pub unsafe fn read_dr7() -> u64 {
    let value: u64;
    core::arch::asm!(
        "mov {0}, dr7",
        out(reg) value,
        options(nostack)
    );
    value
}

#[inline(always)]
pub unsafe fn write_dr7(value: u64) {
    core::arch::asm!(
        "mov dr7, {0}",
        in(reg) value,
        options(nostack)
    );
}

pub unsafe fn set_debug_register(
    dr: DebugRegister,
    action: DebugRegisterType,
    target_address: u64,
) -> bool {
    if dr as u8 >= 4 {
        return false;
    }

    match dr {
        DebugRegister::Dr0 => write_dr0(target_address),
        DebugRegister::Dr1 => write_dr1(target_address),
        DebugRegister::Dr2 => write_dr2(target_address),
        DebugRegister::Dr3 => write_dr3(target_address),
    }

    let mut dr7 = Dr7 { value: read_dr7() };

    dr7.value |= 1 << 10;

    dr7.set_local_exact_breakpoint(true);
    dr7.set_global_exact_breakpoint(true);

    dr7.set_local_breakpoint(dr, true);
    dr7.set_global_breakpoint(dr, true);
    dr7.set_read_write(dr, action);
    dr7.set_length(dr, 1);

    write_dr7(dr7.value);

    true
}

pub unsafe fn clear_debug_register(dr: DebugRegister) -> bool {
    if dr as u8 >= 4 {
        return false;
    }

    let mut dr7 = Dr7 { value: read_dr7() };

    dr7.set_local_breakpoint(dr, false);
    dr7.set_global_breakpoint(dr, false);

    write_dr7(dr7.value);

    true
}

pub unsafe fn get_debug_register_address(dr: DebugRegister) -> u64 {
    match dr {
        DebugRegister::Dr0 => read_dr0(),
        DebugRegister::Dr1 => read_dr1(),
        DebugRegister::Dr2 => read_dr2(),
        DebugRegister::Dr3 => read_dr3(),
    }
}

pub unsafe fn read_dr6_value() -> Dr6 {
    Dr6 { value: read_dr6() }
}

pub unsafe fn write_dr6_value(dr6: Dr6) {
    write_dr6(dr6.value);
}

pub unsafe fn read_dr7_value() -> Dr7 {
    Dr7 { value: read_dr7() }
}

pub unsafe fn write_dr7_value(dr7: Dr7) {
    write_dr7(dr7.value);
}

pub unsafe fn clear_all_debug_registers() {
    write_dr0(0);
    write_dr1(0);
    write_dr2(0);
    write_dr3(0);

    write_dr6(0xFFFF0FF0);
    write_dr7(0x400);
}

pub struct HardwareBreakpoint {
    pub dr: DebugRegister,
    pub address: u64,
    pub action: DebugRegisterType,
    pub length: u8,
    pub enabled: bool,
}

impl HardwareBreakpoint {
    pub fn new(dr: DebugRegister, address: u64, action: DebugRegisterType) -> Self {
        Self {
            dr,
            address,
            action,
            length: 1,
            enabled: false,
        }
    }

    pub unsafe fn apply(&mut self) -> bool {
        if self.enabled {
            return true;
        }

        if set_debug_register(self.dr, self.action, self.address) {
            self.enabled = true;
            true
        } else {
            false
        }
    }

    pub unsafe fn remove(&mut self) -> bool {
        if !self.enabled {
            return true;
        }

        if clear_debug_register(self.dr) {
            self.enabled = false;
            true
        } else {
            false
        }
    }
}

pub static DEBUG_REGISTERS_LOCK: spin::Mutex<()> = spin::Mutex::new(());

pub unsafe fn apply_hardware_breakpoint_safely(bp: &mut HardwareBreakpoint) -> bool {
    let _lock = DEBUG_REGISTERS_LOCK.lock();
    bp.apply()
}

pub unsafe fn remove_hardware_breakpoint_safely(bp: &mut HardwareBreakpoint) -> bool {
    let _lock = DEBUG_REGISTERS_LOCK.lock();
    bp.remove()
}
