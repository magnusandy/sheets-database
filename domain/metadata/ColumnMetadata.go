package metadata

type ColumnMetadata struct {
	columnName string
	columnType ColumnType
	defaultVal string //todo check that the default fits the type //CAN BE nil
	nullable   bool
	//length?
	//unique?
}

func CreateColumnMetadata(
	columnName string,
	columnType ColumnType,
	defaultVal string,
	nullable bool) ColumnMetadata {
	return ColumnMetadata{columnName, columnType, defaultVal, nullable}
}

func (c ColumnMetadata) GetColumnName() string {
	return c.columnName
}

func (c ColumnMetadata) GetColumnType() ColumnType {
	return c.columnType
}

func (c ColumnMetadata) GetDefault() string {
	return c.defaultVal
}

func (c ColumnMetadata) SsNullable() bool {
	return c.nullable
}
