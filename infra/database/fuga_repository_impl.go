package database

import (
	"github.com/ngiyshhk/golang-clean-arch-service/model"
)

type FugaRepositoryImpl struct{}

func (_ *FugaRepositoryImpl) Insert(entity *model.Fuga) (bool, error) {
	return false, nil
}
