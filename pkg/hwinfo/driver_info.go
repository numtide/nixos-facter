package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

enum driver_info_type driver_info_get_type(driver_info_t *info) { return info->any.type; }
driver_info_module_t driver_info_get_module(driver_info_t *info) { return info->module; }
driver_info_mouse_t driver_info_get_mouse(driver_info_t *info) { return info->mouse; }
driver_info_display_t driver_info_get_display(driver_info_t *info) { return info->display; }
driver_info_kbd_t driver_info_get_kbd(driver_info_t *info) { return info->kbd; }
driver_info_dsl_t driver_info_get_dsl(driver_info_t *info) { return info->dsl; }
driver_info_isdn_t driver_info_get_isdn(driver_info_t *info) { return info->isdn; }
driver_info_x11_t driver_info_get_x11(driver_info_t *info) { return info->x11; }
*/
import "C"

import (
	"errors"
)

//go:generate enumer -type=DriverInfoType -json -transform=snake -trimprefix DriverInfoType -output=./driver_info_enum_type.go
type DriverInfoType uint

const (
	DriverInfoTypeAny DriverInfoType = iota
	DriverInfoTypeDisplay
	DriverInfoTypeModule
	DriverInfoTypeMouse
	DriverInfoTypeX11
	DriverInfoTypeIsdn
	DriverInfoTypeKeyboard
	DriverInfoTypeDsl
)

type DriverInfo interface {
	DriverInfoType() DriverInfoType
}

func NewDriverInfo(info *C.driver_info_t) (DriverInfo, error) {
	if info == nil {
		return nil, nil
	}
	infoType := DriverInfoType(C.driver_info_get_type(info))
	switch infoType {
	case DriverInfoTypeModule:
		return NewDriverInfoModule(C.driver_info_get_module(info)), nil
	case DriverInfoTypeMouse:
		return NewDriverInfoMouse(C.driver_info_get_mouse(info)), nil
	case DriverInfoTypeDisplay:
		return NewDriverInfoDisplay(C.driver_info_get_display(info)), nil
	case DriverInfoTypeKeyboard:
		return NewDriverInfoKeyboard(C.driver_info_get_kbd(info)), nil
	case DriverInfoTypeDsl:
		return NewDriverInfoDsl(C.driver_info_get_dsl(info)), nil
	case DriverInfoTypeIsdn:
		return NewDriverInfoIsdn(C.driver_info_get_isdn(info)), nil
	case DriverInfoTypeX11:
		return NewDriverInfoX11(C.driver_info_get_x11(info)), nil
	default:
		return nil, errors.New("unknown driver info type")
	}
}
