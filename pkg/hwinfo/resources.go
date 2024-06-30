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
bool hd_res_link_get_connected(res_link_t *res) { return res->state == 1; }

bool hd_res_fc_get_wwpn_ok(res_fc_t *res) { return res->wwpn_ok; }
bool hd_res_fc_get_fcp_lun_ok(res_fc_t *res) { return res->fcp_lun_ok; }
bool hd_res_fc_get_port_id_ok(res_fc_t *res) { return res->port_id_ok; }
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

type Resource interface {
	ResourceType() ResourceType
}

type ResourceMemory struct {
	Type     ResourceType `json:""`
	Base     uint64       `json:""`
	Range    uint64       `json:""`
	Enabled  bool         `json:""`
	Access   AccessFlags  `json:""`
	Prefetch YesNoFlags   `json:""`
}

func (r ResourceMemory) ResourceType() ResourceType {
	return r.Type
}

func NewResourceMemory(res *C.res_mem_t) *ResourceMemory {
	if res == nil {
		return nil
	}

	return &ResourceMemory{

		Type: ResourceTypeMem,

		Base:     uint64(res.base),
		Range:    uint64(res._range),
		Enabled:  bool(C.hd_res_mem_get_enabled(res)),
		Access:   AccessFlags(C.hd_res_mem_get_access(res)),
		Prefetch: YesNoFlags(C.hd_res_mem_get_prefetch(res)),
	}
}

type ResourcePhysicalMemory struct {
	Type  ResourceType `json:""`
	Range uint64       `json:""`
}

func (r ResourcePhysicalMemory) ResourceType() ResourceType {
	return r.Type
}

func NewResourcePhysicalMemory(res *C.res_phys_mem_t) *ResourcePhysicalMemory {
	if res == nil {
		return nil
	}
	return &ResourcePhysicalMemory{

		Type: ResourceTypePhysMem,

		Range: uint64(res._range),
	}
}

type ResourceIO struct {
	Type    ResourceType `json:""`
	Base    uint64       `json:""`
	Range   uint64       `json:""`
	Enabled bool         `json:""`
	Access  AccessFlags  `json:""`
}

func (r ResourceIO) ResourceType() ResourceType {
	return r.Type
}

func NewResourceIO(res *C.res_io_t) *ResourceIO {
	if res == nil {
		return nil
	}

	return &ResourceIO{

		Type: ResourceTypeIo,

		Base:    uint64(res.base),
		Range:   uint64(res._range),
		Enabled: bool(C.hd_res_io_get_enabled(res)),
		Access:  AccessFlags(C.hd_res_io_get_access(res)),
	}
}

type ResourceIrq struct {
	Type      ResourceType `json:""`
	Base      uint         `json:""`
	Triggered uint         `json:""`
	Enabled   bool         `json:""`
}

func (r ResourceIrq) ResourceType() ResourceType {
	return r.Type
}

func NewResourceIrq(res *C.res_irq_t) *ResourceIrq {
	if res == nil {
		return nil
	}

	return &ResourceIrq{

		Type: ResourceTypeIrq,

		Base:      uint(res.base),
		Triggered: uint(res.triggered),
		Enabled:   bool(C.hd_res_irq_get_enabled(res)),
	}
}

type ResourceDma struct {
	Type    ResourceType `json:""`
	Base    uint         `json:""`
	Enabled bool         `json:""`
}

func (r ResourceDma) ResourceType() ResourceType {
	return r.Type
}

