package tables

import "sheets-database/domain/metadata"

type Row struct {
	id string
	values []string
}

func CreateRow(id string, values []string) Row {
	return Row{id, values}
}

func (r Row) GetId() string {
	return r.id
}

func (r Row) GetValues() []string {
	return r.values
}

func (r Row) ToValueMap(meta []metadata.ColumnMetadata) map[string]string {
	valueMap := map[string]string{metadata.ID_COLUMN: r.GetId()} //insert id as first value

	//we handle id seperately so we pop it out
	withoutId := meta[1:]
	for i, metaColumn := range withoutId {
		valueMap[metaColumn.GetColumnName()] = r.GetValues()[i]
	}

	return valueMap
}
