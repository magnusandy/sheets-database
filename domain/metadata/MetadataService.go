package metadata

type MetadataService interface {
	GetMetadata(tableName string) TableMetadata
}

//todo replace with real service
type StubMetadataService struct {
	allMetadata map[string]TableMetadata //keyed by tablename
}

func (m StubMetadataService) GetMetadata(tableName string) TableMetadata {
	return m.allMetadata[tableName]
}

func CreateStubMetadata() *StubMetadataService {
	users := TableMetadata{
		"users",
		[]ColumnMetadata{
			ColumnMetadata{"id", TEXT, "", false},
			ColumnMetadata{"name", TEXT, "-", false},
			ColumnMetadata{"is_cool", BOOL, "false", false},
			ColumnMetadata{"age", NUMBER, "18", false},
		},
	}
	stubService := StubMetadataService{}
	stubService.allMetadata = map[string]TableMetadata{
		users.TableName: users,
	};
	return &stubService
}
