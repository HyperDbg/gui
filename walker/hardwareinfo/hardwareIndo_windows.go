package hardwareinfo

type (
	Interface any
	object    struct {
		MacInfo *systemScanner
		CpuInfo *cpuInfo
		SsdInfo *ssdInfo
	}
	HardwareInfo = object
)

func New() *HardwareInfo {
	return &HardwareInfo{
		MacInfo: new(systemScanner),
		CpuInfo: new(cpuInfo),
		SsdInfo: new(ssdInfo),
	}
}
