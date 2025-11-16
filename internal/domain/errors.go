package domain

import "errors"

// Domain-level errors, not tied to HTTP or infra.
// These errors describe violations of business invariants.

var (
	ErrEmptyContent     = errors.New("analysis content cannot be empty")
	ErrInvalidWordCount = errors.New("invalid word count")
)
