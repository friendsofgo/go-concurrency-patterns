package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := retry(
		ctx,
		time.Second,
		func() error {
			// some code here, for example call to an API or connect with DB
			return errors.New("error")
		})
	fmt.Println(err)
}

func retry(ctx context.Context, retryInterval time.Duration, fn func() error) error {
	for {
		err := fn()
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout (%v)", err)
		case <-time.After(retryInterval):
			fmt.Println("retrying...")
		}
	}
}
