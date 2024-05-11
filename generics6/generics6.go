// Map w Generics

package main

import "fmt"

type CustomMap[T comparable, V int | string] map[T]V

func main() {

	m1 := make(CustomMap[int, string])
	m1[4] = "hello"

	fmt.Println(m1)
}
