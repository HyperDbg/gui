package Imports

import (
	_ "embed"
	"github.com/ddkwork/golibrary/mylog"

	"github.com/ddkwork/hyperdbgui/sdk/Headers"
)

type nameKind byte

const ( //todo gen it
	HyperDbgLoadVmm nameKind = iota
	HyperDbgUnloadVmm
	HyperDbgInstallVmmDriver
	HyperDbgUninstallVmmDriver
	HyperDbgStopVmmDriver
	HyperDbgInterpreter
	HyperDbgShowSignature
	HyperDbgSetTextMessageCallback
	HyperDbgScriptReadFileAndExecuteCommandline
	HyperDbgContinuePreviousCommand
	HyperDbgCheckMultilineCommand
	ScriptEngineParse
	PrintSymbolBuffer
	PrintSymbol
	RemoveSymbolBuffer
	ScriptEngineSetTextMessageCallback
	ScriptEngineSymbolAbortLoading
	ScriptEngineConvertNameToAddress
	ScriptEngineLoadFileSymbol
	ScriptEngineUnloadAllSymbols
	ScriptEngineUnloadModuleSymbol
	ScriptEngineSearchSymbolForMask
	ScriptEngineGetFieldOffset
	ScriptEngineGetDataTypeSize
	ScriptEngineCreateSymbolTableForDisassembler
	ScriptEngineConvertFileToPdbPath
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails
	ScriptEngineSymbolInitLoad
	ScriptEngineShowDataBasedOnSymbolTypes
	SymSetTextMessageCallback
	SymConvertNameToAddress
	SymLoadFileSymbol
	SymUnloadAllSymbols
	SymUnloadModuleSymbol
	SymSearchSymbolForMask
	SymGetFieldOffset
	SymGetDataTypeSize
	SymCreateSymbolTableForDisassembler
	SymConvertFileToPdbPath
	SymConvertFileToPdbFileAndGuidAndAgeDetails
	SymbolInitLoad
	SymShowDataBasedOnSymbolTypes
	SymQuerySizeof
	SymCastingQueryForFiledsAndTypes
)

