package Headers

import "syscall"

type (
	MODULE_SYMBOL_DETAIL struct {
		IsSymbolDetailsFound bool // TRUE if the details of symbols found, FALSE if not found
		IsLocalSymbolPath    bool // TRUE if the ModuleSymbolPath is a real path
		// and FALSE if ModuleSymbolPath is just a module name
		IsSymbolPDBAvaliable   bool // TRUE if the module's pdb is avilable(if exists in the sympath)
		IsUserMode             bool // TRUE if the module is a user-mode module
		BaseAddress            uint64
		FilePath               [syscall.MAX_PATH]int8
		ModuleSymbolPath       [syscall.MAX_PATH]int8
		ModuleSymbolGuidAndAge [MAXIMUM_GUID_AND_AGE_SIZE]int8
	}
	PMODULE_SYMBOL_DETAIL *MODULE_SYMBOL_DETAIL
)
type (
	USERMODE_LOADED_MODULE_SYMBOLS struct {
		BaseAddress uint64
		Entrypoint  uint64
		FilePath    [syscall.MAX_PATH]rune
	}
	PUSERMODE_LOADED_MODULE_SYMBOLS *USERMODE_LOADED_MODULE_SYMBOLS
)

type (
	USERMODE_LOADED_MODULE_DETAILS struct {
		ProcessId        uint32
		OnlyCountModules bool
		ModulesCount     uint32
		Result           uint32
		// Here is a list of USERMODE_LOADED_MODULE_SYMBOLS (appended)
	}
	PUSERMODE_LOADED_MODULE_DETAILS *USERMODE_LOADED_MODULE_DETAILS
)

type SymbolMapCallback func(Address uint64, ModuleName, ObjectName *int8, ObjectSize uint)

type (
	DEBUGGER_UPDATE_SYMBOL_TABLE struct {
		TotalSymbols       uint32
		CurrentSymbolIndex uint32
		SymbolDetailPacket MODULE_SYMBOL_DETAIL
	}
	PDEBUGGER_UPDATE_SYMBOL_TABLE *DEBUGGER_UPDATE_SYMBOL_TABLE
)

// todo make a init fn check ?
// static_assert(sizeof(DEBUGGER_UPDATE_SYMBOL_TABLE) < PacketChunkSize,
// "err (static_assert), size of PacketChunkSize should be bigger than DEBUGGER_UPDATE_SYMBOL_TABLE (MODULE_SYMBOL_DETAIL)");
type (
	DEBUGGEE_SYMBOL_UPDATE_RESULT struct {
		KernelStatus uint64
	}
	PDEBUGGEE_SYMBOL_UPDATE_RESULT *DEBUGGEE_SYMBOL_UPDATE_RESULT
)
