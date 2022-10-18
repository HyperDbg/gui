#pragma once
typedef enum _NMI_BROADCAST_ACTION_TYPE {
  NMI_BROADCAST_ACTION_NONE = 0,
  NMI_BROADCAST_ACTION_TEST,
  NMI_BROADCAST_ACTION_KD_HALT_CORE,
} NMI_BROADCAST_ACTION_TYPE;
BOOLEAN
VmxBroadcastHandleNmiCallback(PVOID Context, BOOLEAN Handled);
BOOLEAN
VmxBroadcastNmi(UINT32 CurrentCoreIndex,
                NMI_BROADCAST_ACTION_TYPE VmxBroadcastAction);
BOOLEAN
VmxBroadcastNmiHandler(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs,
                       BOOLEAN IsOnVmxNmiHandler);
