package impl

import (
	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

type HogeUsecaseImpl struct {
	FugaRepository repository.FugaRepository
}

func (usecase *HogeUsecaseImpl) Create(entity *model.Fuga) (bool, error) {
	return usecase.FugaRepository.Insert(entity)
}
