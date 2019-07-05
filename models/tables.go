package models

import (
	"errors"
	// "fmt"
	"log"
)

type Tables []Table

func (tabs *Tables) Get(tableName string) (Table, error) {
	for _, tab := range *tabs {
		if tab.TableName == tableName {
			return tab, nil
		}
	}
	return nil, errors.New("Not Found")
}

func (tabs *Tables) AddTable(tableName string) {
	tab, err := tabs.Get(tableName)
	if err == nil {
		return
	}

	tab = Table{TableName: tableName}
	*tabs = append(*tabs, tab)
	return
}

func NewTables() *Tables {
	cmd := `
select
	table_name
,	column_name
from
	information_schema.columns
order by
	table_name
,	ordinal_position
limit
	50
;
`
	rows, err := DB.Query(cmd)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	var tabs Tables
	for rows.Next() {
		var tableName, columnName string
		rows.Scan(&tableName, &columnName)
		tabs.AddTable(tableName)
		// tabs.AddTable(tableName).Columns.AddColumn(columnName)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &tabs
}
