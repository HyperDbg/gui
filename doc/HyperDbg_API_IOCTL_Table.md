# HyperDbg API 和 IOCTL 对照表

## 概述

本文档列出了 HyperDbg 中内核模式和用户模式的所有 API、IOCTL 代码及其源代码位置。

## IOCTL 定义位置

| 项目            | 位置                                                                                                                      |
| ------------- | ----------------------------------------------------------------------------------------------------------------------- |
| IOCTL 定义头文件   | [Ioctls.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\include\SDK\headers\Ioctls.h) |
| 内核模式 IOCTL 处理 | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c)   |

## 内核模式 API 和 IOCTL 对照表

| IOCTL 代码 | IOCTL 名称                                                           | 十六进制值    | 内核处理函数                                                              | 源代码位置                                                                                                                             |
| -------- | ------------------------------------------------------------------ | -------- | ------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------- |
| 0x800    | IOCTL\_REGISTER\_EVENT                                             | 0x222003 | LogRegisterIrpBasedNotification / LogRegisterEventBasedNotification | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L80-L120)    |
| 0x801    | IOCTL\_RETURN\_IRP\_PENDING\_PACKETS\_AND\_DISALLOW\_IOCTL         | 0x222007 | LogCallbackSendBuffer                                               | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L122-L135)   |
| 0x802    | IOCTL\_TERMINATE\_VMX                                              | 0x22200B | DebuggerUninitialize / VmFuncUninitVmm                              | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L137-L148)   |
| 0x803    | IOCTL\_DEBUGGER\_READ\_MEMORY                                      | 0x22200F | DebuggerCommandReadMemory                                           | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L150-L195)   |
| 0x804    | IOCTL\_DEBUGGER\_READ\_OR\_WRITE\_MSR                              | 0x222013 | DebuggerReadOrWriteMsr                                              | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L197-L245)   |
| 0x805    | IOCTL\_DEBUGGER\_READ\_PAGE\_TABLE\_ENTRIES\_DETAILS               | 0x222017 | ExtensionCommandPte                                                 | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L247-L280)   |
| 0x806    | IOCTL\_DEBUGGER\_REGISTER\_EVENT                                   | 0x22201B | DebuggerParseEvent                                                  | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L282-L325)   |
| 0x807    | IOCTL\_DEBUGGER\_ADD\_ACTION\_TO\_EVENT                            | 0x22201F | DebuggerParseAction                                                 | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L327-L370)   |
| 0x808    | IOCTL\_DEBUGGER\_HIDE\_AND\_UNHIDE\_TO\_TRANSPARENT\_THE\_DEBUGGER | 0x222023 | TransparentHideDebuggerWrapper / TransparentUnhideDebuggerWrapper   | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L372-L415)   |
| 0x809    | IOCTL\_DEBUGGER\_VA2PA\_AND\_PA2VA\_COMMANDS                       | 0x222027 | ExtensionCommandVa2paAndPa2va                                       | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L417-L455)   |
| 0x80A    | IOCTL\_DEBUGGER\_EDIT\_MEMORY                                      | 0x22202B | DebuggerCommandEditMemory                                           | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L457-L510)   |
| 0x80B    | IOCTL\_DEBUGGER\_SEARCH\_MEMORY                                    | 0x22202F | DebuggerCommandSearchMemory                                         | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L512-L575)   |
| 0x80C    | IOCTL\_DEBUGGER\_MODIFY\_EVENTS                                    | 0x222033 | DebuggerParseEventsModification                                     | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L577-L620)   |
| 0x80D    | IOCTL\_DEBUGGER\_FLUSH\_LOGGING\_BUFFERS                           | 0x222037 | DebuggerCommandFlush                                                | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L622-L665)   |
| 0x80E    | IOCTL\_DEBUGGER\_ATTACH\_DETACH\_USER\_MODE\_PROCESS               | 0x22203B | AttachingTargetProcess                                              | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L667-L710)   |
| 0x80F    | IOCTL\_DEBUGGER\_PRINT                                             | 0x22203F | (已弃用)                                                               | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c)             |
| 0x810    | IOCTL\_PREPARE\_DEBUGGEE                                           | 0x222043 | SerialConnectionPrepare                                             | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L712-L755)   |
| 0x811    | IOCTL\_PAUSE\_PACKET\_RECEIVED                                     | 0x222047 | KdHaltSystem                                                        | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L757-L800)   |
| 0x812    | IOCTL\_SEND\_SIGNAL\_EXECUTION\_IN\_DEBUGGEE\_FINISHED             | 0x22204B | DebuggerCommandSignalExecutionState                                 | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L802-L845)   |
| 0x813    | IOCTL\_SEND\_USERMODE\_MESSAGES\_TO\_DEBUGGER                      | 0x22204F | DebuggerCommandSendMessage                                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L847-L895)   |
| 0x814    | IOCTL\_SEND\_GENERAL\_BUFFER\_FROM\_DEBUGGEE\_TO\_DEBUGGER         | 0x222053 | DebuggerCommandSendGeneralBufferToDebugger                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L897-L945)   |
| 0x815    | IOCTL\_PERFORM\_KERNEL\_SIDE\_TESTS                                | 0x222057 | TestKernelPerformTests                                              | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L947-L990)   |
| 0x816    | IOCTL\_RESERVE\_PRE\_ALLOCATED\_POOLS                              | 0x22205B | DebuggerCommandReservePreallocatedPools                             | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1000-L1045) |
| 0x817    | IOCTL\_SEND\_USER\_DEBUGGER\_COMMANDS                              | 0x22205F | UdDispatchUsermodeCommands                                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1250-L1295) |
| 0x818    | IOCTL\_GET\_DETAIL\_OF\_ACTIVE\_THREADS\_AND\_PROCESSES            | 0x222063 | AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses           | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1297-L1325) |
| 0x819    | IOCTL\_GET\_USER\_MODE\_MODULE\_DETAILS                            | 0x222067 | UserAccessGetLoadedModules                                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1327-L1375) |
| 0x81A    | IOCTL\_QUERY\_COUNT\_OF\_ACTIVE\_PROCESSES\_OR\_THREADS            | 0x22206B | ProcessQueryCount / ThreadQueryCount                                | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1377-L1435) |
| 0x81B    | IOCTL\_GET\_LIST\_OF\_THREADS\_AND\_PROCESSES                      | 0x22206F | ProcessQueryList / ThreadQueryList                                  | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1437-L1495) |
| 0x81C    | IOCTL\_QUERY\_CURRENT\_PROCESS                                     | 0x222073 | ProcessQueryDetails                                                 | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1520-L1565) |
| 0x81D    | IOCTL\_QUERY\_CURRENT\_THREAD                                      | 0x222077 | ThreadQueryDetails                                                  | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1497-L1518) |
| 0x81E    | IOCTL\_REQUEST\_REV\_MACHINE\_SERVICE                              | 0x22207B | ConfigureInitializeExecTrapOnAllProcessors                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1567-L1610) |
| 0x81F    | IOCTL\_DEBUGGER\_BRING\_PAGES\_IN                                  | 0x22207F | DebuggerCommandBringPagein                                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1612-L1655) |
| 0x820    | IOCTL\_PREACTIVATE\_FUNCTIONALITY                                  | 0x222083 | DebuggerCommandPreactivateFunctionality                             | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1047-L1092) |
| 0x821    | IOCTL\_PCIE\_ENDPOINT\_ENUM                                        | 0x222087 | ExtensionCommandPcitree                                             | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1657-L1700) |
| 0x822    | IOCTL\_PERFORM\_ACTIONS\_ON\_APIC                                  | 0x22208B | ExtensionCommandPerformActionsForApicRequests                       | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1094-L1137) |
| 0x823    | IOCTL\_PCIDEVINFO\_ENUM                                            | 0x22208F | ExtensionCommandPcidevinfo                                          | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1702-L1745) |
| 0x824    | IOCTL\_QUERY\_IDT\_ENTRY                                           | 0x222093 | ExtensionCommandPerformQueryIdtEntriesRequest                       | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1139-L1182) |
| 0x825    | IOCTL\_SET\_BREAKPOINT\_USER\_DEBUGGER                             | 0x222097 | BreakpointAddNew                                                    | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1184-L1227) |
| 0x826    | IOCTL\_PERFORM\_SMI\_OPERATION                                     | 0x22209B | VmFuncSmmPerformSmiOperation                                        | [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c#L1229-L1272) |

## 用户模式 API 和 IOCTL 对照表

### 调试命令 (Debugging Commands)

| 用户命令          | IOCTL 代码                                         | 用户模式函数                     | 通信方式         | 源代码位置                                                                                                                                                                   |
| ------------- | ------------------------------------------------ | ------------------------------ | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `a`           | 无 (本地命令)                                      | CommandAssemble                 | 本地处理         | [a.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\a.cpp)                       |
| `bc`          | 无 (串行通信)                                      | CommandBc                       | 串行数据包         | [bc.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\bc.cpp)                     |
| `bd`          | 无 (串行通信)                                      | CommandBd                       | 串行数据包         | [bd.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\bd.cpp)                     |
| `be`          | 无 (串行通信)                                      | CommandBe                       | 串行数据包         | [be.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\be.cpp)                     |
| `bl`          | 无 (串行通信)                                      | CommandBl                       | 串行数据包         | [bl.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\bl.cpp)                     |
| `bp`          | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)         | CommandBp                       | IOCTL         | [bp.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\bp.cpp)                   |
| `continue`    | 无 (串行通信)                                      | CommandContinue                 | 串行数据包         | [continue.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\continue.cpp)         |
| `core`        | 无 (本地命令)                                      | CommandCore                     | 本地处理         | [core.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\core.cpp)                 |
| `cpu`         | 无 (本地命令)                                      | CommandCpu                      | 本地处理         | [cpu.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\cpu.cpp)                   |
| `e`           | IOCTL\_DEBUGGER\_EDIT\_MEMORY (0x80A)            | CommandE                       | IOCTL         | [e.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\e.cpp)                     |
| `eval`        | 无 (本地命令)                                      | CommandEval                     | 本地处理         | [eval.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\eval.cpp)                 |
| `events`      | IOCTL\_DEBUGGER\_MODIFY\_EVENTS (0x80C)          | CommandEvents                   | IOCTL         | [events.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\events.cpp)           |
| `exit`        | 无 (本地命令)                                      | CommandExit                     | 本地处理         | [exit.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\exit.cpp)                 |
| `flush`       | IOCTL\_DEBUGGER\_FLUSH\_LOGGING\_BUFFERS (0x80D) | CommandFlush                   | IOCTL         | [flush.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\flush.cpp)             |
| `gg`          | 无 (本地命令)                                      | CommandGg                       | 本地处理         | [gg.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\gg.cpp)                     |
| `gu`          | 无 (串行通信)                                      | CommandGu                       | 串行数据包         | [gu.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\gu.cpp)                     |
| `lm`          | IOCTL\_GET\_USER\_MODE\_MODULE\_DETAILS (0x819)  | CommandLm                       | IOCTL         | [lm.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\lm.cpp)                   |
| `load`        | 无 (本地命令)                                      | CommandLoad                     | 本地处理         | [load.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\load.cpp)                 |
| `output`      | 无 (本地命令)                                      | CommandOutput                   | 本地处理         | [output.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\output.cpp)             |
| `pause`       | 无 (串行通信)                                      | CommandPause                    | 串行数据包         | [pause.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\pause.cpp)               |
| `prealloc`    | IOCTL\_RESERVE\_PRE\_ALLOCATED\_POOLS (0x816)    | CommandPrealloc                  | IOCTL         | [prealloc.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\prealloc.cpp)       |
| `preactivate` | IOCTL\_PREACTIVATE\_FUNCTIONALITY (0x820)        | CommandPreactivate               | IOCTL         | [preactivate.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\preactivate.cpp) |
| `print`       | 无 (本地命令)                                      | CommandPrint                    | 本地处理         | [print.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\print.cpp)               |
| `rdmsr`       | IOCTL\_DEBUGGER\_READ\_OR\_WRITE\_MSR (0x804)    | CommandRdmsr                    | IOCTL         | [rdmsr.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\rdmsr.cpp)             |
| `s`           | IOCTL\_DEBUGGER\_SEARCH\_MEMORY (0x80B)          | CommandS                        | IOCTL         | [s.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\s.cpp)                     |
| `settings`    | 无 (本地命令)                                      | CommandSettings                 | 本地处理         | [settings.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\settings.cpp)           |
| `sleep`       | 无 (本地命令)                                      | CommandSleep                    | 本地处理         | [sleep.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\sleep.cpp)               |
| `test`        | IOCTL\_PERFORM\_KERNEL\_SIDE\_TESTS (0x815)      | CommandTest                     | IOCTL         | [test.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\test.cpp)               |
| `unload`      | 无 (本地命令)                                      | CommandUnload                   | 本地处理         | [unload.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\unload.cpp)             |
| `wrmsr`       | IOCTL\_DEBUGGER\_READ\_OR\_WRITE\_MSR (0x804)    | CommandWrmsr                    | IOCTL         | [wrmsr.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\debugging-commands\wrmsr.cpp)             |

