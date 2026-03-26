#![allow(non_snake_case)]

use alloc::string::String;
use alloc::vec::Vec;
use serde::{Serialize, Deserialize};

use crate::models_gen::{RegisterState, ProcessInfo, ThreadInfo, ModuleInfo, BreakpointInfo, BreakpointEvent, ExceptionEvent, DebugPrintEvent, SyscallEvent, Response};

pub use crate::models_gen::{EmptyResponse, MessageType, BreakpointType, DebugState};

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

    #[test]
    fn test_register_state() {
        let regs = RegisterState {
            rax: 1,
            rbx: 2,
            rcx: 3,
            rdx: 4,
            ..Default::default()
        };
        let json = serde_json::to_string(&regs).unwrap();
        assert!(json.contains("\"RAX\":1"));
        assert!(json.contains("\"RBX\":2"));
    }

    #[test]
    fn test_process_info() {
        let info = ProcessInfo {
            process_id: 1234,
            image_name: Some(String::from("test.exe")),
            base_address: Some(String::from("0x7ffe00000000")),
            size: 0x10000,
            cr3: 0x1234567000,
        };
        let json = serde_json::to_string(&info).unwrap();
        assert!(json.contains("\"process_id\":1234"));
        assert!(json.contains("\"image_name\":\"test.exe\""));
        assert!(json.contains("\"base_address\":\"0x7ffe00000000\""));
    }

    #[test]
    fn test_thread_info() {
        let info = ThreadInfo {
            thread_id: 1,
            process_id: 1234,
            start_address: Some(String::from("0x7ffe12340000")),
            teb: Some(String::from("0x7ffe00001000")),
            state: 0,
        };
        let json = serde_json::to_string(&info).unwrap();
        assert!(json.contains("\"thread_id\":1"));
        assert!(json.contains("\"start_address\":\"0x7ffe12340000\""));
    }

    #[test]
    fn test_module_info() {
        let info = ModuleInfo {
            base_address: Some(String::from("0x7ffe00000000")),
            size: 0x10000,
            name: Some(String::from("ntdll.dll")),
            path: Some(String::from("C:\\Windows\\System32\\ntdll.dll")),
        };
        let json = serde_json::to_string(&info).unwrap();
        assert!(json.contains("\"base_address\":\"0x7ffe00000000\""));
        assert!(json.contains("\"name\":\"ntdll.dll\""));
    }

    #[test]
    fn test_breakpoint_info() {
        let info = BreakpointInfo {
            id: 1,
            address: Some(String::from("0x7ffe12345678")),
            bp_type: 1,
            process_id: 1234,
            enabled: true,
            hit_count: 5,
        };
        let json = serde_json::to_string(&info).unwrap();
        assert!(json.contains("\"id\":1"));
        assert!(json.contains("\"address\":\"0x7ffe12345678\""));
        assert!(json.contains("\"type\":1"));
    }

    #[test]
    fn test_syscall_event() {
        let event = SyscallEvent {
            process_id: 1234,
            thread_id: 5678,
            core_id: 0,
            syscall_number: 0x12,
            rax: 0x12,
            rcx: 1,
            rdx: 2,
            r8: 3,
            r9: 4,
            r10: 5,
            rsp: 0x7ffe00001000,
        };
        let json = serde_json::to_string(&event).unwrap();
        assert!(json.contains("\"process_id\":1234"));
        assert!(json.contains("\"syscall_number\":18"));
        assert!(json.contains("\"rax\":18"));
    }
}
