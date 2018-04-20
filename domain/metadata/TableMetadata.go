package metadata

type TableMetadata struct {
	tableName string
	columns   []ColumnMetadata
}

func CreateTableMetadata(tableName string, columns []ColumnMetadata) TableMetadata {
	return TableMetadata{tableName, columns}
}

func (t TableMetadata) GetTableName() string {
	return t.tableName
}

func (t TableMetadata) GetColumns() []ColumnMetadata {
	return t.columns
}
