package impl

import (
	"github.com/ngiyshhk/golang-clean-arch-service/service/repository"
)

type HogeServiceImpl struct {
	FugaRepository repository.FugaRepository
}

func (service *HogeServiceImpl) Create() (bool, error) {
	return service.FugaRepository.Insert()
}
