package usb

import (
	"github.com/google/gousb"
)

type Device struct {
	*gousb.DeviceDesc
}

func Scan() ([]*Device, error) {

	ctx := gousb.NewContext()
	defer ctx.Close()

	var devices []*Device
	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		devices = append(devices, &Device{DeviceDesc: desc})
		return false
	})

	return devices, err
}
