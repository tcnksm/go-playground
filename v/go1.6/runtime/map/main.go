package main

import "sync"

func main() {
	const workers = 2

	var wg sync.WaitGroup
	wg.Add(workers)

	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func() {
			m[0]++
			wg.Done()
		}()
	}
	wg.Wait()
}
