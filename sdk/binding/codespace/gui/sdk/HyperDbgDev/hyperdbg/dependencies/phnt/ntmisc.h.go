package phnt

const (
	VdmStartExecution       = 1  //col:3
	VdmQueueInterrupt       = 2  //col:4
	VdmDelayInterrupt       = 3  //col:5
	VdmInitialize           = 4  //col:6
	VdmFeatures             = 5  //col:7
	VdmSetInt21Handler      = 6  //col:8
	VdmQueryDir             = 7  //col:9
	VdmPrinterDirectIoOpen  = 8  //col:10
	VdmPrinterDirectIoClose = 9  //col:11
	VdmPrinterInitialize    = 10 //col:12
	VdmSetLdtEntries        = 11 //col:13
	VdmSetProcessLdtInfo    = 12 //col:14
	VdmAdlibEmulation       = 13 //col:15
	VdmPMCliControl         = 14 //col:16
	VdmQueryVdmProcess      = 15 //col:17
	VdmPreInitialize        = 16 //col:18
)

const (
	TraceControlStartLogger                        = 1  //col:22
	TraceControlStopLogger                         = 2  //col:23
	TraceControlQueryLogger                        = 3  //col:24
	TraceControlUpdateLogger                       = 4  //col:25
	TraceControlFlushLogger                        = 5  //col:26
	TraceControlIncrementLoggerFile                = 6  //col:27
	TraceControlRealtimeConnect                    = 11 //col:28
	TraceControlActivityIdCreate                   = 12 //col:29
	TraceControlWdiDispatchControl                 = 13 //col:30
	TraceControlRealtimeDisconnectConsumerByHandle = 14 //col:31
	TraceControlRegisterGuidsCode                  = 15 //col:32
	TraceControlReceiveNotification                = 16 //col:33
	TraceControlSendDataBlock                      = 17 //col:34
	TraceControlSendReplyDataBlock                 = 18 //col:35
	TraceControlReceiveReplyDataBlock              = 19 //col:36
	TraceControlWdiUpdateSem                       = 20 //col:37
	TraceControlEnumTraceGuidList                  = 21 //col:38
	TraceControlGetTraceGuidInfo                   = 22 //col:39
	TraceControlEnumerateTraceGuids                = 23 //col:40
	TraceControlRegisterSecurityProv               = 24 //col:41
	TraceControlQueryReferenceTime                 = 25 //col:42
	TraceControlTrackProviderBinary                = 26 //col:43
	TraceControlAddNotificationEvent               = 27 //col:44
	TraceControlUpdateDisallowList                 = 28 //col:45
	TraceControlSetEnableAllKeywordsCode           = 29 //col:46
	TraceControlSetProviderTraitsCode              = 30 //col:47
	TraceControlUseDescriptorTypeCode              = 31 //col:48
	TraceControlEnumTraceGroupList                 = 32 //col:49
	TraceControlGetTraceGroupInfo                  = 33 //col:50
	TraceControlTraceSetDisallowList               = 34 //col:51
	TraceControlSetCompressionSettings             = 35 //col:52
	TraceControlGetCompressionSettings             = 36 //col:53
	TraceControlUpdatePeriodicCaptureState         = 37 //col:54
	TraceControlGetPrivateSessionTraceHandle       = 38 //col:55
	TraceControlRegisterPrivateSession             = 39 //col:56
	TraceControlQuerySessionDemuxObject            = 40 //col:57
	TraceControlSetProviderBinaryTracking          = 41 //col:58
	TraceControlMaxLoggers                         = 42 //col:59
	TraceControlMaxPmcCounter                      = 43 //col:60
)
