package infra

import (
	"sheets-database/domain"
	"sheets-database/domain/metadata"
	"google.golang.org/api/sheets/v4"
	"log"
	"errors"
	"sheets-database/api/dto"
	"encoding/json"
)

const VISIBILITY = "DOCUMENT"
const META_ID = 1
const META_KEY = "tableInfo"


type RestMetadataService struct {
	authService domain.AuthenticationService
	//todo local cache of metadata
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
	//todo the logic should be in the service, this function should just delete and resave new meta

	sheetClient, clientError := createSheetsClient(r.authService)
	if clientError != nil {
		return clientError
	}
	//todo fetch current meta list
	currentMeta, getMetaErr := r.GetDatabaseMetadata(sheetId)
	if getMetaErr != nil {
		return getMetaErr
	}
	if currentMeta == nil { //if there is no current metadata
		currentMeta = map[string]metadata.TableMetadata{}
	}
	existingTable := currentMeta[meta.GetTableName()]

	if existingTable.GetTableName() == meta.GetTableName() { //default value if doesn't exist todo test
		return errors.New("table name: "+meta.GetTableName()+" already exists")
	}

	log.Print(meta)
	//add new table to the current meta
	currentMeta[meta.GetTableName()] = meta

   	deleteErr := r.deleteCurrentMetadata(sheetClient, sheetId)
	if deleteErr != nil {
		return deleteErr
	}

	saveErr := r.saveCurrentMetadata(sheetClient, sheetId, currentMeta)
	if saveErr != nil {
		return saveErr
	}

	return nil
}

func (r RestMetadataService) GetDatabaseMetadata(sheetId string) (map[string]metadata.TableMetadata, error) {
	sheetClient, clientError := createSheetsClient(r.authService)
	if clientError != nil {
		return nil, clientError
	}

	resp, err := sheetClient.Spreadsheets.DeveloperMetadata.Get(sheetId, META_ID).Do()
	log.Print("metadata Resp")
 	log.Print(resp.MetadataValue)
	return nil, err//todo not nil
}

func (r RestMetadataService) UpdateMetadata(sheetId string, meta metadata.TableMetadata) error {
	return nil
}

func (r RestMetadataService) deleteCurrentMetadata(sheetClient *sheets.Service, sheetId string) error {
	devMetaLookup := sheets.DeveloperMetadataLookup{MetadataId:META_ID} //we always use a hardcoded meta id
	dataFilter := sheets.DataFilter{DeveloperMetadataLookup:&devMetaLookup}
	deleteDevMeta := sheets.DeleteDeveloperMetadataRequest{DataFilter:&dataFilter}
	deleteRequest := sheets.Request{DeleteDeveloperMetadata:&deleteDevMeta}
	deleteBatch := sheets.BatchUpdateSpreadsheetRequest{Requests:[]*sheets.Request{&deleteRequest}}
	deleteResp, deleteErr := sheetClient.Spreadsheets.BatchUpdate(sheetId, &deleteBatch).Do()
	if deleteErr != nil{
		domain.LogIfPresent(deleteErr)
		return deleteErr
	}
	log.Print("delete response")
	log.Print(deleteResp)
	return nil;
}

func (r RestMetadataService) createCreateDeveloperMetadataRequest(currentMeta map[string]metadata.TableMetadata) (sheets.Request, error){
	dtoOut := dto.MetadataListDtoFromDomain(currentMeta)
	dtoBytes, marshalErr := json.Marshal(dtoOut)
	if marshalErr != nil {
		return sheets.Request{}, marshalErr
	}

	location := sheets.DeveloperMetadataLocation{Spreadsheet:true}
	devMeta := sheets.DeveloperMetadata{
		Visibility:VISIBILITY,
		Location: &location,
		MetadataId:META_ID,
		MetadataKey:META_KEY,
		MetadataValue:string(dtoBytes),
	}
	createDevMeta := sheets.CreateDeveloperMetadataRequest{DeveloperMetadata:&devMeta}
	return sheets.Request{CreateDeveloperMetadata:&createDevMeta}, nil
}

func (r RestMetadataService) saveCurrentMetadata(sheetClient *sheets.Service, sheetId string, currentMeta map[string]metadata.TableMetadata) error {
	createMetadataRequest, requestCreateError := r.createCreateDeveloperMetadataRequest(currentMeta)
	if requestCreateError != nil {
		return requestCreateError
	}
	batchUpdate := sheets.BatchUpdateSpreadsheetRequest{Requests:[]*sheets.Request{&createMetadataRequest}}
	resp, err := sheetClient.Spreadsheets.BatchUpdate(sheetId, &batchUpdate).Do()
	log.Print(resp)
	return err
}
