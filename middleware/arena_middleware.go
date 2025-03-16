package middleware

import (
	"arena"
	"github.com/gin-gonic/gin"
)

// Key for Gin Context
const ArenaKey = "arena_memory"

// Middleware to allocate and manage arena per request
func ArenaMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create an Arena for the request
		ar := arena.NewArena()
		c.Set(ArenaKey, ar)

		// Continue with the request
		c.Next()

		// Free Arena Memory after response
		ar.Free()
	}
}

// Get Arena from Context
func GetArenaFromContext(c *gin.Context) *arena.Arena {
	ar, exists := c.Get(ArenaKey)
	if !exists {
		return nil
	}
	return ar.(*arena.Arena)
}
