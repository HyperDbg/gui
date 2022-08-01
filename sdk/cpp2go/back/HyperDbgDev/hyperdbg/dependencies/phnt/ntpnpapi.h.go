package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntpnpapi.h.back

const(
_NTPNPAPI_H =  //col:1
)

const(
    HardwareProfileChangeEvent = 1  //col:3
    TargetDeviceChangeEvent = 2  //col:4
    DeviceClassChangeEvent = 3  //col:5
    CustomDeviceEvent = 4  //col:6
    DeviceInstallEvent = 5  //col:7
    DeviceArrivalEvent = 6  //col:8
    PowerEvent = 7  //col:9
    VetoEvent = 8  //col:10
    BlockedDriverEvent = 9  //col:11
    InvalidIDEvent = 10  //col:12
    MaxPlugEventCategory = 11  //col:13
)


const(
    PlugPlayControlEnumerateDevice  = 1  //col:17
    PlugPlayControlRegisterNewDevice  = 2  //col:18
    PlugPlayControlDeregisterDevice  = 3  //col:19
    PlugPlayControlInitializeDevice  = 4  //col:20
    PlugPlayControlStartDevice  = 5  //col:21
    PlugPlayControlUnlockDevice  = 6  //col:22
    PlugPlayControlQueryAndRemoveDevice  = 7  //col:23
    PlugPlayControlUserResponse  = 8  //col:24
    PlugPlayControlGenerateLegacyDevice  = 9  //col:25
    PlugPlayControlGetInterfaceDeviceList  = 10  //col:26
    PlugPlayControlProperty  = 11  //col:27
    PlugPlayControlDeviceClassAssociation  = 12  //col:28
    PlugPlayControlGetRelatedDevice  = 13  //col:29
    PlugPlayControlGetInterfaceDeviceAlias  = 14  //col:30
    PlugPlayControlDeviceStatus  = 15  //col:31
    PlugPlayControlGetDeviceDepth  = 16  //col:32
    PlugPlayControlQueryDeviceRelations  = 17  //col:33
    PlugPlayControlTargetDeviceRelation  = 18  //col:34
    PlugPlayControlQueryConflictList  = 19  //col:35
    PlugPlayControlRetrieveDock  = 20  //col:36
    PlugPlayControlResetDevice  = 21  //col:37
    PlugPlayControlHaltDevice  = 22  //col:38
    PlugPlayControlGetBlockedDriverList  = 23  //col:39
    PlugPlayControlGetDeviceInterfaceEnabled  = 24  //col:40
    MaxPlugPlayControl = 25  //col:41
)



type PLUGPLAY_EVENT_BLOCK struct{
EventGuid GUID
EventCategory PLUGPLAY_EVENT_CATEGORY
Result PULONG
Flags ULONG
TotalSize ULONG
DeviceObject PVOID
Union union
Struct struct
ClassGuid GUID
SymbolicLinkName[1] WCHAR
}




