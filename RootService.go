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
	authenticationService := infra.RestAuthenticationService{Config:configs.Config{}}
	sheetService := infra.RestSheetsService{authenticationService}
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
	http.HandleFunc("/submit-auth", root.api.SubmitAuthCodeHandler)
	//http.HandleFunc("/remove-auth", root.api.TODO)
	http.HandleFunc("/full-data", root.api.FullDataHandler)
	http.HandleFunc("/insert", root.api.InsertDataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//insert
	//update
	//delete
	//find
}
