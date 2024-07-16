package virt

import (
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func detectUml() (Type, error) {
	b, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		if os.IsNotExist(err) {
			log.Debug("/proc/cpuinfo not found, assuming no UML virtualization")
			return TypeNone, nil
		}
		return 0, err
	}

	if strings.Contains(string(b), "vendor_id\t: User Mode Linux") {
		return TypeUml, nil
	}

	return TypeNone, nil
}