### 扩展命令 (Extension Commands)

| 用户命令         | IOCTL 代码                                                                   | 用户模式函数                     | 通信方式         | 源代码位置                                                                                                                                                                         |
| ------------ | -------------------------------------------------------------------------- | ------------------------------ | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `!msrread`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandMsrread                  | IOCTL         | [msrread.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\msrread.cpp)               |
| `!msrwrite`  | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandMsrwrite                 | IOCTL         | [msrwrite.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\msrwrite.cpp)             |
| `!apic`      | IOCTL\_PERFORM\_ACTIONS\_ON\_APIC (0x822)                                  | CommandApic                     | IOCTL         | [apic.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\apic.cpp)                     |
| `!cpuid`     | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandCpuid                    | IOCTL         | [cpuid.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\cpuid.cpp)                   |
| `!crwrite`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandCrwrite                  | IOCTL         | [crwrite.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\crwrite.cpp)               |
| `!dr`        | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandDr                       | IOCTL         | [dr.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\dr.cpp)                         |
| `!epthook`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandEpthook                  | IOCTL         | [epthook.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\epthook.cpp)               |
| `!epthook2`  | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandEpthook2                 | IOCTL         | [epthook2.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\epthook2.cpp)             |
| `!exception` | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandException                | IOCTL         | [exception.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\exception.cpp)           |
| `!hide`      | IOCTL\_DEBUGGER\_HIDE\_AND\_UNHIDE\_TO\_TRANSPARENT\_THE\_DEBUGGER (0x808) | CommandHide                     | IOCTL         | [hide.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\hide.cpp)                     |
| `!idt`       | IOCTL\_QUERY\_IDT\_ENTRY (0x824)                                           | CommandIdt                      | IOCTL         | [idt.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\idt.cpp)                       |
| `!interrupt` | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandInterrupt                | IOCTL         | [interrupt.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\interrupt.cpp)           |
| `!ioin`      | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandIoin                     | IOCTL         | [ioin.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\ioin.cpp)                     |
| `!ioout`     | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandIoout                    | IOCTL         | [ioout.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\ioout.cpp)                   |
| `!ioapic`    | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandIoapic                   | IOCTL         | [ioapic.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\ioapic.cpp)               |
| `!measure`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandMeasure                 | IOCTL         | [measure.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\measure.cpp)             |
| `!mode`      | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandMode                     | IOCTL         | [mode.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\mode.cpp)                     |
| `!monitor`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandMonitor                  | IOCTL         | [monitor.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\monitor.cpp)               |
| `!pa2va`     | IOCTL\_DEBUGGER\_VA2PA\_AND\_PA2VA\_COMMANDS (0x809)                       | CommandPa2va                    | IOCTL         | [pa2va.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\pa2va.cpp)                   |
| `!va2pa`     | IOCTL\_DEBUGGER\_VA2PA\_AND\_PA2VA\_COMMANDS (0x809)                       | CommandVa2pa                    | IOCTL         | [va2pa.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\va2pa.cpp)                   |
| `!pcicam`    | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandPcicam                   | IOCTL         | [pcicam.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\pcicam.cpp)                 |
| `!pcitree`   | IOCTL\_PCIE\_ENDPOINT\_ENUM (0x821)                                        | CommandPcitree                  | IOCTL         | [pcitree.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\pcitree.cpp)               |
| `!pmc`       | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandPmc                      | IOCTL         | [pmc.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\pmc.cpp)                       |
| `!pte`       | IOCTL\_DEBUGGER\_READ\_PAGE\_TABLE\_ENTRIES\_DETAILS (0x805)               | CommandPte                      | IOCTL         | [pte.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\pte.cpp)                       |
| `!rev`       | IOCTL\_REQUEST\_REV\_MACHINE\_SERVICE (0x81E)                               | CommandRev                      | IOCTL         | [rev.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\rev.cpp)                       |
| `!smi`       | IOCTL\_PERFORM\_SMI\_OPERATION (0x826)                                     | CommandSmi                      | IOCTL         | [smi.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\smi.cpp)                       |
| `!syscall`   | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandSyscall                  | IOCTL         | [syscall-sysret.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\syscall-sysret.cpp) |
| `!trace`     | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandTrace                    | IOCTL         | [trace.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\trace.cpp)                   |
| `!track`     | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandTrack                    | IOCTL         | [track.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\track.cpp)                 |
| `!tsc`       | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandTsc                      | IOCTL         | [tsc.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\tsc.cpp)                       |
| `!unhide`    | IOCTL\_DEBUGGER\_HIDE\_AND\_UNHIDE\_TO\_TRANSPARENT\_THE\_DEBUGGER (0x808) | CommandUnhide                   | IOCTL         | [unhide.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\unhide.cpp)                 |
| `!vmcall`    | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandVmcall                   | IOCTL         | [vmcall.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\vmcall.cpp)                 |
| `!xsetbv`    | IOCTL\_DEBUGGER\_REGISTER\_EVENT (0x806)                                   | CommandXsetbv                   | IOCTL         | [xsetbv.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\extension-commands\xsetbv.cpp)                 |

