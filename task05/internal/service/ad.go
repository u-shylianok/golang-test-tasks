package service

import (
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/repository"
)

type AdService struct {
	repo repository.Ad
}

func NewAdService(repo repository.Ad) *AdService {
	return &AdService{repo: repo}
}

func (s *AdService) Create(ad model.Ad) (int, error) {
	return s.repo.Create(ad)
}

func (s *AdService) List(sortBy, order string) ([]model.Ad, error) {
	return s.repo.List(sortBy, order)
}

func (s *AdService) Get(adId int, fields []string) (model.Ad, error) {
	return s.repo.Get(adId, fields)
}
