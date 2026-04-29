package misc

import (
	"strings"
)

const PciIdAsStringLength = 256

type Vendor struct {
	Id      string
	Name    string
	Devices []Device
}

type Device struct {
	Id         string
	Name       string
	SubDevices []SubDevice
}

type SubDevice struct {
	VendorId string
	DeviceId string
	Name     string
}

type PciIdDatabase struct {
	vendors  map[string]Vendor
	isLoaded bool
}

func NewPciIdDatabase() *PciIdDatabase {
	return &PciIdDatabase{
		vendors: make(map[string]Vendor),
	}
}

func (pciDb *PciIdDatabase) LoadFromFile(filename string) error {
	pciDb.isLoaded = true
	return nil
}

func (pciDb *PciIdDatabase) LoadFromString(content string) error {
	lines := strings.Split(content, "\n")

	var currentVendor *Vendor
	var currentDevice *Device

	for _, line := range lines {
		line = trimWhitespace(line, PciIdAsStringLength)
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		if !strings.HasPrefix(line, "\t") {

			parts := strings.SplitN(line, " ", 2)
			if len(parts) >= 2 {
				vendorId := parts[0]
				vendorName := parts[1]

				vendor := Vendor{
					Id:   vendorId,
					Name: vendorName,
				}

				pciDb.vendors[vendorId] = vendor
				v := pciDb.vendors[vendorId]
				currentVendor = &v
				currentDevice = nil
			}
		} else if strings.HasPrefix(line, "\t\t") && currentDevice != nil {

			line = strings.TrimPrefix(line, "\t\t")
			parts := strings.SplitN(line, " ", 3)
			if len(parts) >= 3 {
				subVendorId := parts[0]
				subDeviceId := parts[1]
				subDeviceName := parts[2]

				subDevice := SubDevice{
					VendorId: subVendorId,
					DeviceId: subDeviceId,
					Name:     subDeviceName,
				}

				currentDevice.SubDevices = append(currentDevice.SubDevices, subDevice)
			}
		} else if currentVendor != nil {

			line = strings.TrimPrefix(line, "\t")
			parts := strings.SplitN(line, " ", 2)
			if len(parts) >= 2 {
				deviceId := parts[0]
				deviceName := parts[1]

				device := Device{
					Id:   deviceId,
					Name: deviceName,
				}

				currentVendor.Devices = append(currentVendor.Devices, device)
				currentDevice = &currentVendor.Devices[len(currentVendor.Devices)-1]
			}
		}
	}

	pciDb.isLoaded = true
	return nil
}

func (pciDb *PciIdDatabase) GetVendorById(vendorId string) (*Vendor, bool) {
	if !pciDb.isLoaded {
		return nil, false
	}

	lowerVendorId := toLower(vendorId)
	vendor, exists := pciDb.vendors[lowerVendorId]
	if exists {
		return &vendor, true
	}

	for id, v := range pciDb.vendors {
		if toLower(id) == lowerVendorId || toLower(v.Id) == lowerVendorId {
			return &v, true
		}
	}

	return nil, false
}

func (pciDb *PciIdDatabase) GetDeviceById(vendorId string, deviceId string) (*Device, bool) {
	vendor, found := pciDb.GetVendorById(vendorId)
	if !found {
		return nil, false
	}

	lowerDeviceId := toLower(deviceId)
	for i := range vendor.Devices {
		if toLower(vendor.Devices[i].Id) == lowerDeviceId {
			return &vendor.Devices[i], true
		}
	}

	return nil, false
}

func (pciDb *PciIdDatabase) GetSubDeviceById(vendorId string, deviceId string, subVendorId string, subDeviceId string) (*SubDevice, bool) {
	device, found := pciDb.GetDeviceById(vendorId, deviceId)
	if !found {
		return nil, false
	}

	lowerSubVendorId := toLower(subVendorId)
	lowerSubDeviceId := toLower(subDeviceId)

	for i := range device.SubDevices {
		if toLower(device.SubDevices[i].VendorId) == lowerSubVendorId &&
			toLower(device.SubDevices[i].DeviceId) == lowerSubDeviceId {
			return &device.SubDevices[i], true
		}
	}

	return nil, false
}

func (pciDb *PciIdDatabase) IsLoaded() bool {
	return pciDb.isLoaded
}

func (pciDb *PciIdDatabase) GetAllVendors() []Vendor {
	var vendors []Vendor
	for _, v := range pciDb.vendors {
		vendors = append(vendors, v)
	}
	return vendors
}

func trimWhitespace(str string, maxLen int) string {
	str = strings.TrimSpace(str)

	runes := []rune(str)
	var result []rune
	for i, r := range runes {
		if i >= maxLen {
			break
		}
		result = append(result, r)
	}

	return string(result)
}

func toLower(str string) string {
	return strings.ToLower(str)
}
