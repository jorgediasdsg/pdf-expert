package port

import "github.com/jorgediasdsg/pdf-expert/internal/domain"

// PDFAnalyzerPort defines the interface (port) that
// the application layer uses to analyze PDF files.
//
// Any concrete implementation (adapter) must satisfy
// this contract, but the app layer only depends on this
// interface, not on the library or storage details.
type PDFAnalyzerPort interface {
	AnalyzeFile(path string) (domain.AnalysisResult, error)
}
