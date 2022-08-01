package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\IoHandler.h.back

const(
typedef enum = 1  //col:1
    AccessOut  =  0  //col:3
    AccessIn   =  1  //col:4
)


const(
typedef enum = 1  //col:6
    OpEncodingDx   =  0  //col:8
    OpEncodingImm  =  1  //col:9
)



type (
IoHandler interface{
__inbyte()(ok bool)//col:7
__inword()(ok bool)//col:14
__indword()(ok bool)//col:21
__inbytestring()(ok bool)//col:28
__inwordstring()(ok bool)//col:35
__indwordstring()(ok bool)//col:42
__outbyte()(ok bool)//col:49
__outword()(ok bool)//col:56
__outdword()(ok bool)//col:63
__outbytestring()(ok bool)//col:70
__outwordstring()(ok bool)//col:77
__outdwordstring()(ok bool)//col:84
}
)

func NewIoHandler() { return & ioHandler{} }

func (i *ioHandler)__inbyte()(ok bool){//col:7
/*__inbyte(unsigned short);
#pragma intrinsic(__inbyte)
inline UINT8
IoInByte(UINT16 port)
{
    return __inbyte(port);
}*/
return true
}

func (i *ioHandler)__inword()(ok bool){//col:14
/*__inword(unsigned short);
#pragma intrinsic(__inword)
inline UINT16
IoInWord(UINT16 port)
{
    return __inword(port);
}*/
return true
}

func (i *ioHandler)__indword()(ok bool){//col:21
/*__indword(unsigned short);
#pragma intrinsic(__indword)
inline UINT32
IoInDword(UINT16 port)
{
    return __indword(port);
}*/
return true
}

func (i *ioHandler)__inbytestring()(ok bool){//col:28
/*__inbytestring(unsigned short, unsigned char *, unsigned long);
#pragma intrinsic(__inbytestring)
inline void
IoInByteString(UINT16 port, UINT8 * data, UINT32 size)
{
    __inbytestring(port, data, size);
}*/
return true
}

func (i *ioHandler)__inwordstring()(ok bool){//col:35
/*__inwordstring(unsigned short, unsigned short *, unsigned long);
#pragma intrinsic(__inwordstring)
inline void
IoInWordString(UINT16 port, UINT16 * data, UINT32 size)
{
    __inwordstring(port, data, size);
}*/
return true
}

func (i *ioHandler)__indwordstring()(ok bool){//col:42
/*__indwordstring(unsigned short, unsigned long *, unsigned long);
#pragma intrinsic(__indwordstring)
inline void
IoInDwordString(UINT16 port, UINT32 * data, UINT32 size)
{
    __indwordstring(port, (unsigned long *)data, size);
}*/
return true
}

func (i *ioHandler)__outbyte()(ok bool){//col:49
/*__outbyte(unsigned short, unsigned char);
#pragma intrinsic(__outbyte)
inline void
IoOutByte(UINT16 port, UINT8 value)
{
    __outbyte(port, value);
}*/
return true
}

func (i *ioHandler)__outword()(ok bool){//col:56
/*__outword(unsigned short, unsigned short);
#pragma intrinsic(__outword)
inline void
IoOutWord(UINT16 port, UINT16 value)
{
    __outword(port, value);
}*/
return true
}

func (i *ioHandler)__outdword()(ok bool){//col:63
/*__outdword(unsigned short, unsigned long);
#pragma intrinsic(__outdword)
inline void
IoOutDword(UINT16 port, UINT32 value)
{
    __outdword(port, value);
}*/
return true
}

func (i *ioHandler)__outbytestring()(ok bool){//col:70
/*__outbytestring(unsigned short, unsigned char *, unsigned long);
#pragma intrinsic(__outbytestring)
inline void
IoOutByteString(UINT16 port, UINT8 * data, UINT32 count)
{
    __outbytestring(port, data, count);
}*/
return true
}

func (i *ioHandler)__outwordstring()(ok bool){//col:77
/*__outwordstring(unsigned short, unsigned short *, unsigned long);
#pragma intrinsic(__outwordstring)
inline void
IoOutWordString(UINT16 port, UINT16 * data, UINT32 count)
{
    __outwordstring(port, data, count);
}*/
return true
}

func (i *ioHandler)__outdwordstring()(ok bool){//col:84
/*__outdwordstring(unsigned short, unsigned long *, unsigned long);
#pragma intrinsic(__outdwordstring)
inline void
IoOutDwordString(UINT16 port, UINT32 * data, UINT32 count)
{
    __outdwordstring(port, (unsigned long *)data, count);
}*/
return true
}



