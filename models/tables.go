package models

import (
	"fmt"
	"log"
)

type Tables []Table

func (tabs *Tables) Contains(tableName string) bool {
	for _, tab := range *tabs {
		if tab.TableName == tableName {
			return true
		}
	}
	return false
}

func (tabs *Tables) Index(tableName string) (int, error) {
	for i, tab := range *tabs {
		if tab.TableName == tableName {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Could not find " + tableName)
}

func (tabs *Tables) Find(tableName string) {
	for _, tab := range *tabs {
		if tab.TableName == tableName {
			log.Println("見つかった")
			return
		}
	}
}

func (tabs *Tables) FindTable(tableName string) *Table {
	for _, tab := range *tabs {
		if tab.TableName == tableName {
			log.Println("見つかった!: "+tableName)
			return &tab
		}
	}
	return nil
}

func (tabs *Tables) AddTable(tableName string) *Table {
	i, err := tabs.Index(tableName)
	if err == nil {
		return &tabs[i]
	}

	tab := Table{TableName: tableName}
	*tabs = append(*tabs, tab)
	return &tab
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
		// tabs.AddTable(tableName).AddColumn(columnName)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return &tabs
}
