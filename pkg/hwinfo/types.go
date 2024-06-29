package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"
import "fmt"

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

func (pf ProbeFeature) Set(data *hdData) {
	C.hd_set_probe_feature(data, C.enum_probe_feature(pf))
}

func (pf ProbeFeature) Clear(data *hdData) {
	C.hd_clear_probe_feature(data, C.enum_probe_feature(pf))
}

type Id struct {
	// Id is a numeric id
	Id uint
	// Name (if any)
	Name string
}

func (i Id) String() string {
	return fmt.Sprintf("%d:%s", i.Id, i.Name)
}

// Slot represents a slot and bus number.
// Bits 0-7: slot number, 8-31 bus number
type Slot uint

func (s Slot) Slot() byte {
	return byte(s & 0xFF)
}

func (s Slot) Bus() uint {
	return uint(s & 0xFFFFFF)
}

func (s Slot) String() string {
	return fmt.Sprintf("%d:%d", s.Slot(), s.Bus())
}

type HardwareItem struct {
	// Index is a unique index, starting at 1
	Index uint
	// Bus type (id and name)
	Bus          *Id
	Slot         Slot
	BaseClass    *Id
	SubClass     *Id
	PciInterface *Id
	Vendor       *Id
	SubVendor    *Id
	Device       *Id
	SubDevice    *Id
	Revision     *Id
	Serial       string
	CompatVendor *Id
	CompatDevice *Id
}

func (i HardwareItem) String() string {
	return fmt.Sprintf("bus = %v, name = %v", i.Bus.Id, i.Bus.Name)
}

type HardwareData struct {
	Items []HardwareItem

	// Log contains all messages logged during hardware probing
	Log   string
	Debug uint
}
