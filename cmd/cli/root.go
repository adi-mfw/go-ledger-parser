package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	inputFile  string
	outputFile string
	verbose    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ledger-parser",
	Short: "A CLI tool to parse and process ledger transactions",
	Long: `ledger-parser is a command-line utility that reads transaction data from JSON files,
validates them, and generates aggregated reports by user.

The tool processes transactions and creates summaries including:
- Total credits and debits per user
- Count of valid and invalid transactions`,
	Version: "1.0.0",
	RunE:    runProcess,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "Input JSON file containing transactions (required)")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output JSON file for the report (required)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")

	rootCmd.MarkFlagRequired("input")
	rootCmd.MarkFlagRequired("output")
}
