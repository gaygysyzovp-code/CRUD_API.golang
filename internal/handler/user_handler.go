package handler

import (
	"CRUD_API/internal/models"
	"CRUD_API/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(router *gin.Engine, service service.UserService) {
	handler := &UserHandler{service: service}

	api := router.Group("/api")
	{
		api.POST("/users", handler.CreateUser)
		api.GET("/users", handler.GetUsers)
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUser

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
