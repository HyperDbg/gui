package driver

import (
	"path/filepath"
)

type DriverConfig struct {
	DeviceName     string
	ServiceName    string
	DriverFileName string
}

func NewDriverConfig() *DriverConfig {
	return &DriverConfig{}
}

func NewDriverConfigFromPath(driverPath string) *DriverConfig {
	config := &DriverConfig{}
	config.SetDefaultsFromDriverPath(driverPath)
	return config
}

func (dc *DriverConfig) SetDefaultsFromDriverPath(driverPath string) {
	if driverPath != "" {
		dc.DriverFileName = filepath.Base(driverPath)
		dc.ServiceName = dc.DriverFileName[:len(dc.DriverFileName)-4]
	}
}
