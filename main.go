package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
)

const (
	host     = "e81b45af-1898-41d5-9aaf-077ca48ce2e7.postgresql.eu01.onstackit.cloud"
	port     = 5432
	user     = "jonas"
	password = "l7RkCe5OmB1l78T7hVrbNwu8p9h7VZ9iOxiaCFnybReu8FzguT3XT17bfnWJLK6S"
	dbname   = "blog"
)

func main() {
	handleRequests()
}

func homePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}

func handleRequests() {
	// config := cors.DefaultConfig()
  	// config.AllowOrigins = []string{"http://localhost:10000", "http://localhost:80"}

	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", homePage)
	router.GET("/posts", posts)
	router.POST("/posts", post_posts)
	router.GET("/posts/:id", postsByID)
	router.POST("/posts/:id", post_postsByID)
	router.DELETE("/posts/:id", delete_postsByID)

	router.Run("localhost:10000")
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "null")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}
