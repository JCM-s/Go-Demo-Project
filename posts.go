package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func posts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Endpoint": "posts",
	})
}