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
	Type                DetailType  `json:"type"`
	Bus                 int         `json:"bus"`
	DeviceNumber        int         `json:"device_number"`
	Lev                 int         `json:"lev"` // TODO what is lev short for?
	Parent              int         `json:"parent"`
	Port                int         `json:"port"`
	Count               int         `json:"count"`
	Connections         int         `json:"connections"`
	UsedConnections     int         `json:"used_connections"`
	InterfaceDescriptor int         `json:"interface_descriptor"`
	Speed               uint        `json:"speed"`
	Vendor              uint        `json:"vendor"`
	Device              uint        `json:"device"`
	Revision            uint        `json:"revision"`
	Manufacturer        string      `json:"manufacturer,omitempty"`
	Product             string      `json:"product,omitempty"`
	Serial              string      `json:"serial,omitempty"`
	Driver              string      `json:"driver,omitempty"`
	RawDescriptor       MemoryRange `json:"raw_descriptor"`

	DeviceClass    UsbClass `json:"device_class,omitempty"`
	DeviceSubclass int      `json:"device_subclass,omitempty"`
	DeviceProtocol int      `json:"device_protocol,omitempty"`

	InterfaceClass    UsbClass `json:"interface_class,omitempty"`
	InterfaceSubclass int      `json:"interface_subclass,omitempty"`
	InterfaceProtocol int      `json:"interface_protocol,omitempty"`

	Country uint `json:"country"`
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
