package domain

// TransactionRepository defines the interface for reading and writing transactions
type TransactionRepository interface {
	ReadTransactions(filepath string) ([]Transaction, error)
	WriteReport(report []ReportEntry, filepath string) error
}
