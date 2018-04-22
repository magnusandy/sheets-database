package domain

import "sheets-database/domain/tables"

type SheetsService interface {
	GetAllData(sheetId string) []tables.Table
	GetAllDataForTable(sheetId string, tableName string) (tables.Table, error)
	InsertRowsIntoTable(sheetId string, table tables.Table) error
}
