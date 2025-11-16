package api

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jorgediasdsg/pdf-expert/internal/app/port/mock"
	"github.com/jorgediasdsg/pdf-expert/internal/app/usecase"
	"github.com/jorgediasdsg/pdf-expert/internal/domain"
)

func TestAnalyzePDFHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Fake use case (mocked)
	mockPort := &mock.MockPDFAnalyzer{
		Result: domain.AnalysisResult{
			Content:   "hello world",
			WordCount: 2,
		},
	}

	uc := usecase.NewAnalyzePDFUseCase(mockPort)
	handler := NewHandler(uc)

	// Create test router
	router := gin.New()
	router.POST("/analyze", handler.AnalyzePDF)

	// Create fake PDF file
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.pdf")
	part.Write([]byte("dummy pdf content"))
	writer.Close()

	req := httptest.NewRequest("POST", "/analyze", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if !bytes.Contains(w.Body.Bytes(), []byte(`"word_count":2`)) {
		t.Errorf("expected word_count=2 in response")
	}
}

func TestAnalyzePDFHandler_InvalidRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockPort := &mock.MockPDFAnalyzer{}
	uc := usecase.NewAnalyzePDFUseCase(mockPort)
	handler := NewHandler(uc)

	router := gin.New()
	router.POST("/analyze", handler.AnalyzePDF)

	req := httptest.NewRequest("POST", "/analyze", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fatalf("expected status 400, got %d", w.Code)
	}
}
