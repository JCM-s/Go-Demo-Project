package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	handleRequests()
}

func homePage(c *gin.Context){
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