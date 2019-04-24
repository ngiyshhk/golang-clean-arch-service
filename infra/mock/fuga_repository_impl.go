package mock

type FugaRepositoryImpl struct{}

func (_ *FugaRepositoryImpl) Insert() (bool, error) {
	return true, nil
}
