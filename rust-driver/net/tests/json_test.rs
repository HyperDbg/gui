#![allow(non_snake_case)]

use serde::{Serialize, Deserialize};

pub fn Marshal<T: Serialize>(v: &T) -> Result<Vec<u8>, String> {
    serde_json::to_vec(v).map_err(|e| e.to_string())
}

pub fn Unmarshal<'de, T: Deserialize<'de>>(data: &'de [u8]) -> Result<T, String> {
    serde_json::from_slice(data).map_err(|e| e.to_string())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[derive(Serialize, Deserialize, Debug, PartialEq)]
    struct TestCommand {
        action: String,
        target: Option<String>,
        size: Option<i32>,
    }

    #[test]
    fn test_marshal() {
        let cmd = TestCommand {
            action: String::from("ping"),
            target: Some(String::from("0x1000")),
            size: Some(64),
        };
        let bytes = Marshal(&cmd).unwrap();
        let json_str = String::from_utf8(bytes).unwrap();
        assert!(json_str.contains("\"action\":\"ping\""));
        assert!(json_str.contains("\"target\":\"0x1000\""));
        assert!(json_str.contains("\"size\":64"));
    }

    #[test]
    fn test_unmarshal() {
        let json = b"{\"action\":\"ping\",\"target\":\"0x1000\",\"size\":64}";
        let cmd: TestCommand = Unmarshal(json).unwrap();
        assert_eq!(cmd.action, "ping");
        assert_eq!(cmd.target, Some(String::from("0x1000")));
        assert_eq!(cmd.size, Some(64));
    }

    #[test]
    fn test_unmarshal_with_null() {
        let json = b"{\"action\":\"ping\",\"target\":null,\"size\":null}";
        let cmd: TestCommand = Unmarshal(json).unwrap();
        assert_eq!(cmd.action, "ping");
        assert_eq!(cmd.target, None);
        assert_eq!(cmd.size, None);
    }

    #[test]
    fn test_roundtrip() {
        let original = TestCommand {
            action: String::from("read"),
            target: Some(String::from("0x2000")),
            size: Some(128),
        };
        let bytes = Marshal(&original).unwrap();
        let decoded: TestCommand = Unmarshal(&bytes).unwrap();
        assert_eq!(decoded, original);
    }

    #[test]
    fn test_unmarshal_error() {
        let json = b"invalid json";
        let result: Result<TestCommand, String> = Unmarshal(json);
        assert!(result.is_err());
    }
}
