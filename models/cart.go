package models

import (
	"github.com/google/uuid"
)

type Cart struct {
	GormModel
	UserID   uuid.UUID `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	User     *User     `json:"user_data"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CartID;reference:ID"`
	Amount   int       `json:"amount"`
	IsPaid   bool      `json:"is_paid"`
}

type InsertCart struct {
	UserID uuid.UUID `json:"user_id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	IsPaid string    `json:"is_paid" binding:"required" gorm:"default:false" example:"true"`
}

type UpdateCart struct {
	UserID uuid.UUID `json:"user_id" binding:"omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	IsPaid string    `json:"is_paid" binding:"omitempty"`
}
