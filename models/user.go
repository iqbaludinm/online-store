package models

type User struct {
	GormModel
	Username string `json:"username" binding:"required" gorm:"not null;uniqueIndex"`
	Email    string	`json:"email" binding:"required, email" gorm:"not null;uniqueIndex"`
	Password string `json:"password" binding:"min=6,required" gorm:"not null"`
	Address  string `json:"address" binding:"required" gorm:"not null"`
}


type RegisterUser struct {
	Username string `json:"username" form:"username" example:"iqbaludinm" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email" example:"iqbal@mail.com"`
	Password string `json:"password" form:"password" binding:"required" example:"iqbal123"`
	Address int `json:"address" form:"address" binding:"required" example:"Jakarta"`
}

type LoginUser struct {
	Email    string `json:"email" form:"email" binding:"required,email" example:"iqbal@mail.com"`
	Password string `json:"password" form:"password" binding:"required" example:"iqbal123"`
}
