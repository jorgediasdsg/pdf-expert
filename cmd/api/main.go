package main

import (
	"fmt"
	"log"
	"net/http"

	httpapi "github.com/jorgediasdsg/pdf-expert/internal/api"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

func main() {
	analyzer := pdfanalyzer.NewPDFAnalyzer()
	handler := httpapi.NewHandler(analyzer)

	mux := http.NewServeMux()
	mux.HandleFunc("/analyze", handler.AnalyzePDF)

	wrapped := httpapi.Middleware(mux)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", wrapped))
}
