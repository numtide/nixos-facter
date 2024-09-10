# DetailUsb



## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [DetailType](detail_type.md) | `-` |

### `Bus`



| Type | JSON |
| ---- | -----------|
| int | `bus` |

### `DeviceNumber`



| Type | JSON |
| ---- | -----------|
| int | `device_number` |

### `Lev`



| Type | JSON |
| ---- | -----------|
| int | `lev` |

### `Parent`



| Type | JSON |
| ---- | -----------|
| int | `parent` |

### `Port`



| Type | JSON |
| ---- | -----------|
| int | `port` |

### `Count`



| Type | JSON |
| ---- | -----------|
| int | `count` |

### `Connections`



| Type | JSON |
| ---- | -----------|
| int | `connections` |

### `UsedConnections`



| Type | JSON |
| ---- | -----------|
| int | `used_connections` |

### `InterfaceDescriptor`



| Type | JSON |
| ---- | -----------|
| int | `interface_descriptor` |

### `Speed`



| Type | JSON |
| ---- | -----------|
| uint | `speed` |

### `Manufacturer`



| Type | JSON |
| ---- | -----------|
| string | `manufacturer,omitempty` |

### `Product`



| Type | JSON |
| ---- | -----------|
| string | `product,omitempty` |

### `Driver`



| Type | JSON |
| ---- | -----------|
| string | `driver,omitempty` |

### `DeviceClass`



| Type | JSON |
| ---- | -----------|
| [UsbClass](usb_class.md) | `device_class,omitempty` |

### `DeviceSubclass`



| Type | JSON |
| ---- | -----------|
| int | `device_subclass,omitempty` |

### `DeviceProtocol`



| Type | JSON |
| ---- | -----------|
| int | `device_protocol,omitempty` |

### `InterfaceClass`



| Type | JSON |
| ---- | -----------|
| [UsbClass](usb_class.md) | `interface_class,omitempty` |

### `InterfaceSubclass`



| Type | JSON |
| ---- | -----------|
| int | `interface_subclass,omitempty` |

### `InterfaceProtocol`



| Type | JSON |
| ---- | -----------|
| int | `interface_protocol,omitempty` |

### `Country`



| Type | JSON |
| ---- | -----------|
| uint | `country` |

### `Vendor`

already included in the parent model, so we omit from JSON output


| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Device`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Revision`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `RawDescriptor`

Seems empty and not really needed, omit for now


| Type | JSON |
| ---- | -----------|
| [MemoryRange](memory_range.md) | `-` |

### `Serial`

Sensitive, omit from JSON output


| Type | JSON |
| ---- | -----------|
| string | `-` |
