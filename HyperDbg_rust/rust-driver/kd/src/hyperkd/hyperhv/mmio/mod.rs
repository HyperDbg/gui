use alloc::sync::Arc;
use alloc::vec::Vec;
use core::sync::atomic::{AtomicBool, AtomicU64, Ordering};
use spin::Mutex;

use crate::hyperkd::hyperhv::memory::mapper::physical_address_to_virtual_address;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum MmioError {
    InvalidAddress,
    InvalidSize,
    MappingFailed,
    UnmappingFailed,
    ReadFailed,
    WriteFailed,
    NotMapped,
    AlreadyMapped,
    OutOfMemory,
}

#[derive(Debug, Clone)]
pub struct MmioRegion {
    pub physical_address: u64,
    pub virtual_address: u64,
    pub size: u64,
    pub is_cached: bool,
    pub is_mapped: bool,
}

impl MmioRegion {
    pub fn new(physical_address: u64, size: u64, is_cached: bool) -> Self {
        Self {
            physical_address,
            virtual_address: 0,
            size,
            is_cached,
            is_mapped: false,
        }
    }
}

pub struct MmioManager {
    regions: Vec<Arc<Mutex<MmioRegion>>>,
    total_mapped_size: AtomicU64,
    initialized: AtomicBool,
}

impl MmioManager {
    pub fn new() -> Self {
        Self {
            regions: Vec::new(),
            total_mapped_size: AtomicU64::new(0),
            initialized: AtomicBool::new(false),
        }
    }

    pub fn initialize(&mut self) -> Result<(), MmioError> {
        if self.initialized.load(Ordering::SeqCst) {
            return Ok(());
        }

        self.initialized.store(true, Ordering::SeqCst);
        log::info!("MMIO manager initialized");
        Ok(())
    }

    pub fn map_region(&mut self, physical_address: u64, size: u64, is_cached: bool) -> Result<Arc<Mutex<MmioRegion>>, MmioError> {
        if size == 0 {
            return Err(MmioError::InvalidSize);
        }

        let virtual_address = unsafe { 
            physical_address_to_virtual_address(physical_address)
        };
        if virtual_address == 0 {
            return Err(MmioError::MappingFailed);
        }

        let mut region = MmioRegion::new(physical_address, size, is_cached);
        region.virtual_address = virtual_address;
        region.is_mapped = true;

        let region = Arc::new(Mutex::new(region));
        self.regions.push(region.clone());
        self.total_mapped_size.fetch_add(size, Ordering::SeqCst);

        log::info!("Mapped MMIO region: PA=0x{:X}, VA=0x{:X}, Size=0x{:X}", 
            physical_address, virtual_address, size);

        Ok(region)
    }

    pub fn unmap_region(&mut self, physical_address: u64) -> Result<(), MmioError> {
        let pos = self.regions.iter()
            .position(|r| r.lock().physical_address == physical_address)
            .ok_or(MmioError::NotMapped)?;

        let region = self.regions.remove(pos);
        let mut region = region.lock();
        
        let size = region.size;
        region.is_mapped = false;
        region.virtual_address = 0;

        self.total_mapped_size.fetch_sub(size, Ordering::SeqCst);

        log::info!("Unmapped MMIO region: PA=0x{:X}", physical_address);
        Ok(())
    }

