package header

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\header\common.h.back

const (
	AssertReturn = return ;       //col:1
AssertReturnFalse = return FALSE; //col:2
AssertReturnOne = return 1; //col:3
ASSERT_MESSAGE_DRIVER_NOT_LOADED "handle of the driver not found, probably the driver is not loaded.Did you use 'load' command ?n" = AssertReturnStmt(expr, stmt, rc
) do { if (expr) {  } else { stmt; rc; } } while (0) //col:4
AssertReturnStmt(expr, stmt, rc) = do { if (expr) {  } else { stmt; rc; } } while (0) //col:5
AssertShowMessageReturnStmt(expr, message, rc) = do { if (expr) {  } else { ShowMessages(message); rc; } } while (0) //col:18
PAGE_SIZE = 0x1000                                                                                                   //col:31
PAGE_ALIGN(Va) = ((PVOID)((ULONG_PTR)(Va) & ~(PAGE_SIZE - 1))) //col:32
CPUID_ADDR_WIDTH = 0x80000008                                  //col:33
)
