package nix

import "github.com/numtide/nixos-facter/pkg/hwinfo"

func usb(mg *ModuleGenerator, item *hwinfo.Item) {
	if item.Driver == "" || item.Detail == nil || item.Detail.DetailType() != hwinfo.DetailTypeUsb {
		return
	}

	detail := item.Detail.(hwinfo.DetailUsb)
	if detail.InterfaceClass == hwinfo.UsbClassMassStorage || (detail.InterfaceClass == hwinfo.UsbClassHID && detail.InterfaceProtocol == 1) {
		mg.InitrdAvailableKernelModules = append(mg.InitrdAvailableKernelModules, item.Driver)
	}
}
