package Zycore
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\LibC.h.back

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
/*  Zyan Core Library (Zycore-C)
  Original Author : Florian Bernd, Joel Hoener
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#ifndef ZYCORE_LIBC_H
#define ZYCORE_LIBC_H
#ifndef ZYAN_CUSTOM_LIBC
#ifndef ZYAN_NO_LIBC
#include <errno.h>
#define ZYAN_ERRNO  errno
#include <stdarg.h>
typedef va_list ZyanVAList;
#define ZYAN_VA_START               va_start
#define ZYAN_VA_ARG                 va_arg
#define ZYAN_VA_END                 va_end
#define ZYAN_VA_COPY(dest, source)  va_copy((dest), (source))
#include <stdio.h>
#define ZYAN_FPUTS      fputs
#define ZYAN_FPUTC      fputc
#define ZYAN_FPRINTF    fprintf
#define ZYAN_PRINTF     printf
#define ZYAN_PUTC       putc
#define ZYAN_PUTS       puts
#define ZYAN_SCANF      scanf
#define ZYAN_SSCANF     sscanf
#define ZYAN_VSNPRINTF  vsnprintf
typedef FILE ZyanFile;
#define ZYAN_STDIN      stdin
#define ZYAN_STDOUT     stdout
#define ZYAN_STDERR     stderr
#include <stdlib.h>
#define ZYAN_CALLOC     calloc
#define ZYAN_FREE       free
#define ZYAN_MALLOC     malloc
#define ZYAN_REALLOC    realloc
#include <string.h>
#define ZYAN_MEMCHR     memchr
#define ZYAN_MEMCMP     memcmp
#define ZYAN_MEMCPY     memcpy
#define ZYAN_MEMMOVE    memmove
#define ZYAN_MEMSET     memset
#define ZYAN_STRCAT     strcat
#define ZYAN_STRCHR     strchr
#define ZYAN_STRCMP     strcmp
#define ZYAN_STRCOLL    strcoll
#define ZYAN_STRCPY     strcpy
#define ZYAN_STRCSPN    strcspn
#define ZYAN_STRLEN     strlen
#define ZYAN_STRNCAT    strncat
#define ZYAN_STRNCMP    strncmp
#define ZYAN_STRNCPY    strncpy
#define ZYAN_STRPBRK    strpbrk
#define ZYAN_STRRCHR    strrchr
#define ZYAN_STRSPN     strspn
#define ZYAN_STRSTR     strstr
#define ZYAN_STRTOK     strtok
#define ZYAN_STRXFRM    strxfrm
#include <Zycore/Defines.h>
#include <Zycore/Types.h>
 * These implementations are by no means optimized and will be outperformed by pretty much any
 * libc implementation out there. We do not aim towards providing competetive implementations here,
 * but towards providing a last resort fallback for environments without a working libc.
#if defined(ZYAN_MSVC) || defined(ZYAN_ICC)
typedef char* ZyanVAList;
#   define ZYAN_VA_START __crt_va_start
#   define ZYAN_VA_ARG   __crt_va_arg
#   define ZYAN_VA_END   __crt_va_end
#   define ZYAN_VA_COPY(destination, source) ((destination) = (source))
#elif defined(ZYAN_GNUC)
typedef __builtin_va_list  ZyanVAList;
#   define ZYAN_VA_START(v, l)  __builtin_va_start(v, l)
#   define ZYAN_VA_END(v)       __builtin_va_end(v)
#   define ZYAN_VA_ARG(v, l)    __builtin_va_arg(v, l)
#   define ZYAN_VA_COPY(d, s)   __builtin_va_copy(d, s)
#else
#   error "Unsupported compiler for no-libc mode."
#endif
ZYAN_INLINE void* ZYAN_MEMCHR(const void* str, int c, ZyanUSize n)
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

func (l *libC)ZYAN_INLINE int ZYAN_MEMCMP()(ok bool){//col:261
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

func (l *libC)ZYAN_INLINE void* ZYAN_MEMCPY()(ok bool){//col:272
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

func (l *libC)ZYAN_INLINE void* ZYAN_MEMMOVE()(ok bool){//col:292
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

func (l *libC)ZYAN_INLINE void* ZYAN_MEMSET()(ok bool){//col:302
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRCAT()(ok bool){//col:313
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRCHR()(ok bool){//col:325
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

func (l *libC)ZYAN_INLINE int ZYAN_STRCMP()(ok bool){//col:334
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

func (l *libC)ZYAN_INLINE int ZYAN_STRCOLL()(ok bool){//col:344
/*ZYAN_INLINE int ZYAN_STRCOLL(const char *s1, const char *s2)
{
    ZYAN_UNUSED(s1);
    ZYAN_UNUSED(s2);
    return 0;
}*/
return true
}

func (l *libC)ZYAN_INLINE char* ZYAN_STRCPY()(ok bool){//col:351
/*ZYAN_INLINE char* ZYAN_STRCPY(char* dest, const char* src)
{
    char* ret = dest;
    while ((*dest++ = *src++));
    return ret;
}*/
return true
}

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRCSPN()(ok bool){//col:365
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

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRLEN()(ok bool){//col:375
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRNCAT()(ok bool){//col:393
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

func (l *libC)ZYAN_INLINE int ZYAN_STRNCMP()(ok bool){//col:405
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRNCPY()(ok bool){//col:422
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRPBRK()(ok bool){//col:434
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRRCHR()(ok bool){//col:447
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

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRSPN()(ok bool){//col:457
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRSTR()(ok bool){//col:470
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

func (l *libC)ZYAN_INLINE char* ZYAN_STRTOK()(ok bool){//col:491
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

func (l *libC)ZYAN_INLINE ZyanUSize ZYAN_STRXFRM()(ok bool){//col:501
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



