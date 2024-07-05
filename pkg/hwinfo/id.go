package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import (
	"fmt"
	"slices"
)

//go:generate enumer -type=Vendor -json -transform=snake -trimprefix Vendor -output=./id_vendor_enum.go
type Vendor uint16

const (
	VendorBroadcom         Vendor = 0x14e4
	VendorRedHat           Vendor = 0x1af4
	VendorIntelCorporation Vendor = 0x8086
)

var (
	// DevicesSTA list taken from https://github.com/NixOS/nixpkgs/blob/dac9cdf8c930c0af98a63cbfe8005546ba0125fb/nixos/modules/installer/tools/nixos-generate-config.pl#L152-L158
	DevicesSTA = []uint16{
		0x4311, 0x4312, 0x4313, 0x4315, 0x4327, 0x4328,
		0x4329, 0x432a, 0x432b, 0x432c, 0x432d, 0x4353,
		0x4357, 0x4358, 0x4359, 0x4331, 0x43a0, 0x43b1,
	}

	// DevicesFullMac list taken from https://wireless.wiki.kernel.org/en/users/Drivers/brcm80211#brcmfmac
	DevicesFullMac = []uint16{
		0x43a3, 0x43df, 0x43ec, 0x43d3, 0x43d9, 0x43e9,
		0x43ba, 0x43bb, 0x43bc, 0xaa52, 0x43ca, 0x43cb,
		0x43cc, 0x43c3, 0x43c4, 0x43c5,
	}

	DevicesVirtioSCSI = []uint16{
		0x1004, 0x1048,
	}

	DevicesIntel2200BG = []uint16{
		0x1043, 0x104f, 0x4220, 0x4221, 0x4223, 0x4224,
	}

	DevicesIntel3945ABG = []uint16{
		0x4229, 0x4230, 0x4222, 0x4227,
	}
)

//go:generate enumer -type=IdTag -json -transform=snake -trimprefix IdTag -output=./id_tag_enum.go
type IdTag byte

const (
	IdTagPci IdTag = iota + 1
	IdTagEisa
	IdTagUsb
	IdTagSpecial
	IdTagPcmcia
	IdTagSdio
)

type Id struct {
	Tag   IdTag  `json:",omitempty"`
	Value uint16 `json:",omitempty"`
	// Name (if any)
	Name string `json:",omitempty"`
}

func (i Id) IsEmpty() bool {
	return i.Tag == 0 && i.Value == 0 && i.Name == ""
}

func (i Id) String() string {
	return fmt.Sprintf("%d:%s", i.Value, i.Name)
}

func (i Id) Is(ids ...uint16) bool {
	return slices.Contains(ids, i.Value)
}

func (i Id) IsVendor(vendors ...Vendor) bool {
	for idx := range vendors {
		if i.Value == uint16(vendors[idx]) {
			return true
		}
	}
	return false
}

func NewId(id C.hd_id_t) *Id {
	result := Id{
		/*
			 	Id is actually a combination of some tag to differentiate the various id types and the real id.
				We do the same thing as the ID_VALUE macro in hd.h to get the true value.
		*/
		Tag:   IdTag((id.id >> 16) & 0xf),
		Value: uint16(id.id),
		Name:  C.GoString(id.name),
	}
	if result.IsEmpty() {
		return nil
	}
	return &result
}
