package nix

import "github.com/numtide/nixos-facter/pkg/hwinfo"

func pci(mg *ModuleGenerator, item *hwinfo.Item) {

	if item.Detail == nil || item.Detail.DetailType() != hwinfo.DetailTypePci {
		return
	}

	if item.Driver != "" {
		switch item.HardwareClass {
		// usb controller is needed if we want to use the keyboard when things go wrong in the initrd
		// todo firewire controller, a disk might be attached
		case hwinfo.HardwareClassStorageCtrl, hwinfo.HardwareClassUsbCtrl:
			mg.InitrdAvailableKernelModules = append(mg.InitrdAvailableKernelModules, item.Driver)
		default:
			// do nothing
		}
	}

	vendor := item.Vendor
	device := item.Device

	if vendor.IsVendor(hwinfo.VendorBroadcom) {
		// broadcom STA driver (wl.ko)
		if device.Is(hwinfo.DevicesSTA...) {
			mg.KernelModules = append(mg.KernelModules, "wl")
			mg.ModulePackages = append(mg.ModulePackages, "config.boot.kernelPackages.broadcom_sta")
		}
		// broadcom FullMac driver
		if device.Is(hwinfo.DevicesFullMac...) {
			mg.Imports = append(mg.Imports, `(modulesPath + \"/hardware/network/broadcom-43xx.nix\")`)
		}
	}

	// In case this is a virtio scsi device, we need to explicitly make this available
	if vendor.IsVendor(hwinfo.VendorRedHat) && device.Is(hwinfo.DevicesVirtioSCSI...) {
		mg.InitrdAvailableKernelModules = append(mg.InitrdAvailableKernelModules, "virtio_scsi")
	}

	// Can't rely on module here since it may not be loaded due to missing firmware.
	if vendor.IsVendor(hwinfo.VendorIntelCorporation) {
		if device.Is(hwinfo.DevicesIntel2200BG...) {
			mg.Attrs = append(mg.Attrs, "networking.enableIntel2200BGFirmware = true")
		}

		if device.Is(hwinfo.DevicesIntel3945ABG...) {
			mg.Attrs = append(mg.Attrs, "networking.enableIntel3945ABGFirmware = true")
		}
	}

	// todo review setting the nvidia video driver which is unfree
	// https://github.com/NixOS/nixpkgs/blob/dac9cdf8c930c0af98a63cbfe8005546ba0125fb/nixos/modules/installer/tools/nixos-generate-config.pl#L199-L202
}