### 元命令 (Meta Commands)

| 用户命令       | IOCTL 代码                                  | 用户模式函数        | 通信方式         | 源代码位置                                                                                                                                                       |
| ---------- | ----------------------------------------- | ------------- | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `!attach` | 无 (本地命令)                                  | CommandAttach   | 本地处理         | [attach.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\attach.cpp)       |
| `cls`      | 无 (本地命令)                                  | CommandCls      | 本地处理         | [cls.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\cls.cpp)             |
| `connect`  | 无 (本地命令)                                  | CommandConnect  | 本地处理         | [connect.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\connect.cpp)     |
| `debug`    | IOCTL\_DEBUGGER\_ATTACH\_DETACH\_USER\_MODE\_PROCESS (0x80E) | CommandDebug    | IOCTL         | [debug.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\debug.cpp)         |
| `detach`   | IOCTL\_DEBUGGER\_ATTACH\_DETACH\_USER\_MODE\_PROCESS (0x80E) | CommandDetach   | IOCTL         | [detach.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\detach.cpp)       |
| `disconnect` | 无 (本地命令)                                  | CommandDisconnect | 本地处理         | [disconnect.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\disconnect.cpp) |
| `dump`     | 无 (本地命令)                                  | CommandDump     | 本地处理         | [dump.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\dump.cpp)           |
| `formats`  | 无 (本地命令)                                  | CommandFormats  | 本地处理         | [formats.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\formats.cpp)     |
| `kill`     | 无 (本地命令)                                  | CommandKill     | 本地处理         | [kill.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\kill.cpp)           |
| `listen`   | 无 (本地命令)                                  | CommandListen   | 本地处理         | [listen.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\listen.cpp)       |
| `logclose` | 无 (本地命令)                                  | CommandLogclose | 本地处理         | [logclose.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\logclose.cpp)   |
| `logopen`  | 无 (本地命令)                                  | CommandLogopen  | 本地处理         | [logopen.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\logopen.cpp)     |
| `!pagein` | IOCTL\_DEBUGGER\_BRING\_PAGES\_IN (0x81F) | CommandPagein   | IOCTL         | [pagein.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\pagein.cpp)     |
| `pe`       | 无 (本地命令)                                  | CommandPe       | 本地处理         | [pe.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\pe.cpp)               |
| `process`  | 无 (本地命令)                                  | CommandProcess  | 本地处理         | [process.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\process.cpp)     |
| `restart`  | 无 (本地命令)                                  | CommandRestart  | 本地处理         | [restart.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\restart.cpp)     |
| `script`   | 无 (本地命令)                                  | CommandScript   | 本地处理         | [script.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\script.cpp)       |
| `start`    | 无 (本地命令)                                  | CommandStart    | 本地处理         | [start.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\start.cpp)         |
| `status`   | 无 (本地命令)                                  | CommandStatus   | 本地处理         | [status.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\status.cpp)       |
| `switch`   | 无 (本地命令)                                  | CommandSwitch   | 本地处理         | [switch.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\switch.cpp)       |
| `sym`      | 无 (本地命令)                                  | CommandSym      | 本地处理         | [sym.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\sym.cpp)             |
| `sympath`  | 无 (本地命令)                                  | CommandSympath  | 本地处理         | [sympath.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\sympath.cpp)       |
| `thread`   | 无 (本地命令)                                  | CommandThread   | 本地处理         | [thread.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\commands\meta-commands\thread.cpp)       |

