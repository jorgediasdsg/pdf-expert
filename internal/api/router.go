package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
)

// NewRouter wires routes and middleware.
func NewRouter(uc *usecase.AnalyzePDFUseCase) *gin.Engine {
	router := gin.New()

	router.Use(GinMiddleware())

	handler := NewHandler(uc)

	router.POST("/analyze", handler.AnalyzePDF)

	return router
}
