package domain

import "errors"

var (
	ErrInvalidAmount = errors.New("amount must be greater than 0")
	ErrInvalidDate   = errors.New("invalid date format, expected YYYY-MM-DD")
)