### 核心调试器功能

| 功能      | IOCTL 代码                                      | 源代码位置                                                                                                                                           |
| ------- | --------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| 调试器初始化  | 多个 IOCTL                                      | [debugger.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\core\debugger.cpp)      |
| 内核级调试   | IOCTL\_PREPARE\_DEBUGGEE (0x810)              | [kd.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\kernel-level\kd.cpp)          |
| 用户级调试   | IOCTL\_SEND\_USER\_DEBUGGER\_COMMANDS (0x817) | [ud.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\user-level\ud.cpp)            |
| 内存读取    | IOCTL\_DEBUGGER\_READ\_MEMORY (0x803)         | [readmem.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\misc\readmem.cpp)        |
| 符号处理    | 多个 IOCTL                                      | [symbol.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\debugger\script-engine\symbol.cpp) |
| 对象管理    | 多个 IOCTL                                      | [objects.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\objects\objects.cpp)              |
| 反转机器服务  | IOCTL\_REQUEST\_REV\_MACHINE\_SERVICE (0x81E) | [rev-ctrl.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\rev\rev-ctrl.cpp)                |
| 主应用程序入口 | 多个 IOCTL                                      | [libhyperdbg.cpp](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\libhyperdbg\code\app\libhyperdbg.cpp)          |

