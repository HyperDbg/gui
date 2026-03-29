#![no_std]

extern crate alloc;

use alloc::vec::Vec;
use alloc::string::String;
use alloc::format;

pub type Address = u64;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DisassemblerError {
    InvalidAddress,
    InvalidInstruction,
    BufferTooSmall,
    UnknownError,
}

impl core::fmt::Display for DisassemblerError {
    fn fmt(&self, f: &mut core::fmt::Formatter<'_>) -> core::fmt::Result {
        match self {
            DisassemblerError::InvalidAddress => write!(f, "Invalid address"),
            DisassemblerError::InvalidInstruction => write!(f, "Invalid instruction"),
            DisassemblerError::BufferTooSmall => write!(f, "Buffer too small"),
            DisassemblerError::UnknownError => write!(f, "Unknown error"),
        }
    }
}

impl core::error::Error for DisassemblerError {}

#[derive(Debug, Clone)]
pub struct Instruction {
    pub address: Address,
    pub bytes: Vec<u8>,
    pub mnemonic: String,
    pub operands: String,
    pub length: u8,
    pub category: String,
    pub is_branch: bool,
    pub is_call: bool,
    pub is_ret: bool,
    pub is_interrupt: bool,
}

#[derive(Debug, Clone)]
pub struct Disassembler {
    _dummy: u8,
}

impl Disassembler {
    pub fn new() -> Result<Self, DisassemblerError> {
        Ok(Self { _dummy: 0 })
    }

    pub fn decode(&self, address: Address, bytes: &[u8]) -> Result<Instruction, DisassemblerError> {
        if bytes.is_empty() {
            return Err(DisassemblerError::BufferTooSmall);
        }

        let (mnemonic, operands, length, category, is_branch, is_call, is_ret, is_interrupt) = 
            Self::decode_instruction(bytes);

        Ok(Instruction {
            address,
            bytes: bytes[..length as usize].to_vec(),
            mnemonic,
            operands,
            length,
            category,
            is_branch,
            is_call,
            is_ret,
            is_interrupt,
        })
    }

    pub fn decode_all(&self, start_address: Address, bytes: &[u8], max_instructions: usize) -> Vec<Instruction> {
        let mut instructions = Vec::new();
        let mut offset = 0usize;
        let mut address = start_address;

        while offset < bytes.len() && instructions.len() < max_instructions {
            if let Ok(insn) = self.decode(address, &bytes[offset..]) {
                offset += insn.length as usize;
                address += insn.length as u64;
                instructions.push(insn);
            } else {
                break;
            }
        }

        instructions
    }

