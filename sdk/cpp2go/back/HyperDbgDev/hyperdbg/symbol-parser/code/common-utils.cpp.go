package code
//back\HyperDbgDev\hyperdbg\symbol-parser\code\common-utils.cpp.back

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
/*IsFileExists(const std::string & FileName)
{
    struct stat buffer;
    return (stat(FileName.c_str(), &buffer) == 0);
}*/
return true
}

func (c *commonUtils)IsDirExists()(ok bool){//col:44
/*IsDirExists(const std::string & DirPath)
{
    DWORD Ftyp = GetFileAttributesA(DirPath.c_str());
    if (Ftyp == INVALID_FILE_ATTRIBUTES)
    if (Ftyp & FILE_ATTRIBUTE_DIRECTORY)
}*/
return true
}

func (c *commonUtils)CreateDirectoryRecursive()(ok bool){//col:68
/*CreateDirectoryRecursive(const std::string & Path)
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
}*/
return true
}

func (c *commonUtils)Split()(ok bool){//col:97
/*Split(const std::string & s, const char & c)
{
    string         buff {""};
    vector<string> v;
    for (auto n : s)
    {
        if (n != c)
            buff += n;
        else if (n == c && !buff.empty())
        {
            v.push_back(buff);
            buff.clear();
        }
    }
    if (!buff.empty())
        v.push_back(buff);
    return v;
}*/
return true
}



