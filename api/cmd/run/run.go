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
	"github.com/urfave/cli/v2"
	"log"
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
	client := initEntClient()
	err := createSchema(client)
	if err != nil {
		return err
	}
	graphqlServer := initGraphqlServer(client)
	return runHttpServer(graphqlServer)

}

func runHttpServer(graphqlServer *handler.Server) error {
	httpServer := initHttpServer(graphqlServer)
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

func initHttpServer(graphqlServer *handler.Server) *echo.Echo {
	e := echo.New()
	//e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health", func(c echo.Context) error {
		return nil
	})
	e.GET("/", echo.WrapHandler(playground.Handler("Todo", "/query")))
	e.POST("/query", echo.WrapHandler(graphqlServer))

	return e
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
}
