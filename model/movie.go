package model

import (
	"time"

	"../utils"

	"github.com/jinzhu/gorm"
)

type Movie struct {
	gorm.Model
	Title       string
	Description string
	Date        time.Time
	Image       string

	Articles     []Article
	Applications []Application
}

func FindManyMovies() ([]Movie, error) {
	var movies []Movie
	db := utils.GetDB()
	err := db.Find(&movies).Error
	return movies, err
}

func FindMovieByID(id int) (Movie, error) {
	var movie Movie
	db := utils.GetDB()
	err := db.First(&movie, id).Error
	return movie, err
}
