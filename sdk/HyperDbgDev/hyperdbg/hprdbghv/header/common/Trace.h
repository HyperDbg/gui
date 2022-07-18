#pragma once
#define WPP_CONTROL_GUIDS                                          \
    WPP_DEFINE_CONTROL_GUID(                                       \
        HyperdbgLogger,                                            \
        (2AE39766, AE4B, 46AB, AFC4, 002DB8109721),                \
        WPP_DEFINE_BIT(HVFS_LOG)         /* bit  0 = 0x00000001 */ \
        WPP_DEFINE_BIT(HVFS_LOG_INFO)    /* bit  1 = 0x00000002 */ \
        WPP_DEFINE_BIT(HVFS_LOG_WARNING) /* bit  2 = 0x00000004 */ \
        WPP_DEFINE_BIT(HVFS_LOG_ERROR)   /* bit  3 = 0x00000008 */ \
    )
#define TRACE_LEVEL_NONE                      0 // Tracing is not on
#define TRACE_LEVEL_FATAL                     1 // Abnormal exit or termination
#define TRACE_LEVEL_ERROR                     2 // Severe errors that need logging
#define TRACE_LEVEL_WARNING                   3 // Warnings such as allocation failure
#define TRACE_LEVEL_INFORMATION               4 // Includes non-error cases(for example, Entry-Exit)
#define TRACE_LEVEL_VERBOSE                   5 // Detailed traces from intermediate steps
#define TRACE_LEVEL_RESERVED6                 6
#define TRACE_LEVEL_RESERVED7                 7
#define TRACE_LEVEL_RESERVED8                 8
#define TRACE_LEVEL_RESERVED9                 9
#define WPP_LEVEL_FLAGS_LOGGER(level, flags)  WPP_LEVEL_LOGGER(flags)
#define WPP_LEVEL_FLAGS_ENABLED(level, flags) (WPP_LEVEL_ENABLED(flags) && WPP_CONTROL(WPP_BIT_##flags).Level >= level)
#define WPP_FLAG_EXP_PRE(FLAGS, HR) \
    {                               \
        if (HR != STATUS_SUCCESS) {
#define WPP_FLAG_EXP_POST(FLAGS, HR) \
    ;                                \
    }                                \
    }
#define WPP_FLAG_EXP_ENABLED(FLAGS, HR) WPP_FLAG_ENABLED(FLAGS)
#define WPP_FLAG_EXP_LOGGER(FLAGS, HR)  WPP_FLAG_LOGGER(FLAGS)
