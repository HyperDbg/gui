#pragma once
BOOLEAN
ObjectShowProcessesOrThreadList(BOOLEAN                               IsProcess,
                                PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS SymDetailsForProcessList,
                                UINT64                                Eprocess,
                                PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS  SymDetailsForThreadList);
BOOLEAN
ObjectShowProcessesOrThreadDetails(BOOLEAN IsProcess);
