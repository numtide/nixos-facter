package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import (
	"fmt"
)

//go:generate enumer -type=ProbeFeature -json -trimprefix ProbeFeature
type ProbeFeature int

const (
	ProbeFeatureMemory ProbeFeature = iota + 1
	ProbeFeaturePci
	ProbeFeatureIsapnp
	ProbeFeatureNet
	ProbeFeatureFloppy
	ProbeFeatureMisc

	ProbeFeatureMiscSerial
	ProbeFeatureMiscPar
	ProbeFeatureMiscFloppy
	ProbeFeatureSerial
	ProbeFeatureCpu
	ProbeFeatureBios

	ProbeFeatureMonitor
	ProbeFeatureMouse
	ProbeFeatureScsi
	ProbeFeatureUsb
	ProbeFeatureUsbMods
	ProbeFeatureAdb
	ProbeFeatureModem

	ProbeFeatureModemUsb
	ProbeFeatureParallel
	ProbeFeatureParallelLp
	ProbeFeatureParallelZip
	ProbeFeatureIsa

	ProbeFeatureIsaIsdn
	ProbeFeatureIsdn
	ProbeFeatureKbd
	ProbeFeatureProm
	ProbeFeatureSbus
	ProbeFeatureInt
	ProbeFeatureBraille

	ProbeFeatureBrailleAlva
	ProbeFeatureBrailleFhp
	ProbeFeatureBrailleHt
	ProbeFeatureIgnx11
	ProbeFeatureSys

	ProbeFeatureBiosVbe
	ProbeFeatureIsapnpOld
	ProbeFeatureIsapnpNew
	ProbeFeatureIsapnpMod
	ProbeFeatureBrailleBaum

	ProbeFeatureManual
	ProbeFeatureFb
	ProbeFeatureVeth
	ProbeFeaturePppoe
	ProbeFeatureScan
	ProbeFeaturePcmcia
	ProbeFeatureFork

	ProbeFeatureParallelImm
	ProbeFeatureS390
	ProbeFeatureCpuemu
	ProbeFeatureSysfs
	ProbeFeatureS390disks
	ProbeFeatureUdev

	ProbeFeatureBlock
	ProbeFeatureBlockCdrom
	ProbeFeatureBlockPart
	ProbeFeatureEdd
	ProbeFeatureEddMod
	ProbeFeatureBiosDdc

	ProbeFeatureBiosFb
	ProbeFeatureBiosMode
	ProbeFeatureInput
	ProbeFeatureBlockMods
	ProbeFeatureBiosVesa

	ProbeFeatureCpuemuDebug
	ProbeFeatureScsiNoserial
	ProbeFeatureWlan
	ProbeFeatureBiosCrc
	ProbeFeatureHal

	ProbeFeatureBiosVram
	ProbeFeatureBiosAcpi
	ProbeFeatureBiosDdcPorts
	ProbeFeatureModulesPata

	ProbeFeatureNetEeprom
	ProbeFeatureX86emu

	ProbeFeatureMax
	ProbeFeatureLxrc
	ProbeFeatureDefault

	ProbeFeatureAll
)

//go:generate enumer -type=HardwareItem -json -trimprefix HardwareItem
type HardwareItem int

const (
	HardwareItemNone HardwareItem = iota
	HardwareItemSys
	HardwareItemCpu
	HardwareItemKeyboard
	HardwareItemBraille
	HardwareItemMouse

	HardwareItemJoystick
	HardwareItemPrinter
	HardwareItemScanner
	HardwareItemChipcard
	HardwareItemMonitor
	HardwareItemTv

	HardwareItemDisplay
	HardwareItemFramebuffer
	HardwareItemCamera
	HardwareItemSound
	HardwareItemStorageCtrl

	HardwareItemNetworkCtrl
	HardwareItemIsdn
	HardwareItemModem
	HardwareItemNetwork
	HardwareItemDisk
	HardwareItemPartition

	HardwareItemCdrom
	HardwareItemFloppy
	HardwareItemManual
	HardwareItemUsbCtrl
	HardwareItemUsb
	HardwareItemBios
	HardwareItemPci

	HardwareItemIsapnp
	HardwareItemBridge
	HardwareItemHub
	HardwareItemScsi
	HardwareItemIde
	HardwareItemMemory
	HardwareItemDvb

	HardwareItemPcmcia
	HardwareItemPcmciaCtrl
	HardwareItemIeee1394
	HardwareItemIeee1394Ctrl
	HardwareItemHotplug

	HardwareItemHotplugCtrl
	HardwareItemZip
	HardwareItemPppoe
	HardwareItemWlan
	HardwareItemRedasd
	HardwareItemDsl
	HardwareItemBlock

	HardwareItemTape
	HardwareItemVbe
	HardwareItemBluetooth
	HardwareItemFingerprint
	HardwareItemMmcCtrl
	HardwareItemNvme

	/** append new entries here */
	HardwareItemUnknown
	HardwareItemAll
)

type Id struct {
	// Id is a numeric id
	Id uint `json:",omitempty"`
	// Name (if any)
	Name string `json:",omitempty"`
}

func (i Id) IsEmpty() bool {
	return i.Id == 0 && i.Name == ""
}

func (i Id) String() string {
	return fmt.Sprintf("%d:%s", i.Id, i.Name)
}

func NewId(id C.hd_id_t) *Id {
	result := Id{
		Id:   uint(id.id),
		Name: C.GoString(id.name),
	}
	if result.IsEmpty() {
		return nil
	}
	return &result
}

