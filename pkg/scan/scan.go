package scan

import (
	"fmt"
	"github.com/numtide/nixos-facter/pkg/scan/pci"
	"github.com/numtide/nixos-facter/pkg/scan/usb"
)

type Report struct {
	// PCI is a list of PCI devices
	PCI []*pci.Device
	// USB is a list of USB devices
	USB []*usb.Device
}

type Scanner interface {
	Run(report *Report) error
}

func Run() (report *Report, err error) {

	report = &Report{}

	if report.PCI, err = pci.Scan(); err != nil {
		return nil, fmt.Errorf("failed to scan pci devices: %w", err)
	}

	if report.USB, err = usb.Scan(); err != nil {
		return nil, fmt.Errorf("failed to scan usb devices: %w", err)
	}

	return report, nil
}
