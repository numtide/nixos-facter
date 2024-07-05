package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

bool bios_info_is_apm_supported(bios_info_t *info) { return info->apm_supported; }
bool bios_info_is_apm_enabled(bios_info_t *info) { return info->apm_enabled; }
bool bios_info_is_pnp_bios(bios_info_t *info) { return info->is_pnp_bios; }
bool bios_info_has_lba_support(bios_info_t *info) { return info->lba_support; }
*/
import "C"

type ApmInfo struct {
	Supported  bool `json:""`
	Enabled    bool `json:""`
	Version    uint `json:""`
	SubVersion uint `json:""`
	BiosFlags  uint `json:""`
}

type VbeInfo struct {
	Version     uint `json:""`
	VideoMemory uint `json:""`
}

type DetailBios struct {
	Type    DetailType `json:""`
	ApmInfo ApmInfo    `json:""`
	VbeInfo VbeInfo    `json:""`

	// todo par and ser ports
	PnP           bool `json:""`
	PnPId         uint `json:""` // it is still in big endian format
	LbaSupport    bool `json:""`
	LowMemorySize uint `json:""`
	// todo smp info
	// todo vbe info

	SmbiosVersion uint `json:""`

	// todo lcd
	// todo mouse
	// todo led
	// todo bios32
}

func (d DetailBios) DetailType() DetailType {
	return DetailTypeBios
}

func NewDetailBios(dev C.hd_detail_bios_t) (Detail, error) {
	data := dev.data

	return &DetailBios{
		ApmInfo: ApmInfo{
			Supported:  bool(C.bios_info_is_apm_supported(data)),
			Enabled:    bool(C.bios_info_is_apm_enabled(data)),
			Version:    uint(data.apm_ver),
			SubVersion: uint(data.apm_subver),
			BiosFlags:  uint(data.apm_bios_flags),
		},
		VbeInfo: VbeInfo{
			Version:     uint(data.vbe_ver),
			VideoMemory: uint(data.vbe_video_mem),
		},
		PnP:           bool(C.bios_info_is_pnp_bios(data)),
		PnPId:         uint(data.pnp_id),
		LbaSupport:    bool(C.bios_info_has_lba_support(data)),
		LowMemorySize: uint(data.low_mem_size),
		SmbiosVersion: uint(data.smbios_ver),
	}, nil
}
