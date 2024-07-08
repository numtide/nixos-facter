package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

func Scan() (*Report, error) {
	data := (*C.hd_data_t)(unsafe.Pointer(C.calloc(1, C.size_t(unsafe.Sizeof(C.hd_data_t{})))))

	C.hd_set_probe_feature(data, C.enum_probe_feature(ProbeFeatureAll))
	C.hd_scan(data)
	defer C.hd_free_hd_data(data)

	report := Report{}

	for hd := data.hd; hd != nil; hd = hd.next {
		item, err := NewItem(hd)
		if err != nil {
			return nil, err
		}
		report.Items = append(report.Items, item)
		// get the next item in the list
		hd = hd.next
	}

	return &report, nil
}
