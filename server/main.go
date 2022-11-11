package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aminrashidbeigi/expedition-diaries/config"
	"github.com/aminrashidbeigi/expedition-diaries/endpoints"
	"github.com/aminrashidbeigi/expedition-diaries/middlewares"
	"github.com/aminrashidbeigi/expedition-diaries/storage"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	seed := flag.Bool("seed", false, "This flag is for database seed")
	migrate := flag.Bool("migrate", false, "This flag is for database migrations")
	configFile := flag.String("config", "", "This flag is for config file path")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Couldn't load config: %v", err)
	}

	db := storage.Storage{
		Config: cfg.DBConfig,
	}
	queries, err := db.GetQueries()

	if err != nil {
		log.Fatal(err)
	}

	if *migrate {
		db.Migrate()
		fmt.Println("Database migrated")
	}

	if *seed {
		db.SeedCountries()
		fmt.Println("Database seeded with countries")
	}
	api := endpoints.Router{
		Queries: queries,
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

	router.GET("/country-travels/:code", api.GetCountryTravelsByCode)
	router.GET("/countries", api.GetCountries)
	router.Use(authMiddleware.MiddlewareFunc())
	{
		router.POST("/add-resource", api.AddResource)
		router.POST("/add-traveler", api.AddTraveler)
		router.POST("/add-travel", api.AddTravel)
	}
	err = router.Run(cfg.Host + ":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
}
