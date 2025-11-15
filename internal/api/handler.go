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

// NewHandler creates a new Handler with its dependencies.
func NewHandler(a *pdfanalyzer.PDFAnalyzer) *Handler {
	return &Handler{Analyzer: a}
}

// AnalyzePDF handles POST /analyze
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

	// Create temp file
	tmpPath := "./tmp_" + header.Filename
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create temp file: %v", err))
		return
	}
	defer tmpFile.Close()

	// Copy uploaded file into temp file
	if _, err := io.Copy(tmpFile, file); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to write temp file: %v", err))
		return
	}

	// Perform analysis
	result, err := h.Analyzer.AnalyzeFile(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to analyze PDF: %v", err))
		_ = os.Remove(tmpPath)
		return
	}

	// Return response
	writeJSON(w, http.StatusOK, map[string]any{
		"file":       header.Filename,
		"word_count": result.WordCount,
		"status":     "completed",
	})

	// Clean up
	_ = os.Remove(tmpPath)
}
