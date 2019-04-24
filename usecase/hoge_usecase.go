package usecase

import (
	"github.com/ngiyshhk/golang-clean-arch-service/model"
)

type HogeUsecase interface {
	Create(entity *model.Fuga) (bool, error)
}
