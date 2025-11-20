package domain

// Ledger defines the contract for processing transactions
type Ledger interface {
	ProcessTransactions(transactions []Transaction) ([]ReportEntry, []error, error)
}
