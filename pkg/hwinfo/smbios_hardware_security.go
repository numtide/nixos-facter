package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosHardwareSecurity captures hardware security information.
type SmbiosHardwareSecurity struct {
	Type     SmbiosType `json:"type"`
	Handle   int        `json:"handle"`
	Power    *Id        `json:"power"`    // power-on password status
	Keyboard *Id        `json:"keyboard"` // keyboard password status
	Admin    *Id        `json:"admin"`    // admin password status
	Reset    *Id        `json:"reset"`    // front panel reset status
}

func (s SmbiosHardwareSecurity) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosSecure(info C.smbios_secure_t) (Smbios, error) {
	return SmbiosHardwareSecurity{
		Type:     SmbiosTypeHardwareSecurity,
		Handle:   int(info.handle),
		Power:    NewId(info.power),
		Keyboard: NewId(info.keyboard),
		Admin:    NewId(info.admin),
		Reset:    NewId(info.reset),
	}, nil
}
