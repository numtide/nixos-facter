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
	hd := (*C.hd_data_t)(unsafe.Pointer(C.calloc(1, C.size_t(unsafe.Sizeof(C.hd_data_t{})))))

	C.hd_set_probe_feature(hd, C.enum_probe_feature(ProbeFeatureAll))
	C.hd_scan(hd)

	report := Report{}
	report.Log = C.GoString(hd.log)
	report.Debug = uint(hd.debug)

	for hd.hd != nil {
		item := Item{}
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
		item.HardwareClass = HardwareItem(hd.hd.hw_class)
		item.Model = C.GoString(hd.hd.model)
		item.AttachedTo = uint(hd.hd.attached_to)
		item.SysfsId = C.GoString(hd.hd.sysfs_id)
		item.SysfsBusId = C.GoString(hd.hd.sysfs_bus_id)
		item.SysfsDeviceLink = C.GoString(hd.hd.sysfs_device_link)
		item.UnixDeviceName = C.GoString(hd.hd.unix_dev_name)
		item.UnixDeviceNumber = parseDeviceNumber(hd.hd.unix_dev_num)
		// todo unix dev names
		item.UnixDeviceName2 = C.GoString(hd.hd.unix_dev_name2)
		item.UnixDeviceNumber2 = parseDeviceNumber(hd.hd.unix_dev_num2)

		report.Items = append(report.Items, &item)

		hd.hd = hd.hd.next
	}

	defer C.hd_free_hd_data(hd)

	return &report, nil
}

func parseId(id C.hd_id_t) *Id {
	result := Id{}
	result.Id = uint(id.id)
	result.Name = C.GoString(id.name)

	if result.Id == 0 && result.Name == "" {
		return nil
	}

	return &result
}

func parseDeviceNumber(num C.hd_dev_num_t) *DeviceNumber {
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
