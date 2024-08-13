package virt

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/log"
)

func detectXen() (Type, error) {
	_, err := os.Stat("/proc/xen")
	if os.IsNotExist(err) {
		log.Debug("Virtualisation XEN not found, /proc/xen does not exist")
		return TypeNone, nil
	} else if err != nil {
		return 0, err
	}
	log.Debug("Virtualisation XEN found (/proc/xen exists)")
	return TypeXen, nil
}

const (
	xenFeatDom0      = 11 // xen/include/public/features.h
	featuresPath     = "/sys/hypervisor/properties/features"
	capabilitiesPath = "/proc/xen/capabilities"
)

// detectXenDom0 detects if the system is running as a Xen Dom0.
// It reads the features file and checks if the Dom0 feature flag is set.
// If the features file is not present, it checks the capabilities file to determine if it is a Xen guest.
// Returns a boolean indicating whether the system is running as a Xen Dom0, with false indicate Xen DomU.
func detectXenDom0() (bool, error) {
	l := log.WithPrefix("virt[XEN]")

	var result bool

	b, err := os.ReadFile(featuresPath)
	if err != nil && !os.IsNotExist(err) {
		// read failure
		return false, err
	} else if len(b) > 0 {
		features, err := strconv.ParseUint(string(b), 16, 64)
		if err != nil {
			return false, fmt.Errorf("failed to read %s: %w", featuresPath, err)
		}

		result = (features & (1 << xenFeatDom0)) != 0
		l.Debugf("found %s with value %x", featuresPath, features)
		if result {
			l.Debug("Dom0 is indicated")
		} else {
			l.Debug("DomU is indicated")
		}
		return result, nil
	}

	b, err = os.ReadFile(capabilitiesPath)
	if os.IsNotExist(err) {
		// must be running as a Xen guest
		l.Debugf("%s does not exist, DomU is indicated", capabilitiesPath)
		return false, nil
	} else if err != nil {
		return false, err
	}

	l.Debugf("found %s with value %s", capabilitiesPath, string(b))
	if strings.Contains(string(b), "control_d") {
		result = true
		l.Debug("Dom0 is indicated")
	} else {
		l.Debug("DomU is indicated")
	}

	return result, nil
}
