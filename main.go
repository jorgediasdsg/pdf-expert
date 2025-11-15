package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/ledongthuc/pdf"
)

func main() {
	http.HandleFunc("/analyze", analyzePDFHandler)

	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func analyzePDFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST /analyze", http.StatusMethodNotAllowed)
		return
	}

	// Read uploaded file
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save temporary file
	tmpFilePath := "./tmp_" + header.Filename
	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create temp file: %v", err), http.StatusInternalServerError)
		return
	}
	defer tmpFile.Close()

	_, err = tmpFile.ReadFrom(file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write temp file: %v", err), http.StatusInternalServerError)
		return
	}

	// Extract text from PDF
	content, err := extractPDFText(tmpFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to extract text from PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Count words
	wordCount := countWords(content)

	// Respond JSON
	response := map[string]interface{}{
		"file":       header.Filename,
		"word_count": wordCount,
		"status":     "completed",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = os.Remove(tmpFilePath)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, fmt.Sprintf("Failed to write response: %v", err), http.StatusInternalServerError)
		return
	}
}

func extractPDFText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	reader, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	if _, err = io.Copy(&buf, reader); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func countWords(text string) int {
	count := 0
	inWord := false

	for _, r := range text {
		if r == ' ' || r == '\n' || r == '\t' {
			if inWord {
				count++
				inWord = false
			}
		} else {
			inWord = true
		}
	}

	if inWord {
		count++
	}

	return count
}
