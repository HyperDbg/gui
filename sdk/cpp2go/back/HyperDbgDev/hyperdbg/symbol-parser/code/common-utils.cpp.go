package code


type (
CommonUtils interface{
IsFileExists()(ok bool)//col:26
IsDirExists()(ok bool)//col:45
CreateDirectoryRecursive()(ok bool)//col:70
Split()(ok bool)//col:100
}















)

func NewCommonUtils() { return & commonUtils{} }

func (c *commonUtils)IsFileExists()(ok bool){//col:26





return true
}
















func (c *commonUtils)IsDirExists()(ok bool){//col:45






return true
}
















func (c *commonUtils)CreateDirectoryRecursive()(ok bool){//col:70












return true
}
















func (c *commonUtils)Split()(ok bool){//col:100


















return true
}


















