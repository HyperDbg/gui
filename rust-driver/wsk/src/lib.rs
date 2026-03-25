#![no_std]

extern crate alloc;

mod socket;
mod provider;
mod buffer;

pub use socket::*;
pub use buffer::*;

pub const WSK_NO_WAIT: i64 = 0;
pub const WSK_INFINITE_WAIT: i64 = -1;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SocketError {
    NotConnected,
    ConnectionRefused,
    Timeout,
    BufferTooSmall,
    InvalidAddress,
    OutOfMemory,
    SystemError(i32),
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum SocketType {
    Stream = 1,
    Dgram = 2,
    Raw = 3,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum Protocol {
    Tcp = 6,
    Udp = 17,
}

#[derive(Debug, Clone, Copy)]
pub struct SocketAddr {
    pub ip: [u8; 4],
    pub port: u16,
}

impl SocketAddr {
    pub const fn new(ip: [u8; 4], port: u16) -> Self {
        Self { ip, port }
    }

    pub const fn localhost(port: u16) -> Self {
        Self::new([127, 0, 0, 1], port)
    }

    pub fn to_sockaddr_in(&self) -> [u8; 16] {
        let mut addr = [0u8; 16];
        addr[0..2].copy_from_slice(&2u16.to_be_bytes());
        addr[2..4].copy_from_slice(&self.port.to_be_bytes());
        addr[4..8].copy_from_slice(&self.ip);
        addr
    }
}

impl core::str::FromStr for SocketAddr {
    type Err = SocketError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts: alloc::vec::Vec<&str> = s.split(':').collect();
        if parts.len() != 2 {
            return Err(SocketError::InvalidAddress);
        }

        let ip_parts: alloc::vec::Vec<&str> = parts[0].split('.').collect();
        if ip_parts.len() != 4 {
            return Err(SocketError::InvalidAddress);
        }

        let mut ip = [0u8; 4];
        for i in 0..4 {
            ip[i] = ip_parts[i].parse().map_err(|_| SocketError::InvalidAddress)?;
        }

        let port: u16 = parts[1].parse().map_err(|_| SocketError::InvalidAddress)?;

        Ok(Self { ip, port })
    }
}
