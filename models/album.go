package models

type Album struct {
	ID     uint    `gorm:"primaryKey"` // GORM will use this as the primary key
	Title  string  `gorm:"type:varchar(100)"`
	Artist string  `gorm:"type:varchar(100)"`
	Price  float64 `gorm:"type:numeric(10,2)"`
}
