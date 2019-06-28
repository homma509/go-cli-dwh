package main

import (
	"fmt"

	"dwh/opts"
)

func main() {
	options := opts.Parse()
	fmt.Println(options.Host)
	fmt.Println(options.Port)
}