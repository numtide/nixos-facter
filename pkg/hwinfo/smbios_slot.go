package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosSlot captures system slot information.
type SmbiosSlot struct {
	Type        SmbiosType `json:"type"`
	Handle      int        `json:"handle"`
	Designation string     `json:"designation,omitempty"`
	SlotType    *Id        `json:"slot_type"`
	BusWidth    *Id        `json:"bus_width"`
	Usage       *Id        `json:"usage"`
	Length      *Id        `json:"length"`
	Id          uint       `json:"id"`
	Features    []string   `json:"features,omitempty"`
}

func (s SmbiosSlot) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosSlot(info C.smbios_slot_t) (Smbios, error) {
	return SmbiosSlot{
		Type:        SmbiosTypeSlot,
		Handle:      int(info.handle),
		Designation: C.GoString(info.desig),
		SlotType:    NewId(info.slot_type),
		BusWidth:    NewId(info.bus_width),
		Usage:       NewId(info.usage),
		Length:      NewId(info.length),
		Id:          uint(info.id),
		Features:    ReadStringList(info.feature.str),
	}, nil
}
