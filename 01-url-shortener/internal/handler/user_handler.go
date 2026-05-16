package handler

import (
	"net/http"
	"url-shortener/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func (h *UserHandler) RegisterUser(c *gin.Context) {

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and password are required",
		})
		return
	}

	err := h.UserService.RegisterUser(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	if request.Email == "" || request.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email and password are required",
		})
		return
	}

	token, err := h.UserService.LoginUser(request.Email, request.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
		"token":   token,
	})

}
