package phnt


const(
_NTMISC_H =  //col:13
FLT_PORT_CONNECT = 0x0001 //col:17
FLT_PORT_ALL_ACCESS = (FLT_PORT_CONNECT | STANDARD_RIGHTS_ALL) //col:18
)

const(
    VdmStartExecution = 1  //col:24
    VdmQueueInterrupt = 2  //col:25
    VdmDelayInterrupt = 3  //col:26
    VdmInitialize = 4  //col:27
    VdmFeatures = 5  //col:28
    VdmSetInt21Handler = 6  //col:29
    VdmQueryDir = 7  //col:30
    VdmPrinterDirectIoOpen = 8  //col:31
    VdmPrinterDirectIoClose = 9  //col:32
    VdmPrinterInitialize = 10  //col:33
    VdmSetLdtEntries = 11  //col:34
    VdmSetProcessLdtInfo = 12  //col:35
    VdmAdlibEmulation = 13  //col:36
    VdmPMCliControl = 14  //col:37
    VdmQueryVdmProcess = 15  //col:38
    VdmPreInitialize = 16  //col:39
)


const(
    TraceControlStartLogger  =  1  //col:64
    TraceControlStopLogger  =  2  //col:65
    TraceControlQueryLogger  =  3  //col:66
    TraceControlUpdateLogger  =  4  //col:67
    TraceControlFlushLogger  =  5  //col:68
    TraceControlIncrementLoggerFile  =  6  //col:69
    TraceControlRealtimeConnect  =  11  //col:71
    TraceControlActivityIdCreate  =  12  //col:72
    TraceControlWdiDispatchControl  =  13  //col:73
    TraceControlRealtimeDisconnectConsumerByHandle  =  14  //col:74
    TraceControlRegisterGuidsCode  =  15  //col:75
    TraceControlReceiveNotification  =  16  //col:76
    TraceControlSendDataBlock  =  17 // EnableGuid  //col:77
    TraceControlSendReplyDataBlock  =  18  //col:78
    TraceControlReceiveReplyDataBlock  =  19  //col:79
    TraceControlWdiUpdateSem  =  20  //col:80
    TraceControlEnumTraceGuidList  =  21  //col:81
    TraceControlGetTraceGuidInfo  =  22  //col:82
    TraceControlEnumerateTraceGuids  =  23  //col:83
    TraceControlRegisterSecurityProv  =  24  //col:84
    TraceControlQueryReferenceTime  =  25  //col:85
    TraceControlTrackProviderBinary  =  26  //col:86
    TraceControlAddNotificationEvent  =  27  //col:87
    TraceControlUpdateDisallowList  =  28  //col:88
    TraceControlSetEnableAllKeywordsCode  =  29  //col:89
    TraceControlSetProviderTraitsCode  =  30  //col:90
    TraceControlUseDescriptorTypeCode  =  31  //col:91
    TraceControlEnumTraceGroupList  =  32  //col:92
    TraceControlGetTraceGroupInfo  =  33  //col:93
    TraceControlTraceSetDisallowList =  34  //col:94
    TraceControlSetCompressionSettings  =  35  //col:95
    TraceControlGetCompressionSettings =  36  //col:96
    TraceControlUpdatePeriodicCaptureState  =  37  //col:97
    TraceControlGetPrivateSessionTraceHandle  =  38  //col:98
    TraceControlRegisterPrivateSession  =  39  //col:99
    TraceControlQuerySessionDemuxObject  =  40  //col:100
    TraceControlSetProviderBinaryTracking  =  41  //col:101
    TraceControlMaxLoggers  =  42  //col:102
    TraceControlMaxPmcCounter  =  43  //col:103
)



type (
Ntmisc interface{
NtVdmControl()(ok bool)//col:104
}









































)

func NewNtmisc() { return & ntmisc{} }

func (n *ntmisc)NtVdmControl()(ok bool){//col:104






















































return true
}












































