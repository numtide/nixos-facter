# SmbiosChassis

SmbiosChassis captures motherboard related information.


## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [SmbiosType](smbios_type.md) | `-` |

### `Handle`



| Type | JSON |
| ---- | -----------|
| int | `handle` |

### `Manufacturer`



| Type | JSON |
| ---- | -----------|
| string | `manufacturer` |

### `Version`



| Type | JSON |
| ---- | -----------|
| string | `version` |

### `Serial`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `AssetTag`



| Type | JSON |
| ---- | -----------|
| string | `-` |

### `ChassisType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `chassis_type` |

### `LockPresent`



| Type | JSON |
| ---- | -----------|
| bool | `lock_present` |

### `BootupState`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `bootup_state` |

### `PowerState`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `power_state` |

### `ThermalState`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `thermal_state` |

### `SecurityState`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `security_state` |

### `OEM`



| Type | JSON |
| ---- | -----------|
| string | `oem` |
