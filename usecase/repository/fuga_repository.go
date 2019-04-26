package repository

import "github.com/ngiyshhk/golang-clean-arch-usecase/model"

type FugaRepository interface {
	Insert(entity *model.Fuga) (bool, error)
}
