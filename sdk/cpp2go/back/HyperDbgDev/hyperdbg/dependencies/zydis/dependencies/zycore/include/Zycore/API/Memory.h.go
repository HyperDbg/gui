package API
const(
ZYCORE_MEMORY_H =  //col:33
)
type #if   defined(ZYAN_WINDOWS) uint32
const(
#if   defined(ZYAN_WINDOWS) typedef enum ZyanMemoryPageProtection_ = 1  //col:57
    ZYAN_PAGE_READONLY           typedef enum ZyanMemoryPageProtection_ =  PAGE_READONLY             //col:59
    ZYAN_PAGE_READWRITE          typedef enum ZyanMemoryPageProtection_ =  PAGE_READWRITE  //col:60
    ZYAN_PAGE_EXECUTE            typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE               //col:61
    ZYAN_PAGE_EXECUTE_READ       typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE_READ  //col:62
    ZYAN_PAGE_EXECUTE_READWRITE  typedef enum ZyanMemoryPageProtection_ =  PAGE_EXECUTE_READWRITE  //col:63
#elif defined(ZYAN_POSIX) typedef enum ZyanMemoryPageProtection_ = 7  //col:65
    ZYAN_PAGE_READONLY           typedef enum ZyanMemoryPageProtection_ =  PROT_READ             //col:67
    ZYAN_PAGE_READWRITE          typedef enum ZyanMemoryPageProtection_ =  PROT_READ | PROT_WRITE  //col:68
    ZYAN_PAGE_EXECUTE            typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC               //col:69
    ZYAN_PAGE_EXECUTE_READ       typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC | PROT_READ  //col:70
    ZYAN_PAGE_EXECUTE_READWRITE  typedef enum ZyanMemoryPageProtection_ =  PROT_EXEC | PROT_READ | PROT_WRITE  //col:71
#endif typedef enum ZyanMemoryPageProtection_ = 13  //col:73
)
type (
Memory interface{
  Zyan Core Library ()(ok bool)//col:74
}

)
func NewMemory() { return & memory{} }
func (m *memory)  Zyan Core Library ()(ok bool){//col:74
return true
}

