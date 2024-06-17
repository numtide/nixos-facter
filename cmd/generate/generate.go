package generate

import (
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate reports, nixos modules and so on.",
}

func init() {
	generateCmd.Flags().StringVarP(&OutputPath, "output", "o", "report.json", "Path to write the report")
}

func Register(parent *cobra.Command) {
	// add to parent
	parent.AddCommand(generateCmd)
}
