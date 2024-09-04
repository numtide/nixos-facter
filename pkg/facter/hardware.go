package facter

import (
	"fmt"
	"slices"

	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

type Hardware struct {
	Bios               *hwinfo.DetailBios      `json:"bios,omitempty"`
	BlockDevice        []hwinfo.HardwareDevice `json:"block_device,omitempty"`
	Bluetooth          []hwinfo.HardwareDevice `json:"bluetooth,omitempty"`
	Braille            []hwinfo.HardwareDevice `json:"braille,omitempty"`
	Bridge             []hwinfo.HardwareDevice `json:"bridge,omitempty"`
	Camera             []hwinfo.HardwareDevice `json:"camera,omitempty"`
	Cdrom              []hwinfo.HardwareDevice `json:"cdrom,omitempty"`
	ChipCard           []hwinfo.HardwareDevice `json:"chip_card,omitempty"`
	Cpu                []hwinfo.DetailCpu      `json:"cpu,omitempty"`
	Disk               []hwinfo.HardwareDevice `json:"disk,omitempty"`
	DslAdapter         []hwinfo.HardwareDevice `json:"dsl_adapter,omitempty"`
	DvbCard            []hwinfo.HardwareDevice `json:"dvb_card,omitempty"`
	Fingerprint        []hwinfo.HardwareDevice `json:"fingerprint,omitempty"`
	Firewire           []hwinfo.HardwareDevice `json:"firewire,omitempty"`
	FirewireController []hwinfo.HardwareDevice `json:"firewire_controller,omitempty"`
	Floppy             []hwinfo.HardwareDevice `json:"floppy,omitempty"`
	FrameBuffer        []hwinfo.HardwareDevice `json:"frame_buffer,omitempty"`
	GraphicsCard       []hwinfo.HardwareDevice `json:"graphics_card,omitempty"`
	Hotplug            []hwinfo.HardwareDevice `json:"hotplug,omitempty"`
	HotplugController  []hwinfo.HardwareDevice `json:"hotplug_controller,omitempty"`
	Hub                []hwinfo.HardwareDevice `json:"hub,omitempty"`
	Ide                []hwinfo.HardwareDevice `json:"ide,omitempty"`
	Isapnp             []hwinfo.HardwareDevice `json:"isapnp,omitempty"`
	IsdnAdapter        []hwinfo.HardwareDevice `json:"isdn_adapter,omitempty"`
	Joystick           []hwinfo.HardwareDevice `json:"joystick,omitempty"`
	Keyboard           []hwinfo.HardwareDevice `json:"keyboard,omitempty"`
	Manual             []hwinfo.HardwareDevice `json:"manual,omitempty"`
	Memory             []hwinfo.HardwareDevice `json:"memory,omitempty"`
	MmcController      []hwinfo.HardwareDevice `json:"mmc_controller,omitempty"`
	Modem              []hwinfo.HardwareDevice `json:"modem,omitempty"`
	Monitor            []hwinfo.HardwareDevice `json:"monitor,omitempty"`
	Mouse              []hwinfo.HardwareDevice `json:"mouse,omitempty"`
	NetworkController  []hwinfo.HardwareDevice `json:"network_controller,omitempty"`
	NetworkInterface   []hwinfo.HardwareDevice `json:"network_interface,omitempty"`
	None               []hwinfo.HardwareDevice `json:"none,omitempty"`
	Nvme               []hwinfo.HardwareDevice `json:"nvme,omitempty"`
	Partition          []hwinfo.HardwareDevice `json:"partition,omitempty"`
	Pci                []hwinfo.HardwareDevice `json:"pci,omitempty"`
	Pcmcia             []hwinfo.HardwareDevice `json:"pcmcia,omitempty"`
	PcmciaController   []hwinfo.HardwareDevice `json:"pcmcia_controller,omitempty"`
	Pppoe              []hwinfo.HardwareDevice `json:"pppoe,omitempty"`
	Printer            []hwinfo.HardwareDevice `json:"printer,omitempty"`
	Redasd             []hwinfo.HardwareDevice `json:"redasd,omitempty"`
	Scanner            []hwinfo.HardwareDevice `json:"scanner,omitempty"`
	Scsi               []hwinfo.HardwareDevice `json:"scsi,omitempty"`
	Sound              []hwinfo.HardwareDevice `json:"sound,omitempty"`
	StorageController  []hwinfo.HardwareDevice `json:"storage_controller,omitempty"`
	System             *hwinfo.DetailSys       `json:"system,omitempty"`
	Tape               []hwinfo.HardwareDevice `json:"tape,omitempty"`
	TvCard             []hwinfo.HardwareDevice `json:"tv_card,omitempty"`
	Unknown            []hwinfo.HardwareDevice `json:"unknown,omitempty"`
	Usb                []hwinfo.HardwareDevice `json:"usb,omitempty"`
	UsbController      []hwinfo.HardwareDevice `json:"usb_controller,omitempty"`
	VesaBios           []hwinfo.HardwareDevice `json:"vesa_bios,omitempty"`
	WlanCard           []hwinfo.HardwareDevice `json:"wlan_card,omitempty"`
	Zip                []hwinfo.HardwareDevice `json:"zip,omitempty"`
}

func compareDevice(a hwinfo.HardwareDevice, b hwinfo.HardwareDevice) int {
	return int(a.Index - b.Index)
}

func (h *Hardware) add(device hwinfo.HardwareDevice) error {
	switch device.HardwareClass {
	case hwinfo.HardwareClassBios:
		if h.Bios != nil {
			return fmt.Errorf("bios field is already set")
		} else if bios, ok := device.Detail.(hwinfo.DetailBios); !ok {
			return fmt.Errorf("expected hwinfo.DetailBios, found %T", device.Detail)
		} else {
			h.Bios = &bios
		}
	case hwinfo.HardwareClassBlockDevice:
		h.BlockDevice = append(h.BlockDevice, device)
		slices.SortFunc(h.BlockDevice, compareDevice)
	case hwinfo.HardwareClassBluetooth:
		h.Bluetooth = append(h.Bluetooth, device)
		slices.SortFunc(h.Bluetooth, compareDevice)
	case hwinfo.HardwareClassBraille:
		h.Braille = append(h.Braille, device)
		slices.SortFunc(h.Braille, compareDevice)
	case hwinfo.HardwareClassBridge:
		h.Bridge = append(h.Bridge, device)
		slices.SortFunc(h.Bridge, compareDevice)
	case hwinfo.HardwareClassCamera:
		h.Camera = append(h.Camera, device)
		slices.SortFunc(h.Camera, compareDevice)
	case hwinfo.HardwareClassCdrom:
		h.Cdrom = append(h.Cdrom, device)
		slices.SortFunc(h.Cdrom, compareDevice)
	case hwinfo.HardwareClassChipCard:
		h.ChipCard = append(h.ChipCard, device)
		slices.SortFunc(h.ChipCard, compareDevice)
	case hwinfo.HardwareClassCpu:
		cpu, ok := device.Detail.(hwinfo.DetailCpu)
		if !ok {
			return fmt.Errorf("expected hwinfo.DetailCpu, found %T", device.Detail)
		}

		// We insert by physical id, as we only want one entry per core.
		requiredSize := int(cpu.PhysicalId) + 1
		if len(h.Cpu) < requiredSize {
			newItems := make([]hwinfo.DetailCpu, requiredSize-len(h.Cpu))
			h.Cpu = append(h.Cpu, newItems...)
		}
		h.Cpu[cpu.PhysicalId] = cpu

		// Sort in ascending order to ensure a stable output
		slices.SortFunc(h.Cpu, func(a, b hwinfo.DetailCpu) int {
			return int(a.PhysicalId - b.PhysicalId)
		})

	case hwinfo.HardwareClassDisk:
		h.Disk = append(h.Disk, device)
		slices.SortFunc(h.Disk, compareDevice)
	case hwinfo.HardwareClassDslAdapter:
		h.DslAdapter = append(h.DslAdapter, device)
		slices.SortFunc(h.DslAdapter, compareDevice)
	case hwinfo.HardwareClassDvbCard:
		h.DvbCard = append(h.DvbCard, device)
		slices.SortFunc(h.DvbCard, compareDevice)
	case hwinfo.HardwareClassFingerprint:
		h.Fingerprint = append(h.Fingerprint, device)
		slices.SortFunc(h.Fingerprint, compareDevice)
	case hwinfo.HardwareClassFirewire:
		h.Firewire = append(h.Firewire, device)
		slices.SortFunc(h.Firewire, compareDevice)
	case hwinfo.HardwareClassFirewireController:
		h.FirewireController = append(h.FirewireController, device)
		slices.SortFunc(h.FirewireController, compareDevice)
	case hwinfo.HardwareClassFloppy:
		h.Floppy = append(h.Floppy, device)
		slices.SortFunc(h.Floppy, compareDevice)
	case hwinfo.HardwareClassFrameBuffer:
		h.FrameBuffer = append(h.FrameBuffer, device)
		slices.SortFunc(h.FrameBuffer, compareDevice)
	case hwinfo.HardwareClassGraphicsCard:
		h.GraphicsCard = append(h.GraphicsCard, device)
		slices.SortFunc(h.GraphicsCard, compareDevice)
	case hwinfo.HardwareClassHotplug:
		h.Hotplug = append(h.Hotplug, device)
		slices.SortFunc(h.Hotplug, compareDevice)
	case hwinfo.HardwareClassHotplugController:
		h.HotplugController = append(h.HotplugController, device)
		slices.SortFunc(h.HotplugController, compareDevice)
	case hwinfo.HardwareClassHub:
		h.Hub = append(h.Hub, device)
		slices.SortFunc(h.Hub, compareDevice)
	case hwinfo.HardwareClassIde:
		h.Ide = append(h.Ide, device)
		slices.SortFunc(h.Ide, compareDevice)
	case hwinfo.HardwareClassIsapnp:
		h.Isapnp = append(h.Isapnp, device)
		slices.SortFunc(h.Isapnp, compareDevice)
	case hwinfo.HardwareClassIsdnAdapter:
		h.IsdnAdapter = append(h.IsdnAdapter, device)
		slices.SortFunc(h.IsdnAdapter, compareDevice)
	case hwinfo.HardwareClassJoystick:
		h.Joystick = append(h.Joystick, device)
		slices.SortFunc(h.Joystick, compareDevice)
	case hwinfo.HardwareClassKeyboard:
		h.Keyboard = append(h.Keyboard, device)
		slices.SortFunc(h.Keyboard, compareDevice)
	case hwinfo.HardwareClassManual:
		h.Manual = append(h.Manual, device)
		slices.SortFunc(h.Manual, compareDevice)
	case hwinfo.HardwareClassMemory:
		h.Memory = append(h.Memory, device)
		slices.SortFunc(h.Memory, compareDevice)
	case hwinfo.HardwareClassMmcController:
		h.MmcController = append(h.MmcController, device)
		slices.SortFunc(h.MmcController, compareDevice)
	case hwinfo.HardwareClassModem:
		h.Modem = append(h.Modem, device)
		slices.SortFunc(h.Modem, compareDevice)
	case hwinfo.HardwareClassMonitor:
		h.Monitor = append(h.Monitor, device)
		slices.SortFunc(h.Monitor, compareDevice)
	case hwinfo.HardwareClassMouse:
		h.Mouse = append(h.Mouse, device)
		slices.SortFunc(h.Mouse, compareDevice)
	case hwinfo.HardwareClassNetworkController:
		h.NetworkController = append(h.NetworkController, device)
		slices.SortFunc(h.NetworkController, compareDevice)
	case hwinfo.HardwareClassNetworkInterface:
		h.NetworkInterface = append(h.NetworkInterface, device)
		slices.SortFunc(h.NetworkInterface, compareDevice)
	case hwinfo.HardwareClassNone:
		h.None = append(h.None, device)
		slices.SortFunc(h.None, compareDevice)
	case hwinfo.HardwareClassNvme:
		h.Nvme = append(h.Nvme, device)
		slices.SortFunc(h.Nvme, compareDevice)
	case hwinfo.HardwareClassPartition:
		h.Partition = append(h.Partition, device)
		slices.SortFunc(h.Partition, compareDevice)
	case hwinfo.HardwareClassPci:
		h.Pci = append(h.Pci, device)
		slices.SortFunc(h.Pci, compareDevice)
	case hwinfo.HardwareClassPcmcia:
		h.Pcmcia = append(h.Pcmcia, device)
		slices.SortFunc(h.Pcmcia, compareDevice)
	case hwinfo.HardwareClassPcmciaController:
		h.PcmciaController = append(h.PcmciaController, device)
		slices.SortFunc(h.PcmciaController, compareDevice)
	case hwinfo.HardwareClassPppoe:
		h.Pppoe = append(h.Pppoe, device)
		slices.SortFunc(h.Pppoe, compareDevice)
	case hwinfo.HardwareClassPrinter:
		h.Printer = append(h.Printer, device)
		slices.SortFunc(h.Printer, compareDevice)
	case hwinfo.HardwareClassRedasd:
		h.Redasd = append(h.Redasd, device)
		slices.SortFunc(h.Redasd, compareDevice)
	case hwinfo.HardwareClassScanner:
		h.Scanner = append(h.Scanner, device)
		slices.SortFunc(h.Scanner, compareDevice)
	case hwinfo.HardwareClassScsi:
		h.Scsi = append(h.Scsi, device)
		slices.SortFunc(h.Scsi, compareDevice)
	case hwinfo.HardwareClassSound:
		h.Sound = append(h.Sound, device)
		slices.SortFunc(h.Sound, compareDevice)
	case hwinfo.HardwareClassStorageController:
		h.StorageController = append(h.StorageController, device)
		slices.SortFunc(h.StorageController, compareDevice)
	case hwinfo.HardwareClassSystem:
		if h.System != nil {
			return fmt.Errorf("system field is already set")
		} else if system, ok := device.Detail.(hwinfo.DetailSys); !ok {
			return fmt.Errorf("expected hwinfo.DetailSys, found %T", device.Detail)
		} else {
			h.System = &system
		}
	case hwinfo.HardwareClassTape:
		h.Tape = append(h.Tape, device)
		slices.SortFunc(h.Tape, compareDevice)
	case hwinfo.HardwareClassTvCard:
		h.TvCard = append(h.TvCard, device)
		slices.SortFunc(h.TvCard, compareDevice)
	case hwinfo.HardwareClassUnknown:
		h.Unknown = append(h.Unknown, device)
		slices.SortFunc(h.Unknown, compareDevice)
	case hwinfo.HardwareClassUsb:
		h.Usb = append(h.Usb, device)
		slices.SortFunc(h.Usb, compareDevice)
	case hwinfo.HardwareClassUsbController:
		h.UsbController = append(h.UsbController, device)
		slices.SortFunc(h.UsbController, compareDevice)
	case hwinfo.HardwareClassVesaBios:
		h.VesaBios = append(h.VesaBios, device)
		slices.SortFunc(h.VesaBios, compareDevice)
	case hwinfo.HardwareClassWlanCard:
		h.WlanCard = append(h.WlanCard, device)
		slices.SortFunc(h.WlanCard, compareDevice)
	case hwinfo.HardwareClassZip:
		h.Zip = append(h.Zip, device)
		slices.SortFunc(h.Zip, compareDevice)
	case hwinfo.HardwareClassAll:
		// Do nothing, this is used by the hwinfo cli exclusively.
	}
	return nil
}
