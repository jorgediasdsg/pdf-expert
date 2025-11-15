package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

// Handler groups dependencies for HTTP routes.
type Handler struct {
	Analyzer *pdfanalyzer.PDFAnalyzer
}

// Constructor
func NewHandler(a *pdfanalyzer.PDFAnalyzer) *Handler {
	return &Handler{Analyzer: a}
}

// AnalyzePDF handles the POST /analyze request
func (h *Handler) AnalyzePDF(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Use POST /analyze")
		return
	}

	// Read uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Failed to read file: %v", err))
		return
	}
	defer file.Close()

	// Write to temp file
	tmpPath := "./tmp_" + header.Filename
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create temp file: %v", err))
		return
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, file)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to write temp file: %v", err))
		return
	}

	// Run PDF analysis
	_, wordCount, err := h.Analyzer.Analyze(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to analyze PDF: %v", err))
		_ = os.Remove(tmpPath)
		return
	}

	// Success response
	writeJSON(w, http.StatusOK, map[string]any{
		"file":       header.Filename,
		"word_count": wordCount,
		"status":     "completed",
	})

	// Clean up
	_ = os.Remove(tmpPath)
}
