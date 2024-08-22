package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_detail_type_t hd_detail_get_type(hd_detail_t *det) { return det->type; }
hd_detail_pci_t hd_detail_get_pci(hd_detail_t *det) { return det->pci; }
hd_detail_usb_t hd_detail_get_usb(hd_detail_t *det) { return det->usb; }
hd_detail_isapnp_t hd_detail_get_isapnp(hd_detail_t *det) { return det->isapnp; }
hd_detail_cpu_t hd_detail_get_cpu(hd_detail_t *det) { return det->cpu; }
hd_detail_monitor_t hd_detail_get_monitor(hd_detail_t *det) { return det->monitor; }
hd_detail_bios_t hd_detail_get_bios(hd_detail_t *det) { return det->bios; }
hd_detail_sys_t hd_detail_get_sys(hd_detail_t *det) { return det->sys; }

*/
import "C"

import (
	"encoding/hex"
	"unsafe"
)

//go:generate enumer -type=DetailType -json -transform=snake -trimprefix DetailType -output=./detail_enum_type.go
type DetailType uint

const (
	DetailTypePci DetailType = iota
	DetailTypeUsb
	DetailTypeIsaPnp
	DetailTypeCdrom

	DetailTypeFloppy
	DetailTypeBios
	DetailTypeCpu
	DetailTypeProm

	DetailTypeMonitor
	DetailTypeSys
	DetailTypeScsi
	DetailTypeDevtree

	DetailTypeCcw
	DetailTypeJoystick
)

type Detail interface {
	DetailType() DetailType
}

func NewDetail(detail *C.hd_detail_t) (Detail, error) {
	if detail == nil {
		return nil, nil
	}

	detailType := DetailType(C.hd_detail_get_type(detail))

	switch detailType {
	case DetailTypePci:
		return NewDetailPci(C.hd_detail_get_pci(detail))
	case DetailTypeUsb:
		return NewDetailUsb(C.hd_detail_get_usb(detail))
	case DetailTypeIsaPnp:
		return NewDetailIsaPnpDevice(C.hd_detail_get_isapnp(detail))
	case DetailTypeCpu:
		return NewDetailCpu(C.hd_detail_get_cpu(detail))
	case DetailTypeMonitor:
		return NewDetailMonitor(C.hd_detail_get_monitor(detail))
	case DetailTypeBios:
		return NewDetailBios(C.hd_detail_get_bios(detail))
	case DetailTypeSys:
		return NewDetailSys(C.hd_detail_get_sys(detail))
	default:
		return nil, nil
		// return nil, fmt.Errorf("unexpected detail type: %v", detailType)
	}

	// todo cdrom, floppy, prom, sys, scsi, devtree, ccw, joystick
}

type MemoryRange struct {
	Start uint   `json:"start"`
	Size  uint   `json:"size"`
	Data  string `json:"-"` // hex encoded
}

func NewMemoryRange(mem C.memory_range_t) MemoryRange {
	return MemoryRange{
		Start: uint(mem.start),
		Size:  uint(mem.size),
		Data:  hex.EncodeToString(C.GoBytes(unsafe.Pointer(&mem.data), C.int(mem.size))),
	}
}
