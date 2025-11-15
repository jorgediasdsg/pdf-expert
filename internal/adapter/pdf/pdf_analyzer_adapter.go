package pdf

import (
	"github.com/jorgediasdsg/pdf-expert/internal/app/port"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

// PDFAnalyzerAdapter is the concrete implementation
// of the PDFAnalyzerPort, using the existing
// internal/pdfanalyzer component.
//
// This is part of the "infrastructure" or "adapter"
// layer: it knows about the PDF library and the
// concrete analyzer implementation.
type PDFAnalyzerAdapter struct {
	inner *pdfanalyzer.PDFAnalyzer
}

// NewPDFAnalyzerAdapter creates a new adapter that
// wraps the existing PDFAnalyzer.
func NewPDFAnalyzerAdapter(inner *pdfanalyzer.PDFAnalyzer) port.PDFAnalyzerPort {
	return &PDFAnalyzerAdapter{
		inner: inner,
	}
}

// AnalyzeFile calls the underlying PDFAnalyzer and
// maps its result into the domain.AnalysisResult type.
func (a *PDFAnalyzerAdapter) AnalyzeFile(path string) (domain.AnalysisResult, error) {
	res, err := a.inner.AnalyzeFile(path)
	if err != nil {
		return domain.AnalysisResult{}, err
	}

	return domain.AnalysisResult{
		Content:   res.Content,
		WordCount: res.WordCount,
	}, nil
}
