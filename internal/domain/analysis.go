package domain

// AnalysisResult represents the domain-level result
// of analyzing a PDF file. This type should not depend
// on any framework or external library.
type AnalysisResult struct {
	Content   string
	WordCount int
}
