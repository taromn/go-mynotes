package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	f("direct")

	go f("goroutine 1st")

	go f("goroutine 2nd")

	time.Sleep(3*time.Second)

	fmt.Println("done")
}

// ref https://gobyexample.com/goroutines
