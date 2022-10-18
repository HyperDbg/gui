#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
std::map<std::string, REGS_ENUM> RegistersMap = {
    {"rax", REGISTER_RAX},       {"eax", REGISTER_EAX},
    {"ax", REGISTER_AX},         {"ah", REGISTER_AH},
    {"al", REGISTER_AL},         {"rbx", REGISTER_RBX},
    {"ebx", REGISTER_EBX},       {"bx", REGISTER_BX},
    {"bh", REGISTER_BH},         {"bl", REGISTER_BL},
    {"rcx", REGISTER_RCX},       {"ecx", REGISTER_ECX},
    {"cx", REGISTER_CX},         {"ch", REGISTER_CH},
    {"cl", REGISTER_CL},         {"rdx", REGISTER_RDX},
    {"edx", REGISTER_EDX},       {"dx", REGISTER_DX},
    {"dh", REGISTER_DH},         {"dl", REGISTER_DL},
    {"rsi", REGISTER_RSI},       {"esi", REGISTER_ESI},
    {"si", REGISTER_SI},         {"sil", REGISTER_SIL},
    {"rdi", REGISTER_RDI},       {"edi", REGISTER_EDI},
    {"di", REGISTER_DI},         {"dil", REGISTER_DIL},
    {"rbp", REGISTER_RBP},       {"ebp", REGISTER_EBP},
    {"bp", REGISTER_BP},         {"bpl", REGISTER_BPL},
    {"rsp", REGISTER_RSP},       {"esp", REGISTER_ESP},
    {"sp", REGISTER_SP},         {"spl", REGISTER_SPL},
    {"r8", REGISTER_R8},         {"r8d", REGISTER_R8D},
    {"r8w", REGISTER_R8W},       {"r8h", REGISTER_R8H},
    {"r8l", REGISTER_R8L},       {"r9", REGISTER_R9},
    {"r9d", REGISTER_R9D},       {"r9w", REGISTER_R9W},
    {"r9h", REGISTER_R9H},       {"r9l", REGISTER_R9L},
    {"r10", REGISTER_R10},       {"r10d", REGISTER_R10D},
    {"r10w", REGISTER_R10W},     {"r10h", REGISTER_R10H},
    {"r10l", REGISTER_R10L},     {"r11", REGISTER_R11},
    {"r11d", REGISTER_R11D},     {"r11w", REGISTER_R11W},
    {"r11h", REGISTER_R11H},     {"r11l", REGISTER_R11L},
    {"r12", REGISTER_R12},       {"r12d", REGISTER_R12D},
    {"r12w", REGISTER_R12W},     {"r12h", REGISTER_R12H},
    {"r12l", REGISTER_R12L},     {"r13", REGISTER_R13},
    {"r13d", REGISTER_R13D},     {"r13w", REGISTER_R13W},
    {"r13h", REGISTER_R13H},     {"r13l", REGISTER_R13L},
    {"r14", REGISTER_R14},       {"r14d", REGISTER_R14D},
    {"r14w", REGISTER_R14W},     {"r14h", REGISTER_R14H},
    {"r14l", REGISTER_R14L},     {"r15", REGISTER_R15},
    {"r15d", REGISTER_R15D},     {"r15w", REGISTER_R15W},
    {"r15h", REGISTER_R15H},     {"r15l", REGISTER_R15L},
    {"ds", REGISTER_DS},         {"es", REGISTER_ES},
    {"fs", REGISTER_FS},         {"gs", REGISTER_GS},
    {"cs", REGISTER_CS},         {"ss", REGISTER_SS},
    {"rflags", REGISTER_RFLAGS}, {"eflags", REGISTER_EFLAGS},
    {"flags", REGISTER_FLAGS},   {"cf", REGISTER_CF},
    {"pf", REGISTER_PF},         {"af", REGISTER_AF},
    {"zf", REGISTER_ZF},         {"sf", REGISTER_SF},
    {"tf", REGISTER_TF},         {"if", REGISTER_IF},
    {"df", REGISTER_DF},         {"of", REGISTER_OF},
    {"iopl", REGISTER_IOPL},     {"nt", REGISTER_NT},
    {"rf", REGISTER_RF},         {"vm", REGISTER_VM},
    {"ac", REGISTER_AC},         {"vif", REGISTER_VIF},
    {"vip", REGISTER_VIP},       {"id", REGISTER_ID},
    {"idtr", REGISTER_IDTR},     {"gdtr", REGISTER_GDTR},
    {"ldtr", REGISTER_LDTR},     {"tr", REGISTER_TR},
    {"cr0", REGISTER_CR0},       {"cr2", REGISTER_CR2},
    {"cr3", REGISTER_CR3},       {"cr4", REGISTER_CR4},
    {"cr8", REGISTER_CR8},       {"dr0", REGISTER_DR0},
    {"dr1", REGISTER_DR1},       {"dr2", REGISTER_DR2},
    {"dr3", REGISTER_DR3},       {"dr6", REGISTER_DR6},
    {"dr7", REGISTER_DR7},       {"rip", REGISTER_RIP},
    {"eip", REGISTER_EIP},       {"ip", REGISTER_IP},
};
VOID CommandRHelp() {
  ShowMessages("r : reads or modifies registers.\n\n");
  ShowMessages("syntax : \tr\n");
  ShowMessages("syntax : \tr [Register (string)] [= Expr (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : r\n");
  ShowMessages("\t\te.g : r @rax\n");
  ShowMessages("\t\te.g : r rax\n");
  ShowMessages("\t\te.g : r rax = 0x55\n");
  ShowMessages("\t\te.g : r rax = @rbx + @rcx + 0n10\n");
}

