package main

import (
	"log"
	"net/http"
	"sheets-database/api"
	"sheets-database/infra"
	"sheets-database/infra/configs"
	"sheets-database/domain"
)

type RootService struct {
	api api.Api
}

func initializeRootService() RootService {
	authenticationService := infra.CreateRestAuthenticationService(configs.Config{})
	sheetService := infra.CreateRestSheetService(authenticationService)
	metadataService := infra.CreateRestMetadataService(authenticationService)
	dataService := domain.CreateDataService(sheetService, metadataService)
	api := api.CreateApi(dataService, authenticationService)
	return RootService{api}
}

func main() {
	root := initializeRootService();
	http.HandleFunc("/", root.api.RootHandler) //todo serve help page
	//help
	//stats

	//GET
	http.HandleFunc("/auth-link", root.api.CreateCredentialsHandler)
	//GET
	http.HandleFunc("/submit-auth", root.api.SubmitAuthCodeHandler)
	//remove-auth"

	//POST
	http.HandleFunc("/select", root.api.SelectHandler)//todo maybe should be a get
	//POST
	http.HandleFunc("/insert", root.api.InsertDataHandler)
	//update
	//delete

	//select-query
	//update-query
	//delete-query

	http.HandleFunc("/create-table", root.api.CreateTableHandler)//create-table
	//update-table

	//POST
	http.HandleFunc("/database-info", root.api.GetDatabaseInfoHandler)

	//create-database
	//table-info
	//metadata html page

	log.Fatal(http.ListenAndServe(":8080", nil))
}
