package Zycore

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\Defines.h.back

const (
	ZYCORE_DEFINES_H = //col:1
	ZYAN_MACRO_CONCAT(x, = y
) x ## y //col:2
ZYAN_MACRO_CONCAT_EXPAND(x, = y) ZYAN_MACRO_CONCAT(x, y) //col:3
ZYAN_UNUSED(x) = (void)(x) //col:4
ZYAN_BITFIELD(x) =: x //col:5
ZYAN_REQUIRES_LIBC =  //col:6
ZYAN_ARRAY_LENGTH(a) = (sizeof(a) / sizeof((a)[0])) //col:7
ZYAN_MIN(a, = b) (((a) < (b)) ? (a): (b)) //col:8
ZYAN_MAX(a, = b) (((a) > (b)) ? (a) : (b)) //col:9
ZYAN_ABS(a) = (((a) < 0) ? -(a): (a)) //col:10
ZYAN_IS_POWER_OF_2(x) = (((x) & ((x) - 1)) == 0) //col:11
ZYAN_IS_ALIGNED_TO(x, = align) (((x) & ((align) - 1)) == 0) //col:12
ZYAN_ALIGN_UP(x, = align) (((x) + (align) - 1) & ~((align) - 1)) //col:13
ZYAN_ALIGN_DOWN(x, = align) (((x) - 1) & ~((align) - 1)) //col:14
ZYAN_NEEDS_BIT(n, = b) (((unsigned long)(n) >> (b)) > 0) //col:15
ZYAN_BITS_TO_REPRESENT(n) = ( ZYAN_NEEDS_BIT(n, 0) + ZYAN_NEEDS_BIT(n, 1) + ZYAN_NEEDS_BIT(n, 2) + ZYAN_NEEDS_BIT(n,  3) + ZYAN_NEEDS_BIT(n, 4) + ZYAN_NEEDS_BIT(n, 5) + ZYAN_NEEDS_BIT(n, 6) + ZYAN_NEEDS_BIT(n, 7) + ZYAN_NEEDS_BIT(n, 8) + ZYAN_NEEDS_BIT(n, 9) + ZYAN_NEEDS_BIT(n, 10) + ZYAN_NEEDS_BIT(n, 11) + ZYAN_NEEDS_BIT(n, 12) + ZYAN_NEEDS_BIT(n, 13) + ZYAN_NEEDS_BIT(n, 14) + ZYAN_NEEDS_BIT(n, 15) + ZYAN_NEEDS_BIT(n, 16) + ZYAN_NEEDS_BIT(n, 17) + ZYAN_NEEDS_BIT(n, 18) + ZYAN_NEEDS_BIT(n, 19) + ZYAN_NEEDS_BIT(n, 20) + ZYAN_NEEDS_BIT(n, 21) + ZYAN_NEEDS_BIT(n, 22) + ZYAN_NEEDS_BIT(n, 23) + ZYAN_NEEDS_BIT(n, 24) + ZYAN_NEEDS_BIT(n, 25) + ZYAN_NEEDS_BIT(n, 26) + ZYAN_NEEDS_BIT(n, 27) + ZYAN_NEEDS_BIT(n, 28) + ZYAN_NEEDS_BIT(n, 29) + ZYAN_NEEDS_BIT(n, 30) + ZYAN_NEEDS_BIT(n, 31) ) //col:16
)
