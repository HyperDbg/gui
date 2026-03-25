use core::ffi::c_void;

#[derive(Debug, Clone, Copy)]
// use C ABI due to Rust's ABI instability
#[repr(C)]
pub struct Request {
    pub process_id: u32,

    pub address: *mut c_void,
    pub buffer: *mut c_void,

    pub size: u64,
}
