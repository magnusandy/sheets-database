package domain

type SheetsService interface {
	GetAllData(sheetId string, tableName string) []Table
}
