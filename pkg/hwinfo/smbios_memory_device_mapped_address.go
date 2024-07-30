package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryDeviceMappedAddress captures physical memory array information (consists of several memory devices).
type SmbiosMemoryDeviceMappedAddress struct {
	Type               SmbiosType `json:"type"`
	Handle             int        `json:"handle"`
	MemoryDeviceHandle int        `json:"memory_device_handle"`
	ArrayMapHandle     int        `json:"array_map_handle"`
	StartAddress       uint64     `json:"start_address"`
	EndAddress         uint64     `json:"end_address"`
	RowPosition        uint       `json:"row_position"`        // position of the referenced memory device in a row of the address partition
	InterleavePosition uint       `json:"interleave_position"` // dto, in an interleave
	InterleaveDepth    uint       `json:"interleave_depth"`    // number of consecutive rows
}

func (s SmbiosMemoryDeviceMappedAddress) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemDeviceMap(info C.smbios_memdevicemap_t) (Smbios, error) {
	return SmbiosMemoryDeviceMappedAddress{
		Type:               SmbiosTypeMemoryDeviceMappedAddress,
		Handle:             int(info.handle),
		MemoryDeviceHandle: int(info.memdevice_handle),
		ArrayMapHandle:     int(info.arraymap_handle),
		StartAddress:       uint64(info.start_addr),
		EndAddress:         uint64(info.end_addr),
		RowPosition:        uint(info.row_pos),
		InterleavePosition: uint(info.interleave_pos),
		InterleaveDepth:    uint(info.interleave_depth),
	}, nil
}
