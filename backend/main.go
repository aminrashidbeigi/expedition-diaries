package main

import (
	"flag"
	"fmt"
	"log"

	"example.com/history-travelers/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	queries, err := storage.GetQueries()

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
		storage.SeedCountries()
		fmt.Println("Database seeded with countries")
	}
	api := Router{
		queries: queries,
	}
	router := gin.Default()
	router.GET("/country-travels/:code", api.getCountryTravelsByCode)
	router.GET("/countries", api.getCountries)
	router.POST("/add-resource", api.addResource)
	router.POST("/add-traveler", api.addTraveler)
	router.POST("/add-travel", api.addTravel)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
