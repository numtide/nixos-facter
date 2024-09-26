package virt

import (
	"log/slog"
	"os"
)

func detectHypervisor() (Type, error) {
	b, err := os.ReadFile("/sys/hypervisor/type")

	if os.IsNotExist(err) {
		slog.Debug("failed to read /sys/hypervisor/type")
		return TypeNone, nil
	} else if err != nil {
		return 0, err
	}

	hvType := string(b)
	slog.Debug("Virtualisation found in /sys/hypervisor/type", "type", hvType)

	if hvType == "xen" {
		return TypeXen, nil
	} else {
		return TypeVmOther, nil
	}
}
