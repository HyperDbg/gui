package API


const(
ZYCORE_MEMORY_H =  //col:33
)

const(
#if   defined(ZYAN_WINDOWS) = 1  //col:57
    ZYAN_PAGE_READONLY           =  PAGE_READONLY             //col:59
    ZYAN_PAGE_READWRITE          =  PAGE_READWRITE  //col:60
    ZYAN_PAGE_EXECUTE            =  PAGE_EXECUTE               //col:61
    ZYAN_PAGE_EXECUTE_READ       =  PAGE_EXECUTE_READ  //col:62
    ZYAN_PAGE_EXECUTE_READWRITE  =  PAGE_EXECUTE_READWRITE  //col:63
#elif defined(ZYAN_POSIX) = 7  //col:65
    ZYAN_PAGE_READONLY           =  PROT_READ             //col:67
    ZYAN_PAGE_READWRITE          =  PROT_READ | PROT_WRITE  //col:68
    ZYAN_PAGE_EXECUTE            =  PROT_EXEC               //col:69
    ZYAN_PAGE_EXECUTE_READ       =  PROT_EXEC | PROT_READ  //col:70
    ZYAN_PAGE_EXECUTE_READWRITE  =  PROT_EXEC | PROT_READ | PROT_WRITE  //col:71
#endif = 13  //col:73
)



type (
Memory interface{
#if   defined()(ok bool)//col:74
}

















)

func NewMemory() { return & memory{} }

func (m *memory)#if   defined()(ok bool){//col:74























return true
}




















