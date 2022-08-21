package models

import "time"

// The ProductRefer holds the primary key ie Id of the product model
// The UserRefer holds the primary key ie Id of the user model

type Order struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}
