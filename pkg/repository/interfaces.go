package repository

import (
	"context"
	"cosmosdb-gin/pkg/model"
)

type IRepo interface {
	Hello(ctx context.Context) string
	GetFamily(ctx context.Context, id string) (*model.Family, error)
}
