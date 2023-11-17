package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	gorm *gorm.DB
}

type RepoInterface interface {

}

func NewRepo(gorm *gorm.DB) *BaseRepository {
	return &BaseRepository{gorm: gorm}
}