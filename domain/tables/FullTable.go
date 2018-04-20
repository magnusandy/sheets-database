package tables

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
