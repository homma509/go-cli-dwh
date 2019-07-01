package models

import (
	// "fmt"
	// "log"
)

type Table struct {
	TableName string
	// Columns   []Column
}

// func (tab *Table) AddColumn(columnName string) *Column {
// 	col := *tab.FindColumn(columnName)
// 	if &col != nil {
// 		return &col
// 	}

// 	col = Column{ColumnName: columnName}
// 	tab.Columns = append(tab.Columns, col)
// 	return &col
// }

// func (tab *Table) FindColumn(columnName string) *Column {
// 	for _, col := range tab.Columns {
// 		if col.ColumnName == columnName {
// 			return &col
// 		}
// 	}
// 	return nil
// }

// func (tab *Table) Cmd() string {
// 	cmd := ""
// 	for i, col := range tab.Columns {
// 		if i > 0 {
// 			cmd += ","
// 		}
// 		cmd += col.Cmd()
// 	}
// 	cmd = fmt.Sprintf("SELECT %s FROM %s", cmd, tab.TableName)
// 	return cmd
// }

// func (tab *Table) Exec() {
// 	rows, err := DB.Query(tab.Cmd())
// 	defer rows.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		// TODO
// 		rows.Scan()
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
