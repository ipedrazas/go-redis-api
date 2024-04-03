package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Initialize a Redis client.
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"), // Redis server address host:port
		DB:   0,                       // Default DB
	})

	// Initialize Gin.
	router := gin.Default()

	// Define a handler for the root path.
	router.GET("/api/v1/counter", func(c *gin.Context) {

		// Retrieve the updated visitor count.
		visitors, err := rdb.Get(ctx, "visitors").Result()
		if err != nil && !errors.Is(err, redis.Nil) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving visitor count"})
			return
		}
		if visitors == "" {
			visitors = "0"
		}
		// Send a response to the client.
		c.String(http.StatusOK, fmt.Sprintf("This page has been visited %s time(s)", visitors))
	})

	router.POST("/api/v1/counter", func(c *gin.Context) {
		// Increment the visitor count.
		err := rdb.Incr(ctx, "visitors").Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error incrementing visitor count"})
			return
		}
		visitors, err := rdb.Get(ctx, "visitors").Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving visitor count"})
			return
		}

		// Send a response to the client.
		c.String(http.StatusOK, fmt.Sprintf("This page has been visited %s time(s)", visitors))
	})

	// Start the Gin server on port 8080.
	router.Run(":8080")
}
