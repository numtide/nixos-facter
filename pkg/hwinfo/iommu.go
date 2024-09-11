package hwinfo

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const (
	iommuSysfsPath = "/sys/kernel/iommu_groups"
)

type IOMMUGroup int

func IOMMUGroups() (map[string]IOMMUGroup, error) {
	groups, err := os.ReadDir(iommuSysfsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read iommu groups from sysfs: %w", err)
	}

	result := make(map[string]IOMMUGroup)
	for _, group := range groups {

		if !group.IsDir() {
			return nil, fmt.Errorf("non-directory entry found in %s: %s", iommuSysfsPath, group.Name())
		}

		groupId, err := strconv.Atoi(group.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to parse iommu group id '%s': %w", group.Name(), err)
		}

		devicesPath := filepath.Join(iommuSysfsPath, group.Name(), "devices")

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

			// sysfs id -> iommu group id
			result[resolvedPath[4:]] = IOMMUGroup(groupId)
		}
	}

	return result, nil
}
