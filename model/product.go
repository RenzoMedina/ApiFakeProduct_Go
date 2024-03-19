package model

import "time"

type Product struct {
	Id          uint
	Name        string
	Price       float64
	Quantity    int
	Description string
	Create_at   time.Time
	Update_at   time.Time
}
type Products []*Product

type Storage interface {
	Migrate() error
	//Create(*Product) error
	//Update(*Product) error
	//GetAll() (*Product, error)
	//GetById(uint) (*Product, error)
	//Delete(uint) error
}

type Service struct {
	storage Storage
}

func NewServices(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
