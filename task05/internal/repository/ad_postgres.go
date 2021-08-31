package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/u-shylianok/golang-test-tasks/task05/internal/model"
)

type AdPostgres struct {
	db *sqlx.DB
}

func NewAdPostgres(db *sqlx.DB) *AdPostgres {
	return &AdPostgres{db: db}
}

func (r *AdPostgres) Create(ad model.Ad) (int, error) {
	var id int
	query := "INSERT INTO ads (name, date, price, description) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(query, ad.Name, ad.Date, ad.Price, ad.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AdPostgres) List() ([]model.Ad, error) {
	var ads []model.Ad

	query := "SELECT * FROM ads"
	err := r.db.Select(&ads, query)

	return ads, err
}

func (r *AdPostgres) Get(adId int) (model.Ad, error) {
	var ad model.Ad

	query := "SELECT * FROM ads WHERE id = $1"
	err := r.db.Get(&ad, query, adId)

	return ad, err
}
