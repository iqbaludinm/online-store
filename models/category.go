package models

type Category struct {
	GormModel
	Name     string    `json:"name" binding:"required" gorm:"not null;unique"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CategoryID;reference:ID"`
}

type InsertCategory struct {
	Name string `json:"name" binding:"required" gorm:"not null"`
}

type UpdateCategory struct {
	Name string `json:"name" binding:"omitempty"`
}
