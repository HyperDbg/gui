#[cfg(feature = "secure-device")]
use wdk_sys::{BOOLEAN, GUID, PCUNICODE_STRING, PDEVICE_OBJECT, PDRIVER_OBJECT, ULONG};

use wdk_sys::{
    KPROCESSOR_MODE, NTSTATUS, PEPROCESS, PIO_STACK_LOCATION, PIRP, PSIZE_T, PVOID, SIZE_T,
};

#[allow(non_snake_case)]
pub unsafe fn IoGetCurrentIrpStackLocation(p_irp: PIRP) -> PIO_STACK_LOCATION {
    unsafe {
        assert!((*p_irp).CurrentLocation <= (*p_irp).StackCount + 1);
        (*p_irp)
            .Tail
            .Overlay
            .__bindgen_anon_2
            .__bindgen_anon_1
            .CurrentStackLocation
    }
}

#[allow(non_snake_case)]
unsafe extern "C" {
    pub fn MmCopyVirtualMemory(
        SourceProcess: PEPROCESS,
        SourceAddress: PVOID,
        TargetProcess: PEPROCESS,
        TargetAddress: PVOID,
        BufferSize: SIZE_T,
        PreviousMode: KPROCESSOR_MODE,
        ReturnSize: PSIZE_T,
    ) -> NTSTATUS;

    #[cfg(feature = "secure-device")]
    pub fn WdmlibIoCreateDeviceSecure(
        DriverObject: PDRIVER_OBJECT,
        DeviceExtensionSize: ULONG,
        DeviceName: PCUNICODE_STRING,
        DeviceType: ULONG,
        DeviceCharacteristics: ULONG,
        Exclusive: BOOLEAN,
        DefaultSDDLString: PCUNICODE_STRING,
        DeviceClassGuid: *const GUID,
        DeviceObject: *mut PDEVICE_OBJECT,
    ) -> NTSTATUS;
}
