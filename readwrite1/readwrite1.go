package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)

	}


// ↓defer f.Close()
	defer func(){
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()


	str := "written by golang"
	data := []byte(str)
	count, err := f.Write(data)

        if err != nil {
                fmt.Println(err)

        }

	fmt.Printf("write %d bytes\n", count)
}

// Ref: https://zenn.dev/hsaki/books/golang-io-package/viewer/file
