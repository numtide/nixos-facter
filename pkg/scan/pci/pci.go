package pci

import (
	"fmt"
	"github.com/u-root/u-root/pkg/pci"
	"os"
	"path/filepath"
	"slices"
)

const (
	VendorBroadcom         = 0x14e4
	VendorRedHat           = 0x1af4
	VendorIntelCorporation = 0x8086
)

var (
	DevicesSTA = []uint16{
		0x4311, 0x4312, 0x4313, 0x4315, 0x4327, 0x4328,
		0x4329, 0x432a, 0x432b, 0x432c, 0x432d, 0x4353,
		0x4357, 0x4358, 0x4359, 0x4331, 0x43a0, 0x43b1,
	}

	DevicesFullMac = []uint16{
		0x43a3, 0x43df, 0x43ec, 0x43d3, 0x43d9, 0x43e9,
		0x43ba, 0x43bb, 0x43bc, 0xaa52, 0x43ca, 0x43cb,
		0x43cc, 0x43c3, 0x43c4, 0x43c5,
	}

	DevicesVirtioSCSI = []uint16{
		0x1004, 0x1048,
	}

	DevicesIntel2200BG = []uint16{
		0x1043, 0x104f, 0x4220, 0x4221, 0x4223, 0x4224,
	}

	DevicesIntel3945ABG = []uint16{
		0x4229, 0x4230, 0x4222, 0x4227,
	}
)

type Device struct {
	pci.PCI
	KernelModule string `json:",omitempty"`
}

func (d *Device) IsVendorBroadcom() bool {
	return d.Vendor == VendorBroadcom
}

func (d *Device) IsVendorRedHat() bool {
	return d.Vendor == VendorRedHat
}

func (d *Device) IsVendorIntelCorporation() bool {
	return d.Vendor == VendorIntelCorporation
}

func (d *Device) IsDeviceSTA() bool {
	return slices.Contains(DevicesSTA, d.Device)
}

func (d *Device) IsDeviceFullMac() bool {
	return slices.Contains(DevicesFullMac, d.Device)
}

func (d *Device) IsDeviceVirtioSCSI() bool {
	return slices.Contains(DevicesVirtioSCSI, d.Device)
}

func (d *Device) IsDeviceIntel2200BG() bool {
	return slices.Contains(DevicesIntel2200BG, d.Device)
}

func (d *Device) IsDeviceIntel3945ABG() bool {
	return slices.Contains(DevicesIntel3945ABG, d.Device)
}

func (d *Device) IsClass(classes ...uint32) bool {
	for _, mask := range classes {
		if d.Class&mask == 0x1 {
			return true
		}
	}
	return false
}

func (d *Device) IsStorage() bool {
	return d.IsClass(pci.ClassStorage)
}

func (d *Device) IsSerialFirewire() bool {
	return d.IsClass(pci.ClassSerialFirewire)
}

func (d *Device) IsInputKeyboard() bool {
	return d.IsClass(pci.ClassInputKeyboard)
}

func Scan() ([]*Device, error) {

	reader, err := pci.NewBusReader()
	if err != nil {
		return nil, fmt.Errorf("failed to create bus reader: %w", err)
	}

	devices, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read devices")
	}

	// convert to our device struct and attempt to resolve a kernel module
	var result []*Device
	for idx := range devices {
		dev := &Device{PCI: *devices[idx]}
		result = append(result, dev)

		path, err := os.Readlink(dev.FullPath + "/driver/module")
		if err != nil {
			// todo add some logging and check error
			continue
		}

		dev.KernelModule = filepath.Base(path)
	}

	return result, nil
}
