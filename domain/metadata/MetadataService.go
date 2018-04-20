package metadata

import (
	"errors"
)

const ID_COLUMN string = "id"

type MetadataService interface {
	GetMetadata(tableName string) (TableMetadata, error)
}

//todo replace with real service
type StubMetadataService struct {
	allMetadata map[string]TableMetadata //keyed by tablename
}

func (m StubMetadataService) GetMetadata(tableName string) (TableMetadata, error) {
	 meta := m.allMetadata[tableName]
	if meta.GetTableName() == "" {
		return meta, errors.New("Table metadata not found")
	}
	 return meta, nil
}

func CreateStubMetadata() StubMetadataService {
	users := CreateTableMetadata(
		"users",
		[]ColumnMetadata{
			CreateColumnMetadata(ID_COLUMN, TEXT, "", false),
			CreateColumnMetadata("name", TEXT, "-", false),
			CreateColumnMetadata("is_cool", BOOL, "false", false),
			CreateColumnMetadata("age", NUMBER, "18", false),
		})

	stubService := StubMetadataService{}
	stubService.allMetadata = map[string]TableMetadata{
		users.GetTableName(): users,
	}
	return stubService
}
