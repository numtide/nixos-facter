package hwinfo

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	iommuSysfsPath = "/sys/kernel/iommu_groups"
)

func IOMMUGroups() (map[string]string, error) {
	groups, err := os.ReadDir(iommuSysfsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read iommu groups from sysfs: %w", err)
	}

	result := make(map[string]string)
	for _, group := range groups {

		if !group.IsDir() {
			return nil, fmt.Errorf("non-directory entry found in %s: %s", iommuSysfsPath, group.Name())
		}

		groupPath := filepath.Join(iommuSysfsPath, group.Name())
		devicesPath := filepath.Join(groupPath, "devices")

		devices, err := os.ReadDir(devicesPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read devices from iommu group %s: %w", devicesPath, err)
		}
		for _, device := range devices {
			devicePath := filepath.Join(devicesPath, device.Name())
			resolvedPath, err := filepath.EvalSymlinks(devicePath)
			if err != nil {
				return nil, fmt.Errorf("failed to resolve device symlink '%s': %w", devicePath, err)
			}

			// sysfs id -> iommu group
			// we trim leading '/sys'
			result[resolvedPath[4:]] = groupPath[4:]
		}
	}

	return result, nil
}
