package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_mem_get_enabled(res_mem_t *res) { return res->enabled; }
unsigned hd_res_mem_get_access(res_mem_t *res) { return res->access; }
unsigned hd_res_mem_get_prefetch(res_mem_t *res) { return res->prefetch; }

*/
import "C"

//go:generate enumer -type=ResourceType -json -trimprefix ResourceType
type ResourceType int

const (
	ResourceAny ResourceType = iota
	ResourcePhysMem
	ResourceMem
	ResourceIo
	ResourceIrq
	ResourceDma
	ResourceMonitor

	ResourceSize
	ResourceDisk_geo
	ResourceCache
	ResourceBaud
	ResourceInitStrings
	ResourcePppdOption

	ResourceFramebuffer
	ResourceHwaddr
	ResourceLink
	ResourceWlan
	ResourceFc
	ResourcePhwaddr
)

//go:generate enumer -type=AccessFlags -json -trimprefix AccessFlags
type AccessFlags int

const (
	AccessFlagsUnknown AccessFlags = iota
	AccessFlagsReadOnly
	AccessFlagsWriteOnly
	AccessFlagsReadWrite
)

//go:generate enumer -type=YesNoFlags -json -trimprefix YesNoFlags
type YesNoFlags int

const (
	YesNoFlagsUnknown YesNoFlags = iota
	YesNoFlagsNo
	YesNoFlagsYes
)

type Resource struct {
	Type ResourceType `json:",omitempty"`
}

type ResourceMemory struct {
	Resource
	Base     uint64      `json:",omitempty"`
	Range    uint64      `json:",omitempty"`
	Enabled  bool        `json:",omitempty"`
	Access   AccessFlags `json:",omitempty"`
	Prefetch YesNoFlags  `json:",omitempty"`
}

func NewResourceMemory(res *C.res_mem_t) *ResourceMemory {
	if res == nil {
		return nil
	}

	return &ResourceMemory{
		Resource: Resource{
			Type: ResourceMem,
		},
		Base:  uint64(res.base),
		Range: uint64(res._range),
		// We use custom getters to access bit fields
		Enabled:  bool(C.hd_res_mem_get_enabled(res)),
		Access:   AccessFlags(C.hd_res_mem_get_access(res)),
		Prefetch: YesNoFlags(C.hd_res_mem_get_prefetch(res)),
	}
}

type ResourcePhysicalMemory struct {
	Resource
	Range uint64 `json:",omitempty"`
}

func NewResourcePhysicalMemory(res *C.res_phys_mem_t) *ResourcePhysicalMemory {
	if res == nil {
		return nil
	}
	return &ResourcePhysicalMemory{
		Resource: Resource{
			Type: ResourcePhysMem,
		},
		Range: uint64(res._range),
	}
}
