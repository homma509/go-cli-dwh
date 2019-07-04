package models

type Columns []Column

func (cols *Columns) AddColumn(columnName string) {
	// func (cols *Columns) AddColumn(columnName string) *Column {
	// col := *tab.FindColumn(columnName)
	// if &col != nil {
	// 	return &col
	// }

	// for _, col := range tab.Columns {
	// 	if col.ColumnName == columnName {
	// 		return &col
	// 	}
	// }

	col := Column{ColumnName: columnName}
	*cols = append(*cols, col)
	// return &col
}
