package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postsByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Endpoint": "/posts/{id}",
	})

	/*
	post := [3]string{"Test1", "Test2", "Test3"}
	id, err := strconv.ParseInt(c.Param("id"), 16, 64)

	fmt.Println(err)

	c.JSON(http.StatusOK, gin.H{
		"message": post[id],
	})
	*/
}