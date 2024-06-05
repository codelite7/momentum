package queue

import (
	"context"
	"encoding/json"
	"github.com/acaloiaro/neoq"
	"github.com/acaloiaro/neoq/backends/postgres"
	"github.com/codelite7/momentum/api/config"
)

var queue neoq.Neoq

func Initialize() (err error) {
	queue, err = neoq.New(context.Background(), neoq.WithBackend(postgres.Backend), postgres.WithConnectionString(config.PostgresUri))
	if err != nil {
		return err
	}

	err = startHandlers(queue)
	if err != nil {
		return err
	}

	return nil
}

func startHandlers(queue neoq.Neoq) error {
	err := startWorkosUserCreatedEventHandler(queue)
	if err != nil {
		return err
	}
	err = startMessageEventHandler(queue)
	if err != nil {
		return err
	}
	return nil
}

func convertType(src, dest any) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dest)
}
