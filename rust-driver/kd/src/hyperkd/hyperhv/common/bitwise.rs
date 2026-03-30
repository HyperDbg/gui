#![allow(dead_code)]

#[inline(always)]
pub fn check_bit(value: u64, bit: u32) -> bool {
    (value & (1u64 << bit)) != 0
}

#[inline(always)]
pub fn set_bit(value: u64, bit: u32) -> u64 {
    value | (1u64 << bit)
}

#[inline(always)]
pub fn clear_bit(value: u64, bit: u32) -> u64 {
    value & !(1u64 << bit)
}

#[inline(always)]
pub fn toggle_bit(value: u64, bit: u32) -> u64 {
    value ^ (1u64 << bit)
}

#[inline(always)]
pub fn get_bits(value: u64, start: u32, end: u32) -> u64 {
    let mask = ((1u64 << (end - start + 1)) - 1) << start;
    (value & mask) >> start
}

#[inline(always)]
pub fn set_bits(value: u64, start: u32, end: u32, bits: u64) -> u64 {
    let mask = ((1u64 << (end - start + 1)) - 1) << start;
    (value & !mask) | ((bits << start) & mask)
}

#[inline(always)]
pub fn is_power_of_two(value: u64) -> bool {
    value != 0 && (value & (value - 1)) == 0
}

#[inline(always)]
pub fn align_up(value: u64, alignment: u64) -> u64 {
    (value + alignment - 1) & !(alignment - 1)
}

#[inline(always)]
pub fn align_down(value: u64, alignment: u64) -> u64 {
    value & !(alignment - 1)
}

#[inline(always)]
pub fn page_align(value: u64) -> u64 {
    align_down(value, 0x1000)
}

#[inline(always)]
pub fn page_offset(value: u64) -> u64 {
    value & 0xFFF
}

#[inline(always)]
pub fn page_number(value: u64) -> u64 {
    value >> 12
}

#[inline(always)]
pub fn pml4_index(va: u64) -> u64 {
    (va >> 39) & 0x1FF
}

#[inline(always)]
pub fn pdpt_index(va: u64) -> u64 {
    (va >> 30) & 0x1FF
}

#[inline(always)]
pub fn pd_index(va: u64) -> u64 {
    (va >> 21) & 0x1FF
}

#[inline(always)]
pub fn pt_index(va: u64) -> u64 {
    (va >> 12) & 0x1FF
}

#[inline(always)]
pub fn log2(value: u64) -> u32 {
    63 - value.leading_zeros()
}
