package Zycore
const(
ZYCORE_LIBC_H =  //col:33
ZYAN_ERRNO =  errno //col:52
ZYAN_VA_START =               va_start //col:65
ZYAN_VA_ARG =                 va_arg //col:66
ZYAN_VA_END =                 va_end //col:67
ZYAN_VA_COPY(dest, = source)  va_copy((dest), (source)) //col:68
ZYAN_FPUTS =      fputs //col:76
ZYAN_FPUTC =      fputc //col:77
ZYAN_FPRINTF =    fprintf //col:78
ZYAN_PRINTF =     printf //col:79
ZYAN_PUTC =       putc //col:80
ZYAN_PUTS =       puts //col:81
ZYAN_SCANF =      scanf //col:82
ZYAN_SSCANF =     sscanf //col:83
ZYAN_VSNPRINTF =  vsnprintf //col:84
ZYAN_STDIN =      stdin //col:91
ZYAN_STDOUT =     stdout //col:92
ZYAN_STDERR =     stderr //col:93
ZYAN_CALLOC =     calloc //col:100
ZYAN_FREE =       free //col:101
ZYAN_MALLOC =     malloc //col:102
ZYAN_REALLOC =    realloc //col:103
ZYAN_MEMCHR =     memchr //col:110
ZYAN_MEMCMP =     memcmp //col:111
ZYAN_MEMCPY =     memcpy //col:112
ZYAN_MEMMOVE =    memmove //col:113
ZYAN_MEMSET =     memset //col:114
ZYAN_STRCAT =     strcat //col:115
ZYAN_STRCHR =     strchr //col:116
ZYAN_STRCMP =     strcmp //col:117
ZYAN_STRCOLL =    strcoll //col:118
ZYAN_STRCPY =     strcpy //col:119
ZYAN_STRCSPN =    strcspn //col:120
ZYAN_STRLEN =     strlen //col:121
ZYAN_STRNCAT =    strncat //col:122
ZYAN_STRNCMP =    strncmp //col:123
ZYAN_STRNCPY =    strncpy //col:124
ZYAN_STRPBRK =    strpbrk //col:125
ZYAN_STRRCHR =    strrchr //col:126
ZYAN_STRSPN =     strspn //col:127
ZYAN_STRSTR =     strstr //col:128
ZYAN_STRTOK =     strtok //col:129
ZYAN_STRXFRM =    strxfrm //col:130
)
type (
LibC interface{
  Zyan Core Library ()(ok bool)//col:247
ZYAN_INLINE int ZYAN_MEMCMP()(ok bool)//col:261
ZYAN_INLINE void* ZYAN_MEMCPY()(ok bool)//col:272
ZYAN_INLINE void* ZYAN_MEMMOVE()(ok bool)//col:292
ZYAN_INLINE void* ZYAN_MEMSET()(ok bool)//col:302
ZYAN_INLINE char* ZYAN_STRCAT()(ok bool)//col:313
ZYAN_INLINE char* ZYAN_STRCHR()(ok bool)//col:325
ZYAN_INLINE int ZYAN_STRCMP()(ok bool)//col:334
ZYAN_INLINE int ZYAN_STRCOLL()(ok bool)//col:344
ZYAN_INLINE char* ZYAN_STRCPY()(ok bool)//col:351
ZYAN_INLINE ZyanUSize ZYAN_STRCSPN()(ok bool)//col:365
ZYAN_INLINE ZyanUSize ZYAN_STRLEN()(ok bool)//col:375
ZYAN_INLINE char* ZYAN_STRNCAT()(ok bool)//col:393
ZYAN_INLINE int ZYAN_STRNCMP()(ok bool)//col:405
ZYAN_INLINE char* ZYAN_STRNCPY()(ok bool)//col:422
ZYAN_INLINE char* ZYAN_STRPBRK()(ok bool)//col:434
ZYAN_INLINE char* ZYAN_STRRCHR()(ok bool)//col:447
ZYAN_INLINE ZyanUSize ZYAN_STRSPN()(ok bool)//col:457
ZYAN_INLINE char* ZYAN_STRSTR()(ok bool)//col:470
ZYAN_INLINE char* ZYAN_STRTOK()(ok bool)//col:491
ZYAN_INLINE ZyanUSize ZYAN_STRXFRM()(ok bool)//col:501
}

)
func NewLibC() { return & libC{} }
func (l *libC)  Zyan Core Library ()(ok bool){//col:247
return true
}

func (l *libC)ZYAN_INLINE int ZYAN_MEMCMP()(ok bool){//col:261
return true
}

func (l *libC)ZYAN_INLINE void* ZYAN_MEMCPY()(ok bool){//col:272
return true
}

func (l *libC)ZYAN_INLINE void* ZYAN_MEMMOVE()(ok bool){//col:292
return true
}

func (l *libC)ZYAN_INLINE void* ZYAN_MEMSET()(ok bool){//col:302
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRCAT()(ok bool){//col:313
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRCHR()(ok bool){//col:325
return true
}

func (l *libC)ZYAN_INLINE int ZYAN_STRCMP()(ok bool){//col:334
return true
}

func (l *libC)ZYAN_INLINE int ZYAN_STRCOLL()(ok bool){//col:344
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRCPY()(ok bool){//col:351
return true
}

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRCSPN()(ok bool){//col:365
return true
}

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRLEN()(ok bool){//col:375
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRNCAT()(ok bool){//col:393
return true
}

func (l *libC)ZYAN_INLINE int ZYAN_STRNCMP()(ok bool){//col:405
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRNCPY()(ok bool){//col:422
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRPBRK()(ok bool){//col:434
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRRCHR()(ok bool){//col:447
return true
}

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRSPN()(ok bool){//col:457
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRSTR()(ok bool){//col:470
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRTOK()(ok bool){//col:491
return true
}

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRXFRM()(ok bool){//col:501
return true
}

