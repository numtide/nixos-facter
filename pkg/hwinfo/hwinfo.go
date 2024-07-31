package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdlib.h>

// CGO cannot access union type fields, so we do this as a workaround
hd_smbios_t* hd_smbios_next(hd_smbios_t *sm) { return sm->next; }
*/
import "C"

import (
	"slices"
	"unsafe"
)

func Scan(probes []ProbeFeature) ([]Smbios, []*HardwareItem, error) {
	// initialise the struct to hold scan data
	data := (*C.hd_data_t)(unsafe.Pointer(C.calloc(1, C.size_t(unsafe.Sizeof(C.hd_data_t{})))))

	// set the hardware probes to run
	for _, probe := range probes {
		C.hd_set_probe_feature(data, C.enum_probe_feature(probe))
	}

	// scan
	C.hd_scan(data)
	defer C.hd_free_hd_data(data)

	var smbiosItems []Smbios
	for sm := data.smbios; sm != nil; sm = C.hd_smbios_next(sm) {
		item, err := NewSmbios(sm)
		if err != nil {
			return nil, nil, err
		}
		smbiosItems = append(smbiosItems, item)
	}

	var hardwareItems []*HardwareItem
	for hd := data.hd; hd != nil; hd = hd.next {
		if item, err := NewHardwareItem(hd); err != nil {
			return nil, nil, err
		} else {
			hardwareItems = append(hardwareItems, item)
		}
	}

	// canonically sort by device index
	slices.SortFunc(hardwareItems, func(a, b *HardwareItem) int {
		return int(a.Index) - int(b.Index)
	})

	return smbiosItems, hardwareItems, nil
}
