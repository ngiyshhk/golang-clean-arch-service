package repository

type FugaRepository interface {
	Insert() (bool, error)
}
