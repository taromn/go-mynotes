package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

const defaultPort = 3000

func main() {

	// pattern 3 pflag
	var portC = pflag.Int("portC", defaultPort, "explanation")

	pflag.Parse()

	fmt.Println("PortC:", *portC)

}
