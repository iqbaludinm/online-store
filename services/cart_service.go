package services

import "github.com/iqbaludinm/online-store/models"

type CartService interface {
	CreateCart(cart models.InsertCart) (res models.Cart, err error)
	GetCarts() (carts []models.Cart, err error)
	GetCartById(id int64) (cart models.Cart, err error)
	UpdateCart(id int64, cart models.UpdateCart) (res models.Cart, err error)
	DeleteCart(id int64, userId uint) (res models.Cart, err error)
}

func (s *BaseService) CreateCart(cart models.InsertCart) (models.Cart, error) {
	cartData := models.Cart{
		UserID: cart.UserID,
	}
	return s.repo.CreateCart(cartData)
}

func (s *BaseService) GetCarts() (carts []models.Cart, err error) {
	return s.repo.GetCarts()
}

func (s *BaseService) GetCartById(id int64) (models.Cart, error) {
	return s.repo.GetCartById(id)
}

func (s *BaseService) UpdateCart(id int64, cart models.UpdateCart) (models.Cart, error) {
	cartData := models.Cart{
		UserID: cart.UserID,
	}
	return s.repo.UpdateCart(id, cartData)
}

func (s *BaseService) DeleteCart(id int64, userId uint) (models.Cart, error) {
	return s.repo.DeleteCart(id, userId)
}