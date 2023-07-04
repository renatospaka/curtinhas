package api

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/renatospaka/http-mock/implementation2/httpClient"
)

var (
	// custom client
	client = &httpClient.HTTPClientMock{}

	// api
	api = NewAPI(client, "", 0)
)

func TestApiV1(t *testing.T) {
	// test table
	tt := []struct {
		Body       string
		StatusCode int
		Result     *APIPost
		Error      error
	}{
		{
			Body: `{"userId": 1,"id": 1,"title": "test title","body": "test body"}`,
			StatusCode: 200,
			Result: &APIPost{
				ID: 1,
				UserID: 1,
				Title: "test title",
				Body: "test body",
			},
			Error: nil,
		},
		{
			Body: `{"userId": 2,"id": 2,"title": "test title2","body": "test body2"}`,
			StatusCode: 200,
			Result: &APIPost{
				ID: 2,
				UserID: 2,
				Title: "test title2",
				Body: "test body2",
			},
			Error: nil,
		},
    {
      Body: ``,
      StatusCode: http.StatusNotFound,
      Result: nil,
      Error: fmt.Errorf(http.StatusText(http.StatusNotFound)),
    },
    {
      Body: ``,
      StatusCode: http.StatusBadRequest,
      Result: nil,
      Error: fmt.Errorf(http.StatusText(http.StatusBadRequest)),
    },
	}

	for _, test := range tt {
		// adjust the DoFunc for each test case
		client.DoFunc = func(r *http.Request) (*http.Response, error) {
			return &http.Response{
				// create the custom body
				Body: io.NopCloser(strings.NewReader(test.Body)),
				// createthe custom status code
				StatusCode: test.StatusCode,
			}, nil
		}

		// execute the func
		p, err := api.FetchPostById(context.Background(), 0)

		if err != nil && err.Error() != test.Error.Error() {
			t.Fatalf("expected: %v, got: %v", test.Error, err)
		}

		if !reflect.DeepEqual(p, test.Result) {
			t.Fatalf("expected: %v, got: %v", test.Result, p)
		}
	}
}