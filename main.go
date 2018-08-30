package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Our API will consist of just two routes
	// /jokes - which will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - which will capture likes sent to a particular joke
	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	// Start and run the server
	router.Run(":3000")
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Jokes handler not implemented yet",
	})
}

// LikeJoke increments the likes of a particular joke item
func LikeJoke(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeJoke handler not implemented yet",
	})
}
