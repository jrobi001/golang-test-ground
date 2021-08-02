package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	go func() {
		var wg sync.WaitGroup
		routines := 10
		wg.Add(routines)

		for i := 0; i < routines; i++ {
			go func(i int) {
				for j := 0; j < 10; j++ {
					ch <- j + (i * 10)
				}
				wg.Done()
			}(i)
		}

		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("This is the end.")

}
