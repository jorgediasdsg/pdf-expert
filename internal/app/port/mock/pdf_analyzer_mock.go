package mock

import (
	"github.com/jorgediasdsg/pdf-expert/internal/app/port"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
)

// Ensure interface compliance
var _ port.PDFAnalyzerPort = (*MockPDFAnalyzer)(nil)

type MockPDFAnalyzer struct {
	Result domain.AnalysisResult
	Err    error
}

func (m *MockPDFAnalyzer) AnalyzeFile(path string) (domain.AnalysisResult, error) {
	if m.Err != nil {
		return domain.AnalysisResult{}, m.Err
	}
	return m.Result, nil
}
