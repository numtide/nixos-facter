package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoDsl struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:"db_entry_0,omitempty"`
	DbEntry1 []string `json:"db_entry_1,omitempty"`

	Mode string `json:"mode,omitempty"` // DSL driver types
	Name string `json:"name,omitempty"` // DSL driver name
}

func (d DriverInfoDsl) DriverInfoType() DriverInfoType {
	return DriverInfoTypeDsl
}

func NewDriverInfoDsl(info C.driver_info_dsl_t) DriverInfoDsl {
	return DriverInfoDsl{
		Type:     DriverInfoTypeDsl,
		DbEntry0: ReadStringList(info.hddb0),
		DbEntry1: ReadStringList(info.hddb1),
		Mode:     C.GoString(info.mode),
		Name:     C.GoString(info.name),
	}
}