## IOCTL 代码计算公式

所有 IOCTL 代码使用以下公式计算：

    IOCTL_CODE = ((FILE_DEVICE_UNKNOWN) << 16) | ((FILE_ANY_ACCESS) << 14) | ((FunctionCode) << 2) | (METHOD_BUFFERED)

其中：

*   `FILE_DEVICE_UNKNOWN` \= 0x00000022

*   `FILE_ANY_ACCESS` \= 0

*   `METHOD_BUFFERED` \= 0

*   `FunctionCode` \= 0x800 到 0x826

示例：

    IOCTL_DEBUGGER_READ_MEMORY = (0x22 << 16) | (0 << 14) | (0x803 << 2) | 0
                               = 0x220000 | 0x000000 | 0x200C | 0x0000
                               = 0x22200C

## 相关头文件

| 头文件           | 位置                                                                                                                                                           | 说明       |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ | -------- |
| Ioctls.h      | [hyperdbg/include/SDK/headers/Ioctls.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\include\SDK\headers\Ioctls.h)         | IOCTL 定义 |
| HyperDbgSdk.h | [hyperdbg/include/SDK/HyperDbgSdk.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\include\SDK\HyperDbgSdk.h)               | SDK 主头文件 |
| ErrorCodes.h  | [hyperdbg/include/SDK/headers/ErrorCodes.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\include\SDK\headers\ErrorCodes.h) | 错误代码定义   |

