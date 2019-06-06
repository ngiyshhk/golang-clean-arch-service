package impl

import (
	"errors"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

type HogeUsecaseImpl struct {
	FugaRepository repository.FugaRepository
}

// Select is select hoge
func (usecase *HogeUsecaseImpl) Select(ids []int) ([]*model.Fuga, error) {
	return usecase.FugaRepository.Select(ids)
}

// Create is create hoge
func (usecase *HogeUsecaseImpl) Create(entity *model.Fuga) (bool, error) {
	return usecase.FugaRepository.Insert(entity)
}

// Update is update hoge
func (usecase *HogeUsecaseImpl) Update(entity *model.Fuga) (bool, error) {
	return usecase.FugaRepository.Update(entity)
}

// Delete is delete hoge
func (usecase *HogeUsecaseImpl) Delete(id int) (bool, error) {
	return usecase.FugaRepository.Delete(id)
}

// Get is get a hoge
func (usecase *HogeUsecaseImpl) Get(id int) (*model.Fuga, error) {
	entities, err := usecase.FugaRepository.Select([]int{id})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, errors.New("not exist")
	}
	return entities[0], nil
}
