package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormModel struct {
	ID        uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp;not null"`
}

func (model *GormModel) BeforeCreate(tx *gorm.DB) (err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Error LoadLocation: ", err)
		return time.Now()
	}
	model.CreatedAt = time.Now().In(location)
	model.UpdatedAt = time.Now().In(location)
	return nil
}

func (model *GormModel) BeforeUpdate(tx *gorm.DB) (err error) {
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("Error LoadLocation: ", err)
		return time.Now()
	}
	model.UpdatedAt = time.Now().In(location)
	return nil
}
