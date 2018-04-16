package configs

type Config struct {}

func (c Config) ServerSecretFileName() string {
	return "client_secret.json"
}

func (c Config) ClientOauthFileName() string {
	return "oauth_token.json"
}

func (c Config) ResourcesFolder() string {
	return "resources/"
}

func (c Config) ServerSecretPath() string {
	return c.ResourcesFolder() + c.ServerSecretFileName()
}

func (c Config) ClientOauthPath() string {
	return c.ResourcesFolder() + c.ClientOauthFileName()
}

func (c Config) GoogleScopes() []string {
	return []string {
		"https://www.googleapis.com/auth/spreadsheets.readonly",
		"https://www.googleapis.com/auth/spreadsheets",
		}
}
