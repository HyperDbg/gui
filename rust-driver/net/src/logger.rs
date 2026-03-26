#![allow(non_snake_case)]

extern crate alloc;

use wdk;

#[macro_export]
macro_rules! log_success {
    ($msg:expr) => {
        wdk::println!("[net] ✅ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[net] ✅ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_error {
    ($msg:expr) => {
        wdk::println!("[net] ❌ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[net] ❌ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_info {
    ($msg:expr) => {
        wdk::println!("[net] ℹ️ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[net] ℹ️ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_warn {
    ($msg:expr) => {
        wdk::println!("[net] ⚠️ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[net] ⚠️ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

pub use log_success;
pub use log_error;
pub use log_info;
pub use log_warn;