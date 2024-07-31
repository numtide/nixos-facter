package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosCache captures processor information.
type SmbiosCache struct {
	Type          SmbiosType `json:"type"`
	Handle        int        `json:"handle"`
	Socket        string     `json:"socket"`       // socket designation
	SizeMax       uint       `json:"size_max"`     // max cache size in kbytes
	SizeCurrent   uint       `json:"size_current"` // current size in kbytes
	Speed         uint       `json:"speed"`        // cache speed in nanoseconds
	Mode          *Id        `json:"mode"`         // operational mode
	Enabled       bool       `json:"enabled"`
	Location      *Id        `json:"location"` // cache location
	Socketed      bool       `json:"socketed"`
	Level         uint       `json:"level"`         // cache level (0 = L1, 1 = L2, ...)
	ECC           *Id        `json:"ecc"`           // error correction type
	CacheType     *Id        `json:"cache_type"`    // logical cache type
	Associativity *Id        `json:"associativity"` // cache associativity
	SRAMType      []string   `json:"sram_type_current"`
	SRAMTypes     []string   `json:"sram_type_supported"`
}

func (s SmbiosCache) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosCache(info C.smbios_cache_t) (Smbios, error) {
	return SmbiosCache{
		Type:          SmbiosTypeCache,
		Handle:        int(info.handle),
		Socket:        C.GoString(info.socket),
		SizeMax:       uint(info.max_size),
		SizeCurrent:   uint(info.current_size),
		Speed:         uint(info.speed),
		Mode:          NewId(info.mode),
		Enabled:       uint(info.state) == 1,
		Location:      NewId(info.location),
		Socketed:      uint(info.socketed) == 1,
		Level:         uint(info.level),
		ECC:           NewId(info.ecc),
		CacheType:     NewId(info.cache_type),
		Associativity: NewId(info.assoc),
		SRAMType:      ReadStringList(info.sram.str),
		SRAMTypes:     ReadStringList(info.supp_sram.str),
	}, nil
}
