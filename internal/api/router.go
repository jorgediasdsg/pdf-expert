package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jorgediasdsg/pdf-expert/cmd/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
