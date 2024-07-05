package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool driver_info_module_is_active(driver_info_module_t info) { return info.active; }
bool driver_info_module_is_modprobe(driver_info_module_t info) { return info.modprobe; }
*/
import "C"

type DriverInfoModule struct {
	Type DriverInfoType `json:",omitempty"`
	// actual driver database entries
	DbEntry0   []string `json:",omitempty"`
	DbEntry1   []string `json:",omitempty"`
	Active     bool     `json:""` // if the module is currently active
	Modprobe   bool     `json:""` // modprobe or insmod
	Names      []string `json:""` // (ordered) list of module names
	ModuleArgs []string `json:""` // list of module args (corresponds to the module name list)
	Conf       string   `json:""` // conf.modules entry, if any (e.g., for sb.o)
}

func (d DriverInfoModule) DriverInfoType() DriverInfoType {
	return DriverInfoTypeModule
}

func NewDriverInfoModule(info C.driver_info_module_t) DriverInfoModule {
	return DriverInfoModule{
		Type:       DriverInfoTypeModule,
		DbEntry0:   ReadStringList(info.hddb0),
		DbEntry1:   ReadStringList(info.hddb1),
		Active:     bool(C.driver_info_module_is_active(info)),
		Modprobe:   bool(C.driver_info_module_is_modprobe(info)),
		Names:      ReadStringList(info.names),
		ModuleArgs: ReadStringList(info.mod_args),
		Conf:       C.GoString(info.conf),
	}
}
