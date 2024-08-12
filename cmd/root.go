package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/numtide/nixos-facter/pkg/hwinfo"

	"github.com/numtide/nixos-facter/pkg/facter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile          string
	outputPath       string
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
		// convert the hardware features into probe features
		for _, str := range hardwareFeatures {
			probe, err := hwinfo.ProbeFeatureString(str)
			if err != nil {
				return fmt.Errorf("invalid hardware feature: %w", err)
			}
			scanner.Features = append(scanner.Features, probe)
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
	cobra.OnInitialize(initConfig)

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

	// We currently support all probe features at a high level as they share some generic information,
	// but we do not have mappings for all of their detail sections.
	// These will be added on a priority / need basis.

	defaultFeatures := []string{
		"memory", "pci", "net", "serial", "cpu", "bios", "monitor", "mouse", "scsi", "usb", "prom", "sbus", "sys",
		"sysfs", "udev", "block", "wlan",
	}

	// we strip default and int from the feature list
	allFeatures := hwinfo.ProbeFeatureStrings()
	slices.DeleteFunc(allFeatures, func(str string) bool {
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
			strings.Join(allFeatures, ","),
		),
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nixos-facter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".nixos-facter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
