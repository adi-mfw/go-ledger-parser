package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"go-ledger-parser/internal/application"
	"go-ledger-parser/internal/infrastructure"
)

// runProcess executes the transaction processing workflow
func runProcess(cmd *cobra.Command, args []string) error {
	if verbose {
		fmt.Printf("Reading transactions from: %s\n", inputFile)
	}

	// Initialize dependencies
	repo := infrastructure.NewFileRepository()
	ledgerService := application.NewLedgerService()
	useCase := application.NewProcessTransactionsUseCase(repo, ledgerService)

	// Execute use case
	validationErrors, err := useCase.Execute(inputFile, outputFile)
	if err != nil {
		return fmt.Errorf("processing failed: %w", err)
	}

	// Print validation errors if any
	if len(validationErrors) > 0 {
		fmt.Println("\nValidation errors:")
		for _, err := range validationErrors {
			fmt.Printf("  - %v\n", err)
		}
	}

	if verbose {
		fmt.Printf("\nReport written to: %s\n", outputFile)
	} else {
		fmt.Printf("Successfully processed transactions. Report saved to %s\n", outputFile)
	}

	return nil
}
