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

func Scan(f func(item *HardwareItem) error) error {
	data := (*C.hd_data_t)(unsafe.Pointer(C.calloc(1, C.size_t(unsafe.Sizeof(C.hd_data_t{})))))

	C.hd_set_probe_feature(data, C.enum_probe_feature(ProbeFeatureAll))
	C.hd_scan(data)
	defer C.hd_free_hd_data(data)

	for hd := data.hd; hd != nil; hd = hd.next {
		if item, err := NewHardwareItem(hd); err != nil {
			return err
		} else if err = f(item); err != nil {
			return err
		}
	}

	return nil
}
