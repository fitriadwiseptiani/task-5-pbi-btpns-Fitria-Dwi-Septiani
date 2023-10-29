package models

type User struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	NamaPhoto string `gorm:"type:varchar(300)" json:"nama_photo"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
}
