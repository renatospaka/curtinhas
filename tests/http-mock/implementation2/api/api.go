package api

import "context"

type API interface {
	// this function will do http call to external resource
	FetchPostById(ctx context.Context, id int) (*APIPost, error)
}

type APIPost struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
