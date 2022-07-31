package header


const(
AssertReturn = return; //col:19
AssertReturnFalse = return FALSE; //col:21
AssertReturnOne = return 1; //col:23
ASSERT_MESSAGE_DRIVER_NOT_LOADED "handle of the driver not found, probably the driver is not loaded.Did you use 'load' command ?n" =  //col:25
AssertReturnStmt(expr, stmt, rc) = do { if (expr) { /* likely */ } else { stmt; rc; } } while (0) //col:27
AssertShowMessageReturnStmt(expr, message, rc) = do { if (expr) { /* likely */ } else { ShowMessages(message); rc; } } while (0) //col:41
PAGE_SIZE = 0x1000 //col:59
PAGE_ALIGN(Va) = ((PVOID)((ULONG_PTR)(Va) & ~(PAGE_SIZE - 1))) //col:65
CPUID_ADDR_WIDTH = 0x80000008 //col:71
)


