package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
)

type Handler struct {
	usecase *usecase.AnalyzePDFUseCase
}

func NewHandler(uc *usecase.AnalyzePDFUseCase) *Handler {
	return &Handler{
		usecase: uc,
	}
}

// AnalyzePDF is a thin HTTP adapter. It delegates all logic to the use case.
func (h *Handler) AnalyzePDF(c *gin.Context) {
	cfg := config.Load()

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		writeError(c, 400, fmt.Sprintf("Failed to read file: %v", err))
		return
	}

	tmpPath := fmt.Sprintf("%s/%s", cfg.TempFolder, uploadedFile.Filename)

	if err := c.SaveUploadedFile(uploadedFile, tmpPath); err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to save temp file: %v", err))
		return
	}

	// Build use case input
	input := usecase.AnalyzePDFInput{
		FilePath: tmpPath,
	}

	// Execute the business logic
	output, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to analyze PDF: %v", err))
		_ = os.Remove(tmpPath)
		return
	}

	// Build successful response
	writeSuccess(c, gin.H{
		"file":       uploadedFile.Filename,
		"word_count": output.Analysis.WordCount,
		"status":     "completed",
	})

	_ = os.Remove(tmpPath)
}
