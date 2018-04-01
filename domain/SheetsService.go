package domain

import "sheets-database/api/dto"

type SheetsService interface {
	GetAllData(sheetId string, tableName string) dto.GetAllData
}
