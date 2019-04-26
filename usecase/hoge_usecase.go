package usecase

import (
	"github.com/ngiyshhk/golang-clean-arch-usecase/model"
)

type HogeUsecase interface {
	Create(entity *model.Fuga) (bool, error)
}
