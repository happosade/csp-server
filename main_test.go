package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/happosade/csp-server/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/cspreport", func(c *gin.Context) {
		var report models.Report
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// In a real scenario, we would push to Elasticsearch here
		// For testing purposes, we'll just return a success response

		c.JSON(200, gin.H{
			"status":  "Report received",
			"message": "Thank you for helping us improve security!",
			"id":      "mock-id",
		})
	})

	return r
}

func TestCSPReportEndpoint(t *testing.T) {
	r := setupRouter()

	t.Run("should return success for valid CSP report", func(t *testing.T) {
		// Create a valid CSP report JSON payload
		jsonPayload := []byte(`
{
  "document-uri": "https://example.com/page",
  "referrer": "https://example.com/",
  "blocked-uri": "https://attacker.com/malicious.js",
  "violated-directive": "script-src 'self' https://trusted.cdn.com",
  "original-policy": "default-src 'self'; script-src 'self' https://trusted.cdn.com; object-src 'none'; style-src 'self' 'unsafe-inline'; img-src * data:; font-src 'self' https://fonts.gstatic.com; frame-ancestors 'self';"
}`)

		req, _ := http.NewRequest("POST", "/cspreport", bytes.NewBuffer(jsonPayload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)
		assert.Contains(t, w.Body.String(), `"status":"Report received"`)
	})

	t.Run("should return error for invalid JSON", func(t *testing.T) {
		// Create an invalid JSON payload
		jsonPayload := []byte(`{invalid json}`)

		req, _ := http.NewRequest("POST", "/cspreport", bytes.NewBuffer(jsonPayload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Invalid request payload"`)
	})

	t.Run("should return error for missing fields", func(t *testing.T) {
		// Create a JSON payload with missing required fields
		jsonPayload := []byte(`{
			"document_uri": "https://example.com/page"
		}`)

		req, _ := http.NewRequest("POST", "/cspreport", bytes.NewBuffer(jsonPayload))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
		assert.Contains(t, w.Body.String(), `"error":"Invalid request payload"`)
	})
}