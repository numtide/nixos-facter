# Smbios

Smbios captures comprehensive information about the system's hardware components
as reported by the System Management BIOS (SMBIOS).


## Fields


### `Bios`

Bios provides detailed information about the system's BIOS, including vendor, version, date, features, and ROM size.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosBios](../hwinfo/smbios_bios.md) | `bios,omitempty` |

### `Board`

Board holds motherboard information such as manufacturer, product, and version.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosBoard](../hwinfo/smbios_board.md) | `board,omitempty` |

### `Cache`

Cache provides details about the system's cache components.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `cache,omitempty` |

### `Chassis`

Chassis holds information related to the system's chassis, including manufacturer, version, and lock presence.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosChassis](../hwinfo/smbios_chassis.md) | `chassis,omitempty` |

### `Config`

Config captures system configuration options.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosConfig](../hwinfo/smbios_config.md) | `config,omitempty` |

### `GroupAssociations`

GroupAssociations lists associations between different hardware groups in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `group_associations,omitempty` |

### `HardwareSecurity`

HardwareSecurity provides information on the system's hardware security configurations.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `hardware_security,omitempty` |

### `Language`

Language contains language-related information, including supported and current languages.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosLanguage](../hwinfo/smbios_language.md) | `language,omitempty` |

### `Memory64Error`

Memory64Error provides information on 64-bit memory errors.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_64_error,omitempty` |

### `MemoryArray`

MemoryArray details the physical memory arrays present in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_array,omitempty` |

### `MemoryArrayMappedAddress`

MemoryArrayMappedAddress provides the mapped addresses of memory arrays.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_array_mapped_address,omitempty` |

### `MemoryDevice`

MemoryDevice captures information about individual memory devices in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_device,omitempty` |

### `MemoryDeviceMappedAddress`

MemoryDeviceMappedAddress provides the mapped addresses of memory devices.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_device_mapped_address,omitempty` |

### `MemoryError`

MemoryError provides information on memory errors detected in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `memory_error,omitempty` |

### `Onboard`

Onboard lists the onboard devices present in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `onboard,omitempty` |

### `PointingDevice`

PointingDevice details the system's pointing devices.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `pointing_device,omitempty` |

### `PortConnector`

PortConnector lists the port connectors present on the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `port_connector,omitempty` |

### `PowerControls`

PowerControls provides information on the power control mechanisms in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `power_controls,omitempty` |

### `Processor`

Processor captures details about the processors used in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `processor,omitempty` |

### `Slot`

Slot lists the expansion slots available in the system.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.Smbios](../hwinfo/smbios.md) | `slot,omitempty` |

### `System`

System captures overall system-related information such as manufacturer, product, version, and UUID.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.SmbiosSystem](../hwinfo/smbios_system.md) | `system,omitempty` |
