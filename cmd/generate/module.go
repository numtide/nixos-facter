package generate

import (
	"encoding/json"
	"fmt"
	"github.com/numtide/nixos-facter/pkg/nix"
	"github.com/numtide/nixos-facter/pkg/scan"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var ReportPath string

// scanCmd represents the scan command
var moduleCmd = &cobra.Command{
	Use:   "nixos-module",
	Short: "Generate a nixos module",
	// todo add Long description
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		var report *scan.Report

		if ReportPath == "" {
			if report, err = scan.Run(); err != nil {
				return fmt.Errorf("failed to scan hardware: %w", err)
			}
		} else {
			if file, err := os.Open(ReportPath); err != nil {
				return fmt.Errorf("failed to open report: %w", err)
			} else if b, err := io.ReadAll(file); err != nil {
				return fmt.Errorf("failed to read report: %w", err)
			} else if err = json.Unmarshal(b, &report); err != nil {
				return fmt.Errorf("failed to unmarshal report: %w", err)
			}
		}

		var writer io.WriteCloser
		if OutputPath == "" {
			writer = os.Stdout
			// append newline once we've finished
			defer println()
		} else {
			writer, err = os.OpenFile(OutputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
			if err != nil {
				return fmt.Errorf("failed to open output path for writing: %w", err)
			}
			defer writer.Close()
		}

		gen := nix.ModuleGenerator{Report: report}
		return gen.Generate(writer)
	},
}

func init() {
	f := moduleCmd.Flags()
	f.StringVarP(&OutputPath, "output", "o", "", "Path to write the nixos module")
	f.StringVarP(&ReportPath, "report", "r", "", "Path to read a report json file")

	generateCmd.AddCommand(moduleCmd)
}
