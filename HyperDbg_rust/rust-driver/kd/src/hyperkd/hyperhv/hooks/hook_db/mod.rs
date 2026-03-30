#![allow(non_snake_case)]
#![allow(dead_code)]

mod ept_hook_gen;
mod inline_hook_gen;
mod hook_args_gen;

pub use ept_hook_gen::*;
pub use inline_hook_gen::*;
pub use hook_args_gen::*;
