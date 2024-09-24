package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/numtide/nixos-facter/pkg/facter"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
)

var (
	outputPath       string
	logLevel         string
	hardwareFeatures []string

	scanner = facter.Scanner{}
)

const usage = `nixos-facter [flags]
Hardware report generator

Usage:
  nixos-facter [flags]

Flags:
  --ephemeral          capture all ephemeral properties e.g. swap, filesystems and so on
  --hardware strings   Hardware items to probe. Possible values are memory,pci,isapnp,net,floppy,misc,misc_serial,misc_par,misc_floppy,serial,cpu,bios,monitor,mouse,scsi,usb,usb_mods,adb,modem,modem_usb,parallel,parallel_lp,parallel_zip,isa,isa_isdn,isdn,kbd,prom,sbus,braille,braille_alva,braille_fhp,braille_ht,ignx11,sys,bios_vbe,isapnp_old,isapnp_new,isapnp_mod,braille_baum,manual,fb,veth,pppoe,scan,pcmcia,fork,parallel_imm,s390,cpuemu,sysfs,s390disks,udev,block,block_cdrom,block_part,edd,edd_mod,bios_ddc,bios_fb,bios_mode,input,block_mods,bios_vesa,cpuemu_debug,scsi_noserial,wlan,bios_crc,hal,bios_vram,bios_acpi,bios_ddc_ports,modules_pata,net_eeprom,x86emu,max,lxrc,all,, (default [memory,pci,net,serial,cpu,bios,monitor,scsi,usb,prom,sbus,sys,sysfs,udev,block,wlan])
  -h, --help           help for nixos-facter
  -o, --output string  path to write the report
  --swap               capture swap entries

`

func init() {
	// Define flags
	flag.StringVar(&outputPath, "output", "", "path to write the report")
	flag.StringVar(&outputPath, "o", "", "path to write the report")
	flag.BoolVar(&scanner.Swap, "swap", false, "capture swap entries")
	flag.BoolVar(&scanner.Ephemeral, "ephemeral", false, "capture all ephemeral properties e.g. swap, filesystems and so on")
	flag.StringVar(&logLevel, "log-level", "info", "log level")

	defaultFeatures := []string{
		"memory", "pci", "net", "serial", "cpu", "bios", "monitor", "scsi", "usb", "prom", "sbus", "sys", "sysfs",
		"udev", "block", "wlan",
	}

	probeFeatures := hwinfo.ProbeFeatureStrings()
	filteredFeatures := []string{}
	for _, feature := range probeFeatures {
		if feature != "default" && feature != "int" {
			filteredFeatures = append(filteredFeatures, feature)
		}
	}

	hardwareFeatures = defaultFeatures
	flag.Func("hardware", fmt.Sprintf("Hardware items to probe. Possible values are %s", strings.Join(filteredFeatures, ",")), func(flagValue string) error {
		hardwareFeatures = strings.Split(flagValue, ",")
		return nil
	})

	// Custom usage function
	flag.Usage = func() { fmt.Fprintf(os.Stderr, "%s\n", usage) }
}

func Execute() {
	flag.Parse()

	// Check if the effective user id is 0 e.g. root
	if os.Geteuid() != 0 {
		log.Fatalf("you must run this program as root")
	}

	// Convert the hardware features into probe features
	for _, str := range hardwareFeatures {
		probe, err := hwinfo.ProbeFeatureString(str)
		if err != nil {
			log.Fatalf("invalid hardware feature: %v", err)
		}
		scanner.Features = append(scanner.Features, probe)
	}

	// Set the log level
	switch logLevel {
	case "debug":
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	case "info":
		log.SetFlags(log.LstdFlags)
	case "warn", "error":
		log.SetFlags(0)
	default:
		log.Fatalf("invalid log level: %s", logLevel)
	}

	report, err := scanner.Scan()
	if err != nil {
		log.Fatalf("failed to scan: %v", err)
	}

	bytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("failed to marshal report to json: %v", err)
	}

	// If a file path is provided write the report to it, otherwise output the report on stdout
	if outputPath == "" {
		if _, err = os.Stdout.Write(bytes); err != nil {
			log.Fatalf("failed to write report to stdout: %v", err)
		}
		fmt.Println()
	} else if err = os.WriteFile(outputPath, bytes, 0o644); err != nil {
		log.Fatalf("failed to write report to output path: %v", err)
	}
}
