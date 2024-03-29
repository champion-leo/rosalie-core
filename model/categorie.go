package model

import (
	"../utils"
	"github.com/jinzhu/gorm"
)

type Categorie struct {
	gorm.Model
	Name        string
	Description string

	Awards []Award
}

func FindManyCategories() ([]Categorie, error) {
	var categories []Categorie
	db := utils.GetDB()
	err := db.Find(&categories).Error
	return categories, err
}

func FindCategorieByID(id uint) (Categorie, error) {
	var categorie Categorie
	db := utils.GetDB()
	err := db.First(&categorie, id).Error
	return categorie, err
}
