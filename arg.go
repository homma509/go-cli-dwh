package main

import (
	"flag"
)

func parseArgs(string, string, []string) {
	host := flag.String("h", "127.0.0.1", "hostname (default=127.0.0.1)"
	flag.Parse()
	return *host, flag.Args(0), flag.Args()[1:]
}