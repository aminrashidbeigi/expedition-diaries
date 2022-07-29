package main

import (
	"flag"
	"fmt"
	"log"

	"example.com/history-travelers/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	queries, err := GetQueries()

	if err != nil {
		log.Fatal(err)
	}
	seed := flag.Bool("seed", false, "This flag is for database seed")
	migrate := flag.Bool("migrate", false, "This flag is for database migrations")
	flag.Parse()

	if *migrate {
		storage.Migrate()
		fmt.Println("Database migrated")
	}

	if *seed {
		seeder := Seeder{queries: queries}
		seeder.SeedCountries()
		fmt.Println("Database seeded with countries")
	}
	api := Router{
		queries: queries,
	}
	router := gin.Default()
	router.GET("/countries/:code", api.getCountryByCode)
	router.GET("/countries", api.getCountries)
	router.POST("/add-resource", api.addResource)
	router.POST("/add-traveler", api.addTraveler)
	router.POST("/add-travel", api.addTravel)

	router.Run("localhost:8080")
}
