package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/priyanshoon/go-url-shortener/handler"
	"github.com/priyanshoon/go-url-shortener/store"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// store.InitializeBackupStore()
	store.InitializeStore()

	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Sprintf("failed to start the web server - error : %v", err))
	}
}
