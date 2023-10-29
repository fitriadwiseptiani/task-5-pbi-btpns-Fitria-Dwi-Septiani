package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(300)" json:"username"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `gorm:"type:text" json:"password"`
	//  Photo `gorm:"foreignKey:Photos"`
	// Created_At int64
	// Updated_At int64
}
