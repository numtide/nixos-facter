package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosMemoryDevice captures system slot information.
type SmbiosMemoryDevice struct {
	Type              SmbiosType `json:"-"`
	Handle            int        `json:"handle"`
	Location          string     `json:"location"`      // device location
	BankLocation      string     `json:"bank_location"` // bank location
	Manufacturer      string     `json:"manufacturer"`
	Serial            string     `json:"-"` // omit from json
	AssetTag          string     `json:"asset_tag,omitempty"`
	PartNumber        string     `json:"part_number"`
	ArrayHandle       int        `json:"array_handle"` // memory array this device belongs to
	ErrorHandle       int        `json:"error_handle"` // points to error info record; 0xfffe: not supported, 0xffff: no error
	Width             uint       `json:"width"`        // data width in bits
	ECCBits           uint       `json:"ecc_bits"`     // ecc bits
	Size              uint       `json:"size"`         // kB
	FormFactor        *Id        `json:"form_factor"`
	Set               uint       `json:"set"` // 0: does not belong to a set; 1-0xfe: set number; 0xff: unknown
	MemoryType        *Id        `json:"memory_type"`
	MemoryTypeDetails []string   `json:"memory_type_details"`
	Speed             uint       `json:"speed"` // MHz
}

func (s SmbiosMemoryDevice) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosMemDevice(info C.smbios_memdevice_t) (Smbios, error) {
	return SmbiosMemoryDevice{
		Type:              SmbiosTypeMemoryDevice,
		Handle:            int(info.handle),
		Location:          C.GoString(info.location),
		BankLocation:      C.GoString(info.bank),
		Manufacturer:      C.GoString(info.manuf),
		Serial:            C.GoString(info.serial),
		AssetTag:          C.GoString(info.asset),
		PartNumber:        C.GoString(info.part),
		ArrayHandle:       int(info.array_handle),
		ErrorHandle:       int(info.error_handle),
		Width:             uint(info.width),
		ECCBits:           uint(info.eccbits),
		Size:              uint(info.size),
		FormFactor:        NewId(info.form),
		Set:               uint(info.set),
		MemoryType:        NewId(info.mem_type),
		MemoryTypeDetails: ReadStringList(info.type_detail.str),
		Speed:             uint(info.speed),
	}, nil
}
