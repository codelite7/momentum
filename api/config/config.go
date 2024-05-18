package config

var PostgresUri string
var Port string
var SuperTokensConfig = SuperTokensConfigg{}
var DefaultAgents string
var ApiLangchainBaseUrl string
var SessionRequired bool

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
