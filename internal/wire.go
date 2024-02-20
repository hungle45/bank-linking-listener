//go:build wireinject
// +build wireinject

//go:generate wire
package internal

import (
	"demo/bank-linking-listener/config"
	"demo/bank-linking-listener/internal/delivery/consumer"
	httpHandler "demo/bank-linking-listener/internal/delivery/http"
	"demo/bank-linking-listener/internal/infrastructure/tidb"
	"demo/bank-linking-listener/internal/repository/tidb_repo"
	"demo/bank-linking-listener/internal/server"
	"demo/bank-linking-listener/internal/service"

	"github.com/google/wire"
)

var (
	configSet = wire.NewSet(config.LoadConfig)

	infrastructureSet = wire.NewSet(tidb.NewTiDB)

	repositorySet = wire.NewSet(
		tidb_repo.NewUserRepository,
		tidb_repo.NewBankRepository,
	)

	serviceSet = wire.NewSet(
		service.NewUserService,
		service.NewBankService,
	)

	controllerSet = wire.NewSet(
		httpHandler.NewController,
		consumer.NewController,
	)

	httpServerSet = wire.NewSet(server.NewHTTPServer)

	consumerJobSet = wire.NewSet(server.NewConsumerJob)
)

func InitializeHTTPServer(configFilePath string) server.Server {
	panic(
		wire.Build(
			configSet,
			infrastructureSet,
			repositorySet,
			serviceSet,
			controllerSet,
			httpServerSet,
		),
	)
}

func InitializeConsumerJob(configFilePath string) server.Server {
	panic(
		wire.Build(
			configSet,
			infrastructureSet,
			repositorySet,
			serviceSet,
			controllerSet,
			consumerJobSet,
		),
	)
}
