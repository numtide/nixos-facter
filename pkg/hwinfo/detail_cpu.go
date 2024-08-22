package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

bool cpu_info_fpu(cpu_info_t *info) { return info->fpu; }
bool cpu_info_fpu_exception(cpu_info_t *info) { return info->fpu_exception; }
bool cpu_info_write_protect(cpu_info_t *info) { return info->write_protect; }
*/
import "C"
import "regexp"

//go:generate enumer -type=CpuArch -json -transform=snake -trimprefix CpuArch -output=./detail_enum_cpu_arch.go
type CpuArch uint

const (
	CpuArchUnknown CpuArch = iota
	CpuArchIntel
	CpuArchAlpha
	CpuArchSparc
	CpuArchSparc64
	CpuArchPpc
	CpuArchPpc64
	CpiArch68k
	CpuArchIa64
	CpuArchS390
	CpuArchS390x
	CpuArchArm
	CpuArchMips
	CpuArchx86_64
	CpuArchAarch64
	CpuArchLoongarch
	CpuArchRiscv
)

type AddressSizes struct {
	Physical uint `json:"physical,omitempty"`
	Virtual  uint `json:"virtual,omitempty"`
}

type DetailCpu struct {
	Type DetailType `json:"-"`

	Architecture CpuArch `json:"architecture"`

	VendorName string `json:"vendor_name,omitempty"`
	ModelName  string `json:"model_name,omitempty"`

	Family   uint `json:"family"`
	Model    uint `json:"model"`
	Stepping uint `json:"stepping"`

	Platform string `json:"platform,omitempty"`

	Features        []string `json:"features,omitempty"`
	Bugs            []string `json:"bugs,omitempty"`
	PowerManagement []string `json:"power_management,omitempty"`

	Bogo  float64 `json:"bogo"`
	Cache uint    `json:"cache,omitempty"`
	Units uint    `json:"units,omitempty"`
	Clock uint    `json:"-"`

	// x86 only fields
	PhysicalId     uint         `json:"physical_id"`
	Siblings       uint         `json:"siblings,omitempty"`
	Cores          uint         `json:"cores,omitempty"`
	CoreId         uint         `json:"-"`
	Fpu            bool         `json:"fpu"`
	FpuException   bool         `json:"fpu_exception"`
	CpuidLevel     uint         `json:"cpuid_level,omitempty"`
	WriteProtect   bool         `json:"write_protect"`
	TlbSize        uint         `json:"tlb_size,omitempty"`
	ClflushSize    uint         `json:"clflush_size,omitempty"`
	CacheAlignment int          `json:"cache_alignment,omitempty"`
	AddressSizes   AddressSizes `json:"address_sizes,omitempty"`
	Apicid         uint         `json:"-"`
	ApicidInitial  uint         `json:"-"`
}

var matchCPUFreq = regexp.MustCompile(`, \d+ MHz$`)

func stripCpuFreq(s string) string {
	// strip frequency of the model name as it is not stable.
	return matchCPUFreq.ReplaceAllString(s, "")
}

func NewDetailCpu(cpu C.hd_detail_cpu_t) (Detail, error) {
	data := cpu.data

	return DetailCpu{
		Type: DetailTypeCpu,

		Architecture: CpuArch(data.architecture),
		VendorName:   C.GoString(data.vend_name),
		ModelName:    stripCpuFreq(C.GoString(data.model_name)),

		Family:   uint(data.family),
		Model:    uint(data.model),
		Stepping: uint(data.stepping),

		Platform: C.GoString(data.platform),

		Features:        ReadStringList(data.features),
		Bugs:            ReadStringList(data.bugs),
		PowerManagement: ReadStringList(data.power_management),

		Clock: uint(data.clock),
		Bogo:  float64(data.bogo),
		Cache: uint(data.cache),
		Units: uint(data.units),

		PhysicalId:     uint(data.physical_id),
		Siblings:       uint(data.siblings),
		Cores:          uint(data.cores),
		CoreId:         uint(data.core_id),
		Apicid:         uint(data.apicid),
		ApicidInitial:  uint(data.apicid_initial),
		Fpu:            bool(C.cpu_info_fpu(data)),
		FpuException:   bool(C.cpu_info_fpu_exception(data)),
		CpuidLevel:     uint(data.cpuid_level),
		WriteProtect:   bool(C.cpu_info_write_protect(data)),
		TlbSize:        uint(data.tlb_size),
		ClflushSize:    uint(data.clflush_size),
		CacheAlignment: int(data.cache_alignment),
		AddressSizes: AddressSizes{
			Physical: uint(data.address_size_physical),
			Virtual:  uint(data.address_size_virtual),
		},
	}, nil
}

func (d DetailCpu) DetailType() DetailType {
	return DetailTypeCpu
}
