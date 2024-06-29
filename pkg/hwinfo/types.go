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
	Id uint `json:""`
	// Name (if any)
	Name string `json:""`
}

func (i Id) String() string {
	return fmt.Sprintf("%d:%s", i.Id, i.Name)
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
	UnixDeviceName2   string        `json:","`
	UnixDeviceNumber2 *DeviceNumber `json:",omitempty"`
}

func (i Item) String() string {
	return fmt.Sprintf("bus = %v, name = %v", i.Bus.Id, i.Bus.Name)
}

type Report struct {
	Items []*Item `json:""`

	// Log contains all messages logged during hardware probing
	Log   string `json:",omitempty"`
	Debug uint   `json:",omitempty"`
}
