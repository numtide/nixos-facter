package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

import (
	"encoding/json"
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

//go:generate enumer -type=HardwareClass -json -transform=snake -trimprefix HardwareClass -output=./report_enum_hardware_class.go
type HardwareClass uint

const (
	HardwareClassNone HardwareClass = iota
	HardwareClassSystem
	HardwareClassCpu
	HardwareClassKeyboard
	HardwareClassBraille
	HardwareClassMouse

	HardwareClassJoystick
	HardwareClassPrinter
	HardwareClassScanner
	HardwareClassChipcard
	HardwareClassMonitor
	HardwareClassTvCard

	HardwareClassGraphicsCard
	HardwareClassFramebuffer
	HardwareClassCamera
	HardwareClassSound
	HardwareClassStorageCtrl

	HardwareClassNetwork
	HardwareClassIsdnAdapter
	HardwareClassModem
	HardwareClassNetworkInterface
	HardwareClassDisk
	HardwareClassPartition

	HardwareClassCdrom
	HardwareClassFloppy
	HardwareClassManual
	HardwareClassUsbController
	HardwareClassUsb
	HardwareClassBios
	HardwareClassPci

	HardwareClassIsapnp
	HardwareClassBridge
	HardwareClassHub
	HardwareClassScsi
	HardwareClassIde
	HardwareClassMemory
	HardwareClassDvbCard

	HardwareClassPcmcia
	HardwareClassPcmciaController
	HardwareClassFirewire
	HardwareClassFirewireController
	HardwareClassHotplug

	HardwareClassHotplugController
	HardwareClassZip
	HardwareClassPppoe
	HardwareClassWlanCard
	HardwareClassRedasd
	HardwareClassDslAdapter
	HardwareClassBlockDevice

	HardwareClassTape
	HardwareClassVesaBios
	HardwareClassBluetooth
	HardwareClassFingerprint
	HardwareClassMmcController
	HardwareClassNvme

	HardwareClassUnknown
	HardwareClassAll
)

// Slot represents a slot and bus number.
// Bits 0-7: slot number, 8-31 bus number
type Slot uint

func (s *Slot) Slot() byte {
	return byte(*s & 0xFF)
}

func (s *Slot) Bus() uint {
	return uint((*s & 0xFFFFFF00) >> 8)
}

func (s *Slot) String() string {
	return fmt.Sprintf("%d:%d", s.Slot(), s.Bus())
}

func (s *Slot) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"bus":    s.Bus(),
		"number": s.Slot(),
	})
}

type DeviceNumber struct {
	Type  int  `json:"type"`
	Major uint `json:"major"`
	Minor uint `json:"minor"`
	Range uint `json:"range"`
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

//go:generate enumer -type=Hotplug -json -transform=snake -trimprefix Hotplug -output=./report_enum_hotplug.go
type Hotplug int

const (
	HotplugNone Hotplug = iota
	HotplugPcmcia
	HotplugCardbus
	HotplugPci
	HotplugUsb
	HotplugFirewire
)

type HardwareDevice struct {
	// Index is a unique index, starting at 1
	Index uint `json:"index"`

	// Bus type (id and name)
	Bus               *Id           `json:"bus,omitempty"`
	Slot              Slot          `json:"slot,omitempty"`
	BaseClass         *Id           `json:"base_class,omitempty"`
	SubClass          *Id           `json:"sub_class,omitempty"`
	PciInterface      *Id           `json:"pci_interface,omitempty"`
	Vendor            *Id           `json:"vendor,omitempty"`
	SubVendor         *Id           `json:"sub_vendor,omitempty"`
	Device            *Id           `json:"device,omitempty"`
	SubDevice         *Id           `json:"sub_device,omitempty"`
	Revision          *Id           `json:"revision,omitempty"`
	Serial            string        `json:"-"` // exclude from json output
	CompatVendor      *Id           `json:"compat_vendor,omitempty"`
	CompatDevice      *Id           `json:"compat_device,omitempty"`
	HardwareClass     HardwareClass `json:"hardware_class,omitempty"`
	Model             string        `json:"model,omitempty"`
	AttachedTo        uint          `json:"attached_to,omitempty"`
	SysfsId           string        `json:"sysfs_id,omitempty"`
	SysfsBusId        string        `json:"sysfs_bus_id,omitempty"`
	SysfsDeviceLink   string        `json:"sysfs_device_link,omitempty"`
	UnixDeviceName    string        `json:"unix_device_name,omitempty"`
	UnixDeviceNumber  *DeviceNumber `json:"unix_device_number,omitempty"`
	UnixDeviceNames   []string      `json:"unix_device_names,omitempty"`
	UnixDeviceName2   string        `json:"unix_device_name_2,omitempty"`
	UnixDeviceNumber2 *DeviceNumber `json:"unix_device_number_2,omitempty"`
	RomId             string        `json:"rom_id,omitempty"`
	Udi               string        `json:"udi,omitempty"`
	ParentUdi         string        `json:"parent_udi,omitempty"`

	/*
		UniqueId is a unique string identifying this hardware.
		The string consists of two parts separated by a dot (".").
		The part before the dot describes the location (where the hardware is attached in the system).
		The part after the dot identifies the hardware itself.
		The string must not contain slashes ("/") because we're going to create files with this id as name.
		Apart from this, there are no restrictions on the form of this string.
	*/
	UniqueId  string   `json:"unique_id,omitempty"`
	UniqueIds []string `json:"unique_ids,omitempty"`

	Resources []Resource `json:"resources,omitempty"`
	Detail    Detail     `json:"detail,omitempty"`

	Hotplug     Hotplug `json:"hotplug"`      // indicates what kind of hotplug device (if any) this is
	HotplugSlot uint    `json:"hotplug_slot"` // slot the hotplug device is connected to (e.g. PCMCIA socket), count is 1-based (0: no info available)

	// todo is?
	Driver        string     `json:"driver,omitempty"`         // currently active driver
	DriverModule  string     `json:"driver_module,omitempty"`  // currently active driver module (if any)
	Drivers       []string   `json:"drivers,omitempty"`        // list of currently active drivers
	DriverModules []string   `json:"driver_modules,omitempty"` // list of currently active driver modules
	DriverInfo    DriverInfo `json:"driver_info,omitempty"`    // device driver info
	UsbGuid       string     `json:"usb_guid,omitempty"`       // USB Global Unique Identifier.
	Requires      []string   `json:",omitempty"`               // packages/programs required for this hardware

	// todo hal_prop
	// todo persistent_prop

	ModuleAlias string `json:"module_alias,omitempty"` // module alias
	Label       string `json:"label,omitempty"`        // Consistent Device Name (CDN), pci firmware spec 3.1, chapter 4.6.7
}

func NewHardwareDevice(hd *C.hd_t) (*HardwareDevice, error) {
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
	model := C.GoString(hd.model)
	hwClass := HardwareClass(hd.hw_class)
	if hwClass == HardwareClassCpu {
		model = stripCpuFreq(model)
	}

	return &HardwareDevice{
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
		HardwareClass:    hwClass,
		Model:            model,
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
		Hotplug:           Hotplug(hd.hotplug),
		HotplugSlot:       uint(hd.hotplug_slot),
	}, nil
}
