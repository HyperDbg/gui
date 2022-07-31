package phnt
const(
_NTPNPAPI_H =  //col:13
)
type     HardwareProfileChangeEvent uint32
const(
    HardwareProfileChangeEvent PLUGPLAY_EVENT_CATEGORY = 1  //col:17
    TargetDeviceChangeEvent PLUGPLAY_EVENT_CATEGORY = 2  //col:18
    DeviceClassChangeEvent PLUGPLAY_EVENT_CATEGORY = 3  //col:19
    CustomDeviceEvent PLUGPLAY_EVENT_CATEGORY = 4  //col:20
    DeviceInstallEvent PLUGPLAY_EVENT_CATEGORY = 5  //col:21
    DeviceArrivalEvent PLUGPLAY_EVENT_CATEGORY = 6  //col:22
    PowerEvent PLUGPLAY_EVENT_CATEGORY = 7  //col:23
    VetoEvent PLUGPLAY_EVENT_CATEGORY = 8  //col:24
    BlockedDriverEvent PLUGPLAY_EVENT_CATEGORY = 9  //col:25
    InvalidIDEvent PLUGPLAY_EVENT_CATEGORY = 10  //col:26
    MaxPlugEventCategory PLUGPLAY_EVENT_CATEGORY = 11  //col:27
)
type     PlugPlayControlEnumerateDevice // PLUGPLAY_CONTROL_ENUMERATE_DEVICE_DATA uint32
const(
    PlugPlayControlEnumerateDevice // PLUGPLAY_CONTROL_ENUMERATE_DEVICE_DATA PLUGPLAY_CONTROL_CLASS = 1  //col:86
    PlugPlayControlRegisterNewDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 2  //col:87
    PlugPlayControlDeregisterDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 3  //col:88
    PlugPlayControlInitializeDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 4  //col:89
    PlugPlayControlStartDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 5  //col:90
    PlugPlayControlUnlockDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 6  //col:91
    PlugPlayControlQueryAndRemoveDevice // PLUGPLAY_CONTROL_QUERY_AND_REMOVE_DATA PLUGPLAY_CONTROL_CLASS = 7  //col:92
    PlugPlayControlUserResponse // PLUGPLAY_CONTROL_USER_RESPONSE_DATA PLUGPLAY_CONTROL_CLASS = 8  //col:93
    PlugPlayControlGenerateLegacyDevice // PLUGPLAY_CONTROL_LEGACY_DEVGEN_DATA PLUGPLAY_CONTROL_CLASS = 9  //col:94
    PlugPlayControlGetInterfaceDeviceList // PLUGPLAY_CONTROL_INTERFACE_LIST_DATA PLUGPLAY_CONTROL_CLASS = 10  //col:95
    PlugPlayControlProperty // PLUGPLAY_CONTROL_PROPERTY_DATA PLUGPLAY_CONTROL_CLASS = 11  //col:96
    PlugPlayControlDeviceClassAssociation // PLUGPLAY_CONTROL_CLASS_ASSOCIATION_DATA PLUGPLAY_CONTROL_CLASS = 12  //col:97
    PlugPlayControlGetRelatedDevice // PLUGPLAY_CONTROL_RELATED_DEVICE_DATA PLUGPLAY_CONTROL_CLASS = 13  //col:98
    PlugPlayControlGetInterfaceDeviceAlias // PLUGPLAY_CONTROL_INTERFACE_ALIAS_DATA PLUGPLAY_CONTROL_CLASS = 14  //col:99
    PlugPlayControlDeviceStatus // PLUGPLAY_CONTROL_STATUS_DATA PLUGPLAY_CONTROL_CLASS = 15  //col:100
    PlugPlayControlGetDeviceDepth // PLUGPLAY_CONTROL_DEPTH_DATA PLUGPLAY_CONTROL_CLASS = 16  //col:101
    PlugPlayControlQueryDeviceRelations // PLUGPLAY_CONTROL_DEVICE_RELATIONS_DATA PLUGPLAY_CONTROL_CLASS = 17  //col:102
    PlugPlayControlTargetDeviceRelation // PLUGPLAY_CONTROL_TARGET_RELATION_DATA PLUGPLAY_CONTROL_CLASS = 18  //col:103
    PlugPlayControlQueryConflictList // PLUGPLAY_CONTROL_CONFLICT_LIST PLUGPLAY_CONTROL_CLASS = 19  //col:104
    PlugPlayControlRetrieveDock // PLUGPLAY_CONTROL_RETRIEVE_DOCK_DATA PLUGPLAY_CONTROL_CLASS = 20  //col:105
    PlugPlayControlResetDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 21  //col:106
    PlugPlayControlHaltDevice // PLUGPLAY_CONTROL_DEVICE_CONTROL_DATA PLUGPLAY_CONTROL_CLASS = 22  //col:107
    PlugPlayControlGetBlockedDriverList // PLUGPLAY_CONTROL_BLOCKED_DRIVER_DATA PLUGPLAY_CONTROL_CLASS = 23  //col:108
    PlugPlayControlGetDeviceInterfaceEnabled // PLUGPLAY_CONTROL_DEVICE_INTERFACE_ENABLED PLUGPLAY_CONTROL_CLASS = 24  //col:109
    MaxPlugPlayControl PLUGPLAY_CONTROL_CLASS = 25  //col:110
)
type (
Ntpnpapi interface{
 * Attribution 4.0 International ()(ok bool)//col:28
}

)
func NewNtpnpapi() { return & ntpnpapi{} }
func (n *ntpnpapi) * Attribution 4.0 International ()(ok bool){//col:28
return true
}

