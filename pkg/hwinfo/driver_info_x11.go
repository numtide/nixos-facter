package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool driver_info_x11_supports_3d(driver_info_x11_t info) { return info.x3d; }
char driver_info_x11_colors_all(driver_info_x11_t info) { return info.colors.all; }
char driver_info_x11_colors_c8(driver_info_x11_t info) { return info.colors.c8; }
char driver_info_x11_colors_c15(driver_info_x11_t info) { return info.colors.c15; }
char driver_info_x11_colors_c16(driver_info_x11_t info) { return info.colors.c16; }
char driver_info_x11_colors_c24(driver_info_x11_t info) { return info.colors.c24; }
char driver_info_x11_colors_c32(driver_info_x11_t info) { return info.colors.c32; }
*/
import "C"

type DriverInfoX11 struct {
	Type DriverInfoType `json:",omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:",omitempty"`
	DbEntry1 []string `json:",omitempty"`

	Server      string `json:",omitempty"` // the server/module name
	XF86Version string `json:",omitempty"` // XFree86 version (3 or 4)
	Supports3D  bool   `json:""`           // has 3D support
	Colors      struct {
		// the next 5 entries combined
		All byte `json:""`
		C8  byte `json:""`
		C15 byte `json:""`
		C16 byte `json:""`
		C24 byte `json:""`
		C32 byte `json:""`
	} `json:",omitempty"`
	DacSpeed   uint     `json:""`           // max. ramdac clock
	Extensions []string `json:",omitempty"` // additional X extensions to load ('Module' section)
	Options    []string `json:",omitempty"` // special server options
	Raw        []string `json:",omitempty"` // extra info to add to XF86Config
	Script     string   `json:","`          // 3d script to run
}

func (d DriverInfoX11) DriverInfoType() DriverInfoType {
	return DriverInfoTypeX11
}

func NewDriverInfoX11(info C.driver_info_x11_t) DriverInfoX11 {
	result := DriverInfoX11{
		Type:        DriverInfoTypeX11,
		DbEntry0:    ReadStringList(info.hddb0),
		DbEntry1:    ReadStringList(info.hddb1),
		Server:      C.GoString(info.server),
		XF86Version: C.GoString(info.xf86_ver),
		Supports3D:  bool(C.driver_info_x11_supports_3d(info)),
		DacSpeed:    uint(info.dacspeed),
		Extensions:  ReadStringList(info.extensions),
		Options:     ReadStringList(info.options),
		Raw:         ReadStringList(info.raw),
		Script:      C.GoString(info.script),
	}

	result.Colors.All = byte(C.driver_info_x11_colors_all(info))
	result.Colors.C8 = byte(C.driver_info_x11_colors_c8(info))
	result.Colors.C15 = byte(C.driver_info_x11_colors_c15(info))
	result.Colors.C16 = byte(C.driver_info_x11_colors_c16(info))
	result.Colors.C24 = byte(C.driver_info_x11_colors_c24(info))
	result.Colors.C32 = byte(C.driver_info_x11_colors_c32(info))

	return result
}
