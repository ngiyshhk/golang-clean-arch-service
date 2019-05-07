package repository

import "github.com/ngiyshhk/golang-clean-arch-entity/model"

type FugaRepository interface {
	Insert(entity *model.Fuga) (bool, error)
	Update(entity *model.Fuga) (bool, error)
	Select(ids []int) ([]*model.Fuga, error)
	Delete(id int) (bool, error)
}
