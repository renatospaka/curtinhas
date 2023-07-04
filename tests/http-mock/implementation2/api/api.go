package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/renatospaka/http-mock/implementation2/httpClient"
)

type API interface {
	// this function will do http call to external resource
	FetchPostById(context.Context, int) (*APIPost, error)
}

type APIPost struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func NewAPI(client httpClient.HTTPClient, baseURL string, timeout time.Duration) API {
	return &apiV1{
		c:       client,
		baseURL: baseURL,
		timeout: timeout,
	}
}

type apiV1 struct {
	// we need t put the http.Client here
	// so we can mock it inside the unit test
	c       httpClient.HTTPClient
	baseURL string
	timeout time.Duration
}

func (a *apiV1) FetchPostById(ctx context.Context, id int) (*APIPost, error) {
	u := fmt.Sprintf("%s/posts/%d", a.baseURL, id)

	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
  if err != nil {
    return nil, err
  }

	resp, err := a.c.Do(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

	// there is a parallel with apiMock response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(http.StatusText(resp.StatusCode))
	}

	var result *APIPost
	return result, json.NewDecoder(resp.Body).Decode(&result)
}
