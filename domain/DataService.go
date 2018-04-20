package domain

import "sheets-database/domain/metadata"
import (
	"sheets-database/domain/tables"
	"errors"
	"log"
)

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

func (d DataService) InsertData(sheetId string, unverifiedData tables.FullTable) error {
	//todo verify against metadata
	log.Print(unverifiedData)
	meta, err := d.MetadataService.GetMetadata(unverifiedData.GetTableName())
	if err != nil {
		return errors.New("Table was not found or metadata does not exist, double check the table name")
	}

	validationError := d.validateAgainstMetadata(meta, unverifiedData)
	LogWithMessageIfPresent("VALIDATION ERROR", validationError)
	return validationError
}

func (d DataService) validateAgainstMetadata(tableMetadata metadata.TableMetadata, table tables.FullTable) error {
	if tableMetadata.GetTableName() != table.GetTableName() {
		return errors.New("Table name: "+ table.GetTableName() + " does not match metadata: "+ tableMetadata.GetTableName() )
	}

	for _, rowMap := range table.GetRows() {
		rowError := d.validateRowAgainstMetadata(tableMetadata, rowMap)
		if rowError != nil {
			return rowError
		}
	}

	return nil
}

func (d DataService) validateRowAgainstMetadata(tableMetadata metadata.TableMetadata, rowMap map[string]string) error {
	metaMap := tableMetadata.GetColumnsAsMap()
	//todo all non-null columns must exist and if they don't they need a default
	//todo types need to work

	//all columns in the table need to exist in the meta
	for columnName, _ := range rowMap {
		meta := metaMap[columnName]
		if meta == nil {
			return errors.New("column name: "+ columnName + " does not exist in the meta.")
		}
	}
	return nil
}

func (d DataService) toFullTable(table tables.Table) tables.FullTable {
	meta, err := d.MetadataService.GetMetadata(table.GetTableName())
	LogIfPresent(err)//todo handle error better
	return table.ToFullTable(meta)
}
