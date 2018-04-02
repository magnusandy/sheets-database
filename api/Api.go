package api

import (
	"net/http"
	"sheets-database/domain"
	"encoding/json"
)

type Api struct {
	SheetService domain.SheetsService
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



