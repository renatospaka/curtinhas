package api

import (
	"context"
	"fmt"
	"net/http"
)

type APIMock struct {}


// By doing that, it doesn’t increase the test coverage and it will skip the rest of the code inside the FetchPostByID real implementation.
func (a *APIMock) FetchPostById(ctx context.Context, id int) (*APIPost, error) {
	return nil, fmt.Errorf(http.StatusText(http.StatusNotFound))
}
