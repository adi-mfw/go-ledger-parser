package domain

import (
	"fmt"
	"time"
)

// Transaction represents a single transaction record
type Transaction struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Amount  float64   `json:"amount"`
	DateStr string    `json:"date"`
	UserID  string    `json:"user_id"`
	Date    time.Time `json:"-"` // Parsed date, not serialized
}

// IsValid validates a transaction and parses its date
func (t *Transaction) IsValid() error {
	if t.Amount <= 0 {
		return fmt.Errorf("transaction %s: %w, got %.2f", t.ID, ErrInvalidAmount, t.Amount)
	}

	parsedDate, err := time.Parse("2006-01-02", t.DateStr)
	if err != nil {
		return fmt.Errorf("transaction %s: %w '%s': %v", t.ID, ErrInvalidDate, t.DateStr, err)
	}

	t.Date = parsedDate
	return nil
}
