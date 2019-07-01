package models

import (
	"fmt"
)

type Column struct {
	ColumnName string
	UsedCount  int
}

func (col *Column) Cmd() string {
	return fmt.Sprintf("SUM(CASE WHEN %s IS NULL THEN 0 ELSE 1 END) AS %s", col.ColumnName, col.ColumnName)
}