func NewResourceDma(res *C.res_dma_t) *ResourceDma {
	if res == nil {
		return nil
	}

	return &ResourceDma{

		Type: ResourceTypeDma,

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
	Type   ResourceType `json:""`
	Unit   SizeUnit     `json:""`
	Value1 uint64       `json:""`
	Value2 uint64       `json:",omitempty"`
}

func (r ResourceSize) ResourceType() ResourceType {
	return r.Type
}

func NewResourceSize(res *C.res_size_t) *ResourceSize {
	if res == nil {
		return nil
	}
	return &ResourceSize{

		Type: ResourceTypeSize,

		Unit:   SizeUnit(res.unit),
		Value1: uint64(res.val1),
		Value2: uint64(res.val2),
	}
}

type ResourceBaud struct {
	Type      ResourceType `json:""`
	Speed     uint         `json:""`
	Bits      uint         `json:""`
	StopBits  uint         `json:""`
	Parity    byte         `json:""`
	Handshake byte         `json:""`
}

func (r ResourceBaud) ResourceType() ResourceType {
	return r.Type
}

func NewResourceBaud(res *C.res_baud_t) *ResourceBaud {
	if res == nil {
		return nil
	}
	return &ResourceBaud{

		Type: ResourceTypeBaud,

		Speed:     uint(res.speed),
		Bits:      uint(res.bits),
		StopBits:  uint(res.stopbits),
		Parity:    byte(res.parity),
		Handshake: byte(res.handshake),
	}

}

type ResourceCache struct {
	Type ResourceType `json:""`
	Size uint         `json:""`
}

func (r ResourceCache) ResourceType() ResourceType {
	return r.Type
}

func NewResourceCache(res *C.res_cache_t) *ResourceCache {
	if res == nil {
		return nil
	}
	return &ResourceCache{

		Type: ResourceTypeCache,

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
	Type      ResourceType `json:""`
	Cylinders uint         `json:""`
	Heads     uint         `json:""`
	Sectors   uint         `json:""`
	Size      uint64       `json:""`
	GeoType   GeoType      `json:""`
}

func (r ResourceDiskGeo) ResourceType() ResourceType {
	return r.Type
}

func NewResourceDiskGeo(res *C.res_disk_geo_t) *ResourceDiskGeo {
	if res == nil {
		return nil
	}
	return &ResourceDiskGeo{

		Type: ResourceTypeDiskGeo,

		Cylinders: uint(res.cyls),
		Heads:     uint(res.heads),
		Sectors:   uint(res.sectors),
		Size:      uint64(res.size),
		GeoType:   GeoType(res.geotype),
	}
}

type ResourceMonitor struct {
	Type              ResourceType `json:""`
	Width             uint         `json:""`
	Height            uint         `json:""`
	VerticalFrequency uint         `json:""`
	Interlaced        bool         `json:""`
}

func (r ResourceMonitor) ResourceType() ResourceType {
	return r.Type
}

func NewResourceMonitor(res *C.res_monitor_t) *ResourceMonitor {
	if res == nil {
		return nil
	}
	return &ResourceMonitor{

		Type: ResourceTypeMonitor,

		Width:             uint(res.width),
		Height:            uint(res.height),
		VerticalFrequency: uint(res.vfreq),
		Interlaced:        bool(C.hd_res_monitor_get_interlaced(res)),
	}
}

type ResourceInitStrings struct {
	Type  ResourceType `json:""`
	Init1 string       `json:",omitempty"`
	Init2 string       `json:",omitempty"`
}

func (r ResourceInitStrings) ResourceType() ResourceType {
	return r.Type
}

func NewResourceInitStrings(res *C.res_init_strings_t) *ResourceInitStrings {
	if res == nil {
		return nil
	}
	return &ResourceInitStrings{

		Type: ResourceTypeInitStrings,

		Init1: C.GoString(res.init1),
		Init2: C.GoString(res.init2),
	}
}

type ResourcePppdOption struct {
	Type   ResourceType `json:""`
	Option byte         `json:""`
}

func (r ResourcePppdOption) ResourceType() ResourceType {
	return r.Type
}

func NewResourcePppdOption(res *C.res_pppd_option_t) *ResourcePppdOption {
	if res == nil {
		return nil
	}
	return &ResourcePppdOption{
		Type:   ResourceTypePppdOption,
		Option: byte(*res.option),
	}
}

type ResourceFrameBuffer struct {
	Type         ResourceType `json:""`
	Width        uint         `json:""`
	Height       uint         `json:""`
	BytesPerLine uint         `json:""`
	ColorBits    uint         `json:""`
	Mode         uint         `json:""`
}

func (r ResourceFrameBuffer) ResourceType() ResourceType {
	return r.Type
}

func NewResourceFrameBuffer(res *C.res_framebuffer_t) *ResourceFrameBuffer {
	if res == nil {
		return nil
	}
	return &ResourceFrameBuffer{

		Type: ResourceTypeFramebuffer,

		Width:        uint(res.width),
		Height:       uint(res.height),
		BytesPerLine: uint(res.bytes_p_line),
		ColorBits:    uint(res.colorbits),
		Mode:         uint(res.mode),
	}
}

type ResourceHardwareAddress struct {
	Type    ResourceType `json:""`
	Address byte         `json:""`
}

func (r ResourceHardwareAddress) ResourceType() ResourceType {
	return r.Type
}

func NewResourceHardwareAddress(res *C.res_hwaddr_t) *ResourceHardwareAddress {
	if res == nil {
		return nil
	}
	return &ResourceHardwareAddress{

		Type: ResourceTypeHwaddr,

		Address: byte(*res.addr),
	}
}

type ResourceLink struct {
	Type      ResourceType `json:""`
	Connected bool         `json:""`
}

func (r ResourceLink) ResourceType() ResourceType {
	return r.Type
}

func NewResourceLink(res *C.res_link_t) *ResourceLink {
	if res == nil {
		return nil
	}
	return &ResourceLink{

		Type: ResourceTypeLink,

		Connected: bool(C.hd_res_link_get_connected(res)),
	}
}

type ResourceWlan struct {
	Type        ResourceType `json:""`
	Channels    []string     `json:",omitempty"`
	Frequencies []string     `json:",omitempty"`
	BitRates    []string     `json:",omitempty"`
	AuthModes   []string     `json:",omitempty"`
	EncModes    []string     `json:",omitempty"`
}

func (r ResourceWlan) ResourceType() ResourceType {
	return r.Type
}

func NewResourceWlan(res *C.res_wlan_t) *ResourceWlan {
	if res == nil {
		return nil
	}
	return &ResourceWlan{

		Type: ResourceTypeWlan,

		Channels:    readStringList(res.channels),
		Frequencies: readStringList(res.frequencies),
		BitRates:    readStringList(res.bitrates),
		AuthModes:   readStringList(res.auth_modes),
		EncModes:    readStringList(res.enc_modes),
	}
}

// todo what is FC?
type ResourceFc struct {
	Type         ResourceType `json:""`
	WwpnOk       bool         `json:""`
	FcpLunOk     bool         `json:""`
	PortIdOk     bool         `json:""`
	Wwpn         uint64       `json:""`
	FcpLun       uint64       `json:""`
	PortId       uint         `json:""`
	ControllerId byte         `json:""`
}

func (r ResourceFc) ResourceType() ResourceType {
	return r.Type
}

func NewResourceFc(res *C.res_fc_t) *ResourceFc {
	if res == nil {
		return nil
	}
	return &ResourceFc{

		Type: ResourceTypeFc,

		WwpnOk:       bool(C.hd_res_fc_get_wwpn_ok(res)),
		FcpLunOk:     bool(C.hd_res_fc_get_fcp_lun_ok(res)),
		PortIdOk:     bool(C.hd_res_fc_get_port_id_ok(res)),
		Wwpn:         uint64(res.wwpn),
		FcpLun:       uint64(res.fcp_lun),
		PortId:       uint(res.port_id),
		ControllerId: byte(*res.controller_id),
	}
}

type ResourceAny struct {
	Type ResourceType `json:""`
}

func (r ResourceAny) ResourceType() ResourceType {
	return r.Type
}

func NewResourceAny(res *C.res_any_t) *ResourceAny {
	if res == nil {
		return nil
	}
	return &ResourceAny{

		Type: ResourceTypeAny,
	}
}

func NewResource(res *C.hd_res_t) Resource {
	if res == nil {
		return nil
	}

	switch ResourceType(res._type) {
	case ResourceTypeFc:
		return NewResourceFc(res.fc)
	case ResourceTypeAny:
		return NewResourceAny(res.any)
	case ResourceTypePhysMem:
		return NewResourcePhysicalMemory(res.phys_mem)
	case ResourceTypeMem:
		return NewResourceMemory(res.mem)
	case ResourceTypeIo:
		return NewResourceIO(res.io)
	case ResourceTypeIrq:
		return NewResourceIrq(res.irq)
	case ResourceTypeDma:
		return NewResourceDma(res.dma)
	case ResourceTypeMonitor:
		return NewResourceMonitor(res.monitor)
	case ResourceTypeSize:
		return NewResourceSize(res.size)
	case ResourceTypeDiskGeo:
		return NewResourceDiskGeo(res.disk_geo)
	case ResourceTypeCache:
		return NewResourceCache(res.cache)
	case ResourceTypeBaud:
		return NewResourceBaud(res.baud)
	case ResourceTypeInitStrings:
		return NewResourceInitStrings(res.ini_strings)
	case ResourceTypePppdOption:
		return NewResourcePppdOption(res.pppd_option)
	case ResourceTypeFramebuffer:
		return NewResourceFrameBuffer(res.framebuffer)
	case ResourceTypeHwaddr:
		return NewResourceHardwareAddress(res.hwaddr)
	case ResourceTypeLink:
		return NewResourceLink(res.link)
	case ResourceTypeWlan:
		return NewResourceWlan(res.wlan)
	case ResourceTypePhwaddr:
		// todo
		return nil
		// todo
	}

	// todo maybe return an error?
	return nil
}

func readResources(hd *C.hd_t) []Resource {
	var result []Resource
	for res := hd.res; res != nil; res = res.next {
		result = append(result, NewResource(res))
	}
	return result
}
