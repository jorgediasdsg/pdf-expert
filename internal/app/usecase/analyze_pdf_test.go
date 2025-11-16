package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/jorgediasdsg/pdf-expert/internal/app/dto"
	"github.com/jorgediasdsg/pdf-expert/internal/app/port/mock"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
)

func TestAnalyzePDFUseCase_Success(t *testing.T) {
	mockPort := &mock.MockPDFAnalyzer{
		Result: domain.AnalysisResult{
			Content:   "hello world",
			WordCount: 2,
		},
		Err: nil,
	}

	uc := NewAnalyzePDFUseCase(mockPort)

	input := dto.AnalyzePDFInputDTO{
		FilePath: "/tmp/test.pdf",
	}

	output, err := uc.Execute(context.Background(), input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if output.WordCount != 2 {
		t.Errorf("expected 2 words, got %d", output.WordCount)
	}
}

func TestAnalyzePDFUseCase_InvalidDTO(t *testing.T) {
	mockPort := &mock.MockPDFAnalyzer{}
	uc := NewAnalyzePDFUseCase(mockPort)

	input := dto.AnalyzePDFInputDTO{
		FilePath: "",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected validation error, got nil")
	}
}

func TestAnalyzePDFUseCase_PortError(t *testing.T) {
	mockPort := &mock.MockPDFAnalyzer{
		Err: errors.New("port failure"),
	}
	uc := NewAnalyzePDFUseCase(mockPort)

	input := dto.AnalyzePDFInputDTO{
		FilePath: "/tmp/test.pdf",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected port error, got nil")
	}
}

func TestAnalyzePDFUseCase_DomainError(t *testing.T) {
	mockPort := &mock.MockPDFAnalyzer{
		Result: domain.AnalysisResult{
			Content:   "",
			WordCount: 5,
		},
	}

	uc := NewAnalyzePDFUseCase(mockPort)

	input := dto.AnalyzePDFInputDTO{
		FilePath: "/tmp/test.pdf",
	}

	_, err := uc.Execute(context.Background(), input)
	if err == nil {
		t.Fatalf("expected domain validation error, got nil")
	}
}
