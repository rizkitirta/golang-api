package handler

import (
	"golang-api-gin/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var usersResponse []user.UserResponse
	for _, val := range users {
		userResponse := user.ConvertToResponse(val)
		usersResponse = append(usersResponse, userResponse)
	}

	c.JSON(200, gin.H{
		"data": usersResponse,
	})
}

func (h *userHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.FindById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := user.ConvertToResponse(result)
	c.JSON(200, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) Store(c *gin.Context) {
	var userReq user.UserRequest
	err := c.ShouldBindJSON(&userReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.Create(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := user.ConvertToResponse(result)
	c.JSON(200, gin.H{
		"data": response,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var userReq user.UserRequestUpdate
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.Update(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := user.ConvertToResponse(result)
	c.JSON(200, gin.H{
		"data": userResponse,
	})
}

func (h *userHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.userService.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted",
	})
}
