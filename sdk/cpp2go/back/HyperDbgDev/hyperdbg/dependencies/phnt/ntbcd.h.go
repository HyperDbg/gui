package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntbcd.h.back

const(
_NTBCD_H =  //col:13
MAKE_BCD_OBJECT(ObjectType, ImageType, ApplicationType) = (((ULONG)(ObjectType) << 28) | (((ULONG)(ImageType) & 0xF) << 20) | ((ULONG)(ApplicationType) & 0xFFFFF)) //col:286
MAKE_BCD_APPLICATION_OBJECT(ImageType, ApplicationType) = MAKE_BCD_OBJECT(BCD_OBJECT_TYPE_APPLICATION, (ULONG)(ImageType), (ULONG)(ApplicationType)) //col:291
GET_BCD_OBJECT_TYPE(DataType) = ((BCD_OBJECT_TYPE)(((((ULONG)DataType)) >> 28) & 0xF)) //col:294
GET_BCD_APPLICATION_IMAGE(DataType) = ((BCD_APPLICATION_IMAGE_TYPE)(((((ULONG)DataType)) >> 20) & 0xF)) //col:296
GET_BCD_APPLICATION_OBJECT(DataType) = ((BCD_APPLICATION_OBJECT_TYPE)((((ULONG)DataType)) & 0xFFFFF)) //col:298
BCD_OBJECT_OSLOADER_TYPE = MAKE_BCD_APPLICATION_OBJECT(BCD_APPLICATION_IMAGE_BOOT_APPLICATION, BCD_APPLICATION_OBJECT_WINDOWS_BOOT_LOADER) //col:301
BCD_OBJECT_DESCRIPTION_VERSION = 0x1 //col:336
MAKE_BCDE_DATA_TYPE(Class, Format, Subtype) = (((((ULONG)Class) & 0xF) << 28) | ((((ULONG)Format) & 0xF) << 24) | (((ULONG)Subtype) & 0x00FFFFFF)) //col:495
GET_BCDE_DATA_CLASS(DataType) = ((BCD_ELEMENT_DATATYPE_CLASS)(((((ULONG)DataType)) >> 28) & 0xF)) //col:498
GET_BCDE_DATA_FORMAT(DataType) = ((BCD_ELEMENT_DATATYPE_FORMAT)(((((ULONG)DataType)) >> 24) & 0xF)) //col:500
GET_BCDE_DATA_SUBTYPE(DataType) = ((ULONG)((((ULONG)DataType)) & 0x00FFFFFF)) //col:502
BCD_ELEMENT_DESCRIPTION_VERSION = 0x1 //col:610
)

type     BCD_MESSAGE_TYPE_NONE uint32
const(
    BCD_MESSAGE_TYPE_NONE BCD_MESSAGE_TYPE = 1  //col:84
    BCD_MESSAGE_TYPE_TRACE BCD_MESSAGE_TYPE = 2  //col:85
    BCD_MESSAGE_TYPE_INFORMATION BCD_MESSAGE_TYPE = 3  //col:86
    BCD_MESSAGE_TYPE_WARNING BCD_MESSAGE_TYPE = 4  //col:87
    BCD_MESSAGE_TYPE_ERROR BCD_MESSAGE_TYPE = 5  //col:88
    BCD_MESSAGE_TYPE_MAXIMUM BCD_MESSAGE_TYPE = 6  //col:89
)


type     BCD_IMPORT_NONE uint32
const(
    BCD_IMPORT_NONE BCD_IMPORT_FLAGS = 1  //col:174
    BCD_IMPORT_DELETE_FIRMWARE_OBJECTS BCD_IMPORT_FLAGS = 2  //col:175
)


type     BCD_OPEN_NONE uint32
const(
    BCD_OPEN_NONE BCD_OPEN_FLAGS = 1  //col:203
    BCD_OPEN_OPEN_STORE_OFFLINE BCD_OPEN_FLAGS = 2  //col:204
    BCD_OPEN_SYNC_FIRMWARE_ENTRIES BCD_OPEN_FLAGS = 3  //col:205
)


type     BCD_OBJECT_TYPE_NONE uint32
const(
    BCD_OBJECT_TYPE_NONE BCD_OBJECT_TYPE = 1  //col:247
    BCD_OBJECT_TYPE_APPLICATION BCD_OBJECT_TYPE = 2  //col:248
    BCD_OBJECT_TYPE_INHERITED BCD_OBJECT_TYPE = 3  //col:249
    BCD_OBJECT_TYPE_DEVICE BCD_OBJECT_TYPE = 4  //col:250
)


type     BCD_APPLICATION_OBJECT_NONE = 0 uint32
const(
    BCD_APPLICATION_OBJECT_NONE  BCD_APPLICATION_OBJECT_TYPE =  0  //col:255
    BCD_APPLICATION_OBJECT_FIRMWARE_BOOT_MANAGER  BCD_APPLICATION_OBJECT_TYPE =  1  //col:256
    BCD_APPLICATION_OBJECT_WINDOWS_BOOT_MANAGER  BCD_APPLICATION_OBJECT_TYPE =  2  //col:257
    BCD_APPLICATION_OBJECT_WINDOWS_BOOT_LOADER  BCD_APPLICATION_OBJECT_TYPE =  3  //col:258
    BCD_APPLICATION_OBJECT_WINDOWS_RESUME_APPLICATION  BCD_APPLICATION_OBJECT_TYPE =  4  //col:259
    BCD_APPLICATION_OBJECT_MEMORY_TESTER  BCD_APPLICATION_OBJECT_TYPE =  5  //col:260
    BCD_APPLICATION_OBJECT_LEGACY_NTLDR  BCD_APPLICATION_OBJECT_TYPE =  6  //col:261
    BCD_APPLICATION_OBJECT_LEGACY_SETUPLDR  BCD_APPLICATION_OBJECT_TYPE =  7  //col:262
    BCD_APPLICATION_OBJECT_BOOT_SECTOR  BCD_APPLICATION_OBJECT_TYPE =  8  //col:263
    BCD_APPLICATION_OBJECT_STARTUP_MODULE  BCD_APPLICATION_OBJECT_TYPE =  9  //col:264
    BCD_APPLICATION_OBJECT_GENERIC_APPLICATION  BCD_APPLICATION_OBJECT_TYPE =  10  //col:265
    BCD_APPLICATION_OBJECT_RESERVED  BCD_APPLICATION_OBJECT_TYPE =  0xFFFFF  //col:266
)


type     BCD_APPLICATION_IMAGE_NONE uint32
const(
    BCD_APPLICATION_IMAGE_NONE BCD_APPLICATION_IMAGE_TYPE = 1  //col:271
    BCD_APPLICATION_IMAGE_FIRMWARE_APPLICATION BCD_APPLICATION_IMAGE_TYPE = 2  //col:272
    BCD_APPLICATION_IMAGE_BOOT_APPLICATION BCD_APPLICATION_IMAGE_TYPE = 3  //col:273
    BCD_APPLICATION_IMAGE_LEGACY_LOADER BCD_APPLICATION_IMAGE_TYPE = 4  //col:274
    BCD_APPLICATION_IMAGE_REALMODE_CODE BCD_APPLICATION_IMAGE_TYPE = 5  //col:275
)


type     BCD_INHERITED_CLASS_NONE uint32
const(
    BCD_INHERITED_CLASS_NONE BCD_INHERITED_CLASS_TYPE = 1  //col:280
    BCD_INHERITED_CLASS_LIBRARY BCD_INHERITED_CLASS_TYPE = 2  //col:281
    BCD_INHERITED_CLASS_APPLICATION BCD_INHERITED_CLASS_TYPE = 3  //col:282
    BCD_INHERITED_CLASS_DEVICE BCD_INHERITED_CLASS_TYPE = 4  //col:283
)


type     BCD_COPY_NONE = 0x0 uint32
const(
    BCD_COPY_NONE  BCD_COPY_FLAGS =  0x0  //col:396
    BCD_COPY_COPY_CREATE_NEW_OBJECT_IDENTIFIER  BCD_COPY_FLAGS =  0x1  //col:397
    BCD_COPY_COPY_DELETE_EXISTING_OBJECT  BCD_COPY_FLAGS =  0x2  //col:398
    BCD_COPY_COPY_UNKNOWN_FIRMWARE_APPLICATION  BCD_COPY_FLAGS =  0x4  //col:399
    BCD_COPY_IGNORE_SETUP_TEMPLATE_ELEMENTS  BCD_COPY_FLAGS =  0x8  //col:400
    BCD_COPY_RETAIN_ELEMENT_DATA  BCD_COPY_FLAGS =  0x10  //col:401
    BCD_COPY_MIGRATE_ELEMENT_DATA  BCD_COPY_FLAGS =  0x20  //col:402
)


type     BCD_ELEMENT_DATATYPE_FORMAT_UNKNOWN uint32
const(
    BCD_ELEMENT_DATATYPE_FORMAT_UNKNOWN BCD_ELEMENT_DATATYPE_FORMAT = 1  //col:459
    BCD_ELEMENT_DATATYPE_FORMAT_DEVICE // 0x01000000 BCD_ELEMENT_DATATYPE_FORMAT = 2  //col:460
    BCD_ELEMENT_DATATYPE_FORMAT_STRING // 0x02000000 BCD_ELEMENT_DATATYPE_FORMAT = 3  //col:461
    BCD_ELEMENT_DATATYPE_FORMAT_OBJECT // 0x03000000  BCD_ELEMENT_DATATYPE_FORMAT = 4  //col:462
    BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST // 0x04000000  BCD_ELEMENT_DATATYPE_FORMAT = 5  //col:463
    BCD_ELEMENT_DATATYPE_FORMAT_INTEGER // 0x05000000 BCD_ELEMENT_DATATYPE_FORMAT = 6  //col:464
    BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN // 0x06000000 BCD_ELEMENT_DATATYPE_FORMAT = 7  //col:465
    BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST // 0x07000000 BCD_ELEMENT_DATATYPE_FORMAT = 8  //col:466
    BCD_ELEMENT_DATATYPE_FORMAT_BINARY // 0x08000000 BCD_ELEMENT_DATATYPE_FORMAT = 9  //col:467
)


type     BCD_ELEMENT_DATATYPE_CLASS_NONE uint32
const(
    BCD_ELEMENT_DATATYPE_CLASS_NONE BCD_ELEMENT_DATATYPE_CLASS = 1  //col:472
    BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_CLASS = 2  //col:473
    BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_CLASS = 3  //col:474
    BCD_ELEMENT_DATATYPE_CLASS_DEVICE BCD_ELEMENT_DATATYPE_CLASS = 4  //col:475
    BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_CLASS = 5  //col:476
    BCD_ELEMENT_DATATYPE_CLASS_OEM BCD_ELEMENT_DATATYPE_CLASS = 6  //col:477
)


type     BCD_ELEMENT_DEVICE_TYPE_NONE uint32
const(
    BCD_ELEMENT_DEVICE_TYPE_NONE BCD_ELEMENT_DEVICE_TYPE = 1  //col:482
    BCD_ELEMENT_DEVICE_TYPE_BOOT_DEVICE BCD_ELEMENT_DEVICE_TYPE = 2  //col:483
    BCD_ELEMENT_DEVICE_TYPE_PARTITION BCD_ELEMENT_DEVICE_TYPE = 3  //col:484
    BCD_ELEMENT_DEVICE_TYPE_FILE BCD_ELEMENT_DEVICE_TYPE = 4  //col:485
    BCD_ELEMENT_DEVICE_TYPE_RAMDISK BCD_ELEMENT_DEVICE_TYPE = 5  //col:486
    BCD_ELEMENT_DEVICE_TYPE_UNKNOWN BCD_ELEMENT_DEVICE_TYPE = 6  //col:487
    BCD_ELEMENT_DEVICE_TYPE_QUALIFIED_PARTITION BCD_ELEMENT_DEVICE_TYPE = 7  //col:488
    BCD_ELEMENT_DEVICE_TYPE_VMBUS BCD_ELEMENT_DEVICE_TYPE = 8  //col:489
    BCD_ELEMENT_DEVICE_TYPE_LOCATE_DEVICE BCD_ELEMENT_DEVICE_TYPE = 9  //col:490
    BCD_ELEMENT_DEVICE_TYPE_URI BCD_ELEMENT_DEVICE_TYPE = 10  //col:491
    BCD_ELEMENT_DEVICE_TYPE_COMPOSITE BCD_ELEMENT_DEVICE_TYPE = 11  //col:492
)


type     BCD_FLAG_NONE = 0x0 uint32
const(
    BCD_FLAG_NONE  BCD_FLAGS =  0x0  //col:637
    BCD_FLAG_QUALIFIED_PARTITION  BCD_FLAGS =  0x1  //col:638
    BCD_FLAG_NO_DEVICE_TRANSLATION  BCD_FLAGS =  0x2  //col:639
    BCD_FLAG_ENUMERATE_INHERITED_OBJECTS  BCD_FLAGS =  0x4  //col:640
    BCD_FLAG_ENUMERATE_DEVICE_OPTIONS  BCD_FLAGS =  0x8  //col:641
    BCD_FLAG_OBSERVE_PRECEDENCE  BCD_FLAGS =  0x10  //col:642
    BCD_FLAG_DISABLE_VHD_NT_TRANSLATION  BCD_FLAGS =  0x20  //col:643
    BCD_FLAG_DISABLE_VHD_DEVICE_DETECTION  BCD_FLAGS =  0x40  //col:644
    BCD_FLAG_DISABLE_POLICY_CHECKS  BCD_FLAGS =  0x80  //col:645
)


type     /// <summary> uint32
const(
    /// <summary> BcdBootMgrElementTypes = 1  //col:725
    /// The order in which BCD objects should be displayed. BcdBootMgrElementTypes = 2  //col:726
    /// Objects are displayed using the string specified by the BcdLibraryString_Description element. BcdBootMgrElementTypes = 3  //col:727
    /// </summary> BcdBootMgrElementTypes = 4  //col:728
    /// <remarks>0x24000001</remarks> BcdBootMgrElementTypes = 5  //col:729
    BcdBootMgrObjectList_DisplayOrder  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 1)  //col:730
    /// <summary> BcdBootMgrElementTypes = 7  //col:731
    /// List of boot environment applications the boot manager should execute. BcdBootMgrElementTypes = 8  //col:732
    /// The applications are executed in the order they appear in this list. BcdBootMgrElementTypes = 9  //col:733
    /// If the firmware boot manager does not support loading multiple applications this list cannot contain more than one entry. BcdBootMgrElementTypes = 10  //col:734
    /// </summary> BcdBootMgrElementTypes = 11  //col:735
    /// <remarks>0x24000002</remarks> BcdBootMgrElementTypes = 12  //col:736
    BcdBootMgrObjectList_BootSequence  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 2)  //col:737
    /// <summary> BcdBootMgrElementTypes = 14  //col:738
    /// The default boot environment application to load if the user does not select one. BcdBootMgrElementTypes = 15  //col:739
    /// </summary> BcdBootMgrElementTypes = 16  //col:740
    /// <remarks>0x23000003</remarks> BcdBootMgrElementTypes = 17  //col:741
    BcdBootMgrObject_DefaultObject  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECT 3)  //col:742
    /// <summary> BcdBootMgrElementTypes = 19  //col:743
    /// The maximum number of seconds a boot selection menu is to be displayed to the user. BcdBootMgrElementTypes = 20  //col:744
    /// The menu is displayed until the user selects an option or the time-out expires. BcdBootMgrElementTypes = 21  //col:745
    /// If this value is not specified the boot manager waits for the user to make a selection. BcdBootMgrElementTypes = 22  //col:746
    /// </summary> BcdBootMgrElementTypes = 23  //col:747
    /// <remarks>0x25000004</remarks> BcdBootMgrElementTypes = 24  //col:748
    BcdBootMgrInteger_Timeout  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 4)  //col:749
    /// <summary> BcdBootMgrElementTypes = 26  //col:750
    /// Indicates that a resume operation should be attempted during a system restart. BcdBootMgrElementTypes = 27  //col:751
    /// </summary> BcdBootMgrElementTypes = 28  //col:752
    /// <remarks>0x26000005</remarks> BcdBootMgrElementTypes = 29  //col:753
    BcdBootMgrBoolean_AttemptResume  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 5)  //col:754
    /// <summary> BcdBootMgrElementTypes = 31  //col:755
    /// The resume application object. BcdBootMgrElementTypes = 32  //col:756
    /// </summary> BcdBootMgrElementTypes = 33  //col:757
    /// <remarks>0x23000006</remarks> BcdBootMgrElementTypes = 34  //col:758
    BcdBootMgrObject_ResumeObject  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECT 6)  //col:759
    /// <summary> BcdBootMgrElementTypes = 36  //col:760
    /// BcdBootMgrElementTypes = 37  //col:761
    /// </summary> BcdBootMgrElementTypes = 38  //col:762
    /// <remarks>0x24000007</remarks> BcdBootMgrElementTypes = 39  //col:763
    BcdBootMgrObjectList_StartupSequence  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 7)  //col:764
    /// <summary> BcdBootMgrElementTypes = 41  //col:765
    /// The boot manager tools display order list. BcdBootMgrElementTypes = 42  //col:766
    /// </summary> BcdBootMgrElementTypes = 43  //col:767
    /// <remarks>0x24000010</remarks> BcdBootMgrElementTypes = 44  //col:768
    BcdBootMgrObjectList_ToolsDisplayOrder  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 16)  //col:769
    /// <summary> BcdBootMgrElementTypes = 46  //col:770
    /// Forces the display of the legacy boot menu regardless of the number of OS entries in the BCD store and their BcdOSLoaderInteger_BootMenuPolicy. BcdBootMgrElementTypes = 47  //col:771
    /// </summary> BcdBootMgrElementTypes = 48  //col:772
    /// <remarks>0x26000020</remarks> BcdBootMgrElementTypes = 49  //col:773
    BcdBootMgrBoolean_DisplayBootMenu  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 32)  //col:774
    /// <summary> BcdBootMgrElementTypes = 51  //col:775
    /// Indicates whether the display of errors should be suppressed. BcdBootMgrElementTypes = 52  //col:776
    /// If this setting is enabled the boot manager exits to the multi-OS menu on OS launch error. BcdBootMgrElementTypes = 53  //col:777
    /// </summary> BcdBootMgrElementTypes = 54  //col:778
    /// <remarks>0x26000021</remarks> BcdBootMgrElementTypes = 55  //col:779
    BcdBootMgrBoolean_NoErrorDisplay  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 33)  //col:780
    /// <summary> BcdBootMgrElementTypes = 57  //col:781
    /// The device on which the boot application resides. BcdBootMgrElementTypes = 58  //col:782
    /// </summary> BcdBootMgrElementTypes = 59  //col:783
    /// <remarks>0x21000022</remarks> BcdBootMgrElementTypes = 60  //col:784
    BcdBootMgrDevice_BcdDevice  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 34)  //col:785
    /// <summary> BcdBootMgrElementTypes = 62  //col:786
    /// The boot application. BcdBootMgrElementTypes = 63  //col:787
    /// </summary> BcdBootMgrElementTypes = 64  //col:788
    /// <remarks>0x22000023</remarks> BcdBootMgrElementTypes = 65  //col:789
    BcdBootMgrString_BcdFilePath  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 35)  //col:790
    /// <summary> BcdBootMgrElementTypes = 67  //col:791
    /// BcdBootMgrElementTypes = 68  //col:792
    /// </summary> BcdBootMgrElementTypes = 69  //col:793
    /// <remarks>0x26000024</remarks> BcdBootMgrElementTypes = 70  //col:794
    BcdBootMgrBoolean_HormEnabled  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 36)  //col:795
    /// <summary> BcdBootMgrElementTypes = 72  //col:796
    /// BcdBootMgrElementTypes = 73  //col:797
    /// </summary> BcdBootMgrElementTypes = 74  //col:798
    /// <remarks>0x26000025</remarks> BcdBootMgrElementTypes = 75  //col:799
    BcdBootMgrBoolean_HiberRoot  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 37)  //col:800
    /// <summary> BcdBootMgrElementTypes = 77  //col:801
    /// BcdBootMgrElementTypes = 78  //col:802
    /// </summary> BcdBootMgrElementTypes = 79  //col:803
    /// <remarks>0x22000026</remarks> BcdBootMgrElementTypes = 80  //col:804
    BcdBootMgrString_PasswordOverride  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 38)  //col:805
    /// <summary> BcdBootMgrElementTypes = 82  //col:806
    /// BcdBootMgrElementTypes = 83  //col:807
    /// </summary> BcdBootMgrElementTypes = 84  //col:808
    /// <remarks>0x22000027</remarks> BcdBootMgrElementTypes = 85  //col:809
    BcdBootMgrString_PinpassPhraseOverride  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 39)  //col:810
    /// <summary> BcdBootMgrElementTypes = 87  //col:811
    /// Controls whether custom actions are processed before a boot sequence. BcdBootMgrElementTypes = 88  //col:812
    /// Note This value is supported starting in Windows 8 and Windows Server 2012. BcdBootMgrElementTypes = 89  //col:813
    /// </summary> BcdBootMgrElementTypes = 90  //col:814
    /// <remarks>0x26000028</remarks> BcdBootMgrElementTypes = 91  //col:815
    BcdBootMgrBoolean_ProcessCustomActionsFirst  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 40)  //col:816
    /// <summary> BcdBootMgrElementTypes = 93  //col:817
    /// Custom Bootstrap Actions. BcdBootMgrElementTypes = 94  //col:818
    /// </summary> BcdBootMgrElementTypes = 95  //col:819
    /// <remarks>0x27000030</remarks> BcdBootMgrElementTypes = 96  //col:820
    BcdBootMgrIntegerList_CustomActionsList  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST 48)  //col:821
    /// <summary> BcdBootMgrElementTypes = 98  //col:822
    /// Controls whether a boot sequence persists across multiple boots. BcdBootMgrElementTypes = 99  //col:823
    /// Note This value is supported starting in Windows 8 and Windows Server 2012. BcdBootMgrElementTypes = 100  //col:824
    /// </summary> BcdBootMgrElementTypes = 101  //col:825
    /// <remarks>0x26000031</remarks> BcdBootMgrElementTypes = 102  //col:826
    BcdBootMgrBoolean_PersistBootSequence  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 49)  //col:827
    /// <summary> BcdBootMgrElementTypes = 104  //col:828
    /// BcdBootMgrElementTypes = 105  //col:829
    /// </summary> BcdBootMgrElementTypes = 106  //col:830
    /// <remarks>0x26000032</remarks> BcdBootMgrElementTypes = 107  //col:831
    BcdBootMgrBoolean_SkipStartupSequence  BcdBootMgrElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 50)  //col:832
)


