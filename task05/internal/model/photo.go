package model

type Photo struct {
	Id     int    `json:"id"`
	AdId   int    `json:"ad_id"`
	IsMain bool   `json:"is_main"`
	Link   string `json:"link"`
}
