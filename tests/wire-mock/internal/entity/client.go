package entity

import (
	"errors"

	"github.com/google/uuid"
)

type Client struct {
	ID string
	Name string
	Email string
	Points int
}

type ClientRepositoryInterface interface {
	Save(client *Client) error
}

func NewClient (name, email string) (*Client, error) {
	client := &Client{
		ID:     uuid.New().String(),
		Name:   name,
		Email:  email,
		Points: 0,
	}

	if err := client.Validate(); err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("client name is required")
	}
	if c.Email == "" {
		return errors.New("client email is required")
	}

	return nil
}

func (c *Client) AddPoints(points int) error {
	if points < 0 {
		return errors.New("points cannot be negative")
	}

	c.Points += points
	// if points%2 == 0 {
	// 	c.Points += points
	// }
	return nil
}