package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\code\common-utils.cpp.back

type (
	CommonUtils interface {
		IsFileExists() (ok bool)             //col:13
		CreateDirectoryRecursive() (ok bool) //col:25
	}
	commonUtils struct{}
)

func NewCommonUtils() CommonUtils { return &commonUtils{} }

func (c *commonUtils) IsFileExists() (ok bool) { //col:13
	/*
	   IsFileExists(const std::string & FileName)

	   	{
	   	    struct stat buffer;
	   	    return (stat(FileName.c_str(), &buffer) == 0);

	   IsDirExists(const std::string & DirPath)

	   	{
	   	    DWORD Ftyp = GetFileAttributesA(DirPath.c_str());
	   	    if (Ftyp == INVALID_FILE_ATTRIBUTES)
	   	        return FALSE;
	   	    if (Ftyp & FILE_ATTRIBUTE_DIRECTORY)
	   	        return TRUE;
	   	    return FALSE;
	   	}
	*/
	return true
}

func (c *commonUtils) CreateDirectoryRecursive() (ok bool) { //col:25
	/*
	   CreateDirectoryRecursive(const std::string & Path)

	   	{
	   	    size_t Pos = 0;
	   	    do
	   	    {
	   	        Pos = Path.find_first_of("\\/", Pos + 1);
	   	        CreateDirectoryA(Path.substr(0, Pos).c_str(), NULL);
	   	    } while (Pos != std::string::npos && Pos <= Path.size());
	   	    if (IsDirExists(Path))
	   	        return TRUE;
	   	    return FALSE;
	   	}
	*/
	return true
}

