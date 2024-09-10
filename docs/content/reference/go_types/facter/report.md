# Report

Report represents a detailed report on the system’s hardware, virtualisation, SMBios, and swap entries.


## Fields


### `Version`

Version is a monotonically increasing number,
used to indicate breaking changes or new features in the report output.


| Type | JSON |
| ---- | -----------|
| uint | `version` |

### `System`

System indicates the system architecture e.g. x86_64-linux.


| Type | JSON |
| ---- | -----------|
| string | `system` |

### `Virtualisation`

Virtualisation indicates the type of virtualisation or container environment present on the system.


| Type | JSON |
| ---- | -----------|
| [virt.Type](../virt/type.md) | `virtualisation` |

### `Hardware`

Hardware provides detailed information about the system’s hardware components, such as CPU, memory, and peripherals.


| Type | JSON |
| ---- | -----------|
| [Hardware](hardware.md) | `hardware,omitempty` |

### `Smbios`

Smbios provides detailed information about the system's SMBios data, such as BIOS, board, chassis, memory, and processors.


| Type | JSON |
| ---- | -----------|
| [Smbios](smbios.md) | `smbios,omitempty` |

### `Swap`

Swap contains a list of swap entries representing the system's swap devices or files and their respective details.


| Type | JSON |
| ---- | -----------|
| []*[ephem.SwapEntry](../ephem/swap_entry.md) | `swap,omitempty` |
