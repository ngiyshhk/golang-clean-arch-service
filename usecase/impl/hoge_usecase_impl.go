package impl

import (
	"github.com/ngiyshhk/golang-clean-arch-service/usecase/repository"
)

type HogeUsecaseImpl struct {
	FugaRepository repository.FugaRepository
}

func (usecase *HogeUsecaseImpl) Create() (bool, error) {
	return usecase.FugaRepository.Insert()
}
