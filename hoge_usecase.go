package usecase

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
)

// HogeUsecase is hogeusecase
type HogeUsecase interface {
	Select(ids []int) ([]*model.Fuga, error)
	Create(entity *model.Fuga) (bool, error)
	Update(entity *model.Fuga) (bool, error)
	Delete(id int) (bool, error)
	Get(id int) (*model.Fuga, error)
}
