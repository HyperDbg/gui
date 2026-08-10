/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#pragma once
#include "pluginsdk/BeaEngine.h"
#include "pluginsdk/bridgemain.h"
#include <vector>
#include <string>
#include <map>


struct Instruction_t{
	DISASM BeaStruct;
	unsigned int Length;

	Instruction_t(DISASM *dis, unsigned int len){
		BeaStruct = *dis;
		Length = len;
	}

	Instruction_t(){
		Length = UNKNOWN_OPCODE;
	}
};

struct ArgumentInfo_t
{
	std::string Type;
	std::string Name;

	ArgumentInfo_t(std::string t,std::string n){
		Type=t;
		Name = n;
	}

	ArgumentInfo_t(){}
};

struct FunctionInfo_t
{
	std::string DLLName;
	std::string ReturnType;
	std::string Name;
	std::vector<ArgumentInfo_t> Arguments;
	bool invalid;

	FunctionInfo_t(){
		invalid=false;
	}

	FunctionInfo_t(std::string dll,std::string ret,std::string name, std::vector<ArgumentInfo_t> args){
		DLLName=dll;
		ReturnType=ret;
		Name=name;
		Arguments = args;
		invalid = false;
	}

	bool operator==(const FunctionInfo_t& rhs) const
	{
		return (bool)(stricmp(Name.c_str(), rhs.Name.c_str()) < 0);
	}
	bool operator<(const FunctionInfo_t& rhs) const
	{
		return (bool)stricmp(Name.c_str(), rhs.Name.c_str());
	}

	ArgumentInfo_t arg(int i){
		return Arguments.at(i);
	}

};

namespace std{

	template<typename T, typename D>
	bool contains(map<T,D> s, T key)
	{
		std::map<T,D>::iterator it = s.find(key);
		return (it != s.end());
	}
}

#ifdef _WIN64
#define REGISTER_SIZE 8
#else
#define REGISTER_SIZE 4
#endif