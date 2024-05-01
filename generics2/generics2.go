package main

import "fmt"

type Num interface {
	int | int8 | int16 | float32 | float64
}

func Add[T Num](a T, b T) T {
	return a + b
}

func main() {
	result := Add(1.1, 2)
	fmt.Println(result)

}
