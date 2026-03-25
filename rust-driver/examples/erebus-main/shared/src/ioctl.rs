extern crate alloc;

const FILE_DEVICE_UNKNOWN: u32 = 34u32;
// const METHOD_NEITHER: u32 = 3u32;
const METHOD_BUFFERED: u32 = 0u32;
const FILE_ANY_ACCESS: u32 = 0u32;

macro_rules! ctl_code {
    ($DeviceType:expr, $Function:expr, $Method:expr, $Access:expr) => {
        ($DeviceType << 16) | ($Access << 14) | ($Function << 2) | $Method
    };
}

/* IOCTL CODES */

// read from process memory
pub const EREBUS_IOCTL_READ: u32 =
    ctl_code!(FILE_DEVICE_UNKNOWN, 0x1, METHOD_BUFFERED, FILE_ANY_ACCESS);

// write to process memory
pub const EREBUS_IOCTL_WRITE: u32 =
    ctl_code!(FILE_DEVICE_UNKNOWN, 0x2, METHOD_BUFFERED, FILE_ANY_ACCESS);

// sysdemo compatible codes
pub const IOCTL_SEND_DATA: u32 =
    ctl_code!(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS);

pub const IOCTL_RECEIVE_DATA: u32 =
    ctl_code!(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS);
