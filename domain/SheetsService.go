package domain

import "sheets-database/domain/tables"

type SheetsService interface {
	GetAllData(sheetId string) []tables.Table
	GetAllDataForTable(sheetId string, tableName string) (tables.Table, error)
	InsertRowIntoTable(tableName string, row tables.Row) error
}
