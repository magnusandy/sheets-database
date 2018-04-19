package domain

type SheetsService interface {
	GetAllData(sheetId string) []Table
	GetAllDataForTable(sheetId string, tableName string) (Table, error)
	InsertRowIntoTable(tableName string, row Row) error
}
