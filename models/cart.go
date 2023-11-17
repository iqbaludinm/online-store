package models

type Cart struct {
	GormModel
	UserID   uint       `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User     *User      `json:"user_data"`
	Products []*Product `json:"products" gorm:"many2many:cart_product;"`
}

type InsertCart struct {
	UserID uint `json:"user_id" binding:"required" gorm:"not null" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8" `
}

type UpdateCart struct {
	UserID uint `json:"user_id" binding:"omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
}
