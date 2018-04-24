package infra

import (
	"sheets-database/domain"
	"sheets-database/domain/metadata"
	"google.golang.org/api/sheets/v4"
	"log"
)

const VISIBILITY = "DOCUMENT"

type RestMetadataService struct {
	authService domain.AuthenticationService
}

//todo split out
func CreateRestMetadataService(service domain.AuthenticationService) metadata.MetadataService {
	return RestMetadataService{service}
}

//MetadataService Methods
func (r RestMetadataService) GetTableMetadata(sheetId string, tableName string) (metadata.TableMetadata, error) {
	return metadata.TableMetadata{}, nil //todo
}

func (r RestMetadataService) CreateMetadata(sheetId string, meta metadata.TableMetadata) error {
	sheetClient, clientError := createSheetsClient(r.authService)
	if clientError != nil {
		return clientError
	}
	location := sheets.DeveloperMetadataLocation{Spreadsheet:true}
	devMeta := sheets.DeveloperMetadata{
		Visibility:VISIBILITY,
		Location: &location,
		MetadataId:1,
		MetadataKey:"metaTableInfo",
		MetadataValue:"myMetadata",
	}
	createDevMeta := sheets.CreateDeveloperMetadataRequest{DeveloperMetadata:&devMeta}
	request := sheets.Request{CreateDeveloperMetadata:&createDevMeta}
	batchUpdate := sheets.BatchUpdateSpreadsheetRequest{Requests:[]*sheets.Request{&request}}
	resp, err := sheetClient.Spreadsheets.BatchUpdate(sheetId, &batchUpdate).Do()
	log.Print(resp)
	//todo fetch current meta list
	//find if the value already exists
	//save if it doesnt
	return err
}

func (r RestMetadataService) GetDatabaseMetadata(sheetId string) ([]metadata.TableMetadata, error) {
	sheetClient, clientError := createSheetsClient(r.authService)
	if clientError != nil {
		return nil, clientError
	}

	resp, err := sheetClient.Spreadsheets.DeveloperMetadata.Get(sheetId, 1).Do()
	log.Print("metadata Resp")
	log.Print(resp)
	return nil, err
}

func (r RestMetadataService) UpdateMetadata(sheetId string, meta metadata.TableMetadata) error {
	return nil
}
