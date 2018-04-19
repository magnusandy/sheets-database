package domain

import "sheets-database/domain/metadata"

type Table struct {
	TableName string
	Rows []Row
}

type Row struct {
	Id string
	Values []string
}

func (r Row) toValueMap(meta []metadata.ColumnMetadata) map[string]string {
	valueMap := map[string]string{metadata.ID_COLUMN: r.Id} //insert id as first value

	//we handle id seperately so we pop it out
	withoutId := meta[1:]
	for i, metaColumn := range withoutId {
		valueMap[metaColumn.ColumnName] = r.Values[i]
	}

	return valueMap
}

type FullTable struct {
	TableName string
	Rows []map[string]string
}