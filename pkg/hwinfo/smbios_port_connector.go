package hwinfo

/*
#cgo pkg-config: hwinfo
#include <hd.h>
*/
import "C"

// SmbiosPortConnector captures port connector information.
type SmbiosPortConnector struct {
	Type                        SmbiosType `json:"-"`
	Handle                      int        `json:"handle"`
	PortType                    *Id        `json:"port_type"`
	InternalConnectorType       *Id        `json:"internal_connector_type,omitempty"`
	InternalReferenceDesignator string     `json:"internal_reference_designator,omitempty"`
	ExternalConnectorType       *Id        `json:"external_connector_type,omitempty"`
	ExternalReferenceDesignator string     `json:"external_reference_designator,omitempty"`
}

func (s SmbiosPortConnector) SmbiosType() SmbiosType {
	return s.Type
}

func NewSmbiosConnect(info C.smbios_connect_t) (Smbios, error) {
	return SmbiosPortConnector{
		Type:                        SmbiosTypePortConnector,
		Handle:                      int(info.handle),
		PortType:                    NewId(info.port_type),
		InternalConnectorType:       NewId(info.i_type),
		InternalReferenceDesignator: C.GoString(info.i_des),
		ExternalConnectorType:       NewId(info.x_type),
		ExternalReferenceDesignator: C.GoString(info.x_des),
	}, nil
}
