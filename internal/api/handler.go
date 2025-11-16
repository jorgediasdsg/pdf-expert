package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/dto"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
)

type Handler struct {
	usecase *usecase.AnalyzePDFUseCase
}

func NewHandler(uc *usecase.AnalyzePDFUseCase) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) AnalyzePDF(c *gin.Context) {
	cfg := config.Load()

	fileHeader, err := c.FormFile("file")
	if err != nil {
		writeError(c, 400, fmt.Sprintf("Failed to read file: %v", err))
		return
	}

	tmpPath := fmt.Sprintf("%s/%s", cfg.TempFolder, fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, tmpPath); err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to save file: %v", err))
		return
	}

	input := dto.AnalyzePDFInputDTO{
		FilePath: tmpPath,
	}

	output, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {
		writeError(c, 500, fmt.Sprintf("Failed to analyze PDF: %v", err))
		_ = os.Remove(tmpPath)
		return
	}

	writeSuccess(c, gin.H{
		"file":       fileHeader.Filename,
		"word_count": output.WordCount,
		"status":     "completed",
	})

	_ = os.Remove(tmpPath)
}
