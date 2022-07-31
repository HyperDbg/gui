package phnt


const(
_NTREGAPI_H =  //col:13
REG_INIT_BOOT_SM = 0x0000 //col:17
REG_INIT_BOOT_SETUP = 0x0001 //col:18
REG_INIT_BOOT_ACCEPTED_BASE = 0x0002 //col:19
REG_INIT_BOOT_ACCEPTED_MAX = REG_INIT_BOOT_ACCEPTED_BASE + 999 //col:20
REG_MAX_KEY_VALUE_NAME_LENGTH = 32767 //col:22
REG_MAX_KEY_NAME_LENGTH = 512 //col:23
REG_FLAG_VOLATILE = 0x0001 //col:95
REG_FLAG_LINK = 0x0002 //col:96
REG_KEY_DONT_VIRTUALIZE = 0x0002 //col:99
REG_KEY_DONT_SILENT_FAIL = 0x0004 //col:100
REG_KEY_RECURSE_FLAG = 0x0008 //col:101
)

const(
    KeyBasicInformation // KEY_BASIC_INFORMATION = 1  //col:27
    KeyNodeInformation // KEY_NODE_INFORMATION = 2  //col:28
    KeyFullInformation // KEY_FULL_INFORMATION = 3  //col:29
    KeyNameInformation // KEY_NAME_INFORMATION = 4  //col:30
    KeyCachedInformation // KEY_CACHED_INFORMATION = 5  //col:31
    KeyFlagsInformation // KEY_FLAGS_INFORMATION = 6  //col:32
    KeyVirtualizationInformation // KEY_VIRTUALIZATION_INFORMATION = 7  //col:33
    KeyHandleTagsInformation // KEY_HANDLE_TAGS_INFORMATION = 8  //col:34
    KeyTrustInformation // KEY_TRUST_INFORMATION = 9  //col:35
    KeyLayerInformation // KEY_LAYER_INFORMATION = 10  //col:36
    MaxKeyInfoClass = 11  //col:37
)


const(
    KeyWriteTimeInformation // KEY_WRITE_TIME_INFORMATION = 1  //col:140
    KeyWow64FlagsInformation // KEY_WOW64_FLAGS_INFORMATION = 2  //col:141
    KeyControlFlagsInformation // KEY_CONTROL_FLAGS_INFORMATION = 3  //col:142
    KeySetVirtualizationInformation // KEY_SET_VIRTUALIZATION_INFORMATION = 4  //col:143
    KeySetDebugInformation = 5  //col:144
    KeySetHandleTagsInformation // KEY_HANDLE_TAGS_INFORMATION = 6  //col:145
    KeySetLayerInformation // KEY_SET_LAYER_INFORMATION = 7  //col:146
    MaxKeySetInfoClass = 8  //col:147
)


const(
    KeyValueBasicInformation // KEY_VALUE_BASIC_INFORMATION = 1  //col:189
    KeyValueFullInformation // KEY_VALUE_FULL_INFORMATION = 2  //col:190
    KeyValuePartialInformation // KEY_VALUE_PARTIAL_INFORMATION = 3  //col:191
    KeyValueFullInformationAlign64 = 4  //col:192
    KeyValuePartialInformationAlign64  // KEY_VALUE_PARTIAL_INFORMATION_ALIGN64 = 5  //col:193
    KeyValueLayerInformation // KEY_VALUE_LAYER_INFORMATION = 6  //col:194
    MaxKeyValueInfoClass = 7  //col:195
)


const(
    KeyLoadTrustClassKey  =  1  //col:243
    KeyLoadEvent = 2  //col:244
    KeyLoadToken = 3  //col:245
)


const(
    KeyAdded = 1  //col:269
    KeyRemoved = 2  //col:270
    KeyModified = 3  //col:271
)




