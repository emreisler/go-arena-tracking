package services

import (
	"github.com/emreisler/go-arena-tracking/middleware"
	"github.com/emreisler/go-arena-tracking/models"
	"github.com/gin-gonic/gin"
)

// Service function to create a user with Arena
func CreateUser(c *gin.Context, id int, name string, tags []string) *models.User {
	ar := middleware.GetArenaFromContext(c)
	if ar == nil {
		return nil
	}
	return models.NewUser(ar, id, name, tags)
}
