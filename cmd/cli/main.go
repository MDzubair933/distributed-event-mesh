package main

import (
	"flag"
	"fmt"
)

func main() {
	cmd := flag.String("cmd", "status", "CLI command")
	flag.Parse()
	fmt.Printf("Event Mesh CLI executing: %s\n", *cmd)
}
