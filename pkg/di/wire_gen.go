// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"cosmosdb-gin/pkg/api"
	"cosmosdb-gin/pkg/config"
	"cosmosdb-gin/pkg/handler"
	"cosmosdb-gin/pkg/logger"
	"cosmosdb-gin/pkg/repository"
	"cosmosdb-gin/pkg/usecase"
	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

// Injectors from wire.go:

func CreateApi(cfg *config.Config, db *azcosmos.Client, log logger.ILogger) *api.HttpServer {
	repo := repository.NewRepo(db, log)
	useCase := usecase.NewUseCase(repo, log)
	handlerHandler := handler.NewHandler(log, useCase)
	httpServer := api.NewHttpServer(handlerHandler)
	return httpServer
}
