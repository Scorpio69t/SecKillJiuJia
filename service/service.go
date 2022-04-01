package service

type Service interface {
	New(name string) (Service, error)
}
