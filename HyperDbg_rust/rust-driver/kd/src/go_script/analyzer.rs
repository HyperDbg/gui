//! AST analyzer for extracting hook operations from Go code
//!
//! Analyzes parsed Go code and extracts operations that can be executed
//! in the kernel driver context.

use alloc::string::String;
use alloc::vec::Vec;
use alloc::borrow::ToOwned;
use alloc::format;
use crate::go_script::parser::GoNode;

#[derive(Debug, Clone)]
pub enum HookOperation {
    GetTypeAssertion {
        target: String,
        type_name: String,
    },
    GetBufferPointer {
        buffer_field: String,
    },
    WriteMemory {
        buffer: String,
        offset: usize,
        data: Vec<u8>,
    },
    WriteMemorySwapped {
        buffer: String,
        offset: usize,
        data: Vec<u8>,
    },
    SetReturnValue {
        value: i32,
    },
    CheckCondition {
        left: String,
        operator: String,
        right: String,
    },
    CheckIoControlCode {
        codes: Vec<u32>,
    },
}

pub struct AstAnalyzer {
    operations: Vec<HookOperation>,
}

impl AstAnalyzer {
    pub fn new() -> Self {
        Self {
            operations: Vec::new(),
        }
    }

    pub fn analyze(&mut self, nodes: &[GoNode]) -> Vec<HookOperation> {
        self.operations.clear();

        for node in nodes {
            self.analyze_node(node);
        }

        self.operations.clone()
    }

    fn analyze_node(&mut self, node: &GoNode) {
        match node {
            GoNode::TypeAssertion { target, type_name } => {
                self.operations.push(HookOperation::GetTypeAssertion {
                    target: target.clone(),
                    type_name: type_name.clone(),
                });
            }
            GoNode::Assignment { target, value } => {
                self.analyze_assignment(target, value);
            }
            GoNode::FunctionCall { name, arguments } => {
                self.analyze_function_call(name, arguments);
            }
            GoNode::IfStatement { condition, body } => {
                self.analyze_if_statement(condition, body);
            }
            GoNode::SwitchStatement { value, cases } => {
                self.analyze_switch_statement(value, cases);
            }
            _ => {}
        }
    }

    fn analyze_assignment(&mut self, target: &str, value: &str) {
        if target == "ctx.Return" {
            if let Ok(return_value) = value.parse::<i32>() {
                self.operations.push(HookOperation::SetReturnValue {
                    value: return_value,
                });
            }
        }

        if value.contains("unsafe.Pointer") && value.contains("OutputBuffer") {
            self.operations.push(HookOperation::GetBufferPointer {
                buffer_field: value.to_owned(),
            });
        }
    }

    fn analyze_function_call(&mut self, name: &str, arguments: &[String]) {
        match name {
            "swapBytes" => {
                if arguments.len() == 2 {
                    let target = &arguments[0];
                    let data = &arguments[1];

                    if let Some((buffer, offset)) = self.parse_buffer_offset(target) {
                        if let Some(data_bytes) = self.parse_byte_slice(data) {
                            self.operations.push(HookOperation::WriteMemorySwapped {
                                buffer,
                                offset,
                                data: data_bytes,
                            });
                        }
                    }
                }
            }
            "WriteOutputBytes" | "WriteOutputString" => {
                if arguments.len() == 1 {
                    let data = &arguments[0];
                    if let Some(data_bytes) = self.parse_byte_slice(data) {
                        self.operations.push(HookOperation::WriteMemory {
                            buffer: "args.OutputBuffer".to_owned(),
                            offset: 0,
                            data: data_bytes,
                        });
                    }
                }
            }
            _ => {}
        }
    }

    fn analyze_if_statement(&mut self, condition: &str, body: &[GoNode]) {
        if condition.contains("IoControlCode") {
            let codes = self.extract_io_control_codes(condition);
            if !codes.is_empty() {
                self.operations.push(HookOperation::CheckIoControlCode { codes });
            }
        }

        for node in body {
            self.analyze_node(node);
        }
    }

    fn analyze_switch_statement(&mut self, value: &str, cases: &[super::parser::SwitchCase]) {
        if value.contains("IoControlCode") {
            for case in cases {
                let codes: Vec<u32> = case.values.iter()
                    .filter_map(|v| self.parse_hex_or_const(v))
                    .collect();

                if !codes.is_empty() {
                    self.operations.push(HookOperation::CheckIoControlCode { codes });
                }

                for node in &case.body {
                    self.analyze_node(node);
                }
            }
        }
    }

    fn parse_buffer_offset(&self, target: &str) -> Option<(String, usize)> {
        if target.contains("SSerialNumber") {
            Some(("args.OutputBuffer".to_owned(), 20))
        } else if target.contains("SModelNumber") {
            Some(("args.OutputBuffer".to_owned(), 56))
        } else {
            None
        }
    }

    fn parse_byte_slice(&self, data: &str) -> Option<Vec<u8>> {
        if data.starts_with("[]byte{") {
            let bytes_str = data.trim_start_matches("[]byte{").trim_end_matches("}");
            let bytes: Vec<u8> = bytes_str
                .split(',')
                .filter_map(|b| {
                    b.trim()
                        .trim_start_matches("0x")
                        .parse::<u8>()
                        .ok()
                })
                .collect();
            Some(bytes)
        } else if data.starts_with("\"") {
            let s = data.trim_matches('"');
            Some(s.as_bytes().to_vec())
        } else {
            None
        }
    }

    fn extract_io_control_codes(&self, condition: &str) -> Vec<u32> {
        let mut codes = Vec::new();

        for part in condition.split("||") {
            if let Some(code) = self.parse_hex_or_const(part.trim()) {
                codes.push(code);
            }
        }

        codes
    }

    fn parse_hex_or_const(&self, s: &str) -> Option<u32> {
        let s = s.trim();

        if s.starts_with("0x") || s.starts_with("0X") {
            u32::from_str_radix(s.trim_start_matches("0x").trim_start_matches("0X"), 16).ok()
        } else if s.chars().all(|c| c.is_ascii_digit()) {
            s.parse::<u32>().ok()
        } else {
            match s {
                "SMART_RCV_DRIVE_DATA" => Some(0x0007c088),
                "IOCTL_STORAGE_QUERY_PROPERTY" => Some(0x002d1400),
                "IOCTL_NDIS_QUERY_GLOBAL_STATS" => Some(0x00170202),
                _ => None,
            }
        }
    }
}

impl Default for AstAnalyzer {
    fn default() -> Self {
        Self::new()
    }
}

pub fn analyze_go_code(nodes: &[GoNode]) -> Vec<HookOperation> {
    let mut analyzer = AstAnalyzer::new();
    analyzer.analyze(nodes)
}
