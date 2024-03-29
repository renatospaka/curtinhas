package data

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

var (
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("resource not found")
)

func PresentError(r *http.Request, err error) (*http.Request, interface{}) {
	switch err {
	case ErrUnauthorized:
		render.Status(r, 401)
	case ErrForbidden:
		render.Status(r, 403)
	case ErrNotFound:
		render.Status(r, 404)
	default:
		render.Status(r, 500)
	}
	return r, map[string]string{"error": err.Error()}
}
