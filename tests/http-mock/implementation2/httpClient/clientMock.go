package httpClient

import "net/http"

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type HTTPClientMock struct {
	// DOFunc will be executed whenever Do function is executed
	// so we'll be able to create a custom response
	DoFunc func(*http.Request) (*http.Response, error)
}

func (h *HTTPClientMock) Do(r *http.Request) (*http.Response, error) {
	return h.DoFunc(r)
}
