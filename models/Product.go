package models

type Product struct {
	ID          uint   `diff:"-" gorm:"primarykey" json:"id"`
	Name        string `gorm:"size:255;not null" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}

type Vote struct {
	ID        uint `diff:"-" gorm:"primarykey" json:"id"`
	ProductID uint `gorm:"not null" json:"productId"`
}
