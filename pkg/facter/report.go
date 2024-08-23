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
	Hardware       Hardware  `json:"hardware,omitempty"`
	Smbios         Smbios    `json:"smbios,omitempty"`

	// Ephemeral entries
	Swap []*ephem.SwapEntry `json:"swap,omitempty"`
}

type Scanner struct {
	Swap      bool
	Ephemeral bool
	Features  []hwinfo.ProbeFeature
}

func (s *Scanner) Scan() (*Report, error) {
	var err error
	report := Report{
		Version: build.ReportVersion,
	}

	if build.System == "" {
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	var smbios []hwinfo.Smbios
	var devices []hwinfo.HardwareDevice

	smbios, devices, err = hwinfo.Scan(s.Features)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	for idx := range devices {
		if err = report.Hardware.add(devices[idx]); err != nil {
			return nil, fmt.Errorf("failed to add to hardware report: %w", err)
		}
	}

	for idx := range smbios {
		if err = report.Smbios.add(smbios[idx]); err != nil {
			return nil, fmt.Errorf("failed to add to smbios report: %w", err)
		}
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
