#pragma once
BOOLEAN
IsFileExists(const std::string & FileName);
BOOLEAN
IsDirExists(const std::string & DirPath);
BOOLEAN
CreateDirectoryRecursive(const std::string & Path);
const std::vector<std::string>
Split(const std::string & s, const char & c);
