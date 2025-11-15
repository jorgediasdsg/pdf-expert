package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

type Handler struct {
	Analyzer *pdfanalyzer.PDFAnalyzer
}

func NewHandler(a *pdfanalyzer.PDFAnalyzer) *Handler {
	return &Handler{Analyzer: a}
}

func (h *Handler) AnalyzePDF(c *gin.Context) {
	cfg := config.Load() // for TEMP_FOLDER

	file, err := c.FormFile("file")
	if err != nil {
		writeError(c, 400, fmt.Sprintf("Failed to read file: %v", err))
		return
	}

	tmpPath := fmt.Sprintf("%s/%s", cfg.TempFolder, file.Filename)

	if err := c.SaveUploadedFile(file, tmpPath); err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to save temp file: %v", err))
		return
	}

	result, err := h.Analyzer.AnalyzeFile(tmpPath)
	if err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to analyze PDF: %v", err))
		_ = os.Remove(tmpPath)
		return
	}

	writeSuccess(c, gin.H{
		"file":       file.Filename,
		"word_count": result.WordCount,
		"status":     "completed",
	})

	_ = os.Remove(tmpPath)
}
