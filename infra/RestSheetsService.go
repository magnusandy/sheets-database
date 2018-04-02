package infra

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sheets-database/domain"
	"sheets-database/api/dto"
	"log"
	"errors"
	"bytes"
)

//https://developers.google.com/apis-explorer/?hl=en_GB#p/sheets/v4/sheets.spreadsheets.get?spreadsheetId=1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo&fields=sheets(data(rowData(values(note%252CuserEnteredValue))))%252CspreadsheetId&_h=1&
const GET_ALL_DATA_FIELD_FILTER = "sheets(data(rowData(values(note,userEnteredValue))),properties(sheetId,title)),spreadsheetId"

type RestSheetsService struct {
	SheetID string
	ApiKey  string
}

func (r RestSheetsService) attachKey(url string) string {
	return url + "?key=" + r.ApiKey
}

func (r RestSheetsService) GetAllData() []domain.Table {
	resp, _ := http.Get(r.attachKey(
		"https://sheets.googleapis.com/v4/spreadsheets/"+r.SheetID) + "&fields=" + GET_ALL_DATA_FIELD_FILTER)
	var dto dto.GetAllData = dto.GetAllData{}
	deserializeBody(resp, &dto)
	log.Print(dto)
	return dto.ToDomain()
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
func (r RestSheetsService) InsertRowIntoTable(tableName string, row domain.Row) error {
	json, err := json.Marshal(dto.FromDomain(row))
	domain.LogIfPresent(err)
	url := r.attachKey("https://sheets.googleapis.com/v4/spreadsheets/"+r.SheetID+"/values/"+tableName+":append")+"&valueInputOption=USER_ENTERED"
	log.Print(url)
	resp, postErr := http.Post(
		r.attachKey("https://sheets.googleapis.com/v4/spreadsheets/"+r.SheetID+"/values/"+tableName+":append")+"&valueInputOption=USER_ENTERED",
		"application/json",
		bytes.NewBuffer(json))
	log.Print(resp)
	domain.LogIfPresent(postErr)
	return err
}


func (r RestSheetsService) GetAllDataForTable(tableName string) (domain.Table, error) {
	var allTables []domain.Table = r.GetAllData()
	for i := 0; i < len(allTables); i++ {
		table := allTables[i]
		if table.TableName == tableName {
			return table, nil
		}
	}
	return domain.Table{}, errors.New("table not found")
}

func deserializeBody(response *http.Response, i interface{}) {
	bodyAsBytes, bodyReadError := ioutil.ReadAll(response.Body)
	domain.LogIfPresent(bodyReadError)
	unmarshalError := json.Unmarshal(bodyAsBytes, i)
	domain.LogIfPresent(unmarshalError)
}
