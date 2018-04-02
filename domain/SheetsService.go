package domain

type SheetsService interface {
	GetAllData() []Table
	GetAllDataForTable(tableName string) (Table, error)
	InsertRowIntoTable(tableName string, row Row) error
}
