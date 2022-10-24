package main

import (
	"cosmosdb-gin/pkg/config"
	"cosmosdb-gin/pkg/di"
	"cosmosdb-gin/pkg/logger"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
)

func main() {
	cfg := config.Load()

	cred, err := azcosmos.NewKeyCredential(cfg.DbKey)
	if err != nil {
		log.Fatal(err)
	}

	client, err := azcosmos.NewClientWithKey(cfg.DbEndpoint, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	app := di.CreateApi(cfg, client, logger.Must(logger.NewLogger()))
	app.Run(":8080")
}
