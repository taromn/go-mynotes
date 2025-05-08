package main

import (
	"flag"
	"fmt"
)

const defaultPort = 3000

func main() {

	// pattern 1
	var portA = flag.Int("portA", defaultPort, "explanation")

	// pattern 2
	var portB int
	flag.IntVar(&portB, "portB", defaultPort, "explanation")

	flag.Parse()

	fmt.Println("PortA:", *portA)
	fmt.Println("PortB:", portB)

}
