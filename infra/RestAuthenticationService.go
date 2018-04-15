package infra

import (
	"sheets-database/infra/configs"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"sheets-database/domain"
	"net/http"
)

type RestAuthenticationService struct {
	Config configs.Config
}

func (service RestAuthenticationService) CreateClientConfigLink() string {
	ctx := context.Background()
	oauthConfig := service.configFromSecret(ctx)
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func (service RestAuthenticationService) SubmitClientConfig(authConfig string) {

}

func (service RestAuthenticationService) GetAuthenticatedClient() *http.Client {
	return nil;
}

func (service RestAuthenticationService) configFromSecret(ctx context.Context) *oauth2.Config {
	fileByteArray, fileError := ioutil.ReadFile(service.Config.ClientSecretPath())
	domain.LogWithMessageIfPresent("Unable to read client secret file", fileError);
	config, parseError := google.ConfigFromJSON(fileByteArray, service.Config.GoogleScopes()...)
	domain.LogWithMessageIfPresent("Unable to parse client secret file to config", parseError)
	return config
}
