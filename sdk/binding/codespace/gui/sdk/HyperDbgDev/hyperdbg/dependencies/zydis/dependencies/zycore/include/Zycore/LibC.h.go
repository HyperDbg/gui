package Zycore

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

	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_MEMCMP() (ok bool) { //col:28

	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMCPY() (ok bool) { //col:38

	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMMOVE() (ok bool) { //col:57

	return true
}

func (l *libC) ZYAN_INLINE_voidPtr_ZYAN_MEMSET() (ok bool) { //col:66

	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRCAT() (ok bool) { //col:85

	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRCMP() (ok bool) { //col:93

	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRCOLL() (ok bool) { //col:114

	return true
}

func (l *libC) ZYAN_INLINE_ZyanUSize_ZYAN_STRLEN() (ok bool) { //col:123

	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRNCAT() (ok bool) { //col:140

	return true
}

func (l *libC) ZYAN_INLINE_int_ZYAN_STRNCMP() (ok bool) { //col:163

	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRPBRK() (ok bool) { //col:174

	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRRCHR() (ok bool) { //col:193

	return true
}

func (l *libC) ZYAN_INLINE_charPtr_ZYAN_STRSTR() (ok bool) { //col:221

	return true
}
