package infrastructure

import (
	"encoding/json"
	"fmt"
	"go-ledger-parser/internal/domain"
	"io/ioutil"
)

// FileRepository implements domain.TransactionRepository for file-based I/O
type FileRepository struct{}

// NewFileRepository creates a new instance of FileRepository
func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

// ReadTransactions reads and unmarshals transaction data from a JSON file
func (fr *FileRepository) ReadTransactions(filepath string) ([]domain.Transaction, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filepath, err)
	}

	var transactions []domain.Transaction
	if err := json.Unmarshal(data, &transactions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %w", err)
	}

	return transactions, nil
}

// WriteReport marshals and writes the report to a JSON file
func (fr *FileRepository) WriteReport(report []domain.ReportEntry, filepath string) error {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal report data: %w", err)
	}

	if err := ioutil.WriteFile(filepath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filepath, err)
	}

	return nil
}