    pub fn read_u8(&self, physical_address: u64) -> Result<u8, MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            Ok(core::ptr::read_volatile(addr as *const u8))
        }
    }

    pub fn read_u16(&self, physical_address: u64) -> Result<u16, MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 1 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            Ok(core::ptr::read_volatile(addr as *const u16))
        }
    }

    pub fn read_u32(&self, physical_address: u64) -> Result<u32, MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 3 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            Ok(core::ptr::read_volatile(addr as *const u32))
        }
    }

    pub fn read_u64(&self, physical_address: u64) -> Result<u64, MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 7 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            Ok(core::ptr::read_volatile(addr as *const u64))
        }
    }

    pub fn write_u8(&self, physical_address: u64, value: u8) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::write_volatile(addr as *mut u8, value);
        }

        Ok(())
    }

    pub fn write_u16(&self, physical_address: u64, value: u16) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 1 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::write_volatile(addr as *mut u16, value);
        }

        Ok(())
    }

    pub fn write_u32(&self, physical_address: u64, value: u32) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 3 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::write_volatile(addr as *mut u32, value);
        }

        Ok(())
    }

    pub fn write_u64(&self, physical_address: u64, value: u64) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset + 7 >= region.size {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::write_volatile(addr as *mut u64, value);
        }

        Ok(())
    }

    pub fn read_buffer(&self, physical_address: u64, buffer: &mut [u8]) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset as usize + buffer.len() > region.size as usize {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::copy_nonoverlapping(
                addr as *const u8,
                buffer.as_mut_ptr(),
                buffer.len(),
            );
        }

        Ok(())
    }

    pub fn write_buffer(&self, physical_address: u64, buffer: &[u8]) -> Result<(), MmioError> {
        let region = self.find_region(physical_address)?;
        let region = region.lock();
        
        if !region.is_mapped {
            return Err(MmioError::NotMapped);
        }

        let offset = physical_address - region.physical_address;
        if offset as usize + buffer.len() > region.size as usize {
            return Err(MmioError::InvalidAddress);
        }

        let addr = region.virtual_address + offset;
        unsafe {
            core::ptr::copy_nonoverlapping(
                buffer.as_ptr(),
                addr as *mut u8,
                buffer.len(),
            );
        }

        Ok(())
    }

    fn find_region(&self, physical_address: u64) -> Result<Arc<Mutex<MmioRegion>>, MmioError> {
        self.regions.iter()
            .find(|r| {
                let region = r.lock();
                physical_address >= region.physical_address && 
                physical_address < region.physical_address + region.size
            })
            .cloned()
            .ok_or(MmioError::NotMapped)
    }

    pub fn get_mapped_regions(&self) -> &[Arc<Mutex<MmioRegion>>] {
        &self.regions
    }

    pub fn get_total_mapped_size(&self) -> u64 {
        self.total_mapped_size.load(Ordering::SeqCst)
    }
}

impl Default for MmioManager {
    fn default() -> Self {
        Self::new()
    }
}

pub static MMIO_MANAGER: spin::Once<Mutex<MmioManager>> = spin::Once::new();

pub fn get_mmio_manager() -> &'static Mutex<MmioManager> {
    MMIO_MANAGER.call_once(|| Mutex::new(MmioManager::new()))
}

pub fn mmio_read_u8(physical_address: u64) -> Result<u8, MmioError> {
    get_mmio_manager().lock().read_u8(physical_address)
}

pub fn mmio_read_u16(physical_address: u64) -> Result<u16, MmioError> {
    get_mmio_manager().lock().read_u16(physical_address)
}

pub fn mmio_read_u32(physical_address: u64) -> Result<u32, MmioError> {
    get_mmio_manager().lock().read_u32(physical_address)
}

pub fn mmio_read_u64(physical_address: u64) -> Result<u64, MmioError> {
    get_mmio_manager().lock().read_u64(physical_address)
}

pub fn mmio_write_u8(physical_address: u64, value: u8) -> Result<(), MmioError> {
    get_mmio_manager().lock().write_u8(physical_address, value)
}

pub fn mmio_write_u16(physical_address: u64, value: u16) -> Result<(), MmioError> {
    get_mmio_manager().lock().write_u16(physical_address, value)
}

pub fn mmio_write_u32(physical_address: u64, value: u32) -> Result<(), MmioError> {
    get_mmio_manager().lock().write_u32(physical_address, value)
}

pub fn mmio_write_u64(physical_address: u64, value: u64) -> Result<(), MmioError> {
    get_mmio_manager().lock().write_u64(physical_address, value)
}
