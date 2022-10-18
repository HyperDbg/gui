package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntldr.h.back

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



type  _LDR_SERVICE_TAG_RECORD struct{
_LDR_SERVICE_TAG_RECORD struct //col:7
ServiceTag uint32 //col:8
}


type  _LDRP_CSLIST struct{
Tail PSINGLE_LIST_ENTRY //col:11
}


type  _LDR_DDAG_NODE struct{
Modules *list.List //col:23
ServiceTagList PLDR_SERVICE_TAG_RECORD //col:24
LoadCount uint32 //col:25
LoadWhileUnloadingCount uint32 //col:26
LowestLink uint32 //col:27
Union union //col:28
Dependencies LDRP_CSLIST //col:30
RemovalLink SINGLE_LIST_ENTRY //col:31
}


type  _LDR_DEPENDENCY_RECORD struct{
DependencyLink SINGLE_LIST_ENTRY //col:35
DependencyNode PLDR_DDAG_NODE //col:36
IncomingDependencyLink SINGLE_LIST_ENTRY //col:37
IncomingDependencyNode PLDR_DDAG_NODE //col:38
}


type  _LDR_DATA_TABLE_ENTRY struct{
InLoadOrderLinks *list.List //col:44
InMemoryOrderLinks *list.List //col:45
Union union //col:46
InInitializationOrderLinks *list.List //col:48
InProgressLinks *list.List //col:49
}


type  _LDR_IMPORT_CALLBACK_INFO struct{
ImportCallbackRoutine PLDR_IMPORT_MODULE_CALLBACK //col:116
ImportCallbackParameter uintptr //col:117
}


type  _LDR_SECTION_INFO struct{
SectionHandle uintptr //col:124
DesiredAccess ACCESS_MASK //col:125
ObjA POBJECT_ATTRIBUTES //col:126
SectionPageProtection uint32 //col:127
AllocationAttributes uint32 //col:128
}


type  _LDR_VERIFY_IMAGE_INFO struct{
Size uint32 //col:132
Flags uint32 //col:133
CallbackInfo LDR_IMPORT_CALLBACK_INFO //col:134
SectionInfo LDR_SECTION_INFO //col:135
ImageCharacteristics uint16 //col:136
}


type  _LDR_DLL_LOADED_NOTIFICATION_DATA struct{
Flags uint32 //col:140
FullDllName *uint32 //col:141
BaseDllName *uint32 //col:142
DllBase uintptr //col:143
SizeOfImage uint32 //col:144
}


type  _LDR_DLL_UNLOADED_NOTIFICATION_DATA struct{
Flags uint32 //col:148
FullDllName PCUNICODE_STRING //col:149
BaseDllName PCUNICODE_STRING //col:150
DllBase uintptr //col:151
SizeOfImage uint32 //col:152
}


type  _LDR_FAILURE_DATA struct{
Status NTSTATUS //col:154
DllName[0x20] WCHAR //col:155
AdditionalInfo[0x20] WCHAR //col:156
}


type  _PS_MITIGATION_OPTIONS_MAP struct{
Map[3] ULONG_PTR //col:158
}


type  _PS_MITIGATION_AUDIT_OPTIONS_MAP struct{
Map[3] ULONG_PTR //col:162
}


type  _PS_SYSTEM_DLL_INIT_BLOCK struct{
Size uint32 //col:177
SystemDllWowRelocation ULONG_PTR //col:178
SystemDllNativeRelocation ULONG_PTR //col:179
Wow64SharedInformation[16] ULONG_PTR //col:180
RngData uint32 //col:181
Union union //col:182
Flags uint32 //col:184
Struct struct //col:185
CfgOverride uint32 //col:187
Reserved uint32 //col:188
}


type  _LDR_RESOURCE_INFO struct{
Type ULONG_PTR //col:191
Name ULONG_PTR //col:192
Language ULONG_PTR //col:193
}


type  _LDR_ENUM_RESOURCE_ENTRY struct{
Union union //col:202
NameOrId ULONG_PTR //col:204
Name PIMAGE_RESOURCE_DIRECTORY_STRING //col:205
Struct struct //col:206
Id uint16 //col:208
NameIsPresent uint16 //col:209
}


type  _RTL_PROCESS_MODULE_INFORMATION struct{
Section uintptr //col:220
MappedBase uintptr //col:221
ImageBase uintptr //col:222
ImageSize uint32 //col:223
Flags uint32 //col:224
LoadOrderIndex uint16 //col:225
InitOrderIndex uint16 //col:226
LoadCount uint16 //col:227
OffsetToFileName uint16 //col:228
FullPathName[256] uint8 //col:229
}


type  _RTL_PROCESS_MODULES struct{
NumberOfModules uint32 //col:225
Modules[1] RTL_PROCESS_MODULE_INFORMATION //col:226
}


type  _RTL_PROCESS_MODULE_INFORMATION_EX struct{
NextOffset uint16 //col:233
BaseInfo RTL_PROCESS_MODULE_INFORMATION //col:234
ImageChecksum uint32 //col:235
TimeDateStamp uint32 //col:236
DefaultBase uintptr //col:237
}


type  _DELAYLOAD_PROC_DESCRIPTOR struct{
ImportDescribedByName uint32 //col:241
Union union //col:242
Name PCSTR //col:244
Ordinal uint32 //col:245
}


type  _DELAYLOAD_INFO struct{
Size uint32 //col:253
DelayloadDescriptor PCIMAGE_DELAYLOAD_DESCRIPTOR //col:254
ThunkAddress PIMAGE_THUNK_DATA //col:255
TargetDllName PCSTR //col:256
TargetApiDescriptor DELAYLOAD_PROC_DESCRIPTOR //col:257
TargetModuleBase uintptr //col:258
Unused uintptr //col:259
LastError uint32 //col:260
}




