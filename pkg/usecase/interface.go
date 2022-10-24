package usecase

import (
	"context"
	"cosmosdb-gin/pkg/model"
)

type IUseCase interface {
	GetFamily(ctx context.Context, id string) (*model.Family, error)
}
