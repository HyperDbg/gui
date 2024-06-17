#include "pch.h"
using namespace std;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID CommandPrintHelp(){
    ShowMessages("print : evaluates expressions.\n\n");
    ShowMessages("syntax : \tprint [Expression (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : print dq(poi(@rcx))\n");
}
VOID CommandPrint(vector<string> SplitCommand, string Command){
    PVOID  CodeBuffer;
    UINT64 BufferAddress;
    UINT32 BufferLength;
    UINT32 Pointer;
    if (SplitCommand.size() == 1){
        ShowMessages("incorrect use of the 'print'\n\n");
        CommandPrintHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, SplitCommand.at(0).size());
    Trim(Command);
    Command.insert(0, "print(");
    Command.append(");");
    if (g_IsSerialConnectedToRemoteDebuggee){
        CodeBuffer = ScriptEngineParseWrapper((char *)Command.c_str(), TRUE);
        if (CodeBuffer == NULL){
            return;
        }
        BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
        BufferLength  = ScriptEngineWrapperGetSize(CodeBuffer);
        Pointer       = ScriptEngineWrapperGetPointer(CodeBuffer);
        KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer, FALSE);
        ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
        ShowMessages("\n");
    }
    else{
        ShowMessages("err, you're not connected to any debuggee\n");
    }
}
