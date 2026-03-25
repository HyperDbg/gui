#![no_std]

extern crate alloc;

mod socket;
mod provider;
mod buffer;
mod message;

pub use socket::*;
pub use buffer::*;
pub use message::*;

pub const WSK_NO_WAIT: i64 = 0;
pub const WSK_INFINITE_WAIT: i64 = -1;

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
