#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum LogLevel {
    Info,
    Success,
    Warning,
    Error,
}

impl LogLevel {
    pub const fn as_str(&self) -> &'static str {
        match self {
            LogLevel::Info => "[INFO]",
            LogLevel::Success => "[SUCCESS]",
            LogLevel::Warning => "[WARNING]",
            LogLevel::Error => "[ERROR]",
        }
    }
}

pub struct Logger;

impl Logger {
    #[allow(clippy::needless_pass_by_value)]
    pub fn log(msg: &str, level: LogLevel) {
        wdk::println!("{} {}", level.as_str(), msg);
    }
}

#[macro_export]
macro_rules! log {
    ($level:expr, $($arg:tt)*) => {
        $crate::logger::Logger::log(&alloc::format!($($arg)*), $level)
    };
}
