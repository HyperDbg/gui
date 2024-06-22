typedef unsigned char  UCHAR;
typedef UCHAR     BOOLEAN;  // winnt
typedef BOOLEAN * PBOOLEAN; // winnt
typedef BOOLEAN (*CHECK_VMX_OPERATION)();


typedef unsigned long long QWORD;
typedef unsigned __int64   UINT64, *PUINT64;
typedef unsigned long      DWORD;
typedef int                BOOL;
typedef unsigned char      BYTE;
typedef unsigned short     WORD;
typedef int                INT;
typedef unsigned int       UINT;
typedef unsigned int *     PUINT;
typedef unsigned __int64   ULONG64, *PULONG64;
typedef unsigned __int64   DWORD64, *PDWORD64;
typedef char               CHAR;
//typedef wchar_t            WCHAR;
#define VOID void

typedef unsigned char  UCHAR;
typedef unsigned short USHORT;
typedef unsigned long  ULONG;

typedef UCHAR     BOOLEAN;  // winnt
typedef BOOLEAN * PBOOLEAN; // winnt

typedef signed char      INT8, *PINT8;
typedef signed short     INT16, *PINT16;
typedef signed int       INT32, *PINT32;
typedef signed __int64   INT64, *PINT64;
typedef unsigned char    UINT8, *PUINT8;
typedef unsigned short   UINT16, *PUINT16;
typedef unsigned int     UINT32, *PUINT32;
typedef unsigned __int64 UINT64, *PUINT64;


typedef struct _CR3_TYPE
{
    union
    {
        UINT64 Flags;

        struct
        {
            UINT64 Pcid : 12;
            UINT64 PageFrameNumber : 36;
            UINT64 Reserved1 : 12;
            UINT64 Reserved_2 : 3;
            UINT64 PcidInvalidate : 1;
        } Fields;
    };
} CR3_TYPE, *PCR3_TYPE;




//typedef unsigned char (*CHECK_VMX_OPERATION)();

/*
      "name": "CHECK_VMX_OPERATION",
      "type": {
        "qualType": "BOOLEAN (*)()"  //todo set it func type
      },


package main

import "fmt"

type UCHAR byte
type BOOLEAN UCHAR
type PBOOLEAN *BOOLEAN
type CHECK_VMX_OPERATION func() BOOLEAN

// 示例函数
func checkVMX() BOOLEAN {
    // 示例返回值
    return BOOLEAN(1)
}

func main() {
    var b BOOLEAN = 1
    var pb PBOOLEAN = &b
    var checkVMXOperation CHECK_VMX_OPERATION = checkVMX

    fmt.Println("BOOLEAN值:", b)
    fmt.Println("PBOOLEAN值:", *pb)
    fmt.Println("CHECK_VMX_OPERATION返回值:", checkVMXOperation())
}
*/