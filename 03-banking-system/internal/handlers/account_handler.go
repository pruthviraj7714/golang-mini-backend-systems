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

func (h *AccountHandler) WithdrawMoney(c *gin.Context) {

	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "UserId not found",
		})
		return
	}

	var req struct {
		amount int64
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	parsedUUID, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	msg, err := h.AccountService.WithdrawMoney(parsedUUID, req.amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func (h *AccountHandler) DepositMoney(c *gin.Context) {
	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "UserId not found",
		})
		return
	}

	var req struct {
		amount int64
	}

	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	parsedUUID, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	msg, err := h.AccountService.DepositMoney(parsedUUID, req.amount)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func (h *AccountHandler) TransferMoney(c *gin.Context) {
	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "UserId not found",
		})
		return
	}
	var reqBody struct {
		FromAccountNumber string `json:"fromAccountNumber"`
		ToAccountNumber   string `json:"ToAccountNumber"`
		Amount            int64  `json:"amount"`
	}

	if err := c.ShouldBindBodyWithJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	parsedUUID, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	msg, err := h.AccountService.TransferMoney(parsedUUID, reqBody.Amount, reqBody.FromAccountNumber, reqBody.ToAccountNumber)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})

}

func (h *AccountHandler) GetTransactions(c *gin.Context) {

	userId, exists := c.MustGet("userId").(string)

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "userId not found",
		})
		return
	}

	parsedUserId, err := uuid.Parse(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, err := h.AccountService.GetTransactions(parsedUserId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": data,
	})

}
