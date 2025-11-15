package main

import (
	"fmt"

	"github.com/jorgediasdsg/pdf-expert/internal/api"
	"github.com/jorgediasdsg/pdf-expert/internal/config"
	"github.com/jorgediasdsg/pdf-expert/internal/log"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

func main() {
	cfg := config.Load()

	log.Init(cfg.Env)

	analyzer := pdfanalyzer.NewPDFAnalyzer()
	router := api.NewRouter(analyzer)

	addr := fmt.Sprintf(":%s", cfg.HTTPPort)
	log.Logger.Info("server_started", "addr", addr)

	router.Run(addr)
}
