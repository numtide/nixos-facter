package virt

import (
	"log/slog"
	"os"
	"strings"
)

func detectUml() (Type, error) {
	b, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		if os.IsNotExist(err) {
			slog.Debug("assuming no UML virtualisation", "file", "/proc/cpuinfo")
			return TypeNone, nil
		}
		return 0, err
	}

	if strings.Contains(string(b), "vendor_id\t: User Mode Linux") {
		return TypeUml, nil
	}

	return TypeNone, nil
}
