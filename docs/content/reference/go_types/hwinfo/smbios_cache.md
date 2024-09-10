# SmbiosCache

SmbiosCache captures processor information.


## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [SmbiosType](smbios_type.md) | `-` |

### `Handle`



| Type | JSON |
| ---- | -----------|
| int | `handle` |

### `Socket`



| Type | JSON |
| ---- | -----------|
| string | `socket` |

### `SizeMax`



| Type | JSON |
| ---- | -----------|
| uint | `size_max` |

### `SizeCurrent`



| Type | JSON |
| ---- | -----------|
| uint | `size_current` |

### `Speed`



| Type | JSON |
| ---- | -----------|
| uint | `speed` |

### `Mode`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `mode` |

### `Enabled`



| Type | JSON |
| ---- | -----------|
| bool | `enabled` |

### `Location`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `location` |

### `Socketed`



| Type | JSON |
| ---- | -----------|
| bool | `socketed` |

### `Level`



| Type | JSON |
| ---- | -----------|
| uint | `level` |

### `ECC`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `ecc` |

### `CacheType`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `cache_type` |

### `Associativity`



| Type | JSON |
| ---- | -----------|
| *[Id](id.md) | `associativity` |

### `SRAMType`



| Type | JSON |
| ---- | -----------|
| []string | `sram_type_current` |

### `SRAMTypes`



| Type | JSON |
| ---- | -----------|
| []string | `sram_type_supported` |
