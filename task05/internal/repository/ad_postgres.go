package repository

import (
	"fmt"

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
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var adId int
	createAdQuery := "INSERT INTO ads (name, date, price, description) VALUES ($1, $2, $3, $4) RETURNING id"
	row := tx.QueryRow(createAdQuery, ad.Name, ad.Date, ad.Price, ad.Description)
	err = row.Scan(&adId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	for i := 0; i < len(*ad.Photos); i++ {

		var photoId int
		createPhotoQuery := "INSERT INTO photos (link) VALUES ($1) RETURNING id"
		row1 := tx.QueryRow(createPhotoQuery, (*ad.Photos)[i].Link)
		err = row1.Scan(&photoId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}

		var adsPhotosId int
		createAdsPhotosQuery := "INSERT INTO ads_photos (ad_id, photo_id, is_main) values ($1, $2, $3) RETURNING ad_id"
		row2 := tx.QueryRow(createAdsPhotosQuery, adId, photoId, i == 0) // main photo - first photo
		err = row2.Scan(&adsPhotosId)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return adId, tx.Commit()
}

func (r *AdPostgres) List() ([]model.Ad, error) {
	var ads []model.Ad

	query := "SELECT * FROM ads"

	if err := r.db.Select(&ads, query); err != nil {
		return nil, err
	}

	query2 := "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1"
	for i := 0; i < len(ads); i++ {

		var photos []model.Photo

		if err := r.db.Select(&photos, query2, ads[i].Id); err != nil {
			return nil, err
		}
		ads[i].Photos = &photos
	}

	return ads, nil
}

func (r *AdPostgres) Get(adId int, fields []string) (model.Ad, error) {
	var ad model.Ad

	fieldsToQueries := make(map[string]string)
	for _, a := range fields {
		fieldsToQueries[a] = a
	}

	if dsc, ok := fieldsToQueries["description"]; ok {
		fieldsToQueries["description"] = " , " + dsc
	}

	query := fmt.Sprintf("SELECT id, name, date, price%s FROM ads WHERE id = $1", fieldsToQueries["description"])
	if err := r.db.Get(&ad, query, adId); err != nil {
		return ad, err
	}

	var query2 string
	if _, ok := fieldsToQueries["photos"]; ok {
		query2 = "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1"
	} else {
		query2 = "SELECT photos.id, photos.link FROM photos INNER JOIN ads_photos ON ads_photos.photo_id = photos.id AND ads_photos.ad_id = $1 AND ads_photos.is_main"
	}

	var photos []model.Photo

	if err := r.db.Select(&photos, query2, ad.Id); err != nil {
		return ad, err
	}
	ad.Photos = &photos

	return ad, nil
}
