package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	// Normal context with cancel
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	fmt.Printf("error is: %v\n", ctx.Err())
	fmt.Printf("cause was: %v\n", context.Cause(ctx))

	// New context with cancel and cause
	ctxCause, cancelCause := context.WithCancelCause(context.Background())
	cancelCause(errors.New("something went wrong"))
	fmt.Printf("error is: %v\n", ctxCause.Err())
	fmt.Printf("cause was: %v\n", context.Cause(ctxCause))
}