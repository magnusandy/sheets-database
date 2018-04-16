package api

import (
	"net/http"
	"sheets-database/domain"
	"encoding/json"
	"sheets-database/api/dto"
	"html/template"
	"sheets-database/api/dto/in"
	"io/ioutil"
)

type Api struct {
	SheetService domain.SheetsService
	AuthenticationService domain.AuthenticationService
}

func (api Api) CreateCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	link := api.AuthenticationService.CreateClientConfigLink()
	dto := dto.AuthPageData{link}
	renderTemplate(w, "authLink.html", dto)
}

func (api Api) SubmitAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	var authCode string = r.URL.Query().Get("authCode")
	if authCode != "" {
		tokenSaveError := api.AuthenticationService.SubmitClientConfig(authCode)
		if tokenSaveError != nil {
			renderTemplate(w, "authComplete.html", dto.AuthFailure(tokenSaveError))
		} else {
			renderTemplate(w, "authComplete.html", dto.AuthSuccess())
		}
	}
}

func (api Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	b := api.SheetService.GetAllData()
	json, err := json.Marshal(b)
	domain.LogIfPresent(err);
	w.Write(json)
}

func (api Api) FullDataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	dtoIn := in.GetAllDataIn{}
	json.Unmarshal(body, &dtoIn)
	api.SheetService.GetAllDataForTable(dtoIn.SheetId, dtoIn.TableName);
}

func (api Api) InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	var tableNameQuery string = r.URL.Query().Get("tableName")
	if tableNameQuery != "" {
		api.SheetService.InsertRowIntoTable(tableNameQuery, domain.Row{"XXX", []string{"1", "true", "NULL", "okay hosay"}})
	}
}

func renderTemplate(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles("templates/"+fileName)
	t.Execute(w, data)
}



