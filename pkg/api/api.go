package api

import (
	"cosmosdb-gin/pkg/handler"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	c *gin.Engine
}

func NewHttpServer(h *handler.Handler) *HttpServer {
	router := gin.Default()

	router.GET("/family/:id", h.GetFamily)
	router.GET("/hello", h.Hello)
	return &HttpServer{router}
}

func (s *HttpServer) Run(address string) error {
	return s.c.Run(address)
}
