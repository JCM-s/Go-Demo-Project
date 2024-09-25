package main

//imports
import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// konstanten für die connection zur DB
const (
	host     = "e81b45af-1898-41d5-9aaf-077ca48ce2e7.postgresql.eu01.onstackit.cloud"
	port     = 5432
	user     = "jonas"
	password = "l7RkCe5OmB1l78T7hVrbNwu8p9h7VZ9iOxiaCFnybReu8FzguT3XT17bfnWJLK6S"
	dbname   = "blog"
)

// Post struktur
type Post struct {
	ID     int    `json:"id"`
	Titel  string `json:"title"`
	Autor  string `json:"autor"`
	Inhalt string `json:"inhalt"`
}

// GET /posts - Gibt eine liste aller posts aus
func posts(c *gin.Context) {

	//verbindung etablieren???
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

	//aus Tabelle blog1 in der DB werden die elemente id titel autor und nachricht genommen
	rows, err := conn.Query("SELECT id, title, autor, nachricht FROM blog1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error1": err.Error()})
		return
	}

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Titel, &post.Autor, &post.Inhalt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return

		}
		posts = append(posts, post)
	}
	c.JSON(http.StatusOK, posts)

	rows.Close()
	conn.Close()

}

// verbindung wird etabliert???
func post_posts(c *gin.Context) {
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

	var newPost Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//soll neuen Post anlegen

	err := conn.QueryRow("INSERT INTO blog1 (title, autor, nachricht) VALUES ($1, $2) RETURNING id",
		newPost.Titel, newPost.Autor, newPost.Inhalt).Scan(&newPost.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)

}
