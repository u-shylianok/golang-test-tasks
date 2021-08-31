package service

import (
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/repository"
)

type Ad interface {
	Create(ad model.Ad) (int, error)
	List() ([]model.Ad, error)
	Get(adId int) (model.Ad, error)
}

type Service struct {
	Ad
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ad: NewAdService(repos.Ad),
	}
}
