package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

func NewRouter(analyzer *pdfanalyzer.PDFAnalyzer) *gin.Engine {
	router := gin.New()

	router.Use(GinMiddleware()) // logging + recovery + request id

	handler := NewHandler(analyzer)

	router.POST("/analyze", handler.AnalyzePDF)

	return router
}
