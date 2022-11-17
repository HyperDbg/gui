package common

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\common\common.cpp.back

type (
	Common interface {
		SeparateTo64BitValue() (ok bool)          //col:69
		IsDecimalNotation() (ok bool)             //col:86
		HexToBytes() (ok bool)                    //col:158
		ConvertStringToUInt32() (ok bool)         //col:232
		VmxSupportDetection() (ok bool)           //col:298
		GetConfigFilePath() (ok bool)             //col:395
		CheckCanonicalVirtualAddress() (ok bool)  //col:417
		CheckIfAddressIsValidUsingTsx() (ok bool) //col:507
	}
	common struct{}
)

func NewCommon() Common { return &common{} }

func (c *common) SeparateTo64BitValue() (ok bool) { //col:69
	/*
	   SeparateTo64BitValue(UINT64 Value)

	   	{
	   	    ostringstream OstringStream;
	   	    string        Temp;
	   	    OstringStream << setw(16) << setfill('0') << hex << Value;
	   	    Temp = OstringStream.str();
	   	    Temp.insert(8, 1, '`');

	   PrintBits(size_t const size, void const * const ptr)

	   	{
	   	    unsigned char * b = (unsigned char *)ptr;
	   	    unsigned char   byte;
	   	    int             i, j;
	   	    for (i = size - 1; i >= 0; i--)
	   	    {
	   	        for (j = 7; j >= 0; j--)
	   	        {
	   	            byte = (b[i] >> j) & 1;
	   	            ShowMessages("%u", byte);
	   	        ShowMessages(" ", byte);

	   Replace(std::string & str, const std::string & from, const std::string & to)

	   	{
	   	    size_t start_pos = str.find(from);
	   	    if (start_pos == std::string::npos)
	   	        return FALSE;
	   	    str.replace(start_pos, from.length(), to);

	   ReplaceAll(string & str, const string & from, const string & to)

	   	{
	   	    size_t SartPos = 0;
	   	    if (from.empty())
	   	        return;
	   	    while ((SartPos = str.find(from, SartPos)) != std::string::npos)
	   	    {
	   	        str.replace(SartPos, from.length(), to);
	   	        SartPos += to.length();

	   Split(const string & s, const char & c)

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
	   	    if (!buff.empty())
	   	        v.push_back(buff);

	   IsNumber(const string & str)

	   	{
	   	    return !str.empty() &&
	   	           (str.find_first_not_of("[0123456789]") == std::string::npos);

	   IsHexNotation(const string & s)

	   	{
	   	    BOOLEAN IsAnyThing = FALSE;
	   	    for (auto & CptrChar : s)
	   	    {
	   	        IsAnyThing = TRUE;
	   	        if (!isxdigit(CptrChar))
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    if (IsAnyThing)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (c *common) IsDecimalNotation() (ok bool) { //col:86
	/*
	   IsDecimalNotation(const string & s)

	   	{
	   	    BOOLEAN IsAnyThing = FALSE;
	   	    for (auto & CptrChar : s)
	   	    {
	   	        IsAnyThing = TRUE;
	   	        if (!isdigit(CptrChar))
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    if (IsAnyThing)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (c *common) HexToBytes() (ok bool) { //col:158
	/*
	   HexToBytes(const string & hex)

	   	{
	   	    vector<char> Bytes;
	   	    for (unsigned int i = 0; i < hex.length(); i += 2)
	   	        std::string byteString = hex.substr(i, 2);
	   	        char        byte       = (char)strtol(byteString.c_str(), NULL, 16);
	   	        Bytes.push_back(byte);

	   ConvertStringToUInt64(string TextToConvert, PUINT64 Result)

	   	{
	   	    BOOLEAN IsDecimal = FALSE;
	   	    if (TextToConvert.rfind("0x", 0) == 0 || TextToConvert.rfind("0X", 0) == 0 ||
	   	        TextToConvert.rfind("\\x", 0) == 0 ||
	   	        TextToConvert.rfind("\\X", 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 2);
	   	    else if (TextToConvert.rfind('x', 0) == 0 ||
	   	             TextToConvert.rfind('X', 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 1);
	   	    else if (TextToConvert.rfind("0n", 0) == 0 || TextToConvert.rfind("0N", 0) == 0 ||
	   	             TextToConvert.rfind("\\n", 0) == 0 ||
	   	             TextToConvert.rfind("\\N", 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 2);
	   	    else if (TextToConvert.rfind('n', 0) == 0 ||
	   	             TextToConvert.rfind('N', 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 1);
	   	    TextToConvert.erase(remove(TextToConvert.begin(), TextToConvert.end(), '`'),
	   	                        TextToConvert.end());
	   	    if (IsDecimal)
	   	    {
	   	        if (!IsDecimalNotation(TextToConvert))
	   	        {
	   	            return FALSE;
	   	        }
	   	        else
	   	        {
	   	            errno                                 = 0;
	   	            char *                       unparsed = NULL;
	   	            const char *                 s        = TextToConvert.c_str();
	   	            const unsigned long long int n        = strtoull(s, &unparsed, 10);
	   	            if (errno || (!n && s == unparsed))
	   	            {
	   	                return FALSE;
	   	            }
	   	            *Result = n;
	   	            return TRUE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        if (!IsHexNotation(TextToConvert))
	   	        {
	   	            return FALSE;
	   	        }
	   	        else
	   	        {
	   	            const char * Text         = TextToConvert.c_str();
	   	            unsigned long long result = strtoull(Text, NULL, 16);
	   	            if (errno == EINVAL)
	   	            {
	   	                return FALSE;
	   	            }
	   	            else if (errno == ERANGE)
	   	            {
	   	                return TRUE;
	   	            }
	   	            return TRUE;
	   	        }
	   	    }
	   	}
	*/
	return true
}

