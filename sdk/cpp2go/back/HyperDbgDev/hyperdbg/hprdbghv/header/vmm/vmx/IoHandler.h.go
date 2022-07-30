package vmx
//back\HyperDbgDev\hyperdbg\hprdbghv\header\vmm\vmx\IoHandler.h.back

type typedef enum uint32
const(
typedef enum  = 1  //col:26
    AccessOut   =  0  //col:28
    AccessIn    =  1  //col:29
)


type typedef enum uint32
const(
typedef enum  = 1  //col:36
    OpEncodingDx    =  0  //col:38
    OpEncodingImm   =  1  //col:39
)



type (
IoHandler interface{
__inbyte()(ok bool)//col:54
__inword()(ok bool)//col:64
__indword()(ok bool)//col:74
__inbytestring()(ok bool)//col:84
__inwordstring()(ok bool)//col:94
__indwordstring()(ok bool)//col:104
__outbyte()(ok bool)//col:114
__outword()(ok bool)//col:124
__outdword()(ok bool)//col:134
__outbytestring()(ok bool)//col:144
__outwordstring()(ok bool)//col:154
__outdwordstring()(ok bool)//col:164
}
)

func NewIoHandler() { return & ioHandler{} }

func (i *ioHandler)__inbyte()(ok bool){//col:54
/*__inbyte(unsigned short);
#pragma intrinsic(__inbyte)
inline UINT8
IoInByte(UINT16 port)
{
    return __inbyte(port);
}*/
return true
}

func (i *ioHandler)__inword()(ok bool){//col:64
/*__inword(unsigned short);
#pragma intrinsic(__inword)
inline UINT16
IoInWord(UINT16 port)
{
    return __inword(port);
}*/
return true
}

func (i *ioHandler)__indword()(ok bool){//col:74
/*__indword(unsigned short);
#pragma intrinsic(__indword)
inline UINT32
IoInDword(UINT16 port)
{
    return __indword(port);
}*/
return true
}

func (i *ioHandler)__inbytestring()(ok bool){//col:84
/*__inbytestring(unsigned short, unsigned char *, unsigned long);
#pragma intrinsic(__inbytestring)
inline void
IoInByteString(UINT16 port, UINT8 * data, UINT32 size)
{
    __inbytestring(port, data, size);
}*/
return true
}

func (i *ioHandler)__inwordstring()(ok bool){//col:94
/*__inwordstring(unsigned short, unsigned short *, unsigned long);
#pragma intrinsic(__inwordstring)
inline void
IoInWordString(UINT16 port, UINT16 * data, UINT32 size)
{
    __inwordstring(port, data, size);
}*/
return true
}

func (i *ioHandler)__indwordstring()(ok bool){//col:104
/*__indwordstring(unsigned short, unsigned long *, unsigned long);
#pragma intrinsic(__indwordstring)
inline void
IoInDwordString(UINT16 port, UINT32 * data, UINT32 size)
{
    __indwordstring(port, (unsigned long *)data, size);
}*/
return true
}

func (i *ioHandler)__outbyte()(ok bool){//col:114
/*__outbyte(unsigned short, unsigned char);
#pragma intrinsic(__outbyte)
inline void
IoOutByte(UINT16 port, UINT8 value)
{
    __outbyte(port, value);
}*/
return true
}

func (i *ioHandler)__outword()(ok bool){//col:124
/*__outword(unsigned short, unsigned short);
#pragma intrinsic(__outword)
inline void
IoOutWord(UINT16 port, UINT16 value)
{
    __outword(port, value);
}*/
return true
}

func (i *ioHandler)__outdword()(ok bool){//col:134
/*__outdword(unsigned short, unsigned long);
#pragma intrinsic(__outdword)
inline void
IoOutDword(UINT16 port, UINT32 value)
{
    __outdword(port, value);
}*/
return true
}

func (i *ioHandler)__outbytestring()(ok bool){//col:144
/*__outbytestring(unsigned short, unsigned char *, unsigned long);
#pragma intrinsic(__outbytestring)
inline void
IoOutByteString(UINT16 port, UINT8 * data, UINT32 count)
{
    __outbytestring(port, data, count);
}*/
return true
}

func (i *ioHandler)__outwordstring()(ok bool){//col:154
/*__outwordstring(unsigned short, unsigned short *, unsigned long);
#pragma intrinsic(__outwordstring)
inline void
IoOutWordString(UINT16 port, UINT16 * data, UINT32 count)
{
    __outwordstring(port, data, count);
}*/
return true
}

func (i *ioHandler)__outdwordstring()(ok bool){//col:164
/*__outdwordstring(unsigned short, unsigned long *, unsigned long);
#pragma intrinsic(__outdwordstring)
inline void
IoOutDwordString(UINT16 port, UINT32 * data, UINT32 count)
{
    __outdwordstring(port, (unsigned long *)data, count);
}*/
return true
}



