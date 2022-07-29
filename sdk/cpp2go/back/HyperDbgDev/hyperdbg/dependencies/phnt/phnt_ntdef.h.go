package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\phnt_ntdef.h.back

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

type     NotificationEvent uint32
const(
    NotificationEvent EVENT_TYPE = 1  //col:100
    SynchronizationEvent EVENT_TYPE = 2  //col:101
)


type     NotificationTimer uint32
const(
    NotificationTimer TIMER_TYPE = 1  //col:106
    SynchronizationTimer TIMER_TYPE = 2  //col:107
)


type     WaitAll uint32
const(
    WaitAll WAIT_TYPE = 1  //col:112
    WaitAny WAIT_TYPE = 2  //col:113
    WaitNotification WAIT_TYPE = 3  //col:114
)


type     NtProductWinNt = 1 uint32
const(
    NtProductWinNt  NT_PRODUCT_TYPE =  1  //col:269
    NtProductLanManNt NT_PRODUCT_TYPE = 2  //col:270
    NtProductServer NT_PRODUCT_TYPE = 3  //col:271
)


type     SmallBusiness uint32
const(
    SmallBusiness SUITE_TYPE = 1  //col:276
    Enterprise SUITE_TYPE = 2  //col:277
    BackOffice SUITE_TYPE = 3  //col:278
    CommunicationServer SUITE_TYPE = 4  //col:279
    TerminalServer SUITE_TYPE = 5  //col:280
    SmallBusinessRestricted SUITE_TYPE = 6  //col:281
    EmbeddedNT SUITE_TYPE = 7  //col:282
    DataCenter SUITE_TYPE = 8  //col:283
    SingleUserTS SUITE_TYPE = 9  //col:284
    Personal SUITE_TYPE = 10  //col:285
    Blade SUITE_TYPE = 11  //col:286
    EmbeddedRestricted SUITE_TYPE = 12  //col:287
    SecurityAppliance SUITE_TYPE = 13  //col:288
    StorageServer SUITE_TYPE = 14  //col:289
    ComputeServer SUITE_TYPE = 15  //col:290
    WHServer SUITE_TYPE = 16  //col:291
    PhoneNT SUITE_TYPE = 17  //col:292
    MaxSuiteType SUITE_TYPE = 18  //col:293
)



type (
PhntNtdef interface{
 * Attribution 4.0 International ()(ok bool)//col:34
typedef struct DECLSPEC_ALIGN()(ok bool)//col:41
typedef _Return_type_success_()(ok bool)//col:72
#define NT_SUCCESS()(ok bool)//col:102
    _Field_size_bytes_part_opt_()(ok bool)//col:124
    _Field_size_bytes_part_()(ok bool)//col:138
#define RTL_CONSTANT_STRING()(ok bool)//col:165
#define RTL_BALANCED_NODE_GET_PARENT_POINTER()(ok bool)//col:175
#define InitializeObjectAttributes()(ok bool)//col:249
}
)

func NewPhntNtdef() { return & phntNtdef{} }

func (p *phntNtdef) * Attribution 4.0 International ()(ok bool){//col:34
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _PHNT_NTDEF_H
#define _PHNT_NTDEF_H
#ifndef _NTDEF_
#define _NTDEF_
#ifndef NOTHING
#define NOTHING
#endif
typedef struct _QUAD
{
    union
    {
        __int64 UseThisFieldToCopy;
        double DoNotUseThisField;
    };
} QUAD, *PQUAD;*/
return true
}

func (p *phntNtdef)typedef struct DECLSPEC_ALIGN()(ok bool){//col:41
/*typedef struct DECLSPEC_ALIGN(MEMORY_ALLOCATION_ALIGNMENT) _QUAD_PTR
{
    ULONG_PTR DoNotUseThisField1;
    ULONG_PTR DoNotUseThisField2;
} QUAD_PTR, *PQUAD_PTR;*/
return true
}

func (p *phntNtdef)typedef _Return_type_success_()(ok bool){//col:72
/*typedef _Return_type_success_(return >= 0) LONG NTSTATUS;
typedef NTSTATUS *PNTSTATUS;
typedef char CCHAR;
typedef short CSHORT;
typedef ULONG CLONG;
typedef CCHAR *PCCHAR;
typedef CSHORT *PCSHORT;
typedef CLONG *PCLONG;
typedef PCSTR PCSZ;
typedef UCHAR KIRQL, *PKIRQL;
typedef LONG KPRIORITY, *PKPRIORITY;
typedef USHORT RTL_ATOM, *PRTL_ATOM;
typedef LARGE_INTEGER PHYSICAL_ADDRESS, *PPHYSICAL_ADDRESS;
typedef struct _LARGE_INTEGER_128
{
    LONGLONG QuadPart[2];
} LARGE_INTEGER_128, *PLARGE_INTEGER_128;*/
return true
}

