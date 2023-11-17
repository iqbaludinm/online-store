package repositories

import (
	"github.com/iqbaludinm/online-store/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	CreateCategory(category models.Category) (res models.Category, err error)
	GetCategories() (categories []models.Category, err error)
	GetCategoryById(id int64) (category models.Category, err error)
	UpdateCategory(id int64, category models.Category) (res models.Category, err error)
	DeleteCategory(id int64) (res models.Category, err error)
}

func (r BaseRepository) CreateCategory(category models.Category) (res models.Category, err error) {
	err = r.gorm.Create(&category).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r BaseRepository) GetCategories() (categories []models.Category, err error) {
	err = r.gorm.Find(&categories).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return categories, nil
		}
		return nil, err
	}

	return categories, err
}

func (r BaseRepository) GetCategoryById(id int64) (category models.Category, err error) {
	err = r.gorm.First(&category, id).Error
	if err != nil {
		return category, err
	}
	return category, err
}

func (r BaseRepository) UpdateCategory(id int64, category models.Category) (res models.Category, err error) {
	err = r.gorm.Model(&category).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Model(&res).Updates(category).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r BaseRepository) DeleteCategory(id int64) (res models.Category, err error) {
	category := models.Category{}
	err = r.gorm.First(&category, id).Scan(&res).Error
	if err != nil {
		return res, err
	}
	
	err = r.gorm.Delete(&category, id).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}