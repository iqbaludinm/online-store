package repositories

import (
	"errors"

	"github.com/iqbaludinm/online-store/models"
	"gorm.io/gorm"
)

type CartRepository interface {
	CreateCart(cart models.Cart) (res models.Cart, err error)
	GetCarts() (carts []models.Cart, err error)
	GetCartById(id int64) (cart models.Cart, err error)
	UpdateCart(id int64, cart models.Cart) (res models.Cart, err error)
	DeleteCart(id int64, userId uint) (res models.Cart, err error)
}

// func method nempel dengan struct BaseRepository
func (r BaseRepository) CreateCart(cart models.Cart) (res models.Cart, err error) {
	err = r.gorm.Create(&cart).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r BaseRepository) GetCarts() (carts []models.Cart, err error) {
	err = r.gorm.Find(&carts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return carts, nil
		}
		return nil, err
	}
	return carts, err
}

func (r BaseRepository) GetCartById(id int64) (cart models.Cart, err error) {
	err = r.gorm.First(&cart, id).Error
	if err != nil {
		return cart, err
	}
	return cart, err
}

func (r BaseRepository) UpdateCart(id int64, cart models.Cart) (res models.Cart, err error) {
	err = r.gorm.Model(&cart).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Model(&res).Updates(cart).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r BaseRepository) DeleteCart(id int64, userId uint) (res models.Cart, err error) {
	cart := models.Cart{}
	err = r.gorm.First(&cart, id).Scan(&res).Error

	if err != nil {
		return res, err
	}
	
	if res.UserID != userId {
		return res, errors.New("unauthorized")
	}

	err = r.gorm.Delete(&cart, id, userId).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, err
}