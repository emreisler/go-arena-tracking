package main

import (
	"fmt"
	"github.com/emreisler/go-arena-tracking/constants"
	"github.com/emreisler/go-arena-tracking/handlers"
	"github.com/emreisler/go-arena-tracking/middleware"
	"github.com/emreisler/go-arena-tracking/tracking"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

const memoryLimit = 500 * 1024 * 1024 // 500MB

func monitorMemory() {
	for {
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		if memStats.Alloc > memoryLimit {
			fmt.Printf("Memory exceeded limit! Used: %d bytes\n", memStats.Alloc)
			os.Exit(1) // Force exit
		}

		time.Sleep(1 * time.Second) // Adjust frequency as needed
	}
}

func main() {

	go tracking.StartTracer()

	go monitorMemory()

	if constants.GcOff {
		closeGC()
	}

	r := gin.Default()

	// CORS Middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Allow React frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// --- PPROF ENDPOINTS ---
	r.GET("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))

	r.Use(middleware.ArenaMiddleware())

	r.GET("/jsontrace", tracking.TraceHandler)

	r.POST("/user", handlers.UserHandler)

	r.POST("/order", handlers.OrderHandler)

	r.GET("/mem", MemoryStatsHandler)

	r.Run(":8080")
}

func MemoryStatsHandler(c *gin.Context) {
	ar := middleware.GetArenaFromContext(c)
	data := tracking.GetAllocations(ar)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func closeGC() {
	debug.SetGCPercent(-1)
}
