#pragma once
#define TEST_PROCESS_NAME "hyperdbg-test.exe"
BOOLEAN CreateProcessAndOpenPipeConnection(PHANDLE ConnectionPipeHandle,
                                   PHANDLE ThreadHandle,
                                   PHANDLE ProcessHandle);
VOID CloseProcessAndClosePipeConnection(HANDLE ConnectionPipeHandle,
                                   HANDLE ThreadHandle,
                                   HANDLE ProcessHandle);
