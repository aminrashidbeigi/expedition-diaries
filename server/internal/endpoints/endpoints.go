package endpoints

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aminrashidbeigi/expedition-diaries/internal/sitemap"
	"github.com/aminrashidbeigi/expedition-diaries/storage"
	"github.com/aminrashidbeigi/expedition-diaries/storage/queries"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Queries          *queries.Queries
	SitemapGenerator *sitemap.SitemapGenerator
}

type AddTravelInput struct {
	Title     string   `json:"title" binding:"required"`
	StartedAt string   `json:"started_at"`
	EndedAt   string   `json:"ended_at"`
	Route     string   `json:"route"`
	Slug      string   `json:"slug"`
	Travelers []int    `json:"travelers" binding:"required"`
	Resources []int    `json:"resources" binding:"required"`
	Countries []string `json:"countries" binding:"required"`
}

type AddTravelerInput struct {
	Name        string `json:"name" binding:"required"`
	Link        string `json:"link"`
	Image       string `json:"image"`
	Nationality string `json:"nationality"`
}

type AddResourceInput struct {
	Title    string `json:"title" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Image    string `json:"image"`
	Language string `json:"language"`
	Type     string `json:"type"`
}

type Country struct {
	Code string
	Name string
}

type Traveler struct {
	Name        string
	Link        string
	Image       string
	Nationality string
}

type Resource struct {
	Title    string
	Link     string
	Image    string
	Language string
	Type     string
}

type Travel struct {
	Title     string
	Slug      string
	StartedAt string
	EndedAt   string
	Route     string
	Resources []Resource
	Travelers []Traveler
	Countries []Country
}

type CountryTravels struct {
	Country Country
	Travels []Travel
}

func (r Router) GetCountryTravelsByCode(c *gin.Context) {
	code := c.Param("code")
	code = strings.ToLower(code)
	if len(code) != 2 {
		log.Println("Bad request")
		c.IndentedJSON(http.StatusBadRequest, "Country code not found. it should be 2 letters.")
		return
	}

	countryRecord, err := r.Queries.GetCountryByCode(c, code)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Country code not found.")
		return
	}
	country := Country{
		Name: countryRecord.Name,
		Code: countryRecord.Code,
	}

	var travels []Travel
	countryTravels, err := r.Queries.GetTravlesByCountryCode(c, code)
	if err != nil {
		return
	}
	for _, travel := range countryTravels {
		travelers, err := r.Queries.GetTravelersByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		resources, err := r.Queries.GetResourcesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		countries, err := r.Queries.GetCountriesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		travels = append(travels, Travel{
			Title:     travel.Title,
			Slug:      travel.Slug.String,
			StartedAt: travel.StartedAt,
			EndedAt:   travel.EndedAt,
			Route:     travel.Route.String,
			Resources: resourcesRecordToResourceType(resources),
			Travelers: travelersRecordToTravelerType(travelers),
			Countries: countriesRecordToCountryType(countries),
		})
	}

	c.IndentedJSON(http.StatusOK, CountryTravels{Country: country, Travels: travels})
}

func (r Router) GetTravelBySlug(c *gin.Context) {
	slug := c.Param("slug")

	if slug == "" {
		c.IndentedJSON(http.StatusBadRequest, "slug should not be empty.")
		return
	}
	travelRecords, err := r.Queries.GetTravelBySlug(c, sql.NullString{String: slug, Valid: true})
	if err != nil {
		return
	}
	if len(travelRecords) == 0 {
		return
	}
	travelRecord := travelRecords[0]

	travelers, err := r.Queries.GetTravelersByTravelID(c, travelRecord.ID)
	if err != nil {
		return
	}

	resources, err := r.Queries.GetResourcesByTravelID(c, travelRecord.ID)
	if err != nil {
		return
	}

	countries, err := r.Queries.GetCountriesByTravelID(c, travelRecord.ID)
	if err != nil {
		return
	}

	travel := Travel{
		Title:     travelRecord.Title,
		Slug:      travelRecord.Slug.String,
		StartedAt: travelRecord.StartedAt,
		EndedAt:   travelRecord.EndedAt,
		Route:     travelRecord.Route.String,
		Resources: resourcesRecordToResourceType(resources),
		Travelers: travelersRecordToTravelerType(travelers),
		Countries: countriesRecordToCountryType(countries),
	}

	c.IndentedJSON(http.StatusOK, travel)
}

func (r Router) GetTravels(c *gin.Context) {
	limit64, err := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "limit format is not correct")
		return
	}
	limit := int32(limit64)

	offset64, err := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "offset format is not correct")
		return
	}
	offset := int32(offset64)

	var travels []Travel
	travelsRecords, err := r.Queries.GetTravels(c, queries.GetTravelsParams{
		Offset: offset,
		Limit:  limit,
	})

	if err != nil {
		return
	}
	for _, travel := range travelsRecords {
		travelers, err := r.Queries.GetTravelersByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		resources, err := r.Queries.GetResourcesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		countries, err := r.Queries.GetCountriesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		travels = append(travels, Travel{
			Title:     travel.Title,
			Slug:      travel.Slug.String,
			StartedAt: travel.StartedAt,
			EndedAt:   travel.EndedAt,
			Route:     travel.Route.String,
			Resources: resourcesRecordToResourceType(resources),
			Travelers: travelersRecordToTravelerType(travelers),
			Countries: countriesRecordToCountryType(countries),
		})
	}

	c.IndentedJSON(http.StatusOK, travels)
}

func (r Router) GetCountries(c *gin.Context) {
	countries, err := r.Queries.GetCountries(c)
	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "something bad happend.")
		return
	}
	c.IndentedJSON(http.StatusOK, countries)
}

func (r Router) AddTraveler(c *gin.Context) {
	var input AddTravelerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Input is wrong")
		return
	}

	traveler, err := r.Queries.CreateTraveler(c, queries.CreateTravelerParams{
		Name:        input.Name,
		Link:        input.Link,
		Image:       sql.NullString{String: input.Image, Valid: true},
		Nationality: sql.NullString{String: input.Nationality, Valid: true},
	})

	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "could not create traveler.")
		return
	}
	c.IndentedJSON(http.StatusCreated, traveler)
}

func (r Router) AddResource(c *gin.Context) {
	var input AddResourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Input is wrong")
		return
	}

	traveler, err := r.Queries.CreateResource(c, queries.CreateResourceParams{
		Title:    input.Title,
		Link:     input.Link,
		Image:    input.Image,
		Language: sql.NullString{String: input.Language, Valid: true},
		Type:     sql.NullString{String: input.Type, Valid: true},
	})

	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "could not create resource.")
		return
	}
	c.IndentedJSON(http.StatusCreated, traveler)
}

func (r Router) AddTravel(c *gin.Context) {
	var input AddTravelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Input is wrong: %v", err))
		return
	}

	travel, err := r.Queries.GetTravelByTitle(c, input.Title)
	if err != nil && !storage.IsNoRowError(err) {
		c.IndentedJSON(http.StatusInternalServerError, "could get create travel.")
		return
	}
	travelExistsAlready := false
	if travel.Title == input.Title {
		travelExistsAlready = true
	}

	if !travelExistsAlready {
		travel, err = r.Queries.CreateTravel(c, queries.CreateTravelParams{
			Title:     input.Title,
			StartedAt: input.StartedAt,
			EndedAt:   input.EndedAt,
			Route:     sql.NullString{String: input.Route, Valid: true},
			Slug:      sql.NullString{String: input.Slug, Valid: true},
		})
		if err != nil && !storage.IsNoRowError(err) {
			log.Println("this is error: ", err)
			c.IndentedJSON(http.StatusInternalServerError, "could not create travel.")
			return
		}
	}

	for _, countryInput := range input.Countries {
		country, err := r.Queries.GetCountryByCode(c, countryInput)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "could not find country.")
			return
		}
		_, err = r.Queries.CreateTravelCountry(c, queries.CreateTravelCountryParams{
			TravelID:  travel.ID,
			CountryID: country.ID,
		})
		if err != nil && !storage.IsNoRowError(err) {
			log.Println("this is error: ", err)
			c.IndentedJSON(http.StatusInternalServerError, "could not get country by id.")
			return
		}
	}

	for _, resourceInput := range input.Resources {
		_, err = r.Queries.CreateTravelResource(c, queries.CreateTravelResourceParams{
			TravelID:   travel.ID,
			ResourceID: int32(resourceInput),
		})
		if err != nil && !storage.IsNoRowError(err) {
			log.Println("this is error: ", err)
			c.IndentedJSON(http.StatusInternalServerError, "could not get resource by id.")
			return
		}
	}

	for _, travelerInput := range input.Travelers {
		_, err = r.Queries.CreateTravelTraveler(c, queries.CreateTravelTravelerParams{
			TravelID:   travel.ID,
			TravelerID: int32(travelerInput),
		})
		if err != nil && !storage.IsNoRowError(err) {
			log.Println("this is error: ", err)
			c.IndentedJSON(http.StatusInternalServerError, "could not get traveler by id.")
			return
		}
	}

	c.IndentedJSON(http.StatusCreated, travel)
}

func (r Router) GenerateSitemap(c *gin.Context) {
	r.SitemapGenerator.Generate()
	c.IndentedJSON(http.StatusOK, "Sitemap generated successfuly")
}

func resourcesRecordToResourceType(resourcesRecords []queries.GetResourcesByTravelIDRow) []Resource {
	var resources []Resource
	for _, resourceRecord := range resourcesRecords {
		resources = append(resources, Resource{
			Link:     resourceRecord.Link,
			Image:    resourceRecord.Image,
			Title:    resourceRecord.Title,
			Language: resourceRecord.Language.String,
			Type:     resourceRecord.Type.String,
		})
	}

	return resources
}

func travelersRecordToTravelerType(travelersRecords []queries.GetTravelersByTravelIDRow) []Traveler {
	var travelers []Traveler
	for _, travelerRecord := range travelersRecords {
		travelers = append(travelers, Traveler{
			Link:        travelerRecord.Link,
			Name:        travelerRecord.Name,
			Image:       travelerRecord.Image.String,
			Nationality: travelerRecord.Nationality.String,
		})
	}

	return travelers
}

func countriesRecordToCountryType(countriesRecords []queries.GetCountriesByTravelIDRow) []Country {
	var countries []Country
	for _, countryRecord := range countriesRecords {
		countries = append(countries, Country{
			Code: countryRecord.Code,
			Name: countryRecord.Name,
		})
	}

	return countries
}
