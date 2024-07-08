package nix

import "github.com/numtide/nixos-facter/pkg/hwinfo"

func networking(mg *ModuleGenerator, item *hwinfo.Item) {
	if item.HardwareClass != hwinfo.HardwareClassNetworkCtrl {
		return
	}

	mg.NetworkingInterfaces = append(mg.NetworkingInterfaces, item.UnixDeviceName)
}
