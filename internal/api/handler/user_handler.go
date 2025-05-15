package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/githubdmn/codego/ent"
)

type UserHandler struct {
	client *ent.Client
}

func NewUserHandler(client *ent.Client) *UserHandler {
	return &UserHandler{
		client: client,
	}
}

// @Summary Get all users
// @Description Get all users from the database
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} ent.User
// @Router /api/v1/users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.client.User.Query().All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