type     /// <summary> uint32
const(
    /// <summary> BcdLibrary_FirstMegabytePolicy = 1  //col:837
    /// Use none of the first megabyte of memory. BcdLibrary_FirstMegabytePolicy = 2  //col:838
    /// </summary> BcdLibrary_FirstMegabytePolicy = 3  //col:839
    FirstMegabytePolicyUseNone BcdLibrary_FirstMegabytePolicy = 4  //col:840
    /// <summary> BcdLibrary_FirstMegabytePolicy = 5  //col:841
    /// Use all of the first megabyte of memory. BcdLibrary_FirstMegabytePolicy = 6  //col:842
    /// </summary> BcdLibrary_FirstMegabytePolicy = 7  //col:843
    FirstMegabytePolicyUseAll BcdLibrary_FirstMegabytePolicy = 8  //col:844
    /// <summary> BcdLibrary_FirstMegabytePolicy = 9  //col:845
    /// Reserved for future use. BcdLibrary_FirstMegabytePolicy = 10  //col:846
    /// </summary> BcdLibrary_FirstMegabytePolicy = 11  //col:847
    FirstMegabytePolicyUsePrivate BcdLibrary_FirstMegabytePolicy = 12  //col:848
)


type     DebuggerSerial = 0 uint32
const(
    DebuggerSerial  BcdLibrary_DebuggerType =  0  //col:853
    Debugger1394  BcdLibrary_DebuggerType =  1  //col:854
    DebuggerUsb  BcdLibrary_DebuggerType =  2  //col:855
    DebuggerNet  BcdLibrary_DebuggerType =  3  //col:856
    DebuggerLocal  BcdLibrary_DebuggerType =  4  //col:857
)


type     /// <summary> uint32
const(
    /// <summary> BcdLibrary_DebuggerStartPolicy = 1  //col:862
    /// The debugger will start active. BcdLibrary_DebuggerStartPolicy = 2  //col:863
    /// </summary> BcdLibrary_DebuggerStartPolicy = 3  //col:864
    DebuggerStartActive BcdLibrary_DebuggerStartPolicy = 4  //col:865
    /// <summary> BcdLibrary_DebuggerStartPolicy = 5  //col:866
    /// The debugger will start in the auto-enabled state. BcdLibrary_DebuggerStartPolicy = 6  //col:867
    /// If a debugger is attached it will be used; otherwise the debugger port will be available for other applications. BcdLibrary_DebuggerStartPolicy = 7  //col:868
    /// </summary> BcdLibrary_DebuggerStartPolicy = 8  //col:869
    DebuggerStartAutoEnable BcdLibrary_DebuggerStartPolicy = 9  //col:870
    /// <summary> BcdLibrary_DebuggerStartPolicy = 10  //col:871
    /// The debugger will not start. BcdLibrary_DebuggerStartPolicy = 11  //col:872
    /// </summary> BcdLibrary_DebuggerStartPolicy = 12  //col:873
    DebuggerStartDisable BcdLibrary_DebuggerStartPolicy = 13  //col:874
)


type     /// <summary> uint32
const(
    /// <summary> BcdLibrary_ConfigAccessPolicy = 1  //col:879
    /// Access to PCI configuration space through the memory-mapped region is allowed. BcdLibrary_ConfigAccessPolicy = 2  //col:880
    /// </summary> BcdLibrary_ConfigAccessPolicy = 3  //col:881
    ConfigAccessPolicyDefault BcdLibrary_ConfigAccessPolicy = 4  //col:882
    /// <summary> BcdLibrary_ConfigAccessPolicy = 5  //col:883
    /// Access to PCI configuration space through the memory-mapped region is not allowed. BcdLibrary_ConfigAccessPolicy = 6  //col:884
    /// This setting is used for platforms that implement memory-mapped configuration space incorrectly. BcdLibrary_ConfigAccessPolicy = 7  //col:885
    /// The CFC/CF8 access mechanism can be used to access configuration space on these platforms. BcdLibrary_ConfigAccessPolicy = 8  //col:886
    /// </summary> BcdLibrary_ConfigAccessPolicy = 9  //col:887
    ConfigAccessPolicyDisallowMmConfig BcdLibrary_ConfigAccessPolicy = 10  //col:888
)


type     DisplayMessageTypeDefault = 0 uint32
const(
    DisplayMessageTypeDefault  BcdLibrary_UxDisplayMessageType =  0  //col:893
    DisplayMessageTypeResume  BcdLibrary_UxDisplayMessageType =  1  //col:894
    DisplayMessageTypeHyperV  BcdLibrary_UxDisplayMessageType =  2  //col:895
    DisplayMessageTypeRecovery  BcdLibrary_UxDisplayMessageType =  3  //col:896
    DisplayMessageTypeStartupRepair  BcdLibrary_UxDisplayMessageType =  4  //col:897
    DisplayMessageTypeSystemImageRecovery  BcdLibrary_UxDisplayMessageType =  5  //col:898
    DisplayMessageTypeCommandPrompt  BcdLibrary_UxDisplayMessageType =  6  //col:899
    DisplayMessageTypeSystemRestore  BcdLibrary_UxDisplayMessageType =  7  //col:900
    DisplayMessageTypePushButtonReset  BcdLibrary_UxDisplayMessageType =  8  //col:901
)


type     /// <summary> uint32
const(
    /// <summary> typedef enum BcdLibrary_SafeBoot = 1  //col:906
    /// Load the drivers and services specified by name or group under the following registry key: typedef enum BcdLibrary_SafeBoot = 2  //col:907
    /// HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\SafeBoot\Minimal. typedef enum BcdLibrary_SafeBoot = 3  //col:908
    /// </summary> typedef enum BcdLibrary_SafeBoot = 4  //col:909
    SafemodeMinimal  typedef enum BcdLibrary_SafeBoot =  0  //col:910
    /// <summary> typedef enum BcdLibrary_SafeBoot = 6  //col:911
    /// Load the drivers and services specified by name or group under the following registry key: typedef enum BcdLibrary_SafeBoot = 7  //col:912
    /// HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\SafeBoot\Network typedef enum BcdLibrary_SafeBoot = 8  //col:913
    /// </summary> typedef enum BcdLibrary_SafeBoot = 9  //col:914
    SafemodeNetwork  typedef enum BcdLibrary_SafeBoot =  1  //col:915
    /// <summary> typedef enum BcdLibrary_SafeBoot = 11  //col:916
    /// Boot the system into a repair mode that restores the Active Directory service from backup medium. typedef enum BcdLibrary_SafeBoot = 12  //col:917
    /// </summary> typedef enum BcdLibrary_SafeBoot = 13  //col:918
    SafemodeDsRepair  typedef enum BcdLibrary_SafeBoot =  2  //col:919
)


