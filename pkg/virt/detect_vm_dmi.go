//go:build 386 || amd64 || arm || arm64 || loong64 || riscv

package virt

import (
	"os"
	"strings"

	"log/slog"
)

var (
	dmiVendors = []string{
		"/sys/class/dmi/id/product_name", // Test this before sys_vendor to detect KVM over QEMU
		"/sys/class/dmi/id/sys_vendor",
		"/sys/class/dmi/id/board_vendor",
		"/sys/class/dmi/id/bios_vendor",
		"/sys/class/dmi/id/product_version", // For Hyper-V VMs test
	}

	dmiMapping = map[string]Type{
		"KVM":                TypeKvm,
		"OpenStack":          TypeKvm, // Detect OpenStack instance as KVM in non x86 architecture */
		"KubeVirt":           TypeKvm, // Detect KubeVirt instance as KVM in non x86 architecture */
		"Amazon EC2":         TypeAmazon,
		"QEMU":               TypeQemu,
		"VMware":             TypeVmware, // https://kb.vmware.com/s/article/1009458 */
		"VMW":                TypeVmware,
		"innotek GmbH":       TypeOracle,
		"VirtualBox":         TypeOracle,
		"Oracle Corporation": TypeOracle, // Detect VirtualBox on some proprietary systems via the board_vendor */
		"Xen":                TypeXen,
		"Bochs":              TypeBochs,
		"Parallels":          TypeParallels,
		/* https://wiki.freebsd.org/bhyve */
		"BHYVE":                 TypeBhyve,
		"Hyper-V":               TypeMicrosoft,
		"Apple Virtualization":  TypeApple,
		"Google Compute Engine": TypeGoogle, // https://cloud.google.com/run/d
	}
)

func init() {
	detectVmDmi = _detectVmDmi
}

func _detectVmDmi() (Type, error) {
	vmType := detectDmiVendor()

	// The DMI vendor tables in /sys/class/dmi/id don't help us distinguish between Amazon EC2
	// virtual machines and bare-metal instances, so we need to look at SMBIOS.
	if vmType == TypeAmazon {
		switch detectSmbios() {
		case SmbiosBitSet:
			return TypeAmazon, nil
		case SmbiosBitUnset:
			return TypeNone, nil
		case SmbiosBitUnknown:
			b, err := os.ReadFile("/sys/class/dmi/id/product_name")
			if err != nil {
				// In EC2, virtualized is much more common than metal, so if for some reason
				// we fail to read the DMI data, assume we are virtualized.
				slog.Debug("failed to read /sys/class/dmi/id/product_name, assuming virtualized", "error", err)
				return TypeAmazon, nil
			}
			if strings.Contains(string(b), ".metal") {
				slog.Debug("DMI product name has '.metal', assuming no virtualisation")
				return TypeNone, nil
			}

		}
	}

	// If we haven't identified a VM, but the firmware indicates that there is one, indicate as much. We have no
	// further information about what it is.
	if vmType == TypeNone && detectSmbios() == SmbiosBitSet {
		return TypeVmOther, nil
	}

	return vmType, nil
}

func detectDmiVendor() Type {
	for _, path := range dmiVendors {
		b, err := os.ReadFile(path)
		if err != nil {
			// todo log debug error
			continue
		}

		vendor := string(b)
		for k, v := range dmiMapping {
			if strings.HasPrefix(vendor, k) {
				return v
			}
		}
	}
	return TypeNone
}
