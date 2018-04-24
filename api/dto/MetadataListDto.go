package dto

import "sheets-database/domain/metadata"

type MetadataListDto struct {
	Tables map[string]MetadataTableDto
}

func MetadataListDtoFromDomain(metaList map[string]metadata.TableMetadata) MetadataListDto {
	tableMap := map[string]MetadataTableDto{}
	for _, table := range metaList {
		tableMap[table.GetTableName()] = MetadataTableDtoFromDomain(table)
	}
	return MetadataListDto{
		Tables: tableMap,
	}
}

type MetadataTableDto struct {
	Name    string
	Columns []MetadataColumnDto
}

type MetadataColumnDto struct {
	Name     string
	Type     metadata.ColumnType
	Default  string
	Nullable bool
}

func MetadataTableDtoFromDomain(tableMetadata metadata.TableMetadata) MetadataTableDto {
	columnDtos := []MetadataColumnDto{}
	for _, col := range tableMetadata.GetColumns() {
		columnDtos = append(columnDtos, metadataColumnDtoFromDomain(col))
	}
	return MetadataTableDto{
		Name:    tableMetadata.GetTableName(),
		Columns: columnDtos,
	}
}

func metadataColumnDtoFromDomain(columnMetadata metadata.ColumnMetadata) MetadataColumnDto {
	return MetadataColumnDto{
		Name:     columnMetadata.GetColumnName(),
		Type:     columnMetadata.GetColumnType(),
		Default:  columnMetadata.GetDefault(),
		Nullable: columnMetadata.GetNullable(),
	}
}
