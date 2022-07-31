#include "pch.h"
std::string
Uint64ToString(UINT64 value)
{
    ostringstream Os;
    Os << setw(16) << setfill('0') << hex << value;
    return Os.str();
}

BOOLEAN
StringReplace(std::string & str, const std::string & from, const std::string & to)
{
    size_t start_pos = str.find(from);
    if (start_pos == string::npos)
        return FALSE;
    str.replace(start_pos, from.length(), to);
    return TRUE;
}

std::string
ConvertToString(char * Str)
{
    string s(Str);
    return s;
}

