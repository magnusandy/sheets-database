package tables

import "sheets-database/domain/metadata"

type FullTable struct {
	tableName string
	rows []map[string]string
}

func CreateFullTable(tableName string, rows []map[string]string) FullTable {
	return FullTable{tableName, rows}
}

func (f FullTable) GetTableName() string {
	return f.tableName
}

func (f FullTable) GetRows() []map[string]string {
	return f.rows
}

//ordering is important, based off meta
func (f FullTable) ToTable(meta metadata.TableMetadata) Table {
	rows := []Row{}
	for _, rowMap := range f.GetRows() {
		rows = append(rows, f.rowMapToRow(rowMap, meta.GetColumns()))
	}
	return CreateTable(f.tableName, rows)
}

func (f FullTable) rowMapToRow(rowMap map[string]string, metaColumns []metadata.ColumnMetadata) Row {
	id := rowMap[metadata.ID_COLUMN]
	values := []string{}
	withoutId := metaColumns[1:]
	for _, metaC := range withoutId {
		values = append(values, rowMap[metaC.GetColumnName()])
	}
	return CreateRow(id, values)
}
