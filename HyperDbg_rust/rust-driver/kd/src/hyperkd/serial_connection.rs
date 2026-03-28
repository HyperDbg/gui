#![allow(dead_code)]

use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

pub const SERIAL_PORT_COM1: u16 = 0x3F8;
pub const SERIAL_PORT_COM2: u16 = 0x2F8;
pub const SERIAL_PORT_COM3: u16 = 0x3E8;
pub const SERIAL_PORT_COM4: u16 = 0x2E8;

pub const SERIAL_DATA_REG: u16 = 0;
pub const SERIAL_INT_ENABLE_REG: u16 = 1;
pub const SERIAL_FIFO_CTRL_REG: u16 = 2;
pub const SERIAL_LINE_CTRL_REG: u16 = 3;
pub const SERIAL_MODEM_CTRL_REG: u16 = 4;
pub const SERIAL_LINE_STATUS_REG: u16 = 5;
pub const SERIAL_MODEM_STATUS_REG: u16 = 6;
pub const SERIAL_SCRATCH_REG: u16 = 7;

pub const SERIAL_LSR_DATA_READY: u8 = 0x01;
pub const SERIAL_LSR_OVERRUN_ERROR: u8 = 0x02;
pub const SERIAL_LSR_PARITY_ERROR: u8 = 0x04;
pub const SERIAL_LSR_FRAMING_ERROR: u8 = 0x08;
pub const SERIAL_LSR_BREAK_INDICATOR: u8 = 0x10;
pub const SERIAL_LSR_TX_EMPTY: u8 = 0x20;
pub const SERIAL_LSR_TX_IDLE: u8 = 0x40;
pub const SERIAL_LSR_FIFO_ERROR: u8 = 0x80;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum BaudRate {
    Baud110 = 1047,
    Baud300 = 384,
    Baud1200 = 96,
    Baud2400 = 48,
    Baud4800 = 24,
    Baud9600 = 12,
    Baud19200 = 6,
    Baud38400 = 3,
    Baud57600 = 2,
    Baud115200 = 1,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Parity {
    None = 0,
    Odd = 1,
    Even = 3,
    Mark = 5,
    Space = 7,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
#[repr(u8)]
pub enum StopBits {
    One = 0,
    OnePointFive = 1,
    Two = 2,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum DataBits {
    Five = 5,
    Six = 6,
    Seven = 7,
    Eight = 8,
}

#[derive(Debug, Clone, Copy)]
pub struct SerialConfig {
    pub port: u16,
    pub baud_rate: BaudRate,
    pub parity: Parity,
    pub stop_bits: StopBits,
    pub data_bits: DataBits,
}

impl Default for SerialConfig {
    fn default() -> Self {
        Self {
            port: SERIAL_PORT_COM1,
            baud_rate: BaudRate::Baud115200,
            parity: Parity::None,
            stop_bits: StopBits::One,
            data_bits: DataBits::Eight,
        }
    }
}

pub struct SerialConnection {
    port: AtomicU32,
    is_initialized: AtomicBool,
    receive_buffer: Mutex<[u8; 4096]>,
    receive_index: AtomicU32,
    send_index: AtomicU32,
}

impl SerialConnection {
    pub const fn new() -> Self {
        Self {
            port: AtomicU32::new(SERIAL_PORT_COM1 as u32),
            is_initialized: AtomicBool::new(false),
            receive_buffer: Mutex::new([0; 4096]),
            receive_index: AtomicU32::new(0),
            send_index: AtomicU32::new(0),
        }
    }

    fn get_port(&self) -> u16 {
        self.port.load(Ordering::SeqCst) as u16
    }

    pub unsafe fn initialize(&self, config: SerialConfig) -> bool {
        self.port.store(config.port as u32, Ordering::SeqCst);

        outb(config.port + SERIAL_INT_ENABLE_REG, 0x00);

        outb(config.port + SERIAL_LINE_CTRL_REG, 0x80);

        let divisor = config.baud_rate as u16;
        outb(config.port + SERIAL_DATA_REG, (divisor & 0xFF) as u8);
        outb(config.port + SERIAL_INT_ENABLE_REG, ((divisor >> 8) & 0xFF) as u8);

        let mut line_ctrl = 0u8;
        line_ctrl |= (config.data_bits as u8 - 5) & 0x03;
        line_ctrl |= if config.stop_bits == StopBits::Two { 0x04 } else { 0 };
        line_ctrl |= (config.parity as u8) << 3;
        outb(config.port + SERIAL_LINE_CTRL_REG, line_ctrl);

        outb(config.port + SERIAL_FIFO_CTRL_REG, 0xC7);

        outb(config.port + SERIAL_MODEM_CTRL_REG, 0x0B);

        outb(config.port + SERIAL_INT_ENABLE_REG, 0x01);

        self.is_initialized.store(true, Ordering::SeqCst);
        true
    }

    pub unsafe fn send_byte(&self, byte: u8) -> bool {
        if !self.is_initialized.load(Ordering::SeqCst) {
            return false;
        }

        let port = self.get_port();

        while (inb(port + SERIAL_LINE_STATUS_REG) & SERIAL_LSR_TX_IDLE) == 0 {
            core::hint::spin_loop();
        }

        outb(port + SERIAL_DATA_REG, byte);
        true
    }

    pub unsafe fn receive_byte(&self) -> Option<u8> {
        if !self.is_initialized.load(Ordering::SeqCst) {
            return None;
        }

        let port = self.get_port();

        if (inb(port + SERIAL_LINE_STATUS_REG) & SERIAL_LSR_DATA_READY) != 0 {
            Some(inb(port + SERIAL_DATA_REG))
        } else {
            None
        }
    }

    pub unsafe fn send_buffer(&self, buffer: &[u8]) -> usize {
        let mut sent = 0;
        for &byte in buffer {
            if self.send_byte(byte) {
                sent += 1;
            } else {
                break;
            }
        }
        sent
    }

    pub unsafe fn receive_buffer(&self, buffer: &mut [u8]) -> usize {
        let mut received = 0;
        for byte in buffer.iter_mut() {
            if let Some(b) = self.receive_byte() {
                *byte = b;
                received += 1;
            } else {
                break;
            }
        }
        received
    }

    pub unsafe fn has_data(&self) -> bool {
        if !self.is_initialized.load(Ordering::SeqCst) {
            return false;
        }

        let port = self.get_port();

        (inb(port + SERIAL_LINE_STATUS_REG) & SERIAL_LSR_DATA_READY) != 0
    }

    pub unsafe fn flush(&self) {
        if !self.is_initialized.load(Ordering::SeqCst) {
            return;
        }

        let port = self.get_port();

        while (inb(port + SERIAL_LINE_STATUS_REG) & SERIAL_LSR_TX_EMPTY) == 0 {
            core::hint::spin_loop();
        }
    }
}

impl Default for SerialConnection {
    fn default() -> Self {
        Self::new()
    }
}

pub static SERIAL_CONNECTION: SerialConnection = SerialConnection::new();

pub unsafe fn serial_initialize(config: SerialConfig) -> bool {
    SERIAL_CONNECTION.initialize(config)
}

pub unsafe fn serial_send_byte(byte: u8) -> bool {
    SERIAL_CONNECTION.send_byte(byte)
}

pub unsafe fn serial_receive_byte() -> Option<u8> {
    SERIAL_CONNECTION.receive_byte()
}

pub unsafe fn serial_send_buffer(buffer: &[u8]) -> usize {
    SERIAL_CONNECTION.send_buffer(buffer)
}

pub unsafe fn serial_receive_buffer(buffer: &mut [u8]) -> usize {
    SERIAL_CONNECTION.receive_buffer(buffer)
}

pub unsafe fn serial_has_data() -> bool {
    SERIAL_CONNECTION.has_data()
}

pub unsafe fn serial_flush() {
    SERIAL_CONNECTION.flush()
}

#[inline(always)]
unsafe fn outb(port: u16, value: u8) {
    core::arch::asm!(
        "out dx, al",
        in("dx") port,
        in("al") value,
        options(nostack)
    );
}

#[inline(always)]
unsafe fn inb(port: u16) -> u8 {
    let value: u8;
    core::arch::asm!(
        "in al, dx",
        out("al") value,
        in("dx") port,
        options(nostack)
    );
    value
}

pub unsafe fn serial_print(s: &str) {
    for byte in s.bytes() {
        if byte == b'\n' {
            serial_send_byte(b'\r');
        }
        serial_send_byte(byte);
    }
}

pub unsafe fn serial_println(s: &str) {
    serial_print(s);
    serial_send_byte(b'\r');
    serial_send_byte(b'\n');
}

pub unsafe fn serial_print_hex(value: u64) {
    const HEX_CHARS: &[u8] = b"0123456789ABCDEF";

    serial_print("0x");

    for i in (0..16).rev() {
        let nibble = ((value >> (i * 4)) & 0xF) as usize;
        serial_send_byte(HEX_CHARS[nibble]);
    }
}

pub unsafe fn serial_print_decimal(value: u64) {
    if value == 0 {
        serial_send_byte(b'0');
        return;
    }

    let mut buffer = [0u8; 20];
    let mut index = 0;
    let mut num = value;

    while num > 0 {
        buffer[index] = b'0' + (num % 10) as u8;
        num /= 10;
        index += 1;
    }

    for i in (0..index).rev() {
        serial_send_byte(buffer[i]);
    }
}

pub struct SerialLogger;

impl SerialLogger {
    pub unsafe fn log(level: &str, message: &str) {
        serial_print("[");
        serial_print(level);
        serial_print("] ");
        serial_println(message);
    }

    pub unsafe fn info(message: &str) {
        Self::log("INFO", message);
    }

    pub unsafe fn warn(message: &str) {
        Self::log("WARN", message);
    }

    pub unsafe fn error(message: &str) {
        Self::log("ERROR", message);
    }

    pub unsafe fn debug(message: &str) {
        Self::log("DEBUG", message);
    }
}
