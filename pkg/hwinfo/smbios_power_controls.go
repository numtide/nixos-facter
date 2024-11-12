package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPowerControls captures system power controls information.
type SmbiosPowerControls struct {
	Type   SmbiosType `json:"-"`
	Handle int        `json:"handle"`
	Month  uint       `json:"month"`  // next scheduled power-on month
	Day    uint       `json:"day"`    // dto, day
	Hour   uint       `json:"hour"`   // dto, hour
	Minute uint       `json:"minute"` // dto, minute
	Second uint       `json:"second"` // dto, second
}

func (s SmbiosPowerControls) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosPower(info C.smbios_power_t) (*SmbiosPowerControls, error) {
	return &SmbiosPowerControls{
		Type:   SmbiosTypePowerControls,
		Handle: int(info.handle),
		Month:  uint(info.month),
		Day:    uint(info.day),
		Hour:   uint(info.hour),
		Minute: uint(info.minute),
		Second: uint(info.second),
	}, nil
}
