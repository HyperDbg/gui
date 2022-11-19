package Imports

import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"golang.org/x/sys/windows"
	"os"
	"path/filepath"
	"syscall"
)

type nameKind byte

const ( //todo gen it with rename, api name is too long
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

type (
	Api struct {
		procMap map[nameKind]*syscall.Proc
	}
)

var api = newApi()

func newApi() *Api {
	p := &Api{
		procMap: make(map[nameKind]*syscall.Proc, HyperDbgLoadVmm.Len()),
	}
	for _, kind := range HyperDbgLoadVmm.Names() {
		p.SetProc(kind)
	}
	return p
}
func (a *Api) Proc(name nameKind) *syscall.Proc { return a.procMap[name] }
func (a *Api) SetProc(name nameKind)            { a.procMap[name] = dll.MustFindProc(name.String()) }

var (
	//go:embed HPRDBGCTRL.dll
	dllData []byte
	dll     *syscall.DLL
)

func init() {
	dir, err := os.UserCacheDir()
	if !mylog.Error(err) {
		return
	}
	dir = filepath.Join(dir, "hyperdbgDll", "dll_cache")
	if !mylog.Error(os.MkdirAll(dir, 0755)) {
		return
	}
	windows.SetDllDirectory(dir)
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("hyperdbgDll-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	_, err = os.Stat(filePath)
	if !mylog.Error(err) {
		if err == os.ErrNotExist {
			if !mylog.Error(os.WriteFile(filePath, dllData, 0644)) {
				return
			}
		}
		return
	}
	dll = syscall.MustLoadDLL(dllName)
}
