package main

import "fmt"

func process(v interface{}) (res string) {

	if s, ok := v.(string); ok {
		res = s + " world"
	}

	return

}

func main() {

	test := "hello"

	response := process(test)

	fmt.Println(response)
}

