package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Load HTML files from templates folder
    r.LoadHTMLGlob("/app/templates/*")

    r.Static("/assets", "/app/assets")


    // Serve the base HTML
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "base.html", nil)
    })

    // Serve dynamic content for HTMX
    r.GET("/content", func(c *gin.Context) {
        c.HTML(http.StatusOK, "content.html", gin.H{
            "Message": "Hello from HTMX-powered dynamic content!",
        })
    })

    // Start the server on port 8000
    r.Run(":8000")
}
