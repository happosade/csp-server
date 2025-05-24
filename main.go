// this is a test.
//   - olipa kerran kaljaa
//
// for a go doc
package main

import (
	"context"
	"encoding/json"
	"fmt"
	initializers "github.com/happosade/csp-server/initializers"
	"github.com/happosade/csp-server/models"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ping_gets",
		Help: "The total number of processed events",
	})
)

func init() {
	initializers.ConnectES8()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		go opsProcessed.Inc()
		c.JSON(200, gin.H{
			"message": "pong",
		})

	})

	r.POST("/cspreport", func(c *gin.Context) {
		var report models.Report
		if err := c.ShouldBindJSON(&report); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request payload"})
			return
		}

		// Prepare the report for Elasticsearch
		doc := map[string]interface{}{
			"document_uri":       report.Document_uri,
			"referrer":           report.Referrer,
			"blocked_uri":        report.Blocked_uri,
			"violated_directive": report.Violated_directive,
			"original_policy":    report.Original_policy,
		}

		// Push to Elasticsearch
		res, err := initializers.ES8.Index().
			Index("csp-reports").
			BodyJson(doc).
			Do(context.Background())

		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to save report"})
			fmt.Println("Error indexing document:", err)
			return
		}

		c.JSON(200, gin.H{
			"status":  "Report received",
			"message": "Thank you for helping us improve security!",
			"id":      res.Id,
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run()
}