type     /// <summary> uint32
const(
    /// <summary> BcdLibraryElementTypes = 1  //col:925
    /// Device on which a boot environment application resides. BcdLibraryElementTypes = 2  //col:926
    /// </summary> BcdLibraryElementTypes = 3  //col:927
    /// <remarks>0x11000001</remarks> BcdLibraryElementTypes = 4  //col:928
    BcdLibraryDevice_ApplicationDevice  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 1)  //col:929
    /// <summary> BcdLibraryElementTypes = 6  //col:930
    /// Path to a boot environment application. BcdLibraryElementTypes = 7  //col:931
    /// </summary> BcdLibraryElementTypes = 8  //col:932
    /// <remarks>0x12000002</remarks> BcdLibraryElementTypes = 9  //col:933
    BcdLibraryString_ApplicationPath  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 2)  //col:934
    /// <summary> BcdLibraryElementTypes = 11  //col:935
    /// Display name of the boot environment application. BcdLibraryElementTypes = 12  //col:936
    /// </summary> BcdLibraryElementTypes = 13  //col:937
    /// <remarks>0x12000004</remarks> BcdLibraryElementTypes = 14  //col:938
    BcdLibraryString_Description  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 4)  //col:939
    /// <summary> BcdLibraryElementTypes = 16  //col:940
    /// Preferred locale in RFC 3066 format. BcdLibraryElementTypes = 17  //col:941
    /// </summary> BcdLibraryElementTypes = 18  //col:942
    /// <remarks>0x12000005</remarks> BcdLibraryElementTypes = 19  //col:943
    BcdLibraryString_PreferredLocale  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 5)  //col:944
    /// <summary> BcdLibraryElementTypes = 21  //col:945
    /// List of BCD objects from which the current object should inherit elements. BcdLibraryElementTypes = 22  //col:946
    /// </summary> BcdLibraryElementTypes = 23  //col:947
    /// <remarks>0x14000006</remarks> BcdLibraryElementTypes = 24  //col:948
    BcdLibraryObjectList_InheritedObjects  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 6)  //col:949
    /// <summary> BcdLibraryElementTypes = 26  //col:950
    /// Maximum physical address a boot environment application should recognize. All memory above this address is ignored. BcdLibraryElementTypes = 27  //col:951
    /// </summary> BcdLibraryElementTypes = 28  //col:952
    /// <remarks>0x15000007</remarks> BcdLibraryElementTypes = 29  //col:953
    BcdLibraryInteger_TruncatePhysicalMemory  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 7)  //col:954
    /// <summary> BcdLibraryElementTypes = 31  //col:955
    /// List of boot environment applications to be executed if the associated application fails. The applications are executed in the order they appear in this list. BcdLibraryElementTypes = 32  //col:956
    /// </summary> BcdLibraryElementTypes = 33  //col:957
    /// <remarks>0x14000008</remarks> BcdLibraryElementTypes = 34  //col:958
    BcdLibraryObjectList_RecoverySequence  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST 8)  //col:959
    /// <summary> BcdLibraryElementTypes = 36  //col:960
    /// Indicates whether the recovery sequence executes automatically if the boot application fails. Otherwise the recovery sequence only runs on demand. BcdLibraryElementTypes = 37  //col:961
    /// </summary> BcdLibraryElementTypes = 38  //col:962
    /// <remarks>0x16000009</remarks> BcdLibraryElementTypes = 39  //col:963
    BcdLibraryBoolean_AutoRecoveryEnabled  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 9)  //col:964
    /// <summary> BcdLibraryElementTypes = 41  //col:965
    /// List of page frame numbers describing faulty memory in the system. BcdLibraryElementTypes = 42  //col:966
    /// </summary> BcdLibraryElementTypes = 43  //col:967
    /// <remarks>0x1700000A</remarks> BcdLibraryElementTypes = 44  //col:968
    BcdLibraryIntegerList_BadMemoryList  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST 10)  //col:969
    /// <summary> BcdLibraryElementTypes = 46  //col:970
    /// If TRUE indicates that a boot application can use memory listed in the BcdLibraryIntegerList_BadMemoryList. BcdLibraryElementTypes = 47  //col:971
    /// </summary> BcdLibraryElementTypes = 48  //col:972
    /// <remarks>0x1600000B</remarks> BcdLibraryElementTypes = 49  //col:973
    BcdLibraryBoolean_AllowBadMemoryAccess  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 11)  //col:974
    /// <summary> BcdLibraryElementTypes = 51  //col:975
    /// Indicates how the first megabyte of memory is to be used. The Integer property is one of the values from the BcdLibrary_FirstMegabytePolicy enumeration. BcdLibraryElementTypes = 52  //col:976
    /// </summary> BcdLibraryElementTypes = 53  //col:977
    /// <remarks>0x1500000C</remarks> BcdLibraryElementTypes = 54  //col:978
    BcdLibraryInteger_FirstMegabytePolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 12)  //col:979
    /// <summary> BcdLibraryElementTypes = 56  //col:980
    /// Relocates physical memory on certain AMD processors. BcdLibraryElementTypes = 57  //col:981
    /// This value is not used in Windows 8 or Windows Server 2012. BcdLibraryElementTypes = 58  //col:982
    /// </summary> BcdLibraryElementTypes = 59  //col:983
    /// <remarks>0x1500000D</remarks> BcdLibraryElementTypes = 60  //col:984
    BcdLibraryInteger_RelocatePhysicalMemory  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 13)  //col:985
    /// <summary> BcdLibraryElementTypes = 62  //col:986
    /// Specifies a minimum physical address to use in the boot environment. BcdLibraryElementTypes = 63  //col:987
    /// </summary> BcdLibraryElementTypes = 64  //col:988
    /// <remarks>0x1500000E</remarks> BcdLibraryElementTypes = 65  //col:989
    BcdLibraryInteger_AvoidLowPhysicalMemory  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 14)  //col:990
    /// <summary> BcdLibraryElementTypes = 67  //col:991
    /// BcdLibraryElementTypes = 68  //col:992
    /// </summary> BcdLibraryElementTypes = 69  //col:993
    /// <remarks>0x1600000F</remarks> BcdLibraryElementTypes = 70  //col:994
    BcdLibraryBoolean_TraditionalKsegMappings  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 15)  //col:995
    /// <summary> BcdLibraryElementTypes = 72  //col:996
    /// Indicates whether the boot debugger should be enabled. BcdLibraryElementTypes = 73  //col:997
    /// </summary> BcdLibraryElementTypes = 74  //col:998
    /// <remarks>0x16000010</remarks> BcdLibraryElementTypes = 75  //col:999
    BcdLibraryBoolean_DebuggerEnabled  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 16)  //col:1000
    /// <summary> BcdLibraryElementTypes = 77  //col:1001
    /// Debugger type. The Integer property is one of the values from the BcdLibrary_DebuggerType enumeration. BcdLibraryElementTypes = 78  //col:1002
    /// </summary> BcdLibraryElementTypes = 79  //col:1003
    /// <remarks>0x15000011</remarks> BcdLibraryElementTypes = 80  //col:1004
    BcdLibraryInteger_DebuggerType  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 17)  //col:1005
    /// <summary> BcdLibraryElementTypes = 82  //col:1006
    /// I/O port address for the serial debugger. BcdLibraryElementTypes = 83  //col:1007
    /// </summary> BcdLibraryElementTypes = 84  //col:1008
    /// <remarks>0x15000012</remarks> BcdLibraryElementTypes = 85  //col:1009
    BcdLibraryInteger_SerialDebuggerPortAddress  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 18)  //col:1010
    /// <summary> BcdLibraryElementTypes = 87  //col:1011
    /// Serial port number for serial debugging. BcdLibraryElementTypes = 88  //col:1012
    /// If this value is not specified the default is specified by the DBGP ACPI table settings. BcdLibraryElementTypes = 89  //col:1013
    /// </summary> BcdLibraryElementTypes = 90  //col:1014
    /// <remarks>0x15000013</remarks> BcdLibraryElementTypes = 91  //col:1015
    BcdLibraryInteger_SerialDebuggerPort  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 19)  //col:1016
    /// <summary> BcdLibraryElementTypes = 93  //col:1017
    /// Baud rate for serial debugging. BcdLibraryElementTypes = 94  //col:1018
    /// </summary> BcdLibraryElementTypes = 95  //col:1019
    /// <remarks>0x15000014</remarks> BcdLibraryElementTypes = 96  //col:1020
    BcdLibraryInteger_SerialDebuggerBaudRate  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 20)  //col:1021
    /// <summary> BcdLibraryElementTypes = 98  //col:1022
    /// Channel number for 1394 debugging. BcdLibraryElementTypes = 99  //col:1023
    /// </summary> BcdLibraryElementTypes = 100  //col:1024
    /// <remarks>0x15000015</remarks> BcdLibraryElementTypes = 101  //col:1025
    BcdLibraryInteger_1394DebuggerChannel  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 21)  //col:1026
    /// <summary> BcdLibraryElementTypes = 103  //col:1027
    /// The target name for the USB debugger. The target name is arbitrary but must match between the debugger and the debug target. BcdLibraryElementTypes = 104  //col:1028
    /// </summary> BcdLibraryElementTypes = 105  //col:1029
    /// <remarks>0x12000016</remarks> BcdLibraryElementTypes = 106  //col:1030
    BcdLibraryString_UsbDebuggerTargetName  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 22)  //col:1031
    /// <summary> BcdLibraryElementTypes = 108  //col:1032
    /// If TRUE the debugger will ignore user mode exceptions and only stop for kernel mode exceptions. BcdLibraryElementTypes = 109  //col:1033
    /// </summary> BcdLibraryElementTypes = 110  //col:1034
    /// <remarks>0x16000017</remarks> BcdLibraryElementTypes = 111  //col:1035
    BcdLibraryBoolean_DebuggerIgnoreUsermodeExceptions  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 23)  //col:1036
    /// <summary> BcdLibraryElementTypes = 113  //col:1037
    /// Indicates the debugger start policy. The Integer property is one of the values from the BcdLibrary_DebuggerStartPolicy enumeration. BcdLibraryElementTypes = 114  //col:1038
    /// </summary> BcdLibraryElementTypes = 115  //col:1039
    /// <remarks>0x15000018</remarks> BcdLibraryElementTypes = 116  //col:1040
    BcdLibraryInteger_DebuggerStartPolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 24)  //col:1041
    /// <summary> BcdLibraryElementTypes = 118  //col:1042
    /// Defines the PCI bus device and function numbers of the debugging device. For example 1.5.0 describes the debugging device on bus 1 device 5 function 0. BcdLibraryElementTypes = 119  //col:1043
    /// </summary> BcdLibraryElementTypes = 120  //col:1044
    /// <remarks>0x12000019</remarks> BcdLibraryElementTypes = 121  //col:1045
    BcdLibraryString_DebuggerBusParameters  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 25)  //col:1046
    /// <summary> BcdLibraryElementTypes = 123  //col:1047
    /// Defines the host IP address for the network debugger. BcdLibraryElementTypes = 124  //col:1048
    /// </summary> BcdLibraryElementTypes = 125  //col:1049
    /// <remarks>0x1500001A</remarks> BcdLibraryElementTypes = 126  //col:1050
    BcdLibraryInteger_DebuggerNetHostIP  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 26)  //col:1051
    /// <summary> BcdLibraryElementTypes = 128  //col:1052
    /// Defines the network port for the network debugger. BcdLibraryElementTypes = 129  //col:1053
    /// </summary> BcdLibraryElementTypes = 130  //col:1054
    /// <remarks>0x1500001B</remarks> BcdLibraryElementTypes = 131  //col:1055
    BcdLibraryInteger_DebuggerNetPort  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 27)  //col:1056
    /// <summary> BcdLibraryElementTypes = 133  //col:1057
    /// Controls the use of DHCP by the network debugger. Setting this to false causes the OS to only use link-local addresses. BcdLibraryElementTypes = 134  //col:1058
    /// This value is supported starting in Windows 8 and Windows Server 2012. BcdLibraryElementTypes = 135  //col:1059
    /// </summary> BcdLibraryElementTypes = 136  //col:1060
    /// <remarks>0x1600001C</remarks> BcdLibraryElementTypes = 137  //col:1061
    BcdLibraryBoolean_DebuggerNetDhcp  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 28)  //col:1062
    /// <summary> BcdLibraryElementTypes = 139  //col:1063
    /// Holds the key used to encrypt the network debug connection. BcdLibraryElementTypes = 140  //col:1064
    /// This value is supported starting in Windows 8 and Windows Server 2012. BcdLibraryElementTypes = 141  //col:1065
    /// </summary> BcdLibraryElementTypes = 142  //col:1066
    /// <remarks>0x1200001D</remarks> BcdLibraryElementTypes = 143  //col:1067
    BcdLibraryString_DebuggerNetKey  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 29)  //col:1068
    /// <summary> BcdLibraryElementTypes = 145  //col:1069
    ///  BcdLibraryElementTypes = 146  //col:1070
    /// </summary> BcdLibraryElementTypes = 147  //col:1071
    /// <remarks>0x1600001E</remarks> BcdLibraryElementTypes = 148  //col:1072
    BcdLibraryBoolean_DebuggerNetVM  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 30)  //col:1073
    /// <summary> BcdLibraryElementTypes = 150  //col:1074
    ///  BcdLibraryElementTypes = 151  //col:1075
    /// </summary> BcdLibraryElementTypes = 152  //col:1076
    /// <remarks>0x1200001F</remarks> BcdLibraryElementTypes = 153  //col:1077
    BcdLibraryString_DebuggerNetHostIpv6  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 31)  //col:1078
    /// <summary> BcdLibraryElementTypes = 155  //col:1079
    /// Indicates whether EMS redirection should be enabled. BcdLibraryElementTypes = 156  //col:1080
    /// </summary> BcdLibraryElementTypes = 157  //col:1081
    /// <remarks>0x16000020</remarks> BcdLibraryElementTypes = 158  //col:1082
    BcdLibraryBoolean_EmsEnabled  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 32)  //col:1083
    /// <summary> BcdLibraryElementTypes = 160  //col:1084
    /// COM port number for EMS redirection. BcdLibraryElementTypes = 161  //col:1085
    /// </summary> BcdLibraryElementTypes = 162  //col:1086
    /// <remarks>0x15000022</remarks> BcdLibraryElementTypes = 163  //col:1087
    BcdLibraryInteger_EmsPort  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 34)  //col:1088
    /// <summary> BcdLibraryElementTypes = 165  //col:1089
    /// Baud rate for EMS redirection. BcdLibraryElementTypes = 166  //col:1090
    /// </summary> BcdLibraryElementTypes = 167  //col:1091
    /// <remarks>0x15000023</remarks> BcdLibraryElementTypes = 168  //col:1092
    BcdLibraryInteger_EmsBaudRate  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 35)  //col:1093
    /// <summary> BcdLibraryElementTypes = 170  //col:1094
    /// String that is appended to the load options string passed to the kernel to be consumed by kernel-mode components. BcdLibraryElementTypes = 171  //col:1095
    /// This is useful for communicating with kernel-mode components that are not BCD-aware. BcdLibraryElementTypes = 172  //col:1096
    /// </summary> BcdLibraryElementTypes = 173  //col:1097
    /// <remarks>0x12000030</remarks> BcdLibraryElementTypes = 174  //col:1098
    BcdLibraryString_LoadOptionsString  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 48)  //col:1099
    /// <summary> BcdLibraryElementTypes = 176  //col:1100
    ///  BcdLibraryElementTypes = 177  //col:1101
    /// </summary> BcdLibraryElementTypes = 178  //col:1102
    /// <remarks>0x16000031</remarks> BcdLibraryElementTypes = 179  //col:1103
    BcdLibraryBoolean_AttemptNonBcdStart  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 49)  //col:1104
    /// <summary> BcdLibraryElementTypes = 181  //col:1105
    /// Indicates whether the advanced options boot menu (F8) is displayed. BcdLibraryElementTypes = 182  //col:1106
    /// </summary> BcdLibraryElementTypes = 183  //col:1107
    /// <remarks>0x16000040</remarks> BcdLibraryElementTypes = 184  //col:1108
    BcdLibraryBoolean_DisplayAdvancedOptions  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 64)  //col:1109
    /// <summary> BcdLibraryElementTypes = 186  //col:1110
    /// Indicates whether the boot options editor is enabled. BcdLibraryElementTypes = 187  //col:1111
    /// </summary> BcdLibraryElementTypes = 188  //col:1112
    /// <remarks>0x16000041</remarks> BcdLibraryElementTypes = 189  //col:1113
    BcdLibraryBoolean_DisplayOptionsEdit  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 65)  //col:1114
    /// <summary> BcdLibraryElementTypes = 191  //col:1115
    ///  BcdLibraryElementTypes = 192  //col:1116
    /// </summary> BcdLibraryElementTypes = 193  //col:1117
    /// <remarks>0x15000042</remarks> BcdLibraryElementTypes = 194  //col:1118
    BcdLibraryInteger_FVEKeyRingAddress  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 66)  //col:1119
    /// <summary> BcdLibraryElementTypes = 196  //col:1120
    /// Allows a device override for the bootstat.dat log in the boot manager and winload.exe. BcdLibraryElementTypes = 197  //col:1121
    /// </summary> BcdLibraryElementTypes = 198  //col:1122
    /// <remarks>0x11000043</remarks> BcdLibraryElementTypes = 199  //col:1123
    BcdLibraryDevice_BsdLogDevice  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 67)  //col:1124
    /// <summary> BcdLibraryElementTypes = 201  //col:1125
    /// Allows a path override for the bootstat.dat log file in the boot manager and winload.exe. BcdLibraryElementTypes = 202  //col:1126
    /// </summary> BcdLibraryElementTypes = 203  //col:1127
    /// <remarks>0x12000044</remarks> BcdLibraryElementTypes = 204  //col:1128
    BcdLibraryString_BsdLogPath  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 68)  //col:1129
    /// <summary> BcdLibraryElementTypes = 206  //col:1130
    /// Indicates whether graphics mode is disabled and boot applications must use text mode display. BcdLibraryElementTypes = 207  //col:1131
    /// </summary> BcdLibraryElementTypes = 208  //col:1132
    /// <remarks>0x16000045</remarks> BcdLibraryElementTypes = 209  //col:1133
    BcdLibraryBoolean_BsdPreserveLog  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 69)  //col:1134
    /// <summary> BcdLibraryElementTypes = 211  //col:1135
    ///  BcdLibraryElementTypes = 212  //col:1136
    /// </summary> BcdLibraryElementTypes = 213  //col:1137
    /// <remarks>0x16000046</remarks> BcdLibraryElementTypes = 214  //col:1138
    BcdLibraryBoolean_GraphicsModeDisabled  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 70)  //col:1139
    /// <summary> BcdLibraryElementTypes = 216  //col:1140
    /// Indicates the access policy for PCI configuration space. BcdLibraryElementTypes = 217  //col:1141
    /// </summary> BcdLibraryElementTypes = 218  //col:1142
    /// <remarks>0x15000047</remarks> BcdLibraryElementTypes = 219  //col:1143
    BcdLibraryInteger_ConfigAccessPolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 71)  //col:1144
    /// <summary> BcdLibraryElementTypes = 221  //col:1145
    /// Disables integrity checks. BcdLibraryElementTypes = 222  //col:1146
    /// Cannot be set when secure boot is enabled. BcdLibraryElementTypes = 223  //col:1147
    /// This value is ignored by Windows 7 and Windows 8. BcdLibraryElementTypes = 224  //col:1148
    /// </summary> BcdLibraryElementTypes = 225  //col:1149
    /// <remarks>0x16000048</remarks> BcdLibraryElementTypes = 226  //col:1150
    BcdLibraryBoolean_DisableIntegrityChecks  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 72)  //col:1151
    /// <summary> BcdLibraryElementTypes = 228  //col:1152
    /// Indicates whether the test code signing certificate is supported. BcdLibraryElementTypes = 229  //col:1153
    /// </summary> BcdLibraryElementTypes = 230  //col:1154
    /// <remarks>0x16000049</remarks> BcdLibraryElementTypes = 231  //col:1155
    BcdLibraryBoolean_AllowPrereleaseSignatures  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 73)  //col:1156
    /// <summary> BcdLibraryElementTypes = 233  //col:1157
    /// Overrides the default location of the boot fonts. BcdLibraryElementTypes = 234  //col:1158
    /// </summary> BcdLibraryElementTypes = 235  //col:1159
    /// <remarks>0x1200004A</remarks> BcdLibraryElementTypes = 236  //col:1160
    BcdLibraryString_FontPath  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 74)  //col:1161
    /// <summary> BcdLibraryElementTypes = 238  //col:1162
    ///  BcdLibraryElementTypes = 239  //col:1163
    /// </summary> BcdLibraryElementTypes = 240  //col:1164
    /// <remarks>0x1500004B</remarks> BcdLibraryElementTypes = 241  //col:1165
    BcdLibraryInteger_SiPolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 75)  //col:1166
    /// <summary> BcdLibraryElementTypes = 243  //col:1167
    /// This value (if present) should not be modified. BcdLibraryElementTypes = 244  //col:1168
    /// </summary> BcdLibraryElementTypes = 245  //col:1169
    /// <remarks>0x1500004C</remarks> BcdLibraryElementTypes = 246  //col:1170
    BcdLibraryInteger_FveBandId  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 76)  //col:1171
    /// <summary> BcdLibraryElementTypes = 248  //col:1172
    /// Specifies that legacy BIOS systems should use INT 16h Function 10h for console input instead of INT 16h Function 0h. BcdLibraryElementTypes = 249  //col:1173
    /// </summary> BcdLibraryElementTypes = 250  //col:1174
    /// <remarks>0x16000050</remarks> BcdLibraryElementTypes = 251  //col:1175
    BcdLibraryBoolean_ConsoleExtendedInput  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 80)  //col:1176
    /// <summary> BcdLibraryElementTypes = 253  //col:1177
    ///  BcdLibraryElementTypes = 254  //col:1178
    /// </summary> BcdLibraryElementTypes = 255  //col:1179
    /// <remarks>0x15000051</remarks> BcdLibraryElementTypes = 256  //col:1180
    BcdLibraryInteger_InitialConsoleInput  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 81)  //col:1181
    /// <summary> BcdLibraryElementTypes = 258  //col:1182
    /// Forces a specific graphics resolution at boot. BcdLibraryElementTypes = 259  //col:1183
    /// Possible values include GraphicsResolution1024x768 (0) GraphicsResolution800x600 (1) and GraphicsResolution1024x600 (2). BcdLibraryElementTypes = 260  //col:1184
    /// </summary> BcdLibraryElementTypes = 261  //col:1185
    /// <remarks>0x15000052</remarks> BcdLibraryElementTypes = 262  //col:1186
    BcdLibraryInteger_GraphicsResolution  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 82)  //col:1187
    /// <summary> BcdLibraryElementTypes = 264  //col:1188
    /// If enabled specifies that boot error screens are not shown when OS launch errors occur and the system is reset rather than exiting directly back to the firmware. BcdLibraryElementTypes = 265  //col:1189
    /// </summary> BcdLibraryElementTypes = 266  //col:1190
    /// <remarks>0x16000053</remarks> BcdLibraryElementTypes = 267  //col:1191
    BcdLibraryBoolean_RestartOnFailure  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 83)  //col:1192
    /// <summary> BcdLibraryElementTypes = 269  //col:1193
    /// Forces highest available graphics resolution at boot. BcdLibraryElementTypes = 270  //col:1194
    /// This value can only be used on UEFI systems. BcdLibraryElementTypes = 271  //col:1195
    /// This value is supported starting in Windows 8 and Windows Server 2012. BcdLibraryElementTypes = 272  //col:1196
    /// </summary> BcdLibraryElementTypes = 273  //col:1197
    /// <remarks>0x16000054</remarks> BcdLibraryElementTypes = 274  //col:1198
    BcdLibraryBoolean_GraphicsForceHighestMode  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 84)  //col:1199
    /// <summary> BcdLibraryElementTypes = 276  //col:1200
    /// This setting is used to differentiate between the Windows 7 and Windows 8 implementations of UEFI. BcdLibraryElementTypes = 277  //col:1201
    /// Do not modify this setting. BcdLibraryElementTypes = 278  //col:1202
    /// If this setting is removed from a Windows 8 installation it will not boot. BcdLibraryElementTypes = 279  //col:1203
    /// If this setting is added to a Windows 7 installation it will not boot. BcdLibraryElementTypes = 280  //col:1204
    /// </summary> BcdLibraryElementTypes = 281  //col:1205
    /// <remarks>0x16000060</remarks> BcdLibraryElementTypes = 282  //col:1206
    BcdLibraryBoolean_IsolatedExecutionContext  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 96)  //col:1207
    /// <summary> BcdLibraryElementTypes = 284  //col:1208
    /// This setting disables the progress bar and default Windows logo. If a custom text string has been defined it is also disabled by this setting. BcdLibraryElementTypes = 285  //col:1209
    /// The Integer property is one of the values from the BcdLibrary_UxDisplayMessageType enumeration. BcdLibraryElementTypes = 286  //col:1210
    /// </summary> BcdLibraryElementTypes = 287  //col:1211
    /// <remarks>0x15000065</remarks> BcdLibraryElementTypes = 288  //col:1212
    BcdLibraryInteger_BootUxDisplayMessage  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 101)  //col:1213
    /// <summary> BcdLibraryElementTypes = 290  //col:1214
    ///  BcdLibraryElementTypes = 291  //col:1215
    /// </summary> BcdLibraryElementTypes = 292  //col:1216
    /// <remarks>0x15000066</remarks> BcdLibraryElementTypes = 293  //col:1217
    BcdLibraryInteger_BootUxDisplayMessageOverride  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 102)  //col:1218
    /// <summary> BcdLibraryElementTypes = 295  //col:1219
    /// This setting disables the boot logo. BcdLibraryElementTypes = 296  //col:1220
    /// </summary> BcdLibraryElementTypes = 297  //col:1221
    /// <remarks>0x16000067</remarks> BcdLibraryElementTypes = 298  //col:1222
    BcdLibraryBoolean_BootUxLogoDisable  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 103)  //col:1223
    /// <summary> BcdLibraryElementTypes = 300  //col:1224
    /// This setting disables the boot status text. BcdLibraryElementTypes = 301  //col:1225
    /// </summary> BcdLibraryElementTypes = 302  //col:1226
    /// <remarks>0x16000068</remarks> BcdLibraryElementTypes = 303  //col:1227
    BcdLibraryBoolean_BootUxTextDisable  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 104)  //col:1228
    /// <summary> BcdLibraryElementTypes = 305  //col:1229
    /// This setting disables the boot progress bar. BcdLibraryElementTypes = 306  //col:1230
    /// </summary> BcdLibraryElementTypes = 307  //col:1231
    /// <remarks>0x16000069</remarks> BcdLibraryElementTypes = 308  //col:1232
    BcdLibraryBoolean_BootUxProgressDisable  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 105)  //col:1233
    /// <summary> BcdLibraryElementTypes = 310  //col:1234
    /// This setting disables the boot transition fading. BcdLibraryElementTypes = 311  //col:1235
    /// </summary> BcdLibraryElementTypes = 312  //col:1236
    /// <remarks>0x1600006A</remarks> BcdLibraryElementTypes = 313  //col:1237
    BcdLibraryBoolean_BootUxFadeDisable  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 106)  //col:1238
    /// <summary> BcdLibraryElementTypes = 315  //col:1239
    ///  BcdLibraryElementTypes = 316  //col:1240
    /// </summary> BcdLibraryElementTypes = 317  //col:1241
    /// <remarks>0x1600006B</remarks> BcdLibraryElementTypes = 318  //col:1242
    BcdLibraryBoolean_BootUxReservePoolDebug  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 107)  //col:1243
    /// <summary> BcdLibraryElementTypes = 320  //col:1244
    ///  BcdLibraryElementTypes = 321  //col:1245
    /// </summary> BcdLibraryElementTypes = 322  //col:1246
    /// <remarks>0x1600006C</remarks> BcdLibraryElementTypes = 323  //col:1247
    BcdLibraryBoolean_BootUxDisable  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 108)  //col:1248
    /// <summary> BcdLibraryElementTypes = 325  //col:1249
    ///  BcdLibraryElementTypes = 326  //col:1250
    /// </summary> BcdLibraryElementTypes = 327  //col:1251
    /// <remarks>0x1500006D</remarks> BcdLibraryElementTypes = 328  //col:1252
    BcdLibraryInteger_BootUxFadeFrames  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 109)  //col:1253
    /// <summary> BcdLibraryElementTypes = 330  //col:1254
    ///  BcdLibraryElementTypes = 331  //col:1255
    /// </summary> BcdLibraryElementTypes = 332  //col:1256
    /// <remarks>0x1600006E</remarks> BcdLibraryElementTypes = 333  //col:1257
    BcdLibraryBoolean_BootUxDumpStats  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 110)  //col:1258
    /// <summary> BcdLibraryElementTypes = 335  //col:1259
    ///  BcdLibraryElementTypes = 336  //col:1260
    /// </summary> BcdLibraryElementTypes = 337  //col:1261
    /// <remarks>0x1600006F</remarks> BcdLibraryElementTypes = 338  //col:1262
    BcdLibraryBoolean_BootUxShowStats  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 111)  //col:1263
    /// <summary> BcdLibraryElementTypes = 340  //col:1264
    ///  BcdLibraryElementTypes = 341  //col:1265
    /// </summary> BcdLibraryElementTypes = 342  //col:1266
    /// <remarks>0x16000071</remarks> BcdLibraryElementTypes = 343  //col:1267
    BcdLibraryBoolean_MultiBootSystem  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 113)  //col:1268
    /// <summary> BcdLibraryElementTypes = 345  //col:1269
    ///  BcdLibraryElementTypes = 346  //col:1270
    /// </summary> BcdLibraryElementTypes = 347  //col:1271
    /// <remarks>0x16000072</remarks> BcdLibraryElementTypes = 348  //col:1272
    BcdLibraryBoolean_ForceNoKeyboard  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 114)  //col:1273
    /// <summary> BcdLibraryElementTypes = 350  //col:1274
    ///  BcdLibraryElementTypes = 351  //col:1275
    /// </summary> BcdLibraryElementTypes = 352  //col:1276
    /// <remarks>0x15000073</remarks> BcdLibraryElementTypes = 353  //col:1277
    BcdLibraryInteger_AliasWindowsKey  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 115)  //col:1278
    /// <summary> BcdLibraryElementTypes = 355  //col:1279
    /// Disables the 1-minute timer that triggers shutdown on boot error screens and the F8 menu on UEFI systems. BcdLibraryElementTypes = 356  //col:1280
    /// </summary> BcdLibraryElementTypes = 357  //col:1281
    /// <remarks>0x16000074</remarks> BcdLibraryElementTypes = 358  //col:1282
    BcdLibraryBoolean_BootShutdownDisabled  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 116)  //col:1283
    /// <summary> BcdLibraryElementTypes = 360  //col:1284
    ///  BcdLibraryElementTypes = 361  //col:1285
    /// </summary> BcdLibraryElementTypes = 362  //col:1286
    /// <remarks>0x15000075</remarks> BcdLibraryElementTypes = 363  //col:1287
    BcdLibraryInteger_PerformanceFrequency  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 117)  //col:1288
    /// <summary> BcdLibraryElementTypes = 365  //col:1289
    ///  BcdLibraryElementTypes = 366  //col:1290
    /// </summary> BcdLibraryElementTypes = 367  //col:1291
    /// <remarks>0x15000076</remarks> BcdLibraryElementTypes = 368  //col:1292
    BcdLibraryInteger_SecurebootRawPolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 118)  //col:1293
    /// <summary> BcdLibraryElementTypes = 370  //col:1294
    /// Indicates whether or not an in-memory BCD setting passed between boot apps will trigger BitLocker recovery. BcdLibraryElementTypes = 371  //col:1295
    /// This value should not be modified as it could trigger a BitLocker recovery action. BcdLibraryElementTypes = 372  //col:1296
    /// </summary> BcdLibraryElementTypes = 373  //col:1297
    /// <remarks>0x17000077</remarks> BcdLibraryElementTypes = 374  //col:1298
    BcdLibraryIntegerList_AllowedInMemorySettings  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 119)  //col:1299
    /// <summary> BcdLibraryElementTypes = 376  //col:1300
    ///  BcdLibraryElementTypes = 377  //col:1301
    /// </summary> BcdLibraryElementTypes = 378  //col:1302
    /// <remarks>0x15000079</remarks> BcdLibraryElementTypes = 379  //col:1303
    BcdLibraryInteger_BootUxBitmapTransitionTime  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 121)  //col:1304
    /// <summary> BcdLibraryElementTypes = 381  //col:1305
    ///  BcdLibraryElementTypes = 382  //col:1306
    /// </summary> BcdLibraryElementTypes = 383  //col:1307
    /// <remarks>0x1600007A</remarks> BcdLibraryElementTypes = 384  //col:1308
    BcdLibraryBoolean_TwoBootImages  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 122)  //col:1309
    /// <summary> BcdLibraryElementTypes = 386  //col:1310
    /// Force the use of FIPS cryptography checks on boot applications. BcdLibraryElementTypes = 387  //col:1311
    /// BcdLibraryBoolean_ForceFipsCrypto is documented with wrong value 0x16000079 BcdLibraryElementTypes = 388  //col:1312
    /// </summary> BcdLibraryElementTypes = 389  //col:1313
    /// <remarks>0x1600007B</remarks> BcdLibraryElementTypes = 390  //col:1314
    BcdLibraryBoolean_ForceFipsCrypto  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 123)  //col:1315
    /// <summary> BcdLibraryElementTypes = 392  //col:1316
    ///  BcdLibraryElementTypes = 393  //col:1317
    /// </summary> BcdLibraryElementTypes = 394  //col:1318
    /// <remarks>0x1500007D</remarks> BcdLibraryElementTypes = 395  //col:1319
    BcdLibraryInteger_BootErrorUx  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 125)  //col:1320
    /// <summary> BcdLibraryElementTypes = 397  //col:1321
    ///  BcdLibraryElementTypes = 398  //col:1322
    /// </summary> BcdLibraryElementTypes = 399  //col:1323
    /// <remarks>0x1600007E</remarks> BcdLibraryElementTypes = 400  //col:1324
    BcdLibraryBoolean_AllowFlightSignatures  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 126)  //col:1325
    /// <summary> BcdLibraryElementTypes = 402  //col:1326
    ///  BcdLibraryElementTypes = 403  //col:1327
    /// </summary> BcdLibraryElementTypes = 404  //col:1328
    /// <remarks>0x1500007F</remarks> BcdLibraryElementTypes = 405  //col:1329
    BcdLibraryInteger_BootMeasurementLogFormat  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 127)  //col:1330
    /// <summary> BcdLibraryElementTypes = 407  //col:1331
    ///  BcdLibraryElementTypes = 408  //col:1332
    /// </summary> BcdLibraryElementTypes = 409  //col:1333
    /// <remarks>0x15000080</remarks> BcdLibraryElementTypes = 410  //col:1334
    BcdLibraryInteger_DisplayRotation  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 128)  //col:1335
    /// <summary> BcdLibraryElementTypes = 412  //col:1336
    ///  BcdLibraryElementTypes = 413  //col:1337
    /// </summary> BcdLibraryElementTypes = 414  //col:1338
    /// <remarks>0x15000081</remarks> BcdLibraryElementTypes = 415  //col:1339
    BcdLibraryInteger_LogControl  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 129)  //col:1340
    /// <summary> BcdLibraryElementTypes = 417  //col:1341
    ///  BcdLibraryElementTypes = 418  //col:1342
    /// </summary> BcdLibraryElementTypes = 419  //col:1343
    /// <remarks>0x16000082</remarks> BcdLibraryElementTypes = 420  //col:1344
    BcdLibraryBoolean_NoFirmwareSync  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 130)  //col:1345
    /// <summary> BcdLibraryElementTypes = 422  //col:1346
    ///  BcdLibraryElementTypes = 423  //col:1347
    /// </summary> BcdLibraryElementTypes = 424  //col:1348
    /// <remarks>0x11000084</remarks> BcdLibraryElementTypes = 425  //col:1349
    BcdLibraryDevice_WindowsSystemDevice  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 132)  //col:1350
    /// <summary> BcdLibraryElementTypes = 427  //col:1351
    ///  BcdLibraryElementTypes = 428  //col:1352
    /// </summary> BcdLibraryElementTypes = 429  //col:1353
    /// <remarks>0x16000087</remarks> BcdLibraryElementTypes = 430  //col:1354
    BcdLibraryBoolean_NumLockOn  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 135)  //col:1355
    /// <summary> BcdLibraryElementTypes = 432  //col:1356
    ///  BcdLibraryElementTypes = 433  //col:1357
    /// </summary> BcdLibraryElementTypes = 434  //col:1358
    /// <remarks>0x12000088</remarks> BcdLibraryElementTypes = 435  //col:1359
    BcdLibraryString_AdditionalCiPolicy  BcdLibraryElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY BCD_ELEMENT_DATATYPE_FORMAT_STRING 136)  //col:1360
)


