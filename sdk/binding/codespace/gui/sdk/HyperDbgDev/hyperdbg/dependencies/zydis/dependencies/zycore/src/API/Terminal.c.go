package API
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Terminal.c.back

type (
Terminal interface{
ZyanStatus_ZyanTerminalEnableVT100()(ok bool)//col:41
ZyanStatus_ZyanTerminalIsTTY()(ok bool)//col:93
}
terminal struct{}
)

func NewTerminal()Terminal{ return & terminal{} }

func (t *terminal)ZyanStatus_ZyanTerminalEnableVT100()(ok bool){//col:41
/*ZyanStatus ZyanTerminalEnableVT100(ZyanStandardStream stream)
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

func (t *terminal)ZyanStatus_ZyanTerminalIsTTY()(ok bool){//col:93
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