## 总结

HyperDbg 使用 IOCTL 机制在用户模式和内核模式之间通信。共有 39 个 IOCTL 代码（0x800 到 0x826），涵盖了以下功能类别：

1.  **事件和断点管理** - 注册、修改、添加事件和断点

2.  **内存操作** - 读取、编辑、搜索内存

3.  **寄存器访问** - MSR、CR、DR 等寄存器读写

4.  **虚拟化功能** - EPT Hook、VMCALL、透明模式

5.  **进程和线程管理** - 附加/分离、查询信息

6.  **硬件监控** - IDT、APIC、PCI、SMI

7.  **调试器控制** - 暂停、继续、准备调试

8.  **日志和缓冲区** - 刷新、预分配、消息传递

9.  **用户调试器** - 断点、命令调度

10. **反转机器** - 内存重构服务

每个 IOCTL 都在内核模式的 [Ioctl.c](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\code\driver\Ioctl.c) 中处理，并通过用户模式的各种命令文件调用。

## 补充：内核模式辅助函数（未直接映射到 IOCTL）

以下内核函数在 SDK 头文件中定义，用于支持各种调试功能，但不直接通过 IOCTL 调用：

### 事件应用函数 (ApplyEvents.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ApplyEventMonitorEvent | 应用内存监控事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventEptHookExecCcEvent | 应用 EPT Hook CC 事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventEpthookInlineEvent | 应用 EPT Hook 内联事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventTrapModeChangeEvent | 应用 Trap 模式改变事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventRdmsrExecutionEvent | 应用 RDMSR 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventWrmsrExecutionEvent | 应用 WRMSR 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventInOutExecutionEvent | 应用 IN/OUT 指令事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventTscExecutionEvent | 应用 TSC 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventRdpmcExecutionEvent | 应用 RDPMC 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventMov2DebugRegExecutionEvent | 应用 MOV 到调试寄存器事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventControlRegisterAccessedEvent | 应用控制寄存器访问事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventExceptionEvent | 应用异常事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventInterruptEvent | 应用中断事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventEferSyscallHookEvent | 应用 SYSCALL Hook 事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventEferSysretHookEvent | 应用 SYSRET Hook 事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventVmcallExecutionEvent | 应用 VMCALL 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventCpuidExecutionEvent | 应用 CPUID 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventTracingEvent | 应用跟踪事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |
| ApplyEventXsetbvExecutionEvent | 应用 XSETBV 执行事件 | [ApplyEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ApplyEvents.h) |

