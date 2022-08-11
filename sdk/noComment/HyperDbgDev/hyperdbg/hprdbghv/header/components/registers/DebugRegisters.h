#pragma once
typedef enum _DEBUG_REGISTER_TYPE {
  BREAK_ON_INSTRUCTION_FETCH,
  BREAK_ON_WRITE_ONLY,
  BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED,
  BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH
} DEBUG_REGISTER_TYPE;
BOOLEAN
DebugRegistersSet(UINT32 DebugRegNum, DEBUG_REGISTER_TYPE ActionType,
                  BOOLEAN ApplyToVmcs, UINT64 TargetAddress);
