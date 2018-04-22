package infra

import (
	"sheets-database/domain"
	"google.golang.org/api/sheets/v4"
	"log"
	"sheets-database/domain/tables"
)

//https://developers.google.com/apis-explorer/?hl=en_GB#p/sheets/v4/sheets.spreadsheets.get?spreadsheetId=1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo&fields=sheets(data(rowData(values(note%252CuserEnteredValue))))%252CspreadsheetId&_h=1&
const GET_ALL_DATA_FIELD_FILTER = "sheets(data(rowData(values(note,userEnteredValue))),properties(sheetId,title)),spreadsheetId"

type RestSheetsService struct {
	authService domain.AuthenticationService
}

func CreateRestSheetService(authService domain.AuthenticationService) domain.SheetsService {
	return RestSheetsService{authService}
}

func (r RestSheetsService) GetAllData(sheetId string) []tables.Table {
	_, err := r.createSheetsClient()
	domain.LogIfPresent(err)
	return nil
}

/*
POST https://sheets.googleapis.com/v4/spreadsheets/1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo/values/Sheet2:append?valueInputOption=USER_ENTERED&key={YOUR_API_KEY}

{
 "values": [
  [
   "ID",
   "true",
   "122.3",
   "ID"
  ]
 ]
}
 */
func (r RestSheetsService) InsertRowIntoTable(tableName string, row tables.Row) error {
	return nil
}

func (r RestSheetsService) GetAllDataForTable(sheetId string, tableName string) (tables.Table, error) {
	log.Print(sheetId, " ",tableName)
	sheetClient, err := r.createSheetsClient()
	domain.LogWithMessageIfPresent("problem with sheet client connection", err)
	values, googleError := sheetClient.Spreadsheets.Values.Get(sheetId, tableName).MajorDimension("ROWS").Do()
	domain.LogWithMessageIfPresent("google sheets error", googleError)
	log.Print(values)
	return deserializeValueRangeToDomain(values, tableName), nil
}

func (r RestSheetsService) createSheetsClient() (*sheets.Service, error) {
	httpClient, authError := r.authService.GetAuthenticatedClient()
	if authError != nil {
		domain.LogWithMessageIfPresent("http server error", authError)
		return nil, authError
	} else {
		return sheets.New(httpClient)
	}
}

func deserializeValueRangeToDomain(valueRange *sheets.ValueRange, tableName string) tables.Table {
	var tableRows []tables.Row
	for _, typelessRow := range valueRange.Values {
		typedRow := typeCastRow(typelessRow)
		log.Print(typedRow)
		id := typedRow[0]
		values := make([]string, len(typedRow)-1)//dont need the first value as its already in the id
		copy(values, typedRow[1:])
		log.Print(values)
		domainRow := tables.CreateRow(id, values)
		tableRows = append(tableRows, domainRow)
	}
	return tables.CreateTable(tableName, tableRows)
}

func typeCastRow(row []interface{}) []string {
	newRow :=  []string{}
	for _, value := range row {
		newVal := value.(string)
		newRow = append(newRow, newVal)
	}
	return newRow
}
