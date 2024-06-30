package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_res_mem_get_enabled(res_mem_t *res) { return res->enabled; }
unsigned hd_res_mem_get_access(res_mem_t *res) { return res->access; }
unsigned hd_res_mem_get_prefetch(res_mem_t *res) { return res->prefetch; }

bool hd_res_io_get_enabled(res_io_t *res) { return res->enabled; }
unsigned hd_res_io_get_access(res_io_t *res) { return res->access; }

bool hd_res_irq_get_enabled(res_irq_t *res) { return res->enabled; }
bool hd_res_dma_get_enabled(res_dma_t *res) { return res->enabled; }
bool hd_res_monitor_get_interlaced(res_monitor_t *res) { return res->interlaced; }
*/
import "C"

//go:generate enumer -type=ResourceType -json -trimprefix ResourceType
type ResourceType int

const (
	ResourceTypeAny ResourceType = iota
	ResourceTypePhysMem
	ResourceTypeMem
	ResourceTypeIo
	ResourceTypeIrq
	ResourceTypeDma
	ResourceTypeMonitor

	ResourceTypeSize
	ResourceTypeDiskGeo
	ResourceTypeCache
	ResourceTypeBaud
	ResourceTypeInitStrings
	ResourceTypePppdOption

	ResourceTypeFramebuffer
	ResourceTypeHwaddr
	ResourceTypeLink
	ResourceTypeWlan
	ResourceTypeFc
	ResourceTypePhwaddr
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
	Type ResourceType `json:""`
}

type ResourceMemory struct {
	Resource
	Base     uint64      `json:""`
	Range    uint64      `json:""`
	Enabled  bool        `json:""`
	Access   AccessFlags `json:""`
	Prefetch YesNoFlags  `json:""`
}

func NewResourceMemory(res *C.res_mem_t) *ResourceMemory {
	if res == nil {
		return nil
	}

	return &ResourceMemory{
		Resource: Resource{
			Type: ResourceTypeMem,
		},
		Base:     uint64(res.base),
		Range:    uint64(res._range),
		Enabled:  bool(C.hd_res_mem_get_enabled(res)),
		Access:   AccessFlags(C.hd_res_mem_get_access(res)),
		Prefetch: YesNoFlags(C.hd_res_mem_get_prefetch(res)),
	}
}

type ResourcePhysicalMemory struct {
	Resource
	Range uint64 `json:""`
}

func NewResourcePhysicalMemory(res *C.res_phys_mem_t) *ResourcePhysicalMemory {
	if res == nil {
		return nil
	}
	return &ResourcePhysicalMemory{
		Resource: Resource{
			Type: ResourceTypePhysMem,
		},
		Range: uint64(res._range),
	}
}

type ResourceIO struct {
	Resource
	Base    uint64      `json:""`
	Range   uint64      `json:""`
	Enabled bool        `json:""`
	Access  AccessFlags `json:""`
}

func NewResourceIO(res *C.res_io_t) *ResourceIO {
	if res == nil {
		return nil
	}

	return &ResourceIO{
		Resource: Resource{
			Type: ResourceTypeIo,
		},
		Base:    uint64(res.base),
		Range:   uint64(res._range),
		Enabled: bool(C.hd_res_io_get_enabled(res)),
		Access:  AccessFlags(C.hd_res_io_get_access(res)),
	}
}

type ResourceIrq struct {
	Resource
	Base      uint `json:""`
	Triggered uint `json:""`
	Enabled   bool `json:""`
}

func NewResourceIrq(res *C.res_irq_t) *ResourceIrq {
	if res == nil {
		return nil
	}

	return &ResourceIrq{
		Resource: Resource{
			Type: ResourceTypeIrq,
		},
		Base:      uint(res.base),
		Triggered: uint(res.triggered),
		Enabled:   bool(C.hd_res_irq_get_enabled(res)),
	}
}

type ResourceDma struct {
	Resource
	Base    uint `json:""`
	Enabled bool `json:""`
}

