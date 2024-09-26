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

	flag.Func("hardware", "Hardware items to probe (comma separated).", func(flagValue string) error {
		hardwareFeatures = strings.Split(flagValue, ",")
		return nil
	})
	possibleValues := strings.Join(filteredFeatures, ",")
	defaultValues := strings.Join(defaultFeatures, ",")
	const usage = `nixos-facter [flags]
Hardware report generator

Usage:
  nixos-facter [flags]

Flags:
  --ephemeral          capture all ephemeral properties e.g. swap, filesystems and so on
  -h, --help           help for nixos-facter
  -o, --output string  path to write the report
  --swap               capture swap entries
  --hardware strings   Hardware items to probe.
                       Default: %s
                       Possible values: %s

`

	// Custom usage function
	flag.Usage = func() { fmt.Fprintf(os.Stderr, usage, defaultValues, possibleValues) }
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
