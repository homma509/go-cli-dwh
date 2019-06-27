package main

import (
	"fmt"
)

func main() {
	host, port, username, password, dbname := parseArgs()
	fmt.Println(host, port, username, password, dbname)
}