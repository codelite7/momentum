package run

var PostgresUri string
var Port string
var SuperTokensConfig = SuperTokensConfigg{}

type SuperTokensConfigg struct {
	Uri                string
	ApiKey             string
	ApiDomain          string
	WebsiteDomain      string
	ApiBasePath        string
	WebsiteBasePath    string
	GoogleClientID     string
	GoogleClientSecret string
}
