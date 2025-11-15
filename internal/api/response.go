package api

import "github.com/gin-gonic/gin"

func writeSuccess(c *gin.Context, data interface{}) {
	reqID := c.GetString("request_id")
	c.JSON(200, gin.H{
		"success":    true,
		"data":       data,
		"request_id": reqID,
	})
}

func writeError(c *gin.Context, status int, message string) {
	reqID := c.GetString("request_id")
	c.JSON(status, gin.H{
		"success":    false,
		"error":      message,
		"request_id": reqID,
	})
}
