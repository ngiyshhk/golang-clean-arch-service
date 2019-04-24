package main

import (
	"fmt"

	"github.com/ngiyshhk/golang-clean-arch-service/infra/database"
	"github.com/ngiyshhk/golang-clean-arch-service/usecase"
	"github.com/ngiyshhk/golang-clean-arch-service/usecase/impl"
)

func main() {
	var hogeUsecase usecase.HogeUsecase = &impl.HogeUsecaseImpl{
		FugaRepository: &database.FugaRepositoryImpl{},
	}
	fmt.Println(hogeUsecase.Create())
}
