/* plugin:   (StaticAnalysis) for x64dbg <http://www.x64dbg.com>
* author:   tr4ceflow@gmail.com <http://blog.traceflow.com>
* license:  GLPv3
*/
#ifdef _WIN64


#include "S_IntermodularCalls.h"
#include "AnalysisRunner.h"



#define _isCall(disasm)  ((disasm.Instruction.BranchType==CallType) && (disasm.Instruction.BranchType!=RetType) && !(disasm.Argument1.ArgType &REGISTER_TYPE))


S_IntermodularCalls::S_IntermodularCalls(AnalysisRunner *parent) : ICommand(parent){

}

void S_IntermodularCalls::clear()
{
	numberOfApiCalls = 0;
	numberOfCalls = 0;
}

void S_IntermodularCalls::see(const  Instruction_t* currentInstruction, const StackEmulator* stackState, const RegisterEmulator* regState)
{

	if ((_isCall(currentInstruction->BeaStruct))){
		// current instructions contains a call
		// extract from "call 0x123" --> instruction at 0x123


		// the opcode 0xFF "jmp" tells us that the current call is a call to a dll-function

		numberOfCalls++;
		// does the TitanEngine provides us a label?
		char labelText[MAX_LABEL_SIZE];
		bool hasLabel = DbgGetLabelAt(currentInstruction->BeaStruct.Instruction.AddrValue, SEG_DEFAULT, labelText);
		if (hasLabel){

			// we have a label from TitanEngine --> look up function header in database
			FunctionInfo_t f = mParent->FunctionInformation()->find(labelText);
			if (!f.invalid){
				numberOfApiCalls++;
				// yeah we know everything about the dll-call!
				std::string functionComment;
				functionComment = f.ReturnType + " " + f.Name + "(...)";
				DbgSetCommentAt(currentInstruction->BeaStruct.VirtualAddr, functionComment.c_str());

				if (f.Arguments.size() > 0){
					std::string ArgComment = f.arg(0).Type + " " + f.arg(0).Name;
					DbgSetCommentAt(regState->rcx(), ArgComment.c_str());
				}
				if (f.Arguments.size() > 1){
					std::string ArgComment = f.arg(1).Type + " " + f.arg(1).Name;
					DbgSetCommentAt(regState->rdx(), ArgComment.c_str());
				}
				if (f.Arguments.size() > 2){
					std::string ArgComment = f.arg(2).Type + " " + f.arg(2).Name;
					DbgSetCommentAt(regState->r8(), ArgComment.c_str());
				}
				if (f.Arguments.size() > 3){
					std::string ArgComment = f.arg(3).Type + " " + f.arg(3).Name;
					DbgSetCommentAt(regState->r9(), ArgComment.c_str());
				}
				if (f.Arguments.size() > 4){
					// set comments for the arguments
					for (auto i = 4; i < f.Arguments.size(); i++)
					{
						std::string ArgComment = f.arg(i).Type + " " + f.arg(i).Name;
						duint commentAddr = stackState->lastAccessAtOffset(f.Arguments.size() - i - 1);
						if (commentAddr != STACK_ERROR){
							DbgSetCommentAt(commentAddr, ArgComment.c_str());
						}
						else{
							// we have more arguments in the function descriptions than parameters on the stack
							break;
						}
					}
				}

			}


		}
	}
}

bool S_IntermodularCalls::think()
{
	StackEmulator stack;

	_plugin_logprintf("[StaticAnalysis:IntermodularCalls] found %i calls\n", numberOfCalls);
	_plugin_logprintf("[StaticAnalysis:IntermodularCalls] of which are %i intermodular calls\n", numberOfApiCalls);

	return true;
}

void S_IntermodularCalls::unknownOpCode(const  DISASM* disasm)
{
	// current instruction wasn't correctly disassembled, so assuming worst-case

}



#endif // _WIN64


/* new calling convention in x64

1. arg -> RCX (floating point: XMM0)
2. arg -> RDX (floating point: XMM1)
3. arg -> R8  (floating point: XMM2)
4. arg -> R9  (floating point: XMM3)

additional arguments are pushed onto the stack (right to left)



if (f.Arguments.size() > 0){
std::string ArgComment = f.arg(0).Type + " " + f.arg(0).Name;
DbgSetAutoCommentAt(regState->rcx(), ArgComment.c_str());
}
if (f.Arguments.size() > 1){
std::string ArgComment = f.arg(1).Type + " " + f.arg(1).Name;
DbgSetAutoCommentAt(regState->rdx(), ArgComment.c_str());
}
if (f.Arguments.size() > 2){
std::string ArgComment = f.arg(2).Type + " " + f.arg(2).Name;
DbgSetAutoCommentAt(regState->r8(), ArgComment.c_str());
}
if (f.Arguments.size() > 3){
std::string ArgComment = f.arg(3).Type + " " + f.arg(3).Name;
DbgSetAutoCommentAt(regState->r9(), ArgComment.c_str());
}
if (f.Arguments.size() > 4){
// set comments for the arguments
for (auto i = 4; i < f.Arguments.size(); i++)
{
std::string ArgComment = f.arg(i).Type + " " + f.arg(i).Name;
duint commentAddr = stackState->lastAccessAtOffset(i - 4);
if (commentAddr != STACK_ERROR){
DbgSetAutoCommentAt(commentAddr, ArgComment.c_str());
}
else{
// we have more arguments in the function descriptions than parameters on the stack
break;
}
}
}*/