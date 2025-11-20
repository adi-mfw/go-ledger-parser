package application

import (
	"fmt"
	"go-ledger-parser/internal/domain"
)

// LedgerService implements the domain.Ledger interface
type LedgerService struct{}

// NewLedgerService creates a new instance of LedgerService
func NewLedgerService() *LedgerService {
	return &LedgerService{}
}

// ProcessTransactions processes a slice of transactions and returns aggregated report entries
// It also returns a slice of validation errors encountered during processing
func (ls *LedgerService) ProcessTransactions(transactions []domain.Transaction) ([]domain.ReportEntry, []error, error) {
	summaryMap := make(map[string]*domain.ReportEntry)
	var validationErrors []error

	for _, t := range transactions {
		// Ensure entry exists for this user
		if summaryMap[t.UserID] == nil {
			summaryMap[t.UserID] = &domain.ReportEntry{
				UserID: t.UserID,
			}
		}

		entry := summaryMap[t.UserID]

		// Validate transaction
		if err := t.IsValid(); err != nil {
			validationErrors = append(validationErrors, err)
			entry.InvalidCount++
			continue
		}

		// Process valid transaction
		if t.Type == "CREDIT" {
			entry.TotalCredit += t.Amount
		} else if t.Type == "DEBIT" {
			entry.TotalDebit += t.Amount
		}
		entry.ValidCount++
	}

	// Convert map to slice
	report := make([]domain.ReportEntry, 0, len(summaryMap))
	for _, entry := range summaryMap {
		report = append(report, *entry)
	}

	return report, validationErrors, nil
}

// ProcessTransactionsUseCase handles the complete transaction processing workflow
type ProcessTransactionsUseCase struct {
	repository domain.TransactionRepository
	ledger     domain.Ledger
}

// NewProcessTransactionsUseCase creates a new instance of ProcessTransactionsUseCase
func NewProcessTransactionsUseCase(repo domain.TransactionRepository, ledger domain.Ledger) *ProcessTransactionsUseCase {
	return &ProcessTransactionsUseCase{
		repository: repo,
		ledger:     ledger,
	}
}

// Execute reads transactions, processes them, and writes the report
// Returns validation errors encountered during processing
func (uc *ProcessTransactionsUseCase) Execute(inputFile, outputFile string) ([]error, error) {
	// Read transactions
	transactions, err := uc.repository.ReadTransactions(inputFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read transactions: %w", err)
	}

	// Process transactions
	report, validationErrors, err := uc.ledger.ProcessTransactions(transactions)
	if err != nil {
		return nil, fmt.Errorf("failed to process transactions: %w", err)
	}

	// Write report
	if err := uc.repository.WriteReport(report, outputFile); err != nil {
		return nil, fmt.Errorf("failed to write report: %w", err)
	}

	return validationErrors, nil
}
