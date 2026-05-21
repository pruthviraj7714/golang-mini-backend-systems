package handlers

import (
	"banking-system/internal/models"
	service "banking-system/internal/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	err := h.UserService.RegisterUser(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})

}

func (h *UserHandler) Login(c *gin.Context) {

	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := c.ShouldBindBodyWithJSON(&loginRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	accessToken, refreshToken, err := h.UserService.LoginUser(loginRequest.Email, loginRequest.Password)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid credentials",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "User logged in successfully",
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})

}
