package model

import (
	"time"
)

type Ad struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Date        time.Time `json:"date"`
	Price       int       `json:"price"`
	Description *string   `json:"description,omitempty"`
	Photos      *[]Photo  `json:"photos"`
}
