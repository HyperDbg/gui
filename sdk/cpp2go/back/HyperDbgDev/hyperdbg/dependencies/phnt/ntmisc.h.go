package phnt
//back\HyperDbgDev\hyperdbg\dependencies\phnt\ntmisc.h.back

const(
_NTMISC_H =  //col:13
FLT_PORT_CONNECT = 0x0001 //col:17
FLT_PORT_ALL_ACCESS = (FLT_PORT_CONNECT | STANDARD_RIGHTS_ALL) //col:18
)

type     VdmStartExecution uint32
const(
    VdmStartExecution VDMSERVICECLASS = 1  //col:24
    VdmQueueInterrupt VDMSERVICECLASS = 2  //col:25
    VdmDelayInterrupt VDMSERVICECLASS = 3  //col:26
    VdmInitialize VDMSERVICECLASS = 4  //col:27
    VdmFeatures VDMSERVICECLASS = 5  //col:28
    VdmSetInt21Handler VDMSERVICECLASS = 6  //col:29
    VdmQueryDir VDMSERVICECLASS = 7  //col:30
    VdmPrinterDirectIoOpen VDMSERVICECLASS = 8  //col:31
    VdmPrinterDirectIoClose VDMSERVICECLASS = 9  //col:32
    VdmPrinterInitialize VDMSERVICECLASS = 10  //col:33
    VdmSetLdtEntries VDMSERVICECLASS = 11  //col:34
    VdmSetProcessLdtInfo VDMSERVICECLASS = 12  //col:35
    VdmAdlibEmulation VDMSERVICECLASS = 13  //col:36
    VdmPMCliControl VDMSERVICECLASS = 14  //col:37
    VdmQueryVdmProcess VDMSERVICECLASS = 15  //col:38
    VdmPreInitialize VDMSERVICECLASS = 16  //col:39
)


type     TraceControlStartLogger = 1 uint32
const(
    TraceControlStartLogger  TRACE_CONTROL_INFORMATION_CLASS =  1  //col:64
    TraceControlStopLogger  TRACE_CONTROL_INFORMATION_CLASS =  2  //col:65
    TraceControlQueryLogger  TRACE_CONTROL_INFORMATION_CLASS =  3  //col:66
    TraceControlUpdateLogger  TRACE_CONTROL_INFORMATION_CLASS =  4  //col:67
    TraceControlFlushLogger  TRACE_CONTROL_INFORMATION_CLASS =  5  //col:68
    TraceControlIncrementLoggerFile  TRACE_CONTROL_INFORMATION_CLASS =  6  //col:69
    TraceControlRealtimeConnect  TRACE_CONTROL_INFORMATION_CLASS =  11  //col:71
    TraceControlActivityIdCreate  TRACE_CONTROL_INFORMATION_CLASS =  12  //col:72
    TraceControlWdiDispatchControl  TRACE_CONTROL_INFORMATION_CLASS =  13  //col:73
    TraceControlRealtimeDisconnectConsumerByHandle  TRACE_CONTROL_INFORMATION_CLASS =  14  //col:74
    TraceControlRegisterGuidsCode  TRACE_CONTROL_INFORMATION_CLASS =  15  //col:75
    TraceControlReceiveNotification  TRACE_CONTROL_INFORMATION_CLASS =  16  //col:76
    TraceControlSendDataBlock  TRACE_CONTROL_INFORMATION_CLASS =  17 // EnableGuid  //col:77
    TraceControlSendReplyDataBlock  TRACE_CONTROL_INFORMATION_CLASS =  18  //col:78
    TraceControlReceiveReplyDataBlock  TRACE_CONTROL_INFORMATION_CLASS =  19  //col:79
    TraceControlWdiUpdateSem  TRACE_CONTROL_INFORMATION_CLASS =  20  //col:80
    TraceControlEnumTraceGuidList  TRACE_CONTROL_INFORMATION_CLASS =  21  //col:81
    TraceControlGetTraceGuidInfo  TRACE_CONTROL_INFORMATION_CLASS =  22  //col:82
    TraceControlEnumerateTraceGuids  TRACE_CONTROL_INFORMATION_CLASS =  23  //col:83
    TraceControlRegisterSecurityProv  TRACE_CONTROL_INFORMATION_CLASS =  24  //col:84
    TraceControlQueryReferenceTime  TRACE_CONTROL_INFORMATION_CLASS =  25  //col:85
    TraceControlTrackProviderBinary  TRACE_CONTROL_INFORMATION_CLASS =  26  //col:86
    TraceControlAddNotificationEvent  TRACE_CONTROL_INFORMATION_CLASS =  27  //col:87
    TraceControlUpdateDisallowList  TRACE_CONTROL_INFORMATION_CLASS =  28  //col:88
    TraceControlSetEnableAllKeywordsCode  TRACE_CONTROL_INFORMATION_CLASS =  29  //col:89
    TraceControlSetProviderTraitsCode  TRACE_CONTROL_INFORMATION_CLASS =  30  //col:90
    TraceControlUseDescriptorTypeCode  TRACE_CONTROL_INFORMATION_CLASS =  31  //col:91
    TraceControlEnumTraceGroupList  TRACE_CONTROL_INFORMATION_CLASS =  32  //col:92
    TraceControlGetTraceGroupInfo  TRACE_CONTROL_INFORMATION_CLASS =  33  //col:93
    TraceControlTraceSetDisallowList TRACE_CONTROL_INFORMATION_CLASS =  34  //col:94
    TraceControlSetCompressionSettings  TRACE_CONTROL_INFORMATION_CLASS =  35  //col:95
    TraceControlGetCompressionSettings TRACE_CONTROL_INFORMATION_CLASS =  36  //col:96
    TraceControlUpdatePeriodicCaptureState  TRACE_CONTROL_INFORMATION_CLASS =  37  //col:97
    TraceControlGetPrivateSessionTraceHandle  TRACE_CONTROL_INFORMATION_CLASS =  38  //col:98
    TraceControlRegisterPrivateSession  TRACE_CONTROL_INFORMATION_CLASS =  39  //col:99
    TraceControlQuerySessionDemuxObject  TRACE_CONTROL_INFORMATION_CLASS =  40  //col:100
    TraceControlSetProviderBinaryTracking  TRACE_CONTROL_INFORMATION_CLASS =  41  //col:101
    TraceControlMaxLoggers  TRACE_CONTROL_INFORMATION_CLASS =  42  //col:102
    TraceControlMaxPmcCounter  TRACE_CONTROL_INFORMATION_CLASS =  43  //col:103
)



