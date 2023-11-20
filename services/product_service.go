package services

import "github.com/iqbaludinm/online-store/models"

type ProductService interface {
	CreateProduct(product models.InsertProduct) (res models.Product, err error)
	GetProducts(category string) (products []models.Product, err error)
	GetProductById(id int64) (product models.Product, err error)
	UpdateProduct(id int64, product models.UpdateProduct) (res models.Product, err error)
	DeleteProduct(id int64) (res models.Product, err error)
}

func (s *BaseService) CreateProduct(product models.InsertProduct) (models.Product, error) {
	prod := models.Product{
		Name: product.Name,
		CategoryID: product.CategoryID,
		Stock: product.Stock,
		Price: product.Price,
	}
	return s.repo.CreateProduct(prod)
}

func (s *BaseService) GetProducts(category string) (products []models.Product, err error) {
	return s.repo.GetProducts(category)
}

func (s *BaseService) GetProductById(id int64) (models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *BaseService) UpdateProduct(id int64, product models.UpdateProduct) (models.Product, error) {
	prod := models.Product{
		Name: product.Name,
		CategoryID: product.CategoryID,
		Stock: product.Stock,
		Price: product.Price,
	}
	return s.repo.UpdateProduct(id, prod)
}

func (s *BaseService) DeleteProduct(id int64) (models.Product, error) {
	return s.repo.DeleteProduct(id)
}