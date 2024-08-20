package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "fmt"

// SmbiosChassis captures motherboard related information.
type SmbiosChassis struct {
	Type          SmbiosType `json:"-"`
	Handle        int        `json:"handle"`
	Manufacturer  string     `json:"manufacturer"`
	Version       string     `json:"version"`
	Serial        string     `json:"-"`                   // omit from json output
	AssetTag      string     `json:"asset_tag,omitempty"` // asset tag
	ChassisType   *Id        `json:"chassis_type"`
	LockPresent   bool       `json:"lock_present"` // true: lock present, false: not present or unknown
	BootupState   *Id        `json:"bootup_state"`
	PowerState    *Id        `json:"power_state"`    // power supply state (at last boot)
	ThermalState  *Id        `json:"thermal_state"`  // thermal state (at last boot)
	SecurityState *Id        `json:"security_state"` // security state (at last boot)
	OEM           string     `json:"oem"`            // oem-specific information"
}

func (s SmbiosChassis) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosChassis(info C.smbios_chassis_t) (Smbios, error) {
	return SmbiosChassis{
		Type:          SmbiosTypeChassis,
		Handle:        int(info.handle),
		Manufacturer:  C.GoString(info.manuf),
		Version:       C.GoString(info.version),
		Serial:        C.GoString(info.serial),
		AssetTag:      C.GoString(info.asset),
		ChassisType:   NewId(info.ch_type),
		LockPresent:   uint(info.lock) == 1,
		BootupState:   NewId(info.bootup),
		PowerState:    NewId(info.power),
		ThermalState:  NewId(info.thermal),
		SecurityState: NewId(info.security),
		OEM:           fmt.Sprintf("0x%x", uint(info.oem)),
	}, nil
}
