package domain

import "sheets-database/domain/metadata"
import "sheets-database/domain/tables"

type DataService struct {
	SheetsService   SheetsService
	MetadataService metadata.MetadataService
}

func (d DataService) GetListData(sheetId string, tableName string) tables.Table {
	table, err := d.SheetsService.GetAllDataForTable(sheetId, tableName)
	LogWithMessageIfPresent("error getting all data for table: "+tableName, err)
	return table
}

func (d DataService) GetFullData(sheetId string, tableName string) tables.FullTable {
	table, err := d.SheetsService.GetAllDataForTable(sheetId, tableName)
	LogWithMessageIfPresent("error getting all data for table: "+tableName, err)
	return d.toFullTable(table)
}

func (d DataService) toFullTable(table tables.Table) tables.FullTable {
	meta := d.MetadataService.GetMetadata(table.GetTableName())
	return table.ToFullTable(meta)
}

