package phnt


const(
_PHNT_NTDEF_H =  //col:13
_NTDEF_ =  //col:16
NOTHING =  //col:22
NT_SUCCESS(Status) = (((NTSTATUS)(Status)) >= 0) //col:76
NT_INFORMATION(Status) = ((((ULONG)(Status)) >> 30) == 1) //col:77
NT_WARNING(Status) = ((((ULONG)(Status)) >> 30) == 2) //col:78
NT_ERROR(Status) = ((((ULONG)(Status)) >> 30) == 3) //col:79
NT_FACILITY_MASK = 0xfff //col:81
NT_FACILITY_SHIFT = 16 //col:82
NT_FACILITY(Status) = ((((ULONG)(Status)) >> NT_FACILITY_SHIFT) & NT_FACILITY_MASK) //col:83
NT_NTWIN32(Status) = (NT_FACILITY(Status) == FACILITY_NTWIN32) //col:85
WIN32_FROM_NTSTATUS(Status) = (((ULONG)(Status)) & 0xffff) //col:86
FASTCALL = __fastcall //col:91
FASTCALL =  //col:93
RTL_CONSTANT_STRING(s) = { sizeof(s) - sizeof((s)[0]), sizeof(s), s } //col:142
RTL_BALANCED_NODE_RESERVED_PARENT_MASK = 3 //col:146
RTL_BALANCED_NODE_GET_PARENT_POINTER(Node) = ((PRTL_BALANCED_NODE)((Node)->ParentValue & ~RTL_BALANCED_NODE_RESERVED_PARENT_MASK)) //col:167
OBJ_PROTECT_CLOSE = 0x00000001 //col:199
OBJ_INHERIT = 0x00000002 //col:200
OBJ_AUDIT_OBJECT_CLOSE = 0x00000004 //col:201
OBJ_PERMANENT = 0x00000010 //col:202
OBJ_EXCLUSIVE = 0x00000020 //col:203
OBJ_CASE_INSENSITIVE = 0x00000040 //col:204
OBJ_OPENIF = 0x00000080 //col:205
OBJ_OPENLINK = 0x00000100 //col:206
OBJ_KERNEL_HANDLE = 0x00000200 //col:207
OBJ_FORCE_ACCESS_CHECK = 0x00000400 //col:208
OBJ_IGNORE_IMPERSONATED_DEVICEMAP = 0x00000800 //col:209
OBJ_DONT_REPARSE = 0x00001000 //col:210
OBJ_VALID_ATTRIBUTES = 0x00001ff2 //col:211
InitializeObjectAttributes(p, n, a, r, s) { = (p)->Length = sizeof(OBJECT_ATTRIBUTES); (p)->RootDirectory = r; (p)->Attributes = a; (p)->ObjectName = n; (p)->SecurityDescriptor = s; (p)->SecurityQualityOfService = NULL; } //col:225
RTL_CONSTANT_OBJECT_ATTRIBUTES(n, = a) { sizeof(OBJECT_ATTRIBUTES), NULL, n, a, NULL, NULL } //col:234
RTL_INIT_OBJECT_ATTRIBUTES(n, = a) RTL_CONSTANT_OBJECT_ATTRIBUTES(n, a) //col:235
OBJ_NAME_PATH_SEPARATOR ((WCHAR)L'') =  //col:237
FlagOn(_F, = _SF) ((_F) & (_SF)) //col:329
BooleanFlagOn(F, = SF) ((BOOLEAN)(((F) & (SF)) != 0)) //col:332
SetFlag(_F, = _SF) ((_F) |= (_SF)) //col:335
ClearFlag(_F, = _SF) ((_F) &= ~(_SF)) //col:338
)

const(
    NotificationEvent = 1  //col:100
    SynchronizationEvent = 2  //col:101
)


const(
    NotificationTimer = 1  //col:106
    SynchronizationTimer = 2  //col:107
)


const(
    WaitAll = 1  //col:112
    WaitAny = 2  //col:113
    WaitNotification = 3  //col:114
)


const(
    NtProductWinNt  =  1  //col:269
    NtProductLanManNt = 2  //col:270
    NtProductServer = 3  //col:271
)


const(
    SmallBusiness = 1  //col:276
    Enterprise = 2  //col:277
    BackOffice = 3  //col:278
    CommunicationServer = 4  //col:279
    TerminalServer = 5  //col:280
    SmallBusinessRestricted = 6  //col:281
    EmbeddedNT = 7  //col:282
    DataCenter = 8  //col:283
    SingleUserTS = 9  //col:284
    Personal = 10  //col:285
    Blade = 11  //col:286
    EmbeddedRestricted = 12  //col:287
    SecurityAppliance = 13  //col:288
    StorageServer = 14  //col:289
    ComputeServer = 15  //col:290
    WHServer = 16  //col:291
    PhoneNT = 17  //col:292
    MaxSuiteType = 18  //col:293
)



type (
PhntNtdef interface{
typedef struct DECLSPEC_ALIGN()(ok bool)//col:41
typedef _Return_type_success_()(ok bool)//col:72
    _Field_size_bytes_part_opt_()(ok bool)//col:124
    _Field_size_bytes_part_()(ok bool)//col:138
    ()(ok bool)//col:175
    ()(ok bool)//col:249
}






)

func NewPhntNtdef() { return & phntNtdef{} }

func (p *phntNtdef)typedef struct DECLSPEC_ALIGN()(ok bool){//col:41





return true
}







func (p *phntNtdef)typedef _Return_type_success_()(ok bool){//col:72

















return true
}







func (p *phntNtdef)    _Field_size_bytes_part_opt_()(ok bool){//col:124


return true
}







func (p *phntNtdef)    _Field_size_bytes_part_()(ok bool){//col:138


return true
}







func (p *phntNtdef)    ()(ok bool){//col:175





return true
}







func (p *phntNtdef)    ()(ok bool){//col:249



















return true
}









