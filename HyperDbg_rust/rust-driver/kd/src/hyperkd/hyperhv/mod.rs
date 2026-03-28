pub mod common;
pub mod vmm;
pub mod hooks;
pub mod memory;
pub mod state;
pub mod components;

pub mod assembly;
pub mod broadcast;
pub mod devices;
pub mod disassembler;
pub mod features;
pub mod globals;
pub mod interface;
pub mod mmio;
pub mod processor;

pub use state::*;
pub use components::*;

pub const HYPERDBG_SIGNATURE: u64 = 0x4859504552444247;
