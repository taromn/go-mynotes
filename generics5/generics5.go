// Struct w Generics

package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID   int
	Name string
	Data T
}

func main() {
	u := User[string]{
		ID:   0,
		Name: "test",
		Data: "teststring",
	}
	fmt.Println(u)
}

// ref: https://youtu.be/WpTKqnfp5dY?si=tDYVPJDbiVuJmsHf
