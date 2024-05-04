package run

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/codelite7/momentum/api/ent"
	"github.com/codelite7/momentum/api/resolvers"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
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
	// Create ent.Client and run the schema migration.
	client := Open(PostgresUri)
	err := client.Schema.Create(
		context.Background(),
	)
	if err != nil {
		return err
	}

	// Configure the server and start listening
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	// health endpoint
	http.Handle("/health", healthHandler{})
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	port := fmt.Sprintf(":%s", Port)
	log.Println(fmt.Sprintf("listening on %s", port))
	return http.ListenAndServe(port, nil)
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

type healthHandler struct{}

func (h healthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write(nil)
}
