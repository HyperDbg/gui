package phnt


const(
    LdrModulesMerged  =  -5  //col:3
    LdrModulesInitError  =  -4  //col:4
    LdrModulesSnapError  =  -3  //col:5
    LdrModulesUnloaded  =  -2  //col:6
    LdrModulesUnloading  =  -1  //col:7
    LdrModulesPlaceHolder  =  0  //col:8
    LdrModulesMapping  =  1  //col:9
    LdrModulesMapped  =  2  //col:10
    LdrModulesWaitingForDependencies  =  3  //col:11
    LdrModulesSnapping  =  4  //col:12
    LdrModulesSnapped  =  5  //col:13
    LdrModulesCondensed  =  6  //col:14
    LdrModulesReadyToInit  =  7  //col:15
    LdrModulesInitializing  =  8  //col:16
    LdrModulesReadyToRun  =  9  //col:17
)


const(
    LoadReasonStaticDependency = 1  //col:21
    LoadReasonStaticForwarderDependency = 2  //col:22
    LoadReasonDynamicForwarderDependency = 3  //col:23
    LoadReasonDelayloadDependency = 4  //col:24
    LoadReasonDynamicLoad = 5  //col:25
    LoadReasonAsImageLoad = 6  //col:26
    LoadReasonAsDataLoad = 7  //col:27
    LoadReasonEnclavePrimary  = 8  //col:28
    LoadReasonEnclaveDependency = 9  //col:29
    LoadReasonPatchImage  = 10  //col:30
    LoadReasonUnknown  =  -1  //col:31
)


const(
    LdrHotPatchBaseImage = 1  //col:35
    LdrHotPatchNotApplied = 2  //col:36
    LdrHotPatchAppliedReverse = 3  //col:37
    LdrHotPatchAppliedForward = 4  //col:38
    LdrHotPatchFailedToPatch = 5  //col:39
    LdrHotPatchStateMax = 6  //col:40
)



type LDR_SERVICE_TAG_RECORD struct{
struct // //col:3
ServiceTag uint32 //col:4
}



type LDRP_CSLIST struct{
Tail PSINGLE_LIST_ENTRY //col:8
}



type LDR_DDAG_NODE struct{
Modules *list.List //col:12
ServiceTagList PLDR_SERVICE_TAG_RECORD //col:13
LoadCount uint32 //col:14
LoadWhileUnloadingCount uint32 //col:15
LowestLink uint32 //col:16
Union union //col:17
Dependencies LDRP_CSLIST //col:19
RemovalLink SINGLE_LIST_ENTRY //col:20
}



type LDR_DEPENDENCY_RECORD struct{
DependencyLink SINGLE_LIST_ENTRY //col:29
DependencyNode PLDR_DDAG_NODE //col:30
IncomingDependencyLink SINGLE_LIST_ENTRY //col:31
IncomingDependencyNode PLDR_DDAG_NODE //col:32
}



type LDR_DATA_TABLE_ENTRY struct{
InLoadOrderLinks *list.List //col:36
InMemoryOrderLinks *list.List //col:37
Union union //col:38
InInitializationOrderLinks *list.List //col:40
InProgressLinks *list.List //col:41
}



type LDR_IMPORT_CALLBACK_INFO struct{
ImportCallbackRoutine PLDR_IMPORT_MODULE_CALLBACK //col:112
ImportCallbackParameter PVOID //col:113
}



type LDR_SECTION_INFO struct{
SectionHandle HANDLE //col:117
DesiredAccess ACCESS_MASK //col:118
ObjA POBJECT_ATTRIBUTES //col:119
SectionPageProtection uint32 //col:120
AllocationAttributes uint32 //col:121
}



type LDR_VERIFY_IMAGE_INFO struct{
Size uint32 //col:125
Flags uint32 //col:126
CallbackInfo LDR_IMPORT_CALLBACK_INFO //col:127
SectionInfo LDR_SECTION_INFO //col:128
ImageCharacteristics USHORT //col:129
}



type LDR_DLL_LOADED_NOTIFICATION_DATA struct{
Flags uint32 //col:133
FullDllName PUNICODE_STRING //col:134
BaseDllName PUNICODE_STRING //col:135
DllBase PVOID //col:136
SizeOfImage uint32 //col:137
}



type LDR_DLL_UNLOADED_NOTIFICATION_DATA struct{
Flags uint32 //col:141
FullDllName PCUNICODE_STRING //col:142
BaseDllName PCUNICODE_STRING //col:143
DllBase PVOID //col:144
SizeOfImage uint32 //col:145
}



type LDR_FAILURE_DATA struct{
Status NTSTATUS //col:149
DllName[0x20] WCHAR //col:150
AdditionalInfo[0x20] WCHAR //col:151
}



type PS_MITIGATION_OPTIONS_MAP struct{
Map[3] ULONG_PTR //col:155
}



type PS_MITIGATION_AUDIT_OPTIONS_MAP struct{
Map[3] ULONG_PTR //col:159
}



type PS_SYSTEM_DLL_INIT_BLOCK struct{
Size uint32 //col:163
SystemDllWowRelocation ULONG_PTR //col:164
SystemDllNativeRelocation ULONG_PTR //col:165
Wow64SharedInformation[16] ULONG_PTR //col:166
RngData uint32 //col:167
Union union //col:168
Flags uint32 //col:170
Struct struct //col:171
CfgOverride uint32 //col:173
Reserved uint32 //col:174
}



type LDR_RESOURCE_INFO struct{
Type ULONG_PTR //col:186
Name ULONG_PTR //col:187
Language ULONG_PTR //col:188
}



type LDR_ENUM_RESOURCE_ENTRY struct{
Union union //col:192
NameOrId ULONG_PTR //col:194
Name PIMAGE_RESOURCE_DIRECTORY_STRING //col:195
Struct struct //col:196
Id USHORT //col:198
NameIsPresent USHORT //col:199
}



type RTL_PROCESS_MODULE_INFORMATION struct{
Section HANDLE //col:208
MappedBase PVOID //col:209
ImageBase PVOID //col:210
ImageSize uint32 //col:211
Flags uint32 //col:212
LoadOrderIndex USHORT //col:213
InitOrderIndex USHORT //col:214
LoadCount USHORT //col:215
OffsetToFileName USHORT //col:216
FullPathName[256] uint8 //col:217
}



type RTL_PROCESS_MODULES struct{
NumberOfModules uint32 //col:221
Modules[1] RTL_PROCESS_MODULE_INFORMATION //col:222
}



type RTL_PROCESS_MODULE_INFORMATION_EX struct{
NextOffset USHORT //col:226
BaseInfo RTL_PROCESS_MODULE_INFORMATION //col:227
ImageChecksum uint32 //col:228
TimeDateStamp uint32 //col:229
DefaultBase PVOID //col:230
}



type DELAYLOAD_PROC_DESCRIPTOR struct{
ImportDescribedByName uint32 //col:234
Union union //col:235
Name PCSTR //col:237
Ordinal uint32 //col:238
}



type DELAYLOAD_INFO struct{
Size uint32 //col:243
DelayloadDescriptor PCIMAGE_DELAYLOAD_DESCRIPTOR //col:244
ThunkAddress PIMAGE_THUNK_DATA //col:245
TargetDllName PCSTR //col:246
TargetApiDescriptor DELAYLOAD_PROC_DESCRIPTOR //col:247
TargetModuleBase PVOID //col:248
Unused PVOID //col:249
LastError uint32 //col:250
}