type     /// <summary> uint32
const(
    /// <summary> BcdTemplateElementTypes = 1  //col:1365
    ///  BcdTemplateElementTypes = 2  //col:1366
    /// </summary> BcdTemplateElementTypes = 3  //col:1367
    /// <remarks>0x45000001</remarks> BcdTemplateElementTypes = 4  //col:1368
    BcdSetupInteger_DeviceType  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 1)  //col:1369
    /// <summary> BcdTemplateElementTypes = 6  //col:1370
    ///  BcdTemplateElementTypes = 7  //col:1371
    /// </summary> BcdTemplateElementTypes = 8  //col:1372
    /// <remarks>0x42000002</remarks> BcdTemplateElementTypes = 9  //col:1373
    BcdSetupString_ApplicationRelativePath  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_STRING 2)  //col:1374
    /// <summary> BcdTemplateElementTypes = 11  //col:1375
    ///  BcdTemplateElementTypes = 12  //col:1376
    /// </summary> BcdTemplateElementTypes = 13  //col:1377
    /// <remarks>0x42000003</remarks> BcdTemplateElementTypes = 14  //col:1378
    BcdSetupString_RamdiskDeviceRelativePath  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_STRING 3)  //col:1379
    /// <summary> BcdTemplateElementTypes = 16  //col:1380
    ///  BcdTemplateElementTypes = 17  //col:1381
    /// </summary> BcdTemplateElementTypes = 18  //col:1382
    /// <remarks>0x46000004</remarks> BcdTemplateElementTypes = 19  //col:1383
    BcdSetupBoolean_OmitOsLoaderElements  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 4)  //col:1384
    /// <summary> BcdTemplateElementTypes = 21  //col:1385
    ///  BcdTemplateElementTypes = 22  //col:1386
    /// </summary> BcdTemplateElementTypes = 23  //col:1387
    /// <remarks>0x47000006</remarks> BcdTemplateElementTypes = 24  //col:1388
    BcdSetupIntegerList_ElementsToMigrateList  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST 6)  //col:1389
    /// <summary> BcdTemplateElementTypes = 26  //col:1390
    ///  BcdTemplateElementTypes = 27  //col:1391
    /// </summary> BcdTemplateElementTypes = 28  //col:1392
    /// <remarks>0x46000010</remarks> BcdTemplateElementTypes = 29  //col:1393
    BcdSetupBoolean_RecoveryOs  BcdTemplateElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 16)  //col:1394
)


type     /// <summary> uint32
const(
    /// <summary> BcdOSLoader_NxPolicy = 1  //col:1402
    /// The no-execute page protection is off by default. BcdOSLoader_NxPolicy = 2  //col:1403
    /// </summary> BcdOSLoader_NxPolicy = 3  //col:1404
    NxPolicyOptIn  BcdOSLoader_NxPolicy =  0  //col:1405
    /// <summary> BcdOSLoader_NxPolicy = 5  //col:1406
    /// The no-execute page protection is on by default. BcdOSLoader_NxPolicy = 6  //col:1407
    /// </summary> BcdOSLoader_NxPolicy = 7  //col:1408
    NxPolicyOptOut  BcdOSLoader_NxPolicy =  1  //col:1409
    /// <summary> BcdOSLoader_NxPolicy = 9  //col:1410
    /// The no-execute page protection is always off. BcdOSLoader_NxPolicy = 10  //col:1411
    /// </summary> BcdOSLoader_NxPolicy = 11  //col:1412
    NxPolicyAlwaysOff  BcdOSLoader_NxPolicy =  2  //col:1413
    /// <summary> BcdOSLoader_NxPolicy = 13  //col:1414
    /// The no-execute page protection is always on. BcdOSLoader_NxPolicy = 14  //col:1415
    /// </summary> BcdOSLoader_NxPolicy = 15  //col:1416
    NxPolicyAlwaysOn  BcdOSLoader_NxPolicy =  3  //col:1417
)


type     /// <summary> uint32
const(
    /// <summary> BcdOSLoader_PAEPolicy = 1  //col:1425
    /// Enable PAE if hot-pluggable memory is defined above 4GB. BcdOSLoader_PAEPolicy = 2  //col:1426
    /// </summary> BcdOSLoader_PAEPolicy = 3  //col:1427
    PaePolicyDefault  BcdOSLoader_PAEPolicy =  0  //col:1428
    /// <summary> BcdOSLoader_PAEPolicy = 5  //col:1429
    /// PAE is enabled. BcdOSLoader_PAEPolicy = 6  //col:1430
    /// </summary> BcdOSLoader_PAEPolicy = 7  //col:1431
    PaePolicyForceEnable  BcdOSLoader_PAEPolicy =  1  //col:1432
    /// <summary> BcdOSLoader_PAEPolicy = 9  //col:1433
    /// PAE is disabled. BcdOSLoader_PAEPolicy = 10  //col:1434
    /// </summary> BcdOSLoader_PAEPolicy = 11  //col:1435
    PaePolicyForceDisable  BcdOSLoader_PAEPolicy =  2  //col:1436
)


type     /// <summary> uint32
const(
    /// <summary> BcdOSLoader_BootStatusPolicy = 1  //col:1441
    /// Display all boot failures. BcdOSLoader_BootStatusPolicy = 2  //col:1442
    /// </summary> BcdOSLoader_BootStatusPolicy = 3  //col:1443
    BootStatusPolicyDisplayAllFailures  BcdOSLoader_BootStatusPolicy =  0  //col:1444
    /// <summary> BcdOSLoader_BootStatusPolicy = 5  //col:1445
    /// Ignore all boot failures. BcdOSLoader_BootStatusPolicy = 6  //col:1446
    /// </summary> BcdOSLoader_BootStatusPolicy = 7  //col:1447
    BootStatusPolicyIgnoreAllFailures  BcdOSLoader_BootStatusPolicy =  1  //col:1448
    /// <summary> BcdOSLoader_BootStatusPolicy = 9  //col:1449
    /// Ignore all shutdown failures. BcdOSLoader_BootStatusPolicy = 10  //col:1450
    /// </summary> BcdOSLoader_BootStatusPolicy = 11  //col:1451
    BootStatusPolicyIgnoreShutdownFailures  BcdOSLoader_BootStatusPolicy =  2  //col:1452
    /// <summary> BcdOSLoader_BootStatusPolicy = 13  //col:1453
    /// Ignore all boot failures. BcdOSLoader_BootStatusPolicy = 14  //col:1454
    /// </summary> BcdOSLoader_BootStatusPolicy = 15  //col:1455
    BootStatusPolicyIgnoreBootFailures  BcdOSLoader_BootStatusPolicy =  3  //col:1456
    /// <summary> BcdOSLoader_BootStatusPolicy = 17  //col:1457
    /// Ignore checkpoint failures. BcdOSLoader_BootStatusPolicy = 18  //col:1458
    /// </summary> BcdOSLoader_BootStatusPolicy = 19  //col:1459
    BootStatusPolicyIgnoreCheckpointFailures  BcdOSLoader_BootStatusPolicy =  4  //col:1460
    /// <summary> BcdOSLoader_BootStatusPolicy = 21  //col:1461
    /// Display shutdown failures. BcdOSLoader_BootStatusPolicy = 22  //col:1462
    /// </summary> BcdOSLoader_BootStatusPolicy = 23  //col:1463
    BootStatusPolicyDisplayShutdownFailures  BcdOSLoader_BootStatusPolicy =  5  //col:1464
    /// <summary> BcdOSLoader_BootStatusPolicy = 25  //col:1465
    /// Display boot failures. BcdOSLoader_BootStatusPolicy = 26  //col:1466
    /// </summary> BcdOSLoader_BootStatusPolicy = 27  //col:1467
    BootStatusPolicyDisplayBootFailures  BcdOSLoader_BootStatusPolicy =  6  //col:1468
    /// <summary> BcdOSLoader_BootStatusPolicy = 29  //col:1469
    /// Display checkpoint failures. BcdOSLoader_BootStatusPolicy = 30  //col:1470
    /// </summary> BcdOSLoader_BootStatusPolicy = 31  //col:1471
    BootStatusPolicyDisplayCheckpointFailures  BcdOSLoader_BootStatusPolicy =  7  //col:1472
)


