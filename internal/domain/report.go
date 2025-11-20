package domain

// ReportEntry represents aggregated transaction data for a user
type ReportEntry struct {
	UserID       string  `json:"user_id"`
	TotalCredit  float64 `json:"total_credit"`
	TotalDebit   float64 `json:"total_debit"`
	ValidCount   int     `json:"valid_count"`
	InvalidCount int     `json:"invalid_count"`
}
