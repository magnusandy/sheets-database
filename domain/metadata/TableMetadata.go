package metadata

type TableMetadata struct {
	TableName string
	Columns []ColumnMetadata
}

type ColumnMetadata struct {
	ColumnName string
	Type ColumnType
	Default string //todo check that the default fits the type //CAN BE nil
	Nullable bool
	//length?
}