package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			fmt.Printf("Start %d\n", i)
			time.Sleep(3 * time.Second)
			fmt.Printf("End %d\n", i)
		}()
	}

	wg.Wait()

	fmt.Println("finished all go routines")
}