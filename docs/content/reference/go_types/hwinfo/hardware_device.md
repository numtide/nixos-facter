# HardwareDevice



## Fields


### `Index`

Index is a unique index provided by hwinfo, starting at 1


| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `BusType`

Bus type (id and name)


| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `Slot`



| Type | JSON |
| ---- | -----------|
| [Slot](slot.md) | `` |

### `BaseClass`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `SubClass`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `PciInterface`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `Vendor`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `SubVendor`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `Device`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `SubDevice`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `Revision`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `Serial`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `CompatVendor`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `CompatDevice`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `` |

### `HardwareClass`



| Type | JSON |
| ---- | -----------|
| [HardwareClass](hardware_class.md) | `` |

### `Model`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `AttachedTo`



| Type | JSON |
| ---- | -----------|
| uint | `` |

### `SysfsId`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `SysfsBusId`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `SysfsIOMMUGroupId`



| Type | JSON |
| ---- | -----------|
| int | `` |

### `SysfsDeviceLink`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `UnixDeviceName`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `UnixDeviceNumber`



| Type | JSON |
| ---- | -----------|
| *[DeviceNumber](device_number.md) | `` |

### `UnixDeviceNames`



| Type | JSON |
| ---- | -----------|
| []string | `` |

### `UnixDeviceName2`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `UnixDeviceNumber2`



| Type | JSON |
| ---- | -----------|
| *[DeviceNumber](device_number.md) | `` |

### `RomId`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `Udi`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `ParentUdi`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `UniqueId`

		UniqueId is a unique string identifying this hardware.
		The string consists of two parts separated by a dot (".").
		The part before the dot describes the location (where the hardware is attached in the system).
		The part after the dot identifies the hardware itself.
		The string must not contain slashes ("/") because we're going to create files with this id as name.
		Apart from this, there are no restrictions on the form of this string.


| Type | JSON |
| ---- | -----------|
| string | `` |

### `UniqueIds`



| Type | JSON |
| ---- | -----------|
| []string | `` |

### `Resources`



| Type | JSON |
| ---- | -----------|
| [][Resource](resource.md) | `` |

### `Detail`



| Type | JSON |
| ---- | -----------|
| [Detail](detail.md) | `` |

### `Hotplug`



| Type | JSON |
| ---- | -----------|
| [Hotplug](hotplug.md) | `` |

### `HotplugSlot`



| Type | JSON |
| ---- | -----------|
| uint | `` |

### `Is`



| Type | JSON |
| ---- | -----------|
| [Is](is.md) | `` |

### `Driver`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `DriverModule`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `Drivers`



| Type | JSON |
| ---- | -----------|
| []string | `` |

### `DriverModules`



| Type | JSON |
| ---- | -----------|
| []string | `` |

### `DriverInfo`



| Type | JSON |
| ---- | -----------|
| [DriverInfo](driver_info.md) | `` |

### `UsbGuid`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `Requires`



| Type | JSON |
| ---- | -----------|
| []string | `` |

### `ModuleAlias`



| Type | JSON |
| ---- | -----------|
| string | `` |

### `Label`



| Type | JSON |
| ---- | -----------|
| string | `` |
