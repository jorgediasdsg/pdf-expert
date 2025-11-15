package pdfanalyzer

import (
	"bytes"
	"io"

	"github.com/ledongthuc/pdf"
)

// PDFAnalyzer provides methods to analyze PDF files.
type PDFAnalyzer struct{}

// Constructor
func NewPDFAnalyzer() *PDFAnalyzer {
	return &PDFAnalyzer{}
}

// Analyze reads the PDF file at the given path and returns its text content and word count.
func (a *PDFAnalyzer) Analyze(filePath string) (string, int, error) {
	file, content, err := pdf.Open(filePath)
	if err != nil {
		return "", 0, err
	}
	defer file.Close()

	var buf bytes.Buffer
	reader, err := content.GetPlainText()
	if err != nil {
		return "", 0, err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		return "", 0, err
	}

	contentString := buf.String()

	wordCount := countWords(contentString)
	return contentString, wordCount, nil
}

// Simple, naive word counting
func countWords(text string) int {
	count := 0
	inWord := false

	for _, r := range text {
		if r == ' ' || r == '\n' || r == '\t' {
			if inWord {
				count++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		count++
	}

	return count
}
