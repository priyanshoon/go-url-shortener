package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/priyanshoon/go-url-shortener/shortener"
	"github.com/priyanshoon/go-url-shortener/store"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	// UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, "e0dba740-fc4b-4977-872c-d360239e6b1a")
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, "e0dba740-fc4b-4977-872c-d360239e6b1a")

	host := "http://localhost:3000/"
	c.HTML(200, "url_got.html", gin.H{
		"message":   "Short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")

	initialUrl, pass := store.RetrieveInitialUrl(shortUrl)

	if pass == false {
		c.HTML(404, "404.html", gin.H{})
	}

	c.Redirect(302, initialUrl)
}
