package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntregapi.h.back

const (
	_NTREGAPI_H = //col:1
	REG_INIT_BOOT_SM = 0x0000        //col:2
REG_INIT_BOOT_SETUP = 0x0001         //col:3
REG_INIT_BOOT_ACCEPTED_BASE = 0x0002 //col:4
REG_INIT_BOOT_ACCEPTED_MAX = REG_INIT_BOOT_ACCEPTED_BASE + 999 //col:5
REG_MAX_KEY_VALUE_NAME_LENGTH = 32767                          //col:6
REG_MAX_KEY_NAME_LENGTH = 512 //col:7
REG_FLAG_VOLATILE = 0x0001    //col:8
REG_FLAG_LINK = 0x0002            //col:9
REG_KEY_DONT_VIRTUALIZE = 0x0002  //col:10
REG_KEY_DONT_SILENT_FAIL = 0x0004 //col:11
REG_KEY_RECURSE_FLAG = 0x0008 //col:12
)

const (
	KeyBasicInformation          = 1  //col:3
	KeyNodeInformation           = 2  //col:4
	KeyFullInformation           = 3  //col:5
	KeyNameInformation           = 4  //col:6
	KeyCachedInformation         = 5  //col:7
	KeyFlagsInformation          = 6  //col:8
	KeyVirtualizationInformation = 7  //col:9
	KeyHandleTagsInformation     = 8  //col:10
	KeyTrustInformation          = 9  //col:11
	KeyLayerInformation          = 10 //col:12
	MaxKeyInfoClass              = 11 //col:13
)

const (
	KeyWriteTimeInformation         = 1 //col:17
	KeyWow64FlagsInformation        = 2 //col:18
	KeyControlFlagsInformation      = 3 //col:19
	KeySetVirtualizationInformation = 4 //col:20
	KeySetDebugInformation          = 5 //col:21
	KeySetHandleTagsInformation     = 6 //col:22
	KeySetLayerInformation          = 7 //col:23
	MaxKeySetInfoClass              = 8 //col:24
)

const (
	KeyValueBasicInformation          = 1 //col:28
	KeyValueFullInformation           = 2 //col:29
	KeyValuePartialInformation        = 3 //col:30
	KeyValueFullInformationAlign64    = 4 //col:31
	KeyValuePartialInformationAlign64 = 5 //col:32
	KeyValueLayerInformation          = 6 //col:33
	MaxKeyValueInfoClass              = 7 //col:34
)

const (
	KeyLoadTrustClassKey = 1 //col:38
	KeyLoadEvent         = 2 //col:39
	KeyLoadToken         = 3 //col:40
)

const (
	KeyAdded    = 1 //col:44
	KeyRemoved  = 2 //col:45
	KeyModified = 3 //col:46
)

type KEY_BASIC_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:3
	TitleIndex    uint32        //col:4
	NameLength    uint32        //col:5
	Name          [1]WCHAR      //col:6
}

type KEY_NODE_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:10
	TitleIndex    uint32        //col:11
	ClassOffset   uint32        //col:12
	ClassLength   uint32        //col:13
	NameLength    uint32        //col:14
	Name          [1]WCHAR      //col:15
}

type KEY_FULL_INFORMATION struct {
	LastWriteTime   LARGE_INTEGER //col:19
	TitleIndex      uint32        //col:20
	ClassOffset     uint32        //col:21
	ClassLength     uint32        //col:22
	SubKeys         uint32        //col:23
	MaxNameLen      uint32        //col:24
	MaxClassLen     uint32        //col:25
	Values          uint32        //col:26
	MaxValueNameLen uint32        //col:27
	MaxValueDataLen uint32        //col:28
	Class           [1]WCHAR      //col:29
}

type KEY_NAME_INFORMATION struct {
	NameLength uint32   //col:33
	Name       [1]WCHAR //col:34
}

type KEY_CACHED_INFORMATION struct {
	LastWriteTime   LARGE_INTEGER //col:38
	TitleIndex      uint32        //col:39
	SubKeys         uint32        //col:40
	MaxNameLen      uint32        //col:41
	Values          uint32        //col:42
	MaxValueNameLen uint32        //col:43
	MaxValueDataLen uint32        //col:44
	NameLength      uint32        //col:45
	Name            [1]WCHAR      //col:46
}

