#![no_std]

extern crate alloc;

use alloc::vec::Vec;
use alloc::string::String;
use alloc::format;
use zydis::{Decoder, Formatter, InputData, MachineMode, Opcode, Register, VisibleOperands};

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
    decoder: Decoder<MachineMode>,
    formatter: Formatter,
}

impl Disassembler {
    pub fn new() -> Result<Self, DisassemblerError> {
        let decoder = Decoder::new64();
        let formatter = Formatter::intel();
        
        Ok(Self {
            decoder,
            formatter,
        })
    }

    pub fn decode(&self, address: Address, bytes: &[u8]) -> Result<Instruction, DisassemblerError> {
        if bytes.is_empty() {
            return Err(DisassemblerError::BufferTooSmall);
        }

        let mut input = InputData::new(bytes, address);
        
        match self.decoder.decode::<VisibleOperands>(&mut input) {
            Ok((addr, raw_bytes, insn)) => {
                let mnemonic = insn.mnemonic.to_string();
                let operands = self.format_operands(addr, &insn);
                let length = raw_bytes.len() as u8;
                let category = insn.category.to_string();
                
                let is_branch = insn.is_branch();
                let is_call = insn.mnemonic == Opcode::CALL;
                let is_ret = insn.mnemonic == Opcode::RET;
                let is_interrupt = insn.mnemonic == Opcode::INT || 
                                  insn.mnemonic == Opcode::INT1 || 
                                  insn.mnemonic == Opcode::INT3;

                Ok(Instruction {
                    address,
                    bytes: raw_bytes.to_vec(),
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
            Err(_) => Err(DisassemblerError::InvalidInstruction),
        }
    }

    pub fn decode_all(&self, address: Address, bytes: &[u8], max_instructions: usize) -> Vec<Instruction> {
        let mut instructions = Vec::new();
        let mut current_address = address;
        let mut remaining_bytes = bytes;

        for _ in 0..max_instructions {
            if remaining_bytes.is_empty() {
                break;
            }

            match self.decode(current_address, remaining_bytes) {
                Ok(insn) => {
                    let insn_len = insn.length as usize;
                    current_address += insn_len as u64;
                    remaining_bytes = &remaining_bytes[insn_len..];
                    instructions.push(insn);
                }
                Err(_) => {
                    break;
                }
            }
        }

        instructions
    }

    pub fn decode_function(&self, address: Address, bytes: &[u8]) -> Vec<Instruction> {
        let mut instructions = Vec::new();
        let mut current_address = address;
        let mut remaining_bytes = bytes;
        let mut ret_found = false;

        while !ret_found && !remaining_bytes.is_empty() {
            match self.decode(current_address, remaining_bytes) {
                Ok(insn) => {
                    ret_found = insn.is_ret;
                    let insn_len = insn.length as usize;
                    current_address += insn_len as u64;
                    remaining_bytes = &remaining_bytes[insn_len..];
                    instructions.push(insn);
                }
                Err(_) => {
                    break;
                }
            }
        }

        instructions
    }

    pub fn format(&self, insn: &Instruction) -> String {
        format!("0x{:016X} {}", insn.address, self.format_instruction(insn))
    }

    pub fn format_instruction(&self, insn: &Instruction) -> String {
        if insn.operands.is_empty() {
            insn.mnemonic.clone()
        } else {
            format!("{} {}", insn.mnemonic, insn.operands)
        }
    }

    pub fn format_operands(&self, address: Address, insn: &zydis::DecodedInstruction) -> String {
        match self.formatter.format(Some(address), insn) {
            Ok(formatted) => formatted,
            Err(_) => String::new(),
        }
    }

    pub fn get_branch_target(&self, insn: &Instruction) -> Option<Address> {
        if !insn.is_branch {
            return None;
        }

        if insn.operands.starts_with("0x") {
            let addr_str = insn.operands
                .split_whitespace()
                .next()?
                .trim_start_matches("0x");
            u64::from_str_radix(addr_str, 16).ok()
        } else {
            None
        }
    }

    pub fn find_call_targets(&self, address: Address, bytes: &[u8]) -> Vec<Address> {
        let instructions = self.decode_all(address, bytes, 1000);
        let mut targets = Vec::new();

        for insn in &instructions {
            if insn.is_call {
                if let Some(target) = self.get_branch_target(insn) {
                    targets.push(target);
                }
            }
        }

        targets
    }

    pub fn find_function_bounds(&self, address: Address, bytes: &[u8]) -> Option<(Address, Address)> {
        let instructions = self.decode_function(address, bytes);
        
        if instructions.is_empty() {
            return None;
        }

        let start = instructions[0].address;
        let end = instructions.last()?.address + instructions.last()?.length as u64;

        Some((start, end))
    }
}

impl Default for Disassembler {
    fn default() -> Self {
        Self::new().expect("Failed to create disassembler")
    }
}

pub struct DisassemblerBuilder {
    machine_mode: Option<MachineMode>,
}

impl DisassemblerBuilder {
    pub fn new() -> Self {
        Self {
            machine_mode: None,
        }
    }

    pub fn machine_mode(mut self, mode: MachineMode) -> Self {
        self.machine_mode = Some(mode);
        self
    }

    pub fn build(self) -> Result<Disassembler, DisassemblerError> {
        let machine_mode = self.machine_mode.unwrap_or(MachineMode::Long64);
        let decoder = Decoder::new_with_mode(machine_mode);
        let formatter = Formatter::intel();

        Ok(Disassembler {
            decoder,
            formatter,
        })
    }
}

impl Default for DisassemblerBuilder {
    fn default() -> Self {
        Self::new()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_decode_simple() {
        let disasm = Disassembler::new().unwrap();
        let code: &[u8] = &[0x48, 0x89, 0x5C, 0x24, 0x08];
        
        let insn = disasm.decode(0x1000, code).unwrap();
        
        assert_eq!(insn.address, 0x1000);
        assert_eq!(insn.mnemonic, "mov");
        assert_eq!(insn.length, 5);
    }

    #[test]
    fn test_decode_all() {
        let disasm = Disassembler::new().unwrap();
        let code: &[u8] = &[
            0x48, 0x89, 0x5C, 0x24, 0x08,
            0x48, 0x89, 0x74, 0x24, 0x10,
            0x57,
        ];
        
        let instructions = disasm.decode_all(0x1000, code, 10);
        
        assert_eq!(instructions.len(), 3);
        assert_eq!(instructions[0].mnemonic, "mov");
        assert_eq!(instructions[1].mnemonic, "mov");
        assert_eq!(instructions[2].mnemonic, "push");
    }

    #[test]
    fn test_is_branch() {
        let disasm = Disassembler::new().unwrap();
        let jmp_code: &[u8] = &[0xEB, 0x05];
        
        let insn = disasm.decode(0x1000, jmp_code).unwrap();
        
        assert!(insn.is_branch);
    }

    #[test]
    fn test_is_call() {
        let disasm = Disassembler::new().unwrap();
        let call_code: &[u8] = &[0xE8, 0x00, 0x00, 0x00, 0x00];
        
        let insn = disasm.decode(0x1000, call_code).unwrap();
        
        assert!(insn.is_call);
    }

    #[test]
    fn test_is_ret() {
        let disasm = Disassembler::new().unwrap();
        let ret_code: &[u8] = &[0xC3];
        
        let insn = disasm.decode(0x1000, ret_code).unwrap();
        
        assert!(insn.is_ret);
    }
}
