package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	requestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path"},
	)

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_error_total",
			Help: "Total number of error responses",
		},
		[]string{"method", "path", "status"},
	)

	latencyHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of request durations",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(requestCounter)
	prometheus.MustRegister(errorCounter)
	prometheus.MustRegister(latencyHistogram)
}

func MetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		method := c.Request.Method
		path := c.FullPath()
		status := c.Writer.Status()

		requestCounter.WithLabelValues(method, path).Inc()
		latencyHistogram.WithLabelValues(method, path).Observe(time.Since(start).Seconds())

		if status >= 400 {
			errorCounter.WithLabelValues(method, path, http.StatusText(status)).Inc()
		}
	}
}

func MetricsHandler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
