/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#pragma once
#include "ICommand.h"
#include <list>
#include <set>
#include <map>

#include "meta.h"
#include "StackEmulator.h"
#include "RegisterEmulator.h"

class S_IntermodularCalls : public ICommand
{
	unsigned int numberOfCalls;
	unsigned int numberOfApiCalls;

public:

	S_IntermodularCalls(AnalysisRunner *parent);

	void clear();
	void see(const Instruction_t* disasm, const StackEmulator* stack, const RegisterEmulator* regState);
	bool think();
	void unknownOpCode(const DISASM* disasm);
};

