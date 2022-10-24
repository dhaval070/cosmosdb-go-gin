package main

import (
	"cosmosdb-gin/pkg/config"
	"cosmosdb-gin/pkg/di"
	"cosmosdb-gin/pkg/logger"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
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

	insightsClient := appinsights.NewTelemetryClient(cfg.InstrumentationKey)

	insightsClient.SetIsEnabled(cfg.EnableTelemetry)
	log.Println("telemetry enabled: ", cfg.EnableTelemetry)

	logger := logger.Must(
		logger.NewLogger(insightsClient),
	)

	app := di.CreateApi(cfg, client, logger)
	app.Run(":8080")
}
