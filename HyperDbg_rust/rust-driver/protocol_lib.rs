#![no_std]

extern crate alloc;

mod packet;
mod events;

pub use packet::*;
pub use events::*;

pub const PROTOCOL_VERSION: u32 = 1;
pub const DEFAULT_PORT: u16 = 9527;
pub const MAX_MESSAGE_SIZE: usize = 4096;
