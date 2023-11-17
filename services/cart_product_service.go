package services

import "github.com/iqbaludinm/online-store/models"

type CartProductService interface {
	CreateCartProduct(cartProduct models.InsertCartProduct) (res models.CartProduct, err error)
	GetCartProducts() (cartProducts []models.CartProduct, err error)
	GetCartProductById(id int64) (cartProduct models.CartProduct, err error)
	UpdateCartProduct(id int64, cartProduct models.UpdateCartProduct) (res models.CartProduct, err error)
	DeleteCartProduct(id int64) (res models.CartProduct, err error)
}

func (s *BaseService) CreateCartProduct(cartProduct models.InsertCartProduct) (models.CartProduct, error) {
	cartProd := models.CartProduct{
		CartID: cartProduct.CartID,
		ProductID: cartProduct.ProductID,
		Quantity: cartProduct.Quantity,
		Amount: cartProduct.Amount,
	}
	return s.repo.CreateCartProduct(cartProd)
}

func (s *BaseService) GetCartProducts() (cartProducts []models.CartProduct, err error) {
	return s.repo.GetCartProducts()
}

func (s *BaseService) GetCartProductById(id int64) (models.CartProduct, error) {
	return s.repo.GetCartProductById(id)
}

func (s *BaseService) UpdateCartProduct(id int64, cartProduct models.UpdateCartProduct) (models.CartProduct, error) {
	cartProd := models.CartProduct{
		CartID: cartProduct.CartID,
		ProductID: cartProduct.ProductID,
		Quantity: cartProduct.Quantity,
		Amount: cartProduct.Amount,
	}
	return s.repo.UpdateCartProduct(id, cartProd)
}

func (s *BaseService) DeleteCartProduct(id int64) (models.CartProduct, error) {
	return s.repo.DeleteCartProduct(id)
}