package metadata

import (
	"errors"
)

const ID_COLUMN string = "id"

type MetadataService interface {
	GetTableMetadata(sheetId string, tableName string) (TableMetadata, error)
	GetDatabaseMetadata(sheetId string) ([]TableMetadata, error)
	CreateMetadata(sheetId string, meta TableMetadata) error
	UpdateMetadata(sheetId string, meta TableMetadata) error
}

//todo replace with real service
type StubMetadataService struct {
	allMetadata map[string]map[string]TableMetadata //keyed by sheetid, and then tablename
}

func (m StubMetadataService) GetTableMetadata(sheetId string, tableName string) (TableMetadata, error) {
	meta := m.allMetadata[sheetId][tableName]
	if meta.GetTableName() == "" {
		return meta, errors.New("Table metadata not found")
	}
	return meta, nil
}

func (m StubMetadataService) GetDatabaseeMetadata(sheetId string) ([]TableMetadata, error) {
	return nil, nil
}

func (m StubMetadataService) SaveMetadata(sheetId string, meta TableMetadata) error {
	return nil
}

func (m StubMetadataService) UpdateMetadata(sheetId string, meta TableMetadata) error {
	return nil
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
