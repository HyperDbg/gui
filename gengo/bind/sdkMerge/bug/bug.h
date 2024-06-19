//#include <windows.h>

	#define X86_FLAGS_RESERVED_BITS 0xffc38028
	#define X86_FLAGS_FIXED         0x00000002

	#define IOCTL_PREACTIVATE_FUNCTIONALITY \
	CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)

//	typedef struct _CR3_TYPE
//	{
//	    union
//	    {
//	        UINT64 Flags;
//
//	        struct
//	        {
//	            UINT64 Pcid : 12;
//	            UINT64 PageFrameNumber : 36;
//	            UINT64 Reserved1 : 12;
//	            UINT64 Reserved_2 : 3;
//	            UINT64 PcidInvalidate : 1;
//	        } Fields;
//	    };
//	} CR3_TYPE, *PCR3_TYPE;

//typedef union
//{
//    struct
//    {
//        /**
//         * [Bits 3:0] Segment type.
//         */
//        UINT32 Type : 4;
//
//        /**
//         * [Bit 4] S - Descriptor type (0 = system; 1 = code or data).
//         */
//        UINT32 DescriptorType : 1;
//
//        /**
//         * [Bits 6:5] DPL - Descriptor privilege level.
//         */
//        UINT32 DescriptorPrivilegeLevel : 2;
//
//        /**
//         * [Bit 7] P - Segment present.
//         */
//        UINT32 Present : 1;
//
//        UINT32 Reserved1 : 4;
//
//        /**
//         * [Bit 12] AVL - Available for use by system software.
//         */
//        UINT32 AvailableBit : 1;
//
//        /**
//         * [Bit 13] Reserved (except for CS). L - 64-bit mode active (for CS only).
//         */
//        UINT32 LongMode : 1;
//
//        /**
//         * [Bit 14] D/B - Default operation size (0 = 16-bit segment; 1 = 32-bit segment).
//         */
//        UINT32 DefaultBig : 1;
//
//        /**
//         * [Bit 15] G - Granularity.
//         */
//        UINT32 Granularity : 1;
//        /**
//         * [Bit 16] Segment unusable (0 = usable; 1 = unusable).
//         */
//        UINT32 Unusable : 1;
//        UINT32 Reserved2 : 15;
//    };
//
//    UINT32 AsUInt;
//} VMX_SEGMENT_ACCESS_RIGHTS_TYPE;


//	typedef struct xed_immdis_s {
////        union xed_immdis_values_t value;
//        unsigned int currently_used_space :4; // current number of assigned bytes
//        unsigned int max_allocated_space :4; // max allocation, 4 or 8
//        int present : 1;
//                int immediate_is_unsigned : 1;
//    } xed_immdis_t;
//
