package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	os.Setenv("TAROMN", "DEV")
	os.Setenv("CONFIG", "1")
	os.Setenv("LANG", "Go")

	fmt.Println(os.Getenv("TAROMN")) // get one value

	list1 := os.Environ() // os.Environ() returns all env variables as a slice
	// e.g. [SHELL=/bin/bash COLORTERM=truecolor ...]
	fmt.Println(list1[:3])

	for _, item := range list1 {
		pair := strings.Split(item, "=")
		fmt.Println(pair)
	}

}

// reference: https://gobyexample.com/environment-variables