func NewResourceDma(res *C.res_dma_t) *ResourceDma {
	if res == nil {
		return nil
	}

	return &ResourceDma{
		Resource: Resource{
			Type: ResourceTypeDma,
		},
		Base:    uint(res.base),
		Enabled: bool(C.hd_res_dma_get_enabled(res)),
	}
}

//go:generate enumer -type=SizeUnit -json -trimprefix SizeUnit
type SizeUnit int

const (
	SizeUnitCm SizeUnit = iota
	SizeUnitCinch
	SizeUnitByte
	SizeUnitSectors

	SizeUnitKbyte
	SizeUnitMbyte
	SizeUnitGbyte
	SizeUnitMm
)

type ResourceSize struct {
	Resource
	Unit   SizeUnit `json:""`
	Value1 uint64   `json:""`
	Value2 uint64   `json:",omitempty"`
}

func NewResourceSize(res *C.res_size_t) *ResourceSize {
	if res == nil {
		return nil
	}
	return &ResourceSize{
		Resource: Resource{
			Type: ResourceTypeSize,
		},
		Unit:   SizeUnit(res.unit),
		Value1: uint64(res.val1),
		Value2: uint64(res.val2),
	}
}

type ResourceBaud struct {
	Resource
	Speed     uint `json:""`
	Bits      uint `json:""`
	StopBits  uint `json:""`
	Parity    byte `json:""`
	Handshake byte `json:""`
}

func NewResourceBaud(res *C.res_baud_t) *ResourceBaud {
	if res == nil {
		return nil
	}
	return &ResourceBaud{
		Resource: Resource{
			Type: ResourceTypeBaud,
		},
		Speed:     uint(res.speed),
		Bits:      uint(res.bits),
		StopBits:  uint(res.stopbits),
		Parity:    byte(res.parity),
		Handshake: byte(res.handshake),
	}

}

type ResourceCache struct {
	Resource
	Size uint `json:""`
}

func NewResourceCache(res *C.res_cache_t) *ResourceCache {
	if res == nil {
		return nil
	}
	return &ResourceCache{
		Resource: Resource{
			Type: ResourceTypeCache,
		},
		Size: uint(res.size),
	}
}

//go:generate enumer -type=GeoType -json -trimprefix GeoType
type GeoType int

const (
	GeoTypePhysical GeoType = iota
	GeoTypeLogical
	GeoTypeBiosEdd
	GeoTypeBiosLegacy
)

type ResourceDiskGeo struct {
	Resource
	Cylinders uint    `json:""`
	Heads     uint    `json:""`
	Sectors   uint    `json:""`
	Size      uint64  `json:""`
	GeoType   GeoType `json:""`
}

func NewResourceDiskGeo(res *C.res_disk_geo_t) *ResourceDiskGeo {
	if res == nil {
		return nil
	}
	return &ResourceDiskGeo{
		Resource: Resource{
			Type: ResourceTypeDiskGeo,
		},
		Cylinders: uint(res.cyls),
		Heads:     uint(res.heads),
		Sectors:   uint(res.sectors),
		Size:      uint64(res.size),
		GeoType:   GeoType(res.geotype),
	}
}

type ResourceMonitor struct {
	Resource
	Width             uint `json:""`
	Height            uint `json:""`
	VerticalFrequency uint `json:""`
	Interlaced        bool `json:""`
}

func NewResourceMonitor(res *C.res_monitor_t) *ResourceMonitor {
	if res == nil {
		return nil
	}
	return &ResourceMonitor{
		Resource: Resource{
			Type: ResourceTypeMonitor,
		},
		Width:             uint(res.width),
		Height:            uint(res.height),
		VerticalFrequency: uint(res.vfreq),
		Interlaced:        bool(C.hd_res_monitor_get_interlaced(res)),
	}
}

type ResourceInitStrings struct {
	Resource
	Init1 string `json:",omitempty"`
	Init2 string `json:",omitempty"`
}

func NewResourceInitStrings(res *C.res_init_strings) *ResourceInitStrings {
	if res == nil {
		return nil
	}
	return &ResourceInitStrings{
		Resource: Resource{
			Type: ResourceTypeInitStrings,
		},
		Init1: C.GoString(res.init1),
		Init2: C.GoString(res.init2),
	}
}
