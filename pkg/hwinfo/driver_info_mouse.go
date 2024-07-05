package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoMouse struct {
	Type DriverInfoType `json:",omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:",omitempty"`
	DbEntry1 []string `json:",omitempty"`

	XF86    string `json:",omitempty"` // XF86 protocol name
	GPM     string `json:",omitempty"` // dto, gpm
	Buttons int    `json:",omitempty"` // number of buttons, -1 -> unknown
	Wheels  int    `json:",omitempty"` // dto, wheels
}

func (d DriverInfoMouse) DriverInfoType() DriverInfoType {
	return DriverInfoTypeMouse
}

func NewDriverInfoMouse(info C.driver_info_mouse_t) DriverInfoMouse {
	return DriverInfoMouse{
		Type:     DriverInfoTypeMouse,
		DbEntry0: ReadStringList(info.hddb0),
		DbEntry1: ReadStringList(info.hddb1),
		XF86:     C.GoString(info.xf86),
		GPM:      C.GoString(info.gpm),
		Buttons:  int(info.buttons),
		Wheels:   int(info.wheels),
	}
}
