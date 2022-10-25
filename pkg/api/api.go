package api

import (
	"cosmosdb-gin/pkg/handler"
	"cosmosdb-gin/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type HttpServer struct {
	c *gin.Engine
}

func NewHttpServer(h *handler.Handler, client appinsights.TelemetryClient) *HttpServer {
	router := gin.Default()

	router.Use(middleware.Telemetry(client))
	router.GET("/family/:id", h.GetFamily)
	router.GET("/hello", h.Hello)
	return &HttpServer{router}
}

func (s *HttpServer) Run(address string) error {
	return s.c.Run(address)
}
