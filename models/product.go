package models

type Product struct {
	GormModel
	CategoryID uint      `json:"category_id" gorm:"not null"`
	Category   *Category `json:"category_data"`
	Carts      []*Cart   `json:"carts" gorm:"many2many:cart_product;"`
	Name       string    `json:"name" binding:"required" gorm:"not null"`
	Stock      int       `json:"stock" gorm:"not null"`
	Price      int       `json:"price"`
}

type InsertProduct struct {
	CategoryID uint   `json:"category_id" binding:"required" gorm:"not null"`
	Name       string `json:"name" binding:"required" gorm:"not null"`
	Stock      int    `json:"stock" binding:"required" gorm:"not null"`
	Price      int    `json:"price" gorm:"not null"`
}

type UpdateProduct struct {
	CategoryID uint   `json:"category_id" binding:"omitempty"`
	Name       string `json:"name" binding:"omitempty"`
	Stock      int    `json:"stock" binding:"omitempty"`
	Price      int    `json:"price" binding:"omitempty"`
}
