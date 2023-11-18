package repositories

import (
	"github.com/iqbaludinm/online-store/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product models.Product) (res models.Product, err error)
	GetProducts() (products []models.Product, err error)
	GetProductById(id int64) (product models.Product, err error)
	UpdateProduct(id int64, product models.Product) (res models.Product, err error)
	DeleteProduct(id int64) (res models.Product, err error)
}

func (r BaseRepository) CreateProduct(product models.Product) (res models.Product, err error) {
	err = r.gorm.Create(&product).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}

func (r BaseRepository) GetProducts() (products []models.Product, err error) {
	err = r.gorm.Preload("Category", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name")
	}).Find(&products).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return products, nil
		}
		return nil, err
	}

	return products, err
}

func (r BaseRepository) GetProductById(id int64) (product models.Product, err error) {
	err = r.gorm.First(&product, id).Error
	if err != nil {
		return product, err
	}
	return product, err
}

func (r BaseRepository) UpdateProduct(id int64, product models.Product) (res models.Product, err error) {
	err = r.gorm.Model(&product).Where("id = ?", id).First(&res).Error
	if err != nil {
		return res, err
	}

	err = r.gorm.Model(&res).Updates(product).Error
	if err != nil {
		return res, err
	}

	return res, err
}

func (r BaseRepository) DeleteProduct(id int64) (res models.Product, err error) {
	product := models.Product{}
	err = r.gorm.First(&product, id).Scan(&res).Error
	if err != nil {
		return res, err
	}
	
	err = r.gorm.Delete(&product, id).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, err
}