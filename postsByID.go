package main

import (
	"encoding/json"
	"fmt"
	"log"

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

type posts_struct struct {
	Id string `json:"ID"`
	Title string `json:"Title"`
	Autor string `json:"Autor"`
	Nachricht string `json:"Nachricht"`
}

func postsByID(c *gin.Context) {
	id := c.Param("id")

	conn := connectToDB()

	rows, err := conn.Query("SELECT * FROM blog1 WHERE id=" + id)

	if err != nil {
		log.Fatal(err)
	}

	var title string
	var autor string
	var nachricht string

	for rows.Next() {
		rows.Scan(&id, &title, &autor, &nachricht)
		fmt.Println(autor)
	}

	c.JSON(http.StatusOK, gin.H{
		"ID": id,
		"Title": title,
		"Autor": autor,
		"Nachricht": nachricht,

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
	id := c.Param("id")

	var post posts_struct
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&post); err != nil {
		log.Fatal(err)
	}

	conn := connectToDB()

	rows, err := conn.Query("UPDATE blog1 SET title = '"+ post.Title +"', autor = '"+ post.Autor +"', nachricht = '"+ post.Nachricht +"' WHERE id="+ id)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"Status": "Updated successful",
	})

	rows.Close()
	conn.Close()
}

func delete_postsByID(c *gin.Context) {
	id := c.Param("id")

	conn := connectToDB()

	rows, err := conn.Query("DELETE FROM blog1 WHERE id=" + id)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Deleted",
	})

	rows.Close()
	conn.Close()
}

func connectToDB() *sql.DB {
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
		return conn
	}

	return nil
}
