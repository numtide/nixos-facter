package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import (
	"fmt"
)

//go:generate enumer -type=ProbeFeature -json -transform=snake -trimprefix ProbeFeature -output=./report_enum_probe_feature.go
type ProbeFeature uint

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

//go:generate enumer -type=HardwareItem -json -transform=snake -trimprefix HardwareItem -output=./report_enum_hardware_item.go
type HardwareItem uint

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

	Resources []Resource `json:",omitempty"`
	Detail    Detail     `json:",omitempty"`

	// todo status
	// todo config string
	// todo hotplug
	// todo hotplug_slot
	// todo is?
	Driver        string     `json:",omitempty"` // currently active driver
	DriverModule  string     `json:",omitempty"` // currently active driver module (if any)
	Drivers       []string   `json:",omitempty"` // list of currently active drivers
	DriverModules []string   `json:",omitempty"` // list of currently active driver modules
	DriverInfo    DriverInfo `json:",omitempty"` // device driver info
	UsbGuid       string     `json:",omitempty"` // USB Global Unique Identifier.
	Requires      []string   `json:",omitempty"` // packages/programs required for this hardware

	// todo hal_prop
	// todo persistent_prop

	ModuleAlias string `json:",omitempty"` // module alias
	Label       string `json:",omitempty"` // Consistent Device Name (CDN), pci firmware spec 3.1, chapter 4.6.7
}

func (i Item) String() string {
	return fmt.Sprintf("bus = %v, name = %v", i.Bus.Value, i.Bus.Name)
}

func NewItem(hd *C.hd_t) (*Item, error) {
	if hd == nil {
		return nil, nil
	}

	resources, err := NewResources(hd)
	if err != nil {
		return nil, fmt.Errorf("failed to read resources: %w", err)
	}

	detail, err := NewDetail(hd.detail)
	if err != nil {
		return nil, fmt.Errorf("failed to read detail: %w", err)
	}

	driverInfo, err := NewDriverInfo(hd.driver_info)
	if err != nil {
		return nil, fmt.Errorf("failed to read driver info: %w", err)
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
		UniqueIds:         ReadStringList(hd.unique_ids),
		Resources:         resources,
		Detail:            detail,
		Driver:            C.GoString(hd.driver),
		DriverModule:      C.GoString(hd.driver_module),
		Drivers:           ReadStringList(hd.drivers),
		DriverModules:     ReadStringList(hd.driver_modules),
		UsbGuid:           C.GoString(hd.usb_guid),
		DriverInfo:        driverInfo,
		Requires:          ReadStringList(hd.requires),
		ModuleAlias:       C.GoString(hd.modalias),
		Label:             C.GoString(hd.label),
	}, nil
}

type Report struct {
	Items []*Item `json:""`

	// Log contains all messages logged during hardware probing
	Log   string `json:",omitempty"`
	Debug uint   `json:",omitempty"`
}
