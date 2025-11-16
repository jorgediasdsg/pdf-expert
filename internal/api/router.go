package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
)

func NewRouter(uc *usecase.AnalyzePDFUseCase) *gin.Engine {
	router := gin.New()

	router.Use(GinMiddleware())
	router.Use(MetricsMiddleware())

	handler := NewHandler(uc)

	router.POST("/analyze", handler.AnalyzePDF)

	// Prometheus metrics endpoint
	router.GET("/metrics", MetricsHandler())

	return router
}