// Slot represents a slot and bus number.
// Bits 0-7: slot number, 8-31 bus number
type Slot uint

func (s *Slot) Slot() byte {
	return byte(*s & 0xFF)
}

func (s *Slot) Bus() uint {
	return uint(*s & 0xFFFFFF)
}

func (s *Slot) String() string {
	return fmt.Sprintf("%d:%d", s.Slot(), s.Bus())
}

func (s *Slot) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d:%d\"", s.Slot(), s.Bus())), nil
}

// TODO UnmarshalJSON for Slot

type DeviceNumber struct {
	Type  int  `json:""`
	Major uint `json:""`
	Minor uint `json:""`
	Range uint `json:""`
}

func (d DeviceNumber) IsEmpty() bool {
	return d.Type == 0 && d.Major == 0 && d.Minor == 0 && d.Range == 0
}

func NewDeviceNumber(num C.hd_dev_num_t) *DeviceNumber {
	result := DeviceNumber{
		Type:  int(num._type),
		Major: uint(num.major),
		Minor: uint(num.minor),
		Range: uint(num._range),
	}
	if result.IsEmpty() {
		return nil
	}
	return &result
}

type Item struct {
	// Index is a unique index, starting at 1
	Index uint `json:""`
	// Bus type (id and name)
	Bus               *Id           `json:",omitempty"`
	Slot              Slot          `json:",omitempty"`
	BaseClass         *Id           `json:",omitempty"`
	SubClass          *Id           `json:",omitempty"`
	PciInterface      *Id           `json:",omitempty"`
	Vendor            *Id           `json:",omitempty"`
	SubVendor         *Id           `json:",omitempty"`
	Device            *Id           `json:",omitempty"`
	SubDevice         *Id           `json:",omitempty"`
	Revision          *Id           `json:",omitempty"`
	Serial            string        `json:",omitempty"`
	CompatVendor      *Id           `json:",omitempty"`
	CompatDevice      *Id           `json:",omitempty"`
	HardwareClass     HardwareItem  `json:",omitempty"`
	Model             string        `json:",omitempty"`
	AttachedTo        uint          `json:",omitempty"`
	SysfsId           string        `json:",omitempty"`
	SysfsBusId        string        `json:",omitempty"`
	SysfsDeviceLink   string        `json:",omitempty"`
	UnixDeviceName    string        `json:",omitempty"`
	UnixDeviceNumber  *DeviceNumber `json:",omitempty"`
	UnixDeviceNames   []string      `json:",omitempty"`
	UnixDeviceName2   string        `json:",omitempty"`
	UnixDeviceNumber2 *DeviceNumber `json:",omitempty"`
	RomId             string        `json:",omitempty"`
	Udi               string        `json:",omitempty"`
	ParentUdi         string        `json:",omitempty"`

	/*
		UniqueId is a unique string identifying this hardware.
		The string consists of two parts separated by a dot (".").
		The part before the dot describes the location (where the hardware is attached in the system).
		The part after the dot identifies the hardware itself.
		The string must not contain slashes ("/") because we're going to create files with this id as name.
		Apart from this, there are no restrictions on the form of this string.
	*/
	UniqueId  string   `json:",omitempty"`
	UniqueIds []string `json:",omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf("bus = %v, name = %v", i.Bus.Id, i.Bus.Name)
}

func NewItem(hd *C.hd_t) *Item {
	if hd == nil {
		return nil
	}
	return &Item{
		Index:            uint(hd.idx),
		Bus:              NewId(hd.bus),
		Slot:             Slot(hd.slot),
		BaseClass:        NewId(hd.base_class),
		SubClass:         NewId(hd.sub_class),
		PciInterface:     NewId(hd.prog_if),
		Vendor:           NewId(hd.vendor),
		SubVendor:        NewId(hd.sub_vendor),
		Device:           NewId(hd.device),
		SubDevice:        NewId(hd.sub_device),
		Revision:         NewId(hd.revision),
		Serial:           C.GoString(hd.serial),
		CompatVendor:     NewId(hd.compat_vendor),
		CompatDevice:     NewId(hd.compat_device),
		HardwareClass:    HardwareItem(hd.hw_class),
		Model:            C.GoString(hd.model),
		AttachedTo:       uint(hd.attached_to),
		SysfsId:          C.GoString(hd.sysfs_id),
		SysfsBusId:       C.GoString(hd.sysfs_bus_id),
		SysfsDeviceLink:  C.GoString(hd.sysfs_device_link),
		UnixDeviceName:   C.GoString(hd.unix_dev_name),
		UnixDeviceNumber: NewDeviceNumber(hd.unix_dev_num),
		// todo unix dev names
		UnixDeviceName2:   C.GoString(hd.unix_dev_name2),
		UnixDeviceNumber2: NewDeviceNumber(hd.unix_dev_num2),
		RomId:             C.GoString(hd.rom_id),
		Udi:               C.GoString(hd.udi),
		ParentUdi:         C.GoString(hd.parent_udi),
		UniqueId:          C.GoString(hd.unique_id),
		UniqueIds:         readStringList(hd.unique_ids),
	}
}

type Report struct {
	Items []*Item `json:""`

	// Log contains all messages logged during hardware probing
	Log   string `json:",omitempty"`
	Debug uint   `json:",omitempty"`
}

func readStringList(list *C.str_list_t) (result []string) {
	if list == nil {
		return nil
	}
	for entry := list; entry != nil; entry = entry.next {
		result = append(result, C.GoString(list.str))
		entry = entry.next
	}
	return result
}
