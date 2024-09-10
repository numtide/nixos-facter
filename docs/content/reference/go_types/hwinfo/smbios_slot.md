# SmbiosSlot

SmbiosSlot captures system slot information.


## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [SmbiosType](smbios_type.md) | `-` |

### `Handle`



| Type | JSON |
| ---- | -----------|
| int | `handle` |

### `Designation`



| Type | JSON |
| ---- | -----------|
| string | `designation,omitempty` |

### `SlotType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `slot_type` |

### `BusWidth`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `bus_width` |

### `Usage`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `usage` |

### `Length`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `length` |

### `Id`



| Type | JSON |
| ---- | -----------|
| uint | `id` |

### `Features`



| Type | JSON |
| ---- | -----------|
| []string | `features,omitempty` |
