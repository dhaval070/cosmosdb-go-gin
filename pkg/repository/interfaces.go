package repository

import "cosmosdb-gin/pkg/model"

type IRepo interface {
	Hello() string
	GetFamily(id string) (*model.Family, error)
}
