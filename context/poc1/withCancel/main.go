package main

import (
	"context"
	"errors"
	"net/http"
	"time"
	"fmt"
)

func longRunningProcess(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return errors.New("Failed: Process took too much time")
}

func ghibliReq(ctx context.Context, requestStrig string) {
	req, _ := http.NewRequest(http.MethodGet, requestStrig, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Request failed", err)
	}

	select {
	case <- time.After(500 * time.Millisecond):
		fmt.Println("Response received, status code: ", res.StatusCode)
	case <- ctx.Done():
		fmt.Println("Took too long")
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

	go func() {
		err := longRunningProcess(ctx)
		if err != nil {
			cancel()
		}
	}()

	ghibliReq(ctx, "https://ghibliapi.herokuapp.com/films/58611129-2dbc-4a81-a72f-77ddfc1b1b49")
}