VOID ShowAllRegisters() {
  DEBUGGEE_REGISTER_READ_DESCRIPTION RegState = {0};
  RegState.RegisterID = DEBUGGEE_SHOW_ALL_REGISTERS;
  KdSendReadRegisterPacketToDebuggee(&RegState);
}

VOID CommandR(std::vector<std::string> SplittedCommand, std::string Command) {
  PVOID CodeBuffer;
  UINT64 BufferAddress;
  UINT32 BufferLength;
  UINT32 Pointer;
  REGS_ENUM RegKind;
  std::vector<std::string> Tmp;
  std::string SetRegValue = "SetRegValue";
  if (SplittedCommand[0] != "r") {
    return;
  }
  if (SplittedCommand.size() == 1) {
    if (g_IsSerialConnectedToRemoteDebuggee) {
      ShowAllRegisters();
    } else {
      ShowMessages("err, reading registers (r) is not valid in the current "
                   "context, you should connect to a debuggee\n");
    }
    return;
  }
  if (Command.find('=', 0) == string::npos) {
    Command.erase(0, 1);
    ReplaceAll(Command, "@", "");
    ReplaceAll(Command, " ", "");
    if (RegistersMap.find(Command) != RegistersMap.end()) {
      RegKind = RegistersMap[Command];
    } else {
      RegKind = (REGS_ENUM)-1;
    }
    if (RegKind != -1) {
      DEBUGGEE_REGISTER_READ_DESCRIPTION RegState = {0};
      RegState.RegisterID = RegKind;
      if (g_IsSerialConnectedToRemoteDebuggee) {
        KdSendReadRegisterPacketToDebuggee(&RegState);
      } else {
        ShowMessages("err, reading registers (r) is not valid in the current "
                     "context, you should connect to a debuggee\n");
      }
    } else {
      ShowMessages("err, invalid register\n");
    }
  }
  else if (Command.find('=', 0) != string::npos) {
    Command.erase(0, 1);
    Tmp = Split(Command, '=');
    if (Tmp.size() == 2) {
      ReplaceAll(Tmp[0], " ", "");
      string tmp = Tmp[0];
      if (RegistersMap.find(Tmp[0]) != RegistersMap.end()) {
        RegKind = RegistersMap[Tmp[0]];
      } else {
        ReplaceAll(tmp, "@", "");
        if (RegistersMap.find(tmp) != RegistersMap.end()) {
          RegKind = RegistersMap[tmp];
        } else {
          RegKind = (REGS_ENUM)-1;
        }
      }
      if (RegKind != -1) {
        SetRegValue = "@" + tmp + '=' + Tmp[1] + "; ";
        if (g_IsSerialConnectedToRemoteDebuggee) {
          CodeBuffer =
              ScriptEngineParseWrapper((char *)SetRegValue.c_str(), TRUE);
          if (CodeBuffer == NULL) {
            return;
          }
          BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
          BufferLength = ScriptEngineWrapperGetSize(CodeBuffer);
          Pointer = ScriptEngineWrapperGetPointer(CodeBuffer);
          KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer,
                                       FALSE);
          ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
        } else {
          ShowMessages("err, you're not connected to any debuggee\n");
        }
      } else {
        ShowMessages("err, invalid register\n");
      }
    }
  }
}

