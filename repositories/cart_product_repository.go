package repositories

import (
	"github.com/iqbaludinm/online-store/models"
	"gorm.io/gorm"
)

type CartProductRepository interface {
	CreateCartProduct(cartProduct models.CartProduct) (res models.CartProduct, err error)
	GetCartProducts() (cartProducts []models.CartProduct, err error)
	GetCartProductById(id int64) (cartProduct models.CartProduct, err error)
	UpdateCartProduct(id int64, cartProduct models.CartProduct) (res models.CartProduct, err error)
	DeleteCartProduct(id int64) (res models.CartProduct, err error)
}

func (r BaseRepository) CreateCartProduct(cartProduct models.CartProduct) (res models.CartProduct, err error) {
	err = r.gorm.Create(&cartProduct).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r BaseRepository) GetCartProducts() (cartProducts []models.CartProduct, err error) {
	err = r.gorm.Find(&cartProducts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return cartProducts, nil
		}
		return nil, err
	}

	return cartProducts, err
}

func (r BaseRepository) GetCartProductById(id int64) (cartProduct models.CartProduct, err error) {
	err = r.gorm.First(&cartProduct, id).Error
	if err != nil {
		return cartProduct, err
	}
	return cartProduct, err
}

func (r BaseRepository) UpdateCartProduct(id int64, cartProduct models.CartProduct) (res models.CartProduct, err error) {
	err = r.gorm.Model(&cartProduct).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Model(&res).Updates(cartProduct).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r BaseRepository) DeleteCartProduct(id int64) (res models.CartProduct, err error) {
	cartProduct := models.CartProduct{}
	err = r.gorm.First(&cartProduct, id).Scan(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Delete(&cartProduct, id).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}
