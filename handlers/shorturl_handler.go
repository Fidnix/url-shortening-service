package handlers

import (
	"net/http"
	"url-shortening-service/repository"
	"url-shortening-service/utils"

	"github.com/gin-gonic/gin"
)

type ShorturlHandler struct {
	repo *repository.ShorturlRepository
}

func NewShorturlHandler() *ShorturlHandler {
	return &ShorturlHandler{
		repo: repository.NewShorturlRepository(),
	}
}

func (h *ShorturlHandler) CreateShorturl(c *gin.Context) {
	var request struct {
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if !utils.VerifyUrl(request.Url) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url"})
		return
	}

	data, err := h.repo.InsertShortUrl(request.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create short url"})
		return
	}

	c.IndentedJSON(http.StatusCreated, data)
}

func (h *ShorturlHandler) GetOriginalUrl(c *gin.Context) {
	var shortUrl string = c.Param("shortUrl")

	data, err := h.repo.GetShortUrlByShortCode(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't get original url"})
		return
	}

	c.IndentedJSON(http.StatusAccepted, data)
}

func (h *ShorturlHandler) GetShorturlStats(c *gin.Context) {
	var shortUrl string = c.Param("shortUrl")

	data, err := h.repo.GetShortUrlByShortCode(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't get short url stats"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"id":          data.Id,
		"url":         data.Url,
		"shortUrl":    data.ShortCode,
		"createdAt":   data.CreatedAt,
		"updatedAt":   data.UpdatedAt,
		"accessCount": data.AccessCount,
	})
}

func (h *ShorturlHandler) UpdateShorturl(c *gin.Context) {
	var shortUrl string = c.Param("shortUrl")
	var request struct {
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	if !utils.VerifyUrl(request.Url) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid url"})
		return
	}

	data, err := h.repo.UpdateShortUrl(request.Url, shortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't update short url"})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

func (h *ShorturlHandler) DeleteShorturl(c *gin.Context) {
	var shortUrl string = c.Param("shortUrl")

	_, err := h.repo.DeleteShortUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Couldn't delete short url"})
		return
	}

	c.Status(http.StatusNoContent)
}
