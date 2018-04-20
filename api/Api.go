package api

import (
	"net/http"
	"sheets-database/domain"
	"encoding/json"
	"sheets-database/api/dto"
	"html/template"
	"sheets-database/api/dto/in"
	"io/ioutil"
	"sheets-database/domain/metadata"
	"sheets-database/api/dto/out"
)

type Api struct {
	DataService           domain.DataService
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
	w.Write(nil) //todo help page?
}

func (api Api) SelectHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	dtoIn := in.SelectDto{}
	json.Unmarshal(body, &dtoIn)
	if dtoIn.Format == "" || dtoIn.Format == metadata.LIST {
		tableData := api.DataService.GetListData(dtoIn.SheetId, dtoIn.TableName);
		dtoOut := out.TableDtoFromDomain(tableData)
		json, err := json.Marshal(dtoOut)
		domain.LogIfPresent(err)
		w.Write(json)
	} else if dtoIn.Format == metadata.FULL {
		tableData := api.DataService.GetFullData(dtoIn.SheetId, dtoIn.TableName);
		outDto := out.FullTableDtoFromDomain(tableData)
		json, err := json.Marshal(outDto)
		domain.LogIfPresent(err)
		w.Write(json)
	}

}

func (api Api) InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	dtoIn := in.InsertDto{}
	json.Unmarshal(body, &dtoIn)
	err := api.DataService.InsertData(dtoIn.SheetId, dtoIn.ToDomain())
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func renderTemplate(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles("templates/" + fileName)
	t.Execute(w, data)
}
