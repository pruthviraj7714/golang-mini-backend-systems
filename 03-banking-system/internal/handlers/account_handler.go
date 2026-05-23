package handlers

import (
	"banking-system/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AccountHandler struct {
	AccountService *services.AccountService
}

func (h *AccountHandler) CreateAccount(c *gin.Context) {

	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	parsedUserId, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	accountId, err := h.AccountService.CreateAccount(parsedUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Account created successfully", "accountId": accountId})
}

func (h *AccountHandler) GetAccount(c *gin.Context) {

	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	parsedUserId, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	account, err := h.AccountService.GetAccount(parsedUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": account,
	})

}
