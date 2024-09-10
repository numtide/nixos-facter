# Hardware

Hardware represents various hardware components and their details.
The fields are in alphabetical order, and list types are sorted to ensure a consistent and stable report output.


## Fields


### `Bios`

Bios holds the BIOS details of the hardware.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.DetailBios](../hwinfo/detail_bios.md) | `bios,omitempty` |

### `BlockDevice`

BlockDevice holds the list of block devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `block_device,omitempty` |

### `Bluetooth`

Bluetooth holds the list of Bluetooth devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `bluetooth,omitempty` |

### `Braille`

Braille holds the list of Braille devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `braille,omitempty` |

### `Bridge`

Bridge holds the list of bridge devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `bridge,omitempty` |

### `Camera`

Camera holds the list of camera devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `camera,omitempty` |

### `Cdrom`

Cdrom holds the list of CD-ROM devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `cdrom,omitempty` |

### `ChipCard`

ChipCard holds the list of chip card devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `chip_card,omitempty` |

### `Cpu`

Cpu holds the list of CPU details in the hardware.
There is one entry per physical id.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.DetailCpu](../hwinfo/detail_cpu.md) | `cpu,omitempty` |

### `Disk`

Disk holds the list of disk devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `disk,omitempty` |

### `DslAdapter`

DslAdapter holds the list of DSL adapter devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `dsl_adapter,omitempty` |

### `DvbCard`

DvbCard holds the list of DVB card devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `dvb_card,omitempty` |

### `Fingerprint`

Fingerprint holds the list of fingerprint devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `fingerprint,omitempty` |

### `Firewire`

Firewire holds the list of Firewire devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `firewire,omitempty` |

### `FirewireController`

FirewireController holds the list of Firewire controllers in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `firewire_controller,omitempty` |

### `Floppy`

Floppy holds the list of floppy disk devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `floppy,omitempty` |

### `FrameBuffer`

FrameBuffer holds the list of frame buffer devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `frame_buffer,omitempty` |

### `GraphicsCard`

GraphicsCard holds the list of graphics card devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `graphics_card,omitempty` |

### `Hotplug`

Hotplug holds the list of hotplug devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `hotplug,omitempty` |

### `HotplugController`

HotplugController holds the list of hotplug controllers in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `hotplug_controller,omitempty` |

### `Hub`

Hub holds the list of hub devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `hub,omitempty` |

### `Ide`

Ide holds the list of IDE devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `ide,omitempty` |

### `Isapnp`

Isapnp holds the list of ISAPNP devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `isapnp,omitempty` |

### `IsdnAdapter`

IsdnAdapter holds the list of ISDN adapter devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `isdn_adapter,omitempty` |

### `Joystick`

Joystick holds the list of joystick devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `joystick,omitempty` |

### `Keyboard`

Keyboard holds the list of keyboard devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `keyboard,omitempty` |

### `Manual`

Manual holds the list of manually configured devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `manual,omitempty` |

### `Memory`

Memory holds the list of memory devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `memory,omitempty` |

### `MmcController`

MmcController holds the list of MMC controller devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `mmc_controller,omitempty` |

### `Modem`

Modem holds the list of modem devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `modem,omitempty` |

### `Monitor`

Monitor holds the list of monitor devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `monitor,omitempty` |

### `Mouse`

Mouse holds the list of mouse devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `mouse,omitempty` |

### `NetworkController`

NetworkController holds the list of network controller devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `network_controller,omitempty` |

### `NetworkInterface`

NetworkInterface holds the list of network interface devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `network_interface,omitempty` |

### `None`

None holds the list of devices in the hardware with no specific class.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `none,omitempty` |

### `Nvme`

Nvme holds the list of NVMe devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `nvme,omitempty` |

### `Partition`

Partition holds the list of partition devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `partition,omitempty` |

### `Pci`

Pci holds the list of PCI devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `pci,omitempty` |

### `Pcmcia`

Pcmcia holds the list of PCMCIA devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `pcmcia,omitempty` |

### `PcmciaController`

PcmciaController holds the list of PCMCIA controllers in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `pcmcia_controller,omitempty` |

### `Pppoe`

Pppoe holds the list of PPPoE devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `pppoe,omitempty` |

### `Printer`

Printer holds the list of printer devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `printer,omitempty` |

### `Redasd`

Redasd holds the list of REDASD devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `redasd,omitempty` |

### `Scanner`

Scanner holds the list of scanner devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `scanner,omitempty` |

### `Scsi`

Scsi holds the list of SCSI devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `scsi,omitempty` |

### `Sound`

Sound holds the list of sound devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `sound,omitempty` |

### `StorageController`

StorageController holds the list of storage controller devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `storage_controller,omitempty` |

### `System`

System holds the system details of the hardware.


| Type | JSON |
| ---- | -----------|
| *[hwinfo.DetailSys](../hwinfo/detail_sys.md) | `system,omitempty` |

### `Tape`

Tape holds the list of tape devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `tape,omitempty` |

### `TvCard`

TvCard holds the list of TV card devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `tv_card,omitempty` |

### `Unknown`

Unknown holds the list of unknown devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `unknown,omitempty` |

### `Usb`

Usb holds the list of USB devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `usb,omitempty` |

### `UsbController`

UsbController holds the list of USB controller devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `usb_controller,omitempty` |

### `VesaBios`

VesaBios holds the list of VESA BIOS devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `vesa_bios,omitempty` |

### `WlanCard`

WlanCard holds the list of WLAN card devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `wlan_card,omitempty` |

### `Zip`

Zip holds the list of ZIP drive devices in the hardware.


| Type | JSON |
| ---- | -----------|
| [][hwinfo.HardwareDevice](../hwinfo/hardware_device.md) | `zip,omitempty` |
