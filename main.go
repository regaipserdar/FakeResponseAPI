package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Handle requests to the root URL ("/")
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome VulnApi Happy Hacking",
		})
	})

	// Handle requests to the "/fake-asp-response" endpoint
	r.GET("/fake-asp-response", func(c *gin.Context) {
		// Set response headers to mimic an ASP.NET server
		c.Header("Server", "Microsoft-IIS/10.0")
		c.Header("X-Powered-By", "ASP.NET")
		c.Header("X-AspNet-Version", "4.0.30319")
		c.Header("X-AspNetMvc-Version", "3.0")

		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome ASP.net Fake Response Header",
		})
	})

	// Run the Gin server on port 9090
	r.Run(":9090")
}
