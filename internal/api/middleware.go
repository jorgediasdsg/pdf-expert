package api

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqID := uuid.New().String()
		start := time.Now()

		// Add request ID to context
		r = r.WithContext(withRequestID(r.Context(), reqID))

		// Recover panic
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("[PANIC] req=%s error=%v", reqID, rec)
				writeError(w, http.StatusInternalServerError, "internal server error", reqID)
			}
		}()

		// Log the request
		log.Printf("[REQUEST] %s %s req=%s", r.Method, r.URL.Path, reqID)

		// Continue
		next.ServeHTTP(w, r)

		// Log duration
		log.Printf("[DONE] req=%s duration=%s", reqID, time.Since(start))
	})
}
