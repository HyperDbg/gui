#![allow(non_snake_case)]

use alloc::string::String;
use alloc::vec::Vec;
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Request {
    pub action: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub process_id: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub address: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub breakpoint_id: Option<String>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub size: Option<u32>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub r#type: Option<i32>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<Vec<u8>>,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub regs: Option<RegisterState>,
}

#[derive(Serialize, Deserialize, Debug, Clone, Default)]
pub struct RegisterState {
    #[serde(rename = "RAX")]
    pub rax: u64,
    #[serde(rename = "RBX")]
    pub rbx: u64,
    #[serde(rename = "RCX")]
    pub rcx: u64,
    #[serde(rename = "RDX")]
    pub rdx: u64,
    #[serde(rename = "RSI")]
    pub rsi: u64,
    #[serde(rename = "RDI")]
    pub rdi: u64,
    #[serde(rename = "RBP")]
    pub rbp: u64,
    #[serde(rename = "RSP")]
    pub rsp: u64,
    #[serde(rename = "R8")]
    pub r8: u64,
    #[serde(rename = "R9")]
    pub r9: u64,
    #[serde(rename = "R10")]
    pub r10: u64,
    #[serde(rename = "R11")]
    pub r11: u64,
    #[serde(rename = "R12")]
    pub r12: u64,
    #[serde(rename = "R13")]
    pub r13: u64,
    #[serde(rename = "R14")]
    pub r14: u64,
    #[serde(rename = "R15")]
    pub r15: u64,
    #[serde(rename = "RIP")]
    pub rip: u64,
    #[serde(rename = "RFLAGS")]
    pub rflags: u64,
    #[serde(rename = "CR0")]
    pub cr0: u64,
    #[serde(rename = "CR2")]
    pub cr2: u64,
    #[serde(rename = "CR3")]
    pub cr3: u64,
    #[serde(rename = "CR4")]
    pub cr4: u64,
    #[serde(rename = "DR0")]
    pub dr0: u64,
    #[serde(rename = "DR1")]
    pub dr1: u64,
    #[serde(rename = "DR2")]
    pub dr2: u64,
    #[serde(rename = "DR3")]
    pub dr3: u64,
    #[serde(rename = "DR6")]
    pub dr6: u64,
    #[serde(rename = "DR7")]
    pub dr7: u64,
    #[serde(rename = "GDTR")]
    pub gdtr: u64,
    #[serde(rename = "GSBase")]
    pub gsbase: u64,
    #[serde(rename = "FSBase")]
    pub fsbase: u64,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct Response<T: Serialize> {
    pub success: bool,
    pub message: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<T>,
}

pub type EmptyResponse = Response<()>;

pub fn parse_hex_string(s: &str) -> Option<u64> {
    let s = s.trim();
    if s.starts_with("0x") || s.starts_with("0X") {
        u64::from_str_radix(&s[2..], 16).ok()
    } else {
        u64::from_str_radix(s, 16).ok()
    }
}

pub fn parse_dec_string(s: &str) -> Option<u32> {
    s.trim().parse::<u32>().ok()
}

impl Request {
    pub fn get_process_id(&self) -> Option<u32> {
        self.process_id.as_ref().and_then(|s| parse_dec_string(s))
    }

    pub fn get_address(&self) -> Option<u64> {
        self.address.as_ref().and_then(|s| parse_hex_string(s))
    }

    pub fn get_breakpoint_id(&self) -> Option<u64> {
        self.breakpoint_id.as_ref().and_then(|s| parse_hex_string(s))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_parse_hex_string() {
        assert_eq!(parse_hex_string("0x1234"), Some(0x1234));
        assert_eq!(parse_hex_string("0XABCD"), Some(0xABCD));
        assert_eq!(parse_hex_string("1234"), None);
    }

    #[test]
    fn test_parse_dec_string() {
        assert_eq!(parse_dec_string("1234"), Some(1234));
        assert_eq!(parse_dec_string("5678"), Some(5678));
    }

    #[test]
    fn test_request_initialize() {
        let json = br#"{"action":"initialize"}"#;
        let req: Request = serde_json::from_slice(json).unwrap();
        assert_eq!(req.action, "initialize");
        assert!(req.process_id.is_none());
    }

    #[test]
    fn test_request_attach_process() {
        let json = br#"{"action":"attach_process","process_id":"1234"}"#;
        let req: Request = serde_json::from_slice(json).unwrap();
        assert_eq!(req.action, "attach_process");
        assert_eq!(req.get_process_id(), Some(1234));
    }

    #[test]
    fn test_request_set_breakpoint() {
        let json = br#"{"action":"set_breakpoint","address":"0x7ffe12345678","type":1}"#;
        let req: Request = serde_json::from_slice(json).unwrap();
        assert_eq!(req.action, "set_breakpoint");
        assert_eq!(req.get_address(), Some(0x7ffe12345678));
        assert_eq!(req.r#type, Some(1));
    }

    #[test]
    fn test_request_read_memory() {
        let json = br#"{"action":"read_memory","address":"0x7ffe12345678","size":64}"#;
        let req: Request = serde_json::from_slice(json).unwrap();
        assert_eq!(req.action, "read_memory");
        assert_eq!(req.get_address(), Some(0x7ffe12345678));
        assert_eq!(req.size, Some(64));
    }

    #[test]
    fn test_response_empty() {
        let resp = Response::<()> {
            success: true,
            message: String::from("OK"),
            data: None,
        };
        let json = serde_json::to_string(&resp).unwrap();
        assert!(json.contains("\"success\":true"));
        assert!(json.contains("\"message\":\"OK\""));
    }

    #[test]
    fn test_response_with_data() {
        let resp = Response {
            success: true,
            message: String::from("OK"),
            data: Some(vec![1u8, 2, 3, 4]),
        };
        let json = serde_json::to_string(&resp).unwrap();
        assert!(json.contains("\"data\":[1,2,3,4]"));
    }
}
