package configs

type Config struct {}

func (c Config) ClientSecretFileName() string {
	return "client_secret.json"
}

func (c Config) ClientSecretFileLocation() string {
	return "/Users/magnusandy/Documents/Programming/GO/src/sheets-database/resources/"
}

func (c Config) ClientSecretPath() string {
	return c.ClientSecretFileLocation() + c.ClientSecretFileName()
}

func (c Config) GoogleScopes() []string {
	return []string {
		"https://www.googleapis.com/auth/spreadsheets.readonly",
		"https://www.googleapis.com/auth/spreadsheets",
		}
}
