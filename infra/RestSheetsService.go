package infra

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sheets-database/domain"
	"sheets-database/api/dto"
	"log"
)
//https://developers.google.com/apis-explorer/?hl=en_GB#p/sheets/v4/sheets.spreadsheets.get?spreadsheetId=1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo&fields=sheets(data(rowData(values(note%252CuserEnteredValue))))%252CspreadsheetId&_h=1&
const GET_ALL_DATA_FIELD_FILTER = "sheets(data(rowData(values(note,userEnteredValue))),properties(sheetId,title)),spreadsheetId"

type RestSheetsService struct {
	SheetID string
	ApiKey string
}

func (r RestSheetsService) attachKey(url string) string {
	return url+"?key="+r.ApiKey
}

func attachFields(url string, fields string) {

}

func (r RestSheetsService) GetAllData(sheetId string, tableName string) []domain.Table {
	resp, _ := http.Get(r.attachKey(
		"https://sheets.googleapis.com/v4/spreadsheets/"+r.SheetID)+"&fields="+GET_ALL_DATA_FIELD_FILTER)
	var dto dto.GetAllData = dto.GetAllData{}
	deserializeBody(resp, &dto)
	log.Print(dto)
	return dto.ToDomain()
}

func deserializeBody(response *http.Response, i interface{}) {
	bodyAsBytes, bodyReadError := ioutil.ReadAll(response.Body)
	domain.LogIfPresent(bodyReadError)
	unmarshalError := json.Unmarshal(bodyAsBytes, i)
	domain.LogIfPresent(unmarshalError)
}

