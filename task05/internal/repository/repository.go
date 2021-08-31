package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
)

type Ad interface {
	Create(ad model.Ad) (int, error)
	List() ([]model.Ad, error)
	Get(adId int) (model.Ad, error)
}

type Repository struct {
	Ad
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ad: NewAdPostgres(db),
	}
}
