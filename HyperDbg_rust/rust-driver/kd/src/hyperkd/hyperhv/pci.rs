use alloc::boxed::Box;
use alloc::sync::Arc;
use core::sync::atomic::{AtomicBool, AtomicU32, Ordering};
use spin::Mutex;

extern "system" {
    fn ExAllocatePoolWithTag(pool_type: u32, number_of_bytes: usize, tag: u32) -> *mut u8;
    fn ExFreePool(pool: *mut u8);
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PciError {
    NotInitialized,
    AlreadyInitialized,
    InvalidBus,
    InvalidDevice,
    InvalidFunction,
    InvalidOffset,
    DeviceNotFound,
    OperationFailed,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PciCommand {
    IoSpace = 0x0001,
    MemorySpace = 0x0002,
    BusMaster = 0x0004,
    SpecialCycles = 0x0008,
    MemoryWriteAndInvalidate = 0x0010,
    VgaPaletteSnoop = 0x0020,
    ParityErrorResponse = 0x0040,
    SerrEnable = 0x0080,
    FastBackToBack = 0x0100,
    InterruptDisable = 0x0400,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PciStatus {
    CapabilitiesList = 0x0010,
    Capable66Mhz = 0x0020,
    FastBackToBack = 0x0080,
    MasterDataParityError = 0x0100,
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum PciBarType {
    Memory32,
    Memory64,
    Io,
    Unused,
}

#[derive(Debug, Clone, Copy)]
pub struct PciAddress {
    pub bus: u8,
    pub device: u8,
    pub function: u8,
}

impl PciAddress {
    pub fn new(bus: u8, device: u8, function: u8) -> Self {
        Self {
            bus,
            device,
            function,
        }
    }

    pub fn to_config_address(&self, offset: u8) -> u32 {
        let mut address: u32 = 0x80000000;
        address |= (self.bus as u32) << 16;
        address |= (self.device as u32 & 0x1F) << 11;
        address |= (self.function as u32 & 0x07) << 8;
        address |= (offset as u32 & 0xFC);
        address
    }
}

#[derive(Debug, Clone, Copy)]
pub struct PciBar {
    pub index: u8,
    pub bar_type: PciBarType,
    pub base_address: u64,
    pub size: u32,
    pub prefetchable: bool,
}

impl PciBar {
    pub fn new(index: u8) -> Self {
        Self {
            index,
            bar_type: PciBarType::Unused,
            base_address: 0,
            size: 0,
            prefetchable: false,
        }
    }
}

#[derive(Debug, Clone)]
pub struct PciDevice {
    pub address: PciAddress,
    pub vendor_id: u16,
    pub device_id: u16,
    pub class_code: u8,
    pub subclass_code: u8,
    pub prog_if: u8,
    pub revision_id: u8,
    pub bars: [PciBar; 6],
    pub interrupt_line: u8,
    pub interrupt_pin: u8,
    pub command: u16,
    pub status: u16,
}

impl PciDevice {
    pub fn new(address: PciAddress) -> Self {
        Self {
            address,
            vendor_id: 0xFFFF,
            device_id: 0xFFFF,
            class_code: 0,
            subclass_code: 0,
            prog_if: 0,
            revision_id: 0,
            bars: [
                PciBar::new(0),
                PciBar::new(1),
                PciBar::new(2),
                PciBar::new(3),
                PciBar::new(4),
                PciBar::new(5),
            ],
            interrupt_line: 0xFF,
            interrupt_pin: 0,
            command: 0,
            status: 0,
        }
    }

    pub fn is_valid(&self) -> bool {
        self.vendor_id != 0xFFFF && self.vendor_id != 0x0000
    }

    pub fn get_class_name(&self) -> alloc::string::String {
        match self.class_code {
            0x00 => alloc::string::String::from("Unclassified"),
            0x01 => alloc::string::String::from("Mass Storage Controller"),
            0x02 => alloc::string::String::from("Network Controller"),
            0x03 => alloc::string::String::from("Display Controller"),
            0x04 => alloc::string::String::from("Multimedia Controller"),
            0x05 => alloc::string::String::from("Memory Controller"),
            0x06 => alloc::string::String::from("Bridge"),
            0x07 => alloc::string::String::from("Simple Communication Controller"),
            0x08 => alloc::string::String::from("Base System Peripheral"),
            0x09 => alloc::string::String::from("Input Device Controller"),
            0x0A => alloc::string::String::from("Docking Station"),
            0x0B => alloc::string::String::from("Processor"),
            0x0C => alloc::string::String::from("Serial Bus Controller"),
            0x0D => alloc::string::String::from("Wireless Controller"),
            0x0E => alloc::string::String::from("Intelligent Controller"),
            0x0F => alloc::string::String::from("Satellite Communication Controller"),
            0x10 => alloc::string::String::from("Encryption Controller"),
            0x11 => alloc::string::String::from("Signal Processing Controller"),
            0x12 => alloc::string::String::from("Processing Accelerator"),
            0x13 => alloc::string::String::from("Non-Essential Instrumentation"),
            0x40 => alloc::string::String::from("Co-Processor"),
            _ => alloc::string::String::from("Reserved"),
        }
    }
}

pub struct PciManager {
    initialized: AtomicBool,
    devices: alloc::vec::Vec<PciDevice>,
    bus_count: AtomicU32,
    device_count: AtomicU32,
    scan_count: AtomicU32,
}

impl PciManager {
    pub fn new() -> Self {
        Self {
            initialized: AtomicBool::new(false),
            devices: alloc::vec::Vec::new(),
            bus_count: AtomicU32::new(0),
            device_count: AtomicU32::new(0),
            scan_count: AtomicU32::new(0),
        }
    }

    pub fn initialize(&mut self) -> Result<(), PciError> {
        if self.initialized.load(Ordering::Acquire) {
            return Err(PciError::AlreadyInitialized);
        }

        unsafe {
            self.scan_pci_bus();
        }

        self.initialized.store(true, Ordering::Release);
        Ok(())
    }

    pub fn deinitialize(&mut self) {
        if !self.initialized.load(Ordering::Acquire) {
            return;
        }

        self.devices.clear();
        self.bus_count.store(0, Ordering::Release);
        self.device_count.store(0, Ordering::Release);
        self.scan_count.store(0, Ordering::Release);
        self.initialized.store(false, Ordering::Release);
    }

    pub fn is_initialized(&self) -> bool {
        self.initialized.load(Ordering::Acquire)
    }

    pub fn scan_pci_bus(&mut self) {
        self.devices.clear();

        for bus in 0..256 {
            for device in 0..32 {
                let address = PciAddress::new(bus as u8, device as u8, 0);
                let vendor_device_id = unsafe { self.read_config_dword(&address, 0) };

                let vendor_id = (vendor_device_id & 0xFFFF) as u16;
                let device_id = ((vendor_device_id >> 16) & 0xFFFF) as u16;

                if vendor_id == 0xFFFF {
                    continue;
                }

                let mut pci_device = PciDevice::new(address);
                pci_device.vendor_id = vendor_id;
                pci_device.device_id = device_id;

                let class_revision = unsafe { self.read_config_dword(&address, 8) };
                pci_device.class_code = ((class_revision >> 24) & 0xFF) as u8;
                pci_device.subclass_code = ((class_revision >> 16) & 0xFF) as u8;
                pci_device.prog_if = ((class_revision >> 8) & 0xFF) as u8;
                pci_device.revision_id = (class_revision & 0xFF) as u8;

                let command_status = unsafe { self.read_config_dword(&address, 4) };
                pci_device.command = (command_status & 0xFFFF) as u16;
                pci_device.status = ((command_status >> 16) & 0xFFFF) as u16;

                let interrupt = unsafe { self.read_config_dword(&address, 60) };
                pci_device.interrupt_line = (interrupt & 0xFF) as u8;
                pci_device.interrupt_pin = ((interrupt >> 8) & 0xFF) as u8;

                for i in 0..6 {
                    let bar_offset = 0x10 + (i * 4) as u8;
                    let bar_value = unsafe { self.read_config_dword(&address, bar_offset) };

                    if bar_value == 0 {
                        continue;
                    }

                    pci_device.bars[i as usize].index = i as u8;

                    if (bar_value & 1) != 0 {
                        pci_device.bars[i as usize].bar_type = PciBarType::Io;
                        pci_device.bars[i as usize].base_address = (bar_value & 0xFFFFFFFC) as u64;
                    } else {
                        let bar_type = (bar_value >> 1) & 0x03;
                        match bar_type {
                            0x00 => pci_device.bars[i as usize].bar_type = PciBarType::Memory32,
                            0x02 => pci_device.bars[i as usize].bar_type = PciBarType::Memory64,
                            _ => pci_device.bars[i as usize].bar_type = PciBarType::Unused,
                        }

                        pci_device.bars[i as usize].prefetchable = (bar_value & (1 << 3)) != 0;

                        if pci_device.bars[i as usize].bar_type == PciBarType::Memory64 {
                            let bar_high = unsafe { self.read_config_dword(&address, bar_offset + 4) };
                            pci_device.bars[i as usize].base_address = ((bar_high as u64) << 32) | ((bar_value & 0xFFFFFFF0) as u64);
                        } else {
                            pci_device.bars[i as usize].base_address = (bar_value & 0xFFFFFFF0) as u64;
                        }
                    }
                }

                self.devices.push(pci_device);
            }
        }

        self.device_count.store(self.devices.len() as u32, Ordering::Release);
        self.scan_count.fetch_add(1, Ordering::AcqRel);
    }

    pub fn read_config_byte(&self, address: &PciAddress, offset: u8) -> u8 {
        unsafe {
            self.read_config(address, offset) as u8
        }
    }

    pub fn read_config_word(&self, address: &PciAddress, offset: u8) -> u16 {
        unsafe {
            self.read_config(address, offset) as u16
        }
    }

    pub fn read_config_dword(&self, address: &PciAddress, offset: u8) -> u32 {
        unsafe {
            self.read_config(address, offset)
        }
    }

    pub fn write_config_byte(&self, address: &PciAddress, offset: u8, value: u8) {
        unsafe {
            self.write_config(address, offset, value as u32);
        }
    }

    pub fn write_config_word(&self, address: &PciAddress, offset: u8, value: u16) {
        unsafe {
            self.write_config(address, offset, value as u32);
        }
    }

    pub fn write_config_dword(&self, address: &PciAddress, offset: u8, value: u32) {
        unsafe {
            self.write_config(address, offset, value);
        }
    }

    pub fn find_device_by_vendor_device(&self, vendor_id: u16, device_id: u16) -> Option<&PciDevice> {
        self.devices.iter().find(|d| d.vendor_id == vendor_id && d.device_id == device_id)
    }

    pub fn find_devices_by_class(&self, class_code: u8) -> alloc::vec::Vec<&PciDevice> {
        self.devices.iter().filter(|d| d.class_code == class_code).collect()
    }

    pub fn find_devices_by_bus(&self, bus: u8) -> alloc::vec::Vec<&PciDevice> {
        self.devices.iter().filter(|d| d.address.bus == bus).collect()
    }

    pub fn get_device(&self, bus: u8, device: u8, function: u8) -> Option<&PciDevice> {
        self.devices.iter().find(|d| d.address.bus == bus && d.address.device == device && d.address.function == function)
    }

    pub fn get_all_devices(&self) -> alloc::vec::Vec<&PciDevice> {
        self.devices.iter().collect()
    }

    pub fn get_device_count(&self) -> u32 {
        self.device_count.load(Ordering::Acquire)
    }

    pub fn get_scan_count(&self) -> u32 {
        self.scan_count.load(Ordering::Acquire)
    }

    unsafe fn read_config(&self, address: &PciAddress, offset: u8) -> u32 {
        let config_address = address.to_config_address(offset);

        let config_port: u16 = 0xCF8;
        let data_port: u16 = 0xCFC;

        core::arch::asm!(
            "out dx, eax",
            in("dx") config_port,
            in("eax") config_address,
            options(nomem, nostack),
        );

        let mut data: u32;
        core::arch::asm!(
            "in eax, dx",
            out("eax") data,
            in("dx") data_port,
            options(nomem, nostack),
        );

        data
    }

    unsafe fn write_config(&self, address: &PciAddress, offset: u8, value: u32) {
        let config_address = address.to_config_address(offset);

        let config_port: u16 = 0xCF8;
        let data_port: u16 = 0xCFC;

        core::arch::asm!(
            "out dx, eax",
            in("dx") config_port,
            in("eax") config_address,
            options(nomem, nostack),
        );

        core::arch::asm!(
            "out dx, eax",
            in("dx") data_port,
            in("eax") value,
            options(nomem, nostack),
        );
    }
}

pub static PCI_MANAGER: Mutex<PciManager> = Mutex::new(PciManager::new());

pub fn initialize_pci() -> Result<(), PciError> {
    let mut manager = PCI_MANAGER.lock();
    manager.initialize()
}

pub fn deinitialize_pci() {
    let mut manager = PCI_MANAGER.lock();
    manager.deinitialize();
}

pub fn is_pci_initialized() -> bool {
    let manager = PCI_MANAGER.lock();
    manager.is_initialized()
}

pub fn scan_pci_bus() {
    let mut manager = PCI_MANAGER.lock();
    manager.scan_pci_bus()
}

pub fn read_config_byte(bus: u8, device: u8, function: u8, offset: u8) -> u8 {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.read_config_byte(&address, offset)
}

pub fn read_config_word(bus: u8, device: u8, function: u8, offset: u8) -> u16 {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.read_config_word(&address, offset)
}

pub fn read_config_dword(bus: u8, device: u8, function: u8, offset: u8) -> u32 {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.read_config_dword(&address, offset)
}

pub fn write_config_byte(bus: u8, device: u8, function: u8, offset: u8, value: u8) {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.write_config_byte(&address, offset, value)
}

pub fn write_config_word(bus: u8, device: u8, function: u8, offset: u8, value: u16) {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.write_config_word(&address, offset, value)
}

pub fn write_config_dword(bus: u8, device: u8, function: u8, offset: u8, value: u32) {
    let manager = PCI_MANAGER.lock();
    let address = PciAddress::new(bus, device, function);
    manager.write_config_dword(&address, offset, value)
}

pub fn find_device_by_vendor_device(vendor_id: u16, device_id: u16) -> Option<PciDevice> {
    let manager = PCI_MANAGER.lock();
    manager.find_device_by_vendor_device(vendor_id, device_id).copied()
}

pub fn find_devices_by_class(class_code: u8) -> alloc::vec::Vec<PciDevice> {
    let manager = PCI_MANAGER.lock();
    manager.find_devices_by_class(class_code).iter().map(|d| (*d).clone()).collect()
}

pub fn find_devices_by_bus(bus: u8) -> alloc::vec::Vec<PciDevice> {
    let manager = PCI_MANAGER.lock();
    manager.find_devices_by_bus(bus).iter().map(|d| (*d).clone()).collect()
}

pub fn get_device(bus: u8, device: u8, function: u8) -> Option<PciDevice> {
    let manager = PCI_MANAGER.lock();
    manager.get_device(bus, device, function).copied()
}

pub fn get_all_devices() -> alloc::vec::Vec<PciDevice> {
    let manager = PCI_MANAGER.lock();
    manager.get_all_devices().iter().map(|d| (*d).clone()).collect()
}

pub fn get_device_count() -> u32 {
    let manager = PCI_MANAGER.lock();
    manager.get_device_count()
}

pub fn get_scan_count() -> u32 {
    let manager = PCI_MANAGER.lock();
    manager.get_scan_count()
}
