#ifndef _NTBCD_H
#define _NTBCD_H
#ifndef PHNT_INLINE_BCD_GUIDS
GUID DECLSPEC_SELECTANY GUID_BAD_MEMORY_GROUP                    = {0x5189B25C, 0x5558, 0x4BF2, {0xBC, 0xA4, 0x28, 0x9B, 0x11, 0xBD, 0x29, 0xE2}};
GUID DECLSPEC_SELECTANY GUID_BOOT_LOADER_SETTINGS_GROUP          = {0x6EFB52BF, 0x1766, 0x41DB, {0xA6, 0xB3, 0x0E, 0xE5, 0xEF, 0xF7, 0x2B, 0xD7}};
GUID DECLSPEC_SELECTANY GUID_CURRENT_BOOT_ENTRY                  = {0xFA926493, 0x6F1C, 0x4193, {0xA4, 0x14, 0x58, 0xF0, 0xB2, 0x45, 0x6D, 0x1E}};
GUID DECLSPEC_SELECTANY GUID_DEBUGGER_SETTINGS_GROUP             = {0x4636856E, 0x540F, 0x4170, {0xA1, 0x30, 0xA8, 0x47, 0x76, 0xF4, 0xC6, 0x54}};
GUID DECLSPEC_SELECTANY GUID_DEFAULT_BOOT_ENTRY                  = {0x1CAE1EB7, 0xA0DF, 0x4D4D, {0x98, 0x51, 0x48, 0x60, 0xE3, 0x4E, 0xF5, 0x35}};
GUID DECLSPEC_SELECTANY GUID_EMS_SETTINGS_GROUP                  = {0x0CE4991B, 0xE6B3, 0x4B16, {0xB2, 0x3C, 0x5E, 0x0D, 0x92, 0x50, 0xE5, 0xD9}};
GUID DECLSPEC_SELECTANY GUID_FIRMWARE_BOOTMGR                    = {0xA5A30FA2, 0x3D06, 0x4E9F, {0xB5, 0xF4, 0xA0, 0x1D, 0xF9, 0xD1, 0xFC, 0xBA}};
GUID DECLSPEC_SELECTANY GUID_GLOBAL_SETTINGS_GROUP               = {0x7EA2E1AC, 0x2E61, 0x4728, {0xAA, 0xA3, 0x89, 0x6D, 0x9D, 0x0A, 0x9F, 0x0E}};
GUID DECLSPEC_SELECTANY GUID_HYPERVISOR_SETTINGS_GROUP           = {0x7FF607E0, 0x4395, 0x11DB, {0xB0, 0xDE, 0x08, 0x00, 0x20, 0x0C, 0x9A, 0x66}};
GUID DECLSPEC_SELECTANY GUID_KERNEL_DEBUGGER_SETTINGS_GROUP      = {0x313E8EED, 0x7098, 0x4586, {0xA9, 0xBF, 0x30, 0x9C, 0x61, 0xF8, 0xD4, 0x49}};
GUID DECLSPEC_SELECTANY GUID_RESUME_LOADER_SETTINGS_GROUP        = {0x1AFA9C49, 0x16AB, 0x4A5C, {0x4A, 0x90, 0x21, 0x28, 0x02, 0xDA, 0x94, 0x60}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_BOOTMGR                     = {0x9DEA862C, 0x5CDD, 0x4E70, {0xAC, 0xC1, 0xF3, 0x2B, 0x34, 0x4D, 0x47, 0x95}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_LEGACY_NTLDR                = {0x466F5A88, 0x0AF2, 0x4F76, {0x90, 0x38, 0x09, 0x5B, 0x17, 0x0D, 0xC2, 0x1C}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_MEMORY_TESTER               = {0xB2721D73, 0x1DB4, 0x4C62, {0xBF, 0x78, 0xC5, 0x48, 0xA8, 0x80, 0x14, 0x2D}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_OS_TARGET_TEMPLATE_EFI      = {0xB012B84D, 0xC47C, 0x4ED5, {0xB7, 0x22, 0xC0, 0xC4, 0x21, 0x63, 0xE5, 0x69}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_OS_TARGET_TEMPLATE_PCAT     = {0xA1943BBC, 0xEA85, 0x487C, {0x97, 0xC7, 0xC9, 0xED, 0xE9, 0x08, 0xA3, 0x8A}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_RESUME_TARGET_TEMPLATE_EFI  = {0x0C334284, 0x9A41, 0x4DE1, {0x99, 0xB3, 0xA7, 0xE8, 0x7E, 0x8F, 0xF0, 0x7E}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_RESUME_TARGET_TEMPLATE_PCAT = {0x98B02A23, 0x0674, 0x4CE7, {0xBD, 0xAD, 0xE0, 0xA1, 0x5A, 0x8F, 0xF9, 0x7B}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_EFI                   = {0x7254A080, 0x1510, 0x4E85, {0xAC, 0x0F, 0xE7, 0xFB, 0x3D, 0x44, 0x47, 0x36}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_PCAT                  = {0xCBD971BF, 0xB7B8, 0x4885, {0x95, 0x1A, 0xFA, 0x03, 0x04, 0x4F, 0x5D, 0x71}};
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_RAMDISK_OPTIONS       = {0xAE5534E0, 0xA924, 0x466C, {0xB8, 0x36, 0x75, 0x85, 0x39, 0xA3, 0xEE, 0x3A}};
#else
NTSYSAPI GUID GUID_BAD_MEMORY_GROUP;               // {badmemory}
NTSYSAPI GUID GUID_BOOT_LOADER_SETTINGS_GROUP;     // {bootloadersettings}
NTSYSAPI GUID GUID_CURRENT_BOOT_ENTRY;             // {current}
NTSYSAPI GUID GUID_DEBUGGER_SETTINGS_GROUP;        // {eventsettings} {dbgsettings}
NTSYSAPI GUID GUID_DEFAULT_BOOT_ENTRY;             // {default}
NTSYSAPI GUID GUID_EMS_SETTINGS_GROUP;             // {emssettings}
NTSYSAPI GUID GUID_FIRMWARE_BOOTMGR;               // {fwbootmgr}
NTSYSAPI GUID GUID_GLOBAL_SETTINGS_GROUP;          // {globalsettings}
NTSYSAPI GUID GUID_HYPERVISOR_SETTINGS_GROUP;      // {hypervisorsettings}
NTSYSAPI GUID GUID_KERNEL_DEBUGGER_SETTINGS_GROUP; // {kerneldbgsettings}
NTSYSAPI GUID GUID_RESUME_LOADER_SETTINGS_GROUP;   // {resumeloadersettings}
NTSYSAPI GUID GUID_WINDOWS_BOOTMGR;                // {bootmgr}
NTSYSAPI GUID GUID_WINDOWS_LEGACY_NTLDR;           // {ntldr} {legacy}
NTSYSAPI GUID GUID_WINDOWS_MEMORY_TESTER;          // {memdiag}
NTSYSAPI GUID GUID_WINDOWS_OS_TARGET_TEMPLATE_EFI;
NTSYSAPI GUID GUID_WINDOWS_OS_TARGET_TEMPLATE_PCAT;
NTSYSAPI GUID GUID_WINDOWS_RESUME_TARGET_TEMPLATE_EFI;
NTSYSAPI GUID GUID_WINDOWS_RESUME_TARGET_TEMPLATE_PCAT;
NTSYSAPI GUID GUID_WINDOWS_SETUP_EFI;
NTSYSAPI GUID GUID_WINDOWS_SETUP_PCAT;
NTSYSAPI GUID GUID_WINDOWS_SETUP_RAMDISK_OPTIONS; // {ramdiskoptions}
#endif
typedef enum _BCD_MESSAGE_TYPE {
    BCD_MESSAGE_TYPE_NONE,
    BCD_MESSAGE_TYPE_TRACE,
    BCD_MESSAGE_TYPE_INFORMATION,
    BCD_MESSAGE_TYPE_WARNING,
    BCD_MESSAGE_TYPE_ERROR,
    BCD_MESSAGE_TYPE_MAXIMUM
} BCD_MESSAGE_TYPE;
typedef VOID(NTAPI * BCD_MESSAGE_CALLBACK)(
    _In_ BCD_MESSAGE_TYPE type,
    _In_ PWSTR            Message);
NTSYSAPI
NTSTATUS
NTAPI
BcdSetLogging(
    _In_ BCD_MESSAGE_TYPE     BcdLoggingLevel,
    _In_ BCD_MESSAGE_CALLBACK BcdMessageCallbackRoutine);
NTSYSAPI
VOID
    NTAPI
        BcdInitializeBcdSyncMutant(
            VOID);
NTSYSAPI
NTSTATUS
NTAPI
BcdGetSystemStorePath(
    _Out_ PWSTR * BcdSystemStorePath // RtlFreeHeap(RtlProcessHeap(), 0, BcdSystemStorePath);
);
NTSYSAPI
NTSTATUS
NTAPI
BcdSetSystemStoreDevice(
    _In_ UNICODE_STRING SystemPartition);
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenSystemStore(
    _Out_ PHANDLE BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenStoreFromFile(
    _In_ UNICODE_STRING BcdFilePath,
    _Out_ PHANDLE       BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdCreateStore(
    _In_ UNICODE_STRING BcdFilePath,
    _Out_ PHANDLE       BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdExportStore(
    _In_ UNICODE_STRING BcdFilePath);
NTSYSAPI
NTSTATUS
NTAPI
BcdExportStoreEx(
    _In_ HANDLE         BcdStoreHandle,
    _In_ ULONG          Flags,
    _In_ UNICODE_STRING BcdFilePath);
NTSYSAPI
NTSTATUS
NTAPI
BcdImportStore(
    _In_ UNICODE_STRING BcdFilePath);
typedef enum _BCD_IMPORT_FLAGS {
    BCD_IMPORT_NONE,
    BCD_IMPORT_DELETE_FIRMWARE_OBJECTS
} BCD_IMPORT_FLAGS;
NTSYSAPI
NTSTATUS
NTAPI
BcdImportStoreWithFlags(
    _In_ UNICODE_STRING   BcdFilePath,
    _In_ BCD_IMPORT_FLAGS BcdImportFlags);
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteObjectReferences(
    _In_ HANDLE BcdStoreHandle,
    _In_ PGUID  Identifier);
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteSystemStore(
    VOID);
typedef enum _BCD_OPEN_FLAGS {
    BCD_OPEN_NONE,
    BCD_OPEN_OPEN_STORE_OFFLINE,
    BCD_OPEN_SYNC_FIRMWARE_ENTRIES
} BCD_OPEN_FLAGS;
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenStore(
    _In_ UNICODE_STRING BcdFilePath,
    _In_ BCD_OPEN_FLAGS BcdOpenFlags,
    _Out_ PHANDLE       BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdCloseStore(
    _In_ HANDLE BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdFlushStore(
    _In_ HANDLE BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdForciblyUnloadStore(
    _In_ HANDLE BcdStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdMarkAsSystemStore(
    _In_ HANDLE BcdStoreHandle);
typedef enum _BCD_OBJECT_TYPE {
    BCD_OBJECT_TYPE_NONE,
    BCD_OBJECT_TYPE_APPLICATION,
    BCD_OBJECT_TYPE_INHERITED,
    BCD_OBJECT_TYPE_DEVICE
} BCD_OBJECT_TYPE;
typedef enum _BCD_APPLICATION_OBJECT_TYPE {
    BCD_APPLICATION_OBJECT_NONE                       = 0,
    BCD_APPLICATION_OBJECT_FIRMWARE_BOOT_MANAGER      = 1,
    BCD_APPLICATION_OBJECT_WINDOWS_BOOT_MANAGER       = 2,
    BCD_APPLICATION_OBJECT_WINDOWS_BOOT_LOADER        = 3,
    BCD_APPLICATION_OBJECT_WINDOWS_RESUME_APPLICATION = 4,
    BCD_APPLICATION_OBJECT_MEMORY_TESTER              = 5,
    BCD_APPLICATION_OBJECT_LEGACY_NTLDR               = 6,
    BCD_APPLICATION_OBJECT_LEGACY_SETUPLDR            = 7,
    BCD_APPLICATION_OBJECT_BOOT_SECTOR                = 8,
    BCD_APPLICATION_OBJECT_STARTUP_MODULE             = 9,
    BCD_APPLICATION_OBJECT_GENERIC_APPLICATION        = 10,
    BCD_APPLICATION_OBJECT_RESERVED                   = 0xFFFFF
} BCD_APPLICATION_OBJECT_TYPE;
typedef enum _BCD_APPLICATION_IMAGE_TYPE {
    BCD_APPLICATION_IMAGE_NONE,
    BCD_APPLICATION_IMAGE_FIRMWARE_APPLICATION,
    BCD_APPLICATION_IMAGE_BOOT_APPLICATION,
    BCD_APPLICATION_IMAGE_LEGACY_LOADER,
    BCD_APPLICATION_IMAGE_REALMODE_CODE
} BCD_APPLICATION_IMAGE_TYPE;
typedef enum _BCD_INHERITED_CLASS_TYPE {
    BCD_INHERITED_CLASS_NONE,
    BCD_INHERITED_CLASS_LIBRARY,
    BCD_INHERITED_CLASS_APPLICATION,
    BCD_INHERITED_CLASS_DEVICE
} BCD_INHERITED_CLASS_TYPE;
#define MAKE_BCD_OBJECT(ObjectType, ImageType, ApplicationType) \
    (((ULONG)(ObjectType) << 28) |                              \
     (((ULONG)(ImageType)&0xF) << 20) |                         \
     ((ULONG)(ApplicationType)&0xFFFFF))
#define MAKE_BCD_APPLICATION_OBJECT(ImageType, ApplicationType) \
    MAKE_BCD_OBJECT(BCD_OBJECT_TYPE_APPLICATION, (ULONG)(ImageType), (ULONG)(ApplicationType))
#define GET_BCD_OBJECT_TYPE(DataType) \
    ((BCD_OBJECT_TYPE)(((((ULONG)DataType)) >> 28) & 0xF))
#define GET_BCD_APPLICATION_IMAGE(DataType) \
    ((BCD_APPLICATION_IMAGE_TYPE)(((((ULONG)DataType)) >> 20) & 0xF))
#define GET_BCD_APPLICATION_OBJECT(DataType) \
    ((BCD_APPLICATION_OBJECT_TYPE)((((ULONG)DataType)) & 0xFFFFF))
#define BCD_OBJECT_OSLOADER_TYPE \
    MAKE_BCD_APPLICATION_OBJECT(BCD_APPLICATION_IMAGE_BOOT_APPLICATION, BCD_APPLICATION_OBJECT_WINDOWS_BOOT_LOADER)
typedef union _BCD_OBJECT_DATATYPE {
    ULONG PackedValue;
    union {
        struct
        {
            ULONG           Reserved : 28;
            BCD_OBJECT_TYPE ObjectType : 4;
        };
        struct
        {
            BCD_APPLICATION_OBJECT_TYPE ApplicationType : 20;
            BCD_APPLICATION_IMAGE_TYPE  ImageType : 4;
            ULONG                       Reserved : 4;
            BCD_OBJECT_TYPE             ObjectType : 4;
        } Application;
        struct
        {
            ULONG                    Value : 20;
            BCD_INHERITED_CLASS_TYPE Class : 4;
            ULONG                    Reserved : 4;
            BCD_OBJECT_TYPE          ObjectType : 4;
        } Inherit;
        struct
        {
            ULONG           Reserved : 28;
            BCD_OBJECT_TYPE ObjectType : 4;
        } Device;
    };
} BCD_OBJECT_DATATYPE, *PBCD_OBJECT_DATATYPE;
#define BCD_OBJECT_DESCRIPTION_VERSION 0x1
typedef struct _BCD_OBJECT_DESCRIPTION {
    ULONG Version; // BCD_OBJECT_DESCRIPTION_VERSION
    ULONG Type;    // BCD_OBJECT_DATATYPE
} BCD_OBJECT_DESCRIPTION, *PBCD_OBJECT_DESCRIPTION;
typedef struct _BCD_OBJECT {
    GUID                    Identifer;
    PBCD_OBJECT_DESCRIPTION Description;
} BCD_OBJECT, *PBCD_OBJECT;
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateObjects(
    _In_ HANDLE                               BcdStoreHandle,
    _In_ PBCD_OBJECT_DESCRIPTION              BcdEnumDescriptor,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer, // BCD_OBJECT[]
    _Inout_ PULONG                            BufferSize,
    _Out_ PULONG                              ObjectCount);
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenObject(
    _In_ HANDLE   BcdStoreHandle,
    _In_ PGUID    Identifier,
    _Out_ PHANDLE BcdObjectHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdCreateObject(
    _In_ HANDLE                  BcdStoreHandle,
    _In_ PGUID                   Identifier,
    _In_ PBCD_OBJECT_DESCRIPTION Description,
    _Out_ PHANDLE                BcdObjectHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteObject(
    _In_ HANDLE BcdObjectHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdCloseObject(
    _In_ HANDLE BcdObjectHandle);
typedef enum _BCD_COPY_FLAGS {
    BCD_COPY_NONE                              = 0x0,
    BCD_COPY_COPY_CREATE_NEW_OBJECT_IDENTIFIER = 0x1,
    BCD_COPY_COPY_DELETE_EXISTING_OBJECT       = 0x2,
    BCD_COPY_COPY_UNKNOWN_FIRMWARE_APPLICATION = 0x4,
    BCD_COPY_IGNORE_SETUP_TEMPLATE_ELEMENTS    = 0x8,
    BCD_COPY_RETAIN_ELEMENT_DATA               = 0x10,
    BCD_COPY_MIGRATE_ELEMENT_DATA              = 0x20
} BCD_COPY_FLAGS;
NTSYSAPI
NTSTATUS
NTAPI
BcdCopyObject(
    _In_ HANDLE         BcdStoreHandle,
    _In_ HANDLE         BcdObjectHandle,
    _In_ BCD_COPY_FLAGS BcdCopyFlags,
    _In_ HANDLE         TargetStoreHandle,
    _Out_ PHANDLE       TargetHandleOut);
NTSYSAPI
NTSTATUS
NTAPI
BcdCopyObjectEx(
    _In_ HANDLE         BcdStoreHandle,
    _In_ HANDLE         BcdObjectHandle,
    _In_ BCD_COPY_FLAGS BcdCopyFlags,
    _In_ HANDLE         TargetStoreHandle,
    _In_ PGUID          TargetObjectId,
    _Out_ PHANDLE       TargetHandleOut);
NTSYSAPI
NTSTATUS
NTAPI
BcdCopyObjects(
    _In_ HANDLE                 BcdStoreHandle,
    _In_ BCD_OBJECT_DESCRIPTION Characteristics,
    _In_ BCD_COPY_FLAGS         BcdCopyFlags,
    _In_ HANDLE                 TargetStoreHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdMigrateObjectElementValues(
    _In_ HANDLE TemplateObjectHandle,
    _In_ HANDLE SourceObjectHandle,
    _In_ HANDLE TargetObjectHandle);
NTSYSAPI
NTSTATUS
NTAPI
BcdQueryObject(
    _In_ HANDLE                  BcdObjectHandle,
    _In_ ULONG                   BcdVersion, // BCD_OBJECT_DESCRIPTION_VERSION
    _Out_ BCD_OBJECT_DESCRIPTION Description,
    _Out_ PGUID                  Identifier);
typedef enum _BCD_ELEMENT_DATATYPE_FORMAT {
    BCD_ELEMENT_DATATYPE_FORMAT_UNKNOWN,
    BCD_ELEMENT_DATATYPE_FORMAT_DEVICE,      // 0x01000000
    BCD_ELEMENT_DATATYPE_FORMAT_STRING,      // 0x02000000
    BCD_ELEMENT_DATATYPE_FORMAT_OBJECT,      // 0x03000000
    BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST,  // 0x04000000
    BCD_ELEMENT_DATATYPE_FORMAT_INTEGER,     // 0x05000000
    BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN,     // 0x06000000
    BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, // 0x07000000
    BCD_ELEMENT_DATATYPE_FORMAT_BINARY       // 0x08000000
} BCD_ELEMENT_DATATYPE_FORMAT;
typedef enum _BCD_ELEMENT_DATATYPE_CLASS {
    BCD_ELEMENT_DATATYPE_CLASS_NONE,
    BCD_ELEMENT_DATATYPE_CLASS_LIBRARY,
    BCD_ELEMENT_DATATYPE_CLASS_APPLICATION,
    BCD_ELEMENT_DATATYPE_CLASS_DEVICE,
    BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE,
    BCD_ELEMENT_DATATYPE_CLASS_OEM
} BCD_ELEMENT_DATATYPE_CLASS;
typedef enum _BCD_ELEMENT_DEVICE_TYPE {
    BCD_ELEMENT_DEVICE_TYPE_NONE,
    BCD_ELEMENT_DEVICE_TYPE_BOOT_DEVICE,
    BCD_ELEMENT_DEVICE_TYPE_PARTITION,
    BCD_ELEMENT_DEVICE_TYPE_FILE,
    BCD_ELEMENT_DEVICE_TYPE_RAMDISK,
    BCD_ELEMENT_DEVICE_TYPE_UNKNOWN,
    BCD_ELEMENT_DEVICE_TYPE_QUALIFIED_PARTITION,
    BCD_ELEMENT_DEVICE_TYPE_VMBUS,
    BCD_ELEMENT_DEVICE_TYPE_LOCATE_DEVICE,
    BCD_ELEMENT_DEVICE_TYPE_URI,
    BCD_ELEMENT_DEVICE_TYPE_COMPOSITE
} BCD_ELEMENT_DEVICE_TYPE;
#define MAKE_BCDE_DATA_TYPE(Class, Format, Subtype) \
    (((((ULONG)Class) & 0xF) << 28) | ((((ULONG)Format) & 0xF) << 24) | (((ULONG)Subtype) & 0x00FFFFFF))
#define GET_BCDE_DATA_CLASS(DataType) \
    ((BCD_ELEMENT_DATATYPE_CLASS)(((((ULONG)DataType)) >> 28) & 0xF))
#define GET_BCDE_DATA_FORMAT(DataType) \
    ((BCD_ELEMENT_DATATYPE_FORMAT)(((((ULONG)DataType)) >> 24) & 0xF))
#define GET_BCDE_DATA_SUBTYPE(DataType) \
    ((ULONG)((((ULONG)DataType)) & 0x00FFFFFF))
typedef union _BCD_ELEMENT_DATATYPE {
    ULONG PackedValue;
    struct
    {
        ULONG                       SubType : 24;
        BCD_ELEMENT_DATATYPE_FORMAT Format : 4;
        BCD_ELEMENT_DATATYPE_CLASS  Class : 4;
    };
} BCD_ELEMENT_DATATYPE, *PBCD_ELEMENT_DATATYPE;
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateElementTypes(
    _In_ HANDLE                               BcdObjectHandle,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer, // BCD_ELEMENT_DATATYPE[]
    _Inout_ PULONG                            BufferSize,
    _Out_ PULONG                              ElementCount);
typedef struct _BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION {
    ULONG PartitionStyle;
    ULONG Reserved;
    struct
    {
        union {
            ULONG   DiskSignature;
            ULONG64 PartitionOffset;
        } Mbr;
        union {
            GUID DiskSignature;
            GUID PartitionSignature;
        } Gpt;
    };
} BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION, *PBCD_ELEMENT_DEVICE_QUALIFIED_PARTITION;
typedef struct _BCD_ELEMENT_DEVICE {
    ULONG DeviceType;
    GUID  AdditionalOptions;
    struct
    {
        union {
            ULONG ParentOffset;
            WCHAR Path[ANYSIZE_ARRAY];
        } File;
        union {
            WCHAR Path[ANYSIZE_ARRAY];
        } Partition;
        union {
            ULONG Type;
            ULONG ParentOffset;
            ULONG ElementType;
            WCHAR Path[ANYSIZE_ARRAY];
        } Locate;
        union {
            GUID InterfaceInstance;
        } Vmbus;
        union {
            ULONG Data[ANYSIZE_ARRAY];
        } Unknown;
        BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION QualifiedPartition;
    };
} BCD_ELEMENT_DEVICE, *PBCD_ELEMENT_DEVICE;
typedef struct _BCD_ELEMENT_STRING {
    WCHAR Value[ANYSIZE_ARRAY];
} BCD_ELEMENT_STRING, *PBCD_ELEMENT_STRING;
typedef struct _BCD_ELEMENT_OBJECT {
    GUID Object;
} BCD_ELEMENT_OBJECT, *PBCD_ELEMENT_OBJECT;
typedef struct _BCD_ELEMENT_OBJECT_LIST {
    GUID ObjectList[ANYSIZE_ARRAY];
} BCD_ELEMENT_OBJECT_LIST, *PBCD_ELEMENT_OBJECT_LIST;
typedef struct _BCD_ELEMENT_INTEGER {
    ULONG64 Value;
} BCD_ELEMENT_INTEGER, *PBCD_ELEMENT_INTEGER;
typedef struct _BCD_ELEMENT_INTEGER_LIST {
    ULONG64 Value[ANYSIZE_ARRAY];
} BCD_ELEMENT_INTEGER_LIST, *PBCD_ELEMENT_INTEGER_LIST;
typedef struct _BCD_ELEMENT_BOOLEAN {
    BOOLEAN Value;
} BCD_ELEMENT_BOOLEAN, *PBCD_ELEMENT_BOOLEAN;
#define BCD_ELEMENT_DESCRIPTION_VERSION 0x1
typedef struct BCD_ELEMENT_DESCRIPTION {
    ULONG Version; // BCD_ELEMENT_DESCRIPTION_VERSION
    ULONG Type;
    ULONG DataSize;
} BCD_ELEMENT_DESCRIPTION, *PBCD_ELEMENT_DESCRIPTION;
typedef struct _BCD_ELEMENT {
    PBCD_ELEMENT_DESCRIPTION Description;
    PVOID                    Data;
} BCD_ELEMENT, *PBCD_ELEMENT;
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateElements(
    _In_ HANDLE                               BcdObjectHandle,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer, // BCD_ELEMENT[]
    _Inout_ PULONG                            BufferSize,
    _Out_ PULONG                              ElementCount);
typedef enum _BCD_FLAGS {
    BCD_FLAG_NONE                         = 0x0,
    BCD_FLAG_QUALIFIED_PARTITION          = 0x1,
    BCD_FLAG_NO_DEVICE_TRANSLATION        = 0x2,
    BCD_FLAG_ENUMERATE_INHERITED_OBJECTS  = 0x4,
    BCD_FLAG_ENUMERATE_DEVICE_OPTIONS     = 0x8,
    BCD_FLAG_OBSERVE_PRECEDENCE           = 0x10,
    BCD_FLAG_DISABLE_VHD_NT_TRANSLATION   = 0x20,
    BCD_FLAG_DISABLE_VHD_DEVICE_DETECTION = 0x40,
    BCD_FLAG_DISABLE_POLICY_CHECKS        = 0x80
} BCD_FLAGS;
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateElementsWithFlags(
    _In_ HANDLE                               BcdObjectHandle,
    _In_ BCD_FLAGS                            BcdFlags,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer, // BCD_ELEMENT[]
    _Inout_ PULONG                            BufferSize,
    _Out_ PULONG                              ElementCount);
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateAndUnpackElements(
    _In_ HANDLE                               BcdStoreHandle,
    _In_ HANDLE                               BcdObjectHandle,
    _In_ BCD_FLAGS                            BcdFlags,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer, // BCD_ELEMENT[]
    _Inout_ PULONG                            BufferSize,
    _Out_ PULONG                              ElementCount);
NTSYSAPI
NTSTATUS
NTAPI
BcdGetElementData(
    _In_ HANDLE                               BcdObjectHandle,
    _In_ ULONG                                BcdElement, // BCD_ELEMENT_DATATYPE
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer,
    _Inout_ PULONG                            BufferSize);
NTSYSAPI
NTSTATUS
NTAPI
BcdGetElementDataWithFlags(
    _In_ HANDLE                               BcdObjectHandle,
    _In_ ULONG                                BcdElement, // BCD_ELEMENT_DATATYPE
    _In_ BCD_FLAGS                            BcdFlags,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer,
    _Inout_ PULONG                            BufferSize);
NTSYSAPI
NTSTATUS
NTAPI
BcdSetElementData(
    _In_ HANDLE                            BcdObjectHandle,
    _In_ ULONG                             BcdElement, // BCD_ELEMENT_DATATYPE
    _In_reads_bytes_opt_(BufferSize) PVOID Buffer,
    _In_ ULONG                             BufferSize);
NTSYSAPI
NTSTATUS
NTAPI
BcdSetElementDataWithFlags(
    _In_ HANDLE                            BcdObjectHandle,
    _In_ ULONG                             BcdElement, // BCD_ELEMENT_DATATYPE
    _In_ BCD_FLAGS                         BcdFlags,
    _In_reads_bytes_opt_(BufferSize) PVOID Buffer,
    _In_ ULONG                             BufferSize);
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteElement(
    _In_ HANDLE BcdObjectHandle,
    _In_ ULONG  BcdElement // BCD_ELEMENT_DATATYPE
);
typedef enum _BcdBootMgrElementTypes {
    BcdBootMgrObjectList_DisplayOrder           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 1),
    BcdBootMgrObjectList_BootSequence           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 2),
    BcdBootMgrObject_DefaultObject              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 3),
    BcdBootMgrInteger_Timeout                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 4),
    BcdBootMgrBoolean_AttemptResume             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 5),
    BcdBootMgrObject_ResumeObject               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 6),
    BcdBootMgrObjectList_StartupSequence        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 7),
    BcdBootMgrObjectList_ToolsDisplayOrder      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 16),
    BcdBootMgrBoolean_DisplayBootMenu           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 32),
    BcdBootMgrBoolean_NoErrorDisplay            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 33),
    BcdBootMgrDevice_BcdDevice                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 34),
    BcdBootMgrString_BcdFilePath                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 35),
    BcdBootMgrBoolean_HormEnabled               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 36),
    BcdBootMgrBoolean_HiberRoot                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 37),
    BcdBootMgrString_PasswordOverride           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 38),
    BcdBootMgrString_PinpassPhraseOverride      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 39),
    BcdBootMgrBoolean_ProcessCustomActionsFirst = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 40),
    BcdBootMgrIntegerList_CustomActionsList     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 48),
    BcdBootMgrBoolean_PersistBootSequence       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 49),
    BcdBootMgrBoolean_SkipStartupSequence       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 50),
} BcdBootMgrElementTypes;
typedef enum _BcdLibrary_FirstMegabytePolicy {
    FirstMegabytePolicyUseNone,
    FirstMegabytePolicyUseAll,
    FirstMegabytePolicyUsePrivate
} BcdLibrary_FirstMegabytePolicy;
typedef enum _BcdLibrary_DebuggerType {
    DebuggerSerial = 0,
    Debugger1394   = 1,
    DebuggerUsb    = 2,
    DebuggerNet    = 3,
    DebuggerLocal  = 4
} BcdLibrary_DebuggerType;
typedef enum _BcdLibrary_DebuggerStartPolicy {
    DebuggerStartActive,
    DebuggerStartAutoEnable,
    DebuggerStartDisable
} BcdLibrary_DebuggerStartPolicy;
typedef enum _BcdLibrary_ConfigAccessPolicy {
    ConfigAccessPolicyDefault,
    ConfigAccessPolicyDisallowMmConfig
} BcdLibrary_ConfigAccessPolicy;
typedef enum _BcdLibrary_UxDisplayMessageType {
    DisplayMessageTypeDefault             = 0,
    DisplayMessageTypeResume              = 1,
    DisplayMessageTypeHyperV              = 2,
    DisplayMessageTypeRecovery            = 3,
    DisplayMessageTypeStartupRepair       = 4,
    DisplayMessageTypeSystemImageRecovery = 5,
    DisplayMessageTypeCommandPrompt       = 6,
    DisplayMessageTypeSystemRestore       = 7,
    DisplayMessageTypePushButtonReset     = 8,
} BcdLibrary_UxDisplayMessageType;
typedef enum BcdLibrary_SafeBoot {
    SafemodeMinimal  = 0,
    SafemodeNetwork  = 1,
    SafemodeDsRepair = 2
} BcdLibrary_SafeBoot;
typedef enum _BcdLibraryElementTypes {
    BcdLibraryDevice_ApplicationDevice                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 1),
    BcdLibraryString_ApplicationPath                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdLibraryString_Description                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 4),
    BcdLibraryString_PreferredLocale                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 5),
    BcdLibraryObjectList_InheritedObjects              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 6),
    BcdLibraryInteger_TruncatePhysicalMemory           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 7),
    BcdLibraryObjectList_RecoverySequence              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 8),
    BcdLibraryBoolean_AutoRecoveryEnabled              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 9),
    BcdLibraryIntegerList_BadMemoryList                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 10),
    BcdLibraryBoolean_AllowBadMemoryAccess             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 11),
    BcdLibraryInteger_FirstMegabytePolicy              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 12),
    BcdLibraryInteger_RelocatePhysicalMemory           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 13),
    BcdLibraryInteger_AvoidLowPhysicalMemory           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 14),
    BcdLibraryBoolean_TraditionalKsegMappings          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 15),
    BcdLibraryBoolean_DebuggerEnabled                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
    BcdLibraryInteger_DebuggerType                     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 17),
    BcdLibraryInteger_SerialDebuggerPortAddress        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 18),
    BcdLibraryInteger_SerialDebuggerPort               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 19),
    BcdLibraryInteger_SerialDebuggerBaudRate           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 20),
    BcdLibraryInteger_1394DebuggerChannel              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 21),
    BcdLibraryString_UsbDebuggerTargetName             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 22),
    BcdLibraryBoolean_DebuggerIgnoreUsermodeExceptions = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 23),
    BcdLibraryInteger_DebuggerStartPolicy              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 24),
    BcdLibraryString_DebuggerBusParameters             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 25),
    BcdLibraryInteger_DebuggerNetHostIP                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 26),
    BcdLibraryInteger_DebuggerNetPort                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 27),
    BcdLibraryBoolean_DebuggerNetDhcp                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 28),
    BcdLibraryString_DebuggerNetKey                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 29),
    BcdLibraryBoolean_DebuggerNetVM                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 30),
    BcdLibraryString_DebuggerNetHostIpv6               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 31),
    BcdLibraryBoolean_EmsEnabled                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 32),
    BcdLibraryInteger_EmsPort                          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 34),
    BcdLibraryInteger_EmsBaudRate                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 35),
    BcdLibraryString_LoadOptionsString                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 48),
    BcdLibraryBoolean_AttemptNonBcdStart               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 49),
    BcdLibraryBoolean_DisplayAdvancedOptions           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 64),
    BcdLibraryBoolean_DisplayOptionsEdit               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 65),
    BcdLibraryInteger_FVEKeyRingAddress                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 66),
    BcdLibraryDevice_BsdLogDevice                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 67),
    BcdLibraryString_BsdLogPath                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 68),
    BcdLibraryBoolean_BsdPreserveLog                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 69),
    BcdLibraryBoolean_GraphicsModeDisabled             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 70),
    BcdLibraryInteger_ConfigAccessPolicy               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 71),
    BcdLibraryBoolean_DisableIntegrityChecks           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 72),
    BcdLibraryBoolean_AllowPrereleaseSignatures        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 73),
    BcdLibraryString_FontPath                          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 74),
    BcdLibraryInteger_SiPolicy                         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 75),
    BcdLibraryInteger_FveBandId                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 76),
    BcdLibraryBoolean_ConsoleExtendedInput             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 80),
    BcdLibraryInteger_InitialConsoleInput              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 81),
    BcdLibraryInteger_GraphicsResolution               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 82),
    BcdLibraryBoolean_RestartOnFailure                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 83),
    BcdLibraryBoolean_GraphicsForceHighestMode         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 84),
    BcdLibraryBoolean_IsolatedExecutionContext         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 96),
    BcdLibraryInteger_BootUxDisplayMessage             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 101),
    BcdLibraryInteger_BootUxDisplayMessageOverride     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 102),
    BcdLibraryBoolean_BootUxLogoDisable                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 103),
    BcdLibraryBoolean_BootUxTextDisable                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 104),
    BcdLibraryBoolean_BootUxProgressDisable            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 105),
    BcdLibraryBoolean_BootUxFadeDisable                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 106),
    BcdLibraryBoolean_BootUxReservePoolDebug           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 107),
    BcdLibraryBoolean_BootUxDisable                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 108),
    BcdLibraryInteger_BootUxFadeFrames                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 109),
    BcdLibraryBoolean_BootUxDumpStats                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 110),
    BcdLibraryBoolean_BootUxShowStats                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 111),
    BcdLibraryBoolean_MultiBootSystem                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 113),
    BcdLibraryBoolean_ForceNoKeyboard                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 114),
    BcdLibraryInteger_AliasWindowsKey                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 115),
    BcdLibraryBoolean_BootShutdownDisabled             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 116),
    BcdLibraryInteger_PerformanceFrequency             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 117),
    BcdLibraryInteger_SecurebootRawPolicy              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 118),
    BcdLibraryIntegerList_AllowedInMemorySettings      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 119),
    BcdLibraryInteger_BootUxBitmapTransitionTime       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 121),
    BcdLibraryBoolean_TwoBootImages                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 122),
    BcdLibraryBoolean_ForceFipsCrypto                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 123),
    BcdLibraryInteger_BootErrorUx                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 125),
    BcdLibraryBoolean_AllowFlightSignatures            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 126),
    BcdLibraryInteger_BootMeasurementLogFormat         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 127),
    BcdLibraryInteger_DisplayRotation                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 128),
    BcdLibraryInteger_LogControl                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 129),
    BcdLibraryBoolean_NoFirmwareSync                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 130),
    BcdLibraryDevice_WindowsSystemDevice               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 132),
    BcdLibraryBoolean_NumLockOn                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 135),
    BcdLibraryString_AdditionalCiPolicy                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 136),
} BcdLibraryElementTypes;
typedef enum _BcdTemplateElementTypes {
    BcdSetupInteger_DeviceType                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 1),
    BcdSetupString_ApplicationRelativePath    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdSetupString_RamdiskDeviceRelativePath  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 3),
    BcdSetupBoolean_OmitOsLoaderElements      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 4),
    BcdSetupIntegerList_ElementsToMigrateList = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 6),
    BcdSetupBoolean_RecoveryOs                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
} BcdTemplateElementTypes;
typedef enum _BcdOSLoader_NxPolicy {
    NxPolicyOptIn     = 0,
    NxPolicyOptOut    = 1,
    NxPolicyAlwaysOff = 2,
    NxPolicyAlwaysOn  = 3
} BcdOSLoader_NxPolicy;
typedef enum _BcdOSLoader_PAEPolicy {
    PaePolicyDefault      = 0,
    PaePolicyForceEnable  = 1,
    PaePolicyForceDisable = 2
} BcdOSLoader_PAEPolicy;
typedef enum _BcdOSLoader_BootStatusPolicy {
    BootStatusPolicyDisplayAllFailures        = 0,
    BootStatusPolicyIgnoreAllFailures         = 1,
    BootStatusPolicyIgnoreShutdownFailures    = 2,
    BootStatusPolicyIgnoreBootFailures        = 3,
    BootStatusPolicyIgnoreCheckpointFailures  = 4,
    BootStatusPolicyDisplayShutdownFailures   = 5,
    BootStatusPolicyDisplayBootFailures       = 6,
    BootStatusPolicyDisplayCheckpointFailures = 7
} BcdOSLoaderBootStatusPolicy;
typedef enum _BcdOSLoaderElementTypes {
    BcdOSLoaderDevice_OSDevice                         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 1),
    BcdOSLoaderString_SystemRoot                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdOSLoaderObject_AssociatedResumeObject           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 3),
    BcdOSLoaderBoolean_StampDisks                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 4),
    BcdOSLoaderBoolean_DetectKernelAndHal              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
    BcdOSLoaderString_KernelPath                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 17),
    BcdOSLoaderString_HalPath                          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 18),
    BcdOSLoaderString_DbgTransportPath                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 19),
    BcdOSLoaderInteger_NxPolicy                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 32),
    BcdOSLoaderInteger_PAEPolicy                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 33),
    BcdOSLoaderBoolean_WinPEMode                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 34),
    BcdOSLoaderBoolean_DisableCrashAutoReboot          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 36),
    BcdOSLoaderBoolean_UseLastGoodSettings             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 37),
    BcdOSLoaderBoolean_DisableCodeIntegrityChecks      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 38),
    BcdOSLoaderBoolean_AllowPrereleaseSignatures       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 39),
    BcdOSLoaderBoolean_NoLowMemory                     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 48),
    BcdOSLoaderInteger_RemoveMemory                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 49),
    BcdOSLoaderInteger_IncreaseUserVa                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 50),
    BcdOSLoaderInteger_PerformaceDataMemory            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 51),
    BcdOSLoaderBoolean_UseVgaDriver                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 64),
    BcdOSLoaderBoolean_DisableBootDisplay              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 65),
    BcdOSLoaderBoolean_DisableVesaBios                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 66),
    BcdOSLoaderBoolean_DisableVgaMode                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 67),
    BcdOSLoaderInteger_ClusterModeAddressing           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 80),
    BcdOSLoaderBoolean_UsePhysicalDestination          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 81),
    BcdOSLoaderInteger_RestrictApicCluster             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 82),
    BcdOSLoaderString_OSLoaderTypeEVStore              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 83),
    BcdOSLoaderBoolean_UseLegacyApicMode               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 84),
    BcdOSLoaderInteger_X2ApicPolicy                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 85),
    BcdOSLoaderBoolean_UseBootProcessorOnly            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 96),
    BcdOSLoaderInteger_NumberOfProcessors              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 97),
    BcdOSLoaderBoolean_ForceMaximumProcessors          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 98),
    BcdOSLoaderBoolean_ProcessorConfigurationFlags     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 99),
    BcdOSLoaderBoolean_MaximizeGroupsCreated           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 100),
    BcdOSLoaderBoolean_ForceGroupAwareness             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 101),
    BcdOSLoaderInteger_GroupSize                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 102),
    BcdOSLoaderInteger_UseFirmwarePciSettings          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 112),
    BcdOSLoaderInteger_MsiPolicy                       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 113),
    BcdOSLoaderInteger_PciExpressPolicy                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 114),
    BcdOSLoaderInteger_SafeBoot                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 128),
    BcdOSLoaderBoolean_SafeBootAlternateShell          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 129),
    BcdOSLoaderBoolean_BootLogInitialization           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 144),
    BcdOSLoaderBoolean_VerboseObjectLoadMode           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 145),
    BcdOSLoaderBoolean_KernelDebuggerEnabled           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 160),
    BcdOSLoaderBoolean_DebuggerHalBreakpoint           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 161),
    BcdOSLoaderBoolean_UsePlatformClock                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 162),
    BcdOSLoaderBoolean_ForceLegacyPlatform             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 163),
    BcdOSLoaderBoolean_UsePlatformTick                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 164),
    BcdOSLoaderBoolean_DisableDynamicTick              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 165),
    BcdOSLoaderInteger_TscSyncPolicy                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 166),
    BcdOSLoaderBoolean_EmsEnabled                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 176),
    BcdOSLoaderInteger_ForceFailure                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 192),
    BcdOSLoaderInteger_DriverLoadFailurePolicy         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 193),
    BcdOSLoaderInteger_BootMenuPolicy                  = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 194),
    BcdOSLoaderBoolean_AdvancedOptionsOneTime          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 195),
    BcdOSLoaderBoolean_OptionsEditOneTime              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 196),
    BcdOSLoaderInteger_BootStatusPolicy                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 224),
    BcdOSLoaderBoolean_DisableElamDrivers              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 225),
    BcdOSLoaderInteger_HypervisorLaunchType            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 240),
    BcdOSLoaderString_HypervisorPath                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 241),
    BcdOSLoaderBoolean_HypervisorDebuggerEnabled       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 242),
    BcdOSLoaderInteger_HypervisorDebuggerType          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 243),
    BcdOSLoaderInteger_HypervisorDebuggerPortNumber    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 244),
    BcdOSLoaderInteger_HypervisorDebuggerBaudrate      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 245),
    BcdOSLoaderInteger_HypervisorDebugger1394Channel   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 246),
    BcdOSLoaderInteger_BootUxPolicy                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 247),
    BcdOSLoaderInteger_HypervisorSlatDisabled          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 248),
    BcdOSLoaderString_HypervisorDebuggerBusParams      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 249),
    BcdOSLoaderInteger_HypervisorNumProc               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 250),
    BcdOSLoaderInteger_HypervisorRootProcPerNode       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 251),
    BcdOSLoaderBoolean_HypervisorUseLargeVTlb          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 252),
    BcdOSLoaderInteger_HypervisorDebuggerNetHostIp     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 253),
    BcdOSLoaderInteger_HypervisorDebuggerNetHostPort   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 254),
    BcdOSLoaderInteger_HypervisorDebuggerPages         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 255),
    BcdOSLoaderInteger_TpmBootEntropyPolicy            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 256),
    BcdOSLoaderString_HypervisorDebuggerNetKey         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 272),
    BcdOSLoaderString_HypervisorProductSkuType         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 274),
    BcdOSLoaderInteger_HypervisorRootProc              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 275),
    BcdOSLoaderBoolean_HypervisorDebuggerNetDhcp       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 276),
    BcdOSLoaderInteger_HypervisorIommuPolicy           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 277),
    BcdOSLoaderBoolean_HypervisorUseVApic              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 278),
    BcdOSLoaderString_HypervisorLoadOptions            = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 279),
    BcdOSLoaderInteger_HypervisorMsrFilterPolicy       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 280),
    BcdOSLoaderInteger_HypervisorMmioNxPolicy          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 281),
    BcdOSLoaderInteger_HypervisorSchedulerType         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 282),
    BcdOSLoaderString_HypervisorRootProcNumaNodes      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 283),
    BcdOSLoaderInteger_HypervisorPerfmon               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 284),
    BcdOSLoaderInteger_HypervisorRootProcPerCore       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 285),
    BcdOSLoaderString_HypervisorRootProcNumaNodeLps    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 286),
    BcdOSLoaderInteger_XSavePolicy                     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 288),
    BcdOSLoaderInteger_XSaveAddFeature0                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 289),
    BcdOSLoaderInteger_XSaveAddFeature1                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 290),
    BcdOSLoaderInteger_XSaveAddFeature2                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 291),
    BcdOSLoaderInteger_XSaveAddFeature3                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 292),
    BcdOSLoaderInteger_XSaveAddFeature4                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 293),
    BcdOSLoaderInteger_XSaveAddFeature5                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 294),
    BcdOSLoaderInteger_XSaveAddFeature6                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 295),
    BcdOSLoaderInteger_XSaveAddFeature7                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 296),
    BcdOSLoaderInteger_XSaveRemoveFeature              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 297),
    BcdOSLoaderInteger_XSaveProcessorsMask             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 298),
    BcdOSLoaderInteger_XSaveDisable                    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 299),
    BcdOSLoaderInteger_KernelDebuggerType              = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 300),
    BcdOSLoaderString_KernelDebuggerBusParameters      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 301),
    BcdOSLoaderInteger_KernelDebuggerPortAddress       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 302),
    BcdOSLoaderInteger_KernelDebuggerPortNumber        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 303),
    BcdOSLoaderInteger_ClaimedTpmCounter               = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 304),
    BcdOSLoaderInteger_KernelDebugger1394Channel       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 305),
    BcdOSLoaderString_KernelDebuggerUsbTargetname      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 306),
    BcdOSLoaderInteger_KernelDebuggerNetHostIp         = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 307),
    BcdOSLoaderInteger_KernelDebuggerNetHostPort       = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 308),
    BcdOSLoaderBoolean_KernelDebuggerNetDhcp           = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 309),
    BcdOSLoaderString_KernelDebuggerNetKey             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 310),
    BcdOSLoaderString_IMCHiveName                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 311),
    BcdOSLoaderDevice_IMCDevice                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 312),
    BcdOSLoaderInteger_KernelDebuggerBaudrate          = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 313),
    BcdOSLoaderString_ManufacturingMode                = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 320),
    BcdOSLoaderBoolean_EventLoggingEnabled             = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 321),
    BcdOSLoaderInteger_VsmLaunchType                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 322),
    BcdOSLoaderInteger_HypervisorEnforcedCodeIntegrity = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 324),
    BcdOSLoaderBoolean_DtraceEnabled                   = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 325),
    BcdOSLoaderDevice_SystemDataDevice                 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 336),
    BcdOSLoaderDevice_OsArcDevice                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 337),
    BcdOSLoaderDevice_OsDataDevice                     = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 339),
    BcdOSLoaderDevice_BspDevice                        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 340),
    BcdOSLoaderDevice_BspFilepath                      = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 341),
    BcdOSLoaderString_KernelDebuggerNetHostIpv6        = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 342),
    BcdOSLoaderString_HypervisorDebuggerNetHostIpv6    = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 353),
} BcdOSLoaderElementTypes;
#endif
