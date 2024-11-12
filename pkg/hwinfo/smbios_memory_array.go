package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryArray captures physical memory array information (consists of several memory devices).
type SmbiosMemoryArray struct {
	Type        SmbiosType `json:"-"`
	Handle      int        `json:"handle"`
	Location    *ID        `json:"location"`     // memory device location
	Usage       *ID        `json:"usage"`        // memory usage
	ECC         *ID        `json:"ecc"`          // ECC types
	MaxSize     uint       `json:"max_size"`     // max memory size in KB
	ErrorHandle int        `json:"error_handle"` // points to error info record; 0xfffe: not supported, 0xffff: no error
	Slots       uint       `json:"slots"`        // slots or sockets for this device
}

func (s SmbiosMemoryArray) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemArray(info C.smbios_memarray_t) (*SmbiosMemoryArray, error) {
	return &SmbiosMemoryArray{
		Type:        SmbiosTypeMemoryArray,
		Handle:      int(info.handle),
		Location:    NewID(info.location),
		Usage:       NewID(info.use),
		ECC:         NewID(info.ecc),
		MaxSize:     uint(info.max_size),
		ErrorHandle: int(info.error_handle),
		Slots:       uint(info.slots),
	}, nil
}
