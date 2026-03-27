#![allow(non_snake_case)]

extern crate alloc;

pub mod types_gen;
pub mod handlers_gen;

pub use types_gen::*;
pub use handlers_gen::{DebuggerApi, dispatch_api, NoOpDebugger, EventQueue};
