package Zycore

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\include\Zycore\LibC.h.back

type (
	LibC interface {
		ZYAN_INLINE_voidPtr_ZYAN_MEMCHR() (ok bool)   //col:15
		ZYAN_INLINE_int_ZYAN_MEMCMP() (ok bool)       //col:28
		ZYAN_INLINE_voidPtr_ZYAN_MEMCPY() (ok bool)   //col:38
		ZYAN_INLINE_voidPtr_ZYAN_MEMMOVE() (ok bool)  //col:57
		ZYAN_INLINE_voidPtr_ZYAN_MEMSET() (ok bool)   //col:66
		ZYAN_INLINE_charPtr_ZYAN_STRCAT() (ok bool)   //col:85
		ZYAN_INLINE_int_ZYAN_STRCMP() (ok bool)       //col:93
		ZYAN_INLINE_int_ZYAN_STRCOLL() (ok bool)      //col:114
		ZYAN_INLINE_ZyanUSize_ZYAN_STRLEN() (ok bool) //col:123
		ZYAN_INLINE_charPtr_ZYAN_STRNCAT() (ok bool)  //col:140
		ZYAN_INLINE_int_ZYAN_STRNCMP() (ok bool)      //col:163
		ZYAN_INLINE_charPtr_ZYAN_STRPBRK() (ok bool)  //col:174
		ZYAN_INLINE_charPtr_ZYAN_STRRCHR() (ok bool)  //col:193
		ZYAN_INLINE_charPtr_ZYAN_STRSTR() (ok bool)   //col:221
	}
	libC struct{}
)

func NewLibC() LibC { return &libC{} }

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMCHR() (ok bool) { //col:15
	/*
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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_MEMCMP() (ok bool) { //col:28
	/*
	   ZYAN_INLINE int ZYAN_MEMCMP(const void* s1, const void* s2, ZyanUSize n)

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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMCPY() (ok bool) { //col:38
	/*
	   ZYAN_INLINE void* ZYAN_MEMCPY(void* dst, const void* src, ZyanUSize n)

	   	{
	   	    volatile ZyanU8* dp = dst;
	   	    const ZyanU8* sp = src;
	   	    while (n--)
	   	    {
	   	        *dp++ = *sp++;
	   	    }
	   	    return dst;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMMOVE() (ok bool) { //col:57
	/*
	   ZYAN_INLINE void* ZYAN_MEMMOVE(void* dst, const void* src, ZyanUSize n)

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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMSET() (ok bool) { //col:66
	/*
	   ZYAN_INLINE void* ZYAN_MEMSET(void* dst, int val, ZyanUSize n)

	   	{
	   	    volatile ZyanU8* p = dst;
	   	    while (n--)
	   	    {
	   	        *p++ = (unsigned char)val;
	   	    }
	   	    return dst;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRCAT() (ok bool) { //col:85
	/*
	   ZYAN_INLINE char* ZYAN_STRCAT(char* dest, const char* src)

	   	{
	   	    char* ret = dest;
	   	    while (*dest)
	   	    {
	   	        dest++;
	   	    }
	   	    while ((*dest++ = *src++));

	   ZYAN_INLINE char* ZYAN_STRCHR(const char* s, int c)

	   	{
	   	    while (*s != (char)c)
	   	    {
	   	        if (!*s++)
	   	        {
	   	            return 0;
	   	        }
	   	    }
	   	    return (char*)s;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRCMP() (ok bool) { //col:93
	/*
	   ZYAN_INLINE int ZYAN_STRCMP(const char* s1, const char* s2)

	   	{
	   	    while (*s1 && (*s1 == *s2))
	   	    {
	   	        s1++, s2++;
	   	    }
	   	    return *(const ZyanU8*)s1 - *(const ZyanU8*)s2;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRCOLL() (ok bool) { //col:114
	/*
	   ZYAN_INLINE int ZYAN_STRCOLL(const char *s1, const char *s2)

	   	{
	   	    ZYAN_UNUSED(s1);
	   	    ZYAN_UNUSED(s2);

	   ZYAN_INLINE char* ZYAN_STRCPY(char* dest, const char* src)

	   	{
	   	    char* ret = dest;
	   	    while ((*dest++ = *src++));

	   ZYAN_INLINE ZyanUSize ZYAN_STRCSPN(const char *s1, const char *s2)

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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_ZyanUSize_ZYAN_STRLEN() (ok bool) { //col:123
	/*
	   ZYAN_INLINE ZyanUSize ZYAN_STRLEN(const char* str)

	   	{
	   	    const char* p = str;
	   	    while (*str)
	   	    {
	   	        ++str;
	   	    }
	   	    return str - p;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRNCAT() (ok bool) { //col:140
	/*
	   ZYAN_INLINE char* ZYAN_STRNCAT(char* dest, const char* src, ZyanUSize n)

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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRNCMP() (ok bool) { //col:163
	/*
	   ZYAN_INLINE int ZYAN_STRNCMP(const char* s1, const char* s2, ZyanUSize n)

	   	{
	   	    while (n--)
	   	    {
	   	        if (*s1++ != *s2++)
	   	        {
	   	            return *(unsigned char*)(s1 - 1) - *(unsigned char*)(s2 - 1);

	   ZYAN_INLINE char* ZYAN_STRNCPY(char* dest, const char* src, ZyanUSize n)

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
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRPBRK() (ok bool) { //col:174
	/*
	   ZYAN_INLINE char* ZYAN_STRPBRK(const char* s1, const char* s2)

	   	{
	   	    while (*s1)
	   	    {
	   	        if(ZYAN_STRCHR(s2, *s1++))
	   	        {
	   	            return (char*)--s1;
	   	        }
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRRCHR() (ok bool) { //col:193
	/*
	   ZYAN_INLINE char* ZYAN_STRRCHR(const char* s, int c)

	   	{
	   	    char* ret = 0;
	   	    do
	   	    {
	   	        if (*s == (char)c)
	   	        {
	   	            ret = (char*)s;
	   	        }
	   	    } while (*s++);

	   ZYAN_INLINE ZyanUSize ZYAN_STRSPN(const char* s1, const char* s2)

	   	{
	   	    ZyanUSize ret = 0;
	   	    while (*s1 && ZYAN_STRCHR(s2, *s1++))
	   	    {
	   	        ret++;
	   	    }
	   	    return ret;
	   	}
	*/
	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRSTR() (ok bool) { //col:221
	/*
	   ZYAN_INLINE char* ZYAN_STRSTR(const char* s1, const char* s2)

	   	{
	   	    const ZyanUSize n = ZYAN_STRLEN(s2);
	   	    while (*s1)
	   	    {
	   	        if (!ZYAN_MEMCMP(s1++, s2, n))
	   	        {
	   	            return (char*)(s1 - 1);

	   ZYAN_INLINE char* ZYAN_STRTOK(char* str, const char* delim)

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
	   	}
	*/
	return true
}

