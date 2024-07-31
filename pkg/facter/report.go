package facter

import (
	"fmt"

	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/virt"
)

type Report struct {
	Hardware       []*hwinfo.HardwareItem `json:"hardware"`
	Smbios         []hwinfo.Smbios        `json:"smbios,omitempty"`
	System         string                 `json:"system"`
	Virtualisation virt.Type              `json:"virtualisation"`
}

func GenerateReport(probes []hwinfo.ProbeFeature) (*Report, error) {
	report := Report{}

	if build.System == "" {
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	var err error
	report.Smbios, report.Hardware, err = hwinfo.Scan(probes)
	if err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	if report.Virtualisation, err = virt.Detect(); err != nil {
		return nil, fmt.Errorf("failed to detect virtualisation: %w", err)
	}

	return &report, nil
}
