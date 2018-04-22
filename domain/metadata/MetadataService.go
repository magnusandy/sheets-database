package metadata

import (
	"errors"
)

const ID_COLUMN string = "id"

type MetadataService interface {
	GetMetadata(sheetId string, tableName string) (TableMetadata, error)
}

//todo replace with real service
type StubMetadataService struct {
	allMetadata map[string]map[string]TableMetadata //keyed by sheetid, and then tablename
}

func (m StubMetadataService) GetMetadata(sheetId string, tableName string) (TableMetadata, error) {
	meta := m.allMetadata[sheetId][tableName]
	if meta.GetTableName() == "" {
		return meta, errors.New("Table metadata not found")
	}
	return meta, nil
}

func CreateStubMetadata() StubMetadataService {
	users := CreateTableMetadata(
		"1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo",
		"users",
		[]ColumnMetadata{
			CreateColumnMetadata(ID_COLUMN, TEXT, "", false),
			CreateColumnMetadata("name", TEXT, "-", false),
			CreateColumnMetadata("is_cool", BOOL, "false", false),
			CreateColumnMetadata("age", NUMBER, "18", false),
		})

	table := map[string]TableMetadata{
		users.GetTableName(): users,
	}
	sheetIdMap := map[string]map[string]TableMetadata{
		users.GetSheetId(): table,
	}
	return StubMetadataService{sheetIdMap}
}
