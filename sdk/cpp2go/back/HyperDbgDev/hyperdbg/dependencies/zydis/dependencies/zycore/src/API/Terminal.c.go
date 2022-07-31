package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Terminal.c.back

type (
Terminal interface{
  Zyan Core Library ()(ok bool)//col:96
ZyanStatus ZyanTerminalIsTTY()(ok bool)//col:154
}
)

func NewTerminal() { return & terminal{} }

func (t *terminal)  Zyan Core Library ()(ok bool){//col:96
/*  Zyan Core Library (Zycore-C)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#include <Zycore/API/Terminal.h>
#if   defined(ZYAN_POSIX)
#   include <unistd.h>
#elif defined(ZYAN_WINDOWS)
#   include <windows.h>
#   include <io.h>
#else
#   error "Unsupported platform detected"
#endif
#ifdef ZYAN_WINDOWS
#   ifndef ENABLE_VIRTUAL_TERMINAL_PROCESSING
#       define ENABLE_VIRTUAL_TERMINAL_PROCESSING 0x0004
#   endif
#endif
ZyanStatus ZyanTerminalEnableVT100(ZyanStandardStream stream)
{
    if ((stream != ZYAN_STDSTREAM_OUT) && (stream != ZYAN_STDSTREAM_ERR))
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
#ifdef ZYAN_WINDOWS
    int file;
    switch (stream)
    {
    case ZYAN_STDSTREAM_OUT:
        file = _fileno(ZYAN_STDOUT);
        break;
    case ZYAN_STDSTREAM_ERR:
        file = _fileno(ZYAN_STDERR);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (file < 0)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    HANDLE const handle = (HANDLE)_get_osfhandle(file);
    if (handle == INVALID_HANDLE_VALUE)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    DWORD mode;
    if (!GetConsoleMode(handle, &mode))
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING;
    if (!SetConsoleMode(handle, mode))
    {
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
#endif
    return ZYAN_STATUS_SUCCESS;
}*/
return true
}

func (t *terminal)ZyanStatus ZyanTerminalIsTTY()(ok bool){//col:154
/*ZyanStatus ZyanTerminalIsTTY(ZyanStandardStream stream)
{
    int file;
#ifdef ZYAN_WINDOWS
    switch (stream)
    {
    case ZYAN_STDSTREAM_IN:
        file = _fileno(ZYAN_STDIN);
        break;
    case ZYAN_STDSTREAM_OUT:
        file = _fileno(ZYAN_STDOUT);
        break;
    case ZYAN_STDSTREAM_ERR:
        file = _fileno(ZYAN_STDERR);
        break;
    default:
        ZYAN_UNREACHABLE;
    }
    if (file < 0)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
#else
    switch (stream)
    {
    case ZYAN_STDSTREAM_IN:
        file = STDIN_FILENO;
        break;
    case ZYAN_STDSTREAM_OUT:
        file = STDOUT_FILENO;
        break;
    case ZYAN_STDSTREAM_ERR:
        file = STDERR_FILENO;
        break;
    default:
        ZYAN_UNREACHABLE;
    }
#endif
#ifdef ZYAN_WINDOWS
    if (_isatty(file))
#else
    if ( isatty(file))
#endif
    {
        return ZYAN_STATUS_TRUE;
    }
    if (ZYAN_ERRNO == EBADF)
    {
        return ZYAN_STATUS_INVALID_ARGUMENT;
    }
    return ZYAN_STATUS_FALSE;
}*/
return true
}



