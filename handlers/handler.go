package handlers

import (
	"github.com/emreisler/go-arena-tracking/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Request Payload
type CreateUserRequest struct {
	ID   int      `json:"id"`
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

// Request payload struct
type OrderRequest struct {
	ID       int      `json:"id"`
	Quantity int      `json:"quantity"`
	Price    float64  `json:"price"`
	Items    []string `json:"items"`
}

func UserHandler(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := services.CreateUser(c, req.ID, req.Name, req.Tags)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "name": user.Name, "tags": user.Tags})
}

func OrderHandler(c *gin.Context) {
	var req OrderRequest

	// Parse JSON request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Create order using Go arenas
	order := services.CreateOrder(c, req.ID, req.Quantity, req.Price, req.Items)

	// Respond with created order
	c.JSON(http.StatusOK, gin.H{
		"message": "Order created successfully",
		"order":   order,
	})
}
