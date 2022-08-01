package Zycore
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\LibC.h.back

const(
ZYCORE_LIBC_H =  //col:1
ZYAN_ERRNO =  errno //col:2
ZYAN_VA_START =               va_start //col:3
ZYAN_VA_ARG =                 va_arg //col:4
ZYAN_VA_END =                 va_end //col:5
ZYAN_VA_COPY(dest, = source)  va_copy((dest), (source)) //col:6
ZYAN_FPUTS =      fputs //col:7
ZYAN_FPUTC =      fputc //col:8
ZYAN_FPRINTF =    fprintf //col:9
ZYAN_PRINTF =     printf //col:10
ZYAN_PUTC =       putc //col:11
ZYAN_PUTS =       puts //col:12
ZYAN_SCANF =      scanf //col:13
ZYAN_SSCANF =     sscanf //col:14
ZYAN_VSNPRINTF =  vsnprintf //col:15
ZYAN_STDIN =      stdin //col:16
ZYAN_STDOUT =     stdout //col:17
ZYAN_STDERR =     stderr //col:18
ZYAN_CALLOC =     calloc //col:19
ZYAN_FREE =       free //col:20
ZYAN_MALLOC =     malloc //col:21
ZYAN_REALLOC =    realloc //col:22
ZYAN_MEMCHR =     memchr //col:23
ZYAN_MEMCMP =     memcmp //col:24
ZYAN_MEMCPY =     memcpy //col:25
ZYAN_MEMMOVE =    memmove //col:26
ZYAN_MEMSET =     memset //col:27
ZYAN_STRCAT =     strcat //col:28
ZYAN_STRCHR =     strchr //col:29
ZYAN_STRCMP =     strcmp //col:30
ZYAN_STRCOLL =    strcoll //col:31
ZYAN_STRCPY =     strcpy //col:32
ZYAN_STRCSPN =    strcspn //col:33
ZYAN_STRLEN =     strlen //col:34
ZYAN_STRNCAT =    strncat //col:35
ZYAN_STRNCMP =    strncmp //col:36
ZYAN_STRNCPY =    strncpy //col:37
ZYAN_STRPBRK =    strpbrk //col:38
ZYAN_STRRCHR =    strrchr //col:39
ZYAN_STRSPN =     strspn //col:40
ZYAN_STRSTR =     strstr //col:41
ZYAN_STRTOK =     strtok //col:42
ZYAN_STRXFRM =    strxfrm //col:43
)

type (
LibC interface{
ZYAN_INLINE_voidPtr_ZYAN_MEMCHR()(ok bool)//col:15
ZYAN_INLINE_int_ZYAN_MEMCMP()(ok bool)//col:28
ZYAN_INLINE_voidPtr_ZYAN_MEMCPY()(ok bool)//col:38
ZYAN_INLINE_voidPtr_ZYAN_MEMMOVE()(ok bool)//col:57
ZYAN_INLINE_voidPtr_ZYAN_MEMSET()(ok bool)//col:66
ZYAN_INLINE_charPtr_ZYAN_STRCAT()(ok bool)//col:76
ZYAN_INLINE_charPtr_ZYAN_STRCHR()(ok bool)//col:87
ZYAN_INLINE_int_ZYAN_STRCMP()(ok bool)//col:95
ZYAN_INLINE_int_ZYAN_STRCOLL()(ok bool)//col:101
ZYAN_INLINE_charPtr_ZYAN_STRCPY()(ok bool)//col:107
ZYAN_INLINE_ZyanUSize_ZYAN_STRCSPN()(ok bool)//col:120
ZYAN_INLINE_ZyanUSize_ZYAN_STRLEN()(ok bool)//col:129
ZYAN_INLINE_charPtr_ZYAN_STRNCAT()(ok bool)//col:146
ZYAN_INLINE_int_ZYAN_STRNCMP()(ok bool)//col:157
ZYAN_INLINE_charPtr_ZYAN_STRNCPY()(ok bool)//col:173
ZYAN_INLINE_charPtr_ZYAN_STRPBRK()(ok bool)//col:184
ZYAN_INLINE_charPtr_ZYAN_STRRCHR()(ok bool)//col:196
ZYAN_INLINE_ZyanUSize_ZYAN_STRSPN()(ok bool)//col:205
ZYAN_INLINE_charPtr_ZYAN_STRSTR()(ok bool)//col:217
ZYAN_INLINE_charPtr_ZYAN_STRTOK()(ok bool)//col:237
ZYAN_INLINE_ZyanUSize_ZYAN_STRXFRM()(ok bool)//col:246
}
)

