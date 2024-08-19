package facter

import (
	"fmt"

	"github.com/numtide/nixos-facter/pkg/ephem"

	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/virt"
)

type Report struct {
	// monotonically increasing number, used to indicate breaking changes or new features in the report output
	Version        uint      `json:"version"`
	System         string    `json:"system"`
	Virtualisation virt.Type `json:"virtualisation"`

	Smbios []hwinfo.Smbios    `json:"smbios,omitempty"`
	Swap   []*ephem.SwapEntry `json:"swap,omitempty"`

	// grouped by hardware class
	Hardware map[string][]*hwinfo.HardwareDevice `json:"hardware,omitempty"`
}

func (r Report) addDevice(device *hwinfo.HardwareDevice) {
	key := device.HardwareClass.String()
	r.Hardware[key] = append(r.Hardware[key], device)
}

type Scanner struct {
	Swap      bool
	Ephemeral bool
	Features  []hwinfo.ProbeFeature
}

func (s *Scanner) Scan() (*Report, error) {
	var err error
	report := Report{
		Version:  build.ReportVersion,
		Hardware: make(map[string][]*hwinfo.HardwareDevice),
	}

	if build.System == "" {
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	var devices []*hwinfo.HardwareDevice
	report.Smbios, devices, err = hwinfo.Scan(s.Features)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	for idx := range devices {
		report.addDevice(devices[idx])
	}

	if report.Virtualisation, err = virt.Detect(); err != nil {
		return nil, fmt.Errorf("failed to detect virtualisation: %w", err)
	}

	if s.Ephemeral || s.Swap {
		if report.Swap, err = ephem.SwapEntries(); err != nil {
			return nil, fmt.Errorf("failed to detect swap devices: %w", err)
		}
	}

	return &report, nil
}
