use alloc::vec::Vec;
use alloc::vec;

pub struct Buffer {
    data: Vec<u8>,
    read_pos: usize,
    write_pos: usize,
}

impl Buffer {
    pub fn new(capacity: usize) -> Self {
        Self {
            data: vec![0u8; capacity],
            read_pos: 0,
            write_pos: 0,
        }
    }

    pub fn from_slice(data: &[u8]) -> Self {
        Self {
            data: data.to_vec(),
            read_pos: 0,
            write_pos: data.len(),
        }
    }

    pub fn len(&self) -> usize {
        self.write_pos - self.read_pos
    }

    pub fn is_empty(&self) -> bool {
        self.len() == 0
    }

    pub fn capacity(&self) -> usize {
        self.data.len()
    }

    pub fn read(&mut self, buf: &mut [u8]) -> usize {
        let available = self.len();
        let to_read = core::cmp::min(available, buf.len());

        buf[..to_read].copy_from_slice(&self.data[self.read_pos..self.read_pos + to_read]);
        self.read_pos += to_read;

        if self.read_pos == self.write_pos {
            self.read_pos = 0;
            self.write_pos = 0;
        }

        to_read
    }

    pub fn write(&mut self, data: &[u8]) -> usize {
        let available = self.capacity() - self.write_pos;
        let to_write = core::cmp::min(available, data.len());

        self.data[self.write_pos..self.write_pos + to_write].copy_from_slice(&data[..to_write]);
        self.write_pos += to_write;

        to_write
    }

    pub fn peek(&self, buf: &mut [u8]) -> usize {
        let available = self.len();
        let to_peek = core::cmp::min(available, buf.len());

        buf[..to_peek].copy_from_slice(&self.data[self.read_pos..self.read_pos + to_peek]);

        to_peek
    }

    pub fn skip(&mut self, n: usize) {
        let to_skip = core::cmp::min(n, self.len());
        self.read_pos += to_skip;
    }

    pub fn clear(&mut self) {
        self.read_pos = 0;
        self.write_pos = 0;
    }

    pub fn as_slice(&self) -> &[u8] {
        &self.data[self.read_pos..self.write_pos]
    }

    pub fn as_mut_slice(&mut self) -> &mut [u8] {
        &mut self.data[self.write_pos..]
    }

    pub fn compact(&mut self) {
        if self.read_pos > 0 {
            let len = self.len();
            self.data.copy_within(self.read_pos..self.write_pos, 0);
            self.read_pos = 0;
            self.write_pos = len;
        }
    }
}

pub struct BufferPool {
    buffers: Vec<Buffer>,
    buffer_size: usize,
    max_buffers: usize,
}

impl BufferPool {
    pub fn new(buffer_size: usize, initial_count: usize, max_buffers: usize) -> Self {
        let mut buffers = Vec::with_capacity(max_buffers);
        for _ in 0..initial_count {
            buffers.push(Buffer::new(buffer_size));
        }

        Self {
            buffers,
            buffer_size,
            max_buffers,
        }
    }

    pub fn acquire(&mut self) -> Option<Buffer> {
        self.buffers.pop().or_else(|| {
            if self.buffers.len() < self.max_buffers {
                Some(Buffer::new(self.buffer_size))
            } else {
                None
            }
        })
    }

    pub fn release(&mut self, buffer: Buffer) {
        if self.buffers.len() < self.max_buffers {
            let mut buffer = buffer;
            buffer.clear();
            self.buffers.push(buffer);
        }
    }

    pub fn available(&self) -> usize {
        self.buffers.len()
    }
}

pub struct RingBuffer {
    data: Vec<u8>,
    head: usize,
    tail: usize,
    full: bool,
}

impl RingBuffer {
    pub fn new(capacity: usize) -> Self {
        Self {
            data: vec![0u8; capacity],
            head: 0,
            tail: 0,
            full: false,
        }
    }

    pub fn len(&self) -> usize {
        if self.full {
            self.data.len()
        } else if self.tail >= self.head {
            self.tail - self.head
        } else {
            self.data.len() - self.head + self.tail
        }
    }

    pub fn is_empty(&self) -> bool {
        !self.full && self.head == self.tail
    }

    pub fn is_full(&self) -> bool {
        self.full
    }

    pub fn capacity(&self) -> usize {
        self.data.len()
    }

    pub fn write(&mut self, data: &[u8]) -> usize {
        let mut written = 0;

        for &byte in data {
            if self.full {
                break;
            }

            self.data[self.tail] = byte;
            self.tail = (self.tail + 1) % self.data.len();
            written += 1;

            if self.tail == self.head {
                self.full = true;
            }
        }

        written
    }

    pub fn read(&mut self, buf: &mut [u8]) -> usize {
        let mut read = 0;

        for byte in buf.iter_mut() {
            if self.is_empty() {
                break;
            }

            *byte = self.data[self.head];
            self.head = (self.head + 1) % self.data.len();
            read += 1;
            self.full = false;
        }

        read
    }

    pub fn clear(&mut self) {
        self.head = 0;
        self.tail = 0;
        self.full = false;
    }
}
