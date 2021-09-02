package model

import (
	"time"
)

type Ad struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Date        time.Time `json:"date" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	Description *string   `json:"description,omitempty" binding:"required"`
	Photos      *[]Photo  `json:"photos"`
}
