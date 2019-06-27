package main

import (
	"fmt"

	"dwh/args"
)

func main() {
	params := args.Parse()
	fmt.Println(params.Host)
}