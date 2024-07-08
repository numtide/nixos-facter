package facter

import (
	"encoding/json"
	"fmt"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"slices"
)

type Report struct {
	// Hardware is a map of HardwareClass (in string form) to a list of HardwareItem entries.
	Hardware map[hwinfo.HardwareClass][]*hwinfo.HardwareItem `json:"hardware"`
}

func (r *Report) MarshalJSON() ([]byte, error) {
	// convert enum keys to strings
	hardware := make(map[string][]*hwinfo.HardwareItem, len(r.Hardware))
	for key, value := range r.Hardware {
		hardware[key.String()] = value
	}

	// note: when golang marshals a map, the keys are serialized lexicographically https://go.dev/src/encoding/json/encode.go#L359
	// thus allowing for a stable output when combined with sorting of the hardware items
	return json.Marshal(&struct {
		Hardware map[string][]*hwinfo.HardwareItem `json:"hardware"`
	}{
		Hardware: hardware,
	})
}

func (r *Report) AddHardwareItem(item *hwinfo.HardwareItem) {
	key := item.HardwareClass
	r.Hardware[key] = append(r.Hardware[key], item)

	// canonically sort by device index
	slices.SortFunc(r.Hardware[key], func(a, b *hwinfo.HardwareItem) int {
		return int(a.Index) - int(b.Index)
	})
}

func GenerateReport() (*Report, error) {
	report := Report{}
	report.Hardware = make(map[hwinfo.HardwareClass][]*hwinfo.HardwareItem)

	if err := hwinfo.Scan(func(item *hwinfo.HardwareItem) error {
		report.AddHardwareItem(item)
		return nil
	}); err != nil {
		return nil, fmt.Errorf("failed to scan hardware: %w", err)
	}

	return &report, nil
}
