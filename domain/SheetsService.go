package domain

type SheetsService interface {
	GetAllData() []Table
	GetAllDataForTable(sheetId string, tableName string) (Table, error)
	InsertRowIntoTable(tableName string, row Row) error
}
