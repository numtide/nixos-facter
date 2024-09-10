# DetailPci



## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [DetailType](detail_type.md) | `-` |

### `Flags`



| Type | JSON |
| ---- | -----------|
| [][PciFlag](pci_flag.md) | `flags,omitempty` |

### `Function`



| Type | JSON |
| ---- | -----------|
| uint | `function` |

### `Command`

todo map pci constants from pci.h?


| Type | JSON |
| ---- | -----------|
| uint | `command` |

### `HeaderType`



| Type | JSON |
| ---- | -----------|
| uint | `header_type` |

### `SecondaryBus`



| Type | JSON |
| ---- | -----------|
| uint | `secondary_bus` |

### `Irq`



| Type | JSON |
| ---- | -----------|
| uint | `irq` |

### `ProgIf`

Programming Interface Byte: a read-only register that specifies a register-level programming interface for the device.


| Type | JSON |
| ---- | -----------|
| uint | `prog_if` |

### `Bus`

already included in the parent model, so we omit from JSON output


| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Slot`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `BaseClass`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `SubClass`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Device`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Vendor`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `SubDevice`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `SubVendor`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Revision`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `BaseAddress`



| Type | JSON |
| ---- | -----------|
| []uint64 | `-` |

### `BaseLength`



| Type | JSON |
| ---- | -----------|
| []uint64 | `-` |

### `AddressFlags`



| Type | JSON |
| ---- | -----------|
| []uint | `-` |

### `RomBaseAddress`



| Type | JSON |
| ---- | -----------|
| uint64 | `-` |

### `RomBaseLength`



| Type | JSON |
| ---- | -----------|
| uint64 | `-` |

### `SysfsId`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `SysfsBusId`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `ModuleAlias`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `Label`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `Data`

Omit from JSON output


| Type | JSON |
| ---- | -----------|
| string | `-` |

### `DataLength`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `DataExtLength`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Log`



| Type | JSON |
| ---- | -----------|
| string | `-` |
