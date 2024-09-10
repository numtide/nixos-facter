# SmbiosPortConnector

SmbiosPortConnector captures port connector information.


## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [SmbiosType](smbios_type.md) | `-` |

### `Handle`



| Type | JSON |
| ---- | -----------|
| int | `handle` |

### `PortType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `port_type` |

### `InternalConnectorType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `internal_connector_type,omitempty` |

### `InternalReferenceDesignator`



| Type | JSON |
| ---- | -----------|
| string | `internal_reference_designator,omitempty` |

### `ExternalConnectorType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `external_connector_type,omitempty` |

### `ExternalReferenceDesignator`



| Type | JSON |
| ---- | -----------|
| string | `external_reference_designator,omitempty` |
