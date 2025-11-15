package pdfanalyzer

import (
	"path/filepath"
	"testing"
)

func TestAnalyzeFile_SimplePDF(t *testing.T) {
	pdfPath := filepath.Join("testdata", "simple.pdf")

	analyzer := NewPDFAnalyzer()

	result, err := analyzer.AnalyzeFile(pdfPath)
	if err != nil {
		t.Fatalf("AnalyzeFile returned error: %v", err)
	}

	if result.WordCount == 0 {
		t.Errorf("expected > 0 words, got %d", result.WordCount)
	}

	if len(result.Content) == 0 {
		t.Errorf("expected content to be non-empty")
	}
}