func (c *common) ConvertStringToUInt32() (ok bool) { //col:232
	/*
	   ConvertStringToUInt32(string TextToConvert, PUINT32 Result)

	   	{
	   	    BOOLEAN IsDecimal = FALSE;
	   	    if (TextToConvert.rfind("0x", 0) == 0 || TextToConvert.rfind("0X", 0) == 0 ||
	   	        TextToConvert.rfind("\\x", 0) == 0 ||
	   	        TextToConvert.rfind("\\X", 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 2);
	   	    else if (TextToConvert.rfind('x', 0) == 0 ||
	   	             TextToConvert.rfind('X', 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 1);
	   	    else if (TextToConvert.rfind("0n", 0) == 0 || TextToConvert.rfind("0N", 0) == 0 ||
	   	             TextToConvert.rfind("\\n", 0) == 0 ||
	   	             TextToConvert.rfind("\\N", 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 2);
	   	    else if (TextToConvert.rfind('n', 0) == 0 ||
	   	             TextToConvert.rfind('N', 0) == 0)
	   	    {
	   	        TextToConvert = TextToConvert.erase(0, 1);
	   	    TextToConvert.erase(remove(TextToConvert.begin(), TextToConvert.end(), '`'),
	   	                        TextToConvert.end());
	   	    if (IsDecimal)
	   	    {
	   	        if (!IsDecimalNotation(TextToConvert))
	   	        {
	   	            return FALSE;
	   	        }
	   	        else
	   	        {
	   	            try
	   	            {
	   	                int i   = std::stoi(TextToConvert);
	   	            catch (std::invalid_argument const & e)
	   	            {
	   	                return FALSE;
	   	            }
	   	            catch (std::out_of_range const & e)
	   	            {
	   	                return FALSE;
	   	            }
	   	            return FALSE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        if (!IsHexNotation(TextToConvert))
	   	        {
	   	            return FALSE;
	   	        }
	   	        else
	   	        {
	   	            UINT32 TempResult;
	   	            TempResult = stoi(TextToConvert, nullptr, 16);

	   HasEnding(string const & fullString, string const & ending)

	   	{
	   	    if (fullString.length() >= ending.length())
	   	    {
	   	        return (0 == fullString.compare(fullString.length() - ending.length(),
	   	                                        ending.length(),
	   	                                        ending));

	   ValidateIP(const string & ip)

	   	{
	   	    vector<string> list = Split(ip, '.');
	   	    if (list.size() != 4)
	   	        return FALSE;
	   	    for (string str : list)
	   	    {
	   	        if (!IsNumber(str) || stoi(str) > 255 || stoi(str) < 0)
	   	            return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (c *common) VmxSupportDetection() (ok bool) { //col:298
	/*
	   VmxSupportDetection()

	   	{
	   	    return AsmVmxSupportDetection();

	   SetPrivilege(HANDLE  hToken,

	   	LPCTSTR lpszPrivilege,
	   	BOOL    bEnablePrivilege

	   )

	   	{
	   	    TOKEN_PRIVILEGES tp;
	   	    LUID             luid;
	   	    if (!LookupPrivilegeValue(NULL,
	   	                              lpszPrivilege,
	   	                              &luid))
	   	    {
	   	        ShowMessages("err, in LookupPrivilegeValue (%x)\n", GetLastError());
	   	    if (bEnablePrivilege)
	   	        tp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
	   	    else
	   	        tp.Privileges[0].Attributes = 0;
	   	    if (!AdjustTokenPrivileges(hToken, FALSE, &tp, sizeof(TOKEN_PRIVILEGES), (PTOKEN_PRIVILEGES)NULL, (PDWORD)NULL))
	   	    {
	   	        ShowMessages("err, in AdjustTokenPrivileges (%x)\n", GetLastError());
	   	    if (GetLastError() == ERROR_NOT_ALL_ASSIGNED)
	   	    {
	   	        ShowMessages("err, the token does not have the specified privilege (ACCESS DENIED!)\n");
	   	        ShowMessages("make sure to run it with administrator privileges\n");

	   ltrim(std::string & s)

	   	{
	   	    s.erase(s.begin(), std::find_if(s.begin(), s.end(), [](int ch) { return !std::isspace(ch); }));

	   rtrim(std::string & s)

	   	{
	   	    s.erase(std::find_if(s.rbegin(), s.rend(), [](int ch) { return !std::isspace(ch); })
	   	                .base(),
	   	            s.end());

	   Trim(std::string & s)

	   	{
	   	    ltrim(s);
	   	    rtrim(s);

	   RemoveSpaces(std::string str)

	   	{
	   	    str.erase(remove(str.begin(), str.end(), ' '), str.end());

	   IsFileExistA(const char * FileName)

	   	{
	   	    struct stat buffer;
	   	    return (stat(FileName, &buffer) == 0);

	   IsFileExistW(const wchar_t * FileName)

	   	{
	   	    struct _stat64i32 buffer;
	   	    return (_wstat(FileName, &buffer) == 0);

	   IsEmptyString(char * Text)

	   	{
	   	    size_t Len;
	   	    if (Text == NULL || Text[0] == '\0')
	   	    {
	   	        return TRUE;
	   	    }
	   	    Len = strlen(Text);
	   	    for (size_t i = 0; i < Len; i++)
	   	    {
	   	        if (Text[i] != ' ' && Text[i] != '\t' && Text[i] != ' \n')
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (c *common) GetConfigFilePath() (ok bool) { //col:395
	/*
	   GetConfigFilePath(PWCHAR ConfigPath)

	   	{
	   	    WCHAR CurrentPath[260] = {0};
	   	    GetModuleFileNameW(NULL, CurrentPath, MAX_PATH);
	   	    PathRemoveFileSpecW(CurrentPath);
	   	    PathCombineW(ConfigPath, CurrentPath, CONFIG_FILE_NAME);

	   ListDirectory(const std::string & Directory, const std::string & Extension)

	   	{
	   	    WIN32_FIND_DATAA         FindData;
	   	    HANDLE                   Translate     = INVALID_HANDLE_VALUE;
	   	    std::string              FullPath = Directory + "\\" + Extension;
	   	    std::vector<std::string> DirList;
	   	    Translate = FindFirstFileA(FullPath.c_str(), &FindData);
	   	    if (Translate == INVALID_HANDLE_VALUE)
	   	        throw std::runtime_error("invalid handle value! please check your path...");
	   	    while (FindNextFileA(Translate, &FindData) != 0)
	   	    {
	   	        DirList.push_back(Directory + "\\" + std::string(FindData.cFileName));
	   	    FindClose(Translate);

	   StringToWString(std::wstring & ws, const std::string & s)

	   	{
	   	    std::wstring wsTmp(s.begin(), s.end());

	   SplitPathAndArgs(std::vector<std::string> & Qargs, const std::string & Command)

	   	{
	   	    int  Len = Command.length();
	   	    for (int i = 0; i < Len; i++)
	   	    {
	   	        int start = i;
	   	        if (Command[i] == '\"')
	   	        {
	   	            Qot = true;
	   	        }
	   	        else if (Command[i] == '\'')
	   	            Sqot = true;
	   	        if (Qot)
	   	        {
	   	            i++;
	   	            start++;
	   	            while (i < Len && Command[i] != '\"')
	   	                i++;
	   	            if (i < Len)
	   	                Qot = false;
	   	            ArgLen = i - start;
	   	            i++;
	   	        }
	   	        else if (Sqot)
	   	        {
	   	            i++;
	   	            while (i < Len && Command[i] != '\'')
	   	                i++;
	   	            if (i < Len)
	   	                Sqot = false;
	   	            ArgLen = i - start;
	   	            i++;
	   	        }
	   	        else
	   	        {
	   	            while (i < Len && Command[i] != ' ')
	   	                i++;
	   	            ArgLen = i - start;
	   	        }
	   	        string Temp = Command.substr(start, ArgLen);
	   	        if (!Temp.empty() && Temp != " ")
	   	        {
	   	            Qargs.push_back(Temp);

	   FindCaseInsensitive(std::string Input, std::string ToSearch, size_t Pos)

	   	{
	   	    std::transform(Input.begin(), Input.end(), Input.begin(), ::tolower);
	   	    std::transform(ToSearch.begin(), ToSearch.end(), ToSearch.begin(), ::tolower);
	   	    return Input.find(ToSearch, Pos);

	   FindCaseInsensitiveW(std::wstring Input, std::wstring ToSearch, size_t Pos)

	   	{
	   	    std::transform(Input.begin(), Input.end(), Input.begin(), ::tolower);
	   	    std::transform(ToSearch.begin(), ToSearch.end(), ToSearch.begin(), ::tolower);
	   	    return Input.find(ToSearch, Pos);

	   ConvertStringVectorToCharPointerArray(const std::string & s)

	   	{
	   	    char * pc = new char[s.size() + 1];
	   	    std::strcpy(pc, s.c_str());

	   GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo)

	   	{
	   	    __cpuidex(CpuInfo, Func, SubFunc);

	   Getx86VirtualAddressWidth()

	   	{
	   	    int Regs[4];
	   	    GetCpuid(CPUID_ADDR_WIDTH, 0, Regs);
	   	    return ((Regs[0] >> 8) & 0x0ff);

	   CheckCpuSupportRtm()

	   	{
	   	    int     Regs1[4];
	   	    int     Regs2[4];
	   	    BOOLEAN Result;
	   	    GetCpuid(0, 0, Regs1);
	   	    GetCpuid(7, 0, Regs2);
	   	    Result = Regs1[0] >= 0x7 && (Regs2[1] & 0x4800) == 0x4800;
	   	    return Result;
	   	}
	*/
	return true
}

