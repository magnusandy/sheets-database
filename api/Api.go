package api

import (
	"net/http"
	"sheets-database/domain"
	"encoding/json"
	"html/template"
	"sheets-database/api/dto/in"
	"io/ioutil"
	"sheets-database/domain/metadata"
	"sheets-database/api/dto/out"
	"sheets-database/api/dto/out/auth"
)

type Api struct {
	DataService           domain.DataService
	AuthenticationService domain.AuthenticationService
}

func (api Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write(nil) //todo help page?
}

func (api Api) CreateCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	link := api.AuthenticationService.CreateClientConfigLink()
	dto := auth.AuthPageData{link}
	renderTemplate(w, "authLink.html", dto)
}

func (api Api) SubmitAuthCodeHandler(w http.ResponseWriter, r *http.Request) {
	var authCode string = r.URL.Query().Get("authCode")
	if authCode != "" {
		tokenSaveError := api.AuthenticationService.SubmitClientConfig(authCode)

		dto := auth.AuthSuccess()
		if tokenSaveError != nil {
			dto = auth.AuthFailure(tokenSaveError)
		}
		renderTemplate(w, "authComplete.html", dto)
	}
}

func (api Api) SelectHandler(w http.ResponseWriter, r *http.Request) {

	dtoIn := in.SelectDto{}
	readBodyIntoDto(r, dtoIn)

	if dtoIn.Format == "" || dtoIn.Format == metadata.LIST {
		tableData := api.DataService.GetListData(dtoIn.SheetId, dtoIn.TableName);
		writeResponse(w, out.TableDtoFromDomain(tableData))
	} else if dtoIn.Format == metadata.FULL {
		tableData := api.DataService.GetFullData(dtoIn.SheetId, dtoIn.TableName);
		writeResponse(w, out.FullTableDtoFromDomain(tableData))
	}
}

func (api Api) InsertDataHandler(w http.ResponseWriter, r *http.Request) {

	dtoIn := in.InsertDto{}
	readBodyIntoDto(r, dtoIn)

	err := api.DataService.InsertData(dtoIn.SheetId, dtoIn.ToDomain())
	if err != nil {
		http.Error(w, err.Error(), 400)
	}
}

func renderTemplate(w http.ResponseWriter, fileName string, data interface{}) {
	t, _ := template.ParseFiles("templates/" + fileName)
	t.Execute(w, data)
}

func writeResponse(w http.ResponseWriter, outDto interface{}) {
	json, err := json.Marshal(outDto)
	domain.LogWithMessageIfPresent("JSON Marshal Error: ", err)
	w.Write(json)
}

func readBodyIntoDto(r *http.Request, inDtoAddress interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	domain.LogWithMessageIfPresent("Read Body Error: ", err)
	unmarshalError := json.Unmarshal(body, inDtoAddress)
	domain.LogWithMessageIfPresent("JSON Unmarshal Error: ", unmarshalError)
}
