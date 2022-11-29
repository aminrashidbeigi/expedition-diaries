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
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func main() {

	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://5c2865f9e4ca429db523bb3a503bbb56@o303173.ingest.sentry.io/4504140873072640",
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	seed := flag.Bool("seed", false, "This flag is for database seed")
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

	db.Migrate()

	if *seed {
		db.SeedCountries()
		fmt.Println("Database seeded with countries")
	}
	api := endpoints.Router{
		Queries: queries,
	}

	GenerateSitemap(queries, cfg.Url, cfg.Sitemaplocation)

	router := gin.Default()

	authMiddleware, err := middlewares.AuthMiddleware(cfg.User)
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
	router.GET("/travels", api.GetTravels)
	router.GET("/travels/:slug", api.GetTravelBySlug)
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
