package ark

import (
	"strings"

	"golang.org/x/exp/constraints"
)

// Code generated by GeneratedFile enum - DO NOT EDIT.

type ArksKind byte

const (
	KernelTablesKind ArksKind = iota
	ExplorerKind
	TaskManagerKind
	DriverToolKind
	RegistryEditorKind
	HardwareMonitorKind
	HardwareHookKind
	RandomHookKind
	EnvironmentEditorKind
	VstartKind
	CryptKind
	PackerKind
	InvalidArksKind
)

func ConvertInteger2ArksKind[T constraints.Integer](v T) ArksKind {
	return ArksKind(v)
}

func (k ArksKind) AssertKind(kinds string) ArksKind {
	for _, kind := range k.Kinds() {
		if strings.ToLower(kinds) == strings.ToLower(kind.String()) {
			return kind
		}
	}
	return InvalidArksKind
}

func (k ArksKind) String() string {
	switch k {
	case KernelTablesKind:
		return "KernelTables"
	case ExplorerKind:
		return "Explorer"
	case TaskManagerKind:
		return "TaskManager"
	case DriverToolKind:
		return "DriverTool"
	case RegistryEditorKind:
		return "RegistryEditor"
	case HardwareMonitorKind:
		return "HardwareMonitor"
	case HardwareHookKind:
		return "HardwareHook"
	case RandomHookKind:
		return "RandomHook"
	case EnvironmentEditorKind:
		return "EnvironmentEditor"
	case VstartKind:
		return "Vstart"
	case CryptKind:
		return "Crypt"
	case PackerKind:
		return "Packer"
	default:
		return "InvalidArksKind"
	}
}

func (k ArksKind) Keys() []string {
	return []string{
		"KernelTables",
		"Explorer",
		"TaskManager",
		"DriverTool",
		"RegistryEditor",
		"HardwareMonitor",
		"HardwareHook",
		"RandomHook",
		"EnvironmentEditor",
		"Vstart",
		"Crypt",
		"Packer",
	}
}

func (k ArksKind) Kinds() []ArksKind {
	return []ArksKind{
		KernelTablesKind,
		ExplorerKind,
		TaskManagerKind,
		DriverToolKind,
		RegistryEditorKind,
		HardwareMonitorKind,
		HardwareHookKind,
		RandomHookKind,
		EnvironmentEditorKind,
		VstartKind,
		CryptKind,
		PackerKind,
	}
}

func (k ArksKind) PngFileName() string {
	switch k {
	case KernelTablesKind:
		return "KernelTables"
	case ExplorerKind:
		return "Explorer"
	case TaskManagerKind:
		return "TaskManager"
	case DriverToolKind:
		return "DriverTool"
	case RegistryEditorKind:
		return "RegistryEditor"
	case HardwareMonitorKind:
		return "HardwareMonitor"
	case HardwareHookKind:
		return "HardwareHook"
	case RandomHookKind:
		return "RandomHook"
	case EnvironmentEditorKind:
		return "EnvironmentEditor"
	case VstartKind:
		return "Vstart"
	case CryptKind:
		return "Crypt"
	case PackerKind:
		return "Packer"
	default:
		return "InvalidArksKind"
	}
}
