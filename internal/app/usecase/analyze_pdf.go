package usecase

import (
	"context"

	"github.com/jorgediasdsg/pdf-expert/internal/app/port"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
)

// AnalyzePDFInput represents the input data required
// to execute the AnalyzePDFUseCase.
type AnalyzePDFInput struct {
	FilePath string
}

// AnalyzePDFOutput represents the structured output
// of the AnalyzePDFUseCase.
type AnalyzePDFOutput struct {
	Analysis domain.AnalysisResult
}

// AnalyzePDFUseCase orchestrates the flow to analyze
// a PDF file using a PDFAnalyzerPort.
//
// It does not know anything about HTTP, Gin,
// filesystem details, or the PDF library itself.
type AnalyzePDFUseCase struct {
	analyzer port.PDFAnalyzerPort
}

func NewAnalyzePDFUseCase(analyzer port.PDFAnalyzerPort) *AnalyzePDFUseCase {
	return &AnalyzePDFUseCase{
		analyzer: analyzer,
	}
}

// Execute runs the use case: given a file path, it
// asks the analyzer port to analyze the file and
// returns a domain-level result.
func (uc *AnalyzePDFUseCase) Execute(ctx context.Context, in AnalyzePDFInput) (AnalyzePDFOutput, error) {
	analysis, err := uc.analyzer.AnalyzeFile(in.FilePath)
	if err != nil {
		return AnalyzePDFOutput{}, err
	}

	return AnalyzePDFOutput{
		Analysis: analysis,
	}, nil
}
