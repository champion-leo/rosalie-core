package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./model"
	"./routers"
	"./utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/joho/godotenv/autoload"
)

func testInsert(db *gorm.DB) {
	news := model.News{
		Title:           "News1",
		Author:          "Author",
		PublicationDate: time.Now(),
		UpdateDate:      time.Now(),
		Summary:         "News about a random movie",
		Content:         "News content blablabla\n blablablablabalandahuyegesgfegfuegfuegfeyges",
	}

	db.Create(&news)
	categorie1 := model.Categorie{
		Name:        "Meilleure Interprétation",
		Description: "",
	}
	db.Create(&categorie1)
	categorie2 := model.Categorie{
		Name:        "Meilleur Espoir",
		Description: "",
	}
	db.Create(&categorie2)
	categorie3 := model.Categorie{
		Name:        "Meilleur Scénario Original",
		Description: "",
	}
	db.Create(&categorie3)
	award1 := model.Award{
		CategorieID: 1,
	}
	db.Create(&award1)
	award2 := model.Award{
		CategorieID: 1,
	}
	db.Create(&award2)
	awardsResult := []model.Award{}
	db.Find(&awardsResult)
	for i, _ := range awardsResult {
		db.Model(awardsResult[i]).Related(&awardsResult[i].Categorie)
	}
	fmt.Println(awardsResult)
}

func createSchema(db *gorm.DB) {
	for _, model := range []interface{}{&model.Application{}, &model.Article{}, &model.Award{}, &model.Categorie{}, &model.Edition{}, &model.Movie{}, &model.News{}, &model.Person{}, &model.Question{}} {
		db.DropTableIfExists(model)
		db.CreateTable(model)
	}
}

func main() {
	db, err := utils.InitDB()
	if err != nil {
		log.Fatalf("DB: Cannot connect: %s\n", err)
	}
	defer db.Close()
	db.LogMode(true)
	createSchema(db)
	time.Sleep(2 * time.Second)
	testInsert(db)

	engine := gin.Default()
	router := engine.Group("/v1")
	routers.NewsRegister(router)
	routers.CategorieRegister(router)
	routers.EditionRegister(router)
	routers.QuestionRegister(router)
	routers.ArticleRegister(router)
	routers.MovieRegister(router)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "4000"
	}
	engine.Run(fmt.Sprintf(":%s", port))
}
