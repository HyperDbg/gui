package Imports

//todo goto windows modify this file

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

type api byte

const ( //todo gen it
	HyperDbgLoadVmm api = iota
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

func (a api) String() string {
	switch a {
	case HyperDbgLoadVmm:
		return "HyperDbgLoadVmm"
	default:
		return "error fn name from dll, please check update" //this should panic?
	}
}

var (
	//go:embed HPRDBGCTRL.dll
	dllData                                                  []byte
	HyperDbgLoadVmmProc                                      *syscall.Proc
	HyperDbgUnloadVmmProc                                    *syscall.Proc
	HyperDbgInstallVmmDriverProc                             *syscall.Proc
	HyperDbgUninstallVmmDriverProc                           *syscall.Proc
	HyperDbgStopVmmDriverProc                                *syscall.Proc
	HyperDbgInterpreterProc                                  *syscall.Proc
	HyperDbgShowSignatureProc                                *syscall.Proc
	HyperDbgSetTextMessageCallbackProc                       *syscall.Proc
	HyperDbgScriptReadFileAndExecuteCommandlineProc          *syscall.Proc
	HyperDbgContinuePreviousCommandProc                      *syscall.Proc
	HyperDbgCheckMultilineCommandProc                        *syscall.Proc
	ScriptEngineParseProc                                    *syscall.Proc
	PrintSymbolBufferProc                                    *syscall.Proc
	PrintSymbolProc                                          *syscall.Proc
	RemoveSymbolBufferProc                                   *syscall.Proc
	ScriptEngineSetTextMessageCallbackProc                   *syscall.Proc
	ScriptEngineSymbolAbortLoadingProc                       *syscall.Proc
	ScriptEngineConvertNameToAddressProc                     *syscall.Proc
	ScriptEngineLoadFileSymbolProc                           *syscall.Proc
	ScriptEngineUnloadAllSymbolsProc                         *syscall.Proc
	ScriptEngineUnloadModuleSymbolProc                       *syscall.Proc
	ScriptEngineSearchSymbolForMaskProc                      *syscall.Proc
	ScriptEngineGetFieldOffsetProc                           *syscall.Proc
	ScriptEngineGetDataTypeSizeProc                          *syscall.Proc
	ScriptEngineCreateSymbolTableForDisassemblerProc         *syscall.Proc
	ScriptEngineConvertFileToPdbPathProc                     *syscall.Proc
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsProc *syscall.Proc
	ScriptEngineSymbolInitLoadProc                           *syscall.Proc
	ScriptEngineShowDataBasedOnSymbolTypesProc               *syscall.Proc
	SymSetTextMessageCallbackProc                            *syscall.Proc
	SymConvertNameToAddressProc                              *syscall.Proc
	SymLoadFileSymbolProc                                    *syscall.Proc
	SymUnloadAllSymbolsProc                                  *syscall.Proc
	SymUnloadModuleSymbolProc                                *syscall.Proc
	SymSearchSymbolForMaskProc                               *syscall.Proc
	SymGetFieldOffsetProc                                    *syscall.Proc
	SymGetDataTypeSizeProc                                   *syscall.Proc
	SymCreateSymbolTableForDisassemblerProc                  *syscall.Proc
	SymConvertFileToPdbPathProc                              *syscall.Proc
	SymConvertFileToPdbFileAndGuidAndAgeDetailsProc          *syscall.Proc
	SymbolInitLoadProc                                       *syscall.Proc
	SymShowDataBasedOnSymbolTypesProc                        *syscall.Proc
	SymQuerySizeofProc                                       *syscall.Proc
	SymCastingQueryForFiledsAndTypesProc                     *syscall.Proc
)

func init() {
	dir, err := os.UserCacheDir()
	if !mylog.Error(err) {
		return
	}
	dir = filepath.Join(dir, "hyperdbg", "dll_cache")
	if !mylog.Error(os.MkdirAll(dir, 0755)) {
		return
	}
	windows.SetDllDirectory(dir)
	sha := sha256.Sum256(dllData)
	dllName := fmt.Sprintf("hyperdbg-%s.dll", base64.RawURLEncoding.EncodeToString(sha[:]))
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
	hyperdbg := syscall.MustLoadDLL(dllName)
	HyperDbgLoadVmmProc = hyperdbg.MustFindProc(HyperDbgLoadVmm.String())
	HyperDbgUnloadVmmProc = hyperdbg.MustFindProc(HyperDbgUnloadVmm.String())
	HyperDbgInstallVmmDriverProc = hyperdbg.MustFindProc(HyperDbgInstallVmmDriver.String())
	HyperDbgUninstallVmmDriverProc = hyperdbg.MustFindProc(HyperDbgUninstallVmmDriver.String())
	HyperDbgStopVmmDriverProc = hyperdbg.MustFindProc(HyperDbgStopVmmDriver.String())
	HyperDbgInterpreterProc = hyperdbg.MustFindProc(HyperDbgInterpreter.String())
	HyperDbgShowSignatureProc = hyperdbg.MustFindProc(HyperDbgShowSignature.String())
	HyperDbgSetTextMessageCallbackProc = hyperdbg.MustFindProc("") //todo gen another
	HyperDbgScriptReadFileAndExecuteCommandlineProc = hyperdbg.MustFindProc("")
	HyperDbgContinuePreviousCommandProc = hyperdbg.MustFindProc("")
	HyperDbgCheckMultilineCommandProc = hyperdbg.MustFindProc("")
	ScriptEngineParseProc = hyperdbg.MustFindProc("")
	PrintSymbolBufferProc = hyperdbg.MustFindProc("")
	PrintSymbolProc = hyperdbg.MustFindProc("")
	RemoveSymbolBufferProc = hyperdbg.MustFindProc("")
	ScriptEngineSetTextMessageCallbackProc = hyperdbg.MustFindProc("")
	ScriptEngineSymbolAbortLoadingProc = hyperdbg.MustFindProc("")
	ScriptEngineConvertNameToAddressProc = hyperdbg.MustFindProc("")
	ScriptEngineLoadFileSymbolProc = hyperdbg.MustFindProc("")
	ScriptEngineUnloadAllSymbolsProc = hyperdbg.MustFindProc("")
	ScriptEngineUnloadModuleSymbolProc = hyperdbg.MustFindProc("")
	ScriptEngineSearchSymbolForMaskProc = hyperdbg.MustFindProc("")
	ScriptEngineGetFieldOffsetProc = hyperdbg.MustFindProc("")
	ScriptEngineGetDataTypeSizeProc = hyperdbg.MustFindProc("")
	ScriptEngineCreateSymbolTableForDisassemblerProc = hyperdbg.MustFindProc("")
	ScriptEngineConvertFileToPdbPathProc = hyperdbg.MustFindProc("")
	ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsProc = hyperdbg.MustFindProc("")
	ScriptEngineSymbolInitLoadProc = hyperdbg.MustFindProc("")
	ScriptEngineShowDataBasedOnSymbolTypesProc = hyperdbg.MustFindProc("")
	SymSetTextMessageCallbackProc = hyperdbg.MustFindProc("")
	SymConvertNameToAddressProc = hyperdbg.MustFindProc("")
	SymLoadFileSymbolProc = hyperdbg.MustFindProc("")
	SymUnloadAllSymbolsProc = hyperdbg.MustFindProc("")
	SymUnloadModuleSymbolProc = hyperdbg.MustFindProc("")
	SymSearchSymbolForMaskProc = hyperdbg.MustFindProc("")
	SymGetFieldOffsetProc = hyperdbg.MustFindProc("")
	SymGetDataTypeSizeProc = hyperdbg.MustFindProc("")
	SymCreateSymbolTableForDisassemblerProc = hyperdbg.MustFindProc("")
	SymConvertFileToPdbPathProc = hyperdbg.MustFindProc("")
	SymConvertFileToPdbFileAndGuidAndAgeDetailsProc = hyperdbg.MustFindProc("")
	SymbolInitLoadProc = hyperdbg.MustFindProc("")
	SymShowDataBasedOnSymbolTypesProc = hyperdbg.MustFindProc("")
	SymQuerySizeofProc = hyperdbg.MustFindProc("")
	SymCastingQueryForFiledsAndTypesProc = hyperdbg.MustFindProc("")
}
