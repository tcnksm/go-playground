package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	f4()
}

func f1() {
	for i := 0; i < 10; i++ {
		worker(i)
	}
}

func f2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			worker(i)
		}(i)
	}

	wg.Wait()
}

func f3() {
	var wg sync.WaitGroup
	errCh := make(chan error, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			errCh <- worker(i)
		}(i)
	}

	go func() {
		wg.Wait()
		close(errCh)
	}()

	for err := range errCh {
		if err != nil {
			log.Println(err)
		}
	}
}

func f4() {
	eg := errgroup.Group{}
	semaphore := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			semaphore <- struct{}{}
			defer func() {
				<-semaphore
			}()
			return worker(i)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Println(err)
	}
}

func f5() {
	eg, ctx := errgroup.WithContext(context.TODO())
	for i := 0; i < 10; i++ {
		i := i
		eg.Go(func() error {
			return workerContext(ctx, i)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Println(err)
	}

}

func worker(i int) error {

	if i%3 == 0 {
		// return fmt.Errorf("worker failed: %d", i)
	}

	time.Sleep(1 * time.Second)
	log.Println(i)
	return nil
}

func workerContext(ctx context.Context, i int) error {
	select {
	case <-ctx.Done():
		log.Println(i, ctx.Err())
		return ctx.Err()
	default:
	}

	if i%3 == 0 {
		return fmt.Errorf("worker failed: %d", i)
	}

	time.Sleep(1 * time.Second)
	log.Println(i)

	return nil
}
