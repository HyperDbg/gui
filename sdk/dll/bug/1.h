
typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;

typedef unsigned char      uint8_t;   // 无符号8位整数
typedef unsigned short     uint16_t;  // 无符号16位整数
typedef unsigned int       uint32_t;  // 无符号32位整数
typedef unsigned long long uint64_t;  // 无符号64位整数
typedef signed char        int8_t;    // 有符号8位整数
typedef signed short       int16_t;   // 有符号16位整数
typedef signed int         int32_t;   // 有符号32位整数
typedef signed long long   int64_t;   // 有符号64位整数
typedef int bool;           // 使用 typedef 定义 bool 类型

typedef int* intptr_t;
//typedef unsigned char  UCHAR;
//typedef UCHAR     BOOLEAN;  // winnt
//typedef BOOLEAN * PBOOLEAN; // winnt
//typedef BOOLEAN (*CHECK_VMX_OPERATION)();
typedef unsigned char (*CHECK_VMX_OPERATION)();

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








