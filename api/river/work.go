package river

import (
	"context"
	"github.com/codelite7/momentum/api/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/riverqueue/river"
	"github.com/riverqueue/river/riverdriver/riverpgxv5"
)

var RiverClient *river.Client[pgx.Tx]

func StartRiverWorkers() (*river.Workers, error) {
	workers := river.NewWorkers()
	// AddWorker panics if the worker is already registered or invalid:
	err := river.AddWorkerSafely[MessageArgs](workers, &MessageWorker{})
	return workers, err
}

func StartRiverClient() error {
	dbPool, err := pgxpool.New(context.Background(), config.PostgresUri)
	if err != nil {
		return err
	}

	workers, err := StartRiverWorkers()
	RiverClient, err = river.NewClient[pgx.Tx](riverpgxv5.New(dbPool), &river.Config{
		Queues: map[string]river.QueueConfig{
			river.QueueDefault: {MaxWorkers: 10},
		},
		Workers: workers,
	})
	if err != nil {
		return err
	}

	// Run the client inline. All executed jobs will inherit from ctx:
	err = RiverClient.Start(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func StopRiverClient() error {
	return RiverClient.Stop(context.Background())
}
