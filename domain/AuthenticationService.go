package domain

import "net/http"

type AuthenticationService interface {
	CreateClientConfigLink() string
	SubmitClientConfig(authConfig string) error
	GetAuthenticatedClient() (*http.Client, error)
}