# 1 "D:\\workspace\\workspace\\branch\\gui\\ux\\sdk\\bug\\1.c"
# 1 "<built-in>" 1
# 1 "<built-in>" 3
# 374 "<built-in>" 3
# 1 "<command line>" 1
# 1 "<built-in>" 2
# 1 "D:\\workspace\\workspace\\branch\\gui\\ux\\sdk\\bug\\1.c" 2
typedef unsigned char UCHAR;
typedef UCHAR BOOLEAN;
typedef BOOLEAN * PBOOLEAN;
typedef BOOLEAN (*CHECK_VMX_OPERATION)();


typedef unsigned long long QWORD;
typedef unsigned __int64 UINT64, *PUINT64;
typedef unsigned long DWORD;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int * PUINT;
typedef unsigned __int64 ULONG64, *PULONG64;
typedef unsigned __int64 DWORD64, *PDWORD64;
typedef char CHAR;



typedef unsigned char UCHAR;
typedef unsigned short USHORT;
typedef unsigned long ULONG;

typedef UCHAR BOOLEAN;
typedef BOOLEAN * PBOOLEAN;

typedef signed char INT8, *PINT8;
typedef signed short INT16, *PINT16;
typedef signed int INT32, *PINT32;
typedef signed __int64 INT64, *PINT64;
typedef unsigned char UINT8, *PUINT8;
typedef unsigned short UINT16, *PUINT16;
typedef unsigned int UINT32, *PUINT32;
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
