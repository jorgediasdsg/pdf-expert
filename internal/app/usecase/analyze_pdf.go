package usecase

import (
	"context"

	"github.com/jorgediasdsg/pdf-expert/internal/app/dto"
	"github.com/jorgediasdsg/pdf-expert/internal/app/port"
)

type AnalyzePDFUseCase struct {
	analyzer port.PDFAnalyzerPort
}

func NewAnalyzePDFUseCase(analyzer port.PDFAnalyzerPort) *AnalyzePDFUseCase {
	return &AnalyzePDFUseCase{analyzer: analyzer}
}

// Execute applies validation at the DTO and domain levels.
func (uc *AnalyzePDFUseCase) Execute(ctx context.Context, input dto.AnalyzePDFInputDTO) (dto.AnalyzePDFOutputDTO, error) {

	// 1. DTO validation
	if err := input.Validate(); err != nil {
		return dto.AnalyzePDFOutputDTO{}, err
	}

	// 2. Port call → returns domain object
	domainResult, err := uc.analyzer.AnalyzeFile(input.FilePath)
	if err != nil {
		return dto.AnalyzePDFOutputDTO{}, err
	}

	// 3. Domain validation
	if err := domainResult.Validate(); err != nil {
		return dto.AnalyzePDFOutputDTO{}, err
	}

	// 4. Map domain → DTO
	out := dto.AnalyzePDFOutputDTO{
		Content:   domainResult.Content,
		WordCount: domainResult.WordCount,
	}

	return out, nil
}
