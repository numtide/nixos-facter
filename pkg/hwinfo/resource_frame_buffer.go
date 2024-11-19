package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_framebuffer_t hd_res_get_framebuffer(hd_res_t *res) { return res->framebuffer; }
*/
import "C"

import (
	"errors"
	"fmt"
)

type ResourceFrameBuffer struct {
	Type         ResourceType `json:"type"`
	Width        uint         `json:"width"`
	Height       uint         `json:"height"`
	BytesPerLine uint         `json:"bytes_per_line"`
	ColorBits    uint         `json:"color_bits"`
	Mode         uint         `json:"mode"`
}

func (r ResourceFrameBuffer) ResourceType() ResourceType {
	return ResourceTypeFramebuffer
}

func NewResourceFrameBuffer(res *C.hd_res_t, resType ResourceType) (*ResourceFrameBuffer, error) {
	if res == nil {
		return nil, errors.New("res is nil")
	}

	if resType != ResourceTypeFramebuffer {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeFramebuffer, resType)
	}

	fb := C.hd_res_get_framebuffer(res)

	return &ResourceFrameBuffer{
		Type:         resType,
		Width:        uint(fb.width),
		Height:       uint(fb.height),
		BytesPerLine: uint(fb.bytes_p_line),
		ColorBits:    uint(fb.colorbits),
		Mode:         uint(fb.mode),
	}, nil
}
