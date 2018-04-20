package dto

import "sheets-database/domain/tables"

type AppendRow struct {
	Values [][]string `json:"values"`
}

func FromDomain(row tables.Row) AppendRow {
	var insertRow []string
	insertRow = append(insertRow, row.GetId())
	for i := 0; i < len(row.GetValues()); i++ {
		insertRow = append(insertRow, row.GetValues()[i])
	}
	var fullValues [][]string
	fullValues = append(fullValues, insertRow)
	return AppendRow{
		fullValues,
	}
}
