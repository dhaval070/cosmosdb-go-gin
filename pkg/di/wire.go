//go:build wireinject

package di

import (
	"cosmosdb-gin/pkg/api"
	"cosmosdb-gin/pkg/config"
	"cosmosdb-gin/pkg/handler"
	"cosmosdb-gin/pkg/logger"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/wire"
)

func CreateApi(cfg *config.Config, db *azcosmos.Client, log logger.ILogger) (http *api.HttpServer) {
	wire.Build(api.NewHttpServer, handler.Wired)

	return http
}
