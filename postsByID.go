package main

import (
	"fmt"
	"log"
	"strconv"

	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "e81b45af-1898-41d5-9aaf-077ca48ce2e7.postgresql.eu01.onstackit.cloud"
	port     = 5432
	user     = "jonas"
	password = "l7RkCe5OmB1l78T7hVrbNwu8p9h7VZ9iOxiaCFnybReu8FzguT3XT17bfnWJLK6S"
	dbname   = "blog"
)

func postsByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected!")
	}

	rows, err := conn.Query("SELECT autor FROM blog1 WHERE id=" + c.Param("id"))

	if err != nil {
		log.Fatal(err)
	}

	var autor string

	for rows.Next() {
		rows.Scan(&autor)
		fmt.Println(autor)
	}

	c.JSON(http.StatusOK, gin.H{
		"ID":       id,
		"Autor":    autor,
		"Endpoint": "/posts/" + c.Param("id"),
	})

	rows.Close()
	conn.Close()

	/*
		post := [3]string{"Test1", "Test2", "Test3"}
		id, err := strconv.ParseInt(c.Param("id"), 16, 64)

		fmt.Println(err)

		c.JSON(http.StatusOK, gin.H{
			"message": post[id],
		})
	*/
}

func post_postsByID(c *gin.Context) {

}

func delete_postsByID(c *gin.Context) {
	
}