### 事件验证函数 (ValidateEvents.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ValidateEventMonitor | 验证内存监控事件 | [ValidateEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ValidateEvents.h) |
| ValidateEventException | 验证异常事件 | [ValidateEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ValidateEvents.h) |
| ValidateEventInterrupt | 验证中断事件 | [ValidateEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ValidateEvents.h) |
| ValidateEventTrapExec | 验证 Trap 执行事件 | [ValidateEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ValidateEvents.h) |
| ValidateEventEptHookHiddenBreakpointAndInlineHooks | 验证 EPT Hook 和内联 Hook | [ValidateEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\ValidateEvents.h) |

### 扩展命令函数 (ExtensionCommands.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ExtensionCommandChangeAllMsrBitmapReadAllCores | 修改所有核心的 MSR 读取位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandResetChangeAllMsrBitmapReadAllCores | 重置所有核心的 MSR 读取位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandChangeAllMsrBitmapWriteAllCores | 修改所有核心的 MSR 写入位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandResetAllMsrBitmapWriteAllCores | 重置所有核心的 MSR 写入位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandEnableRdtscExitingAllCores | 启用所有核心的 RDTSC 退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableRdtscExitingAllCores | 禁用所有核心的 RDTSC 退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableRdtscExitingForClearingEventsAllCores | 清除事件时禁用 RDTSC 退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores | 清除事件时禁用 MOV 到调试寄存器退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandEnableRdpmcExitingAllCores | 启用所有核心的 RDPMC 退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableRdpmcExitingAllCores | 禁用所有核心的 RDPMC 退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandSetExceptionBitmapAllCores | 设置所有核心的异常位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandUnsetExceptionBitmapAllCores | 取消设置所有核心的异常位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandResetExceptionBitmapAllCores | 重置所有核心的异常位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandEnableMovDebugRegistersExitingAllCores | 启用所有核心的 MOV 调试寄存器退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableMovDebugRegistersExitingAllCores | 禁用所有核心的 MOV 调试寄存器退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandSetExternalInterruptExitingAllCores | 设置所有核心的外部中断退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores | 清除中断事件时取消外部中断退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandIoBitmapChangeAllCores | 修改所有核心的 IO 位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandIoBitmapResetAllCores | 重置所有核心的 IO 位图 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandEnableMovControlRegisterExitingAllCores | 启用所有核心的 MOV 控制寄存器退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |
| ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores | 清除事件时禁用 MOV 到控制寄存器退出 | [ExtensionCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\ExtensionCommands.h) |

### 断点管理函数 (BreakpointCommands.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| BreakpointRemoveAllBreakpoints | 移除所有断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointAddNew | 添加新断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointListOrModify | 列出或修改断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointHandleBreakpoints | 处理断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointCheckAndHandleDebuggerDefinedBreakpoints | 检查并处理调试器定义的断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointCheckAndHandleReApplyingBreakpoint | 检查并处理重新应用断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointCheckAndHandleDebugBreakpoint | 检查并处理调试断点 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |
| BreakpointRestoreTheTrapFlagOnceTriggered | 触发后恢复 Trap 标志 | [BreakpointCommands.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\commands\BreakpointCommands.h) |

### 进程管理函数 (Process.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ProcessEnableOrDisableThreadChangeMonitor | 启用或禁用线程更改监控 | [Process.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Process.h) |
| ProcessTriggerCr3ProcessChange | 触发 CR3 进程更改 | [Process.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Process.h) |
| ProcessHandleProcessChange | 处理进程更改 | [Process.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Process.h) |
| ProcessInterpretProcess | 解释进程命令 | [Process.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Process.h) |
| ProcessCheckIfEprocessIsValid | 检查 EPROCESS 是否有效 | [Process.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Process.h) |

