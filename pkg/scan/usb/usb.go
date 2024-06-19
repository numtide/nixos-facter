package usb

import (
	"fmt"
	"github.com/google/gousb"
	"os"
	"path/filepath"
	"strconv"
)

type Device struct {
	*gousb.DeviceDesc

	KernelModule string `json:",omitempty"`
}

func (d *Device) Paths() []string {
	var hubPorts string
	for _, path := range d.Path {
		hubPorts += "." + strconv.Itoa(path)
	}
	if hubPorts != "" {
		hubPorts = hubPorts[1:]
	}

	var result []string
	for _, cfg := range d.Configs {

		for _, iface := range cfg.Interfaces {
			result = append(result, fmt.Sprintf("/sys/bus/usb/devices/%d-%s:%d.%d", d.Bus, hubPorts, cfg.Number, iface.Number))
		}
	}
	return result
}

func Scan() ([]*Device, error) {

	ctx := gousb.NewContext()
	defer ctx.Close()

	var devices []*Device
	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {

		devices = append(devices, &Device{DeviceDesc: desc})

		for _, dev := range devices {
			for _, path := range dev.Paths() {
				module, err := os.Readlink(filepath.Join(path, "driver/module"))
				if err != nil {
					// todo add some logging and check error
					continue
				}

				dev.KernelModule = filepath.Base(module)
			}
		}

		return false
	})

	return devices, err
}
