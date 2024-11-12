package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type SyncRange struct {
	Min uint `json:"min"`
	Max uint `json:"max"`
}

type SyncTimings struct {
	Disp      uint `json:"disp"` // todo what's the proper name for this?
	SyncStart uint `json:"sync_start"`
	SyncEnd   uint `json:"sync_end"`
	Total     uint `json:"total"` // todo what's a better name for this?
}

type DetailMonitor struct {
	Type                  DetailType  `json:"-"`
	ManufactureYear       uint        `json:"manufacture_year"`
	ManufactureWeek       uint        `json:"manufacture_week"`
	VerticalSync          SyncRange   `json:"vertical_sync"`
	HorizontalSync        SyncRange   `json:"horizontal_sync"`
	HorizontalSyncTimings SyncTimings `json:"horizontal_sync_timings"`
	VerticalSyncTimings   SyncTimings `json:"vertical_sync_timings"`
	Clock                 uint        `json:"clock"`
	Width                 uint        `json:"width"`
	Height                uint        `json:"height"`
	WidthMillimetres      uint        `json:"width_millimetres"`
	HeightMillimetres     uint        `json:"height_millimetres"`
	HorizontalFlag        byte        `json:"horizontal_flag"`
	VerticalFlag          byte        `json:"vertical_flag"`
	Vendor                string      `json:"vendor"`
	Name                  string      `json:"name"`

	Serial string `json:"-"`
}

func (d DetailMonitor) DetailType() DetailType {
	return DetailTypeMonitor
}

func NewDetailMonitor(mon C.hd_detail_monitor_t) (*DetailMonitor, error) {
	data := mon.data

	return &DetailMonitor{
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