type     /// <summary> uint32
const(
    /// <summary> BcdOSLoaderElementTypes = 1  //col:1478
    /// The device on which the operating system resides. BcdOSLoaderElementTypes = 2  //col:1479
    /// </summary> BcdOSLoaderElementTypes = 3  //col:1480
    /// <remarks>0x21000001</remarks> BcdOSLoaderElementTypes = 4  //col:1481
    BcdOSLoaderDevice_OSDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 1)  //col:1482
    /// <summary> BcdOSLoaderElementTypes = 6  //col:1483
    /// The file path to the operating system (%SystemRoot% minus the volume). BcdOSLoaderElementTypes = 7  //col:1484
    /// </summary> BcdOSLoaderElementTypes = 8  //col:1485
    /// <remarks>0x22000002</remarks> BcdOSLoaderElementTypes = 9  //col:1486
    BcdOSLoaderString_SystemRoot  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 2)  //col:1487
    /// <summary> BcdOSLoaderElementTypes = 11  //col:1488
    /// The resume application associated with the operating system. BcdOSLoaderElementTypes = 12  //col:1489
    /// </summary> BcdOSLoaderElementTypes = 13  //col:1490
    /// <remarks>0x23000003</remarks> BcdOSLoaderElementTypes = 14  //col:1491
    BcdOSLoaderObject_AssociatedResumeObject  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_OBJECT 3)  //col:1492
    /// <summary> BcdOSLoaderElementTypes = 16  //col:1493
    ///  BcdOSLoaderElementTypes = 17  //col:1494
    /// </summary> BcdOSLoaderElementTypes = 18  //col:1495
    /// <remarks>0x26000004</remarks> BcdOSLoaderElementTypes = 19  //col:1496
    BcdOSLoaderBoolean_StampDisks  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 4)  //col:1497
    /// <summary> BcdOSLoaderElementTypes = 21  //col:1498
    /// Indicates whether the operating system loader should determine the kernel and HAL to load based on the platform features. BcdOSLoaderElementTypes = 22  //col:1499
    /// </summary> BcdOSLoaderElementTypes = 23  //col:1500
    /// <remarks>0x26000010</remarks> BcdOSLoaderElementTypes = 24  //col:1501
    BcdOSLoaderBoolean_DetectKernelAndHal  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 16)  //col:1502
    /// <summary> BcdOSLoaderElementTypes = 26  //col:1503
    /// The kernel to be loaded by the operating system loader. This value overrides the default kernel. BcdOSLoaderElementTypes = 27  //col:1504
    /// </summary> BcdOSLoaderElementTypes = 28  //col:1505
    /// <remarks>0x22000011</remarks> BcdOSLoaderElementTypes = 29  //col:1506
    BcdOSLoaderString_KernelPath  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 17)  //col:1507
    /// <summary> BcdOSLoaderElementTypes = 31  //col:1508
    /// The HAL to be loaded by the operating system loader. This value overrides the default HAL. BcdOSLoaderElementTypes = 32  //col:1509
    /// </summary> BcdOSLoaderElementTypes = 33  //col:1510
    /// <remarks>0x22000012</remarks> BcdOSLoaderElementTypes = 34  //col:1511
    BcdOSLoaderString_HalPath  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 18)  //col:1512
    /// <summary> BcdOSLoaderElementTypes = 36  //col:1513
    /// The transport DLL to be loaded by the operating system loader. This value overrides the default Kdcom.dll. BcdOSLoaderElementTypes = 37  //col:1514
    /// </summary> BcdOSLoaderElementTypes = 38  //col:1515
    /// <remarks>0x22000013</remarks> BcdOSLoaderElementTypes = 39  //col:1516
    BcdOSLoaderString_DbgTransportPath  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 19)  //col:1517
    /// <summary> BcdOSLoaderElementTypes = 41  //col:1518
    /// The no-execute page protection policy. The Integer property is one of the values from the BcdOSLoader_NxPolicy enumeration. BcdOSLoaderElementTypes = 42  //col:1519
    /// </summary> BcdOSLoaderElementTypes = 43  //col:1520
    /// <remarks>0x25000020</remarks> BcdOSLoaderElementTypes = 44  //col:1521
    BcdOSLoaderInteger_NxPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 32)  //col:1522
    /// <summary> BcdOSLoaderElementTypes = 46  //col:1523
    /// The Physical Address Extension (PAE) policy. The Integer property is one of the values from the BcdOSLoader_PAEPolicy enumeration. BcdOSLoaderElementTypes = 47  //col:1524
    /// </summary> BcdOSLoaderElementTypes = 48  //col:1525
    /// <remarks>0x25000021</remarks> BcdOSLoaderElementTypes = 49  //col:1526
    BcdOSLoaderInteger_PAEPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 33)  //col:1527
    /// <summary> BcdOSLoaderElementTypes = 51  //col:1528
    /// Indicates that the system should be started in Windows Preinstallation Environment (Windows PE) mode. BcdOSLoaderElementTypes = 52  //col:1529
    /// </summary> BcdOSLoaderElementTypes = 53  //col:1530
    /// <remarks>0x26000022</remarks> BcdOSLoaderElementTypes = 54  //col:1531
    BcdOSLoaderBoolean_WinPEMode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 34)  //col:1532
    /// <summary> BcdOSLoaderElementTypes = 56  //col:1533
    /// Indicates that the system should not automatically reboot when it crashes. BcdOSLoaderElementTypes = 57  //col:1534
    /// </summary> BcdOSLoaderElementTypes = 58  //col:1535
    /// <remarks>0x26000024</remarks> BcdOSLoaderElementTypes = 59  //col:1536
    BcdOSLoaderBoolean_DisableCrashAutoReboot  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 36)  //col:1537
    /// <summary> BcdOSLoaderElementTypes = 61  //col:1538
    /// Indicates that the system should use the last-known good settings. BcdOSLoaderElementTypes = 62  //col:1539
    /// </summary> BcdOSLoaderElementTypes = 63  //col:1540
    /// <remarks>0x26000025</remarks> BcdOSLoaderElementTypes = 64  //col:1541
    BcdOSLoaderBoolean_UseLastGoodSettings  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 37)  //col:1542
    /// <summary> BcdOSLoaderElementTypes = 66  //col:1543
    /// BcdOSLoaderElementTypes = 67  //col:1544
    /// </summary> BcdOSLoaderElementTypes = 68  //col:1545
    /// <remarks>0x26000026</remarks> BcdOSLoaderElementTypes = 69  //col:1546
    BcdOSLoaderBoolean_DisableCodeIntegrityChecks  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 38)  //col:1547
    /// <summary> BcdOSLoaderElementTypes = 71  //col:1548
    /// Indicates whether the test code signing certificate is supported. BcdOSLoaderElementTypes = 72  //col:1549
    /// </summary> BcdOSLoaderElementTypes = 73  //col:1550
    /// <remarks>0x26000027</remarks> BcdOSLoaderElementTypes = 74  //col:1551
    BcdOSLoaderBoolean_AllowPrereleaseSignatures  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 39)  //col:1552
    /// <summary> BcdOSLoaderElementTypes = 76  //col:1553
    /// Indicates whether the system should utilize the first 4GB of physical memory. BcdOSLoaderElementTypes = 77  //col:1554
    /// This option requires 5GB of physical memory and on x86 systems it requires PAE to be enabled. BcdOSLoaderElementTypes = 78  //col:1555
    /// </summary> BcdOSLoaderElementTypes = 79  //col:1556
    /// <remarks>0x26000030</remarks> BcdOSLoaderElementTypes = 80  //col:1557
    BcdOSLoaderBoolean_NoLowMemory  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 48)  //col:1558
    /// <summary> BcdOSLoaderElementTypes = 82  //col:1559
    /// The amount of memory the system should ignore. BcdOSLoaderElementTypes = 83  //col:1560
    /// </summary> BcdOSLoaderElementTypes = 84  //col:1561
    /// <remarks>0x25000031</remarks> BcdOSLoaderElementTypes = 85  //col:1562
    BcdOSLoaderInteger_RemoveMemory  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 49)  //col:1563
    /// <summary> BcdOSLoaderElementTypes = 87  //col:1564
    /// The amount of memory that should be utilized by the process address space in bytes. BcdOSLoaderElementTypes = 88  //col:1565
    /// This value should be between 2GB and 3GB. BcdOSLoaderElementTypes = 89  //col:1566
    /// Increasing this value from the default 2GB decreases the amount of virtual address space available to the system and device drivers. BcdOSLoaderElementTypes = 90  //col:1567
    /// </summary> BcdOSLoaderElementTypes = 91  //col:1568
    /// <remarks>0x25000032</remarks> BcdOSLoaderElementTypes = 92  //col:1569
    BcdOSLoaderInteger_IncreaseUserVa  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 50)  //col:1570
    /// <summary> BcdOSLoaderElementTypes = 94  //col:1571
    /// BcdOSLoaderElementTypes = 95  //col:1572
    /// </summary> BcdOSLoaderElementTypes = 96  //col:1573
    /// <remarks>0x25000033</remarks> BcdOSLoaderElementTypes = 97  //col:1574
    BcdOSLoaderInteger_PerformaceDataMemory  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 51)  //col:1575
    /// <summary> BcdOSLoaderElementTypes = 99  //col:1576
    /// Indicates whether the system should use the standard VGA display driver instead of a high-performance display driver. BcdOSLoaderElementTypes = 100  //col:1577
    /// </summary> BcdOSLoaderElementTypes = 101  //col:1578
    /// <remarks>0x26000040</remarks> BcdOSLoaderElementTypes = 102  //col:1579
    BcdOSLoaderBoolean_UseVgaDriver  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 64)  //col:1580
    /// <summary> BcdOSLoaderElementTypes = 104  //col:1581
    /// Indicates whether the system should initialize the VGA driver responsible for displaying simple graphics during the boot process. BcdOSLoaderElementTypes = 105  //col:1582
    /// If not there is no display is presented during the boot process. BcdOSLoaderElementTypes = 106  //col:1583
    /// </summary> BcdOSLoaderElementTypes = 107  //col:1584
    /// <remarks>0x26000041</remarks> BcdOSLoaderElementTypes = 108  //col:1585
    BcdOSLoaderBoolean_DisableBootDisplay  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 65)  //col:1586
    /// <summary> BcdOSLoaderElementTypes = 110  //col:1587
    /// Indicates whether the VGA driver should avoid VESA BIOS calls. BcdOSLoaderElementTypes = 111  //col:1588
    /// Note This value is ignored by Windows 8 and Windows Server 2012. BcdOSLoaderElementTypes = 112  //col:1589
    /// </summary> BcdOSLoaderElementTypes = 113  //col:1590
    /// <remarks>0x26000042</remarks> BcdOSLoaderElementTypes = 114  //col:1591
    BcdOSLoaderBoolean_DisableVesaBios  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 66)  //col:1592
    /// <summary> BcdOSLoaderElementTypes = 116  //col:1593
    /// Disables the use of VGA modes in the OS. BcdOSLoaderElementTypes = 117  //col:1594
    /// </summary> BcdOSLoaderElementTypes = 118  //col:1595
    /// <remarks>0x26000043</remarks> BcdOSLoaderElementTypes = 119  //col:1596
    BcdOSLoaderBoolean_DisableVgaMode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 67)  //col:1597
    /// <summary> BcdOSLoaderElementTypes = 121  //col:1598
    /// Indicates that cluster-mode APIC addressing should be utilized and the value is the maximum number of processors per cluster. BcdOSLoaderElementTypes = 122  //col:1599
    /// </summary> BcdOSLoaderElementTypes = 123  //col:1600
    /// <remarks>0x25000050</remarks> BcdOSLoaderElementTypes = 124  //col:1601
    BcdOSLoaderInteger_ClusterModeAddressing  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 80)  //col:1602
    /// <summary> BcdOSLoaderElementTypes = 126  //col:1603
    /// Indicates whether to enable physical-destination mode for all APIC messages. BcdOSLoaderElementTypes = 127  //col:1604
    /// </summary> BcdOSLoaderElementTypes = 128  //col:1605
    /// <remarks>0x26000051</remarks> BcdOSLoaderElementTypes = 129  //col:1606
    BcdOSLoaderBoolean_UsePhysicalDestination  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 81)  //col:1607
    /// <summary> BcdOSLoaderElementTypes = 131  //col:1608
    /// The maximum number of APIC clusters that should be used by cluster-mode addressing. BcdOSLoaderElementTypes = 132  //col:1609
    /// </summary> BcdOSLoaderElementTypes = 133  //col:1610
    /// <remarks>0x25000052</remarks> BcdOSLoaderElementTypes = 134  //col:1611
    BcdOSLoaderInteger_RestrictApicCluster  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 82)  //col:1612
    /// <summary> BcdOSLoaderElementTypes = 136  //col:1613
    /// BcdOSLoaderElementTypes = 137  //col:1614
    /// </summary> BcdOSLoaderElementTypes = 138  //col:1615
    /// <remarks>0x22000053</remarks> BcdOSLoaderElementTypes = 139  //col:1616
    BcdOSLoaderString_OSLoaderTypeEVStore  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 83)  //col:1617
    /// <summary> BcdOSLoaderElementTypes = 141  //col:1618
    /// Used to force legacy APIC mode even if the processors and chipset support extended APIC mode. BcdOSLoaderElementTypes = 142  //col:1619
    /// </summary> BcdOSLoaderElementTypes = 143  //col:1620
    /// <remarks>0x26000054</remarks> BcdOSLoaderElementTypes = 144  //col:1621
    BcdOSLoaderBoolean_UseLegacyApicMode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 84)  //col:1622
    /// <summary> BcdOSLoaderElementTypes = 146  //col:1623
    /// Enables the use of extended APIC mode if supported. BcdOSLoaderElementTypes = 147  //col:1624
    /// Zero (0) indicates default behavior one (1) indicates that extended APIC mode is disabled and two (2) indicates that extended APIC mode is enabled. BcdOSLoaderElementTypes = 148  //col:1625
    /// The system defaults to using extended APIC mode if available. BcdOSLoaderElementTypes = 149  //col:1626
    /// </summary> BcdOSLoaderElementTypes = 150  //col:1627
    /// <remarks>0x25000055</remarks> BcdOSLoaderElementTypes = 151  //col:1628
    BcdOSLoaderInteger_X2ApicPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 85)  //col:1629
    /// <summary> BcdOSLoaderElementTypes = 153  //col:1630
    /// Indicates whether the operating system should initialize or start non-boot processors. BcdOSLoaderElementTypes = 154  //col:1631
    /// </summary> BcdOSLoaderElementTypes = 155  //col:1632
    /// <remarks>0x26000060</remarks> BcdOSLoaderElementTypes = 156  //col:1633
    BcdOSLoaderBoolean_UseBootProcessorOnly  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 96)  //col:1634
    /// <summary> BcdOSLoaderElementTypes = 158  //col:1635
    /// The maximum number of processors that can be utilized by the system; all other processors are ignored. BcdOSLoaderElementTypes = 159  //col:1636
    /// </summary> BcdOSLoaderElementTypes = 160  //col:1637
    /// <remarks>0x25000061</remarks> BcdOSLoaderElementTypes = 161  //col:1638
    BcdOSLoaderInteger_NumberOfProcessors  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 97)  //col:1639
    /// <summary> BcdOSLoaderElementTypes = 163  //col:1640
    /// Indicates whether the system should use the maximum number of processors. BcdOSLoaderElementTypes = 164  //col:1641
    /// </summary> BcdOSLoaderElementTypes = 165  //col:1642
    /// <remarks>0x26000062</remarks> BcdOSLoaderElementTypes = 166  //col:1643
    BcdOSLoaderBoolean_ForceMaximumProcessors  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 98)  //col:1644
    /// <summary> BcdOSLoaderElementTypes = 168  //col:1645
    /// Indicates whether processor specific configuration flags are to be used. BcdOSLoaderElementTypes = 169  //col:1646
    /// </summary> BcdOSLoaderElementTypes = 170  //col:1647
    /// <remarks>0x25000063</remarks> BcdOSLoaderElementTypes = 171  //col:1648
    BcdOSLoaderBoolean_ProcessorConfigurationFlags  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 99)  //col:1649
    /// <summary> BcdOSLoaderElementTypes = 173  //col:1650
    /// Maximizes the number of groups created when assigning nodes to processor groups. BcdOSLoaderElementTypes = 174  //col:1651
    /// </summary> BcdOSLoaderElementTypes = 175  //col:1652
    /// <remarks>0x26000064</remarks> BcdOSLoaderElementTypes = 176  //col:1653
    BcdOSLoaderBoolean_MaximizeGroupsCreated  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 100)  //col:1654
    /// <summary> BcdOSLoaderElementTypes = 178  //col:1655
    /// This setting makes drivers group aware and can be used to determine improper group usage. BcdOSLoaderElementTypes = 179  //col:1656
    /// </summary> BcdOSLoaderElementTypes = 180  //col:1657
    /// <remarks>0x26000065</remarks> BcdOSLoaderElementTypes = 181  //col:1658
    BcdOSLoaderBoolean_ForceGroupAwareness  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 101)  //col:1659
    /// <summary> BcdOSLoaderElementTypes = 183  //col:1660
    /// Specifies the size of all processor groups. Must be set to a power of 2. BcdOSLoaderElementTypes = 184  //col:1661
    /// </summary> BcdOSLoaderElementTypes = 185  //col:1662
    /// <remarks>0x25000066</remarks> BcdOSLoaderElementTypes = 186  //col:1663
    BcdOSLoaderInteger_GroupSize  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 102)  //col:1664
    /// <summary> BcdOSLoaderElementTypes = 188  //col:1665
    /// Indicates whether the system should use I/O and IRQ resources created by the system firmware instead of using dynamically configured resources. BcdOSLoaderElementTypes = 189  //col:1666
    /// </summary> BcdOSLoaderElementTypes = 190  //col:1667
    /// <remarks>0x26000070</remarks> BcdOSLoaderElementTypes = 191  //col:1668
    BcdOSLoaderInteger_UseFirmwarePciSettings  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 112)  //col:1669
    /// <summary> BcdOSLoaderElementTypes = 193  //col:1670
    /// The PCI Message Signaled Interrupt (MSI) policy. Zero (0) indicates default and one (1) indicates that MSI interrupts are disabled. BcdOSLoaderElementTypes = 194  //col:1671
    /// </summary> BcdOSLoaderElementTypes = 195  //col:1672
    /// <remarks>0x25000071</remarks> BcdOSLoaderElementTypes = 196  //col:1673
    BcdOSLoaderInteger_MsiPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 113)  //col:1674
    /// <summary> BcdOSLoaderElementTypes = 198  //col:1675
    /// Undocumented. Zero (0) indicates default and one (1) indicates that PCI Express is forcefully disabled. BcdOSLoaderElementTypes = 199  //col:1676
    /// </summary> BcdOSLoaderElementTypes = 200  //col:1677
    /// <remarks>0x25000072</remarks> BcdOSLoaderElementTypes = 201  //col:1678
    BcdOSLoaderInteger_PciExpressPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 114)  //col:1679
    /// <summary> BcdOSLoaderElementTypes = 203  //col:1680
    /// The Integer property is one of the values from the BcdLibrary_SafeBoot enumeration. BcdOSLoaderElementTypes = 204  //col:1681
    /// </summary> BcdOSLoaderElementTypes = 205  //col:1682
    /// <remarks>0x25000080</remarks> BcdOSLoaderElementTypes = 206  //col:1683
    BcdOSLoaderInteger_SafeBoot  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 128)  //col:1684
    /// <summary> BcdOSLoaderElementTypes = 208  //col:1685
    /// Indicates whether the system should use the shell specified under the following registry key instead of the default shell: BcdOSLoaderElementTypes = 209  //col:1686
    /// HKEY_LOCAL_MACHINE\SYSTEM\CurrentControlSet\Control\SafeBoot\AlternateShell. BcdOSLoaderElementTypes = 210  //col:1687
    /// </summary> BcdOSLoaderElementTypes = 211  //col:1688
    /// <remarks>0x26000081</remarks> BcdOSLoaderElementTypes = 212  //col:1689
    BcdOSLoaderBoolean_SafeBootAlternateShell  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 129)  //col:1690
    /// <summary> BcdOSLoaderElementTypes = 214  //col:1691
    /// Indicates whether the system should write logging information to %SystemRoot%\Ntbtlog.txt during initialization. BcdOSLoaderElementTypes = 215  //col:1692
    /// </summary> BcdOSLoaderElementTypes = 216  //col:1693
    /// <remarks>0x26000090</remarks> BcdOSLoaderElementTypes = 217  //col:1694
    BcdOSLoaderBoolean_BootLogInitialization  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 144)  //col:1695
    /// <summary> BcdOSLoaderElementTypes = 219  //col:1696
    /// Indicates whether the system should display verbose information. BcdOSLoaderElementTypes = 220  //col:1697
    /// </summary> BcdOSLoaderElementTypes = 221  //col:1698
    /// <remarks>0x26000091</remarks> BcdOSLoaderElementTypes = 222  //col:1699
    BcdOSLoaderBoolean_VerboseObjectLoadMode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 145)  //col:1700
    /// <summary> BcdOSLoaderElementTypes = 224  //col:1701
    /// Indicates whether the kernel debugger should be enabled using the settings in the inherited debugger object. BcdOSLoaderElementTypes = 225  //col:1702
    /// </summary> BcdOSLoaderElementTypes = 226  //col:1703
    /// <remarks>0x260000A0</remarks> BcdOSLoaderElementTypes = 227  //col:1704
    BcdOSLoaderBoolean_KernelDebuggerEnabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 160)  //col:1705
    /// <summary> BcdOSLoaderElementTypes = 229  //col:1706
    /// Indicates whether the HAL should call DbgBreakPoint at the start of HalInitSystem for phase 0 initialization of the kernel. BcdOSLoaderElementTypes = 230  //col:1707
    /// </summary> BcdOSLoaderElementTypes = 231  //col:1708
    /// <remarks>0x260000A1</remarks> BcdOSLoaderElementTypes = 232  //col:1709
    BcdOSLoaderBoolean_DebuggerHalBreakpoint  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 161)  //col:1710
    /// <summary> BcdOSLoaderElementTypes = 234  //col:1711
    /// Forces the use of the platform clock as the system's performance counter. BcdOSLoaderElementTypes = 235  //col:1712
    /// </summary> BcdOSLoaderElementTypes = 236  //col:1713
    /// <remarks>0x260000A2</remarks> BcdOSLoaderElementTypes = 237  //col:1714
    BcdOSLoaderBoolean_UsePlatformClock  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 162)  //col:1715
    /// <summary> BcdOSLoaderElementTypes = 239  //col:1716
    /// Forces the OS to assume the presence of legacy PC devices like CMOS and keyboard controllers. BcdOSLoaderElementTypes = 240  //col:1717
    /// This value should only be used for debugging. BcdOSLoaderElementTypes = 241  //col:1718
    /// </summary> BcdOSLoaderElementTypes = 242  //col:1719
    /// <remarks>0x260000A3</remarks> BcdOSLoaderElementTypes = 243  //col:1720
    BcdOSLoaderBoolean_ForceLegacyPlatform  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 163)  //col:1721
    /// <summary> BcdOSLoaderElementTypes = 245  //col:1722
    /// BcdOSLoaderElementTypes = 246  //col:1723
    /// </summary> BcdOSLoaderElementTypes = 247  //col:1724
    /// <remarks>0x260000A4</remarks> BcdOSLoaderElementTypes = 248  //col:1725
    BcdOSLoaderBoolean_UsePlatformTick  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 164)  //col:1726
    /// <summary> BcdOSLoaderElementTypes = 250  //col:1727
    /// BcdOSLoaderElementTypes = 251  //col:1728
    /// </summary> BcdOSLoaderElementTypes = 252  //col:1729
    /// <remarks>0x260000A5</remarks> BcdOSLoaderElementTypes = 253  //col:1730
    BcdOSLoaderBoolean_DisableDynamicTick  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 165)  //col:1731
    /// <summary> BcdOSLoaderElementTypes = 255  //col:1732
    /// Controls the TSC synchronization policy. Possible values include default (0) legacy (1) or enhanced (2). BcdOSLoaderElementTypes = 256  //col:1733
    /// This value is supported starting in Windows 8 and Windows Server 2012. BcdOSLoaderElementTypes = 257  //col:1734
    /// </summary> BcdOSLoaderElementTypes = 258  //col:1735
    /// <remarks>0x250000A6</remarks> BcdOSLoaderElementTypes = 259  //col:1736
    BcdOSLoaderInteger_TscSyncPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 166)  //col:1737
    /// <summary> BcdOSLoaderElementTypes = 261  //col:1738
    /// Indicates whether EMS should be enabled in the kernel. BcdOSLoaderElementTypes = 262  //col:1739
    /// </summary> BcdOSLoaderElementTypes = 263  //col:1740
    /// <remarks>0x260000B0</remarks> BcdOSLoaderElementTypes = 264  //col:1741
    BcdOSLoaderBoolean_EmsEnabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 176)  //col:1742
    /// <summary> BcdOSLoaderElementTypes = 266  //col:1743
    /// BcdOSLoaderElementTypes = 267  //col:1744
    /// </summary> BcdOSLoaderElementTypes = 268  //col:1745
    /// <remarks>0x250000C0</remarks> BcdOSLoaderElementTypes = 269  //col:1746
    BcdOSLoaderInteger_ForceFailure  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 192)  //col:1747
    /// <summary> BcdOSLoaderElementTypes = 271  //col:1748
    /// Indicates the driver load failure policy. Zero (0) indicates that a failed driver load is fatal and the boot will not continue BcdOSLoaderElementTypes = 272  //col:1749
    /// one (1) indicates that the standard error control is used. BcdOSLoaderElementTypes = 273  //col:1750
    /// </summary> BcdOSLoaderElementTypes = 274  //col:1751
    /// <remarks>0x250000C1</remarks> BcdOSLoaderElementTypes = 275  //col:1752
    BcdOSLoaderInteger_DriverLoadFailurePolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 193)  //col:1753
    /// <summary> BcdOSLoaderElementTypes = 277  //col:1754
    /// Defines the type of boot menus the system will use. Possible values include menupolicylegacy (0) or menupolicystandard (1). BcdOSLoaderElementTypes = 278  //col:1755
    /// The default value is menupolicylegacy (0). BcdOSLoaderElementTypes = 279  //col:1756
    /// </summary> BcdOSLoaderElementTypes = 280  //col:1757
    /// <remarks>0x250000C2</remarks> BcdOSLoaderElementTypes = 281  //col:1758
    BcdOSLoaderInteger_BootMenuPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 194)  //col:1759
    /// <summary> BcdOSLoaderElementTypes = 283  //col:1760
    /// Controls whether the system boots to the legacy menu (F8 menu) on the next boot. BcdOSLoaderElementTypes = 284  //col:1761
    /// Note This value is supported starting in Windows 8 and Windows Server 2012. BcdOSLoaderElementTypes = 285  //col:1762
    /// </summary> BcdOSLoaderElementTypes = 286  //col:1763
    /// <remarks>0x260000C3</remarks> BcdOSLoaderElementTypes = 287  //col:1764
    BcdOSLoaderBoolean_AdvancedOptionsOneTime  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 195)  //col:1765
    /// <summary> BcdOSLoaderElementTypes = 289  //col:1766
    /// BcdOSLoaderElementTypes = 290  //col:1767
    /// </summary> BcdOSLoaderElementTypes = 291  //col:1768
    /// <remarks>0x260000C4</remarks> BcdOSLoaderElementTypes = 292  //col:1769
    BcdOSLoaderBoolean_OptionsEditOneTime  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 196)  //col:1770
    /// <summary> BcdOSLoaderElementTypes = 294  //col:1771
    /// The boot status policy. The Integer property is one of the values from the BcdOSLoaderBootStatusPolicy enumeration BcdOSLoaderElementTypes = 295  //col:1772
    /// </summary> BcdOSLoaderElementTypes = 296  //col:1773
    /// <remarks>0x250000E0</remarks> BcdOSLoaderElementTypes = 297  //col:1774
    BcdOSLoaderInteger_BootStatusPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 224)  //col:1775
    /// <summary> BcdOSLoaderElementTypes = 299  //col:1776
    /// The OS loader removes this entry for security reasons. This option can only be triggered by using the F8 menu; a user must be physically present to trigger this option. BcdOSLoaderElementTypes = 300  //col:1777
    /// This value is supported starting in Windows 8 and Windows Server 2012. BcdOSLoaderElementTypes = 301  //col:1778
    /// </summary> BcdOSLoaderElementTypes = 302  //col:1779
    /// <remarks>0x260000E1</remarks> BcdOSLoaderElementTypes = 303  //col:1780
    BcdOSLoaderBoolean_DisableElamDrivers  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 225)  //col:1781
    /// <summary> BcdOSLoaderElementTypes = 305  //col:1782
    /// Controls the hypervisor launch type. Options are HyperVisorLaunchOff (0) and HypervisorLaunchAuto (1). BcdOSLoaderElementTypes = 306  //col:1783
    /// </summary> BcdOSLoaderElementTypes = 307  //col:1784
    /// <remarks>0x250000F0</remarks> BcdOSLoaderElementTypes = 308  //col:1785
    BcdOSLoaderInteger_HypervisorLaunchType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 240)  //col:1786
    /// <summary> BcdOSLoaderElementTypes = 310  //col:1787
    /// BcdOSLoaderElementTypes = 311  //col:1788
    /// </summary> BcdOSLoaderElementTypes = 312  //col:1789
    /// <remarks>0x250000F1</remarks> BcdOSLoaderElementTypes = 313  //col:1790
    BcdOSLoaderString_HypervisorPath  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 241)  //col:1791
    /// <summary> BcdOSLoaderElementTypes = 315  //col:1792
    /// Controls whether the hypervisor debugger is enabled. BcdOSLoaderElementTypes = 316  //col:1793
    /// </summary> BcdOSLoaderElementTypes = 317  //col:1794
    /// <remarks>0x260000F2</remarks> BcdOSLoaderElementTypes = 318  //col:1795
    BcdOSLoaderBoolean_HypervisorDebuggerEnabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 242)  //col:1796
    /// <summary> BcdOSLoaderElementTypes = 320  //col:1797
    /// Controls the hypervisor debugger type. Can be set to SERIAL (0) 1394 (1) or NET (2). BcdOSLoaderElementTypes = 321  //col:1798
    /// </summary> BcdOSLoaderElementTypes = 322  //col:1799
    /// <remarks>0x250000F3</remarks> BcdOSLoaderElementTypes = 323  //col:1800
    BcdOSLoaderInteger_HypervisorDebuggerType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 243)  //col:1801
    /// <summary> BcdOSLoaderElementTypes = 325  //col:1802
    /// Specifies the serial port number for serial debugging. BcdOSLoaderElementTypes = 326  //col:1803
    /// </summary> BcdOSLoaderElementTypes = 327  //col:1804
    /// <remarks>0x250000F4</remarks> BcdOSLoaderElementTypes = 328  //col:1805
    BcdOSLoaderInteger_HypervisorDebuggerPortNumber  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 244)  //col:1806
    /// <summary> BcdOSLoaderElementTypes = 330  //col:1807
    /// Specifies the baud rate for serial debugging. BcdOSLoaderElementTypes = 331  //col:1808
    /// </summary> BcdOSLoaderElementTypes = 332  //col:1809
    /// <remarks>0x250000F5</remarks> BcdOSLoaderElementTypes = 333  //col:1810
    BcdOSLoaderInteger_HypervisorDebuggerBaudrate  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 245)  //col:1811
    /// <summary> BcdOSLoaderElementTypes = 335  //col:1812
    /// Specifies the channel number for 1394 debugging. BcdOSLoaderElementTypes = 336  //col:1813
    /// </summary> BcdOSLoaderElementTypes = 337  //col:1814
    /// <remarks>0x250000F6</remarks> BcdOSLoaderElementTypes = 338  //col:1815
    BcdOSLoaderInteger_HypervisorDebugger1394Channel  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 246)  //col:1816
    /// <summary> BcdOSLoaderElementTypes = 340  //col:1817
    /// Values are Disabled (0) Basic (1) and Standard (2). BcdOSLoaderElementTypes = 341  //col:1818
    /// </summary> BcdOSLoaderElementTypes = 342  //col:1819
    /// <remarks>0x250000F7</remarks> BcdOSLoaderElementTypes = 343  //col:1820
    BcdOSLoaderInteger_BootUxPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 247)  //col:1821
    /// <summary> BcdOSLoaderElementTypes = 345  //col:1822
    /// BcdOSLoaderElementTypes = 346  //col:1823
    /// </summary> BcdOSLoaderElementTypes = 347  //col:1824
    /// <remarks>0x220000F8</remarks> BcdOSLoaderElementTypes = 348  //col:1825
    BcdOSLoaderInteger_HypervisorSlatDisabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 248)  //col:1826
    /// <summary> BcdOSLoaderElementTypes = 350  //col:1827
    /// Defines the PCI bus device and function numbers of the debugging device used with the hypervisor. BcdOSLoaderElementTypes = 351  //col:1828
    /// For example 1.5.0 describes the debugging device on bus 1 device 5 function 0. BcdOSLoaderElementTypes = 352  //col:1829
    /// </summary> BcdOSLoaderElementTypes = 353  //col:1830
    /// <remarks>0x220000F9</remarks> BcdOSLoaderElementTypes = 354  //col:1831
    BcdOSLoaderString_HypervisorDebuggerBusParams  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 249)  //col:1832
    /// <summary> BcdOSLoaderElementTypes = 356  //col:1833
    /// BcdOSLoaderElementTypes = 357  //col:1834
    /// </summary> BcdOSLoaderElementTypes = 358  //col:1835
    /// <remarks>0x250000FA</remarks> BcdOSLoaderElementTypes = 359  //col:1836
    BcdOSLoaderInteger_HypervisorNumProc  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 250)  //col:1837
    /// <summary> BcdOSLoaderElementTypes = 361  //col:1838
    /// BcdOSLoaderElementTypes = 362  //col:1839
    /// </summary> BcdOSLoaderElementTypes = 363  //col:1840
    /// <remarks>0x250000FB</remarks> BcdOSLoaderElementTypes = 364  //col:1841
    BcdOSLoaderInteger_HypervisorRootProcPerNode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 251)  //col:1842
    /// <summary> BcdOSLoaderElementTypes = 366  //col:1843
    /// BcdOSLoaderElementTypes = 367  //col:1844
    /// </summary> BcdOSLoaderElementTypes = 368  //col:1845
    /// <remarks>0x260000FC</remarks> BcdOSLoaderElementTypes = 369  //col:1846
    BcdOSLoaderBoolean_HypervisorUseLargeVTlb  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 252)  //col:1847
    /// <summary> BcdOSLoaderElementTypes = 371  //col:1848
    /// BcdOSLoaderElementTypes = 372  //col:1849
    /// </summary> BcdOSLoaderElementTypes = 373  //col:1850
    /// <remarks>0x250000FD</remarks> BcdOSLoaderElementTypes = 374  //col:1851
    BcdOSLoaderInteger_HypervisorDebuggerNetHostIp  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 253)  //col:1852
    /// <summary> BcdOSLoaderElementTypes = 376  //col:1853
    /// BcdOSLoaderElementTypes = 377  //col:1854
    /// </summary> BcdOSLoaderElementTypes = 378  //col:1855
    /// <remarks>0x250000FE</remarks> BcdOSLoaderElementTypes = 379  //col:1856
    BcdOSLoaderInteger_HypervisorDebuggerNetHostPort  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 254)  //col:1857
    /// <summary> BcdOSLoaderElementTypes = 381  //col:1858
    /// BcdOSLoaderElementTypes = 382  //col:1859
    /// </summary> BcdOSLoaderElementTypes = 383  //col:1860
    /// <remarks>0x250000FF</remarks> BcdOSLoaderElementTypes = 384  //col:1861
    BcdOSLoaderInteger_HypervisorDebuggerPages  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 255)  //col:1862
    /// <summary> BcdOSLoaderElementTypes = 386  //col:1863
    /// BcdOSLoaderElementTypes = 387  //col:1864
    /// </summary> BcdOSLoaderElementTypes = 388  //col:1865
    /// <remarks>0x25000100</remarks> BcdOSLoaderElementTypes = 389  //col:1866
    BcdOSLoaderInteger_TpmBootEntropyPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 256)  //col:1867
    /// <summary> BcdOSLoaderElementTypes = 391  //col:1868
    /// BcdOSLoaderElementTypes = 392  //col:1869
    /// </summary> BcdOSLoaderElementTypes = 393  //col:1870
    /// <remarks>0x22000110</remarks> BcdOSLoaderElementTypes = 394  //col:1871
    BcdOSLoaderString_HypervisorDebuggerNetKey  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 272)  //col:1872
    /// <summary> BcdOSLoaderElementTypes = 396  //col:1873
    /// BcdOSLoaderElementTypes = 397  //col:1874
    /// </summary> BcdOSLoaderElementTypes = 398  //col:1875
    /// <remarks>0x22000112</remarks> BcdOSLoaderElementTypes = 399  //col:1876
    BcdOSLoaderString_HypervisorProductSkuType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 274)  //col:1877
    /// <summary> BcdOSLoaderElementTypes = 401  //col:1878
    /// BcdOSLoaderElementTypes = 402  //col:1879
    /// </summary> BcdOSLoaderElementTypes = 403  //col:1880
    /// <remarks>0x22000113</remarks> BcdOSLoaderElementTypes = 404  //col:1881
    BcdOSLoaderInteger_HypervisorRootProc  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 275)  //col:1882
    /// <summary> BcdOSLoaderElementTypes = 406  //col:1883
    /// BcdOSLoaderElementTypes = 407  //col:1884
    /// </summary> BcdOSLoaderElementTypes = 408  //col:1885
    /// <remarks>0x26000114</remarks> BcdOSLoaderElementTypes = 409  //col:1886
    BcdOSLoaderBoolean_HypervisorDebuggerNetDhcp  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 276)  //col:1887
    /// <summary> BcdOSLoaderElementTypes = 411  //col:1888
    /// BcdOSLoaderElementTypes = 412  //col:1889
    /// </summary> BcdOSLoaderElementTypes = 413  //col:1890
    /// <remarks>0x25000115</remarks> BcdOSLoaderElementTypes = 414  //col:1891
    BcdOSLoaderInteger_HypervisorIommuPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 277)  //col:1892
    /// <summary> BcdOSLoaderElementTypes = 416  //col:1893
    /// BcdOSLoaderElementTypes = 417  //col:1894
    /// </summary> BcdOSLoaderElementTypes = 418  //col:1895
    /// <remarks>0x26000116</remarks> BcdOSLoaderElementTypes = 419  //col:1896
    BcdOSLoaderBoolean_HypervisorUseVApic  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 278)  //col:1897
    /// <summary> BcdOSLoaderElementTypes = 421  //col:1898
    /// BcdOSLoaderElementTypes = 422  //col:1899
    /// </summary> BcdOSLoaderElementTypes = 423  //col:1900
    /// <remarks>0x22000117</remarks> BcdOSLoaderElementTypes = 424  //col:1901
    BcdOSLoaderString_HypervisorLoadOptions  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 279)  //col:1902
    /// <summary> BcdOSLoaderElementTypes = 426  //col:1903
    /// BcdOSLoaderElementTypes = 427  //col:1904
    /// </summary> BcdOSLoaderElementTypes = 428  //col:1905
    /// <remarks>0x25000118</remarks> BcdOSLoaderElementTypes = 429  //col:1906
    BcdOSLoaderInteger_HypervisorMsrFilterPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 280)  //col:1907
    /// <summary> BcdOSLoaderElementTypes = 431  //col:1908
    /// BcdOSLoaderElementTypes = 432  //col:1909
    /// </summary> BcdOSLoaderElementTypes = 433  //col:1910
    /// <remarks>0x25000119</remarks> BcdOSLoaderElementTypes = 434  //col:1911
    BcdOSLoaderInteger_HypervisorMmioNxPolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 281)  //col:1912
    /// <summary> BcdOSLoaderElementTypes = 436  //col:1913
    /// BcdOSLoaderElementTypes = 437  //col:1914
    /// </summary> BcdOSLoaderElementTypes = 438  //col:1915
    /// <remarks>0x2500011A</remarks> BcdOSLoaderElementTypes = 439  //col:1916
    BcdOSLoaderInteger_HypervisorSchedulerType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 282)  //col:1917
    /// <summary> BcdOSLoaderElementTypes = 441  //col:1918
    /// BcdOSLoaderElementTypes = 442  //col:1919
    /// </summary> BcdOSLoaderElementTypes = 443  //col:1920
    /// <remarks>0x2200011B</remarks> BcdOSLoaderElementTypes = 444  //col:1921
    BcdOSLoaderString_HypervisorRootProcNumaNodes  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 283)  //col:1922
    /// <summary> BcdOSLoaderElementTypes = 446  //col:1923
    /// BcdOSLoaderElementTypes = 447  //col:1924
    /// </summary> BcdOSLoaderElementTypes = 448  //col:1925
    /// <remarks>0x2500011C</remarks> BcdOSLoaderElementTypes = 449  //col:1926
    BcdOSLoaderInteger_HypervisorPerfmon  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 284)  //col:1927
    /// <summary> BcdOSLoaderElementTypes = 451  //col:1928
    /// BcdOSLoaderElementTypes = 452  //col:1929
    /// </summary> BcdOSLoaderElementTypes = 453  //col:1930
    /// <remarks>0x2500011D</remarks> BcdOSLoaderElementTypes = 454  //col:1931
    BcdOSLoaderInteger_HypervisorRootProcPerCore  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 285)  //col:1932
    /// <summary> BcdOSLoaderElementTypes = 456  //col:1933
    /// BcdOSLoaderElementTypes = 457  //col:1934
    /// </summary> BcdOSLoaderElementTypes = 458  //col:1935
    /// <remarks>0x2200011E</remarks> BcdOSLoaderElementTypes = 459  //col:1936
    BcdOSLoaderString_HypervisorRootProcNumaNodeLps  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 286)  //col:1937
    /// <summary> BcdOSLoaderElementTypes = 461  //col:1938
    /// BcdOSLoaderElementTypes = 462  //col:1939
    /// </summary> BcdOSLoaderElementTypes = 463  //col:1940
    /// <remarks>0x25000120</remarks> BcdOSLoaderElementTypes = 464  //col:1941
    BcdOSLoaderInteger_XSavePolicy  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 288)  //col:1942
    /// <summary> BcdOSLoaderElementTypes = 466  //col:1943
    /// BcdOSLoaderElementTypes = 467  //col:1944
    /// </summary> BcdOSLoaderElementTypes = 468  //col:1945
    /// <remarks>0x25000121</remarks> BcdOSLoaderElementTypes = 469  //col:1946
    BcdOSLoaderInteger_XSaveAddFeature0  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 289)  //col:1947
    /// <summary> BcdOSLoaderElementTypes = 471  //col:1948
    /// BcdOSLoaderElementTypes = 472  //col:1949
    /// </summary> BcdOSLoaderElementTypes = 473  //col:1950
    /// <remarks>0x25000122</remarks> BcdOSLoaderElementTypes = 474  //col:1951
    BcdOSLoaderInteger_XSaveAddFeature1  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 290)  //col:1952
    /// <summary> BcdOSLoaderElementTypes = 476  //col:1953
    /// BcdOSLoaderElementTypes = 477  //col:1954
    /// </summary> BcdOSLoaderElementTypes = 478  //col:1955
    /// <remarks>0x25000123</remarks> BcdOSLoaderElementTypes = 479  //col:1956
    BcdOSLoaderInteger_XSaveAddFeature2  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 291)  //col:1957
    /// <summary> BcdOSLoaderElementTypes = 481  //col:1958
    /// BcdOSLoaderElementTypes = 482  //col:1959
    /// </summary> BcdOSLoaderElementTypes = 483  //col:1960
    /// <remarks>0x25000124</remarks> BcdOSLoaderElementTypes = 484  //col:1961
    BcdOSLoaderInteger_XSaveAddFeature3  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 292)  //col:1962
    /// <summary> BcdOSLoaderElementTypes = 486  //col:1963
    /// BcdOSLoaderElementTypes = 487  //col:1964
    /// </summary> BcdOSLoaderElementTypes = 488  //col:1965
    /// <remarks>0x25000125</remarks> BcdOSLoaderElementTypes = 489  //col:1966
    BcdOSLoaderInteger_XSaveAddFeature4  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 293)  //col:1967
    /// <summary> BcdOSLoaderElementTypes = 491  //col:1968
    /// BcdOSLoaderElementTypes = 492  //col:1969
    /// </summary> BcdOSLoaderElementTypes = 493  //col:1970
    /// <remarks>0x25000126</remarks> BcdOSLoaderElementTypes = 494  //col:1971
    BcdOSLoaderInteger_XSaveAddFeature5  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 294)  //col:1972
    /// <summary> BcdOSLoaderElementTypes = 496  //col:1973
    /// BcdOSLoaderElementTypes = 497  //col:1974
    /// </summary> BcdOSLoaderElementTypes = 498  //col:1975
    /// <remarks>0x25000127</remarks> BcdOSLoaderElementTypes = 499  //col:1976
    BcdOSLoaderInteger_XSaveAddFeature6  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 295)  //col:1977
    /// <summary> BcdOSLoaderElementTypes = 501  //col:1978
    /// BcdOSLoaderElementTypes = 502  //col:1979
    /// </summary> BcdOSLoaderElementTypes = 503  //col:1980
    /// <remarks>0x25000128</remarks> BcdOSLoaderElementTypes = 504  //col:1981
    BcdOSLoaderInteger_XSaveAddFeature7  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 296)  //col:1982
    /// <summary> BcdOSLoaderElementTypes = 506  //col:1983
    /// BcdOSLoaderElementTypes = 507  //col:1984
    /// </summary> BcdOSLoaderElementTypes = 508  //col:1985
    /// <remarks>0x25000129</remarks> BcdOSLoaderElementTypes = 509  //col:1986
    BcdOSLoaderInteger_XSaveRemoveFeature  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 297)  //col:1987
    /// <summary> BcdOSLoaderElementTypes = 511  //col:1988
    /// BcdOSLoaderElementTypes = 512  //col:1989
    /// </summary> BcdOSLoaderElementTypes = 513  //col:1990
    /// <remarks>0x2500012A</remarks> BcdOSLoaderElementTypes = 514  //col:1991
    BcdOSLoaderInteger_XSaveProcessorsMask  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 298)  //col:1992
    /// <summary> BcdOSLoaderElementTypes = 516  //col:1993
    /// BcdOSLoaderElementTypes = 517  //col:1994
    /// </summary> BcdOSLoaderElementTypes = 518  //col:1995
    /// <remarks>0x2500012B</remarks> BcdOSLoaderElementTypes = 519  //col:1996
    BcdOSLoaderInteger_XSaveDisable  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 299)  //col:1997
    /// <summary> BcdOSLoaderElementTypes = 521  //col:1998
    /// BcdOSLoaderElementTypes = 522  //col:1999
    /// </summary> BcdOSLoaderElementTypes = 523  //col:2000
    /// <remarks>0x2500012C</remarks> BcdOSLoaderElementTypes = 524  //col:2001
    BcdOSLoaderInteger_KernelDebuggerType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 300)  //col:2002
    /// <summary> BcdOSLoaderElementTypes = 526  //col:2003
    /// BcdOSLoaderElementTypes = 527  //col:2004
    /// </summary> BcdOSLoaderElementTypes = 528  //col:2005
    /// <remarks>0x2200012D</remarks> BcdOSLoaderElementTypes = 529  //col:2006
    BcdOSLoaderString_KernelDebuggerBusParameters  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 301)  //col:2007
    /// <summary> BcdOSLoaderElementTypes = 531  //col:2008
    /// BcdOSLoaderElementTypes = 532  //col:2009
    /// </summary> BcdOSLoaderElementTypes = 533  //col:2010
    /// <remarks>0x2500012E</remarks> BcdOSLoaderElementTypes = 534  //col:2011
    BcdOSLoaderInteger_KernelDebuggerPortAddress  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 302)  //col:2012
    /// <summary> BcdOSLoaderElementTypes = 536  //col:2013
    /// BcdOSLoaderElementTypes = 537  //col:2014
    /// </summary> BcdOSLoaderElementTypes = 538  //col:2015
    /// <remarks>0x2500012F</remarks> BcdOSLoaderElementTypes = 539  //col:2016
    BcdOSLoaderInteger_KernelDebuggerPortNumber  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 303)  //col:2017
    /// <summary> BcdOSLoaderElementTypes = 541  //col:2018
    /// BcdOSLoaderElementTypes = 542  //col:2019
    /// </summary> BcdOSLoaderElementTypes = 543  //col:2020
    /// <remarks>0x25000130</remarks> BcdOSLoaderElementTypes = 544  //col:2021
    BcdOSLoaderInteger_ClaimedTpmCounter  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 304)  //col:2022
    /// <summary> BcdOSLoaderElementTypes = 546  //col:2023
    /// BcdOSLoaderElementTypes = 547  //col:2024
    /// </summary> BcdOSLoaderElementTypes = 548  //col:2025
    /// <remarks>0x25000131</remarks> BcdOSLoaderElementTypes = 549  //col:2026
    BcdOSLoaderInteger_KernelDebugger1394Channel  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 305)  //col:2027
    /// <summary> BcdOSLoaderElementTypes = 551  //col:2028
    /// BcdOSLoaderElementTypes = 552  //col:2029
    /// </summary> BcdOSLoaderElementTypes = 553  //col:2030
    /// <remarks>0x22000132</remarks> BcdOSLoaderElementTypes = 554  //col:2031
    BcdOSLoaderString_KernelDebuggerUsbTargetname  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 306)  //col:2032
    /// <summary> BcdOSLoaderElementTypes = 556  //col:2033
    /// BcdOSLoaderElementTypes = 557  //col:2034
    /// </summary> BcdOSLoaderElementTypes = 558  //col:2035
    /// <remarks>0x25000133</remarks> BcdOSLoaderElementTypes = 559  //col:2036
    BcdOSLoaderInteger_KernelDebuggerNetHostIp  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 307)  //col:2037
    /// <summary> BcdOSLoaderElementTypes = 561  //col:2038
    /// BcdOSLoaderElementTypes = 562  //col:2039
    /// </summary> BcdOSLoaderElementTypes = 563  //col:2040
    /// <remarks>0x25000134</remarks> BcdOSLoaderElementTypes = 564  //col:2041
    BcdOSLoaderInteger_KernelDebuggerNetHostPort  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 308)  //col:2042
    /// <summary> BcdOSLoaderElementTypes = 566  //col:2043
    /// BcdOSLoaderElementTypes = 567  //col:2044
    /// </summary> BcdOSLoaderElementTypes = 568  //col:2045
    /// <remarks>0x26000135</remarks> BcdOSLoaderElementTypes = 569  //col:2046
    BcdOSLoaderBoolean_KernelDebuggerNetDhcp  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 309)  //col:2047
    /// <summary> BcdOSLoaderElementTypes = 571  //col:2048
    /// BcdOSLoaderElementTypes = 572  //col:2049
    /// </summary> BcdOSLoaderElementTypes = 573  //col:2050
    /// <remarks>0x22000136</remarks> BcdOSLoaderElementTypes = 574  //col:2051
    BcdOSLoaderString_KernelDebuggerNetKey  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 310)  //col:2052
    /// <summary> BcdOSLoaderElementTypes = 576  //col:2053
    /// BcdOSLoaderElementTypes = 577  //col:2054
    /// </summary> BcdOSLoaderElementTypes = 578  //col:2055
    /// <remarks>0x22000137</remarks> BcdOSLoaderElementTypes = 579  //col:2056
    BcdOSLoaderString_IMCHiveName  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 311)  //col:2057
    /// <summary> BcdOSLoaderElementTypes = 581  //col:2058
    /// BcdOSLoaderElementTypes = 582  //col:2059
    /// </summary> BcdOSLoaderElementTypes = 583  //col:2060
    /// <remarks>0x21000138</remarks> BcdOSLoaderElementTypes = 584  //col:2061
    BcdOSLoaderDevice_IMCDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 312)  //col:2062
    /// <summary> BcdOSLoaderElementTypes = 586  //col:2063
    /// BcdOSLoaderElementTypes = 587  //col:2064
    /// </summary> BcdOSLoaderElementTypes = 588  //col:2065
    /// <remarks>0x25000139</remarks> BcdOSLoaderElementTypes = 589  //col:2066
    BcdOSLoaderInteger_KernelDebuggerBaudrate  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 313)  //col:2067
    /// <summary> BcdOSLoaderElementTypes = 591  //col:2068
    /// BcdOSLoaderElementTypes = 592  //col:2069
    /// </summary> BcdOSLoaderElementTypes = 593  //col:2070
    /// <remarks>0x22000140</remarks> BcdOSLoaderElementTypes = 594  //col:2071
    BcdOSLoaderString_ManufacturingMode  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 320)  //col:2072
    /// <summary> BcdOSLoaderElementTypes = 596  //col:2073
    /// BcdOSLoaderElementTypes = 597  //col:2074
    /// </summary> BcdOSLoaderElementTypes = 598  //col:2075
    /// <remarks>0x26000141</remarks> BcdOSLoaderElementTypes = 599  //col:2076
    BcdOSLoaderBoolean_EventLoggingEnabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 321)  //col:2077
    /// <summary> BcdOSLoaderElementTypes = 601  //col:2078
    /// BcdOSLoaderElementTypes = 602  //col:2079
    /// </summary> BcdOSLoaderElementTypes = 603  //col:2080
    /// <remarks>0x25000142</remarks> BcdOSLoaderElementTypes = 604  //col:2081
    BcdOSLoaderInteger_VsmLaunchType  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 322)  //col:2082
    /// <summary> BcdOSLoaderElementTypes = 606  //col:2083
    /// Undocumented. Zero (0) indicates default one (1) indicates that disabled and two (2) indicates strict mode. BcdOSLoaderElementTypes = 607  //col:2084
    /// </summary> BcdOSLoaderElementTypes = 608  //col:2085
    /// <remarks>0x25000144</remarks> BcdOSLoaderElementTypes = 609  //col:2086
    BcdOSLoaderInteger_HypervisorEnforcedCodeIntegrity  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_INTEGER 324)  //col:2087
    /// <summary> BcdOSLoaderElementTypes = 611  //col:2088
    /// BcdOSLoaderElementTypes = 612  //col:2089
    /// </summary> BcdOSLoaderElementTypes = 613  //col:2090
    /// <remarks>0x26000145</remarks> BcdOSLoaderElementTypes = 614  //col:2091
    BcdOSLoaderBoolean_DtraceEnabled  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN 325)  //col:2092
    /// <summary> BcdOSLoaderElementTypes = 616  //col:2093
    /// BcdOSLoaderElementTypes = 617  //col:2094
    /// </summary> BcdOSLoaderElementTypes = 618  //col:2095
    /// <remarks>0x21000150</remarks> BcdOSLoaderElementTypes = 619  //col:2096
    BcdOSLoaderDevice_SystemDataDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 336)  //col:2097
    /// <summary> BcdOSLoaderElementTypes = 621  //col:2098
    /// BcdOSLoaderElementTypes = 622  //col:2099
    /// </summary> BcdOSLoaderElementTypes = 623  //col:2100
    /// <remarks>0x21000151</remarks> BcdOSLoaderElementTypes = 624  //col:2101
    BcdOSLoaderDevice_OsArcDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 337)  //col:2102
    /// <summary> BcdOSLoaderElementTypes = 626  //col:2103
    /// BcdOSLoaderElementTypes = 627  //col:2104
    /// </summary> BcdOSLoaderElementTypes = 628  //col:2105
    /// <remarks>0x21000153</remarks> BcdOSLoaderElementTypes = 629  //col:2106
    BcdOSLoaderDevice_OsDataDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 339)  //col:2107
    /// <summary> BcdOSLoaderElementTypes = 631  //col:2108
    /// BcdOSLoaderElementTypes = 632  //col:2109
    /// </summary> BcdOSLoaderElementTypes = 633  //col:2110
    /// <remarks>0x21000154</remarks> BcdOSLoaderElementTypes = 634  //col:2111
    BcdOSLoaderDevice_BspDevice  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 340)  //col:2112
    /// <summary> BcdOSLoaderElementTypes = 636  //col:2113
    /// BcdOSLoaderElementTypes = 637  //col:2114
    /// </summary> BcdOSLoaderElementTypes = 638  //col:2115
    /// <remarks>0x21000155</remarks> BcdOSLoaderElementTypes = 639  //col:2116
    BcdOSLoaderDevice_BspFilepath  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_DEVICE 341)  //col:2117
    /// <summary> BcdOSLoaderElementTypes = 641  //col:2118
    /// BcdOSLoaderElementTypes = 642  //col:2119
    /// </summary> BcdOSLoaderElementTypes = 643  //col:2120
    /// <remarks>0x22000156</remarks> BcdOSLoaderElementTypes = 644  //col:2121
    BcdOSLoaderString_KernelDebuggerNetHostIpv6  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 342)  //col:2122
    /// <summary> BcdOSLoaderElementTypes = 646  //col:2123
    /// BcdOSLoaderElementTypes = 647  //col:2124
    /// </summary> BcdOSLoaderElementTypes = 648  //col:2125
    /// <remarks>0x22000161</remarks> BcdOSLoaderElementTypes = 649  //col:2126
    BcdOSLoaderString_HypervisorDebuggerNetHostIpv6  BcdOSLoaderElementTypes =  MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION BCD_ELEMENT_DATATYPE_FORMAT_STRING 353)  //col:2127
)




