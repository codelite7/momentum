package common

import (
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/codelite7/momentum/api/config"
	"github.com/codelite7/momentum/api/ent"
	"strings"
)

var EntClient *ent.Client

func InitializeEntClient() error {
	var err error
	EntClient, err = Open(config.PostgresUri)
	return err
}

// Open new connection
func Open(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "not found")
}
