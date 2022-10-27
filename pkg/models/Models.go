package models

import "time"

type Hero struct {
	Id           int64     `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Power        int64     `json:"power"`
	Durability   int64     `json:"durability"`
	Intelligence int64     `json:"intelligence"`
	Img          string    `json:"img"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type PublicUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TeamMember struct {
	Id     int64 `json:"id" gorm:"primaryKey"`
	UserId int64 `json:"userId"`
	HeroId int64 `json:"heroId"`
}
