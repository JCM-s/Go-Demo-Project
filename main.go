package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/", homePage)
	router.GET("/posts", posts)
	router.POST("/posts", post_posts)
	router.GET("/posts/:id", postsByID)
	router.POST("/posts/:id", post_postsByID)
	router.DELETE("/posts/:id", delete_postsByID)

	router.Run("localhost:10000")
}
