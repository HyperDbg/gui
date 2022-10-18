package phnt

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntregapi.h.back

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

type _KEY_BASIC_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:9
	TitleIndex    uint32        //col:10
	NameLength    uint32        //col:11
	Name          [1]WCHAR      //col:12
}

type _KEY_NODE_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:18
	TitleIndex    uint32        //col:19
	ClassOffset   uint32        //col:20
	ClassLength   uint32        //col:21
	NameLength    uint32        //col:22
	Name          [1]WCHAR      //col:23
}

type _KEY_FULL_INFORMATION struct {
	LastWriteTime   LARGE_INTEGER //col:32
	TitleIndex      uint32        //col:33
	ClassOffset     uint32        //col:34
	ClassLength     uint32        //col:35
	SubKeys         uint32        //col:36
	MaxNameLen      uint32        //col:37
	MaxClassLen     uint32        //col:38
	Values          uint32        //col:39
	MaxValueNameLen uint32        //col:40
	MaxValueDataLen uint32        //col:41
	Class           [1]WCHAR      //col:42
}

type _KEY_NAME_INFORMATION struct {
	NameLength uint32   //col:37
	Name       [1]WCHAR //col:38
}

type _KEY_CACHED_INFORMATION struct {
	LastWriteTime   LARGE_INTEGER //col:49
	TitleIndex      uint32        //col:50
	SubKeys         uint32        //col:51
	MaxNameLen      uint32        //col:52
	Values          uint32        //col:53
	MaxValueNameLen uint32        //col:54
	MaxValueDataLen uint32        //col:55
	NameLength      uint32        //col:56
	Name            [1]WCHAR      //col:57
}

type _KEY_FLAGS_INFORMATION struct {
	Wow64Flags   uint32 //col:55
	KeyFlags     uint32 //col:56
	ControlFlags uint32 //col:57
}

type _KEY_VIRTUALIZATION_INFORMATION struct {
	VirtualizationCandidate uint32 //col:64
	VirtualizationEnabled   uint32 //col:65
	VirtualTarget           uint32 //col:66
	VirtualStore            uint32 //col:67
	VirtualSource           uint32 //col:68
	Reserved                uint32 //col:69
}

type _KEY_TRUST_INFORMATION struct {
	TrustedKey uint32 //col:69
	Reserved   uint32 //col:70
}

type _KEY_LAYER_INFORMATION struct {
	IsTombstone      uint32 //col:77
	IsSupersedeLocal uint32 //col:78
	IsSupersedeTree  uint32 //col:79
	ClassIsInherited uint32 //col:80
	Reserved         uint32 //col:81
}

type _KEY_WRITE_TIME_INFORMATION struct {
	LastWriteTime LARGE_INTEGER //col:81
}

type _KEY_WOW64_FLAGS_INFORMATION struct {
	UserFlags uint32 //col:85
}

type _KEY_HANDLE_TAGS_INFORMATION struct {
	HandleTags uint32 //col:89
}

type _KEY_SET_LAYER_INFORMATION struct {
	IsTombstone      uint32 //col:97
	IsSupersedeLocal uint32 //col:98
	IsSupersedeTree  uint32 //col:99
	ClassIsInherited uint32 //col:100
	Reserved         uint32 //col:101
}

type _KEY_CONTROL_FLAGS_INFORMATION struct {
	ControlFlags uint32 //col:101
}

type _KEY_SET_VIRTUALIZATION_INFORMATION struct {
	VirtualTarget uint32 //col:108
	VirtualStore  uint32 //col:109
	VirtualSource uint32 //col:110
	Reserved      uint32 //col:111
}

type _KEY_VALUE_BASIC_INFORMATION struct {
	TitleIndex uint32   //col:115
	Type       uint32   //col:116
	NameLength uint32   //col:117
	Name       [1]WCHAR //col:118
}

type _KEY_VALUE_FULL_INFORMATION struct {
	TitleIndex uint32   //col:124
	Type       uint32   //col:125
	DataOffset uint32   //col:126
	DataLength uint32   //col:127
	NameLength uint32   //col:128
	Name       [1]WCHAR //col:129
}

type _KEY_VALUE_PARTIAL_INFORMATION struct {
	TitleIndex uint32   //col:131
	Type       uint32   //col:132
	DataLength uint32   //col:133
	Data       [1]uint8 //col:134
}

type _KEY_VALUE_PARTIAL_INFORMATION_ALIGN64 struct {
	Type       uint32   //col:137
	DataLength uint32   //col:138
	Data       [1]uint8 //col:139
}

type _KEY_VALUE_LAYER_INFORMATION struct {
	IsTombstone uint32 //col:142
	Reserved    uint32 //col:143
}

type _KEY_LOAD_ENTRY struct {
	EntryType KEY_LOAD_ENTRY_TYPE //col:150
	Union     union               //col:151
	Handle    uintptr             //col:153
	Value     ULONG_PTR           //col:154
}

type _KEY_VALUE_ENTRY struct {
	ValueName  *uint32 //col:158
	DataLength uint32  //col:159
	DataOffset uint32  //col:160
	Type       uint32  //col:161
}

type _REG_NOTIFY_INFORMATION struct {
	NextEntryOffset uint32     //col:165
	Action          REG_ACTION //col:166
	KeyLength       uint32     //col:167
	Key             [1]WCHAR   //col:168
}

type _KEY_PID_ARRAY struct {
	ProcessId uintptr //col:170
	KeyName   *int16  //col:171
}

type _KEY_OPEN_SUBKEYS_INFORMATION struct {
	Count    uint32           //col:175
	KeyArray [1]KEY_PID_ARRAY //col:176
}

