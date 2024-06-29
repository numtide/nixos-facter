package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type hdData = C.hd_data_t
type hdId = C.hd_id_t
type hd = C.hd_t

func Scan() {
	hd := (*hdData)(unsafe.Pointer(C.calloc(1, C.size_t(unsafe.Sizeof(hdData{})))))

	ProbeFeatureAll.Set(hd)
	C.hd_scan(hd)

	fmt.Println("Scan complete")

	data := HardwareData{}
	data.Log = C.GoString(hd.log)
	data.Debug = uint(hd.debug)

	for hd.hd != nil {
		item := HardwareItem{}
		item.Index = uint(hd.hd.idx)
		item.Bus = parseId(hd.hd.bus)
		item.Slot = Slot(hd.hd.slot)
		item.BaseClass = parseId(hd.hd.base_class)
		item.SubClass = parseId(hd.hd.sub_class)
		item.PciInterface = parseId(hd.hd.prog_if)
		item.Vendor = parseId(hd.hd.vendor)
		item.SubVendor = parseId(hd.hd.sub_vendor)
		item.Device = parseId(hd.hd.device)
		item.SubDevice = parseId(hd.hd.sub_device)
		item.Revision = parseId(hd.hd.revision)
		item.Serial = C.GoString(hd.hd.serial)
		item.CompatVendor = parseId(hd.hd.compat_vendor)
		item.CompatDevice = parseId(hd.hd.compat_device)

		data.Items = append(data.Items, item)
		hd.hd = hd.hd.next
	}

	defer C.hd_free_hd_data(hd)

	for _, item := range data.Items {
		fmt.Printf("New item: %d\n", item.Index)
		fmt.Printf("Bus = %v\n", item.Bus)
		fmt.Printf("Slot = %v\n", item.Slot)
		fmt.Printf("Base Class = %v\n", item.BaseClass)
		fmt.Printf("Sub Class = %v\n", item.SubClass)
		fmt.Printf("PCI Interface = %v\n", item.PciInterface)
		fmt.Printf("Vendor = %v\n", item.Vendor)
		fmt.Printf("Sub Vendor = %v\n", item.SubVendor)
		fmt.Printf("Device = %v\n", item.Device)
		fmt.Printf("Sub Device = %v\n", item.SubDevice)
		fmt.Printf("Revision = %v\n", item.Revision)
		fmt.Printf("Serial = %v\n", item.Serial)
		fmt.Println()
	}
}

func parseId(id hdId) *Id {
	result := Id{}
	result.Id = uint(id.id)
	result.Name = C.GoString(id.name)
	return &result
}
