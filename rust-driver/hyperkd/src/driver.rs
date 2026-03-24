use alloc::boxed::Box;
use alloc::sync::Arc;
use spin::Mutex;
use wdk_sys::{
    DRIVER_OBJECT,
    PDRIVER_OBJECT,
    UNICODE_STRING,
    NTSTATUS,
    STATUS_SUCCESS,
    STATUS_UNSUCCESSFUL,
    DEVICE_TYPE,
    FILE_DEVICE_UNKNOWN,
    FILE_DEVICE_SECURE_OPEN,
    IRP_MJ_CREATE,
    IRP_MJ_CLOSE,
    IRP_MJ_DEVICE_CONTROL,
    IRP_MJ_READ,
    IRP_MJ_WRITE,
    IO_STATUS_BLOCK,
    PIRP,
    PIO_STACK_LOCATION,
    ULONG,
    PVOID,
};

use crate::*;

static DRIVER_CONTEXT: Mutex<Option<Arc<Mutex<HyperDbgDriver>>>> = Mutex::new(None);

#[no_mangle]
pub extern "system" fn DriverEntry(
    driver_object: PDRIVER_OBJECT,
    registry_path: *mut UNICODE_STRING,
) -> NTSTATUS {
    unsafe {
        (*driver_object).DriverUnload = Some(driver_unload);

        for i in 0..IRP_MJ_MAXIMUM_FUNCTION as usize {
            (*driver_object).MajorFunction[i] = Some(driver_default_dispatch);
        }

        (*driver_object).MajorFunction[IRP_MJ_CREATE as usize] = Some(driver_create);
        (*driver_object).MajorFunction[IRP_MJ_CLOSE as usize] = Some(driver_close);
        (*driver_object).MajorFunction[IRP_MJ_DEVICE_CONTROL as usize] = Some(driver_device_control);
    }

    match initialize_driver() {
        Ok(driver) => {
            let mut ctx = DRIVER_CONTEXT.lock();
            *ctx = Some(Arc::new(Mutex::new(driver)));
            STATUS_SUCCESS
        }
        Err(_) => STATUS_UNSUCCESSFUL,
    }
}

const IRP_MJ_MAXIMUM_FUNCTION: ULONG = 28;

extern "system" {
    fn IoCreateDevice(
        driver_object: PDRIVER_OBJECT,
        device_extension_size: ULONG,
        device_name: *mut UNICODE_STRING,
        device_type: DEVICE_TYPE,
        device_characteristics: ULONG,
        exclusive: bool,
        device_object: *mut PVOID,
    ) -> NTSTATUS;

    fn IoCreateSymbolicLink(
        symbolic_link_name: *mut UNICODE_STRING,
        device_name: *mut UNICODE_STRING,
    ) -> NTSTATUS;

    fn IoDeleteDevice(device_object: PVOID);

    fn IoCompleteRequest(irp: PIRP, priority_boost: i32);

    fn IoGetCurrentIrpStackLocation(irp: PIRP) -> PIO_STACK_LOCATION;
}

extern "system" fn driver_unload(driver_object: PDRIVER_OBJECT) {
    let mut ctx = DRIVER_CONTEXT.lock();
    if let Some(driver) = ctx.take() {
        let mut driver = driver.lock();
        let _ = driver.terminate();
    }
    drop(ctx);

    unsafe {
        IoDeleteDevice((*driver_object).DeviceObject);
    }
}

