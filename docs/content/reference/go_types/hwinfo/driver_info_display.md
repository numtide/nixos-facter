# DriverInfoDisplay



## Fields


### `Type`



| Type | JSON |
| ---- | -----------|
| [DriverInfoType](driver_info_type.md) | `type,omitempty` |

### `DbEntry0`

actual driver database entries


| Type | JSON |
| ---- | -----------|
| []string | `db_entry_0,omitempty` |

### `DbEntry1`



| Type | JSON |
| ---- | -----------|
| []string | `db_entry_1,omitempty` |

### `Width`



| Type | JSON |
| ---- | -----------|
| uint | `width` |

### `Height`



| Type | JSON |
| ---- | -----------|
| uint | `height` |

### `VerticalSync`



| Type | JSON |
| ---- | -----------|
| [SyncRange](sync_range.md) | `vertical_sync` |

### `HorizontalSync`



| Type | JSON |
| ---- | -----------|
| [SyncRange](sync_range.md) | `horizontal_sync` |

### `Bandwidth`



| Type | JSON |
| ---- | -----------|
| uint | `bandwidth` |

### `HorizontalSyncTimings`



| Type | JSON |
| ---- | -----------|
| [SyncTimings](sync_timings.md) | `horizontal_sync_timings` |

### `VerticalSyncTimings`



| Type | JSON |
| ---- | -----------|
| [SyncTimings](sync_timings.md) | `vertical_sync_timings` |

### `HorizontalFlag`



| Type | JSON |
| ---- | -----------|
| byte | `horizontal_flag` |

### `VerticalFlag`



| Type | JSON |
| ---- | -----------|
| byte | `vertical_flag` |
