package API
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\API\Memory.h.back

const(
ZYCORE_MEMORY_H =  //col:1
)

const(
#if   defined(ZYAN_WINDOWS) = 1  //col:3
    ZYAN_PAGE_READONLY           =  PAGE_READONLY             //col:4
    ZYAN_PAGE_READWRITE          =  PAGE_READWRITE  //col:5
    ZYAN_PAGE_EXECUTE            =  PAGE_EXECUTE               //col:6
    ZYAN_PAGE_EXECUTE_READ       =  PAGE_EXECUTE_READ  //col:7
    ZYAN_PAGE_EXECUTE_READWRITE  =  PAGE_EXECUTE_READWRITE  //col:8
#elif defined(ZYAN_POSIX) = 7  //col:9
    ZYAN_PAGE_READONLY           =  PROT_READ             //col:10
    ZYAN_PAGE_READWRITE          =  PROT_READ | PROT_WRITE  //col:11
    ZYAN_PAGE_EXECUTE            =  PROT_EXEC               //col:12
    ZYAN_PAGE_EXECUTE_READ       =  PROT_EXEC | PROT_READ  //col:13
    ZYAN_PAGE_EXECUTE_READWRITE  =  PROT_EXEC | PROT_READ | PROT_WRITE  //col:14
#endif = 13  //col:15
)




