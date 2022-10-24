package repository

import (
	"context"
	"cosmosdb-gin/pkg/logger"
	"cosmosdb-gin/pkg/model"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/google/wire"
)

type repo struct {
	client *azcosmos.Client
	logger.ILogger
}

func NewRepo(c *azcosmos.Client, log logger.ILogger) *repo {
	return &repo{
		c,
		log,
	}
}

func (r *repo) GetFamily(id string) (*model.Family, error) {
	c, err := r.client.NewContainer("ToDoList", "families")

	if err != nil {
		r.Error(err)
		return nil, fmt.Errorf("%v", err)
	}

	pk := azcosmos.NewPartitionKeyString(id)

	item, err := c.ReadItem(context.Background(), pk, id, nil)

	if err != nil {
		r.Error(err)
		return nil, fmt.Errorf("%v", err)
	}
	log.Println(string(item.Value))

	var family model.Family
	err = json.Unmarshal(item.Value, &family)

	if err != nil {
		r.Error(err)
		return nil, err
	}

	return &family, nil
}

func (r *repo) Hello() string {
	return "hi from repo"
}

var Wired = wire.NewSet(NewRepo, wire.Bind(new(IRepo), new(*repo)))
