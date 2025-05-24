// this is a test.
//   - olipa kerran kaljaa
//
// for a go doc
package main

import (
	initializers "github.com/happosade/csp-server/initializers"
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

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.Run()
}
