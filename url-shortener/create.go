package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var url URL
	if err := c.ShouldBindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// validate the URL
	if !isValid(url.LongURL) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
	}

	// Insert the long URL into the database and get the generated ID
	var id int
	err := db.QueryRow("INSERT INTO urls(long_url) VALUES ($1) RETURNING id", url.LongURL).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}

	// Convert the ID to a short URL
	url.ShortURL = base62Encode(id)

	// Update the record with the short URL
	_, err = db.Exec("UPDATE urls SET short_url=$1 WHERE id=$2", url.ShortURL, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}

	c.JSON(http.StatusOK, url)
}