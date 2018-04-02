package main

import (
	"log"
	"net/http"
	"sheets-database/api"
	"sheets-database/infra"
)

type RootService struct {
	api api.Api
}

func initializeService() RootService {
	service := infra.RestSheetsService{"1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo", "AIzaSyC0fuoGxv66q_ZAQRJybm2jcRfGD9XGmHI"}
	return RootService{
		api.Api{SheetService: service},
	}
}

func main() {
	root := initializeService();
	http.HandleFunc("/", root.api.RootHandler)
	http.HandleFunc("/full-data", root.api.FullDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
