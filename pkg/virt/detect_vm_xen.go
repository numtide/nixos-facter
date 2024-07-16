package virt

import (
	"os"

	"github.com/charmbracelet/log"
)

func detectXen() (Type, error) {
	_, err := os.Stat("/proc/xen")
	if os.IsNotExist(err) {
		log.Debug("Virtualization XEN not found, /proc/xen does not exist")
		return TypeNone, nil
	} else if err != nil {
		return 0, err
	}
	log.Debug("Virtualization XEN found (/proc/xen exists)")
	return TypeXen, nil
}

func detectXenDom0() (bool, error) {
	// todo implement
	panic("implement me")
}
