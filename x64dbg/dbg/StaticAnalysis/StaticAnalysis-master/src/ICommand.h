/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#pragma once
#include "pluginsdk/bridgemain.h"
#include "pluginmain.h"
#include "pluginsdk/BeaEngine.h"

#include "meta.h"

class AnalysisRunner;
class StackEmulator;
class RegisterEmulator;

class ICommand
{
public:
	ICommand(AnalysisRunner *parent);
	virtual ~ICommand(void);


protected:
	duint mBase;
	duint mSize;
	AnalysisRunner *mParent;

public:

	// initialization before any analysis
	void initialise(const duint Base,const duint Size);
	// clear all extracted informations
	virtual void clear() = 0;
	// each sub-plugin will get a simulated flow (it gets an instruction and the state of the stack)
	virtual void see(const Instruction_t* disasm, const StackEmulator *stack, const RegisterEmulator* regState) = 0;
	// this methods process all gathered informations
	virtual bool think() = 0;

	void unknownOpCode(const DISASM* disasm);
};

