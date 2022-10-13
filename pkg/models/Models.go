package models

import "time"

type Hero struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Power        int64     `json:"power"`
	Durability   int64     `json:"durability"`
	Intelligence int64     `json:"intelligence"`
	ImgURL       string    `json:"imgURL"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
