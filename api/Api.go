package api

import (
	"net/http"
	"sheets-database/domain"
	"encoding/json"
	"fmt"
)

type Api struct {
	SheetService domain.SheetsService
	AuthenticationService domain.AuthenticationService
}

func (api Api) CreateCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	link := api.AuthenticationService.CreateClientConfigLink()
	fmt.Fprint(w, link)
}

func (api Api) RootHandler(w http.ResponseWriter, r *http.Request) {
	b := api.SheetService.GetAllData()
	json, err := json.Marshal(b)
	domain.LogIfPresent(err);
	w.Write(json)
}

func (api Api) FullDataHandler(w http.ResponseWriter, r *http.Request) {
	var tableNameQuery string = r.URL.Query().Get("tableName")
	if tableNameQuery != "" {
		table, err := api.SheetService.GetAllDataForTable(tableNameQuery)
		domain.LogIfPresent(err)
		json, err := json.Marshal(table)
		w.Write(json)
	}
}

func (api Api) InsertDataHandler(w http.ResponseWriter, r *http.Request) {
	var tableNameQuery string = r.URL.Query().Get("tableName")
	if tableNameQuery != "" {
		api.SheetService.InsertRowIntoTable(tableNameQuery, domain.Row{"XXX", []string{"1", "true", "NULL", "okay hosay"}})
	}
}



