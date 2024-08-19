package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPointingDevice captures pointing device (aka 'mouse') information.
type SmbiosPointingDevice struct {
	Type      SmbiosType `json:"-"`
	Handle    int        `json:"handle"`
	MouseType *Id        `json:"mouse_type"`
	Interface *Id        `json:"interface"`
	Buttons   uint       `json:"buttons"`
}

func (s SmbiosPointingDevice) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMouse(info C.smbios_mouse_t) (Smbios, error) {
	return SmbiosPointingDevice{
		Type:      SmbiosTypePointingDevice,
		Handle:    int(info.handle),
		MouseType: NewId(info.mtype),
		Interface: NewId(info._interface),
		Buttons:   uint(info.buttons),
	}, nil
}
