package domain

import "sheets-database/domain/metadata"

type DataService struct {
	SheetsService   SheetsService
	MetadataService metadata.MetadataService
}

func (d DataService) GetListData(sheetId string, tableName string) Table {
	table, err := d.SheetsService.GetAllDataForTable(sheetId, tableName)
	LogWithMessageIfPresent("error getting all data for table: "+tableName, err)
	return table
}

func (d DataService) GetFullData(sheetId string, tableName string) FullTable {
	table, err := d.SheetsService.GetAllDataForTable(sheetId, tableName)
	LogWithMessageIfPresent("error getting all data for table: "+tableName, err)
	return d.toFullTable(table)
}

func (d DataService) toFullTable(table Table) FullTable {
	meta := d.MetadataService.GetMetadata(table.TableName)
	rowValueMaps := []map[string]string{}
	for _, row := range table.Rows {
		rowValueMaps = append(rowValueMaps, row.toValueMap(meta.Columns))
	}
	return FullTable{
		table.TableName,
		rowValueMaps,
	}
}

