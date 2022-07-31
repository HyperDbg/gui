package code
type (
CommonUtils interface{
IsFileExists()(ok bool)//col:26
IsDirExists()(ok bool)//col:44
CreateDirectoryRecursive()(ok bool)//col:68
Split()(ok bool)//col:97
}

)
func NewCommonUtils() { return & commonUtils{} }
func (c *commonUtils)IsFileExists()(ok bool){//col:26
return true
}

func (c *commonUtils)IsDirExists()(ok bool){//col:44
return true
}

func (c *commonUtils)CreateDirectoryRecursive()(ok bool){//col:68
return true
}

func (c *commonUtils)Split()(ok bool){//col:97
return true
}

