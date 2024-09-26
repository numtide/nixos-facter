// Package ephem contains utilities for capturing ephemeral aspects of a target machine.
//
// Currently, it only supports capturing swap configurations.
// Eventually it will capture more things, such as filesystems.
package ephem

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var deviceGlobs = []string{
	"/dev/stratis/*/*",
	"/dev/disk/by-uuid/*",
	"/dev/mapper/*",
	"/dev/disk/by-label/*",
}

// StableDevicePath takes a device path and converts it into a more stable form.
// For example, /dev/nvme* is assigned on startup by detection order which is not consistent.
// A disk path of the form /dev/disk/by-uuid/* is not startup-dependent.
func StableDevicePath(device string) (string, error) {
	l := slog.With("prefix", "stableDevicePath")

	if !strings.HasPrefix("/", device) {
		return device, nil
	}
	stat, err := os.Stat(device)
	if err != nil {
		return "", err
	}

	for idx := range deviceGlobs {
		glob := deviceGlobs[idx]
		l.Debug("searching glob", "glob", glob)

		matches, err := filepath.Glob(glob)
		if err != nil {
			// the only possible error is ErrBadPattern
			return "", err
		}
		for _, match := range matches {
			matchStat, err := os.Stat(match)
			if err != nil {
				l.Debug("failed to stat match", "match", match, "error", err)
				continue
			}
			if os.SameFile(stat, matchStat) {
				l.Debug("match found for device", "match", match, "device", device)
				return match, nil
			}
		}
	}

	l.Debug("no match found for device", "device", device)
	// if no match was found, we return the original device path
	return device, nil
}
