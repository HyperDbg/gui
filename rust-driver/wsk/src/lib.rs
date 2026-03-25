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
pub struct SocketAddrV4 {
    pub ip: [u8; 4],
    pub port: u16,
}

impl SocketAddrV4 {
    pub const fn new(ip: [u8; 4], port: u16) -> Self {
        Self { ip, port }
    }

    pub const fn localhost(port: u16) -> Self {
        Self::new([127, 0, 0, 1], port)
    }

    pub fn to_sockaddr_in(&self) -> [u8; 16] {
        let mut addr = [0u8; 16];
        addr[0..2].copy_from_slice(&(2u16 as u16).to_be_bytes());
        addr[2..4].copy_from_slice(&self.port.to_be_bytes());
        addr[4..8].copy_from_slice(&self.ip);
        addr
    }
}

#[derive(Debug, Clone, Copy)]
pub struct SocketAddrV6 {
    pub ip: [u8; 16],
    pub port: u16,
    pub flowinfo: u32,
    pub scope_id: u32,
}

impl SocketAddrV6 {
    pub const fn new(ip: [u8; 16], port: u16, flowinfo: u32, scope_id: u32) -> Self {
        Self { ip, port, flowinfo, scope_id }
    }

    pub const fn any(port: u16) -> Self {
        Self::new([0u8; 16], port, 0, 0)
    }

    pub const fn localhost(port: u16) -> Self {
        let mut ip = [0u8; 16];
        ip[15] = 1;
        Self::new(ip, port, 0, 0)
    }

    pub fn to_sockaddr_in6(&self) -> [u8; 28] {
        let mut addr = [0u8; 28];
        addr[0..2].copy_from_slice(&(23u16 as u16).to_be_bytes());
        addr[2..4].copy_from_slice(&self.port.to_be_bytes());
        addr[4..8].copy_from_slice(&self.flowinfo.to_be_bytes());
        addr[8..24].copy_from_slice(&self.ip);
        addr[24..28].copy_from_slice(&self.scope_id.to_be_bytes());
        addr
    }
}

pub type SocketAddr = SocketAddrV4;
