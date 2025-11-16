package usecase

import (
	"context"

	"github.com/jorgediasdsg/pdf-expert/internal/app/dto"
	"github.com/jorgediasdsg/pdf-expert/internal/app/port"
)

// AnalyzePDFUseCase orchestrates PDF analysis using a Port.
type AnalyzePDFUseCase struct {
	analyzer port.PDFAnalyzerPort
}

func NewAnalyzePDFUseCase(analyzer port.PDFAnalyzerPort) *AnalyzePDFUseCase {
	return &AnalyzePDFUseCase{
		analyzer: analyzer,
	}
}

func (uc *AnalyzePDFUseCase) Execute(ctx context.Context, input dto.AnalyzePDFInputDTO) (dto.AnalyzePDFOutputDTO, error) {

	// Call domain-capable port
	result, err := uc.analyzer.AnalyzeFile(input.FilePath)
	if err != nil {
		return dto.AnalyzePDFOutputDTO{}, err
	}

	// Map domain → DTO
	out := dto.AnalyzePDFOutputDTO{
		Content:   result.Content,
		WordCount: result.WordCount,
	}

	return out, nil
}
