package domain

// AnalysisResult represents the pure domain entity.
type AnalysisResult struct {
	Content   string
	WordCount int
}

// Validate enforces domain invariants.
func (a AnalysisResult) Validate() error {
	if a.Content == "" {
		return ErrEmptyContent
	}
	if a.WordCount < 0 {
		return ErrInvalidWordCount
	}
	return nil
}
