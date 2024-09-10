# Scanner

Scanner defines a type responsible for scanning and reporting system hardware information.


## Fields


### `Swap`

Swap indicates whether the system swap information should be reported.


| Type | JSON |
| ---- | -----------|
| bool | `` |

### `Ephemeral`

Ephemeral indicates whether the scanner should report ephemeral details,
such as swap.


| Type | JSON |
| ---- | -----------|
| bool | `` |

### `Features`

Features is a list of ProbeFeature types that should be scanned for.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.ProbeFeature](../hwinfo/probe_feature.md) | `` |
