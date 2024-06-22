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