type KEY_FLAGS_INFORMATION struct {
	Wow64Flags   uint32 //col:50
	KeyFlags     uint32 //col:51
	ControlFlags uint32 //col:52
}

type KEY_VIRTUALIZATION_INFORMATION struct {
	VirtualizationCandidate uint32 //col:56
	VirtualizationEnabled   uint32 //col:57
	VirtualTarget           uint32 //col:58
	VirtualStore            uint32 //col:59
	VirtualSource           uint32 //col:60
	Reserved                uint32 //col:61
}

type KEY_TRUST_INFORMATION struct {
	TrustedKey uint32 //col:65
	Reserved   uint32 //col:66
}

type KEY_LAYER_INFORMATION struct {
	IsTombstone      uint32 //col:70
	IsSupersedeLocal uint32 //col:71
	IsSupersedeTree  uint32 //col:72
	ClassIsInherited uint32 //col:73
	Reserved         uint32 //col:74
}

type KEY_WRITE_TIME_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:78
}

type KEY_WOW64_FLAGS_INFORMATION struct {
	UserFlags uint32 //col:82
}

type KEY_HANDLE_TAGS_INFORMATION struct {
	HandleTags uint32 //col:86
}

type KEY_SET_LAYER_INFORMATION struct {
	IsTombstone      uint32 //col:90
	IsSupersedeLocal uint32 //col:91
	IsSupersedeTree  uint32 //col:92
	ClassIsInherited uint32 //col:93
	Reserved         uint32 //col:94
}

type KEY_CONTROL_FLAGS_INFORMATION struct {
	ControlFlags uint32 //col:98
}

type KEY_SET_VIRTUALIZATION_INFORMATION struct {
	VirtualTarget uint32 //col:102
	VirtualStore  uint32 //col:103
	VirtualSource uint32 //col:104
	Reserved      uint32 //col:105
}

type KEY_VALUE_BASIC_INFORMATION struct {
	TitleIndex uint32   //col:109
	Type       uint32   //col:110
	NameLength uint32   //col:111
	Name       [1]WCHAR //col:112
}

type KEY_VALUE_FULL_INFORMATION struct {
	TitleIndex uint32   //col:116
	Type       uint32   //col:117
	DataOffset uint32   //col:118
	DataLength uint32   //col:119
	NameLength uint32   //col:120
	Name       [1]WCHAR //col:121
}

type KEY_VALUE_PARTIAL_INFORMATION struct {
	TitleIndex uint32   //col:125
	Type       uint32   //col:126
	DataLength uint32   //col:127
	Data       [1]uint8 //col:128
}

type KEY_VALUE_PARTIAL_INFORMATION_ALIGN64 struct {
	Type       uint32   //col:132
	DataLength uint32   //col:133
	Data       [1]uint8 //col:134
}

type KEY_VALUE_LAYER_INFORMATION struct {
	IsTombstone uint32 //col:138
	Reserved    uint32 //col:139
}

type KEY_LOAD_ENTRY struct {
	EntryType KEY_LOAD_ENTRY_TYPE //col:143
	Union     union               //col:144
	Handle    HANDLE              //col:146
	Value     ULONG_PTR           //col:147
}

type KEY_VALUE_ENTRY struct {
	ValueName  PUNICODE_STRING //col:152
	DataLength uint32          //col:153
	DataOffset uint32          //col:154
	Type       uint32          //col:155
}

type REG_NOTIFY_INFORMATION struct {
	NextEntryOffset uint32     //col:159
	Action          REG_ACTION //col:160
	KeyLength       uint32     //col:161
	Key             [1]WCHAR   //col:162
}

type KEY_PID_ARRAY struct {
	ProcessId HANDLE         //col:166
	KeyName   UNICODE_STRING //col:167
}

type KEY_OPEN_SUBKEYS_INFORMATION struct {
	Count    uint32           //col:171
	KeyArray [1]KEY_PID_ARRAY //col:172
}
