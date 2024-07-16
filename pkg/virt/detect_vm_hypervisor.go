package virt

import (
	"os"

	"github.com/charmbracelet/log"
)

func detectHypervisor() (Type, error) {
	b, err := os.ReadFile("/sys/hypervisor/type")

	if os.IsNotExist(err) {
		log.Debug("failed to read /sys/hypervisor/type")
		return TypeNone, nil
	} else if err != nil {
		return 0, err
	}

	hvType := string(b)
	log.Debug("Virtualization %s found in /sys/hypervisor/type", hvType)

	if hvType == "xen" {
		return TypeXen, nil
	} else {
		return TypeVmOther, nil
	}
}
