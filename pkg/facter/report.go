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

	// grouped by hardware class
	Hardware map[string][]*hwinfo.HardwareDevice `json:"hardware,omitempty"`

	// grouped by SmbiosType
	Smbios map[string][]hwinfo.Smbios `json:"smbios,omitempty"`

	// Ephemeral entries
	Swap []*ephem.SwapEntry `json:"swap,omitempty"`
}

func (r Report) addDevice(device *hwinfo.HardwareDevice) {
	key := device.HardwareClass.String()
	r.Hardware[key] = append(r.Hardware[key], device)
}

func (r Report) addSmbios(smbios hwinfo.Smbios) {
	key := smbios.SmbiosType().String()
	r.Smbios[key] = append(r.Smbios[key], smbios)
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
		Smbios:   make(map[string][]hwinfo.Smbios),
		Hardware: make(map[string][]*hwinfo.HardwareDevice),
	}

	if build.System == "" {
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	var smbios []hwinfo.Smbios
	var devices []*hwinfo.HardwareDevice

	smbios, devices, err = hwinfo.Scan(s.Features)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	for idx := range devices {
		report.addDevice(devices[idx])
	}

	for idx := range smbios {
		report.addSmbios(smbios[idx])
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
