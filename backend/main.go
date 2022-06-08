package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db := GetDB()

	seed := flag.Bool("seed", false, "This flag is for database seed")
	flag.Parse()
	if *seed {
		seeder := Seeder{db: db}
		seeder.SeedCountries()
		fmt.Println("Database seeded with countries")
	}

	api := Router{
		db: db,
	}
	router := gin.Default()
	router.GET("/countries/:code", api.getCountryByCode)
	router.GET("/countries", api.getCountries)
	router.POST("/add-travel", api.addTravel)

	router.Run("localhost:8080")
}
