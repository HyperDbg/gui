pub use wdk_sys::_LARGE_INTEGER;

pub mod memory_caching_type {
    pub const MM_NON_CACHED: i32 = 0;
    pub const MM_CACHED: i32 = 1;
    pub const MM_WRITE_COMBINED: i32 = 2;
    pub const MM_HARDWARE_COHERENT_CACHED: i32 = 3;
}

pub use memory_caching_type::*;

pub type MEMORY_CACHING_TYPE = i32;
