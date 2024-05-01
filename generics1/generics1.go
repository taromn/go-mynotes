package main
import "fmt"

//func Add(a int, b int) int {
//	return a + b
//}

func Add[T int|float64](a T, b T) T {
      return a + b
}

func main() {
	result := Add(1.1, 2)
	fmt.Println(result)

}
