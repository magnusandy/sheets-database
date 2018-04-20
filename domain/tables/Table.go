package tables

import "sheets-database/domain/metadata"

type Table struct {
	tableName string
	rows []Row
}

func CreateTable(tableName string, rows []Row) Table {
	return Table{tableName:tableName, rows:rows}
}

func (t Table) GetTableName() string {
	return t.tableName
}

func (t Table) GetRows() []Row {
	return t.rows
}

func (t Table) ToFullTable(tableMetadata metadata.TableMetadata) FullTable {
	rowValueMaps := []map[string]string{}
	for _, row := range t.GetRows() {
		rowValueMaps = append(rowValueMaps, row.ToValueMap(tableMetadata.GetColumns()))
	}
	return CreateFullTable(t.GetTableName(), rowValueMaps)
}