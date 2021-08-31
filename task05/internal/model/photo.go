package model

type Photo struct {
	Id   int    `json:"id"`
	AdId int    `json:"ad_id"`
	Link string `json:"link"`
}
