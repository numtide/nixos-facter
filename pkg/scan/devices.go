package scan

import (
	"fmt"
	"github.com/u-root/u-root/pkg/pci"
	"os"
	"path/filepath"
)

type DeviceScanner struct {
}

func (p *DeviceScanner) Run(report *Report) error {
	// read devices
	var devices pci.Devices

	if reader, err := pci.NewBusReader(); err != nil {
		return fmt.Errorf("failed to create bus reader: %w", err)
	} else if devices, err = reader.Read(); err != nil {
		return fmt.Errorf("failed to read devices")
	}

	// convert to our device struct and attempt to resolve a kernel module
	report.Devices = make([]*Device, len(devices))
	for idx := range devices {
		dev := &Device{PCI: *devices[idx]}
		report.Devices[idx] = dev

		path, err := os.Readlink(dev.FullPath + "/driver/module")
		if err != nil {
			// todo add some logging and check error
			continue
		}

		dev.Module = filepath.Base(path)
	}

	return nil
}
