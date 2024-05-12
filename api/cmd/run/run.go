package run

import (
	"context"
	"database/sql"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/resolvers"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword"
	"github.com/supertokens/supertokens-golang/recipe/thirdpartyemailpassword/tpepmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"time"
)

var RunCommand = &cli.Command{
	Name:  "run",
	Usage: "Run the api",
	Flags: flags,
	Action: func(context *cli.Context) error {
		return run()
	},
}

var Client *ent.Client

func run() error {
	Client = initEntClient()
	err := createSchema(Client)
	if err != nil {
		return err
	}
	graphqlServer := initGraphqlServer(Client)
	return runHttpServer(graphqlServer)

}

func runHttpServer(graphqlServer *handler.Server) error {
	httpServer, err := initHttpServer(graphqlServer)
	if err != nil {
		return err
	}
	port := fmt.Sprintf(":%s", Port)
	return httpServer.Start(port)
}

func initEntClient() *ent.Client {
	return Open(PostgresUri)
}

func createSchema(client *ent.Client) error {
	return client.Schema.Create(
		context.Background(),
	)
}

func initGraphqlServer(client *ent.Client) *handler.Server {
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}

func initHttpServer(graphqlServer *handler.Server) (*echo.Echo, error) {
	err := initSuperTokens()
	if err != nil {
		return nil, err
	}
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{"*"},
		AllowMethods:                             []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:                             append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...),
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))
	e.Use(echo.WrapMiddleware(sessionMiddleware))
	e.Use(echo.WrapMiddleware(supertokens.Middleware))
	e.GET("/health", func(c echo.Context) error {
		return nil
	})
	e.GET("/", echo.WrapHandler(playground.Handler("Todo", "/query")))
	e.POST("/query", echo.WrapHandler(graphqlServer))

	return e, nil
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func initSuperTokens() error {
	return supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// These are the connection details of the app you created on supertokens.com
			ConnectionURI: SuperTokensConfig.Uri,
			APIKey:        SuperTokensConfig.ApiKey,
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "momentum",
			APIDomain:       SuperTokensConfig.ApiDomain,
			WebsiteDomain:   SuperTokensConfig.WebsiteDomain,
			APIBasePath:     &SuperTokensConfig.ApiBasePath,
			WebsiteBasePath: &SuperTokensConfig.WebsiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			thirdpartyemailpassword.Init(&tpepmodels.TypeInput{
				Providers: []tpmodels.ProviderInput{
					{
						Config: tpmodels.ProviderConfig{
							ThirdPartyId: "google",
							Clients: []tpmodels.ProviderClientConfig{
								{
									ClientID:     SuperTokensConfig.GoogleClientID,
									ClientSecret: SuperTokensConfig.GoogleClientSecret,
								},
							},
						},
					},
				},
				Override: &tpepmodels.OverrideStruct{
					Functions: func(originalImplementation tpepmodels.RecipeInterface) tpepmodels.RecipeInterface {
						original := *originalImplementation.EmailPasswordSignUp
						*originalImplementation.EmailPasswordSignUp = func(email string, password string, tenantId string, userContext supertokens.UserContext) (tpepmodels.SignUpResponse, error) {
							input := ent.CreateUserInput{
								CreatedAt: lo.ToPtr(time.Now()),
								UpdatedAt: lo.ToPtr(time.Now()),
								Email:     email,
							}
							var response tpepmodels.SignUpResponse
							tx, err := Client.Tx(context.Background())
							if err != nil {
								return response, err
							}
							defer tx.Rollback()
							user, err := tx.User.Create().SetInput(input).Save(context.Background())
							if err != nil {
								return response, err
							}
							response, err = original(email, password, tenantId, userContext)
							if err != nil {
								return response, err
							}

							if response.OK != nil {
								if response.OK.User.ID != "" {
									externalUserId := user.ID.String()
									_, err := supertokens.CreateUserIdMapping(response.OK.User.ID, externalUserId, nil, nil)
									if err != nil {
										return response, err
									}
									response.OK.User.ID = externalUserId
								}

							}
							err = tx.Commit()
							if err != nil {
								return response, err
							}
							return response, nil
						}
						return originalImplementation
					},
				},
			}),
			session.Init(&sessmodels.TypeInput{
				Override: &sessmodels.OverrideStruct{
					Functions: func(originalImplementation sessmodels.RecipeInterface) sessmodels.RecipeInterface {
						// First we copy the original implementation func
						originalCreateNewSession := *originalImplementation.CreateNewSession

						// Now we override the CreateNewSession function
						*originalImplementation.CreateNewSession = func(userID string, accessTokenPayload, sessionDataInDatabase map[string]interface{}, disableAntiCsrf *bool, tenantId string, userContext supertokens.UserContext) (sessmodels.SessionContainer, error) {
							// This goes in the access token, and is available to read on the frontend.
							if accessTokenPayload == nil {
								accessTokenPayload = map[string]interface{}{}
							}
							user, err := thirdpartyemailpassword.GetUserById(userID)
							if err != nil {
								return &sessmodels.TypeSessionContainer{}, err
							}
							accessTokenPayload["email"] = user.Email

							return originalCreateNewSession(userID, accessTokenPayload, sessionDataInDatabase, disableAntiCsrf, tenantId, userContext)
						}

						return originalImplementation
					},
				},
			}),
			dashboard.Init(nil),
		},
	})
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "postgres-uri",
		Aliases:     []string{"pu"},
		Value:       "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		Usage:       "postgres connection string",
		EnvVars:     []string{"POSTGRES_URI"},
		Destination: &PostgresUri,
	},
	&cli.StringFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Value:       "3000",
		Usage:       "port to run the server on",
		EnvVars:     []string{"PORT"},
		Destination: &Port,
	},
	&cli.StringFlag{
		Name:        "supertokens-uri",
		Aliases:     []string{"su"},
		Value:       "",
		Usage:       "uri to supertokens instance",
		EnvVars:     []string{"SUPERTOKENS_URI"},
		Destination: &SuperTokensConfig.Uri,
	},
	&cli.StringFlag{
		Name:        "supertokens-api-key",
		Aliases:     []string{"sak"},
		Value:       "",
		Usage:       "supertokens api key",
		EnvVars:     []string{"SUPERTOKENS_API_KEY"},
		Destination: &SuperTokensConfig.ApiKey,
	},
	&cli.StringFlag{
		Name:        "supertokens-api-domain",
		Aliases:     []string{"sad"},
		Value:       "http://localhost:3000",
		Usage:       "api address, this is self referential",
		EnvVars:     []string{"SUPERTOKENS_API_DOMAIN"},
		Destination: &SuperTokensConfig.ApiDomain,
	},
	&cli.StringFlag{
		Name:        "supertokens-website-domain",
		Aliases:     []string{"swd"},
		Value:       "http://localhost:4200",
		Usage:       "ui address",
		EnvVars:     []string{"SUPERTOKENS_WEBSITE_DOMAIN"},
		Destination: &SuperTokensConfig.WebsiteDomain,
	},
	&cli.StringFlag{
		Name:        "supertokens-api-base-path",
		Aliases:     []string{"sabp"},
		Value:       "/auth",
		Usage:       "api path where supertokens features are served from",
		EnvVars:     []string{"SUPERTOKENS_API_BASE_PATH"},
		Destination: &SuperTokensConfig.ApiBasePath,
	},
	&cli.StringFlag{
		Name:        "supertokens-website-base-path",
		Aliases:     []string{"swbp"},
		Value:       "/auth",
		Usage:       "ui path where supertokens features are handled",
		EnvVars:     []string{"SUPERTOKENS_WEBSITE_BASE_PATH"},
		Destination: &SuperTokensConfig.WebsiteBasePath,
	},
	&cli.StringFlag{
		Name:        "supertokens-google-client-id",
		Aliases:     []string{"sgci"},
		Value:       "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
		Usage:       "client id for google social login",
		EnvVars:     []string{"SUPERTOKENS_GOOGLE_CLIENT_ID"},
		Destination: &SuperTokensConfig.GoogleClientID,
	},
	&cli.StringFlag{
		Name:        "supertokens-google-client-secret",
		Aliases:     []string{"sgcs"},
		Value:       "",
		Usage:       "client secret for google social login",
		EnvVars:     []string{"SUPERTOKENS_GOOGLE_CLIENT_SECRET"},
		Destination: &SuperTokensConfig.GoogleClientSecret,
	},
}

func sessionMiddleware(next http.Handler) http.Handler {
	return session.VerifySession(&sessmodels.VerifySessionOptions{SessionRequired: lo.ToPtr(false)}, next.ServeHTTP)
}
