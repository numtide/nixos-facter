package facter

import (
	"fmt"
	"slices"

	"github.com/numtide/nixos-facter/pkg/virt"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/build"
)


type Report struct {
	Virtualisation virt.Type              `json:"virtualisation"`
	Hardware       []*hwinfo.HardwareItem `json:"hardware"`
	System         string                 `json:"system"`
}

func (r *Report) AddHardwareItem(item *hwinfo.HardwareItem) {
	r.Hardware = append(r.Hardware, item)

	// canonically sort by device index
	slices.SortFunc(r.Hardware, func(a, b *hwinfo.HardwareItem) int {
		return int(a.Index) - int(b.Index)
	})
}

func GenerateReport() (*Report, error) {
	report := Report{}

	if build.System == "" {
		return nil, fmt.Errorf("system is not set")
	}
	report.System = build.System

	if err := hwinfo.Scan(func(item *hwinfo.HardwareItem) error {
		report.AddHardwareItem(item)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	var err error
	if report.Virtualisation, err = virt.Detect(); err != nil {
		return nil, fmt.Errorf("failed to detect virtualisation: %w", err)
	}

	return &report, nil
}