extern "system" fn driver_default_dispatch(
    device_object: PVOID,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

extern "system" fn driver_create(
    device_object: PVOID,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

extern "system" fn driver_close(
    device_object: PVOID,
    irp: PIRP,
) -> NTSTATUS {
    unsafe {
        (*irp).IoStatus.Status = STATUS_SUCCESS;
        (*irp).IoStatus.Information = 0;
        IoCompleteRequest(irp, 0);
    }
    STATUS_SUCCESS
}

extern "system" fn driver_device_control(
    device_object: PVOID,
    irp: PIRP,
) -> NTSTATUS {
    let ctx = DRIVER_CONTEXT.lock();
    let driver = match ctx.as_ref() {
        Some(d) => d.clone(),
        None => {
            unsafe {
                (*irp).IoStatus.Status = STATUS_UNSUCCESSFUL;
                (*irp).IoStatus.Information = 0;
                IoCompleteRequest(irp, 0);
            }
            return STATUS_UNSUCCESSFUL;
        }
    };
    drop(ctx);

    let mut driver = driver.lock();

    unsafe {
        let stack = IoGetCurrentIrpStackLocation(irp);
        let control_code = (*stack).Parameters.DeviceIoControl.IoControlCode;
        let input_buffer = (*irp).AssociatedIrp.SystemBuffer as *const u8;
        let input_length = (*stack).Parameters.DeviceIoControl.InputBufferLength as usize;
        let output_buffer = (*irp).AssociatedIrp.SystemBuffer as *mut u8;
        let output_length = (*stack).Parameters.DeviceIoControl.OutputBufferLength as usize;

        let result = handle_ioctl(
            &mut driver,
            control_code,
            core::slice::from_raw_parts(input_buffer, input_length),
            core::slice::from_raw_parts_mut(output_buffer, output_length),
        );

        match result {
            Ok(bytes_written) => {
                (*irp).IoStatus.Status = STATUS_SUCCESS;
                (*irp).IoStatus.Information = bytes_written as u64;
            }
            Err(_) => {
                (*irp).IoStatus.Status = STATUS_UNSUCCESSFUL;
                (*irp).IoStatus.Information = 0;
            }
        }

        IoCompleteRequest(irp, 0);
    }

    STATUS_SUCCESS
}

fn handle_ioctl(
    driver: &mut HyperDbgDriver,
    control_code: u32,
    input: &[u8],
    output: &mut [u8],
) -> Result<usize, DriverError> {
    match control_code {
        0x800 => {
            driver.initialize()?;
            Ok(0)
        }
        0x801 => {
            driver.terminate()?;
            Ok(0)
        }
        0x802 => {
            if input.len() < 6 {
                return Err(DriverError::InvalidAddress);
            }
            let port = u16::from_le_bytes([input[0], input[1]]);
            let addr_len = input[2] as usize;
            let addr = core::str::from_utf8(&input[3..3+addr_len])
                .map_err(|_| DriverError::InvalidAddress)?;
            driver.connect(addr, port)?;
            Ok(0)
        }
        0x803 => {
            driver.disconnect();
            Ok(0)
        }
        0x804 => {
            if input.len() < 4 {
                return Err(DriverError::ProcessNotFound);
            }
            let pid = u32::from_le_bytes([input[0], input[1], input[2], input[3]]);
            driver.attach_process(pid)?;
            Ok(0)
        }
        0x805 => {
            driver.detach_process()?;
            Ok(0)
        }
        0x806 => {
            if input.len() < 12 {
                return Err(DriverError::InvalidAddress);
            }
            let address = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            let bp_type = u32::from_le_bytes([input[8], input[9], input[10], input[11]]);
            let bp_type = match bp_type {
                0 => BreakpointType::Software,
                1 => BreakpointType::Hardware,
                2 => BreakpointType::Hidden,
                3 => BreakpointType::Ept,
                _ => return Err(DriverError::InvalidAddress),
            };
            let id = driver.set_breakpoint(address, bp_type)?;
            if output.len() >= 8 {
                output[..8].copy_from_slice(&id.to_le_bytes());
                Ok(8)
            } else {
                Ok(0)
            }
        }
        0x807 => {
            if input.len() < 8 {
                return Err(DriverError::BreakpointNotFound);
            }
            let id = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            driver.remove_breakpoint(id)?;
            Ok(0)
        }
        0x808 => {
            driver.continue_execution()?;
            Ok(0)
        }
        0x809 => {
            driver.pause_execution()?;
            Ok(0)
        }
        0x80A => {
            driver.step_into()?;
            Ok(0)
        }
        0x80B => {
            driver.step_over()?;
            Ok(0)
        }
        0x80C => {
            driver.step_out()?;
            Ok(0)
        }
        0x80D => {
            if input.len() < 4 {
                return Err(DriverError::InvalidCoreId);
            }
            let core_id = u32::from_le_bytes([input[0], input[1], input[2], input[3]]);
            let regs = driver.read_registers(core_id)?;
            if output.len() >= 272 {
                let offset = serialize_registers(&regs, output);
                Ok(offset)
            } else {
                Ok(0)
            }
        }
        0x80E => {
            if input.len() < 12 {
                return Err(DriverError::InvalidAddress);
            }
            let address = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            let size = u32::from_le_bytes([input[8], input[9], input[10], input[11]]) as usize;
            let data = driver.read_memory(address, size)?;
            let copy_size = core::cmp::min(data.len(), output.len());
            output[..copy_size].copy_from_slice(&data[..copy_size]);
            Ok(copy_size)
        }
        0x80F => {
            if input.len() < 8 {
                return Err(DriverError::InvalidAddress);
            }
            let address = u64::from_le_bytes([
                input[0], input[1], input[2], input[3],
                input[4], input[5], input[6], input[7],
            ]);
            let data = &input[8..];
            driver.write_memory(address, data)?;
            Ok(0)
        }
        _ => Err(DriverError::InvalidAddress),
    }
}

fn serialize_registers(regs: &RegisterState, output: &mut [u8]) -> usize {
    let mut offset = 0;
    output[offset..offset+8].copy_from_slice(&regs.rax.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rbx.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rcx.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rdx.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rsi.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rdi.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rbp.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rsp.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r8.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r9.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r10.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r11.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r12.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r13.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r14.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.r15.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rip.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.rflags.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.cr0.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.cr2.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.cr3.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.cr4.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr0.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr1.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr2.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr3.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr6.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.dr7.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.gdtr.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.gs_base.to_le_bytes()); offset += 8;
    output[offset..offset+8].copy_from_slice(&regs.fs_base.to_le_bytes());
    offset + 8
}

fn initialize_driver() -> Result<HyperDbgDriver, DriverError> {
    HyperDbgDriver::new()
}
