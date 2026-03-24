use crate::memory::MemoryManager;
use crate::process::Process;
use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
pub enum ExtensionCommandError {
    NotInitialized,
    InvalidParameter,
    InvalidAddress,
    InvalidProcessId,
    OperationFailed,
    AccessDenied,
    InsufficientMemory,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ApicRequest {
    pub apic_type: u32,
    pub kernel_status: u32,
    pub is_using_x2apic: bool,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct LapicPage {
    pub reserved: [u8; 0x400],
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct IoApicEntryPackets {
    pub entries: [u32; 24],
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct InterruptDescriptorTableEntriesPackets {
    pub entries: [u64; 256],
    pub kernel_status: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct Va2PaAndPa2VaCommands {
    pub is_virtual2physical: bool,
    pub virtual_address: u64,
    pub physical_address: u64,
    pub process_id: u32,
    pub kernel_status: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct ReadPageTableEntriesDetails {
    pub virtual_address: u64,
    pub process_id: u32,
    pub pml4e_virtual_address: u64,
    pub pml4e_value: u64,
    pub pdpte_virtual_address: u64,
    pub pdpte_value: u64,
    pub pde_virtual_address: u64,
    pub pde_value: u64,
    pub pte_virtual_address: u64,
    pub pte_value: u64,
    pub kernel_status: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct PciDeviceInfo {
    pub bus: u8,
    pub device: u8,
    pub function: u8,
    pub config_space: [u8; 256],
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct PcitreeRequestResponsePacket {
    pub device_info_list: [PciDeviceInfo; 256],
    pub device_info_list_num: u32,
    pub kernel_status: u32,
}

#[repr(C)]
#[derive(Debug, Clone, Copy)]
pub struct PcidevinfoRequestResponsePacket {
    pub device_info: PciDeviceInfo,
    pub print_raw: bool,
    pub kernel_status: u32,
}

pub unsafe fn perform_actions_for_apic_requests(
    apic_request: &mut ApicRequest,
    lapic_buffer: Option<&mut LapicPage>,
    io_apic_buffer: Option<&mut IoApicEntryPackets>,
) -> Result<(), ExtensionCommandError> {
    match apic_request.apic_type {
        0 => {
            if let Some(lapic) = lapic_buffer {
                if vm_func_apic_store_local_apic_fields(lapic, &mut apic_request.is_using_x2apic) {
                    apic_request.kernel_status = 0;
                    Ok(())
                } else {
                    apic_request.kernel_status = 1;
                    Err(ExtensionCommandError::OperationFailed)
                }
            } else {
                Err(ExtensionCommandError::InvalidParameter)
            }
        }
        1 => {
            if let Some(io_apic) = io_apic_buffer {
                if vm_func_apic_store_io_apic_fields(io_apic) {
                    apic_request.kernel_status = 0;
                    Ok(())
                } else {
                    apic_request.kernel_status = 1;
                    Err(ExtensionCommandError::OperationFailed)
                }
            } else {
                Err(ExtensionCommandError::InvalidParameter)
            }
        }
        _ => {
            apic_request.kernel_status = 1;
            Err(ExtensionCommandError::InvalidParameter)
        }
    }
}

unsafe fn vm_func_apic_store_local_apic_fields(
    lapic: &mut LapicPage,
    is_using_x2apic: &mut bool,
) -> bool {
    extern "C" {
        fn rdmsr(msr: u32) -> u64;
    }

    let apic_base = rdmsr(0x1B);
    *is_using_x2apic = (apic_base & (1 << 10)) != 0;

    if *is_using_x2apic {
        for i in 0..0x400 {
            let offset = i * 8;
            let msr_value = rdmsr(0x800 + (i as u32));
            let ptr = (lapic as *mut LapicPage) as *mut u64;
            *ptr.add(i) = msr_value;
        }
    } else {
        let apic_base = apic_base & 0xFFFFF000;
        for i in 0..0x400 {
            let offset = i * 4;
            let ptr = (apic_base + offset as u64) as *const u32;
            let value = *ptr;
            let lapic_ptr = (lapic as *mut LapicPage) as *mut u32;
            *lapic_ptr.add(i) = value;
        }
    }

    true
}

unsafe fn vm_func_apic_store_io_apic_fields(
    io_apic: &mut IoApicEntryPackets,
) -> bool {
    extern "C" {
        fn read_io_port(port: u16) -> u32;
    }

    let io_apic_base = 0xFEC00000u64;

    for i in 0..24 {
        let reg_select = (io_apic_base + 0x00) as *mut u32;
        let reg_window = (io_apic_base + 0x10) as *mut u32;

        *reg_select = 0x10 + (i as u32) * 2;
        let lower = *reg_window;

        *reg_select = 0x11 + (i as u32) * 2;
        let upper = *reg_window;

        io_apic.entries[i] = ((upper as u64) << 32) | (lower as u64);
    }

    true
}

pub unsafe fn perform_query_idt_entries_request(
    idt_query_request: &mut InterruptDescriptorTableEntriesPackets,
    read_from_vmx_root: bool,
) -> Result<(), ExtensionCommandError> {
    vm_func_idt_query_entries(idt_query_request, read_from_vmx_root);
    idt_query_request.kernel_status = 0;
    Ok(())
}

unsafe fn vm_func_idt_query_entries(
    idt_query_request: &mut InterruptDescriptorTableEntriesPackets,
    read_from_vmx_root: bool,
) {
    extern "C" {
        fn read_idt_entry(index: u32) -> u64;
    }

    for i in 0..256 {
        idt_query_request.entries[i] = read_idt_entry(i);
    }
}

pub unsafe fn va2pa_and_pa2va(
    address_details: &mut Va2PaAndPa2VaCommands,
    operate_on_vmx_root: bool,
) -> Result<(), ExtensionCommandError> {
    if operate_on_vmx_root {
        if address_details.is_virtual2physical {
            address_details.physical_address = 
                virtual_address_to_physical_address_on_target_process(address_details.virtual_address);

            if address_details.physical_address == 0 {
                address_details.kernel_status = 1;
                return Err(ExtensionCommandError::InvalidAddress);
            } else {
                address_details.kernel_status = 0;
                return Ok(());
            }
        } else {
            address_details.virtual_address = 
                physical_address_to_virtual_address_on_target_process(address_details.physical_address);
            address_details.kernel_status = 0;
            return Ok(());
        }
    } else {
        extern "C" {
            fn ps_get_current_process_id() -> u32;
        }

        if address_details.process_id == ps_get_current_process_id() {
            if address_details.is_virtual2physical {
                address_details.physical_address = 
                    virtual_address_to_physical_address(address_details.virtual_address);

                if address_details.physical_address == 0 {
                    address_details.kernel_status = 1;
                    return Err(ExtensionCommandError::InvalidAddress);
                } else {
                    address_details.kernel_status = 0;
                    return Ok(());
                }
            } else {
                address_details.virtual_address = 
                    physical_address_to_virtual_address(address_details.physical_address);
                address_details.kernel_status = 0;
                return Ok(());
            }
        } else {
            if !is_process_exist(address_details.process_id) {
                address_details.kernel_status = 2;
                return Err(ExtensionCommandError::InvalidProcessId);
            }

            if address_details.is_virtual2physical {
                address_details.physical_address = 
                    virtual_address_to_physical_address_by_process_id(
                        address_details.virtual_address,
                        address_details.process_id,
                    );

                if address_details.physical_address == 0 {
                    address_details.kernel_status = 1;
                    return Err(ExtensionCommandError::InvalidAddress);
                } else {
                    address_details.kernel_status = 0;
                    return Ok(());
                }
            } else {
                address_details.virtual_address = 
                    physical_address_to_virtual_address_by_process_id(
                        address_details.physical_address,
                        address_details.process_id,
                    );
                address_details.kernel_status = 0;
                return Ok(());
            }
        }
    }
}

unsafe fn virtual_address_to_physical_address_on_target_process(
    virtual_address: u64,
) -> u64 {
    extern "C" {
        fn mmlayout_virtual_to_physical(virtual_address: u64) -> u64;
    }

    mmlayout_virtual_to_physical(virtual_address)
}

unsafe fn physical_address_to_virtual_address_on_target_process(
    physical_address: u64,
) -> u64 {
    extern "C" {
        fn mmlayout_physical_to_virtual(physical_address: u64) -> u64;
    }

    mmlayout_physical_to_virtual(physical_address)
}

unsafe fn virtual_address_to_physical_address(virtual_address: u64) -> u64 {
    extern "C" {
        fn mmlayout_virtual_to_physical(virtual_address: u64) -> u64;
    }

    mmlayout_virtual_to_physical(virtual_address)
}

unsafe fn physical_address_to_virtual_address(physical_address: u64) -> u64 {
    extern "C" {
        fn mmlayout_physical_to_virtual(physical_address: u64) -> u64;
    }

    mmlayout_physical_to_virtual(physical_address)
}

unsafe fn virtual_address_to_physical_address_by_process_id(
    virtual_address: u64,
    process_id: u32,
) -> u64 {
    extern "C" {
        fn mmlayout_switch_to_process(process_id: u32) -> u64;
        fn mmlayout_switch_to_previous() -> u64;
    }

    let restore_cr3 = mmlayout_switch_to_process(process_id);
    let physical_address = mmlayout_virtual_to_physical(virtual_address);
    mmlayout_switch_to_previous();

    physical_address
}

unsafe fn physical_address_to_virtual_address_by_process_id(
    physical_address: u64,
    process_id: u32,
) -> u64 {
    extern "C" {
        fn mmlayout_switch_to_process(process_id: u32) -> u64;
        fn mmlayout_switch_to_previous() -> u64;
    }

    let restore_cr3 = mmlayout_switch_to_process(process_id);
    let virtual_address = mmlayout_physical_to_virtual(physical_address);
    mmlayout_switch_to_previous();

    virtual_address
}

unsafe fn is_process_exist(process_id: u32) -> bool {
    extern "C" {
        fn ps_lookup_process_by_process_id(process_id: u32, eprocess: *mut *mut u8) -> u32;
        fn ob_dereference_object(object: *mut u8);
    }

    let mut eprocess: *mut u8 = core::ptr::null_mut();
    let status = ps_lookup_process_by_process_id(process_id, &mut eprocess);

    if status == 0 {
        ob_dereference_object(eprocess);
        return true;
    }

    false
}

pub unsafe fn extension_command_pte(
    pte_details: &mut ReadPageTableEntriesDetails,
    is_operating_in_vmx_root: bool,
) -> Result<(), ExtensionCommandError> {
    let mut restore_cr3: u64 = 0;

    if is_operating_in_vmx_root {
        if virtual_address_to_physical_address_on_target_process(pte_details.virtual_address) == 0 {
            pte_details.kernel_status = 1;
            return Err(ExtensionCommandError::InvalidAddress);
        }

        extern "C" {
            fn mmlayout_switch_to_current_process() -> u64;
        }

        restore_cr3 = mmlayout_switch_to_current_process();
    } else {
        extern "C" {
            fn ps_get_current_process_id() -> u32;
        }

        if pte_details.process_id != ps_get_current_process_id() {
            if !is_process_exist(pte_details.process_id) {
                pte_details.kernel_status = 2;
                return Err(ExtensionCommandError::InvalidProcessId);
            }

            extern "C" {
                fn mmlayout_switch_to_process(process_id: u32) -> u64;
            }

            restore_cr3 = mmlayout_switch_to_process(pte_details.process_id);
        }

        if virtual_address_to_physical_address(pte_details.virtual_address) == 0 {
            pte_details.kernel_status = 1;
            let _ = restore_cr3;
            return Err(ExtensionCommandError::InvalidAddress);
        }
    }

    extern "C" {
        fn memory_mapper_get_pte_va(virtual_address: u64, level: u32) -> u64;
    }

    const PAGING_LEVEL_PAGE_MAP_LEVEL_4: u32 = 3;
    const PAGING_LEVEL_PAGE_DIRECTORY_POINTER_TABLE: u32 = 2;
    const PAGING_LEVEL_PAGE_DIRECTORY: u32 = 1;
    const PAGING_LEVEL_PAGE_TABLE: u32 = 0;

    let pml4e_va = memory_mapper_get_pte_va(pte_details.virtual_address, PAGING_LEVEL_PAGE_MAP_LEVEL_4);
    if pml4e_va != 0 {
        pte_details.pml4e_virtual_address = pml4e_va;
        pte_details.pml4e_value = *((pml4e_va) as *const u64);
    }

    let pdpte_va = memory_mapper_get_pte_va(pte_details.virtual_address, PAGING_LEVEL_PAGE_DIRECTORY_POINTER_TABLE);
    if pdpte_va != 0 {
        pte_details.pdpte_virtual_address = pdpte_va;
        pte_details.pdpte_value = *((pdpte_va) as *const u64);
    }

    let pde_va = memory_mapper_get_pte_va(pte_details.virtual_address, PAGING_LEVEL_PAGE_DIRECTORY);
    if pde_va != 0 {
        pte_details.pde_virtual_address = pde_va;
        pte_details.pde_value = *((pde_va) as *const u64);
    }

    let pte_va = memory_mapper_get_pte_va(pte_details.virtual_address, PAGING_LEVEL_PAGE_TABLE);
    if pte_va != 0 {
        pte_details.pte_virtual_address = pte_va;
        pte_details.pte_value = *((pte_va) as *const u64);
    }

    pte_details.kernel_status = 0;

    if restore_cr3 != 0 {
        extern "C" {
            fn mmlayout_switch_to_previous();
        }

        mmlayout_switch_to_previous();
    }

    Ok(())
}

pub unsafe fn extension_command_change_all_msr_bitmap_read_all_cores(bitmap_mask: u64) {
    crate::halted_core::broadcast_change_all_msr_bitmap_read_all_cores(bitmap_mask);
}

pub unsafe fn extension_command_reset_change_all_msr_bitmap_read_all_cores() {
    crate::halted_core::broadcast_reset_change_all_msr_bitmap_read_all_cores();
}

pub unsafe fn extension_command_change_all_msr_bitmap_write_all_cores(bitmap_mask: u64) {
    crate::halted_core::broadcast_change_all_msr_bitmap_write_all_cores(bitmap_mask);
}

pub unsafe fn extension_command_reset_all_msr_bitmap_write_all_cores() {
    crate::halted_core::broadcast_reset_all_msr_bitmap_write_all_cores();
}

pub unsafe fn extension_command_enable_rdtsc_exiting_all_cores() {
    crate::halted_core::broadcast_enable_rdtsc_exiting_all_cores();
}

pub unsafe fn extension_command_disable_rdtsc_exiting_all_cores() {
    crate::halted_core::broadcast_disable_rdtsc_exiting_all_cores();
}

pub unsafe fn extension_command_disable_rdtsc_exiting_for_clearing_events_all_cores() {
    crate::halted_core::broadcast_disable_rdtsc_exiting_for_clearing_events_all_cores();
}

pub unsafe fn extension_command_set_exception_bitmap_all_cores(exception_index: u64) {
    crate::halted_core::broadcast_set_exception_bitmap_all_cores(exception_index);
}

pub unsafe fn extension_command_unset_exception_bitmap_all_cores(exception_index: u64) {
    crate::halted_core::broadcast_unset_exception_bitmap_all_cores(exception_index);
}

pub unsafe fn extension_command_reset_exception_bitmap_all_cores() {
    crate::halted_core::broadcast_reset_exception_bitmap_all_cores();
}

pub unsafe fn extension_command_enable_mov_control_register_exiting_all_cores(options: u64) {
    crate::halted_core::broadcast_enable_mov_control_register_exiting_all_cores(options);
}

pub unsafe fn extension_command_disable_mov_to_control_registers_exiting_all_cores(options: u64) {
    crate::halted_core::broadcast_disable_mov_to_control_registers_exiting_all_cores(options);
}

pub unsafe fn extension_command_enable_mov_debug_registers_exiting_all_cores() {
    crate::halted_core::broadcast_enable_mov_debug_registers_exiting_all_cores();
}

pub unsafe fn extension_command_disable_mov_debug_registers_exiting_all_cores() {
    crate::halted_core::broadcast_disable_mov_debug_registers_exiting_all_cores();
}

pub unsafe fn extension_command_set_external_interrupt_exiting_all_cores() {
    crate::halted_core::broadcast_set_external_interrupt_exiting_all_cores();
}

pub unsafe fn extension_command_unset_external_interrupt_exiting_only_on_clearing_interrupt_events_all_cores() {
    crate::halted_core::broadcast_unset_external_interrupt_exiting_only_on_clearing_interrupt_events_all_cores();
}

pub unsafe fn extension_command_io_bitmap_change_all_cores(port: u64) {
    crate::halted_core::broadcast_io_bitmap_change_all_cores(port);
}

pub unsafe fn extension_command_io_bitmap_reset_all_cores() {
    crate::halted_core::broadcast_io_bitmap_reset_all_cores();
}

pub unsafe fn extension_command_pcitree(
    pcitree_packet: &mut PcitreeRequestResponsePacket,
    operate_on_vmx_root: bool,
) -> Result<(), ExtensionCommandError> {
    extern "C" {
        fn pci_read_cam(bus: u8, device: u8, function: u8, offset: u8, size: u32) -> u32;
    }

    const BUS_MAX_NUM: u8 = 256;
    const DEVICE_MAX_NUM: u8 = 32;
    const FUNCTION_MAX_NUM: u8 = 8;
    const DEV_MAX_NUM: u32 = 256;

    let mut dev_num: u32 = 0;

    for b in 0..BUS_MAX_NUM {
        for d in 0..DEVICE_MAX_NUM {
            for f in 0..FUNCTION_MAX_NUM {
                let device_id_vendor_id = pci_read_cam(b, d, f, 0, 4);

                if device_id_vendor_id != 0xFFFFFFFF {
                    pcitree_packet.device_info_list[dev_num as usize].bus = b;
                    pcitree_packet.device_info_list[dev_num as usize].device = d;
                    pcitree_packet.device_info_list[dev_num as usize].function = f;
                    
                    let vendor_id = (device_id_vendor_id & 0xFFFF) as u16;
                    let device_id = ((device_id_vendor_id >> 16) & 0xFFFF) as u16;
                    
                    let config_space_ptr = &mut pcitree_packet.device_info_list[dev_num as usize].config_space;
                    config_space_ptr[0] = (vendor_id & 0xFF) as u8;
                    config_space_ptr[1] = ((vendor_id >> 8) & 0xFF) as u8;
                    config_space_ptr[2] = (device_id & 0xFF) as u8;
                    config_space_ptr[3] = ((device_id >> 8) & 0xFF) as u8;

                    let class_code = pci_read_cam(b, d, f, 8, 4);
                    config_space_ptr[9] = ((class_code >> 24) & 0xFF) as u8;
                    config_space_ptr[10] = ((class_code >> 16) & 0xFF) as u8;
                    config_space_ptr[11] = ((class_code >> 8) & 0xFF) as u8;

                    dev_num += 1;
                    if dev_num == DEV_MAX_NUM {
                        return Ok(());
                    }
                }
            }
        }
    }

    pcitree_packet.device_info_list_num = dev_num;

    if dev_num > 0 {
        pcitree_packet.kernel_status = 0;
        Ok(())
    } else {
        pcitree_packet.kernel_status = 1;
        Err(ExtensionCommandError::InvalidAddress)
    }
}

pub unsafe fn extension_command_pcidevinfo(
    pcidevinfo_packet: &mut PcidevinfoRequestResponsePacket,
    operate_on_vmx_root: bool,
) -> Result<(), ExtensionCommandError> {
    extern "C" {
        fn pci_read_cam(bus: u8, device: u8, function: u8, offset: u8, size: u32) -> u32;
    }

    const CAM_CONFIG_SPACE_LENGTH: u16 = 256;

    let device_id_vendor_id = pci_read_cam(
        pcidevinfo_packet.device_info.bus,
        pcidevinfo_packet.device_info.device,
        pcidevinfo_packet.device_info.function,
        0,
        4,
    );

    if device_id_vendor_id != 0xFFFFFFFF {
        let config_space_ptr = &mut pcidevinfo_packet.device_info.config_space;
        for i in (0..CAM_CONFIG_SPACE_LENGTH).step_by(4) {
            let value = pci_read_cam(
                pcidevinfo_packet.device_info.bus,
                pcidevinfo_packet.device_info.device,
                pcidevinfo_packet.device_info.function,
                i as u8,
                4,
            );
            
            config_space_ptr[i as usize] = (value & 0xFF) as u8;
            config_space_ptr[(i + 1) as usize] = ((value >> 8) & 0xFF) as u8;
            config_space_ptr[(i + 2) as usize] = ((value >> 16) & 0xFF) as u8;
            config_space_ptr[(i + 3) as usize] = ((value >> 24) & 0xFF) as u8;
        }

        pcidevinfo_packet.kernel_status = 0;
        Ok(())
    } else {
        pcidevinfo_packet.device_info.config_space[0] = 0xFF;
        pcidevinfo_packet.device_info.config_space[1] = 0xFF;
        pcidevinfo_packet.device_info.config_space[2] = 0xFF;
        pcidevinfo_packet.device_info.config_space[3] = 0xFF;
        pcidevinfo_packet.kernel_status = 1;
        Err(ExtensionCommandError::InvalidAddress)
    }
}