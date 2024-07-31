package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryError captures 32-bit memory error information.
type SmbiosMemoryError struct {
	Type          SmbiosType `json:"type"`
	Handle        int        `json:"handle"`
	ErrorType     *Id        `json:"error_type"`     // error type memory
	Granularity   *Id        `json:"granularity"`    // memory array or memory partition
	Operation     *Id        `json:"operation"`      // mem operation causing the rror
	Syndrome      uint       `json:"syndrome"`       // vendor-specific ECC syndrome; 0: unknown
	ArrayAddress  uint       `json:"array_address"`  // fault address relative to mem array; 0x80000000: unknown
	DeviceAddress uint       `json:"device_address"` // fault address relative to mem array; 0x80000000: unknown
	Range         uint       `json:"range"`          // range, within which the error can be determined; 0x80000000: unknown
}

func (s SmbiosMemoryError) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemError(info C.smbios_memerror_t) (Smbios, error) {
	return SmbiosMemoryError{
		Type:          SmbiosTypeMemoryError,
		Handle:        int(info.handle),
		ErrorType:     NewId(info.err_type),
		Granularity:   NewId(info.granularity),
		Operation:     NewId(info.operation),
		Syndrome:      uint(info.syndrome),
		ArrayAddress:  uint(info.array_addr),
		DeviceAddress: uint(info.device_addr),
		Range:         uint(info._range),
	}, nil
}
