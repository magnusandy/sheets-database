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

func (m MetadataListDto) ToDomain(sheetId string) map[string]metadata.TableMetadata {
	metaMap := map[string]metadata.TableMetadata{}
	for _, tableDto := range m.Tables {
		metaMap[tableDto.Name] = tableDto.ToDomain(sheetId)
	}
	return metaMap
}

type MetadataTableDto struct {
	Name    string
	Columns []MetadataColumnDto
}

func MetadataTableDtoFromDomain(tableMetadata metadata.TableMetadata) MetadataTableDto {
	columnDtos := []MetadataColumnDto{}
	for _, col := range tableMetadata.GetColumns() {
		columnDtos = append(columnDtos, MetadataColumnDtoFromDomain(col))
	}
	return MetadataTableDto{
		Name:    tableMetadata.GetTableName(),
		Columns: columnDtos,
	}
}

func (m MetadataTableDto) ToDomain(sheetId string) metadata.TableMetadata {
	domainColumns := []metadata.ColumnMetadata{}
	for _, colDto := range m.Columns {
		domainColumns = append(domainColumns, colDto.ToDomain())
	}

	return metadata.CreateTableMetadata(
		sheetId,
		m.Name,
		domainColumns,
	);
}

type MetadataColumnDto struct {
	Name     string
	Type     metadata.ColumnType
	Default  string
	Nullable bool
}

func MetadataColumnDtoFromDomain(columnMetadata metadata.ColumnMetadata) MetadataColumnDto {
	return MetadataColumnDto{
		Name:     columnMetadata.GetColumnName(),
		Type:     columnMetadata.GetColumnType(),
		Default:  columnMetadata.GetDefault(),
		Nullable: columnMetadata.GetNullable(),
	}
}

func (m MetadataColumnDto) ToDomain() metadata.ColumnMetadata {
	return metadata.CreateColumnMetadata(
		m.Name,
		m.Type,
		m.Default,
		m.Nullable,
	);
}
