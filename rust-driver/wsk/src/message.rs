use alloc::string::{String, ToString};
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Command {
    pub action: String,
    #[serde(default)]
    pub target: Option<String>,
    #[serde(default)]
    pub value: Option<u64>,
    #[serde(default)]
    pub size: Option<usize>,
    #[serde(default)]
    pub data: Option<String>,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Response {
    pub status: String,
    #[serde(default)]
    pub data: Option<String>,
    #[serde(default)]
    pub error: Option<String>,
}

impl Response {
    #[inline(never)]
    pub fn ok(data: &str) -> Self {
        Self {
            status: "ok".to_string(),
            data: Some(data.to_string()),
            error: None,
        }
    }

    #[inline(never)]
    pub fn error(msg: &str) -> Self {
        Self {
            status: "error".to_string(),
            data: None,
            error: Some(msg.to_string()),
        }
    }
}

#[inline(never)]
pub fn parse_command(data: &[u8]) -> Option<Command> {
    serde_json::from_slice(data).ok()
}

#[inline(never)]
pub fn serialize_response(resp: &Response) -> alloc::vec::Vec<u8> {
    serde_json::to_vec(resp).unwrap_or_else(|_| alloc::vec::Vec::new())
}
