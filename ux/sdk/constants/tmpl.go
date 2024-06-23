package constants

type ErrorKind byte

const (
	ERROR_SUCCESS           ErrorKind = 0
	ERROR_INVALID_PARAMS              = 1
	ERROR_INVALID_STATE               = 2
	ERROR_INVALID_FORMAT              = 3
	ERROR_INVALID_DATA                = 4
	ERROR_INVALID_OPERATION           = 5
	ERROR_NOT_FOUND                   = 6
	ERROR_ALREADY_EXISTS              = 7
	ERROR_ACCESS_DENIED               = 8
	ERROR_NOT_SUPPORTED               = 9
	ERROR_INTERNAL                    = 10
	ERROR_TIMEOUT                     = 11
	ERROR_BUSY                        = 12
	ERROR_NOT_INITIALIZED             = 13
	ERROR_NOT_IMPLEMENTED             = 14
	ERROR_UNKNOWN                     = 15
)

func (k ErrorKind) String() string {
	switch k {
	case ERROR_SUCCESS:
		return "success"
	case ERROR_INVALID_PARAMS:
		return "invalid params"
	case ERROR_INVALID_STATE:
		return "invalid state"
	case ERROR_INVALID_FORMAT:
		return "invalid format"
	case ERROR_INVALID_DATA:
		return "invalid data"
	case ERROR_INVALID_OPERATION:
		return "invalid operation"
	case ERROR_NOT_FOUND:
		return "not found"
	case ERROR_ALREADY_EXISTS:
		return "already exists"
	case ERROR_ACCESS_DENIED:
		return "access denied"
	case ERROR_NOT_SUPPORTED:
		return "not supported"
	case ERROR_INTERNAL:
		return "internal error"
	case ERROR_TIMEOUT:
		return "timeout"
	case ERROR_BUSY:
		return "busy"
	case ERROR_NOT_INITIALIZED:
		return "not initialized"
	case ERROR_NOT_IMPLEMENTED:
		return "not implemented"
	default:
		return "unknown error"
	}
}

func (k ErrorKind) Elements() []ErrorKind {
	return []ErrorKind{
		ERROR_SUCCESS,
		ERROR_INVALID_PARAMS,
		ERROR_INVALID_STATE,
		ERROR_INVALID_FORMAT,
		ERROR_INVALID_DATA,
		ERROR_INVALID_OPERATION,
	}
}
