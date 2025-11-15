package pdfanalyzer

import (
	"bytes"
	"io"

	"github.com/ledongthuc/pdf"
)

// PDFAnalyzer processes PDF files and extracts text and metadata.
type PDFAnalyzer struct{}

// Constructor
func NewPDFAnalyzer() *PDFAnalyzer {
	return &PDFAnalyzer{}
}

// AnalyzeFile extracts text from the PDF at the given path and returns an AnalysisResult.
func (a *PDFAnalyzer) AnalyzeFile(filePath string) (AnalysisResult, error) {
	file, content, err := pdf.Open(filePath)
	if err != nil {
		return AnalysisResult{}, err
	}
	defer file.Close()

	var buf bytes.Buffer
	reader, err := content.GetPlainText()
	if err != nil {
		return AnalysisResult{}, err
	}

	if _, err := io.Copy(&buf, reader); err != nil {
		return AnalysisResult{}, err
	}

	text := buf.String()
	wordCount := countWords(text)

	return AnalysisResult{
		Content:   text,
		WordCount: wordCount,
	}, nil
}
