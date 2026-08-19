/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
 * author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
 * license:  GLPv3
 */
#pragma once
#include "pluginsdk/bridgemain.h"
#include "pluginmain.h"
#include "pluginsdk/BeaEngine.h"

#include "S_IntermodularCalls.h"

#include "ApiDB.h"

class AnalysisRunner
{
	std::map<UInt64,Instruction_t> mInstructionsBuffer;
	duint mBaseAddress;
	duint mSize;

	S_IntermodularCalls *_Calls;
	ApiDB *mApiDb;


	unsigned char* mCodeMemory;
	UIntPtr currentEIP;
	UInt64 currentVirtualAddr;

protected:
	// forwarding
	void see(const Instruction_t *disasm, const  StackEmulator *stack, const RegisterEmulator* regState);
	void clear();
	void think();
	void initialise();
	void unknownOpCode(const DISASM *disasm);

private:
	void run();
	void publishInstructions();

public:

	ApiDB* FunctionInformation() const;
	void setFunctionInformation(ApiDB *api);

	AnalysisRunner(duint BaseAddress,duint Size);
	~AnalysisRunner(void);
	
	void start();
	


	int instruction(UInt64 va, Instruction_t* instr) const;

	
};

