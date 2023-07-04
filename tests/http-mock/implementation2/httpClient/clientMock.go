package httpClient

import "net/http"

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPCientMock struct {
	// DOFunc will be executed whenever Do function is executed
	// so we'll be able to create a custom response
	DoFunc func(*http.Request) (*http.Response, error)
}

func (h *HTTPCientMock) Do(r *http.Request) (*http.Response, error) {
	return h.DoFunc(r)
}

