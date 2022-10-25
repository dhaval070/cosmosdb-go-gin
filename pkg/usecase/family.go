package usecase

import (
	"context"
	"cosmosdb-gin/pkg/logger"
	"cosmosdb-gin/pkg/model"
	"cosmosdb-gin/pkg/repository"

	"github.com/google/wire"
)

type UseCase struct {
	repo repository.IRepo
	log  logger.ILogger
}

func NewUseCase(repo repository.IRepo, l logger.ILogger) *UseCase {
	return &UseCase{repo, l}
}

func (uc *UseCase) GetFamily(ctx context.Context, id string) (*model.Family, error) {
	var err error

	uc.log.Debug("getfamily")
	row, err := uc.repo.GetFamily(ctx, id)
	return row, err
}

var Wired = wire.NewSet(NewUseCase, repository.Wired, wire.Bind(new(IUseCase), new(*UseCase)))
