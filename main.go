package main

import (
	"fmt"

	"github.com/ngiyshhk/golang-clean-arch-service/service"
	"github.com/ngiyshhk/golang-clean-arch-service/service/impl"
)

type fugaRepositoryMockImpl struct{}

func (a *fugaRepositoryMockImpl) Insert() (bool, error) {
	return false, nil
}

func main() {
	var hogeService service.HogeService = &impl.HogeServiceImpl{
		FugaRepository: &fugaRepositoryMockImpl{},
	}
	fmt.Println(hogeService.Create())
}
