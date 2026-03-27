#![no_std]
#![allow(non_snake_case)]

extern crate alloc;

#[macro_export]
macro_rules! log_success {
    ($msg:expr) => {
        wdk::println!("[hyperdbg] ✅ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[hyperdbg] ✅ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_error {
    ($msg:expr) => {
        wdk::println!("[hyperdbg] ❌ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[hyperdbg] ❌ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_info {
    ($msg:expr) => {
        wdk::println!("[hyperdbg] ℹ️ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[hyperdbg] ℹ️ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}

#[macro_export]
macro_rules! log_warn {
    ($msg:expr) => {
        wdk::println!("[hyperdbg] ⚠️ {} [{}:{}]", $msg, file!(), line!());
    };
    ($fmt:expr, $($arg:expr),*) => {
        wdk::println!("[hyperdbg] ⚠️ {} [{}:{}]", alloc::format!($fmt, $($arg),*), file!(), line!());
    };
}
