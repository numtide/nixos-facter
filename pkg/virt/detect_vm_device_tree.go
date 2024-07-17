//go:build (arm || arm64 || ppc || ppc64 || riscv) && !noasm && !appengine

package virt

func init() {
	detectVmDeviceTree = _detectVmDeviceTree
}

func _detectVmDeviceTree() (Type, error) {
	b, err := os.ReadFile("/proc/device-tree/hypervisor/compatible")
	if os.IsNotExist(err) {

		ibm, _ := os.Stat("/proc/device-tree/ibm,partition-name")
		hmc, _ := os.Stat("/proc/device-tree/hmc-managed")
		qemu, _ := os.Stat("/proc/device-tree/chosen/qemu,graphic-width")

		if !(ibm == nil || hmc == nil) && qemu == nil {
			return TypePowerVM, nil
		}

		dir, err := os.Stat("/proc/device-tree")
		if os.IsNotExist(err) || !dir.IsDir() {
			log.Debug("/proc/device-tree directory does not exist")
			return TypeNone, nil
		}

		entries, err := os.ReadDir("/proc/device-tree")
		if err != nil {
			return 0, err
		}

		for _, entry := range entries {
			if strings.Contains(entry.Name(), "fw-cfg") {
				log.Debug("Virtualisation QEMU: \"fw-cfg\" present in /proc/device-tree/%s", entry.Name())
				return TypeQemu, nil
			}
		}

		b, err := os.ReadFile("/proc/device-tree/compatible")
		if err != nil {
			return 0, err
		} else if string(b) == "qemu,pseries" {
			log.Debug("Virtualisation %s found in /proc/device-tree/compatible", string(b))
			return TypeQemu, nil
		}

		log.Debug("No virtualisation found in /proc/device-tree/*")
		return TypeNone, nil

	} else if err != nil {
		return 0, err
	}

	log.Debug("Virtualisation %s found in /proc/device-tree/hypervisor/compatible", string(b))

	switch string(b) {
	case "linux,kvm":
		return TypeKvm, nil
	case "xen":
		return TypeXen, nil
	case "vmware":
		return TypeVmware, nil
	default:
		return TypeVmOther, nil
	}
}
