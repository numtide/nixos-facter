package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>

// CGO cannot access union type fields, so we do this as a workaround
res_baud_t hd_res_get_baud(hd_res_t *res) { return res->baud; }
*/
import "C"
import "fmt"

type ResourceBaud struct {
	Type      ResourceType `json:"type"`
	Speed     uint         `json:"speed"`
	Bits      uint         `json:"bits"`
	StopBits  uint         `json:"stop_bits"`
	Parity    byte         `json:"parity"`
	Handshake byte         `json:"handshake"`
}

func (r ResourceBaud) ResourceType() ResourceType {
	return ResourceTypeBaud
}

func NewResourceBaud(res *C.hd_res_t, resType ResourceType) (*ResourceBaud, error) {
	if res == nil {
		return nil, nil
	}

	if resType != ResourceTypeBaud {
		return nil, fmt.Errorf("expected resource type '%s', found '%s'", ResourceTypeBaud, resType)
	}

	baud := C.hd_res_get_baud(res)

	return &ResourceBaud{
		Type:      resType,
		Speed:     uint(baud.speed),
		Bits:      uint(baud.bits),
		StopBits:  uint(baud.stopbits),
		Parity:    byte(baud.parity),
		Handshake: byte(baud.handshake),
	}, nil
}
