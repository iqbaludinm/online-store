package services

import "github.com/iqbaludinm/online-store/repositories"

type BaseService struct {
	repo repositories.RepoInterface
}

type ServiceInterface interface {
	UserService
	CartService
	CategoryService
	ProductService
	CartProductService
}

func NewService(repo repositories.RepoInterface) ServiceInterface {
	return &BaseService{repo: repo}
}