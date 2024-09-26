package main

//imports
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// Post struktur
type Post struct {
	ID     int    `json:"id"`
	Titel  string `json:"title"`
	Autor  string `json:"autor"`
	Inhalt string `json:"nachricht"`
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
	rows, err := conn.Query("SELECT id, title, autor, nachricht FROM blog1 ORDER BY id asc")
	if err != nil {
		log.Fatal(err)
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
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&newPost); err != nil {
		log.Fatal(err)
	}

	//soll neuen Post anlegen
	rows, err := conn.Query("INSERT INTO blog1 (title, autor, nachricht) VALUES ('"+ newPost.Titel +"', '"+ newPost.Autor +"', '"+ newPost.Inhalt +"') RETURNING id")

	for rows.Next() {
		rows.Scan(&newPost.ID)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPost)

	conn.Close()
	rows.Close()
}
