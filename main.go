package main

import (
	"url-shortening-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var shorturl_hanlder = handlers.NewShorturlHandler()

	router := gin.Default()
	router.POST("/shorten", shorturl_hanlder.CreateShorturl)
	router.GET("/shorten/:shortUrl", shorturl_hanlder.GetOriginalUrl)
	router.PUT("/shorten/:shortUrl", shorturl_hanlder.UpdateShorturl)
	router.DELETE("/shorten/:shortUrl", shorturl_hanlder.DeleteShorturl)
	router.GET("/shorten/:shortUrl/stats", shorturl_hanlder.GetShorturlStats)

	router.Run("localhost:8000")
}
