package handler

import (
	"fmt"
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UrlHandler struct {
	UrlService *services.UrlService
}

func (h *UrlHandler) CreateUrl(c *gin.Context) {

	userId := c.MustGet("userId")

	userIdStr, ok := userId.(string)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid User Id",
		})
		c.Abort()
		return
	}

	parsedUUID, err := uuid.Parse(userIdStr)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid User Id",
		})
		c.Abort()
		return
	}

	var request struct {
		Url string `json:"url"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	shortUrl, err := h.UrlService.CreateUrl(parsedUUID, request.Url)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "url created",
		"url":     shortUrl,
	})
}

func (h *UrlHandler) RedirectToShortUrl(c *gin.Context) {
	// userId := c.MustGet("userId")

	// userIdStr, ok := userId.(string)

	// if !ok {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid User Id",
	// 	})
	// 	c.Abort()
	// 	return
	// }

	// parsedUUID, err := uuid.Parse(userIdStr)

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid User Id",
	// 	})
	// 	c.Abort()
	// 	return
	// }

	shortId := c.Param("shortId")

	fmt.Println("param: " + shortId)

	resp, err := h.UrlService.RedirectToShortUrl(shortId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, resp)
}
