package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
)

type Ad interface {
	Create(ad model.Ad) (int, error)
	List(sortBy, order string) ([]model.Ad, error)
	Get(adId int, fields []string) (model.Ad, error)
}

type Repository struct {
	Ad
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ad: NewAdPostgres(db),
	}
}
