package metadata

type TableMetadata struct {
	sheetId string
	tableName string
	columns   []ColumnMetadata
}

func CreateTableMetadata(sheetId string, tableName string, columns []ColumnMetadata) TableMetadata {
	return TableMetadata{sheetId,tableName, columns}
}

func (t TableMetadata) GetSheetId() string {
	return t.sheetId
}

func (t TableMetadata) GetTableName() string {
	return t.tableName
}

func (t TableMetadata) GetColumns() []ColumnMetadata {
	return t.columns
}

func (t TableMetadata) GetColumnsAsMap() map[string]*ColumnMetadata {
	returnMap := map[string]*ColumnMetadata{}
	for _, meta := range t.GetColumns() {
		returnMap[meta.GetColumnName()] = &meta
	}
	return returnMap
}
