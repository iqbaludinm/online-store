package services

import "github.com/iqbaludinm/online-store/models"

type CategoryService interface {
	CreateCategory(category models.InsertCategory) (res models.Category, err error)
	GetCategories() (categories []models.Category, err error)
	GetCategoryById(id int64) (category models.Category, err error)
	UpdateCategory(id int64, category models.UpdateCategory) (res models.Category, err error)
	DeleteCategory(id int64) (res models.Category, err error)
}

func (s *BaseService) CreateCategory(category models.InsertCategory) (models.Category, error) {
	cat := models.Category{
		Name: category.Name,
	}
	return s.repo.CreateCategory(cat)
}

func (s *BaseService) GetCategories() (categories []models.Category, err error) {
	return s.repo.GetCategories()
}

func (s *BaseService) GetCategoryById(id int64) (models.Category, error) {
	return s.repo.GetCategoryById(id)
}

func (s *BaseService) UpdateCategory(id int64, category models.UpdateCategory) (models.Category, error) {
	cat := models.Category{
		Name: category.Name,
	}
	return s.repo.UpdateCategory(id, cat)
}

func (s *BaseService) DeleteCategory(id int64) (models.Category, error) {
	return s.repo.DeleteCategory(id)
}