func (c *common) CheckCanonicalVirtualAddress() (ok bool) { //col:417
	/*
	   CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress)

	   	{
	   	    UINT64 Addr = (UINT64)VAddr;
	   	    UINT64 MaxVirtualAddrLowHalf, MinVirtualAddressHighHalf;
	   	    UINT32 AddrWidth = g_VirtualAddressWidth;
	   	    MaxVirtualAddrLowHalf = ((UINT64)1ull << (AddrWidth - 1)) - 1;
	   	    MinVirtualAddressHighHalf = ~MaxVirtualAddrLowHalf;
	   	    if ((Addr > MaxVirtualAddrLowHalf) && (Addr < MinVirtualAddressHighHalf))
	   	    {
	   	        *IsKernelAddress = FALSE;
	   	        return FALSE;
	   	    }
	   	    if (MinVirtualAddressHighHalf < Addr)
	   	    {
	   	        *IsKernelAddress = TRUE;
	   	    }
	   	    else
	   	    {
	   	        *IsKernelAddress = FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (c *common) CheckIfAddressIsValidUsingTsx() (ok bool) { //col:507
	/*
	   CheckIfAddressIsValidUsingTsx(UINT64 Address)

	   	{
	   	    UINT32  Status = 0;
	   	    BOOLEAN Result = FALSE;
	   	    CHAR    TempContent;
	   	    if ((Status = _xbegin()) == _XBEGIN_STARTED)
	   	    {
	   	        TempContent = *(CHAR *)Address;
	   	        _xend();

	   CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size)

	   	{
	   	    BOOLEAN IsKernelAddress;
	   	    BOOLEAN Result = FALSE;
	   	    if (!CheckCanonicalVirtualAddress(TargetAddress, &IsKernelAddress))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (IsKernelAddress)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (g_RtmSupport)
	   	    {
	   	        UINT64 AddressToCheck =
	   	            (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
	   	        if (AddressToCheck > PAGE_SIZE)
	   	        {
	   	            UINT64 ReadSize = AddressToCheck;
	   	            while (Size != 0)
	   	            {
	   	                ReadSize = (UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) - TargetAddress;
	   	                if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE)
	   	                {
	   	                    ReadSize = Size;
	   	                }
	   	                if (!CheckIfAddressIsValidUsingTsx(TargetAddress))
	   	                {
	   	                    return FALSE;
	   	                }
	   	                Size          = Size - ReadSize;
	   	                TargetAddress = TargetAddress + ReadSize;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            if (!CheckIfAddressIsValidUsingTsx(TargetAddress))
	   	            {
	   	                return FALSE;
	   	            }
	   	        }
	   	    }
	   	    else
	   	    {
	   	        return FALSE;
	   	        UINT64 AddressToCheck =
	   	            (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
	   	        if (AddressToCheck > PAGE_SIZE)
	   	        {
	   	            UINT64 ReadSize = AddressToCheck;
	   	            while (Size != 0)
	   	            {
	   	                ReadSize = (UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) - TargetAddress;
	   	                if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE)
	   	                {
	   	                    ReadSize = Size;
	   	                }
	   	                try
	   	                {
	   	                    UINT64 ReadingTest = *((UINT64 *)TargetAddress);
	   	                catch (...)
	   	                {
	   	                    return FALSE;
	   	                }
	   	                Size          = Size - ReadSize;
	   	                TargetAddress = TargetAddress + ReadSize;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            try
	   	            {
	   	                UINT64 ReadingTest = *((UINT64 *)TargetAddress);
	   	            catch (...)
	   	            {
	   	                return FALSE;
	   	            }
	   	        }
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}