func NewLibC() { return & libC{} }

func (l *libC)ZYAN_INLINE_voidPtr_ZYAN_MEMCHR()(ok bool){//col:15
/*ZYAN_INLINE void* ZYAN_MEMCHR(const void* str, int c, ZyanUSize n)
{
    const ZyanU8* p = (ZyanU8*)str;
    while (n--)
    {
        if (*p != (ZyanU8)c)
        {
            p++;
        } else
        {
            return (void*)p;
        }
    }
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_int_ZYAN_MEMCMP()(ok bool){//col:28
/*ZYAN_INLINE int ZYAN_MEMCMP(const void* s1, const void* s2, ZyanUSize n)
{
    const ZyanU8* p1 = s1, *p2 = s2;
    while (n--)
    {
        if (*p1 != *p2)
        {
            return *p1 - *p2;
        }
        p1++, p2++;
    }
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_voidPtr_ZYAN_MEMCPY()(ok bool){//col:38
/*ZYAN_INLINE void* ZYAN_MEMCPY(void* dst, const void* src, ZyanUSize n)
{
    volatile ZyanU8* dp = dst;
    const ZyanU8* sp = src;
    while (n--)
    {
        *dp++ = *sp++;
    }
    return dst;
}*/
return true
}

func (l *libC)ZYAN_INLINE_voidPtr_ZYAN_MEMMOVE()(ok bool){//col:57
/*ZYAN_INLINE void* ZYAN_MEMMOVE(void* dst, const void* src, ZyanUSize n)
{
    volatile ZyanU8* pd = dst;
    const ZyanU8* ps = src;
    if (ps < pd)
    {
        for (pd += n, ps += n; n--;)
        {
            *--pd = *--ps;
        }
    } else
    {
        while (n--)
        {
            *pd++ = *ps++;
        }
    }
    return dst;
}*/
return true
}

func (l *libC)ZYAN_INLINE_voidPtr_ZYAN_MEMSET()(ok bool){//col:66
/*ZYAN_INLINE void* ZYAN_MEMSET(void* dst, int val, ZyanUSize n)
{
    volatile ZyanU8* p = dst;
    while (n--)
    {
        *p++ = (unsigned char)val;
    }
    return dst;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRCAT()(ok bool){//col:76
/*ZYAN_INLINE char* ZYAN_STRCAT(char* dest, const char* src)
{
    char* ret = dest;
    while (*dest)
    {
        dest++;
    }
    while ((*dest++ = *src++));
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRCHR()(ok bool){//col:87
/*ZYAN_INLINE char* ZYAN_STRCHR(const char* s, int c)
{
    while (*s != (char)c)
    {
        if (!*s++)
        {
            return 0;
        }
    }
    return (char*)s;
}*/
return true
}

func (l *libC)ZYAN_INLINE_int_ZYAN_STRCMP()(ok bool){//col:95
/*ZYAN_INLINE int ZYAN_STRCMP(const char* s1, const char* s2)
{
    while (*s1 && (*s1 == *s2))
    {
        s1++, s2++;
    }
    return *(const ZyanU8*)s1 - *(const ZyanU8*)s2;
}*/
return true
}

