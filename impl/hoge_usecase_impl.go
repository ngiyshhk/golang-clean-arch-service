package impl

import (
	"errors"

	"github.com/ngiyshhk/golang-clean-arch-entity/model"
	"github.com/ngiyshhk/golang-clean-arch-usecase/repository"
)

// HogeUsecaseImpl is
type HogeUsecaseImpl struct {
	UserRepository repository.UserRepository
}

// Select is select hoge
func (usecase *HogeUsecaseImpl) Select(ids []int) ([]*model.User, error) {
	return usecase.UserRepository.Select(ids)
}

// Create is create hoge
func (usecase *HogeUsecaseImpl) Create(entity *model.User) (bool, error) {
	return usecase.UserRepository.Insert(entity)
}

// Update is update hoge
func (usecase *HogeUsecaseImpl) Update(entity *model.User) (bool, error) {
	return usecase.UserRepository.Update(entity)
}

// Delete is delete hoge
func (usecase *HogeUsecaseImpl) Delete(id int) (bool, error) {
	return usecase.UserRepository.Delete(id)
}

// Get is get a hoge
func (usecase *HogeUsecaseImpl) Get(id int) (*model.User, error) {
	entities, err := usecase.UserRepository.Select([]int{id})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, errors.New("not exist")
	}
	return entities[0], nil
}
