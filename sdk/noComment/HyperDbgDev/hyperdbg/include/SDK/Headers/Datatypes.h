#pragma once
typedef int (*Callback)(const char *Text);
typedef struct _DEBUGGEE_USER_INPUT_PACKET {
  UINT32 CommandLen;
  BOOLEAN IgnoreFinishedSignal;
  UINT32 Result;
} DEBUGGEE_USER_INPUT_PACKET, *PDEBUGGEE_USER_INPUT_PACKET;
typedef struct _DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET {
  UINT32 Length;
} DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET,
    *PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET;
#define SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED                                  \
  sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED)
typedef struct _DEBUGGER_PAUSE_PACKET_RECEIVED {
  UINT32 Result; 
} DEBUGGER_PAUSE_PACKET_RECEIVED, *PDEBUGGER_PAUSE_PACKET_RECEIVED;
typedef struct _DEBUGGEE_KD_PAUSED_PACKET {
  UINT64 Rip;
  BOOLEAN Is32BitAddress; 
  DEBUGGEE_PAUSING_REASON PausingReason;
  ULONG CurrentCore;
  UINT64 EventTag;
  UINT64 Rflags;
  BYTE InstructionBytesOnRip[MAXIMUM_INSTR_SIZE];
  UINT16 ReadInstructionLen;
} DEBUGGEE_KD_PAUSED_PACKET, *PDEBUGGEE_KD_PAUSED_PACKET;
typedef struct _DEBUGGEE_UD_PAUSED_PACKET {
  UINT64 Rip;
  UINT64 ProcessDebuggingToken;
  BOOLEAN Is32Bit; 
  DEBUGGEE_PAUSING_REASON PausingReason;
  UINT32 ProcessId;
  UINT32 ThreadId;
  UINT64 EventTag;
  UINT64 Rflags;
  BYTE InstructionBytesOnRip[MAXIMUM_INSTR_SIZE];
  UINT16 ReadInstructionLen;
  GUEST_REGS GuestRegs;
} DEBUGGEE_UD_PAUSED_PACKET, *PDEBUGGEE_UD_PAUSED_PACKET;
static_assert(sizeof(DEBUGGEE_UD_PAUSED_PACKET) < PacketChunkSize,
              "err (static_assert), size of PacketChunkSize should be bigger "
              "than DEBUGGEE_UD_PAUSED_PACKET");
typedef struct _DEBUGGEE_MESSAGE_PACKET {
  UINT32 OperationCode;
  CHAR Message[PacketChunkSize];
} DEBUGGEE_MESSAGE_PACKET, *PDEBUGGEE_MESSAGE_PACKET;
