package facter

import (
	"fmt"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

type Smbios struct {
	Bios                      *hwinfo.SmbiosBios     `json:"bios,omitempty"`
	Board                     *hwinfo.SmbiosBoard    `json:"board,omitempty"`
	Cache                     []hwinfo.Smbios        `json:"cache,omitempty"`
	Chassis                   *hwinfo.SmbiosChassis  `json:"chassis,omitempty"`
	Config                    *hwinfo.SmbiosConfig   `json:"config,omitempty"`
	GroupAssociations         []hwinfo.Smbios        `json:"group_associations,omitempty"`
	HardwareSecurity          []hwinfo.Smbios        `json:"hardware_security,omitempty"`
	Language                  *hwinfo.SmbiosLanguage `json:"language,omitempty"`
	Memory64Error             []hwinfo.Smbios        `json:"memory_64_error,omitempty"`
	MemoryArray               []hwinfo.Smbios        `json:"memory_array,omitempty"`
	MemoryArrayMappedAddress  []hwinfo.Smbios        `json:"memory_array_mapped_address,omitempty"`
	MemoryDevice              []hwinfo.Smbios        `json:"memory_device,omitempty"`
	MemoryDeviceMappedAddress []hwinfo.Smbios        `json:"memory_device_mapped_address,omitempty"`
	MemoryError               []hwinfo.Smbios        `json:"memory_error,omitempty"`
	Onboard                   []hwinfo.Smbios        `json:"onboard,omitempty"`
	PointingDevice            []hwinfo.Smbios        `json:"pointing_device,omitempty"`
	PortConnector             []hwinfo.Smbios        `json:"port_connector,omitempty"`
	PowerControls             []hwinfo.Smbios        `json:"power_controls,omitempty"`
	Processor                 []hwinfo.Smbios        `json:"processor,omitempty"`
	Slot                      []hwinfo.Smbios        `json:"slot,omitempty"`
	System                    *hwinfo.SmbiosSystem   `json:"system,omitempty"`
}

func (s *Smbios) add(item hwinfo.Smbios) error {
	switch item.SmbiosType() {
	case hwinfo.SmbiosTypeBios:
		if s.Bios != nil {
			return fmt.Errorf("bios field is already set")
		} else if bios, ok := item.(hwinfo.SmbiosBios); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosBios, found %T", item)
		} else {
			s.Bios = &bios
		}
	case hwinfo.SmbiosTypeBoard:
		if s.Board != nil {
			return fmt.Errorf("board field is already set")
		} else if board, ok := item.(hwinfo.SmbiosBoard); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosBoard, found %T", item)
		} else {
			s.Board = &board
		}
	case hwinfo.SmbiosTypeCache:
		s.Cache = append(s.Cache, item)
	case hwinfo.SmbiosTypeChassis:
		if s.Chassis != nil {
			return fmt.Errorf("chassis field is already set")
		} else if chassis, ok := item.(hwinfo.SmbiosChassis); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosChassis, found %T", item)
		} else {
			s.Chassis = &chassis
		}
	case hwinfo.SmbiosTypeConfig:
		if s.Config != nil {
			return fmt.Errorf("config field is already set")
		} else if config, ok := item.(hwinfo.SmbiosConfig); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosConfig, found %T", item)
		} else {
			s.Config = &config
		}
	case hwinfo.SmbiosTypeGroupAssociations:
		s.GroupAssociations = append(s.GroupAssociations, item)
	case hwinfo.SmbiosTypeHardwareSecurity:
		s.GroupAssociations = append(s.GroupAssociations, item)
	case hwinfo.SmbiosTypeLanguage:
		if s.Language != nil {
			return fmt.Errorf("language field is already set")
		} else if language, ok := item.(hwinfo.SmbiosLanguage); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosLanguage, found %T", item)
		} else {
			s.Language = &language
		}
	case hwinfo.SmbiosTypeMemory64Error:
		s.Memory64Error = append(s.Memory64Error, item)
	case hwinfo.SmbiosTypeMemoryArray:
		s.MemoryArray = append(s.MemoryArray, item)
	case hwinfo.SmbiosTypeMemoryArrayMappedAddress:
		s.MemoryArrayMappedAddress = append(s.MemoryArrayMappedAddress, item)
	case hwinfo.SmbiosTypeMemoryDevice:
		s.MemoryDevice = append(s.MemoryDevice, item)
	case hwinfo.SmbiosTypeMemoryDeviceMappedAddress:
		s.MemoryDeviceMappedAddress = append(s.MemoryDeviceMappedAddress, item)
	case hwinfo.SmbiosTypeMemoryError:
		s.MemoryError = append(s.MemoryError, item)
	case hwinfo.SmbiosTypeOnboard:
		s.Onboard = append(s.Onboard, item)
	case hwinfo.SmbiosTypePointingDevice:
		s.PointingDevice = append(s.PointingDevice, item)
	case hwinfo.SmbiosTypePortConnector:
		s.PortConnector = append(s.PortConnector, item)
	case hwinfo.SmbiosTypePowerControls:
		s.PowerControls = append(s.PowerControls, item)
	case hwinfo.SmbiosTypeProcessor:
		s.Processor = append(s.Processor, item)
	case hwinfo.SmbiosTypeSlot:
		s.Slot = append(s.Slot, item)
	case hwinfo.SmbiosTypeSystem:
		if s.System != nil {
			return fmt.Errorf("system field is already set")
		} else if system, ok := item.(hwinfo.SmbiosSystem); !ok {
			return fmt.Errorf("expected hwinfo.SmbiosSystem, found %T", item)
		} else {
			s.System = &system
		}
	default:
		// Do nothing for the rest of the types, we currently don't map them
	}

	return nil
}
