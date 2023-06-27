package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
		_ "github.com/lib/pq"
)

type URL struct {
	LongURL  string
	ShortURL string
}

var cache = &sync.Map{}
var db *sql.DB

func main() {
	conn := "postgresql://postgres:postgres@localhost:3201/urls?sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	r := gin.Default()
	r.GET("/:shortURL", redirect)
	r.POST("/create", create)
	r.Run(":8000")
}
