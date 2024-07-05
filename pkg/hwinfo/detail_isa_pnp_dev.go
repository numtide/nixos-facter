package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
#include <stdbool.h>

// custom getters to get around the problem with bitfields https://github.com/golang/go/issues/43261
bool hd_isapnp_card_get_broken(isapnp_card_t *card) { return card->broken; }
*/
import "C"
import (
	"encoding/hex"
	"unsafe"
)

type IsaPnpResource struct {
	Length int    `json:""`
	Type   int    `json:""`
	Data   string `json:""` // hex encoded
}

func NewIsaPnpResource(res *C.isapnp_res_t) *IsaPnpResource {
	if res == nil {
		return nil
	}
	return &IsaPnpResource{
		Length: int(res.len),
		Type:   int(res._type),
		Data:   hex.EncodeToString(C.GoBytes(unsafe.Pointer(&res.data), res.len)),
	}
}

type IsaPnpCard struct {
	Csn      int             `json:""`
	LogDevs  int             `json:""` // todo full name?
	Serial   string          `json:""`
	CardRegs string          `json:""` // todo full name?
	LdevRegs string          `json:""` // todo full name? hex encoded
	ResLen   int             `json:""` // todo full name?
	Broken   bool            `json:""` // mark a broken card
	Resource *IsaPnpResource `json:""`
}

func NewIsaPnpCard(card *C.isapnp_card_t) (*IsaPnpCard, error) {
	if card == nil {
		return nil, nil
	}
	return &IsaPnpCard{
		Csn:     int(card.csn),
		LogDevs: int(card.log_devs),
		//Serial:   C.GoString(card.serial),	todo
		//CardRegs: C.GoString(card.card_regs), todo
		LdevRegs: hex.EncodeToString(C.GoBytes(unsafe.Pointer(&card.ldev_regs), C.int(0xd0))),
		ResLen:   int(card.res_len),
		Broken:   bool(C.hd_isapnp_card_get_broken(card)),
		Resource: NewIsaPnpResource(card.res),
	}, nil
}

type DetailIsaPnpDevice struct {
	Type   DetailType  `json:""`
	Card   *IsaPnpCard `json:""`
	Device int         `json:""`
	Flags  uint        `json:""`
}

func (d DetailIsaPnpDevice) DetailType() DetailType {
	return DetailTypeIsaPnp
}

func NewDetailIsaPnpDevice(pnp C.hd_detail_isapnp_t) (Detail, error) {
	data := pnp.data

	card, err := NewIsaPnpCard(data.card)
	if err != nil {
		return nil, err
	}

	return DetailIsaPnpDevice{
		Type:   DetailTypeIsaPnp,
		Card:   card,
		Device: int(data.dev),
		Flags:  uint(data.flags),
	}, nil
}
