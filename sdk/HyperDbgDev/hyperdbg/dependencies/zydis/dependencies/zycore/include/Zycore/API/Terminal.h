#ifndef ZYCORE_TERMINAL_H
#define ZYCORE_TERMINAL_H
#include <ZycoreExportConfig.h>
#include <Zycore/LibC.h>
#include <Zycore/Status.h>
#ifdef __cplusplus
extern "C" {
#endif
#ifndef ZYAN_NO_LIBC
#    define ZYAN_VT100SGR_RESET             "\033[0m"
#    define ZYAN_VT100SGR_FG_DEFAULT        "\033[39m"
#    define ZYAN_VT100SGR_FG_BLACK          "\033[30m"
#    define ZYAN_VT100SGR_FG_RED            "\033[31m"
#    define ZYAN_VT100SGR_FG_GREEN          "\033[32m"
#    define ZYAN_VT100SGR_FG_YELLOW         "\033[33m"
#    define ZYAN_VT100SGR_FG_BLUE           "\033[34m"
#    define ZYAN_VT100SGR_FG_MAGENTA        "\033[35m"
#    define ZYAN_VT100SGR_FG_CYAN           "\033[36m"
#    define ZYAN_VT100SGR_FG_WHITE          "\033[37m"
#    define ZYAN_VT100SGR_FG_BRIGHT_BLACK   "\033[90m"
#    define ZYAN_VT100SGR_FG_BRIGHT_RED     "\033[91m"
#    define ZYAN_VT100SGR_FG_BRIGHT_GREEN   "\033[92m"
#    define ZYAN_VT100SGR_FG_BRIGHT_YELLOW  "\033[93m"
#    define ZYAN_VT100SGR_FG_BRIGHT_BLUE    "\033[94m"
#    define ZYAN_VT100SGR_FG_BRIGHT_MAGENTA "\033[95m"
#    define ZYAN_VT100SGR_FG_BRIGHT_CYAN    "\033[96m"
#    define ZYAN_VT100SGR_FG_BRIGHT_WHITE   "\033[97m"
#    define ZYAN_VT100SGR_BG_DEFAULT        "\033[49m"
#    define ZYAN_VT100SGR_BG_BLACK          "\033[40m"
#    define ZYAN_VT100SGR_BG_RED            "\033[41m"
#    define ZYAN_VT100SGR_BG_GREEN          "\033[42m"
#    define ZYAN_VT100SGR_BG_YELLOW         "\033[43m"
#    define ZYAN_VT100SGR_BG_BLUE           "\033[44m"
#    define ZYAN_VT100SGR_BG_MAGENTA        "\033[45m"
#    define ZYAN_VT100SGR_BG_CYAN           "\033[46m"
#    define ZYAN_VT100SGR_BG_WHITE          "\033[47m"
#    define ZYAN_VT100SGR_BG_BRIGHT_BLACK   "\033[100m"
#    define ZYAN_VT100SGR_BG_BRIGHT_RED     "\033[101m"
#    define ZYAN_VT100SGR_BG_BRIGHT_GREEN   "\033[102m"
#    define ZYAN_VT100SGR_BG_BRIGHT_YELLOW  "\033[103m"
#    define ZYAN_VT100SGR_BG_BRIGHT_BLUE    "\033[104m"
#    define ZYAN_VT100SGR_BG_BRIGHT_MAGENTA "\033[105m"
#    define ZYAN_VT100SGR_BG_BRIGHT_CYAN    "\033[106m"
#    define ZYAN_VT100SGR_BG_BRIGHT_WHITE   "\033[107m"
typedef enum ZyanStandardStream_ {
    ZYAN_STDSTREAM_IN,
    ZYAN_STDSTREAM_OUT,
    ZYAN_STDSTREAM_ERR
} ZyanStandardStream;
ZYCORE_EXPORT ZyanStatus
ZyanTerminalEnableVT100(ZyanStandardStream stream);
ZYCORE_EXPORT ZyanStatus
ZyanTerminalIsTTY(ZyanStandardStream stream);
#endif // ZYAN_NO_LIBC
#ifdef __cplusplus
}

#endif
#endif /* ZYCORE_TERMINAL_H */