func (l *libC)ZYAN_INLINE_int_ZYAN_STRCOLL()(ok bool){//col:101
/*ZYAN_INLINE int ZYAN_STRCOLL(const char *s1, const char *s2)
{
    ZYAN_UNUSED(s1);
    ZYAN_UNUSED(s2);
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRCPY()(ok bool){//col:107
/*ZYAN_INLINE char* ZYAN_STRCPY(char* dest, const char* src)
{
    char* ret = dest;
    while ((*dest++ = *src++));
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_ZyanUSize_ZYAN_STRCSPN()(ok bool){//col:120
/*ZYAN_INLINE ZyanUSize ZYAN_STRCSPN(const char *s1, const char *s2)
{
    ZyanUSize ret = 0;
    while (*s1)
    {
        if (ZYAN_STRCHR(s2, *s1))
        {
            return ret;
        }
        s1++, ret++;
    }
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_ZyanUSize_ZYAN_STRLEN()(ok bool){//col:129
/*ZYAN_INLINE ZyanUSize ZYAN_STRLEN(const char* str)
{
    const char* p = str;
    while (*str)
    {
        ++str;
    }
    return str - p;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRNCAT()(ok bool){//col:146
/*ZYAN_INLINE char* ZYAN_STRNCAT(char* dest, const char* src, ZyanUSize n)
{
    char* ret = dest;
    while (*dest)
    {
        dest++;
    }
    while (n--)
    {
        if (!(*dest++ = *src++))
        {
            return ret;
        }
    }
    *dest = 0;
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_int_ZYAN_STRNCMP()(ok bool){//col:157
/*ZYAN_INLINE int ZYAN_STRNCMP(const char* s1, const char* s2, ZyanUSize n)
{
    while (n--)
    {
        if (*s1++ != *s2++)
        {
            return *(unsigned char*)(s1 - 1) - *(unsigned char*)(s2 - 1);
        }
    }
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRNCPY()(ok bool){//col:173
/*ZYAN_INLINE char* ZYAN_STRNCPY(char* dest, const char* src, ZyanUSize n)
{
    char* ret = dest;
    do
    {
        if (!n--)
        {
            return ret;
        }
    } while ((*dest++ = *src++));
    while (n--)
    {
        *dest++ = 0;
    }
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRPBRK()(ok bool){//col:184
/*ZYAN_INLINE char* ZYAN_STRPBRK(const char* s1, const char* s2)
{
    while (*s1)
    {
        if(ZYAN_STRCHR(s2, *s1++))
        {
            return (char*)--s1;
        }
    }
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRRCHR()(ok bool){//col:196
/*ZYAN_INLINE char* ZYAN_STRRCHR(const char* s, int c)
{
    char* ret = 0;
    do
    {
        if (*s == (char)c)
        {
            ret = (char*)s;
        }
    } while (*s++);
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_ZyanUSize_ZYAN_STRSPN()(ok bool){//col:205
/*ZYAN_INLINE ZyanUSize ZYAN_STRSPN(const char* s1, const char* s2)
{
    ZyanUSize ret = 0;
    while (*s1 && ZYAN_STRCHR(s2, *s1++))
    {
        ret++;
    }
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRSTR()(ok bool){//col:217
/*ZYAN_INLINE char* ZYAN_STRSTR(const char* s1, const char* s2)
{
    const ZyanUSize n = ZYAN_STRLEN(s2);
    while (*s1)
    {
        if (!ZYAN_MEMCMP(s1++, s2, n))
        {
            return (char*)(s1 - 1);
        }
    }
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE_charPtr_ZYAN_STRTOK()(ok bool){//col:237
/*ZYAN_INLINE char* ZYAN_STRTOK(char* str, const char* delim)
{
    static char* p = 0;
    if (str)
    {
        p = str;
    } else
    if (!p)
    {
        return 0;
    }
    str = p + ZYAN_STRSPN(p, delim);
    p = str + ZYAN_STRCSPN(str, delim);
    if (p == str)
    {
        return p = 0;
    }
    p = *p ? *p = 0, p + 1 : 0;
    return str;
}*/
return true
}

func (l *libC)ZYAN_INLINE_ZyanUSize_ZYAN_STRXFRM()(ok bool){//col:246
/*ZYAN_INLINE ZyanUSize ZYAN_STRXFRM(char* dest, const char* src, ZyanUSize n)
{
    const ZyanUSize n2 = ZYAN_STRLEN(src);
    if (n > n2)
    {
        ZYAN_STRCPY(dest, src);
    }
    return n2;
}*/
return true
}



