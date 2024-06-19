package nix

import (
	"embed"
	"fmt"
	"github.com/numtide/nixos-facter/pkg/scan"
	"github.com/u-root/u-root/pkg/pci"
	"io"
	"slices"
	"text/template"
)

var (
	//go:embed templates
	templatesFS embed.FS
)

type ModuleGenerator struct {
	Report *scan.Report

	Attrs                        []string
	Imports                      []string
	KernelModules                []string
	ModulePackages               []string
	InitrdKernelModules          []string
	InitrdAvailableKernelModules []string
}

func (mg *ModuleGenerator) Generate(writer io.Writer) error {
	mg.processDevices()

	// remove duplicates
	mg.Attrs = slices.Compact(mg.Attrs)
	mg.Imports = slices.Compact(mg.Imports)
	mg.KernelModules = slices.Compact(mg.KernelModules)
	mg.ModulePackages = slices.Compact(mg.ModulePackages)
	mg.InitrdKernelModules = slices.Compact(mg.InitrdKernelModules)
	mg.InitrdAvailableKernelModules = slices.Compact(mg.InitrdAvailableKernelModules)

	funcMap := template.FuncMap{
		"nixList":       ToNixList,
		"nixStringList": ToNixStringList,
		"multiLineList": MultiLineList,
	}

	tmpl, err := template.
		New("hw_config.tmpl").
		Funcs(funcMap).
		ParseFS(templatesFS, "**/*.tmpl")

	if err != nil {
		return fmt.Errorf("failed to load hardware config template: %w", err)
	}

	return tmpl.Funcs(funcMap).Execute(writer, mg)
}

func (mg *ModuleGenerator) processDevices() {
	for _, dev := range mg.Report.PCI {

		if dev.KernelModule != "" {
			if dev.IsClass(
				// mass-storage controller
				pci.ClassStorage,
				// firewire controller, a disk might be attached
				pci.ClassSerialFirewire,
				// usb controller, needed if we want to use the keyboard when things go wrong in the initrd
				pci.ClassSerialUSB,
			) {
				mg.InitrdAvailableKernelModules = append(mg.InitrdAvailableKernelModules, dev.KernelModule)
			}
		}

		if dev.IsVendorBroadcom() {

			// broadcom STA driver (wl.ko)
			// list taken from https://github.com/NixOS/nixpkgs/blob/dac9cdf8c930c0af98a63cbfe8005546ba0125fb/nixos/modules/installer/tools/nixos-generate-config.pl#L152-L158
			if dev.IsDeviceSTA() {
				mg.KernelModules = append(mg.KernelModules, "wl")
				mg.ModulePackages = append(mg.ModulePackages, "config.boot.kernelPackages.broadcom_sta")
			}

			// broadcom FullMac driver
			// list taken from https://wireless.wiki.kernel.org/en/users/Drivers/brcm80211#brcmfmac
			if dev.IsDeviceFullMac() {
				mg.Imports = append(mg.Imports, `(modulesPath + \"/hardware/network/broadcom-43xx.nix\")`)
			}

		}

		// In case this is a virtio scsi device, we need to explicitly make this available
		if dev.IsVendorRedHat() && dev.IsDeviceVirtioSCSI() {
			mg.InitrdAvailableKernelModules = append(mg.InitrdAvailableKernelModules, "virtio_scsi")
		}

		// Can't rely on module here since it may not be loaded due to missing firmware.
		if dev.IsVendorIntelCorporation() {

			if dev.IsDeviceIntel2200BG() {
				mg.Attrs = append(mg.Attrs, "networking.enableIntel2200BGFirmware = true")
			}

			if dev.IsDeviceIntel3945ABG() {
				mg.Attrs = append(mg.Attrs, "networking.enableIntel3945ABGFirmware = true")
			}

		}

		// todo review setting the nvidia video driver which is unfree
		// https://github.com/NixOS/nixpkgs/blob/dac9cdf8c930c0af98a63cbfe8005546ba0125fb/nixos/modules/installer/tools/nixos-generate-config.pl#L199-L202
	}
}
