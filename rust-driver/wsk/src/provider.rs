use crate::{SocketError, SocketType, Protocol};

pub struct WskProvider {
}

impl WskProvider {
    pub fn new() -> Result<Self, SocketError> {
        Ok(Self {})
    }
}
