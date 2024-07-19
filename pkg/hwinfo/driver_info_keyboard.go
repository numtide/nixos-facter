package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoKeyboard struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:"db_entry_0,omitempty"`
	DbEntry1 []string `json:"db_entry_1,omitempty"`

	XkbRules  string `json:"xkb_rules,omitempty"`
	XkbModel  string `json:"xkb_model,omitempty"`
	XkbLayout string `json:"xkb_layout,omitempty"`
	Keymap    string `json:"keymap,omitempty"`
}

func (d DriverInfoKeyboard) DriverInfoType() DriverInfoType {
	return DriverInfoTypeKeyboard
}

func NewDriverInfoKeyboard(info C.driver_info_kbd_t) DriverInfoKeyboard {
	return DriverInfoKeyboard{
		Type:      DriverInfoTypeKeyboard,
		DbEntry0:  ReadStringList(info.hddb0),
		DbEntry1:  ReadStringList(info.hddb1),
		XkbRules:  C.GoString(info.XkbRules),
		XkbModel:  C.GoString(info.XkbModel),
		XkbLayout: C.GoString(info.XkbLayout),
		Keymap:    C.GoString(info.keymap),
	}
}
