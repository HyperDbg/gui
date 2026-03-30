pub use wdk_sys::_LARGE_INTEGER;

pub mod memory_caching_type {
    pub const MM_NON_CACHED: i32 = 0;
    pub const MM_CACHED: i32 = 1;
    pub const MM_WRITE_COMBINED: i32 = 2;
    pub const MM_HARDWARE_COHERENT_CACHED: i32 = 3;
}

pub use memory_caching_type::*;

pub type MEMORY_CACHING_TYPE = i32;

pub mod pool_flags {
    pub const POOL_FLAG_NON_PAGED: u64 = 0x0000000000000040;
    pub const POOL_FLAG_NON_PAGED_EXECUTE: u64 = 0x0000000000000080;
    pub const POOL_FLAG_PAGED: u64 = 0x0000000000000100;
    pub const POOL_FLAG_USE_QUOTA: u64 = 0x0000000000000001;
    pub const POOL_FLAG_UNINITIALIZED: u64 = 0x0000000000000002;
    pub const POOL_FLAG_CACHE_ALIGNED: u64 = 0x0000000000000008;
    pub const POOL_FLAG_RAISE_ON_FAILURE: u64 = 0x0000000000000020;
}

pub use pool_flags::*;
