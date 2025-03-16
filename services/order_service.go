package services

import (
	"github.com/emreisler/go-arena-tracking/middleware"
	"github.com/emreisler/go-arena-tracking/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context, id, quantity int, price float64, items []string) *models.Order {
	ar := middleware.GetArenaFromContext(c)
	if ar == nil {
		return nil
	}
	return models.NewOrder(ar, id, quantity, price, items)
}
