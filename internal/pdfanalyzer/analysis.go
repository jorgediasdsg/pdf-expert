package pdfanalyzer

// AnalysisResult represents the outcome of analyzing a PDF file.
type AnalysisResult struct {
	Content   string // raw extracted text (Phase 2: still basic)
	WordCount int    // naive word count
}
