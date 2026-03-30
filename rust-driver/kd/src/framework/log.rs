extern crate alloc;


#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LogLevel {
    Info,
    Error,
    Success,
    Warn,
}

#[macro_export]
macro_rules! log {
    (LogLevel::Info, $($arg:tt)*) => {
        $crate::log_info!($($arg)*)
    };
    (LogLevel::Error, $($arg:tt)*) => {
        $crate::log_error!($($arg)*)
    };
    (LogLevel::Success, $($arg:tt)*) => {
        $crate::log_success!($($arg)*)
    };
    (LogLevel::Warn, $($arg:tt)*) => {
        $crate::log_warn!($($arg)*)
    };
}
