package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type SyncRange struct {
	Min uint `json:""`
	Max uint `json:""`
}

type SyncTimings struct {
	Disp      uint `json:""` // todo what's the proper name for this?
	SyncStart uint `json:""`
	SyncEnd   uint `json:""`
	Total     uint `json:""` // todo what's a better name for this?
}

type DetailMonitor struct {
	Type                  DetailType  `json:""`
	ManufactureYear       uint        `json:""`
	ManufactureWeek       uint        `json:""`
	VerticalSync          SyncRange   `json:""`
	HorizontalSync        SyncRange   `json:""`
	HorizontalSyncTimings SyncTimings `json:""`
	VerticalSyncTimings   SyncTimings `json:""`
	Clock                 uint        `json:""`
	Width                 uint        `json:""`
	Height                uint        `json:""`
	WidthMillimetres      uint        `json:""`
	HeightMillimetres     uint        `json:""`
	HorizontalFlag        byte        `json:""`
	VerticalFlag          byte        `json:""`
	Vendor                string      `json:""`
	Name                  string      `json:""`
	Serial                string      `json:""`
}

func (d DetailMonitor) DetailType() DetailType {
	return DetailTypeMonitor
}

func NewDetailMonitor(mon C.hd_detail_monitor_t) (Detail, error) {
	data := mon.data

	return DetailMonitor{
		Type:            DetailTypeMonitor,
		ManufactureYear: uint(data.manu_year),
		ManufactureWeek: uint(data.manu_week),
		VerticalSync: SyncRange{
			Min: uint(data.min_vsync),
			Max: uint(data.max_vsync),
		},
		HorizontalSync: SyncRange{
			Min: uint(data.min_hsync),
			Max: uint(data.max_hsync),
		},
		Clock:             uint(data.clock),
		Width:             uint(data.width),
		Height:            uint(data.height),
		WidthMillimetres:  uint(data.width_mm),
		HeightMillimetres: uint(data.height_mm),
		HorizontalSyncTimings: SyncTimings{
			Disp:      uint(data.hdisp),
			SyncStart: uint(data.hsyncstart),
			SyncEnd:   uint(data.hsyncend),
			Total:     uint(data.htotal),
		},
		VerticalSyncTimings: SyncTimings{
			Disp:      uint(data.vdisp),
			SyncStart: uint(data.vsyncstart),
			SyncEnd:   uint(data.vsyncend),
			Total:     uint(data.vtotal),
		},
		HorizontalFlag: byte(data.hflag),
		VerticalFlag:   byte(data.vflag),
		Vendor:         C.GoString(data.vendor),
		Name:           C.GoString(data.name),
		Serial:         C.GoString(data.serial),
	}, nil
}
