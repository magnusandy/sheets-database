package dto

import "sheets-database/domain"

type AppendRow struct {
	Values [][]string `json:"values"`
}

func FromDomain(row domain.Row) AppendRow {
	var insertRow []string
	insertRow = append(insertRow, row.Id)
	for i := 0; i < len(row.Values); i++ {
		insertRow = append(insertRow, row.Values[i])
	}
	var fullValues [][]string
	fullValues = append(fullValues, insertRow)
	return AppendRow{
		fullValues,
	}
}
