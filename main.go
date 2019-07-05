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
	for _, tab := range *tabs {
		log.Println(tab.TableName)
		log.Println(tab.Columns)
	}
}
