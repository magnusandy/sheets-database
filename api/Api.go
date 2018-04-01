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
	b := api.SheetService.GetAllData("", "")
	json, err := json.Marshal(b)
	domain.LogIfPresent(err);
	w.Write(json)
}
