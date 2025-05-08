package main

import (
	"flag"
	"fmt"
)

const defaultPort = 3000

func main() {

	// pattern 1
	var portAddr = flag.Int("portA", defaultPort, "explanation")

	// pattern 2
	var port int
	flag.IntVar(&port, "port", defaultPort, "explanation")

	flag.Parse()

	fmt.Println("PortA:", *portAddr)
	fmt.Println("Port:", port)
}
