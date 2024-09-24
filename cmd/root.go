package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/numtide/nixos-facter/pkg/hwinfo"

	"github.com/numtide/nixos-facter/pkg/facter"

	"github.com/spf13/cobra"
)

var (
	cfgFile          string
	outputPath       string
	logLevel         string
	hardwareFeatures []string

	scanner = facter.Scanner{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nixos-facter",
	Short: "Hardware report generator",
	// todo Long description
	// todo add Long description
	RunE: func(cmd *cobra.Command, args []string) error {
		// check the effective user id is 0 e.g. root
		if os.Geteuid() != 0 {
			cmd.SilenceUsage = true
			return fmt.Errorf("you must run this program as root")
		}

		// convert the hardware features into probe features
		for _, str := range hardwareFeatures {
			probe, err := hwinfo.ProbeFeatureString(str)
			if err != nil {
				return fmt.Errorf("invalid hardware feature: %w", err)
			}
			scanner.Features = append(scanner.Features, probe)
		}

		// set the log level
		if logLevel != "" {
			if logLevel == "debug" {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			} else if logLevel == "info" {
				slog.SetLogLoggerLevel(slog.LevelInfo)
			} else if logLevel == "warn" {
				slog.SetLogLoggerLevel(slog.LevelWarn)
			} else if logLevel == "error" {
				slog.SetLogLoggerLevel(slog.LevelError)
			} else {
				return fmt.Errorf("invalid log level: %s", logLevel)
			}
		}

		report, err := scanner.Scan()
		if err != nil {
			return err
		}

		bytes, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal report to json: %w", err)
		}

		// if a file path is provided write the report to it, otherwise output the report on stdout
		if outputPath == "" {
			if _, err = os.Stdout.Write(bytes); err != nil {
				return fmt.Errorf("failed to write report to stdout: %w", err)
			}
			fmt.Println()
		} else if err = os.WriteFile(outputPath, bytes, 0o644); err != nil {
			return fmt.Errorf("failed to write report to output path: %w", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.s
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nixos-facter.yaml)")

	// Cobra also supports local flags, which will only run when this action is called directly.
	f := rootCmd.Flags()
	f.StringVarP(&outputPath, "output", "o", "", "path to write the report")

	// Options for optional ephemeral system properties.
	f.BoolVarP(&scanner.Swap, "swap", "s", false, "capture swap entries")
	f.BoolVarP(&scanner.Ephemeral, "ephemeral", "e", false, "capture all ephemeral properties e.g. swap, filesystems and so on")
	f.StringVarP(&logLevel, "log-level", "l", "info", "log level")

	// We currently support all probe features at a high level as they share some generic information,
	// but we do not have mappings for all of their detail sections.
	// These will be added on a priority / need basis.

	defaultFeatures := []string{
		"memory", "pci", "net", "serial", "cpu", "bios", "monitor", "scsi", "usb", "prom", "sbus", "sys", "sysfs",
		"udev", "block", "wlan",
	}

	// we strip default and int from the feature list
	probeFeatures := hwinfo.ProbeFeatureStrings()
	slices.DeleteFunc(probeFeatures, func(str string) bool {
		switch str {
		case "default", "int":
			return true
		default:
			return false
		}
	})

	f.StringSliceVarP(
		&hardwareFeatures,
		"hardware",
		"f",
		defaultFeatures,
		fmt.Sprintf(
			"Hardware items to probe. Possible values are %s",
			strings.Join(probeFeatures, ","),
		),
	)
}
