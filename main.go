package main

import (
	// "fmt"
	"log"

	// "dwh/opts"
	"dwh/models"
)

func main() {
	// options := opts.Parse()
	// fmt.Println(options.Dbname)
	tabs := models.NewTables()
	log.Println(tabs)
}