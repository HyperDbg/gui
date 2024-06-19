##include <windows.h>

	#define X86_FLAGS_RESERVED_BITS 0xffc38028
	#define X86_FLAGS_FIXED         0x00000002

	#define IOCTL_PREACTIVATE_FUNCTIONALITY \
	CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)

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