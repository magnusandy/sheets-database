package main

import (
	"log"
	"net/http"
	"sheets-database/api"
	"sheets-database/infra"
	"sheets-database/infra/configs"
	"sheets-database/domain"
	"sheets-database/domain/metadata"
)

type RootService struct {
	api api.Api
}

func initializeService() RootService {
	authenticationService := infra.RestAuthenticationService{Config:configs.Config{}}
	sheetService := infra.RestSheetsService{authenticationService}
	metadataService := metadata.CreateStubMetadata()
	dataService := domain.CreateDataService(sheetService, metadataService)
	return RootService{
		api.Api{DataService: dataService,
		AuthenticationService: authenticationService,
		},
	}
}

func main() {
	root := initializeService();
	http.HandleFunc("/", root.api.RootHandler)//todo serve help page
	//help
	//stats

	http.HandleFunc("/auth-link", root.api.CreateCredentialsHandler)
	http.HandleFunc("/submit-auth", root.api.SubmitAuthCodeHandler)
	//remove-auth"

	http.HandleFunc("/select", root.api.SelectHandler)
	http.HandleFunc("/insert", root.api.InsertDataHandler)
	//update
	//delete

	//select-query
	//update-query
	//delete-query

	//create-table
	//update-table

	//metadata-json
	//metadata html page

	log.Fatal(http.ListenAndServe(":8080", nil))
}
