package models

import (
	// "fmt"
	"log"
)

type Column struct {
	ColumnName string
	UsedCount  int
}

type Table struct {
	TableName string
	Columns   []Column
}

var Tables []Table

func GetColumnInfo() {
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

	i := -1
	for rows.Next() {
		var tableName, columnName string
		rows.Scan(&tableName, &columnName)

		if i == -1 || Tables[i].TableName != tableName {
			table := Table{}
			table.TableName = tableName
			Tables = append(Tables, table)
			i += 1
		}

		column := Column{}
		column.ColumnName = columnName
		Tables[i].Columns = append(Tables[i].Columns, column)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(Tables)
}
