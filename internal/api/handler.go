package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/dto"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
)

type Handler struct {
	usecase *usecase.AnalyzePDFUseCase
}

func NewHandler(uc *usecase.AnalyzePDFUseCase) *Handler {
	return &Handler{usecase: uc}
}

// AnalyzePDF godoc
// @Summary Analyze a PDF and count its words
// @Description Upload a PDF file and receive the word count
// @Tags analysis
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "PDF file"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 422 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /analyze [post]
func (h *Handler) AnalyzePDF(c *gin.Context) {
	cfg := config.Load()

	fileHeader, err := c.FormFile("file")
	if err != nil {
		writeError(c, 400, "file is required")
		return
	}

	tmpPath := fmt.Sprintf("%s/%s", cfg.TempFolder, fileHeader.Filename)
	if err := c.SaveUploadedFile(fileHeader, tmpPath); err != nil {
		writeError(c, 500, fmt.Sprintf("failed to save file: %v", err))
		return
	}

	input := dto.AnalyzePDFInputDTO{FilePath: tmpPath}

	output, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {

		switch err {
		case dto.ErrInvalidPath:
			writeError(c, 400, err.Error())
		case domain.ErrEmptyContent, domain.ErrInvalidWordCount:
			writeError(c, 422, err.Error())
		default:
			writeError(c, 500, err.Error())
		}

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