    fn decode_instruction(bytes: &[u8]) -> (String, String, u8, String, bool, bool, bool, bool) {
        if bytes.is_empty() {
            return (String::from("db"), format!("0x{:02x}", bytes[0]), 1, String::from("data"), false, false, false, false);
        }

        let b = bytes[0];
        
        match b {
            0x90 => (String::from("nop"), String::new(), 1, String::from("misc"), false, false, false, false),
            0xC3 => (String::from("ret"), String::new(), 1, String::from("ret"), false, false, true, false),
            0xCC => (String::from("int3"), String::new(), 1, String::from("interrupt"), false, false, false, true),
            0xCD if bytes.len() > 1 => (String::from("int"), format!("0x{:02x}", bytes[1]), 2, String::from("interrupt"), false, false, false, true),
            0xCE => (String::from("into"), String::new(), 1, String::from("interrupt"), false, false, false, true),
            0xF1 => (String::from("int1"), String::new(), 1, String::from("interrupt"), false, false, false, true),
            
            0xE8 if bytes.len() >= 5 => {
                let offset = i32::from_le_bytes([bytes[1], bytes[2], bytes[3], bytes[4]]);
                let target = offset.wrapping_add(5);
                (String::from("call"), format!("0x{:x}", target as u64), 5, String::from("call"), true, true, false, false)
            }
            0xE9 if bytes.len() >= 5 => {
                let offset = i32::from_le_bytes([bytes[1], bytes[2], bytes[3], bytes[4]]);
                let target = offset.wrapping_add(5);
                (String::from("jmp"), format!("0x{:x}", target as u64), 5, String::from("branch"), true, false, false, false)
            }
            0xEB if bytes.len() >= 2 => {
                let offset = bytes[1] as i8;
                let target = offset.wrapping_add(2);
                (String::from("jmp"), format!("0x{:x}", target as u64), 2, String::from("branch"), true, false, false, false)
            }
            
            0x50..=0x57 => (String::from("push"), format!("r{}", (b - 0x50) as u8), 1, String::from("stack"), false, false, false, false),
            0x58..=0x5F => (String::from("pop"), format!("r{}", (b - 0x58) as u8), 1, String::from("stack"), false, false, false, false),
            
            0xB0..=0xB7 if bytes.len() >= 2 => {
                (String::from("mov"), format!("al, 0x{:02x}", bytes[1]), 2, String::from("data"), false, false, false, false)
            }
            0xB8..=0xBF if bytes.len() >= 5 => {
                let imm = u32::from_le_bytes([bytes[1], bytes[2], bytes[3], bytes[4]]);
                (String::from("mov"), format!("r{}, 0x{:x}", b - 0xB8, imm), 5, String::from("data"), false, false, false, false)
            }
            0x48 if bytes.len() > 1 && bytes[1] >= 0xB8 && bytes[1] <= 0xBF && bytes.len() >= 10 => {
                let imm = u64::from_le_bytes([bytes[2], bytes[3], bytes[4], bytes[5], bytes[6], bytes[7], bytes[8], bytes[9]]);
                (String::from("mov"), format!("r{}, 0x{:x}", bytes[1] - 0xB8, imm), 10, String::from("data"), false, false, false, false)
            }
            
            0x01 | 0x03 | 0x09 | 0x0B | 0x21 | 0x23 | 0x29 | 0x2B | 0x31 | 0x33 => {
                if bytes.len() >= 2 {
                    let modrm = bytes[1];
                    let reg = (modrm >> 3) & 0x7;
                    let rm = modrm & 0x7;
                    let mnemonic = match b {
                        0x01 | 0x03 => "add",
                        0x09 | 0x0B => "or",
                        0x21 | 0x23 => "and",
                        0x29 | 0x2B => "sub",
                        0x31 | 0x33 => "xor",
                        _ => "???",
                    };
                    let operands = if b & 0x2 == 0 { 
                        format!("r{}, r{}", reg, rm) 
                    } else { 
                        format!("r{}, r{}", rm, reg) 
                    };
                    (String::from(mnemonic), operands, 2, String::from("alu"), false, false, false, false)
                } else {
                    (String::from("db"), format!("0x{:02x}", b), 1, String::from("data"), false, false, false, false)
                }
            }
            
            0x74..=0x7F if bytes.len() >= 2 => {
                let offset = bytes[1] as i8;
                let mnemonic = match b {
                    0x74 => "jz",
                    0x75 => "jnz",
                    0x76 => "jbe",
                    0x77 => "ja",
                    0x78 => "js",
                    0x79 => "jns",
                    0x7A => "jp",
                    0x7B => "jnp",
                    0x7C => "jl",
                    0x7D => "jge",
                    0x7E => "jle",
                    0x7F => "jg",
                    _ => "j??",
                };
                let target = 2i32.wrapping_add(offset as i32);
                (String::from(mnemonic), format!("0x{:x}", target as u64), 2, String::from("branch"), true, false, false, false)
            }
            
            0x0F if bytes.len() > 1 => {
                match bytes[1] {
                    0x80..=0x8F if bytes.len() >= 6 => {
                        let offset = i32::from_le_bytes([bytes[2], bytes[3], bytes[4], bytes[5]]);
                        let mnemonic = match bytes[1] {
                            0x80 => "jo",
                            0x81 => "jno",
                            0x82 => "jb",
                            0x83 => "jae",
                            0x84 => "je",
                            0x85 => "jne",
                            0x86 => "jbe",
                            0x87 => "ja",
                            0x88 => "js",
                            0x89 => "jns",
                            0x8A => "jp",
                            0x8B => "jnp",
                            0x8C => "jl",
                            0x8D => "jge",
                            0x8E => "jle",
                            0x8F => "jg",
                            _ => "j??",
                        };
                        let target = 6i32.wrapping_add(offset);
                        (String::from(mnemonic), format!("0x{:x}", target as u64), 6, String::from("branch"), true, false, false, false)
                    }
                    _ => (String::from("db"), format!("0x{:02x}, 0x{:02x}", bytes[0], bytes[1]), 2, String::from("data"), false, false, false, false)
                }
            }
            
            0xFF if bytes.len() >= 2 => {
                let modrm = bytes[1];
                let ext = (modrm >> 3) & 0x7;
                match ext {
                    0 => (String::from("inc"), String::new(), 2, String::from("alu"), false, false, false, false),
                    1 => (String::from("dec"), String::new(), 2, String::from("alu"), false, false, false, false),
                    2 => (String::from("call"), String::new(), 2, String::from("call"), true, true, false, false),
                    4 => (String::from("jmp"), String::new(), 2, String::from("branch"), true, false, false, false),
                    6 => (String::from("push"), String::new(), 2, String::from("stack"), false, false, false, false),
                    _ => (String::from("ff"), format!("/{}", ext), 2, String::from("unknown"), false, false, false, false)
                }
            }
            
            _ => (String::from("db"), format!("0x{:02x}", b), 1, String::from("data"), false, false, false, false)
        }
    }
}

#[derive(Debug, Clone)]
pub struct DisassemblerBuilder {
    _dummy: u8,
}

impl DisassemblerBuilder {
    pub fn new() -> Self {
        Self { _dummy: 0 }
    }

    pub fn build(&self) -> Result<Disassembler, DisassemblerError> {
        Disassembler::new()
    }
}

impl Default for DisassemblerBuilder {
    fn default() -> Self {
        Self::new()
    }
}
