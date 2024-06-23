package constants

type IoctlKind int

const (
	MONITOR_IOCTL_ENABLE_MONITOR       IoctlKind = 0x00120004
	MONITOR_IOCTL_DISABLE_MONITOR                = 0x00120008
	IOCTL_INTERNAL_KEYBOARD_CONNECT              = 0x000B0203
	IOCTL_INTERNAL_KEYBOARD_DISCONNECT           = 0x000B0403
)

func (k IoctlKind) String() string {
	switch k {
	case 0x00120004:
		return "MonitorIoctlEnableMonitor"
	case 0x00120008:
		return "MonitorIoctlDisableMonitor"
	case 0x000B0203:
		return "IoctlInternalKeyboardConnect"
	case 0x000B0403:
		return "IoctlInternalKeyboardDisconnect"
	default:
		return "unknown"
	}
}

func (k IoctlKind) Elements() []IoctlKind {
	return []IoctlKind{
		MONITOR_IOCTL_ENABLE_MONITOR,
		MONITOR_IOCTL_DISABLE_MONITOR,
		IOCTL_INTERNAL_KEYBOARD_CONNECT,
		IOCTL_INTERNAL_KEYBOARD_DISCONNECT,
	}
}
