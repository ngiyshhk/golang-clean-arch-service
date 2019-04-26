package usecase

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

type HogeUsecase interface {
	Create(entity *model.Fuga) (bool, error)
}
