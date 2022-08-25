package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aminrashidbeigi/history-travels/middlewares"
	"github.com/aminrashidbeigi/history-travels/storage"
	jwt "github.com/appleboy/gin-jwt/v2"
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

	authMiddleware, err := middlewares.AuthMiddleware()
	if err != nil {
		log.Fatal(err)
	}

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth")
	auth.POST("/login", authMiddleware.LoginHandler)
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)

	router.GET("/country-travels/:code", api.getCountryTravelsByCode)
	router.GET("/countries", api.getCountries)
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.POST("/add-resource", api.addResource)
		router.POST("/add-traveler", api.addTraveler)
		router.POST("/add-travel", api.addTravel)
	}
	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
