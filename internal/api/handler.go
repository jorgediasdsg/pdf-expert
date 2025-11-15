package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

type Handler struct {
	Analyzer *pdfanalyzer.PDFAnalyzer
}

func NewHandler(a *pdfanalyzer.PDFAnalyzer) *Handler {
	return &Handler{Analyzer: a}
}

func (h *Handler) AnalyzePDF(w http.ResponseWriter, r *http.Request) {
	reqID := getRequestID(r.Context())

	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "Use POST /analyze", reqID)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		writeError(w, http.StatusBadRequest, fmt.Sprintf("Failed to read file: %v", err), reqID)
		return
	}
	defer file.Close()

	tmpPath := "./tmp_" + header.Filename
	tmpFile, err := os.Create(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create temp file: %v", err), reqID)
		return
	}
	defer tmpFile.Close()

	if _, err := io.Copy(tmpFile, file); err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to write temp file: %v", err), reqID)
		return
	}

	result, err := h.Analyzer.AnalyzeFile(tmpPath)
	if err != nil {
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to analyze PDF: %v", err), reqID)
		_ = os.Remove(tmpPath)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"file":       header.Filename,
		"word_count": result.WordCount,
		"status":     "completed",
	}, reqID)

	_ = os.Remove(tmpPath)
}
