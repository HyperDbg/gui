package vmx
type typedef enum
uint32
const (
	typedef   enum = 1 //col:26
	AccessOut      = 0 //col:28
	AccessIn       = 1 //col:29
)
type typedef enum
uint32
const (
	typedef       enum = 1 //col:36
	OpEncodingDx       = 0 //col:38
	OpEncodingImm      = 1 //col:39
)
type (
	IoHandler interface {
		__inbyte() (ok bool)         //col:54
		__inword() (ok bool)         //col:64
		__indword() (ok bool)        //col:74
		__inbytestring() (ok bool)   //col:84
		__inwordstring() (ok bool)   //col:94
		__indwordstring() (ok bool)  //col:104
		__outbyte() (ok bool)        //col:114
		__outword() (ok bool)        //col:124
		__outdword() (ok bool)       //col:134
		__outbytestring() (ok bool)  //col:144
		__outwordstring() (ok bool)  //col:154
		__outdwordstring() (ok bool) //col:164
	}
)
func NewIoHandler() { return &ioHandler{} }
func (i *ioHandler) __inbyte() (ok bool) { //col:54
	return true
}

func (i *ioHandler) __inword() (ok bool) { //col:64
	return true
}

func (i *ioHandler) __indword() (ok bool) { //col:74
	return true
}

func (i *ioHandler) __inbytestring() (ok bool) { //col:84
	return true
}

func (i *ioHandler) __inwordstring() (ok bool) { //col:94
	return true
}

func (i *ioHandler) __indwordstring() (ok bool) { //col:104
	return true
}

func (i *ioHandler) __outbyte() (ok bool) { //col:114
	return true
}

func (i *ioHandler) __outword() (ok bool) { //col:124
	return true
}

func (i *ioHandler) __outdword() (ok bool) { //col:134
	return true
}

func (i *ioHandler) __outbytestring() (ok bool) { //col:144
	return true
}

func (i *ioHandler) __outwordstring() (ok bool) { //col:154
	return true
}

func (i *ioHandler) __outdwordstring() (ok bool) { //col:164
	return true
}

