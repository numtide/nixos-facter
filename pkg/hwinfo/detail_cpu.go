package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

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

type DetailCpu struct {
	Type         DetailType `json:""`
	Architecture CpuArch    `json:""`
	Family       uint       `json:""`
	Model        uint       `json:""`
	Stepping     uint       `json:""`
	Cache        uint       `json:""`
	Clock        uint       `json:""`
	Units        uint       `json:""`
	VendorName   string     `json:""`
	ModelName    string     `json:""`
	Platform     string     `json:""`
	Features     []string   `json:""`
	Bogo         float64    `json:""`
}

func NewDetailCpu(cpu C.hd_detail_cpu_t) (Detail, error) {
	data := cpu.data

	return DetailCpu{
		Type:         DetailTypeCpu,
		Architecture: CpuArch(data.architecture),
		Family:       uint(data.family),
		Model:        uint(data.model),
		Stepping:     uint(data.stepping),
		Cache:        uint(data.cache),
		Clock:        uint(data.clock),
		Units:        uint(data.units),
		VendorName:   C.GoString(data.vend_name),
		ModelName:    C.GoString(data.model_name),
		Platform:     C.GoString(data.platform),
		Features:     ReadStringList(data.features),
		Bogo:         float64(data.bogo),
	}, nil
}

func (d DetailCpu) DetailType() DetailType {
	return DetailTypeCpu
}
