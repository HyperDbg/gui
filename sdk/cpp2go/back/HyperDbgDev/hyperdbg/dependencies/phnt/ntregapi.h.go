package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntregapi.h.back

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
// = REG_FORCE_UNLOAD            1 //col:599
// = REG_UNLOAD_LEGAL_FLAGS      (REG_FORCE_UNLOAD) //col:600
)

type     KeyBasicInformation // KEY_BASIC_INFORMATION uint32
const(
    KeyBasicInformation // KEY_BASIC_INFORMATION KEY_INFORMATION_CLASS = 1  //col:27
    KeyNodeInformation // KEY_NODE_INFORMATION KEY_INFORMATION_CLASS = 2  //col:28
    KeyFullInformation // KEY_FULL_INFORMATION KEY_INFORMATION_CLASS = 3  //col:29
    KeyNameInformation // KEY_NAME_INFORMATION KEY_INFORMATION_CLASS = 4  //col:30
    KeyCachedInformation // KEY_CACHED_INFORMATION KEY_INFORMATION_CLASS = 5  //col:31
    KeyFlagsInformation // KEY_FLAGS_INFORMATION KEY_INFORMATION_CLASS = 6  //col:32
    KeyVirtualizationInformation // KEY_VIRTUALIZATION_INFORMATION KEY_INFORMATION_CLASS = 7  //col:33
    KeyHandleTagsInformation // KEY_HANDLE_TAGS_INFORMATION KEY_INFORMATION_CLASS = 8  //col:34
    KeyTrustInformation // KEY_TRUST_INFORMATION KEY_INFORMATION_CLASS = 9  //col:35
    KeyLayerInformation // KEY_LAYER_INFORMATION KEY_INFORMATION_CLASS = 10  //col:36
    MaxKeyInfoClass KEY_INFORMATION_CLASS = 11  //col:37
)


type     KeyWriteTimeInformation // KEY_WRITE_TIME_INFORMATION uint32
const(
    KeyWriteTimeInformation // KEY_WRITE_TIME_INFORMATION KEY_SET_INFORMATION_CLASS = 1  //col:140
    KeyWow64FlagsInformation // KEY_WOW64_FLAGS_INFORMATION KEY_SET_INFORMATION_CLASS = 2  //col:141
    KeyControlFlagsInformation // KEY_CONTROL_FLAGS_INFORMATION KEY_SET_INFORMATION_CLASS = 3  //col:142
    KeySetVirtualizationInformation // KEY_SET_VIRTUALIZATION_INFORMATION KEY_SET_INFORMATION_CLASS = 4  //col:143
    KeySetDebugInformation KEY_SET_INFORMATION_CLASS = 5  //col:144
    KeySetHandleTagsInformation // KEY_HANDLE_TAGS_INFORMATION KEY_SET_INFORMATION_CLASS = 6  //col:145
    KeySetLayerInformation // KEY_SET_LAYER_INFORMATION KEY_SET_INFORMATION_CLASS = 7  //col:146
    MaxKeySetInfoClass KEY_SET_INFORMATION_CLASS = 8  //col:147
)


type     KeyValueBasicInformation // KEY_VALUE_BASIC_INFORMATION uint32
const(
    KeyValueBasicInformation // KEY_VALUE_BASIC_INFORMATION KEY_VALUE_INFORMATION_CLASS = 1  //col:189
    KeyValueFullInformation // KEY_VALUE_FULL_INFORMATION KEY_VALUE_INFORMATION_CLASS = 2  //col:190
    KeyValuePartialInformation // KEY_VALUE_PARTIAL_INFORMATION KEY_VALUE_INFORMATION_CLASS = 3  //col:191
    KeyValueFullInformationAlign64 KEY_VALUE_INFORMATION_CLASS = 4  //col:192
    KeyValuePartialInformationAlign64  // KEY_VALUE_PARTIAL_INFORMATION_ALIGN64 KEY_VALUE_INFORMATION_CLASS = 5  //col:193
    KeyValueLayerInformation // KEY_VALUE_LAYER_INFORMATION KEY_VALUE_INFORMATION_CLASS = 6  //col:194
    MaxKeyValueInfoClass KEY_VALUE_INFORMATION_CLASS = 7  //col:195
)


type     KeyLoadTrustClassKey = 1 uint32
const(
    KeyLoadTrustClassKey  KEY_LOAD_ENTRY_TYPE =  1  //col:243
    KeyLoadEvent KEY_LOAD_ENTRY_TYPE = 2  //col:244
    KeyLoadToken KEY_LOAD_ENTRY_TYPE = 3  //col:245
)


type     KeyAdded uint32
const(
    KeyAdded REG_ACTION = 1  //col:269
    KeyRemoved REG_ACTION = 2  //col:270
    KeyModified REG_ACTION = 3  //col:271
)



type (
Ntregapi interface{
 * Attribution 4.0 International ()(ok bool)//col:38
}
)

func NewNtregapi() { return & ntregapi{} }

func (n *ntregapi) * Attribution 4.0 International ()(ok bool){//col:38
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTREGAPI_H
#define _NTREGAPI_H
#define REG_INIT_BOOT_SM 0x0000
#define REG_INIT_BOOT_SETUP 0x0001
#define REG_INIT_BOOT_ACCEPTED_BASE 0x0002
#define REG_INIT_BOOT_ACCEPTED_MAX REG_INIT_BOOT_ACCEPTED_BASE + 999
#define REG_MAX_KEY_VALUE_NAME_LENGTH 32767
#define REG_MAX_KEY_NAME_LENGTH 512
typedef enum _KEY_INFORMATION_CLASS
{
    MaxKeyInfoClass
} KEY_INFORMATION_CLASS;*/
return true
}