type BCD_OBJECT struct{
Identifer GUID
Description PBCD_OBJECT_DESCRIPTION
}


type BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION struct{
PartitionStyle ULONG
Reserved ULONG
Struct struct
Union union
DiskSignature ULONG
PartitionOffset ULONG64
}


type BCD_ELEMENT_DEVICE struct{
DeviceType ULONG
AdditionalOptions GUID
Struct struct
Union union
ParentOffset ULONG
Path[ANYSIZE_ARRAY] WCHAR
}


type BCD_ELEMENT_STRING struct{
Value[ANYSIZE_ARRAY] WCHAR
}


type BCD_ELEMENT_OBJECT struct{
Object GUID
}


type BCD_ELEMENT_OBJECT_LIST struct{
ObjectList[ANYSIZE_ARRAY] GUID
}


type BCD_ELEMENT_INTEGER struct{
Value ULONG64
}


type BCD_ELEMENT_INTEGER_LIST struct{
Value[ANYSIZE_ARRAY] ULONG64
}


type BCD_ELEMENT_BOOLEAN struct{
Value bool
}


type typedef struct BCD_ELEMENT_DESCRIPTION struct{
Type ULONG
DataSize ULONG
}


type BCD_ELEMENT struct{
Description PBCD_ELEMENT_DESCRIPTION
Data PVOID
}



