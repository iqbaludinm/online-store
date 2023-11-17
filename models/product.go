package models

import (
	"github.com/google/uuid"
)

type Product struct {
	GormModel
	CategoryID uuid.UUID `json:"category_id" gorm:"not null"`
	Category   *Category `json:"category_data"`
	CartID     uuid.UUID `json:"cart_id" gorm:"not null"`
	Cart       *Cart     `json:"cart_data"`
	Name       string    `json:"name" binding:"required" gorm:"not null"`
	Stock      int       `json:"stock" gorm:"not null"`
	Image      string    `json:"image"`
}

type InsertProduct struct {
	CategoryID uuid.UUID `json:"category_id" binding:"required" gorm:"not null"`
	CartID     uuid.UUID `json:"cart_id" binding:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name       string    `json:"name" binding:"required" gorm:"not null"`
	Stock      int       `json:"stock" binding:"required" gorm:"not null"`
	Image      string    `json:"image"`
}

type UpdateProduct struct {
	CategoryID uuid.UUID `json:"category_id" binding:"omitempty" gorm:"not null"`
	CartID     uuid.UUID `json:"cart_id" binding:"omitempty" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name       string    `json:"name" binding:"omitempty" gorm:"not null"`
	Stock      int       `json:"stock" binding:"omitempty" gorm:"not null"`
	Image      string    `json:"image"`
}
