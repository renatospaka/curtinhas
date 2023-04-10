package main

import (
	"context"
	"log"
	"time"
)

type Service func(ctx context.Context) (int64, error)

func main() {
	services := []Service{
		Service1,
		Service2,
		Service3,
	}

	var total int64 = 0
	for _, service := range services {
		v, err := service(context.Background())
		if err != nil {
			log.Fatalln("Service", err)
		}

		total += v
	}

	log.Printf("Total is: %d\n", total)
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
