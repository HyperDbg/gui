package kerneldump

import (
	"strings"
)

type StackAnalysis struct {
	ExceptionCode    uint32
	ExceptionAddress uint64
	ExceptionFound   bool
	ModuleBase       uint64
	ModuleSize       uint32
	ModuleName       string
	ModuleFound      bool
}

func (kd *KernelDump) AnalyzeStackForExceptionAndModule() (*StackAnalysis, error) {
	frames, err := kd.GetStackTrace()
	if err != nil {
		return nil, err
	}

	result := &StackAnalysis{
		ExceptionFound: false,
		ModuleFound:    false,
	}

	for _, frame := range frames {
		retAddr := frame.RetAddr

		if retAddr >= 0 && retAddr <= 0xFFFFFFFF {
			exceptCode := uint32(retAddr)

			if exceptCode >= 0x80000000 && exceptCode <= 0x800000FF {
				result.ExceptionCode = exceptCode
				result.ExceptionFound = true
				break
			}
		}
	}

	if result.ExceptionFound {
		modules, err := kd.GetModules()
		if err == nil && len(modules) > 0 {
			for _, mod := range modules {
				if strings.Contains(strings.ToLower(mod.GetName()), "hyperkd") {
					result.ModuleBase = mod.DllBase
					result.ModuleSize = mod.SizeOfImage
					result.ModuleName = mod.GetName()
					result.ModuleFound = true
					break
				}
			}
		}
	}

	return result, nil
}

func (kd *KernelDump) ExtractExceptionCodeFromStack() (uint32, bool) {
	analysis, err := kd.AnalyzeStackForExceptionAndModule()
	if err != nil {
		return 0, false
	}

	return analysis.ExceptionCode, analysis.ExceptionFound
}

func (kd *KernelDump) ExtractModuleInfoFromStack() (string, uint64, uint32, bool) {
	analysis, err := kd.AnalyzeStackForExceptionAndModule()
	if err != nil {
		return "", 0, 0, false
	}

	return analysis.ModuleName, analysis.ModuleBase, analysis.ModuleSize, analysis.ModuleFound
}
