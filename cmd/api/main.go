package main

import (
	"fmt"

	_ "github.com/jorgediasdsg/pdf-expert/cmd/api/docs"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"

	_ "github.com/jorgediasdsg/pdf-expert/cmd/api/docs"
	"github.com/jorgediasdsg/pdf-expert/internal/adapter/pdf"
	"github.com/jorgediasdsg/pdf-expert/internal/api"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
	"github.com/jorgediasdsg/pdf-expert/internal/log"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

func main() {
	cfg := config.Load()

	// Initialize global logger (dev or prod)
	log.Init(cfg.Env)

	// Infra analyzer (old implementation)
	infraAnalyzer := pdfanalyzer.NewPDFAnalyzer()

	// Adapter wrapping the infra analyzer as a Port implementation
	analyzerAdapter := pdf.NewPDFAnalyzerAdapter(infraAnalyzer)

	// Use case
	analyzeUseCase := usecase.NewAnalyzePDFUseCase(analyzerAdapter)

	// Router (Gin) receives ONLY the use case
	router := api.NewRouter(analyzeUseCase)

	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Logger.Info("server_started", "addr", addr)

	router.Run(addr)
}
