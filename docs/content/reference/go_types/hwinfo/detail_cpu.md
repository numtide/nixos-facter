# DetailCpu



## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [DetailType](detail_type.md) | `-` |

### `Architecture`



| Type | JSON |
| ---- | -----------|
| [CpuArch](cpu_arch.md) | `architecture` |

### `VendorName`



| Type | JSON |
| ---- | -----------|
| string | `vendor_name,omitempty` |

### `ModelName`



| Type | JSON |
| ---- | -----------|
| string | `model_name,omitempty` |

### `Family`



| Type | JSON |
| ---- | -----------|
| uint | `family` |

### `Model`



| Type | JSON |
| ---- | -----------|
| uint | `model` |

### `Stepping`



| Type | JSON |
| ---- | -----------|
| uint | `stepping` |

### `Platform`



| Type | JSON |
| ---- | -----------|
| string | `platform,omitempty` |

### `Features`



| Type | JSON |
| ---- | -----------|
| []string | `features,omitempty` |

### `Bugs`



| Type | JSON |
| ---- | -----------|
| []string | `bugs,omitempty` |

### `PowerManagement`



| Type | JSON |
| ---- | -----------|
| []string | `power_management,omitempty` |

### `Bogo`



| Type | JSON |
| ---- | -----------|
| float64 | `bogo` |

### `Cache`



| Type | JSON |
| ---- | -----------|
| uint | `cache,omitempty` |

### `Units`



| Type | JSON |
| ---- | -----------|
| uint | `units,omitempty` |

### `Clock`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `PhysicalId`

x86 only fields


| Type | JSON |
| ---- | -----------|
| uint | `physical_id` |

### `Siblings`



| Type | JSON |
| ---- | -----------|
| uint | `siblings,omitempty` |

### `Cores`



| Type | JSON |
| ---- | -----------|
| uint | `cores,omitempty` |

### `CoreId`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `Fpu`



| Type | JSON |
| ---- | -----------|
| bool | `fpu` |

### `FpuException`



| Type | JSON |
| ---- | -----------|
| bool | `fpu_exception` |

### `CpuidLevel`



| Type | JSON |
| ---- | -----------|
| uint | `cpuid_level,omitempty` |

### `WriteProtect`



| Type | JSON |
| ---- | -----------|
| bool | `write_protect` |

### `TlbSize`



| Type | JSON |
| ---- | -----------|
| uint | `tlb_size,omitempty` |

### `ClflushSize`



| Type | JSON |
| ---- | -----------|
| uint | `clflush_size,omitempty` |

### `CacheAlignment`



| Type | JSON |
| ---- | -----------|
| int | `cache_alignment,omitempty` |

### `AddressSizes`



| Type | JSON |
| ---- | -----------|
| [AddressSizes](address_sizes.md) | `address_sizes,omitempty` |

### `Apicid`



| Type | JSON |
| ---- | -----------|
| uint | `-` |

### `ApicidInitial`



| Type | JSON |
| ---- | -----------|
| uint | `-` |
