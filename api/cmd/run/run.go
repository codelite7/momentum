package run

import (
	"context"
	"encoding/json"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/codelite7/momentum/api/cmd/run/queue"
	"github.com/codelite7/momentum/api/common"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/ent/agent"
	"github.com/codelite7/momentum/api/resolvers"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/workos/workos-go/v4/pkg/events"
	"github.com/workos/workos-go/v4/pkg/organizations"
	"github.com/workos/workos-go/v4/pkg/usermanagement"
	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
	"strings"
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

func run() error {
	err := common.InitializeLogging()
	if err != nil {
		return err
	}
	initializeWorkos()
	err = common.InitializeEntClient()
	if err != nil {
		return err
	}
	err = createSchema(common.EntClient)
	if err != nil {
		return err
	}
	err = ensureDefaultAgents()
	if err != nil {
		return err
	}
	graphqlServer := initGraphqlServer(common.EntClient)
	err = queue.Initialize()
	if err != nil {
		return err
	}
	go HandleAuthEvents()
	err = runHttpServer(graphqlServer)
	if err != nil {
		return err
	}

	return nil
}

func ensureDefaultAgents() error {
	var defaultAgents []*ent.Agent
	err := json.Unmarshal([]byte(config.DefaultAgents), &defaultAgents)
	if err != nil {
		return err
	}
	for _, agent := range defaultAgents {
		existingAgent, err := getAgentByProviderAndModel(agent.Provider, agent.Model)
		if err != nil && !strings.Contains(err.Error(), "not found") {
			return err
		}
		if existingAgent == nil {
			_, err = common.EntClient.Agent.Create().
				SetProvider(agent.Provider).
				SetModel(agent.Model).
				SetAPIKey(agent.APIKey).
				Save(context.Background())
			if err != nil {
				return err
			}
		} else {
			existingAgent.APIKey = agent.APIKey
			_, err = common.EntClient.Agent.UpdateOne(existingAgent).Save(context.Background())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func getAgentByProviderAndModel(provider, model string) (*ent.Agent, error) {
	return common.EntClient.Agent.Query().Where(agent.Provider(provider), agent.Model(model)).First(context.Background())
}
func runHttpServer(graphqlServer *handler.Server) error {
	httpServer, err := initHttpServer(graphqlServer)
	if err != nil {
		return err
	}
	port := fmt.Sprintf(":%s", config.Port)
	return httpServer.Start(port)
}

func createSchema(client *ent.Client) error {
	return client.Schema.Create(
		context.Background(),
	)
}

func initGraphqlServer(client *ent.Client) *handler.Server {
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		if asErr, ok := err.(error); ok {
			common.Logger.Error("recovered from panic", zap.Error(tracerr.Wrap(asErr)))
		} else {
			common.Logger.Error("recovered from panic", zap.Any("error", err))
		}
		return gqlerror.Errorf("Internal server error!")
	})

	return srv
}

type JwksResponse struct {
	Keys []Key `json:"keys"`
}
type Key struct {
	Alg     string   `json:"alg"`
	Kty     string   `json:"kty"`
	Use     string   `json:"use"`
	N       string   `json:"n"`
	E       string   `json:"e"`
	Kid     string   `json:"kid"`
	X5C     []string `json:"x5c"`
	X5TS256 string   `json:"x5t#S256"`
}

func initHttpServer(graphqlServer *handler.Server) (*echo.Echo, error) {
	//signingKeys, err := getJWKSKeys()
	//if err != nil {
	//	return nil, err
	//}
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(AuthMiddleware())
	//e.Use(echojwt.WithConfig(echojwt.Config{
	//	Skipper: func(c echo.Context) bool {
	//		if c.Request().Method == echo.OPTIONS {
	//			return true
	//		}
	//		if c.Request().URL.Path == "/health" {
	//			return true
	//		}
	//		return false
	//	},
	//	SigningKeys:            signingKeys,
	//	SigningMethod:          "RS256",
	//	ContinueOnIgnoredError: !config.ValidateJwt,
	//	ErrorHandler: func(c echo.Context, err error) error {
	//		common.Logger.Error("error getting jwt", zap.Error(err))
	//		return nil
	//	},
	//	SuccessHandler: func(c echo.Context) {
	//
	//	},
	//}))
	e.GET("/health", func(c echo.Context) error {
		return nil
	})
	e.GET("/", echo.WrapHandler(playground.Handler("Todo", "/query")))
	e.POST("/query", echo.WrapHandler(graphqlServer))

	return e, nil
}

func initializeWorkos() {
	events.SetAPIKey(config.WorkosApiKey)
	usermanagement.SetAPIKey(config.WorkosApiKey)
	organizations.SetAPIKey(config.WorkosApiKey)
}

var flags = []cli.Flag{
	&cli.StringFlag{
		Name:        "postgres-uri",
		Aliases:     []string{"pu"},
		Value:       "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
		Usage:       "postgres connection string",
		EnvVars:     []string{"POSTGRES_URI"},
		Destination: &config.PostgresUri,
	},
	&cli.StringFlag{
		Name:        "port",
		Aliases:     []string{"p"},
		Value:       "3001",
		Usage:       "port to run the server on",
		EnvVars:     []string{"PORT"},
		Destination: &config.Port,
	},
	&cli.StringFlag{
		Name:        "default-agents",
		Aliases:     []string{"da"},
		Value:       "[]",
		Usage:       "json array of agent objects that will be created on startup",
		EnvVars:     []string{"DEFAULT_AGENTS"},
		Destination: &config.DefaultAgents,
	},
	&cli.StringFlag{
		Name:        "api-langchain-base-url",
		Aliases:     []string{"albu"},
		Value:       "http://localhost:6543",
		Usage:       "base url to api-langchain",
		EnvVars:     []string{"API_LANGCHAIN_BASE_URL"},
		Destination: &config.ApiLangchainBaseUrl,
	},
	&cli.BoolFlag{
		Name:        "session-required",
		Aliases:     []string{"sr"},
		Value:       false,
		Usage:       "when true a session is required to reach the api",
		EnvVars:     []string{"SESSION_REQUIRED"},
		Destination: &config.SessionRequired,
	},
	&cli.StringFlag{
		Name:        "workos-client-id",
		Aliases:     []string{"wci"},
		Usage:       "workos client id",
		EnvVars:     []string{"WORKOS_CLIENT_ID"},
		Destination: &config.WorkosClientId,
	},
	&cli.StringFlag{
		Name:        "workos-api-key",
		Aliases:     []string{"wak"},
		Usage:       "workos api key",
		EnvVars:     []string{"WORKOS_API_KEY"},
		Destination: &config.WorkosApiKey,
	},
	&cli.DurationFlag{
		Name:        "workos-poll-interval",
		Aliases:     []string{"wpo"},
		Value:       10 * time.Second,
		Usage:       "How often to poll workos for events",
		EnvVars:     []string{"WORKOS_POLL_INTERVAL"},
		Destination: &config.WorkosPollInterval,
	},
	&cli.BoolFlag{
		Name:        "validate-jwt",
		Aliases:     []string{"vj"},
		Value:       true,
		Usage:       "When true, jwt tokens are validated against the workos JWKS",
		EnvVars:     []string{"VALIDATE_JWT"},
		Destination: &config.ValidateJwt,
	},
}
