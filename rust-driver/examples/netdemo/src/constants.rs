pub static NT_DRIVER_NAME: &str = "\\Driver\\netDemo";
pub static NT_DEVICE_NAME: &str = "\\Device\\netDemo";
pub static DOS_DEVICE_NAME: &str = "\\??\\netDemo";
pub static DRIVER_UM_NAME: &str = "\\\\.\\netDemo";

pub const DEFAULT_SERVER_IP: [u8; 4] = [127, 0, 0, 1];
pub const DEFAULT_SERVER_PORT: u16 = 9527;
pub const MAX_MESSAGE_SIZE: usize = 4096;
