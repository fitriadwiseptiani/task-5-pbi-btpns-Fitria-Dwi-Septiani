package models

type Photos struct {
	Id       int64  `gorm:"primaryKey" json:"id_photo"`
	Title    string `gorm:"type:varchar(300)" json:"title"`
	Caption  string `gorm:"type:varchar(300)" json:"caption"`
	PhotoURL string `gorm:"type:varchar(300)" json:"photo"`
	UserID   int64  `json:"id"`
}
