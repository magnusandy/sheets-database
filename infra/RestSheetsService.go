package infra

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sheets-database/domain"
	"sheets-database/api/dto"
	"log"
)

type RestSheetsService struct {
	SheetID string
	ApiKey string
}

func (r RestSheetsService) attachKey(url string) string {
	return url+"?key="+r.ApiKey
}

func (r RestSheetsService) GetAllData(sheetId string, tableName string) dto.GetAllData {
	resp, _ := http.Get(r.attachKey("https://sheets.googleapis.com/v4/spreadsheets/"+r.SheetID))
	var dto dto.GetAllData = dto.GetAllData{}
	deserializeBody(resp, &dto)
	log.Print(dto)
	return dto
}

func deserializeBody(response *http.Response, i interface{}) {
	bodyAsBytes, bodyReadError := ioutil.ReadAll(response.Body)
	domain.LogIfPresent(bodyReadError)
	unmarshalError := json.Unmarshal(bodyAsBytes, i)
	domain.LogIfPresent(unmarshalError)
}

