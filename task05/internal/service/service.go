package service

import (
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Ad interface {
	Create(ad model.Ad) (int, error)
	List(sortBy, order string) ([]model.Ad, error)
	Get(adId int, fields []string) (model.Ad, error)
}

type Service struct {
	Ad
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ad: NewAdService(repos.Ad),
	}
}