func (p *phntNtdef)#define NT_SUCCESS()(ok bool){//col:102
/*#define NT_SUCCESS(Status) (((NTSTATUS)(Status)) >= 0)
#define NT_INFORMATION(Status) ((((ULONG)(Status)) >> 30) == 1)
#define NT_WARNING(Status) ((((ULONG)(Status)) >> 30) == 2)
#define NT_ERROR(Status) ((((ULONG)(Status)) >> 30) == 3)
#define NT_FACILITY_MASK 0xfff
#define NT_FACILITY_SHIFT 16
#define NT_FACILITY(Status) ((((ULONG)(Status)) >> NT_FACILITY_SHIFT) & NT_FACILITY_MASK)
#define NT_NTWIN32(Status) (NT_FACILITY(Status) == FACILITY_NTWIN32)
#define WIN32_FROM_NTSTATUS(Status) (((ULONG)(Status)) & 0xffff)
#ifndef _WIN64
#define FASTCALL __fastcall
#else
#define FASTCALL
#endif
typedef enum _EVENT_TYPE
{
    NotificationEvent,
    SynchronizationEvent
} EVENT_TYPE;*/
return true
}

func (p *phntNtdef)    _Field_size_bytes_part_opt_()(ok bool){//col:124
/*    _Field_size_bytes_part_opt_(MaximumLength, Length) PCHAR Buffer;
} STRING, *PSTRING, ANSI_STRING, *PANSI_STRING, OEM_STRING, *POEM_STRING;*/
return true
}

func (p *phntNtdef)    _Field_size_bytes_part_()(ok bool){//col:138
/*    _Field_size_bytes_part_(MaximumLength, Length) PWCH Buffer;
} UNICODE_STRING, *PUNICODE_STRING;*/
return true
}

func (p *phntNtdef)#define RTL_CONSTANT_STRING()(ok bool){//col:165
/*#define RTL_CONSTANT_STRING(s) { sizeof(s) - sizeof((s)[0]), sizeof(s), s }
#define RTL_BALANCED_NODE_RESERVED_PARENT_MASK 3
typedef struct _RTL_BALANCED_NODE
{
    union
    {
        struct _RTL_BALANCED_NODE *Children[2];
        struct
        {
            struct _RTL_BALANCED_NODE *Left;
            struct _RTL_BALANCED_NODE *Right;
        };
    };
    union
    {
        UCHAR Red : 1;
        UCHAR Balance : 2;
        ULONG_PTR ParentValue;
    };
} RTL_BALANCED_NODE, *PRTL_BALANCED_NODE;*/
return true
}

func (p *phntNtdef)#define RTL_BALANCED_NODE_GET_PARENT_POINTER()(ok bool){//col:175
/*#define RTL_BALANCED_NODE_GET_PARENT_POINTER(Node) \
    ((PRTL_BALANCED_NODE)((Node)->ParentValue & ~RTL_BALANCED_NODE_RESERVED_PARENT_MASK))
typedef struct _SINGLE_LIST_ENTRY32
{
    ULONG Next;
} SINGLE_LIST_ENTRY32, *PSINGLE_LIST_ENTRY32;*/
return true
}

func (p *phntNtdef)#define InitializeObjectAttributes()(ok bool){//col:249
/*#define InitializeObjectAttributes(p, n, a, r, s) { \
    (p)->Length = sizeof(OBJECT_ATTRIBUTES); \
    (p)->RootDirectory = r; \
    (p)->Attributes = a; \
    (p)->ObjectName = n; \
    (p)->SecurityDescriptor = s; \
    (p)->SecurityQualityOfService = NULL; \
    }
#define RTL_CONSTANT_OBJECT_ATTRIBUTES(n, a) { sizeof(OBJECT_ATTRIBUTES), NULL, n, a, NULL, NULL }
#define RTL_INIT_OBJECT_ATTRIBUTES(n, a) RTL_CONSTANT_OBJECT_ATTRIBUTES(n, a)
#define OBJ_NAME_PATH_SEPARATOR ((WCHAR)L'\\')
typedef struct _OBJECT_ATTRIBUTES64
{
    ULONG Length;
    ULONG64 RootDirectory;
    ULONG64 ObjectName;
    ULONG Attributes;
    ULONG64 SecurityDescriptor;
    ULONG64 SecurityQualityOfService;
} OBJECT_ATTRIBUTES64, *POBJECT_ATTRIBUTES64;*/
return true
}



