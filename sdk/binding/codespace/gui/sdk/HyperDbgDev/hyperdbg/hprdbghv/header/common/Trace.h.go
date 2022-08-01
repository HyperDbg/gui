package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\common\Trace.h.back

const (
	WPP_CONTROL_GUIDS = WPP_DEFINE_CONTROL_GUID(HyperdbgLogger, (2AE39766, AE4B, 46AB, AFC4, 002DB8109721
), WPP_DEFINE_BIT(HVFS_LOG) WPP_DEFINE_BIT(HVFS_LOG_INFO) WPP_DEFINE_BIT(HVFS_LOG_WARNING) WPP_DEFINE_BIT(HVFS_LOG_ERROR) ) //col:1
TRACE_LEVEL_NONE =        0 //col:10
TRACE_LEVEL_FATAL = 1       //col:11
TRACE_LEVEL_ERROR = 2       //col:12
TRACE_LEVEL_WARNING = 3     //col:13
TRACE_LEVEL_INFORMATION = 4 //col:14
TRACE_LEVEL_VERBOSE =     5 //col:15
TRACE_LEVEL_RESERVED6 = 6   //col:16
TRACE_LEVEL_RESERVED7 = 7   //col:17
TRACE_LEVEL_RESERVED8 = 8 //col:18
TRACE_LEVEL_RESERVED9 = 9 //col:19
WPP_LEVEL_FLAGS_LOGGER(level, = flags)  WPP_LEVEL_LOGGER(flags) //col:20
WPP_LEVEL_FLAGS_ENABLED(level, = flags) (WPP_LEVEL_ENABLED(flags) && WPP_CONTROL(WPP_BIT_##flags).Level >= level) //col:21
WPP_FLAG_EXP_PRE(FLAGS, HR) = { if (HR != STATUS_SUCCESS) { //col:22
WPP_FLAG_EXP_POST(FLAGS, HR) =; } } //col:26
WPP_FLAG_EXP_ENABLED(FLAGS, = HR) WPP_FLAG_ENABLED(FLAGS) //col:30
WPP_FLAG_EXP_LOGGER(FLAGS, = HR)  WPP_FLAG_LOGGER(FLAGS) //col:31
)
