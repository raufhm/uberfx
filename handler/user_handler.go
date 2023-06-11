package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/raufhm/learning-uberfx/domain"
	"github.com/raufhm/learning-uberfx/service"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		// Handle the error
		// ...
		return
	}

	response := domain.User{
		UID:       user.UID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}

	c.JSON(http.StatusOK, response)
}
