//! Rust code generator for hook operations
//!
//! Converts analyzed hook operations into executable Rust code.

use alloc::string::String;
use alloc::vec::Vec;
use alloc::format;
use crate::go_script::analyzer::HookOperation;

pub struct RustGenerator {
    code: String,
}

impl RustGenerator {
    pub fn new() -> Self {
        Self {
            code: String::new(),
        }
    }

    pub fn generate(&mut self, operations: &[HookOperation]) -> String {
        self.code.clear();

        for op in operations {
            self.generate_operation(op);
        }

        self.code.clone()
    }

    fn generate_operation(&mut self, op: &HookOperation) {
        match op {
            HookOperation::GetTypeAssertion { target, type_name } => {
                self.code.push_str(&format!(
                    "let {} = unsafe {{ &*({} as *const {}) }};\n",
                    target, target, type_name
                ));
            }
            HookOperation::GetBufferPointer { buffer_field } => {
                self.code.push_str(&format!(
                    "let buffer_ptr = {};\n",
                    buffer_field
                ));
            }
            HookOperation::WriteMemory { buffer, offset, data } => {
                self.code.push_str(&format!(
                    "unsafe {{\n    core::ptr::copy_nonoverlapping(\n        &{:?} as *const u8,\n        ({} as *mut u8).add({}),\n        {}\n    );\n}}\n",
                    data, buffer, offset, data.len()
                ));
            }
            HookOperation::WriteMemorySwapped { buffer, offset, data } => {
                let swapped = swap_bytes(data);
                self.code.push_str(&format!(
                    "unsafe {{\n    core::ptr::copy_nonoverlapping(\n        &{:?} as *const u8,\n        ({} as *mut u8).add({}),\n        {}\n    );\n}}\n",
                    swapped, buffer, offset, swapped.len()
                ));
            }
            HookOperation::SetReturnValue { value } => {
                self.code.push_str(&format!(
                    "regs.rax = {} as u64;\n",
                    value
                ));
            }
            HookOperation::CheckCondition { left, operator, right } => {
                self.code.push_str(&format!(
                    "if {} {} {} {{\n",
                    left, operator, right
                ));
            }
            HookOperation::CheckIoControlCode { codes } => {
                let codes_str: Vec<String> = codes.iter()
                    .map(|c| format!("0x{:X}", c))
                    .collect();
                self.code.push_str(&format!(
                    "match args.io_control_code {{\n    {} => {{\n",
                    codes_str.join(" | ")
                ));
            }
        }
    }

    pub fn generate_function(&mut self, operations: &[HookOperation], func_name: &str) -> String {
        let body = self.generate(operations);

        format!(
            "pub unsafe fn {}(regs: &mut crate::hyperkd::hyperhv::assembly::debugger_asm::GuestContext) {{\n    let args = &*(regs.rdx as *const crate::go_script::executor::NtDeviceIoControlFileArgs);\n{}\n}}",
            func_name,
            body
        )
    }
}

impl Default for RustGenerator {
    fn default() -> Self {
        Self::new()
    }
}

fn swap_bytes(data: &[u8]) -> Vec<u8> {
    let mut result = data.to_vec();
    for i in (0..result.len()).step_by(2) {
        if i + 1 < result.len() {
            result.swap(i, i + 1);
        }
    }
    result
}

pub fn generate_rust_code(operations: &[HookOperation]) -> String {
    let mut generator = RustGenerator::new();
    generator.generate(operations)
}

pub fn generate_hook_function(operations: &[HookOperation], func_name: &str) -> String {
    let mut generator = RustGenerator::new();
    generator.generate_function(operations, func_name)
}
