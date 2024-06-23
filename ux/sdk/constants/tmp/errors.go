package constants

type ErrorsKind int

const (
	ERROR_SUCCESS          ErrorsKind = 0
	ERROR_INVALID_FUNCTION            = 1
	ERROR_FILE_NOT_FOUND              = 2
	ERROR_PATH_NOT_FOUND              = 3
)

func (k ErrorsKind) String() string {
	switch k {
	case 0:
		return "ERROR_SUCCESS"
	case 1:
		return "ERROR_INVALID_FUNCTION"
	case 2:
		return "ERROR_FILE_NOT_FOUND"
	case 3:
		return "ERROR_PATH_NOT_FOUND"
	default:
		return "unknown"
	}
}

func (k ErrorsKind) Elements() []ErrorsKind {
	return []ErrorsKind{
		ERROR_SUCCESS,
		ERROR_INVALID_FUNCTION,
		ERROR_FILE_NOT_FOUND,
		ERROR_PATH_NOT_FOUND,
	}
}
