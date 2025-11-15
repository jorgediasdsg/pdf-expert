package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/jorgediasdsg/pdf-expert/internal/api"
	"github.com/jorgediasdsg/pdf-expert/internal/pdfanalyzer"
)

func main() {
	// Create the analyzer component
	analyzer := pdfanalyzer.NewPDFAnalyzer()

	// Create HTTP handler grouping
	handler := api.NewHandler(analyzer)

	// Register routes
	http.HandleFunc("/analyze", handler.AnalyzePDF)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
