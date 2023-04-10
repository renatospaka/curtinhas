package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

type Service func(ctx context.Context) (int64, error)

func main() {
	services := []Service{
		Service1,
		Service2,
		Service3,
	}

	var total int64 = 0
	
	start := time.Now()
	grp, ctx := errgroup.WithContext(context.Background())

	for _, service := range services {
		grp.Go(func() error {
			v, err := service(ctx)
			if err != nil {
				return err
			}
			_ = atomic.AddInt64(&total, v)

			return nil
		})
	}
	if err := grp.Wait(); err != nil {
		log.Fatalln("Service", err)
	}		

	duration := time.Since(start)
	log.Println("Total is: ", total, " Duration of", duration)
}

func Service1(ctx context.Context) (int64, error) {
	time.Sleep(100 * time.Millisecond)
	return 1, nil
}

func Service2(ctx context.Context) (int64, error) {
	time.Sleep(100 * time.Millisecond)
	return 2, nil
}

func Service3(ctx context.Context) (int64, error) {
	time.Sleep(200 * time.Millisecond)
	return 3, nil
}
