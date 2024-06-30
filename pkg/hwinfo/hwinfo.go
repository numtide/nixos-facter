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
	report.Log = C.GoString(data.log)
	report.Debug = uint(data.debug)

	// get first item in the list
	hd := data.hd

	for hd != nil {
		item := Item{}
		item.Index = uint(hd.idx)
		item.Bus = readId(hd.bus)
		item.Slot = Slot(hd.slot)
		item.BaseClass = readId(hd.base_class)
		item.SubClass = readId(hd.sub_class)
		item.PciInterface = readId(hd.prog_if)
		item.Vendor = readId(hd.vendor)
		item.SubVendor = readId(hd.sub_vendor)
		item.Device = readId(hd.device)
		item.SubDevice = readId(hd.sub_device)
		item.Revision = readId(hd.revision)
		item.Serial = C.GoString(hd.serial)
		item.CompatVendor = readId(hd.compat_vendor)
		item.CompatDevice = readId(hd.compat_device)
		item.HardwareClass = HardwareItem(hd.hw_class)
		item.Model = C.GoString(hd.model)
		item.AttachedTo = uint(hd.attached_to)
		item.SysfsId = C.GoString(hd.sysfs_id)
		item.SysfsBusId = C.GoString(hd.sysfs_bus_id)
		item.SysfsDeviceLink = C.GoString(hd.sysfs_device_link)
		item.UnixDeviceName = C.GoString(hd.unix_dev_name)
		item.UnixDeviceNumber = readDeviceNumber(hd.unix_dev_num)
		// todo unix dev names
		item.UnixDeviceName2 = C.GoString(hd.unix_dev_name2)
		item.UnixDeviceNumber2 = readDeviceNumber(hd.unix_dev_num2)

		report.Items = append(report.Items, &item)

		// get next item in the list
		hd = hd.next
	}

	return &report, nil
}

func readId(id C.hd_id_t) *Id {
	result := Id{}
	result.Id = uint(id.id)
	result.Name = C.GoString(id.name)

	if result.Id == 0 && result.Name == "" {
		return nil
	}

	return &result
}

func readDeviceNumber(num C.hd_dev_num_t) *DeviceNumber {
	result := DeviceNumber{}
	result.Type = int(num._type)
	result.Major = uint(num.major)
	result.Minor = uint(num.minor)
	result.Range = uint(num._range)

	if result.Type == 0 && result.Major == 0 && result.Minor == 0 && result.Range == 0 {
		return nil
	}
	return &result
}
