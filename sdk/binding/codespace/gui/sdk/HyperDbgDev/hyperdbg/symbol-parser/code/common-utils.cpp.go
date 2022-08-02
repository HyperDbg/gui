package code



type (
	CommonUtils interface {
		IsFileExists() (ok bool)             //col:13
		CreateDirectoryRecursive() (ok bool) //col:25
	}
	commonUtils struct{}
)

func NewCommonUtils() CommonUtils { return &commonUtils{} }

func (c *commonUtils) IsFileExists() (ok bool) { //col:13


















	return true
}


func (c *commonUtils) CreateDirectoryRecursive() (ok bool) { //col:25















	return true
}


