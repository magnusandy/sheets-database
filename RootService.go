package main

import (
	"log"
	"net/http"
	"sheets-database/api"
	"sheets-database/infra"
	"sheets-database/infra/configs"
)

type RootService struct {
	api api.Api
}

func initializeService() RootService {
	sheetService := infra.RestSheetsService{"1lLhDVyufI4GmiCNk3N1pibyRfQZ0nfXttLD6wKNb_Xo", "AIzaSyC0fuoGxv66q_ZAQRJybm2jcRfGD9XGmHI"}
	authenticationService := infra.RestAuthenticationService{Config:configs.Config{}}
	return RootService{
		api.Api{SheetService: sheetService,
		AuthenticationService: authenticationService,
		},
	}
}

func main() {
	root := initializeService();
	http.HandleFunc("/", root.api.RootHandler)
	http.HandleFunc("/auth-link", root.api.CreateCredentialsHandler)
	http.HandleFunc("/submit-auth", root.api.CreateCredentialsHandler)
	http.HandleFunc("/full-data", root.api.FullDataHandler)
	http.HandleFunc("/insert", root.api.InsertDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
