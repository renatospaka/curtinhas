package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/lib/pq"
)

func redirect(c *gin.Context) {
	shortURL := c.Param("shortURL")
	id := base62Decode(shortURL)

	// Try to get the URL from the cache
	if longURL, ok := cache.Load(shortURL); ok {
		c.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}
	
	var url URL
	err := db.QueryRow("SELECT long_url FROM urls WHERE id=$1", id).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	cache.Store(shortURL, url.LongURL)
	c.Redirect(http.StatusMovedPermanently, url.LongURL)
}
