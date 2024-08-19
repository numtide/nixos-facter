package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
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

type DetailCpu struct {
	Type         DetailType `json:"-"`
	Architecture CpuArch    `json:"architecture"`
	Family       uint       `json:"family"`
	Model        uint       `json:"model"`
	Stepping     uint       `json:"stepping"`
	Cache        uint       `json:"cache"`
	// This field changes as the CPU up/down scales, so we do not export it
	Clock      uint     `json:"-"`
	Units      uint     `json:"units"`
	VendorName string   `json:"vendor_name"`
	ModelName  string   `json:"model_name"`
	Platform   string   `json:"platform"`
	Features   []string `json:"features"`
	Bogo       float64  `json:"bogo"`
}

var matchCPUFreq = regexp.MustCompile(`, \d+ MHz$`)

func stripCpuFreq(s string) string {
	// strip frequency of the model name as it is not stable.
	return matchCPUFreq.ReplaceAllString(s, "")
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
		ModelName:    stripCpuFreq(C.GoString(data.model_name)),
		Platform:     C.GoString(data.platform),
		Features:     ReadStringList(data.features),
		Bogo:         float64(data.bogo),
	}, nil
}

func (d DetailCpu) DetailType() DetailType {
	return DetailTypeCpu
}
