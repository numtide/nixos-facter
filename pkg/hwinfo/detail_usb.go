package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

//go:generate enumer -type=UsbClass -json -transform=snake -trimprefix UsbClass -output=./detail_usb_enum_usb_class.go
type UsbClass uint16

const (
	UsbClassPerInterface       UsbClass = 0x00
	UsbClassAudio              UsbClass = 0x01
	UsbClassComm               UsbClass = 0x02
	UsbClassHID                UsbClass = 0x03
	UsbClassPhysical           UsbClass = 0x05
	UsbClassImage              UsbClass = 0x06
	UsbClassPTP                UsbClass = UsbClassImage // legacy name for image
	UsbClassPrinter            UsbClass = 0x07
	UsbClassMassStorage        UsbClass = 0x08
	UsbClassHub                UsbClass = 0x09
	UsbClassData               UsbClass = 0x0a
	UsbClassSmartCard          UsbClass = 0x0b
	UsbClassContentSecurity    UsbClass = 0x0d
	UsbClassVideo              UsbClass = 0x0e
	UsbClassPersonalHealthcare UsbClass = 0x0f
	UsbClassAudioVideo         UsbClass = 0x10
	UsbClassBillboard          UsbClass = 0x11
	UsbClassUSBTypeCBridge     UsbClass = 0x12
	UsbClassDiagnosticDevice   UsbClass = 0xdc
	UsbClassWireless           UsbClass = 0xe0
	UsbClassMiscellaneous      UsbClass = 0xef
	UsbClassApplication        UsbClass = 0xfe
	UsbClassVendorSpec         UsbClass = 0xff
)

type DetailUsb struct {
	Type                DetailType  `json:""`
	Bus                 int         `json:""`
	DeviceNumber        int         `json:""`
	Lev                 int         `json:""` // TODO what is lev short for?
	Parent              int         `json:""`
	Port                int         `json:""`
	Count               int         `json:""`
	Connections         int         `json:""`
	UsedConnections     int         `json:""`
	InterfaceDescriptor int         `json:""`
	Speed               uint        `json:""`
	Vendor              uint        `json:""`
	Device              uint        `json:""`
	Revision            uint        `json:""`
	Manufacturer        string      `json:",omitempty"`
	Product             string      `json:",omitempty"`
	Serial              string      `json:",omitempty"`
	Driver              string      `json:",omitempty"`
	RawDescriptor       MemoryRange `json:""`

	DeviceClass    UsbClass `json:",omitempty"`
	DeviceSubclass int      `json:",omitempty"`
	DeviceProtocol int      `json:",omitempty"`

	InterfaceClass    UsbClass `json:",omitempty"`
	InterfaceSubclass int      `json:",omitempty"`
	InterfaceProtocol int      `json:",omitempty"`

	Country uint `json:""`
}

func (d DetailUsb) DetailType() DetailType {
	return DetailTypeUsb
}

func NewDetailUsb(usb C.hd_detail_usb_t) (Detail, error) {

	data := usb.data

	if data.next != nil {
		println("usb next is not nil")
	}

	return DetailUsb{
		Type:                DetailTypeUsb,
		Bus:                 int(data.bus),
		DeviceNumber:        int(data.dev_nr),
		Lev:                 int(data.lev),
		Parent:              int(data.parent),
		Port:                int(data.port),
		Count:               int(data.count),
		Connections:         int(data.conns),
		UsedConnections:     int(data.used_conns),
		InterfaceDescriptor: int(data.ifdescr),
		Speed:               uint(data.speed),
		Vendor:              uint(data.vendor),
		Device:              uint(data.device),
		Revision:            uint(data.rev),
		Manufacturer:        C.GoString(data.manufact),
		Product:             C.GoString(data.product),
		Serial:              C.GoString(data.serial),
		Driver:              C.GoString(data.driver),
		RawDescriptor:       NewMemoryRange(data.raw_descr),
		DeviceClass:         UsbClass(data.d_cls),
		DeviceSubclass:      int(data.d_sub),
		DeviceProtocol:      int(data.d_prot),
		// todo data.i_alt??
		InterfaceClass:    UsbClass(data.i_cls),
		InterfaceSubclass: int(data.i_sub),
		InterfaceProtocol: int(data.i_prot),
	}, nil
}
