#define __int64 long long
#define __iamcu__
#define __int32 int
#define NTDDI_WIN7 0x06010000
#define __forceinline __attribute__((always_inline))
#define _AMD64_
#define _M_AMD64
#define _M_X64
#define __unaligned
#define _MSC_FULL_VER 192930133 # https://dev.to/yumetodo/list-of-mscver-and-mscfullver-8nd
#define _CTYPE_DISABLE_MACROS

#define XSTR(x) STR(x)
#define STR(x) #x

#include<sal.h>

# https://github.com/nemequ/salieri/blob/master/salieri.h

#if defined(_In_)
#undef _In_
#define _In_  __attribute__((anno("_In_")))
#endif

#if defined(_In_opt_)
#undef _In_opt_
#define _In_opt_  __attribute__((anno("_In_opt_")))
#endif

#if defined(_In_z_)
#undef _In_z_
#define _In_z_  __attribute__((anno("_In_z_")))
#endif

#if defined(_In_opt_z_)
#undef _In_opt_z_
#define _In_opt_z_  __attribute__((anno("_In_opt_z_")))
#endif

#if defined(_In_reads_)
#undef _In_reads_
#define _In_reads_(s)  __attribute__((anno("_In_reads_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_opt_)
#undef _In_reads_opt_
#define _In_reads_opt_(s)  __attribute__((anno("_In_reads_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_bytes_)
#undef _In_reads_bytes_
#define _In_reads_bytes_(s)  __attribute__((anno("_In_reads_bytes_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_bytes_opt_)
#undef _In_reads_bytes_opt_
#define _In_reads_bytes_opt_(s)  __attribute__((anno("_In_reads_bytes_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_z_)
#undef _In_reads_z_
#define _In_reads_z_(s)  __attribute__((anno("_In_reads_z_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_opt_z_)
#undef _In_reads_opt_z_
#define _In_reads_opt_z_(s)  __attribute__((anno("_In_reads_opt_z_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_or_z_)
#undef _In_reads_or_z_
#define _In_reads_or_z_(s)  __attribute__((anno("_In_reads_or_z_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_or_z_opt_)
#undef _In_reads_or_z_opt_
#define _In_reads_or_z_opt_(s)  __attribute__((anno("_In_reads_or_z_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_to_ptr_)
#undef _In_reads_to_ptr_
#define _In_reads_to_ptr_(s)  __attribute__((anno("_In_reads_to_ptr_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_to_ptr_opt_)
#undef _In_reads_to_ptr_opt_
#define _In_reads_to_ptr_opt_(s)  __attribute__((anno("_In_reads_to_ptr_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_to_ptr_z_)
#undef _In_reads_to_ptr_z_
#define _In_reads_to_ptr_z_(s)  __attribute__((anno("_In_reads_to_ptr_z_"))) __attribute__((size(#s)))
#endif

#if defined(_In_reads_to_ptr_opt_z_)
#undef _In_reads_to_ptr_opt_z_
#define _In_reads_to_ptr_opt_z_(s)  __attribute__((anno("_In_reads_to_ptr_opt_z_"))) __attribute__((size(#s)))
#endif


//////////////////////// OUT///////////////////////////////////

#if defined(_Out_)
#undef _Out_
#define _Out_  __attribute__((anno("_Out_")))
#endif

#if defined(_Out_opt_)
#undef _Out_opt_
#define _Out_opt_  __attribute__((anno("_Out_opt_")))
#endif

#if defined(_Out_writes_)
#undef _Out_writes_
#define _Out_writes_(s)  __attribute__((anno("_Out_writes_")))  __attribute__((size(#s)))
#endif

#if defined(_Out_writes_opt_)
#undef _Out_writes_opt_
#define _Out_writes_opt_(s)  __attribute__((anno("_Out_writes_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_bytes_)
#undef _Out_writes_bytes_
#define _Out_writes_bytes_(s)  __attribute__((anno("_Out_writes_bytes_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_bytes_opt_)
#undef _Out_writes_bytes_opt_
#define _Out_writes_bytes_opt_(s)  __attribute__((anno("_Out_writes_bytes_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_z_)
#undef _Out_writes_z_
#define _Out_writes_z_(s)  __attribute__((anno("_Out_writes_z_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_opt_z_)
#undef _Out_writes_opt_z_
#define _Out_writes_opt_z_(s)  __attribute__((anno("_Out_writes_opt_z_", s)))
#endif