func (k nameKind) Names() []nameKind {
	return []nameKind{
		HyperDbgLoadVmm,
		HyperDbgUnloadVmm,
		HyperDbgInstallVmmDriver,
		HyperDbgUninstallVmmDriver,
		HyperDbgStopVmmDriver,
		HyperDbgInterpreter,
		HyperDbgShowSignature,
		HyperDbgSetTextMessageCallback,
		HyperDbgScriptReadFileAndExecuteCommandline,
		HyperDbgContinuePreviousCommand,
		HyperDbgCheckMultilineCommand,
		ScriptEngineParse,
		PrintSymbolBuffer,
		PrintSymbol,
		RemoveSymbolBuffer,
		ScriptEngineSetTextMessageCallback,
		ScriptEngineSymbolAbortLoading,
		ScriptEngineConvertNameToAddress,
		ScriptEngineLoadFileSymbol,
		ScriptEngineUnloadAllSymbols,
		ScriptEngineUnloadModuleSymbol,
		ScriptEngineSearchSymbolForMask,
		ScriptEngineGetFieldOffset,
		ScriptEngineGetDataTypeSize,
		ScriptEngineCreateSymbolTableForDisassembler,
		ScriptEngineConvertFileToPdbPath,
		ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails,
		ScriptEngineSymbolInitLoad,
		ScriptEngineShowDataBasedOnSymbolTypes,
		SymSetTextMessageCallback,
		SymConvertNameToAddress,
		SymLoadFileSymbol,
		SymUnloadAllSymbols,
		SymUnloadModuleSymbol,
		SymSearchSymbolForMask,
		SymGetFieldOffset,
		SymGetDataTypeSize,
		SymCreateSymbolTableForDisassembler,
		SymConvertFileToPdbPath,
		SymConvertFileToPdbFileAndGuidAndAgeDetails,
		SymbolInitLoad,
		SymShowDataBasedOnSymbolTypes,
		SymQuerySizeof,
		SymCastingQueryForFiledsAndTypes,
	}
}
func (k nameKind) Len() int { return len(k.Names()) } //means dll export ? number of Api
func (k nameKind) String() string {
	switch k {
	case HyperDbgLoadVmm:
		return "HyperDbgLoadVmm"
	case HyperDbgUnloadVmm:
		return "HyperDbgUnloadVmm"
	case HyperDbgInstallVmmDriver:
		return "HyperDbgInstallVmmDriver"
	case HyperDbgUninstallVmmDriver:
		return "HyperDbgUninstallVmmDriver"
	case HyperDbgStopVmmDriver:
		return "HyperDbgStopVmmDriver"
	case HyperDbgInterpreter:
		return "HyperDbgInterpreter"
	case HyperDbgShowSignature:
		return "HyperDbgShowSignature"
	case HyperDbgSetTextMessageCallback:
		return "HyperDbgSetTextMessageCallback"
	case HyperDbgScriptReadFileAndExecuteCommandline:
		return "HyperDbgScriptReadFileAndExecuteCommandline"
	case HyperDbgContinuePreviousCommand:
		return "HyperDbgContinuePreviousCommand"
	case HyperDbgCheckMultilineCommand:
		return "HyperDbgCheckMultilineCommand"
	case ScriptEngineParse:
		return "ScriptEngineParse"
	case PrintSymbolBuffer:
		return "PrintSymbolBuffer"
	case PrintSymbol:
		return "PrintSymbol"
	case RemoveSymbolBuffer:
		return "RemoveSymbolBuffer"
	case ScriptEngineSetTextMessageCallback:
		return "ScriptEngineSetTextMessageCallback"
	case ScriptEngineSymbolAbortLoading:
		return "ScriptEngineSymbolAbortLoading"
	case ScriptEngineConvertNameToAddress:
		return "ScriptEngineConvertNameToAddress"
	case ScriptEngineLoadFileSymbol:
		return "ScriptEngineLoadFileSymbol"
	case ScriptEngineUnloadAllSymbols:
		return "ScriptEngineUnloadAllSymbols"
	case ScriptEngineUnloadModuleSymbol:
		return "ScriptEngineUnloadModuleSymbol"
	case ScriptEngineSearchSymbolForMask:
		return "ScriptEngineSearchSymbolForMask"
	case ScriptEngineGetFieldOffset:
		return "ScriptEngineGetFieldOffset"
	case ScriptEngineGetDataTypeSize:
		return "ScriptEngineGetDataTypeSize"
	case ScriptEngineCreateSymbolTableForDisassembler:
		return "ScriptEngineCreateSymbolTableForDisassembler"
	case ScriptEngineConvertFileToPdbPath:
		return "ScriptEngineConvertFileToPdbPath"
	case ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails:
		return "ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails"
	case ScriptEngineSymbolInitLoad:
		return "ScriptEngineSymbolInitLoad"
	case ScriptEngineShowDataBasedOnSymbolTypes:
		return "ScriptEngineShowDataBasedOnSymbolTypes"
	case SymSetTextMessageCallback:
		return "SymSetTextMessageCallback"
	case SymConvertNameToAddress:
		return "SymConvertNameToAddress"
	case SymLoadFileSymbol:
		return "SymLoadFileSymbol"
	case SymUnloadAllSymbols:
		return "SymUnloadAllSymbols"
	case SymUnloadModuleSymbol:
		return "SymUnloadModuleSymbol"
	case SymSearchSymbolForMask:
		return "SymSearchSymbolForMask"
	case SymGetFieldOffset:
		return "SymGetFieldOffset"
	case SymGetDataTypeSize:
		return "SymGetDataTypeSize"
	case SymCreateSymbolTableForDisassembler:
		return "SymCreateSymbolTableForDisassembler"
	case SymConvertFileToPdbPath:
		return "SymConvertFileToPdbPath"
	case SymConvertFileToPdbFileAndGuidAndAgeDetails:
		return "SymConvertFileToPdbFileAndGuidAndAgeDetails"
	case SymbolInitLoad:
		return "SymbolInitLoad"
	case SymShowDataBasedOnSymbolTypes:
		return "SymShowDataBasedOnSymbolTypes"
	case SymQuerySizeof:
		return "SymQuerySizeof"
	case SymCastingQueryForFiledsAndTypes:
		return "SymCastingQueryForFiledsAndTypes"
	default:
		return "error fn name from dll, please check update" //this should panic?
	}
}

func Bool2UintPtr(b bool) uintptr {
	v := 1
	if b {
		v = 0
	}
	return uintptr(v)
}

func DecodeErrorCode(code uintptr) string {
	status := Headers.ErrorCodes(code).String()
	if status != "" {
		mylog.Error(status)
	}
	return status
}