### 线程管理函数 (Thread.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ThreadInterpretThread | 解释线程命令 | [Thread.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Thread.h) |
| ThreadEnableOrDisableThreadChangeMonitor | 启用或禁用线程更改监控 | [Thread.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Thread.h) |
| ThreadHandleThreadChange | 处理线程更改 | [Thread.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\objects\Thread.h) |

### 用户调试器函数 (Ud.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| UdInitializeUserDebugger | 初始化用户调试器 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdUninitializeUserDebugger | 反初始化用户调试器 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdHandleInstantBreak | 处理立即中断 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdApplyHardwareDebugRegister | 应用硬件调试寄存器 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdCheckAndHandleBreakpointsAndDebugBreaks | 检查并处理断点和调试中断 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdDispatchUsermodeCommands | 分发用户模式命令 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdCheckForCommand | 检查命令 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdHandleDebugEventsWhenUserDebuggerIsAttached | 用户调试器附加时处理调试事件 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |
| UdSendFormatsFunctionResult | 发送格式化函数结果 | [Ud.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Ud.h) |

### 附加/分离函数 (Attaching.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| AttachingInitialize | 初始化附加功能 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingCheckThreadInterceptionWithUserDebugger | 检查线程拦截 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingConfigureInterceptingThreads | 配置拦截线程 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingHandleEntrypointInterception | 处理入口点拦截 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingRemoveAndFreeAllProcessDebuggingDetails | 移除并释放所有进程调试详情 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingFindProcessDebuggingDetailsByToken | 通过 Token 查找进程调试详情 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingFindProcessDebuggingDetailsByProcessId | 通过 PID 查找进程调试详情 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingCheckUnhandledEptViolation | 检查未处理的 EPT 违规 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |
| AttachingReachedToValidLoadedModule | 到达有效的已加载模块 | [Attaching.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\user-level\Attaching.h) |

### 脚本引擎函数 (ScriptEngine.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| ScriptEngineWrapperGetInstructionPointer | 获取指令指针 | [ScriptEngine.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\script-engine\ScriptEngine.h) |
| ScriptEngineWrapperGetAddressOfReservedBuffer | 获取保留缓冲区地址 | [ScriptEngine.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\script-engine\ScriptEngine.h) |
| ScriptEngineUpdateTargetCoreDateTime | 更新目标核心日期时间 | [ScriptEngine.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\script-engine\ScriptEngine.h) |
| ScriptEngineGetTargetCoreTime | 获取目标核心时间 | [ScriptEngine.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\script-engine\ScriptEngine.h) |
| ScriptEngineGetTargetCoreDate | 获取目标核心日期 | [ScriptEngine.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\script-engine\ScriptEngine.h) |

### 调试器事件函数 (DebuggerEvents.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| DebuggerEventEnableEferOnAllProcessors | 在所有处理器上启用 EFER | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |
| DebuggerEventDisableEferOnAllProcessors | 在所有处理器上禁用 EFER | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |
| DebuggerEventEnableMovToCr3ExitingOnAllProcessors | 在所有处理器上启用 MOV 到 CR3 退出 | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |
| DebuggerEventDisableMovToCr3ExitingOnAllProcessors | 在所有处理器上禁用 MOV 到 CR3 退出 | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |
| DebuggerEventEnableMonitorReadWriteExec | 启用读写执行监控 | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |
| DebuggerCheckProcessOrThreadChange | 检查进程或线程更改 | [DebuggerEvents.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\events\DebuggerEvents.h) |

### 调试器核心函数 (Debugger.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| DebuggerGetRegValueWrapper | 获取寄存器值包装器 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |
| DebuggerGetLastError | 获取最后一个错误 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |
| DebuggerSetLastError | 设置最后一个错误 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |
| DebuggerInitialize | 初始化调试器 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |
| DebuggerUninitialize | 反初始化调试器 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |
| DebuggerCreateEvent | 创建事件 | [Debugger.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\core\Debugger.h) |

### 串行连接函数 (SerialConnection.h)

| 函数名称 | 功能描述 | 源代码位置 |
| ----- | ---- | ---- |
| SerialConnectionTest | 测试串行连接 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionPrepare | 准备串行连接 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionCheckPort | 检查串口 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionCheckBaudrate | 检查波特率 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionSend | 发送数据 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionRecvBuffer | 接收缓冲区 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionSendTwoBuffers | 发送两个缓冲区 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |
| SerialConnectionSendThreeBuffers | 发送三个缓冲区 | [SerialConnection.h](file:///d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\hyperkd\header\debugger\communication\SerialConnection.h) |

***

***