#if defined(_Out_writes_to_)
#undef _Out_writes_to_
#define _Out_writes_to_(s, c)  __attribute__((anno("_Out_writes_to_"))) __attribute__((size(#s))) __attribute__((count(#c)))
#endif

#if defined(_Out_writes_to_opt_)
#undef _Out_writes_to_opt_
#define _Out_writes_to_opt_(s, c)  __attribute__((anno("_Out_writes_to_opt_"))) __attribute__((size(#s))) __attribute__((count(#c)))
#endif

#if defined(_Out_writes_all_)
#undef _Out_writes_all_
#define _Out_writes_all_(s)  __attribute__((anno("_Out_writes_all_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_all_opt_)
#undef _Out_writes_all_opt_
#define _Out_writes_all_opt_(s)  __attribute__((anno("_Out_writes_all_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_bytes_to_)
#undef _Out_writes_bytes_to_
#define _Out_writes_bytes_to_(s, c)  __attribute__((anno("_Out_writes_bytes_to_"))) __attribute__((size(#s))) __attribute__((count(#c)))
#endif

#if defined(_Out_writes_bytes_to_opt_)
#undef _Out_writes_bytes_to_opt_
#define _Out_writes_bytes_to_opt_(s, c)  __attribute__((anno("_Out_writes_bytes_to_opt_"))) __attribute__((size(#s))) __attribute__((count(#c)))
#endif

#if defined(_Out_writes_bytes_all_)
#undef _Out_writes_bytes_all_
#define _Out_writes_bytes_all_(s)  __attribute__((anno("_Out_writes_bytes_all_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_bytes_all_opt_)
#undef _Out_writes_bytes_all_opt_
#define _Out_writes_bytes_all_opt_(s)  __attribute__((anno("_Out_writes_bytes_all_opt_"))) __attribute__((size(#s)))
#endif


#if defined(_Out_writes_to_ptr_)
#undef _Out_writes_to_ptr_
#define _Out_writes_to_ptr_(s)  __attribute__((anno("_Out_writes_to_ptr_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_to_ptr_opt_)
#undef _Out_writes_to_ptr_opt_
#define _Out_writes_to_ptr_opt_(s)  __attribute__((anno("_Out_writes_to_ptr_opt_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_to_ptr_z_)
#undef _Out_writes_to_ptr_z_
#define _Out_writes_to_ptr_z_(s)  __attribute__((anno("_Out_writes_to_ptr_z_"))) __attribute__((size(#s)))
#endif

#if defined(_Out_writes_to_ptr_opt_z_)
#undef _Out_writes_to_ptr_opt_z_
#define _Out_writes_to_ptr_opt_z_(s)  __attribute__((anno("_Out_writes_to_ptr_opt_z_"))) __attribute__((size(#s)))
#endif



#if defined(_Outptr_)
#undef _Outptr_
#define _Outptr_  __attribute__((anno("_Outptr_")))
#endif


////////////////////////IN OUT /////////////////////////////////


#if defined(_Inout_)
#undef _Inout_
#define _Inout_  __attribute__((anno("_Inout_")))
#endif

#include <windows.h>
#include <tlhelp32.h>
#include <wininet.h>
#include <shlwapi.h>
#include <shellapi.h>
#include <Shlobj.h>
#include <string.h>
#include <wchar.h>
#include <winsock2.h>
#include <wincrypt.h>
#include <ws2tcpip.h>
#include <phnt_windows.h> // for phnt
#include <phnt.h>
//#pragma message "The value of Windows:" XSTR(_WIN32_WINNT)

// Testing
//#include <minwindef.h>
//#include <minwinbase.h>
//int myfunc (_In_ int param1, _Out_opt_ double param2, _Out_writes_(whatever) PUCHAR param3);

// typedef void VOID;
// void myfunc (_In_ int* param1, _In_ int param2);


