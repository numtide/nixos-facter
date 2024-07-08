package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import (
	"encoding/hex"
	"unsafe"
)

//go:generate enumer -type=PciFlag -json -transform=snake -trimprefix PciFlag -output=./detail_enum_pci_flag.go
type PciFlag uint

const (
	PciFlagOk PciFlag = iota
	PciFlagPm
	PciFlagAgp
)

func ParsePciFlags(flags uint) (result []PciFlag) {
	for _, flag := range PciFlagValues() {
		if (flag & (1 << flags)) == 1 {
			result = append(result, flag)
		}
	}
	return
}

type DetailPci struct {
	Type          DetailType `json:"type"`
	Data          string     `json:"data"`            // the PCI data, hex encoded
	DataLength    uint       `json:"data_length"`     // holds the actual length of Data
	DataExtLength uint       `json:"data_ext_length"` // max. accessed config byte

	Log string `json:""` // log messages

	Flags []PciFlag `json:"flags,omitempty"` //
	// todo map pci constants from pci.h?
	Command      uint `json:"command"`       // PCI_COMMAND
	HeaderType   uint `json:"header_type"`   // PCI_HEADER_TYPE
	SecondaryBus uint `json:"secondary_bus"` // > 0 for PCI & CB bridges

	Bus      uint `json:"bus"` // PCI bus number
	Slot     uint `json:"slot"`
	Function uint `json:"function"`

	// PCI device classes
	BaseClass uint `json:"base_class"`
	SubClass  uint `json:"sub_class"`
	ProgIf    uint `json:"prog_if"`

	// Vendor and device ids
	Device    uint `json:"device"`
	Vendor    uint `json:"vendor"`
	SubDevice uint `json:"sub_device"`
	SubVendor uint `json:"sub_vendor"`
	Revision  uint `json:"revision"`

	// used irq if any
	Irq uint `json:"irq"`

	BaseAddress  [7]uint64 `json:"base_address"`  // I/O or memory base
	BaseLength   [7]uint64 `json:"base_length"`   // I/O or memory ranges
	AddressFlags [7]uint   `json:"address_flags"` // I/O or memory address flags

	RomBaseAddress uint64 `json:"rom_base_address"` // memory base for card ROM
	RomBaseLength  uint64 `json:"rom_base_length"`  // memory range for card ROM

	SysfsId     string `json:"sysfs_id,omitempty"`     // sysfs path
	SysfsBusId  string `json:"sysfs_bus_id,omitempty"` // sysfs bus id
	ModuleAlias string `json:"module_alias,omitempty"` // module alias
	Label       string `json:"label,omitempty"`        // Consistent Device Name (CDN), pci firmware 3.1, chapter 4.6.7

	// todo edid data
	//EdidData   [6][0x80]byte `json:",omitempty"` // edid record
	//EdidLength [6]uint       `json:",omitempty"` // edid record length
}

func (p DetailPci) DetailType() DetailType {
	return DetailTypePci
}

func NewDetailPci(pci C.hd_detail_pci_t) (Detail, error) {

	data := pci.data

	return DetailPci{
		Type:           DetailTypePci,
		Data:           hex.EncodeToString(C.GoBytes(unsafe.Pointer(&data.data), 256)),
		DataLength:     uint(data.data_len),
		DataExtLength:  uint(data.data_ext_len),
		Log:            C.GoString(data.log),
		Flags:          ParsePciFlags(uint(data.flags)),
		Command:        uint(data.cmd),
		HeaderType:     uint(data.hdr_type),
		SecondaryBus:   uint(data.secondary_bus),
		Bus:            uint(data.bus),
		Slot:           uint(data.slot),
		Function:       uint(data._func),
		BaseClass:      uint(data.base_class),
		SubClass:       uint(data.sub_class),
		ProgIf:         uint(data.prog_if),
		Device:         uint(data.dev),
		Vendor:         uint(data.vend),
		SubDevice:      uint(data.sub_dev),
		SubVendor:      uint(data.sub_vend),
		Revision:       uint(data.rev),
		Irq:            uint(data.irq),
		BaseAddress:    [7]uint64(ReadUint64Array(unsafe.Pointer(&data.base_addr), 7)),
		BaseLength:     [7]uint64(ReadUint64Array(unsafe.Pointer(&data.base_len), 7)),
		AddressFlags:   [7]uint(ReadUintArray(unsafe.Pointer(&data.addr_flags), 7)),
		RomBaseAddress: uint64(data.rom_base_addr),
		RomBaseLength:  uint64(data.rom_base_len),
		SysfsId:        C.GoString(data.sysfs_id),
		SysfsBusId:     C.GoString(data.sysfs_bus_id),
		ModuleAlias:    C.GoString(data.modalias),
		Label:          C.GoString(data.label),
		// todo edid data
	}, nil
}
