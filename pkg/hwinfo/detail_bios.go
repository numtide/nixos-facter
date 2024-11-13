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
	Supported  bool `json:"supported"`
	Enabled    bool `json:"enabled"`
	Version    uint `json:"version"`
	SubVersion uint `json:"sub_version"`
	BiosFlags  uint `json:"bios_flags"`
}

type VbeInfo struct {
	Version     uint `json:"version"`
	VideoMemory uint `json:"video_memory"`
}

type DetailBios struct {
	Type    DetailType `json:"-"`
	ApmInfo ApmInfo    `json:"apm_info"`
	VbeInfo VbeInfo    `json:"vbe_info"`

	// todo par and ser ports
	PnP           bool `json:"pnp"`
	PnPId         uint `json:"pnp_id"` // it is still in big endian format
	LbaSupport    bool `json:"lba_support"`
	LowMemorySize uint `json:"low_memory_size"`
	// todo smp info
	// todo vbe info

	SmbiosVersion uint `json:"smbios_version"`

	// todo lcd
	// todo mouse
	// todo led
	// todo bios32
}

func (d DetailBios) DetailType() DetailType {
	return DetailTypeBios
}

func NewDetailBios(dev C.hd_detail_bios_t) (*DetailBios, error) {
	data := dev.data

	return &DetailBios{
		Type: DetailTypeBios,
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
