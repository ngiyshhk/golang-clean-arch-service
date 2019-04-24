package main

import (
	"fmt"

	"github.com/ngiyshhk/golang-clean-arch-service/infra/mock"
	"github.com/ngiyshhk/golang-clean-arch-service/model"

	"github.com/ngiyshhk/golang-clean-arch-service/usecase"
	"github.com/ngiyshhk/golang-clean-arch-service/usecase/impl"
)

func main() {
	var hogeUsecase usecase.HogeUsecase = &impl.HogeUsecaseImpl{
		FugaRepository: &mock.FugaRepositoryImpl{},
	}
	fmt.Println(hogeUsecase.Create(&model.Fuga{Id: 123, Name: "aaa", Age: 23}))
}
