package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_hwaddr_t hd_res_get_hwaddr(hd_res_t *res) { return res->hwaddr; }
*/
import "C"
import "fmt"

type ResourceHardwareAddress struct {
	Type    ResourceType `json:""`
	Address byte         `json:""`
}

func (r ResourceHardwareAddress) ResourceType() ResourceType {
	return r.Type
}

func NewResourceHardwareAddress(res *C.hd_res_t, resType ResourceType) (*ResourceHardwareAddress, error) {

	if !(resType == ResourceTypeHwaddr || resType == ResourceTypePhwaddr) {
		return nil, fmt.Errorf("invalid resource type %s, must be either %s or %s", resType, ResourceTypeHwaddr, ResourceTypePhwaddr)
	}

	if res == nil {
		return nil, nil
	}

	hwaddr := C.hd_res_get_hwaddr(res)

	return &ResourceHardwareAddress{
		Type:    resType,
		Address: byte(*hwaddr.addr),
	}, nil
}
