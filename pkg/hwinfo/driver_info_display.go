package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

type DriverInfoDisplay struct {
	Type DriverInfoType `json:"type,omitempty"`
	// actual driver database entries
	DbEntry0 []string `json:"db_entry_0,omitempty"`
	DbEntry1 []string `json:"db_entry_1,omitempty"`

	Width                 uint        `json:"width"`
	Height                uint        `json:"height"`
	VerticalSync          SyncRange   `json:"vertical_sync"`
	HorizontalSync        SyncRange   `json:"horizontal_sync"`
	Bandwidth             uint        `json:"bandwidth"`
	HorizontalSyncTimings SyncTimings `json:"horizontal_sync_timings"`
	VerticalSyncTimings   SyncTimings `json:"vertical_sync_timings"`
	HorizontalFlag        byte        `json:"horizontal_flag"`
	VerticalFlag          byte        `json:"vertical_flag"`
}

func (d DriverInfoDisplay) DriverInfoType() DriverInfoType {
	return DriverInfoTypeDisplay
}

func NewDriverInfoDisplay(info C.driver_info_display_t) DriverInfoDisplay {
	return DriverInfoDisplay{
		Type:     DriverInfoTypeDisplay,
		DbEntry0: ReadStringList(info.hddb0),
		DbEntry1: ReadStringList(info.hddb1),
		Width:    uint(info.width),
		Height:   uint(info.height),
		VerticalSync: SyncRange{
			Min: uint(info.min_vsync),
			Max: uint(info.max_vsync),
		},
		HorizontalSync: SyncRange{
			Min: uint(info.min_hsync),
			Max: uint(info.max_hsync),
		},
		HorizontalSyncTimings: SyncTimings{
			Disp:      uint(info.hdisp),
			SyncStart: uint(info.hsyncstart),
			SyncEnd:   uint(info.hsyncend),
			Total:     uint(info.htotal),
		},
		VerticalSyncTimings: SyncTimings{
			Disp:      uint(info.vdisp),
			SyncStart: uint(info.vsyncstart),
			SyncEnd:   uint(info.vsyncend),
			Total:     uint(info.vtotal),
		},
		HorizontalFlag: byte(info.hflag),
		VerticalFlag:   byte(info.vflag),
	}
}
