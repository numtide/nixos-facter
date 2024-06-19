package scan

import (
	"fmt"
	"github.com/numtide/nixos-facter/pkg/scan/block"
	"github.com/numtide/nixos-facter/pkg/scan/pci"
	"github.com/numtide/nixos-facter/pkg/scan/usb"
)

type Report struct {
	// PCIDevices is a list of PCI devices
	PCIDevices []*pci.Device
	// USBDevices is a list of USB devices
	USBDevices []*usb.Device
	// BlockDevices is a list of Block devices
	BlockDevices []*block.Device
}

type Scanner interface {
	Run(report *Report) error
}

func Run() (report *Report, err error) {

	report = &Report{}

	if report.PCIDevices, err = pci.Scan(); err != nil {
		return nil, fmt.Errorf("failed to scan pci devices: %w", err)
	}

	if report.USBDevices, err = usb.Scan(); err != nil {
		return nil, fmt.Errorf("failed to scan usb devices: %w", err)
	}

	if report.BlockDevices, err = block.Scan(); err != nil {
		return nil, fmt.Errorf("failed to scan block devices: %w", err)
	}

	return report, nil
}