type (
Ntmisc interface{
 * Attribution 4.0 International ()(ok bool)//col:40
NtVdmControl()(ok bool)//col:104
}
)

func NewNtmisc() { return & ntmisc{} }

func (n *ntmisc) * Attribution 4.0 International ()(ok bool){//col:40
/* * Attribution 4.0 International (CC BY 4.0) license. 
 * 
 * You must give appropriate credit, provide a link to the license, and 
 * indicate if changes were made. You may do so in any reasonable manner, but 
 * not in any way that suggests the licensor endorses you or your use.
#ifndef _NTMISC_H
#define _NTMISC_H
#define FLT_PORT_CONNECT 0x0001
#define FLT_PORT_ALL_ACCESS (FLT_PORT_CONNECT | STANDARD_RIGHTS_ALL)
typedef enum _VDMSERVICECLASS
{
    VdmStartExecution,
    VdmQueueInterrupt,
    VdmDelayInterrupt,
    VdmInitialize,
    VdmFeatures,
    VdmSetInt21Handler,
    VdmQueryDir,
    VdmPrinterDirectIoOpen,
    VdmPrinterDirectIoClose,
    VdmPrinterInitialize,
    VdmSetLdtEntries,
    VdmSetProcessLdtInfo,
    VdmAdlibEmulation,
    VdmPMCliControl,
    VdmQueryVdmProcess,
    VdmPreInitialize
} VDMSERVICECLASS, *PVDMSERVICECLASS;*/
return true
}

func (n *ntmisc)NtVdmControl()(ok bool){//col:104
/*NtVdmControl(
    _In_ VDMSERVICECLASS Service,
    _Inout_ PVOID ServiceData
    );
NTSYSCALLAPI
NTSTATUS
NTAPI
NtTraceEvent(
    _In_ HANDLE TraceHandle,
    _In_ ULONG Flags,
    _In_ ULONG FieldSize,
    _In_ PVOID Fields
    );
typedef enum _TRACE_CONTROL_INFORMATION_CLASS
{
    TraceControlStartLogger = 1,
    TraceControlStopLogger = 2,
    TraceControlQueryLogger = 3,
    TraceControlUpdateLogger = 4,
    TraceControlFlushLogger = 5,
    TraceControlIncrementLoggerFile = 6,
    TraceControlRealtimeConnect = 11,
    TraceControlActivityIdCreate = 12,
    TraceControlWdiDispatchControl = 13,
    TraceControlRealtimeDisconnectConsumerByHandle = 14,
    TraceControlRegisterGuidsCode = 15,
    TraceControlReceiveNotification = 16,
    TraceControlSendReplyDataBlock = 18,
    TraceControlReceiveReplyDataBlock = 19,
    TraceControlWdiUpdateSem = 20,
    TraceControlEnumTraceGuidList = 21,
    TraceControlGetTraceGuidInfo = 22,
    TraceControlEnumerateTraceGuids = 23,
    TraceControlRegisterSecurityProv = 24,
    TraceControlQueryReferenceTime = 25,
    TraceControlTrackProviderBinary = 26,
    TraceControlAddNotificationEvent = 27,
    TraceControlUpdateDisallowList = 28,
    TraceControlSetEnableAllKeywordsCode = 29,
    TraceControlSetProviderTraitsCode = 30,
    TraceControlUseDescriptorTypeCode = 31,
    TraceControlEnumTraceGroupList = 32,
    TraceControlGetTraceGroupInfo = 33,
    TraceControlTraceSetDisallowList= 34,
    TraceControlSetCompressionSettings = 35,
    TraceControlGetCompressionSettings= 36,
    TraceControlUpdatePeriodicCaptureState = 37,
    TraceControlGetPrivateSessionTraceHandle = 38,
    TraceControlRegisterPrivateSession = 39,
    TraceControlQuerySessionDemuxObject = 40,
    TraceControlSetProviderBinaryTracking = 41,
    TraceControlMaxLoggers = 42,
    TraceControlMaxPmcCounter = 43
} TRACE_CONTROL_INFORMATION_CLASS;*/
return true
}



