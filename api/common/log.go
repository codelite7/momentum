package common

import "go.uber.org/zap"

var Logger *zap.Logger

func InitializeLogging() (err error) {
	Logger, err = zap.NewProduction()
	return
}
