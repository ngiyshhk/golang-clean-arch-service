package repository

import "github.com/ngiyshhk/golang-clean-arch-entity/model"

// FugaRepository is
type FugaRepository interface {
	Insert(entity *model.Fuga) (bool, error)
	Update(entity *model.Fuga) (bool, error)
	Select(ids []int) ([]*model.Fuga, error)
	Delete(id int) (bool, error)
}

// FugaRepoMock is
type FugaRepoMock struct {
	FugaRepository
	MockInsert func(entity *model.Fuga) (bool, error)
	MockUpdate func(entity *model.Fuga) (bool, error)
	MockSelect func(ids []int) ([]*model.Fuga, error)
	MockDelete func(id int) (bool, error)
}

// Insert is
func (m *FugaRepoMock) Insert(entity *model.Fuga) (bool, error) {
	return m.MockInsert(entity)
}

// Update is
func (m *FugaRepoMock) Update(entity *model.Fuga) (bool, error) {
	return m.MockUpdate(entity)
}

// Select is
func (m *FugaRepoMock) Select(ids []int) ([]*model.Fuga, error) {
	return m.MockSelect(ids)
}

// Delete is
func (m *FugaRepoMock) Delete(id int) (bool, error) {
	return m.MockDelete(id)
}
