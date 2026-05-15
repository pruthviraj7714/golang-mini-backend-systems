package handler

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	UrlServices *services.UrlServices
}

func (h *UrlHandler) CreateUrl(c *gin.Context) {

	userId := c.MustGet("userId").(int)

	var url string

	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	shortUrl, err := h.UrlServices.CreateUrl(userId, url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "url already exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "url created",
		"url":     shortUrl,
	})
}
