package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryArrayMappedAddress captures physical memory array information (consists of several memory devices).
type SmbiosMemoryArrayMappedAddress struct {
	Type         SmbiosType `json:"-"`
	Handle       int        `json:"handle"`
	ArrayHandle  int        `json:"array_handle"`  // memory array this mapping belongs to
	StartAddress uint64     `json:"start_address"` // memory range start address
	EndAddress   uint64     `json:"end_address"`   // end address
	PartWidth    uint       `json:"part_width"`    // number of memory devices
}

func (s SmbiosMemoryArrayMappedAddress) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemArrayMap(info C.smbios_memarraymap_t) (Smbios, error) {
	return SmbiosMemoryArrayMappedAddress{
		Type:         SmbiosTypeMemoryArrayMappedAddress,
		Handle:       int(info.handle),
		ArrayHandle:  int(info.array_handle),
		StartAddress: uint64(info.start_addr),
		EndAddress:   uint64(info.end_addr),
		PartWidth:    uint(info.part_width),
	}, nil
}
