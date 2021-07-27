package main

import (
	"fmt"
	"sync"
)

func main() {
	count := 0
	const gos = 100
	var wg sync.WaitGroup
	wg.Add(gos)

	var mu sync.Mutex

	for i := 0; i < gos; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()

			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
