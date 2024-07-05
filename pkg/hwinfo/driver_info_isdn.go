package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoIsdn struct {
	Type DriverInfoType `json:",omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:",omitempty"`
	DbEntry1 []string `json:",omitempty"`

	I4lType    int    `json:""`
	I4lSubtype int    `json:""`
	I4lName    string `json:",omitempty"`
	// todo isdn params
}

func (d DriverInfoIsdn) DriverInfoType() DriverInfoType {
	return DriverInfoTypeIsdn
}

func NewDriverInfoIsdn(info C.driver_info_isdn_t) DriverInfoIsdn {
	return DriverInfoIsdn{
		Type:       DriverInfoTypeIsdn,
		DbEntry0:   ReadStringList(info.hddb0),
		DbEntry1:   ReadStringList(info.hddb1),
		I4lType:    int(info.i4l_type),
		I4lSubtype: int(info.i4l_subtype),
		I4lName:    C.GoString(info.i4l_name),
	}
}
