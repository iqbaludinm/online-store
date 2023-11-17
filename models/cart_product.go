package models

type CartProduct struct {
	GormModel
	CartID    uint `json:"cart_id" gorm:"foreignKey:CartID"`
	ProductID uint `json:"product_id" gorm:"foreignKey:ProductID"`
	Quantity  int  `json:"quantity"`
	Amount    int  `json:"amount"`
}

type InsertCartProduct struct {
	ProductID uint `json:"product_id" binding:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CartID    uint `json:"cart_id" binding:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Quantity  int  `json:"quantity" binding:"required"`
	Amount    int  `json:"amount" binding:"required"`
}

type UpdateCartProduct struct {
	ProductID uint `json:"product_id" binding:"omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CartID    uint `json:"cart_id" binding:"omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Quantity  int  `json:"quantity" binding:"omitempty"`
	Amount    int  `json:"amount" binding:"omitempty"`
}
