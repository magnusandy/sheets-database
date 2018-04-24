package infra

import (
	"google.golang.org/api/sheets/v4"
	"sheets-database/domain"
)

func createSheetsClient(service domain.AuthenticationService) (*sheets.Service, error) {
	httpClient, authError := service.GetAuthenticatedClient()
	if authError != nil {
		domain.LogWithMessageIfPresent("http server error", authError)
		return nil, authError
	} else {
		return sheets.New(httpClient)
	}
}