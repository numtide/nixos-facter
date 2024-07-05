package generate

import (
	"encoding/json"
	"fmt"
	"github.com/numtide/nixos-facter/pkg/hwinfo"
	"github.com/spf13/cobra"
	"os"
)

var OutputPath string
var PrettyPrint bool

// scanCmd represents the scan command
var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Scan the system and produce a report",
	// todo add Long description
	RunE: func(cmd *cobra.Command, args []string) error {
		report, err := hwinfo.Scan()
		if err != nil {
			return err
		}

		var b []byte
		if PrettyPrint {
			b, err = json.MarshalIndent(report, "", "  ")
		} else {
			b, err = json.Marshal(report)
		}

		if err != nil {
			return fmt.Errorf("failed to marshal report to json: %w", err)
		}

		// if a file path is provided write the report to it, otherwise output the report on stdout
		if OutputPath == "" {
			if _, err = os.Stdout.Write(b); err != nil {
				return fmt.Errorf("failed to write report to stdout: %w", err)
			}
			fmt.Println()
		} else if err = os.WriteFile(OutputPath, b, 0644); err != nil {
			return fmt.Errorf("failed to write report to output path: %w", err)
		}

		return nil
	},
}

func init() {
	f := reportCmd.Flags()
	f.StringVarP(&OutputPath, "output", "o", "", "Path to write the report")
	f.BoolVarP(&PrettyPrint, "pretty-print", "p", false, "Pretty print json")

	generateCmd.AddCommand(reportCmd)
}
