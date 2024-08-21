package facter

import (
	"encoding/json"
	"fmt"

	"github.com/numtide/nixos-facter/pkg/ephem"

	"github.com/numtide/nixos-facter/pkg/build"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/numtide/nixos-facter/pkg/virt"
)

type HardwareMap map[hwinfo.HardwareClass][]*hwinfo.HardwareDevice

func (h HardwareMap) MarshalJSON() ([]byte, error) {
	values := make(map[string]any)

	// We iterate through each hardware class, making some changes to the structure as appropriate before marshalling
	// to JSON.
	for class, list := range h {
		switch class {

		case hwinfo.HardwareClassBios, hwinfo.HardwareClassSystem:
			// we know there's only one entry for these classes, so we ensure the entry exists and flatten the list
			switch len(list) {
			case 0:
				// no entries, continue
				continue
			case 1:
				// we found the entry we know to be there, flatten the value
				values[class.String()] = list[0]
			default:
				// this should not happen
				return nil, fmt.Errorf("unexpected number of entries %d for hardware class %s, expected 1", len(list), class.String())
			}

		case hwinfo.HardwareClassCpu:
			// For cpu entries, we group by physical id and take the last entry.
			// We can do this because most of the fields are the same across cores and virtual threads, even when dealing
			// with performance and efficiency cores on the same die.
			// The fields that are core-specific have already been filtered out from the JSON output in the underlying
			// model.
			// This logic only really applies to x86 for now.
			// There's nothing in the kernel which writes out physical id for aarch64 or riscv.
			// In those cases, we just take the first entry.

			if len(list) == 0 {
				// empty list, continue
				continue
			}

			// Determine the architecture we're dealing with by looking at the first entry.
			first := list[0].Detail.(hwinfo.DetailCpu)

			switch first.Architecture {
			case hwinfo.CpuArchx86_64:
				// Group based on physical id and reduce the value to one entry per processor.
				byPhysicalId := make(map[uint]hwinfo.DetailCpu)
				for idx := range list {
					detail := list[idx].Detail.(hwinfo.DetailCpu)
					byPhysicalId[detail.PhysicalId] = detail
				}

				var summary []hwinfo.DetailCpu
				for _, v := range byPhysicalId {
					summary = append(summary, v)
				}

				values[class.String()] = summary

				break

			default:
				// take the first entry as a summary of all cores
				values[class.String()] = first
			}

		default:
			// For all other hardware classes, we just leave the value as is, converting the enum key to a string.
			values[class.String()] = list
		}
	}

	// Marshal to JSON.
	return json.Marshal(values)
}

type Report struct {
	// monotonically increasing number, used to indicate breaking changes or new features in the report output
	Version        uint      `json:"version"`
	System         string    `json:"system"`
	Virtualisation virt.Type `json:"virtualisation"`

	// grouped by hardware class
	Hardware HardwareMap `json:"hardware,omitempty"`

	// grouped by SmbiosType
	// todo flatten some entries like with hardware
	Smbios map[string][]hwinfo.Smbios `json:"smbios,omitempty"`

	// Ephemeral entries
	Swap []*ephem.SwapEntry `json:"swap,omitempty"`
}

func (r Report) addDevice(device *hwinfo.HardwareDevice) {
	key := device.HardwareClass
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
		Hardware: make(HardwareMap),
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
