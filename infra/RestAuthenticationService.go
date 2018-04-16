package infra

import (
	"sheets-database/infra/configs"
	"golang.org/x/oauth2"
	"golang.org/x/net/context"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"sheets-database/domain"
	"net/http"
	"os"
	"encoding/json"
)

type RestAuthenticationService struct {
	Config configs.Config
}

func (service RestAuthenticationService) CreateClientConfigLink() string {
	ctx := context.Background()
	oauthConfig := service.configFromServerSecretFile(ctx)
	authURL := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func (service RestAuthenticationService) SubmitClientConfig(authConfig string) error {
	ctx := context.Background()
	oauthConfig := service.configFromServerSecretFile(ctx)
	oauthToken, webRetrieveError := oauthConfig.Exchange(ctx, authConfig)
	if webRetrieveError != nil {
		domain.LogWithMessageIfPresent("unable to retrieve client token", webRetrieveError)
		return webRetrieveError
	} else {
		return service.saveClientOauthToken(oauthToken)
	}
}

func (service RestAuthenticationService) GetAuthenticatedClient() (*http.Client, error) {
	ctx := context.Background()
	oauthConfig := service.configFromServerSecretFile(ctx)
	token, err := service.clientTokenFromFile()
	return oauthConfig.Client(ctx, token), err
}

func (service RestAuthenticationService) configFromServerSecretFile(ctx context.Context) *oauth2.Config {
	fileByteArray, fileError := ioutil.ReadFile(service.Config.ServerSecretPath())
	domain.LogWithMessageIfPresent("Unable to read client secret file", fileError);
	config, parseError := google.ConfigFromJSON(fileByteArray, service.Config.GoogleScopes()...)
	domain.LogWithMessageIfPresent("Unable to parse client secret file to config", parseError)
	return config
}

func (service RestAuthenticationService) clientTokenFromFile() (*oauth2.Token, error) {
	oauthFile, err := os.Open(service.Config.ClientOauthPath())
	if err != nil {
		return nil, err
	}
	oauthToken := &oauth2.Token{}
	err = json.NewDecoder(oauthFile).Decode(oauthToken)
	defer oauthFile.Close()
	return oauthToken, err
}

func (service RestAuthenticationService) saveClientOauthToken(token *oauth2.Token) error {
	file, err := os.OpenFile(service.Config.ClientOauthPath(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		domain.LogWithMessageIfPresent("problem saving oauth file", err)
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)

	return nil

}
