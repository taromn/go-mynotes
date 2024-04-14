package main

import (
	"os"
	"fmt"
	"bufio"
)

func check(e error) {
	if e != nil {
		panic(e)
	}

}

func main() {
	fmt.Println("↓os.ReadFile")
	dat, err := os.ReadFile("test.txt")  // read entire file
	check(err)
	fmt.Print(string(dat))
	
	fmt.Println("↓f.Seek")	
	f, err := os.Open("test.txt")
	check(err)

	defer f.Close()

	_, err = f.Seek(6,0)  // next time read from 6
	check(err)
	
	bt := make([]byte,5)
	f.Read(bt)
	fmt.Print(string(bt))

	fmt.Println()

	_, err = f.Seek(0,0)  // next time read from 0
        check(err)

	bt2 := make([]byte,12)
	f.Read(bt2)
	fmt.Println(string(bt2))


	fmt.Println("↓bufio")
	
	_, err = f.Seek(0,0)
        check(err)

	rb := bufio.NewReader(f)
	outb, err := rb.Peek(10)
	check(err)
	fmt.Println(string(outb))

}