type (
Ntbcd interface{
 * Attribution 4.0 International ()(ok bool)//col:90
typedef VOID ()(ok bool)//col:176
BcdImportStoreWithFlags()(ok bool)//col:206
BcdOpenStore()(ok bool)//col:251
#define MAKE_BCD_OBJECT()(ok bool)//col:334
BcdEnumerateObjects()(ok bool)//col:403
BcdCopyObject()(ok bool)//col:468
#define MAKE_BCDE_DATA_TYPE()(ok bool)//col:514
BcdEnumerateElementTypes()(ok bool)//col:543
BcdEnumerateElements()(ok bool)//col:646
BcdEnumerateElementsWithFlags()(ok bool)//col:833
    BcdLibraryDevice_ApplicationDevice = MAKE_BCDE_DATA_TYPE()(ok bool)//col:1361
    BcdSetupInteger_DeviceType = MAKE_BCDE_DATA_TYPE()(ok bool)//col:1395
    BcdOSLoaderDevice_OSDevice = MAKE_BCDE_DATA_TYPE()(ok bool)//col:2128
}
)

func NewNtbcd() { return & ntbcd{} }

func (n *ntbcd) * Attribution 4.0 International ()(ok bool){//col:90
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTBCD_H
#define _NTBCD_H
#ifndef PHNT_INLINE_BCD_GUIDS
GUID DECLSPEC_SELECTANY GUID_BAD_MEMORY_GROUP = { 0x5189B25C, 0x5558, 0x4BF2, { 0xBC, 0xA4, 0x28, 0x9B, 0x11, 0xBD, 0x29, 0xE2 } };
GUID DECLSPEC_SELECTANY GUID_BOOT_LOADER_SETTINGS_GROUP = { 0x6EFB52BF, 0x1766, 0x41DB, { 0xA6, 0xB3, 0x0E, 0xE5, 0xEF, 0xF7, 0x2B, 0xD7 } };
GUID DECLSPEC_SELECTANY GUID_CURRENT_BOOT_ENTRY = { 0xFA926493, 0x6F1C, 0x4193, { 0xA4, 0x14, 0x58, 0xF0, 0xB2, 0x45, 0x6D, 0x1E } };
GUID DECLSPEC_SELECTANY GUID_DEBUGGER_SETTINGS_GROUP = { 0x4636856E, 0x540F, 0x4170, { 0xA1, 0x30, 0xA8, 0x47, 0x76, 0xF4, 0xC6, 0x54 } };
GUID DECLSPEC_SELECTANY GUID_DEFAULT_BOOT_ENTRY = { 0x1CAE1EB7, 0xA0DF, 0x4D4D, { 0x98, 0x51, 0x48, 0x60, 0xE3, 0x4E, 0xF5, 0x35 } };
GUID DECLSPEC_SELECTANY GUID_EMS_SETTINGS_GROUP = { 0x0CE4991B, 0xE6B3, 0x4B16, { 0xB2, 0x3C, 0x5E, 0x0D, 0x92, 0x50, 0xE5, 0xD9 } };
GUID DECLSPEC_SELECTANY GUID_FIRMWARE_BOOTMGR = { 0xA5A30FA2, 0x3D06, 0x4E9F, { 0xB5, 0xF4, 0xA0, 0x1D, 0xF9, 0xD1, 0xFC, 0xBA } };
GUID DECLSPEC_SELECTANY GUID_GLOBAL_SETTINGS_GROUP = { 0x7EA2E1AC, 0x2E61, 0x4728, { 0xAA, 0xA3, 0x89, 0x6D, 0x9D, 0x0A, 0x9F, 0x0E } };
GUID DECLSPEC_SELECTANY GUID_HYPERVISOR_SETTINGS_GROUP = { 0x7FF607E0, 0x4395, 0x11DB, { 0xB0, 0xDE, 0x08, 0x00, 0x20, 0x0C, 0x9A, 0x66 } };
GUID DECLSPEC_SELECTANY GUID_KERNEL_DEBUGGER_SETTINGS_GROUP = { 0x313E8EED, 0x7098, 0x4586, { 0xA9, 0xBF, 0x30, 0x9C, 0x61, 0xF8, 0xD4, 0x49 } };
GUID DECLSPEC_SELECTANY GUID_RESUME_LOADER_SETTINGS_GROUP = { 0x1AFA9C49, 0x16AB, 0x4A5C, { 0x4A, 0x90, 0x21, 0x28, 0x02, 0xDA, 0x94, 0x60 } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_BOOTMGR = { 0x9DEA862C, 0x5CDD, 0x4E70, { 0xAC, 0xC1, 0xF3, 0x2B, 0x34, 0x4D, 0x47, 0x95 } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_LEGACY_NTLDR = { 0x466F5A88, 0x0AF2, 0x4F76, { 0x90, 0x38, 0x09, 0x5B, 0x17, 0x0D, 0xC2, 0x1C } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_MEMORY_TESTER = { 0xB2721D73, 0x1DB4, 0x4C62, { 0xBF, 0x78, 0xC5, 0x48, 0xA8, 0x80, 0x14, 0x2D } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_OS_TARGET_TEMPLATE_EFI = { 0xB012B84D, 0xC47C, 0x4ED5, { 0xB7, 0x22, 0xC0, 0xC4, 0x21, 0x63, 0xE5, 0x69 } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_OS_TARGET_TEMPLATE_PCAT = { 0xA1943BBC, 0xEA85, 0x487C, { 0x97, 0xC7, 0xC9, 0xED, 0xE9, 0x08, 0xA3, 0x8A } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_RESUME_TARGET_TEMPLATE_EFI = { 0x0C334284, 0x9A41, 0x4DE1, { 0x99, 0xB3, 0xA7, 0xE8, 0x7E, 0x8F, 0xF0, 0x7E } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_RESUME_TARGET_TEMPLATE_PCAT = { 0x98B02A23, 0x0674, 0x4CE7, { 0xBD, 0xAD, 0xE0, 0xA1, 0x5A, 0x8F, 0xF9, 0x7B } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_EFI = { 0x7254A080, 0x1510, 0x4E85, { 0xAC, 0x0F, 0xE7, 0xFB, 0x3D, 0x44, 0x47, 0x36 } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_PCAT = { 0xCBD971BF, 0xB7B8, 0x4885, { 0x95, 0x1A, 0xFA, 0x03, 0x04, 0x4F, 0x5D, 0x71 } };
GUID DECLSPEC_SELECTANY GUID_WINDOWS_SETUP_RAMDISK_OPTIONS = { 0xAE5534E0, 0xA924, 0x466C, { 0xB8, 0x36, 0x75, 0x85, 0x39, 0xA3, 0xEE, 0x3A } };
#else
NTSYSAPI GUID GUID_WINDOWS_OS_TARGET_TEMPLATE_EFI;
NTSYSAPI GUID GUID_WINDOWS_OS_TARGET_TEMPLATE_PCAT;
NTSYSAPI GUID GUID_WINDOWS_RESUME_TARGET_TEMPLATE_EFI;
NTSYSAPI GUID GUID_WINDOWS_RESUME_TARGET_TEMPLATE_PCAT;
NTSYSAPI GUID GUID_WINDOWS_SETUP_EFI;
NTSYSAPI GUID GUID_WINDOWS_SETUP_PCAT;
#endif
typedef enum _BCD_MESSAGE_TYPE
{
    BCD_MESSAGE_TYPE_NONE,
    BCD_MESSAGE_TYPE_TRACE,
    BCD_MESSAGE_TYPE_INFORMATION,
    BCD_MESSAGE_TYPE_WARNING,
    BCD_MESSAGE_TYPE_ERROR,
    BCD_MESSAGE_TYPE_MAXIMUM
} BCD_MESSAGE_TYPE;*/
return true
}

func (n *ntbcd)typedef VOID ()(ok bool){//col:176
/*typedef VOID (NTAPI* BCD_MESSAGE_CALLBACK)(
    _In_ BCD_MESSAGE_TYPE type,
    _In_ PWSTR Message
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdSetLogging(
    _In_ BCD_MESSAGE_TYPE BcdLoggingLevel,
    _In_ BCD_MESSAGE_CALLBACK BcdMessageCallbackRoutine
    );
NTSYSAPI
VOID
NTAPI
BcdInitializeBcdSyncMutant(
    VOID
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdGetSystemStorePath(
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdSetSystemStoreDevice(
    _In_ UNICODE_STRING SystemPartition
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenSystemStore(
    _Out_ PHANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenStoreFromFile(
    _In_ UNICODE_STRING BcdFilePath,
    _Out_ PHANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCreateStore(
    _In_ UNICODE_STRING BcdFilePath,
    _Out_ PHANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdExportStore(
    _In_ UNICODE_STRING BcdFilePath
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdExportStoreEx(
    _In_ HANDLE BcdStoreHandle,
    _In_ ULONG Flags,
    _In_ UNICODE_STRING BcdFilePath
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdImportStore(
    _In_ UNICODE_STRING BcdFilePath
    );
typedef enum _BCD_IMPORT_FLAGS
{
    BCD_IMPORT_NONE,
    BCD_IMPORT_DELETE_FIRMWARE_OBJECTS
} BCD_IMPORT_FLAGS;*/
return true
}

func (n *ntbcd)BcdImportStoreWithFlags()(ok bool){//col:206
/*BcdImportStoreWithFlags(
    _In_ UNICODE_STRING BcdFilePath,
    _In_ BCD_IMPORT_FLAGS BcdImportFlags
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteObjectReferences(
    _In_ HANDLE BcdStoreHandle,
    _In_ PGUID Identifier
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteSystemStore(
    VOID
    );
typedef enum _BCD_OPEN_FLAGS
{
    BCD_OPEN_NONE,
    BCD_OPEN_OPEN_STORE_OFFLINE,
    BCD_OPEN_SYNC_FIRMWARE_ENTRIES
} BCD_OPEN_FLAGS;*/
return true
}

func (n *ntbcd)BcdOpenStore()(ok bool){//col:251
/*BcdOpenStore(
    _In_ UNICODE_STRING BcdFilePath,
    _In_ BCD_OPEN_FLAGS BcdOpenFlags,
    _Out_ PHANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCloseStore(
    _In_ HANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdFlushStore(
    _In_ HANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdForciblyUnloadStore(
    _In_ HANDLE BcdStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdMarkAsSystemStore(
    _In_ HANDLE BcdStoreHandle
    );
typedef enum _BCD_OBJECT_TYPE
{
    BCD_OBJECT_TYPE_NONE,
    BCD_OBJECT_TYPE_APPLICATION,
    BCD_OBJECT_TYPE_INHERITED,
    BCD_OBJECT_TYPE_DEVICE
} BCD_OBJECT_TYPE;*/
return true
}

func (n *ntbcd)#define MAKE_BCD_OBJECT()(ok bool){//col:334
/*#define MAKE_BCD_OBJECT(ObjectType, ImageType, ApplicationType) \
    (((ULONG)(ObjectType) << 28) | \
    (((ULONG)(ImageType) & 0xF) << 20) | \
    ((ULONG)(ApplicationType) & 0xFFFFF))
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
typedef union _BCD_OBJECT_DATATYPE
{
    ULONG PackedValue;
    union
    {
        struct
        {
            ULONG Reserved : 28;
            BCD_OBJECT_TYPE ObjectType : 4;
        };
        struct
        {
            BCD_APPLICATION_OBJECT_TYPE ApplicationType : 20;
            BCD_APPLICATION_IMAGE_TYPE ImageType : 4;
            ULONG Reserved : 4;
            BCD_OBJECT_TYPE ObjectType : 4;
        } Application;
        struct
        {
            ULONG Value : 20;
            BCD_INHERITED_CLASS_TYPE Class : 4;
            ULONG Reserved : 4;
            BCD_OBJECT_TYPE ObjectType : 4;
        } Inherit;
        struct
        {
            ULONG Reserved : 28;
            BCD_OBJECT_TYPE ObjectType : 4;
        } Device;
    };
} BCD_OBJECT_DATATYPE, *PBCD_OBJECT_DATATYPE;*/
return true
}

func (n *ntbcd)BcdEnumerateObjects()(ok bool){//col:403
/*BcdEnumerateObjects(
    _In_ HANDLE BcdStoreHandle,
    _In_ PBCD_OBJECT_DESCRIPTION BcdEnumDescriptor,
    _Inout_ PULONG BufferSize,
    _Out_ PULONG ObjectCount
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdOpenObject(
    _In_ HANDLE BcdStoreHandle,
    _In_ PGUID Identifier,
    _Out_ PHANDLE BcdObjectHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCreateObject(
    _In_ HANDLE BcdStoreHandle,
    _In_ PGUID Identifier,
    _In_ PBCD_OBJECT_DESCRIPTION Description,
    _Out_ PHANDLE BcdObjectHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteObject(
    _In_ HANDLE BcdObjectHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCloseObject(
    _In_ HANDLE BcdObjectHandle
    );
typedef enum _BCD_COPY_FLAGS
{
    BCD_COPY_NONE = 0x0,
    BCD_COPY_COPY_CREATE_NEW_OBJECT_IDENTIFIER = 0x1,
    BCD_COPY_COPY_DELETE_EXISTING_OBJECT = 0x2,
    BCD_COPY_COPY_UNKNOWN_FIRMWARE_APPLICATION = 0x4,
    BCD_COPY_IGNORE_SETUP_TEMPLATE_ELEMENTS = 0x8,
    BCD_COPY_RETAIN_ELEMENT_DATA = 0x10,
    BCD_COPY_MIGRATE_ELEMENT_DATA = 0x20
} BCD_COPY_FLAGS;*/
return true
}

func (n *ntbcd)BcdCopyObject()(ok bool){//col:468
/*BcdCopyObject(
    _In_ HANDLE BcdStoreHandle,
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_COPY_FLAGS BcdCopyFlags,
    _In_ HANDLE TargetStoreHandle,
    _Out_ PHANDLE TargetHandleOut
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCopyObjectEx(
    _In_ HANDLE BcdStoreHandle,
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_COPY_FLAGS BcdCopyFlags,
    _In_ HANDLE TargetStoreHandle,
    _In_ PGUID TargetObjectId,
    _Out_ PHANDLE TargetHandleOut
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdCopyObjects(
    _In_ HANDLE BcdStoreHandle,
    _In_ BCD_OBJECT_DESCRIPTION Characteristics,
    _In_ BCD_COPY_FLAGS BcdCopyFlags,
    _In_ HANDLE TargetStoreHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdMigrateObjectElementValues(
    _In_ HANDLE TemplateObjectHandle,
    _In_ HANDLE SourceObjectHandle,
    _In_ HANDLE TargetObjectHandle
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdQueryObject(
    _In_ HANDLE BcdObjectHandle,
    _Out_ BCD_OBJECT_DESCRIPTION Description,
    _Out_ PGUID Identifier
    );
typedef enum _BCD_ELEMENT_DATATYPE_FORMAT
{
    BCD_ELEMENT_DATATYPE_FORMAT_UNKNOWN,
} BCD_ELEMENT_DATATYPE_FORMAT;*/
return true
}

func (n *ntbcd)#define MAKE_BCDE_DATA_TYPE()(ok bool){//col:514
/*#define MAKE_BCDE_DATA_TYPE(Class, Format, Subtype) \
    (((((ULONG)Class) & 0xF) << 28) | ((((ULONG)Format) & 0xF) << 24) | (((ULONG)Subtype) & 0x00FFFFFF))
#define GET_BCDE_DATA_CLASS(DataType) \
    ((BCD_ELEMENT_DATATYPE_CLASS)(((((ULONG)DataType)) >> 28) & 0xF))
#define GET_BCDE_DATA_FORMAT(DataType) \
    ((BCD_ELEMENT_DATATYPE_FORMAT)(((((ULONG)DataType)) >> 24) & 0xF))
#define GET_BCDE_DATA_SUBTYPE(DataType) \
    ((ULONG)((((ULONG)DataType)) & 0x00FFFFFF))
typedef union _BCD_ELEMENT_DATATYPE
{
    ULONG PackedValue;
    struct
    {
        ULONG SubType : 24;
        BCD_ELEMENT_DATATYPE_FORMAT Format : 4;
        BCD_ELEMENT_DATATYPE_CLASS Class : 4;
    };
} BCD_ELEMENT_DATATYPE, *PBCD_ELEMENT_DATATYPE;*/
return true
}

func (n *ntbcd)BcdEnumerateElementTypes()(ok bool){//col:543
/*BcdEnumerateElementTypes(
    _In_ HANDLE BcdObjectHandle,
    _Inout_ PULONG BufferSize,
    _Out_ PULONG ElementCount
    );
typedef struct _BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION
{
    ULONG PartitionStyle;
    ULONG Reserved;
    struct
    {
        union
        {
            ULONG DiskSignature;
            ULONG64 PartitionOffset;
        } Mbr;
        union
        {
            GUID DiskSignature;
            GUID PartitionSignature;
        } Gpt;
    };
} BCD_ELEMENT_DEVICE_QUALIFIED_PARTITION, *PBCD_ELEMENT_DEVICE_QUALIFIED_PARTITION;*/
return true
}

func (n *ntbcd)BcdEnumerateElements()(ok bool){//col:646
/*BcdEnumerateElements(
    _In_ HANDLE BcdObjectHandle,
    _Inout_ PULONG BufferSize,
    _Out_ PULONG ElementCount
    );
typedef enum _BCD_FLAGS
{
    BCD_FLAG_NONE = 0x0,
    BCD_FLAG_QUALIFIED_PARTITION = 0x1,
    BCD_FLAG_NO_DEVICE_TRANSLATION = 0x2,
    BCD_FLAG_ENUMERATE_INHERITED_OBJECTS = 0x4,
    BCD_FLAG_ENUMERATE_DEVICE_OPTIONS = 0x8,
    BCD_FLAG_OBSERVE_PRECEDENCE = 0x10,
    BCD_FLAG_DISABLE_VHD_NT_TRANSLATION = 0x20,
    BCD_FLAG_DISABLE_VHD_DEVICE_DETECTION = 0x40,
    BCD_FLAG_DISABLE_POLICY_CHECKS = 0x80
} BCD_FLAGS;*/
return true
}

func (n *ntbcd)BcdEnumerateElementsWithFlags()(ok bool){//col:833
/*BcdEnumerateElementsWithFlags(
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_FLAGS BcdFlags,
    _Inout_ PULONG BufferSize,
    _Out_ PULONG ElementCount
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdEnumerateAndUnpackElements(
    _In_ HANDLE BcdStoreHandle,
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_FLAGS BcdFlags,
    _Inout_ PULONG BufferSize,
    _Out_ PULONG ElementCount
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdGetElementData(
    _In_ HANDLE BcdObjectHandle,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer,
    _Inout_ PULONG BufferSize
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdGetElementDataWithFlags(
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_FLAGS BcdFlags,
    _Out_writes_bytes_opt_(*BufferSize) PVOID Buffer,
    _Inout_ PULONG BufferSize
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdSetElementData(
    _In_ HANDLE BcdObjectHandle,
    _In_reads_bytes_opt_(BufferSize) PVOID Buffer,
    _In_ ULONG BufferSize
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdSetElementDataWithFlags(
    _In_ HANDLE BcdObjectHandle,
    _In_ BCD_FLAGS BcdFlags,
    _In_reads_bytes_opt_(BufferSize) PVOID Buffer,
    _In_ ULONG BufferSize
    );
NTSYSAPI
NTSTATUS
NTAPI
BcdDeleteElement(
    _In_ HANDLE BcdObjectHandle,
    );
typedef enum _BcdBootMgrElementTypes
{
    BcdBootMgrObjectList_DisplayOrder = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 1),
    BcdBootMgrObjectList_BootSequence = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 2),
    BcdBootMgrObject_DefaultObject = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 3),
    BcdBootMgrInteger_Timeout = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 4),
    BcdBootMgrBoolean_AttemptResume = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 5),
    BcdBootMgrObject_ResumeObject = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 6),
    BcdBootMgrObjectList_StartupSequence = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 7),
    BcdBootMgrObjectList_ToolsDisplayOrder = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 16),
    BcdBootMgrBoolean_DisplayBootMenu = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 32),
    BcdBootMgrBoolean_NoErrorDisplay = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 33),
    BcdBootMgrDevice_BcdDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 34),
    BcdBootMgrString_BcdFilePath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 35),
    BcdBootMgrBoolean_HormEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 36),
    BcdBootMgrBoolean_HiberRoot = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 37),
    BcdBootMgrString_PasswordOverride = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 38),
    BcdBootMgrString_PinpassPhraseOverride = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 39),
    BcdBootMgrBoolean_ProcessCustomActionsFirst = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 40),
    BcdBootMgrIntegerList_CustomActionsList = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 48),
    BcdBootMgrBoolean_PersistBootSequence = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 49),
    BcdBootMgrBoolean_SkipStartupSequence = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 50),
} BcdBootMgrElementTypes;*/
return true
}

func (n *ntbcd)    BcdLibraryDevice_ApplicationDevice = MAKE_BCDE_DATA_TYPE()(ok bool){//col:1361
/*    BcdLibraryDevice_ApplicationDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 1),
    BcdLibraryString_ApplicationPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdLibraryString_Description = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 4),
    BcdLibraryString_PreferredLocale = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 5),
    BcdLibraryObjectList_InheritedObjects = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 6),
    BcdLibraryInteger_TruncatePhysicalMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 7),
    BcdLibraryObjectList_RecoverySequence = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_OBJECTLIST, 8),
    BcdLibraryBoolean_AutoRecoveryEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 9),
    BcdLibraryIntegerList_BadMemoryList = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 10),
    BcdLibraryBoolean_AllowBadMemoryAccess = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 11),
    BcdLibraryInteger_FirstMegabytePolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 12),
    BcdLibraryInteger_RelocatePhysicalMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 13),
    BcdLibraryInteger_AvoidLowPhysicalMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 14),
    BcdLibraryBoolean_TraditionalKsegMappings = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 15),
    BcdLibraryBoolean_DebuggerEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
    BcdLibraryInteger_DebuggerType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 17),
    BcdLibraryInteger_SerialDebuggerPortAddress = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 18),
    BcdLibraryInteger_SerialDebuggerPort = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 19),
    BcdLibraryInteger_SerialDebuggerBaudRate = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 20),
    BcdLibraryInteger_1394DebuggerChannel = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 21),
    BcdLibraryString_UsbDebuggerTargetName = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 22),
    BcdLibraryBoolean_DebuggerIgnoreUsermodeExceptions = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 23),
    BcdLibraryInteger_DebuggerStartPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 24),
    BcdLibraryString_DebuggerBusParameters = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 25),
    BcdLibraryInteger_DebuggerNetHostIP = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 26),
    BcdLibraryInteger_DebuggerNetPort = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 27),
    BcdLibraryBoolean_DebuggerNetDhcp = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 28),
    BcdLibraryString_DebuggerNetKey = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 29),
    BcdLibraryBoolean_DebuggerNetVM = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 30),
    BcdLibraryString_DebuggerNetHostIpv6 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 31),
    BcdLibraryBoolean_EmsEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 32),
    BcdLibraryInteger_EmsPort = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 34),
    BcdLibraryInteger_EmsBaudRate = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 35),
    BcdLibraryString_LoadOptionsString = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 48),
    BcdLibraryBoolean_AttemptNonBcdStart = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 49),
    BcdLibraryBoolean_DisplayAdvancedOptions = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 64),
    BcdLibraryBoolean_DisplayOptionsEdit = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 65),
    BcdLibraryInteger_FVEKeyRingAddress = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 66),
    BcdLibraryDevice_BsdLogDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 67),
    BcdLibraryString_BsdLogPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 68),
    BcdLibraryBoolean_BsdPreserveLog = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 69),
    BcdLibraryBoolean_GraphicsModeDisabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 70),
    BcdLibraryInteger_ConfigAccessPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 71),
    BcdLibraryBoolean_DisableIntegrityChecks = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 72),
    BcdLibraryBoolean_AllowPrereleaseSignatures = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 73),
    BcdLibraryString_FontPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 74),
    BcdLibraryInteger_SiPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 75),
    BcdLibraryInteger_FveBandId = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 76),
    BcdLibraryBoolean_ConsoleExtendedInput = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 80),
    BcdLibraryInteger_InitialConsoleInput = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 81),
    BcdLibraryInteger_GraphicsResolution = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 82),
    BcdLibraryBoolean_RestartOnFailure = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 83),
    BcdLibraryBoolean_GraphicsForceHighestMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 84),
    BcdLibraryBoolean_IsolatedExecutionContext = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 96),
    BcdLibraryInteger_BootUxDisplayMessage = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 101),
    BcdLibraryInteger_BootUxDisplayMessageOverride = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 102),
    BcdLibraryBoolean_BootUxLogoDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 103),
    BcdLibraryBoolean_BootUxTextDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 104),
    BcdLibraryBoolean_BootUxProgressDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 105),
    BcdLibraryBoolean_BootUxFadeDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 106),
    BcdLibraryBoolean_BootUxReservePoolDebug = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 107),
    BcdLibraryBoolean_BootUxDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 108),
    BcdLibraryInteger_BootUxFadeFrames = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 109),
    BcdLibraryBoolean_BootUxDumpStats = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 110),
    BcdLibraryBoolean_BootUxShowStats = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 111),
    BcdLibraryBoolean_MultiBootSystem = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 113),
    BcdLibraryBoolean_ForceNoKeyboard = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 114),
    BcdLibraryInteger_AliasWindowsKey = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 115),
    BcdLibraryBoolean_BootShutdownDisabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 116),
    BcdLibraryInteger_PerformanceFrequency = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 117),
    BcdLibraryInteger_SecurebootRawPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 118),
    BcdLibraryIntegerList_AllowedInMemorySettings = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 119),
    BcdLibraryInteger_BootUxBitmapTransitionTime = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 121),
    BcdLibraryBoolean_TwoBootImages = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 122),
    BcdLibraryBoolean_ForceFipsCrypto = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 123),
    BcdLibraryInteger_BootErrorUx = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 125),
    BcdLibraryBoolean_AllowFlightSignatures = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 126),
    BcdLibraryInteger_BootMeasurementLogFormat = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 127),
    BcdLibraryInteger_DisplayRotation = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 128),
    BcdLibraryInteger_LogControl = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 129),
    BcdLibraryBoolean_NoFirmwareSync = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 130),
    BcdLibraryDevice_WindowsSystemDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 132),
    BcdLibraryBoolean_NumLockOn = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 135),
    BcdLibraryString_AdditionalCiPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_LIBRARY, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 136),
} BcdLibraryElementTypes;*/
return true
}

func (n *ntbcd)    BcdSetupInteger_DeviceType = MAKE_BCDE_DATA_TYPE()(ok bool){//col:1395
/*    BcdSetupInteger_DeviceType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 1),
    BcdSetupString_ApplicationRelativePath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdSetupString_RamdiskDeviceRelativePath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 3),
    BcdSetupBoolean_OmitOsLoaderElements = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 4),
    BcdSetupIntegerList_ElementsToMigrateList = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_INTEGERLIST, 6),
    BcdSetupBoolean_RecoveryOs = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_SETUPTEMPLATE, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
} BcdTemplateElementTypes;*/
return true
}

func (n *ntbcd)    BcdOSLoaderDevice_OSDevice = MAKE_BCDE_DATA_TYPE()(ok bool){//col:2128
/*    BcdOSLoaderDevice_OSDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 1),
    BcdOSLoaderString_SystemRoot = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 2),
    BcdOSLoaderObject_AssociatedResumeObject = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_OBJECT, 3),
    BcdOSLoaderBoolean_StampDisks = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 4),
    BcdOSLoaderBoolean_DetectKernelAndHal = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 16),
    BcdOSLoaderString_KernelPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 17),
    BcdOSLoaderString_HalPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 18),
    BcdOSLoaderString_DbgTransportPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 19),
    BcdOSLoaderInteger_NxPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 32),
    BcdOSLoaderInteger_PAEPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 33),
    BcdOSLoaderBoolean_WinPEMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 34),
    BcdOSLoaderBoolean_DisableCrashAutoReboot = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 36),
    BcdOSLoaderBoolean_UseLastGoodSettings = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 37),
    BcdOSLoaderBoolean_DisableCodeIntegrityChecks = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 38),
    BcdOSLoaderBoolean_AllowPrereleaseSignatures = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 39),
    BcdOSLoaderBoolean_NoLowMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 48),
    BcdOSLoaderInteger_RemoveMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 49),
    BcdOSLoaderInteger_IncreaseUserVa = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 50),
    BcdOSLoaderInteger_PerformaceDataMemory = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 51),
    BcdOSLoaderBoolean_UseVgaDriver = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 64),
    BcdOSLoaderBoolean_DisableBootDisplay = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 65),
    BcdOSLoaderBoolean_DisableVesaBios = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 66),
    BcdOSLoaderBoolean_DisableVgaMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 67),
    BcdOSLoaderInteger_ClusterModeAddressing = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 80),
    BcdOSLoaderBoolean_UsePhysicalDestination = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 81),
    BcdOSLoaderInteger_RestrictApicCluster = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 82),
    BcdOSLoaderString_OSLoaderTypeEVStore = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 83),
    BcdOSLoaderBoolean_UseLegacyApicMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 84),
    BcdOSLoaderInteger_X2ApicPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 85),
    BcdOSLoaderBoolean_UseBootProcessorOnly = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 96),
    BcdOSLoaderInteger_NumberOfProcessors = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 97),
    BcdOSLoaderBoolean_ForceMaximumProcessors = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 98),
    BcdOSLoaderBoolean_ProcessorConfigurationFlags = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 99),
    BcdOSLoaderBoolean_MaximizeGroupsCreated = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 100),
    BcdOSLoaderBoolean_ForceGroupAwareness = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 101),
    BcdOSLoaderInteger_GroupSize = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 102),
    BcdOSLoaderInteger_UseFirmwarePciSettings = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 112),
    BcdOSLoaderInteger_MsiPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 113),
    BcdOSLoaderInteger_PciExpressPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 114),
    BcdOSLoaderInteger_SafeBoot = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 128),
    BcdOSLoaderBoolean_SafeBootAlternateShell = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 129),
    BcdOSLoaderBoolean_BootLogInitialization = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 144),
    BcdOSLoaderBoolean_VerboseObjectLoadMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 145),
    BcdOSLoaderBoolean_KernelDebuggerEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 160),
    BcdOSLoaderBoolean_DebuggerHalBreakpoint = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 161),
    BcdOSLoaderBoolean_UsePlatformClock = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 162),
    BcdOSLoaderBoolean_ForceLegacyPlatform = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 163),
    BcdOSLoaderBoolean_UsePlatformTick = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 164),
    BcdOSLoaderBoolean_DisableDynamicTick = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 165),
    BcdOSLoaderInteger_TscSyncPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 166),
    BcdOSLoaderBoolean_EmsEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 176),
    BcdOSLoaderInteger_ForceFailure = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 192),
    BcdOSLoaderInteger_DriverLoadFailurePolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 193),
    BcdOSLoaderInteger_BootMenuPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 194),
    BcdOSLoaderBoolean_AdvancedOptionsOneTime = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 195),
    BcdOSLoaderBoolean_OptionsEditOneTime = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 196),
    BcdOSLoaderInteger_BootStatusPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 224),
    BcdOSLoaderBoolean_DisableElamDrivers = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 225),
    BcdOSLoaderInteger_HypervisorLaunchType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 240),
    BcdOSLoaderString_HypervisorPath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 241),
    BcdOSLoaderBoolean_HypervisorDebuggerEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 242),
    BcdOSLoaderInteger_HypervisorDebuggerType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 243),
    BcdOSLoaderInteger_HypervisorDebuggerPortNumber = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 244),
    BcdOSLoaderInteger_HypervisorDebuggerBaudrate = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 245),
    BcdOSLoaderInteger_HypervisorDebugger1394Channel = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 246),
    BcdOSLoaderInteger_BootUxPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 247),
    BcdOSLoaderInteger_HypervisorSlatDisabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 248),
    BcdOSLoaderString_HypervisorDebuggerBusParams = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 249),
    BcdOSLoaderInteger_HypervisorNumProc = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 250),
    BcdOSLoaderInteger_HypervisorRootProcPerNode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 251),
    BcdOSLoaderBoolean_HypervisorUseLargeVTlb = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 252),
    BcdOSLoaderInteger_HypervisorDebuggerNetHostIp = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 253),
    BcdOSLoaderInteger_HypervisorDebuggerNetHostPort = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 254),
    BcdOSLoaderInteger_HypervisorDebuggerPages = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 255),
    BcdOSLoaderInteger_TpmBootEntropyPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 256),
    BcdOSLoaderString_HypervisorDebuggerNetKey = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 272),
    BcdOSLoaderString_HypervisorProductSkuType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 274),
    BcdOSLoaderInteger_HypervisorRootProc = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 275),
    BcdOSLoaderBoolean_HypervisorDebuggerNetDhcp = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 276),
    BcdOSLoaderInteger_HypervisorIommuPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 277),
    BcdOSLoaderBoolean_HypervisorUseVApic = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 278),
    BcdOSLoaderString_HypervisorLoadOptions = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 279),
    BcdOSLoaderInteger_HypervisorMsrFilterPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 280),
    BcdOSLoaderInteger_HypervisorMmioNxPolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 281),
    BcdOSLoaderInteger_HypervisorSchedulerType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 282),
    BcdOSLoaderString_HypervisorRootProcNumaNodes = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 283),
    BcdOSLoaderInteger_HypervisorPerfmon = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 284),
    BcdOSLoaderInteger_HypervisorRootProcPerCore = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 285),
    BcdOSLoaderString_HypervisorRootProcNumaNodeLps = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 286),
    BcdOSLoaderInteger_XSavePolicy = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 288),
    BcdOSLoaderInteger_XSaveAddFeature0 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 289),
    BcdOSLoaderInteger_XSaveAddFeature1 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 290),
    BcdOSLoaderInteger_XSaveAddFeature2 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 291),
    BcdOSLoaderInteger_XSaveAddFeature3 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 292),
    BcdOSLoaderInteger_XSaveAddFeature4 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 293),
    BcdOSLoaderInteger_XSaveAddFeature5 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 294),
    BcdOSLoaderInteger_XSaveAddFeature6 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 295),
    BcdOSLoaderInteger_XSaveAddFeature7 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 296),
    BcdOSLoaderInteger_XSaveRemoveFeature = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 297),
    BcdOSLoaderInteger_XSaveProcessorsMask = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 298),
    BcdOSLoaderInteger_XSaveDisable = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 299),
    BcdOSLoaderInteger_KernelDebuggerType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 300),
    BcdOSLoaderString_KernelDebuggerBusParameters = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 301),
    BcdOSLoaderInteger_KernelDebuggerPortAddress = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 302),
    BcdOSLoaderInteger_KernelDebuggerPortNumber = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 303),
    BcdOSLoaderInteger_ClaimedTpmCounter = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 304),
    BcdOSLoaderInteger_KernelDebugger1394Channel = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 305),
    BcdOSLoaderString_KernelDebuggerUsbTargetname = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 306),
    BcdOSLoaderInteger_KernelDebuggerNetHostIp = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 307),
    BcdOSLoaderInteger_KernelDebuggerNetHostPort = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 308),
    BcdOSLoaderBoolean_KernelDebuggerNetDhcp = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 309),
    BcdOSLoaderString_KernelDebuggerNetKey = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 310),
    BcdOSLoaderString_IMCHiveName = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 311),
    BcdOSLoaderDevice_IMCDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 312),
    BcdOSLoaderInteger_KernelDebuggerBaudrate = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 313),
    BcdOSLoaderString_ManufacturingMode = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 320),
    BcdOSLoaderBoolean_EventLoggingEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 321),
    BcdOSLoaderInteger_VsmLaunchType = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 322),
    BcdOSLoaderInteger_HypervisorEnforcedCodeIntegrity = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_INTEGER, 324),
    BcdOSLoaderBoolean_DtraceEnabled = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_BOOLEAN, 325),
    BcdOSLoaderDevice_SystemDataDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 336),
    BcdOSLoaderDevice_OsArcDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 337),
    BcdOSLoaderDevice_OsDataDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 339),
    BcdOSLoaderDevice_BspDevice = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 340),
    BcdOSLoaderDevice_BspFilepath = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_DEVICE, 341),
    BcdOSLoaderString_KernelDebuggerNetHostIpv6 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 342),
    BcdOSLoaderString_HypervisorDebuggerNetHostIpv6 = MAKE_BCDE_DATA_TYPE(BCD_ELEMENT_DATATYPE_CLASS_APPLICATION, BCD_ELEMENT_DATATYPE_FORMAT_STRING, 353),
} BcdOSLoaderElementTypes;*/
return true
}



