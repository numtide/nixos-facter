# SwapEntry

SwapEntry represents a swap entry.


## Fields


### `Type`

Type is the type of swap e.g. partition or file.


| Type | JSON |
| ---- | -----------|
| [SwapType](swap_type.md) | `type` |

### `Filename`

Filename is the path to the swap device or file.


| Type | JSON |
| ---- | -----------|
| string | `path` |

### `Size`

Size is the total size of the swap in kilobytes.


| Type | JSON |
| ---- | -----------|
| uint64 | `size` |

### `Used`

Used is the amount of swap space currently in use, in kilobytes.


| Type | JSON |
| ---- | -----------|
| uint64 | `used` |

### `Priority`

Priority determines the order in which swap spaces are used.
Higher numbers have higher priority.


| Type | JSON |
| ---- | -----------|
| int32 | `priority` |
