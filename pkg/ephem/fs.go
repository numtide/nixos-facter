package ephem

import (
	"encoding/json"
	"regexp"
	"strconv"
	"strings"

	"github.com/moby/sys/mountinfo"
)

var specialFsRegex = regexp.MustCompile(`^(/proc|/dev|/sys|/run|/var/lib/docker|/var/lib/nfs/rpc_pipefs).*`)

// MountInfo represents the information about a mount point.
// It is just a type alias for mountinfo.Info to allow us to add JSON marshalling.
type MountInfo struct {
	mountinfo.Info
}

func (i MountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]string{
		"id":              strconv.Itoa(i.ID),
		"parent_id":       strconv.Itoa(i.Parent),
		"major":           strconv.Itoa(i.Major),
		"minor":           strconv.Itoa(i.Minor),
		"root":            i.Root,
		"mount_point":     i.Mountpoint,
		"mount_options":   i.Options,
		"optional":        i.Optional,
		"filesystem_type": i.FSType,
		"mount_source":    i.Source,
		"super_options":   i.VFSOptions,
	})
}

func Mounts() ([]*MountInfo, error) {
	info, err := mountinfo.GetMounts(mountFilter)
	if err != nil {
		return nil, err
	}
	var result []*MountInfo
	for idx := range info {
		result = append(result, &MountInfo{Info: *info[idx]})
	}
	return result, nil
}

func mountFilter(info *mountinfo.Info) (skip, stop bool) {
	if skip = specialFsRegex.MatchString(info.Mountpoint); skip {
		return true, false
	}

	// skip the read-only bind-mount on /nix/store
	// https://github.com/NixOS/nixpkgs/blob/dac9cdf8c930c0af98a63cbfe8005546ba0125fb/nixos/modules/installer/tools/nixos-generate-config.pl#L395
	if info.Mountpoint == "/nix/store" && strings.Contains(info.VFSOptions, "rw") && strings.Contains(info.Options, "ro") {
		return true, false
	}

	return false, false
}
