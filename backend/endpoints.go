package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/aminrashidbeigi/history-travels/storage"
	"github.com/aminrashidbeigi/history-travels/storage/queries"
	"github.com/gin-gonic/gin"
)

type Router struct {
	queries *queries.Queries
}

type AddTravelInput struct {
	Title     string   `json:"title" binding:"required"`
	StartedAt string   `json:"started_at"`
	EndedAt   string   `json:"ended_at"`
	Travelers []int    `json:"travelers" binding:"required"`
	Resources []int    `json:"resources" binding:"required"`
	Countries []string `json:"countries" binding:"required"`
}

type AddTravelerInput struct {
	Name string `json:"name" binding:"required"`
	Link string `json:"link" binding:"required"`
}

type AddResourceInput struct {
	Title string `json:"title" binding:"required"`
	Link  string `json:"link" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type Country struct {
	Code string
	Name string
}

type Traveler struct {
	Name string
	Link string
}

type Resource struct {
	Title string
	Link  string
	Image string
}

type Travel struct {
	Resources []Resource
	Travelers []Traveler
	Countries []Country
}

type CountryTravels struct {
	Country Country
	Travels []Travel
}

func (r Router) getCountryTravelsByCode(c *gin.Context) {
	code := c.Param("code")
	code = strings.ToLower(code)
	if len(code) != 2 {
		log.Println("Bad request")
		c.IndentedJSON(http.StatusBadRequest, "Country code not found. it should be 2 letters.")
		return
	}

	countryRecord, err := r.queries.GetCountryByCode(c, code)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, "Country code not found.")
		return
	}
	country := Country{
		Name: countryRecord.Name,
		Code: countryRecord.Code,
	}

	var travels []Travel
	countryTravels, err := r.queries.GetTravlesByCountryCode(c, code)
	if err != nil {
		return
	}
	for _, travel := range countryTravels {
		travelers, err := r.queries.GetTravelersByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		resources, err := r.queries.GetResourcesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		countries, err := r.queries.GetCountriesByTravelID(c, travel.ID)
		if err != nil {
			return
		}

		travels = append(travels, Travel{
			Resources: resourcesRecordToResourceType(resources),
			Travelers: travelersRecordToTravelerType(travelers),
			Countries: countriesRecordToCountryType(countries),
		})
	}

	c.IndentedJSON(http.StatusOK, CountryTravels{Country: country, Travels: travels})
}

func (r Router) getCountries(c *gin.Context) {
	countries, err := r.queries.GetCountries(c)
	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "something bad happend.")
		return
	}
	c.IndentedJSON(http.StatusOK, countries)
}

func (r Router) addTraveler(c *gin.Context) {
	var input AddTravelerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Input is wrong")
		return
	}

	traveler, err := r.queries.CreateTraveler(c, queries.CreateTravelerParams{
		Name: input.Name,
		Link: input.Link,
	})

	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "could not create traveler.")
		return
	}
	c.IndentedJSON(http.StatusCreated, traveler)
}

func (r Router) addResource(c *gin.Context) {
	var input AddResourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Input is wrong")
		return
	}

	traveler, err := r.queries.CreateResource(c, queries.CreateResourceParams{
		Title: input.Title,
		Link:  input.Link,
		Image: input.Image,
	})

	if err != nil {
		log.Println("this is error: ", err)
		c.IndentedJSON(http.StatusInternalServerError, "could not create resource.")
		return
	}
	c.IndentedJSON(http.StatusCreated, traveler)
}

func (r Router) addTravel(c *gin.Context) {
	var input AddTravelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, fmt.Sprintf("Input is wrong: %v", err))
		return
	}

	travel, err := r.queries.GetTravelByTitle(c, input.Title)
	if err != nil && !storage.IsNoRowError(err) {
		c.IndentedJSON(http.StatusInternalServerError, "could get create travel.")
		return
	}
	travelExistsAlready := false
	if travel.Title == input.Title {
		travelExistsAlready = true
	}

	if !travelExistsAlready {
		travel, err = r.queries.CreateTravel(c, queries.CreateTravelParams{
			Title:     input.Title,
			StartedAt: input.StartedAt,
			EndedAt:   input.EndedAt,
		})
		if err != nil && !storage.IsNoRowError(err) {
			log.Println("this is error: ", err)
			c.IndentedJSON(http.StatusInternalServerError, "could not create travel.")
			return
		}
	}

	for _, countryInput := range input.Countries {
		country, err := r.queries.GetCountryByCode(c, countryInput)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, "could not find country.")
			return
		}
		_, err = r.queries.CreateTravelCountry(c, queries.CreateTravelCountryParams{
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
		_, err = r.queries.CreateTravelResource(c, queries.CreateTravelResourceParams{
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
		_, err = r.queries.CreateTravelTraveler(c, queries.CreateTravelTravelerParams{
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

func resourcesRecordToResourceType(resourcesRecords []queries.GetResourcesByTravelIDRow) []Resource {
	var resources []Resource
	for _, resourceRecord := range resourcesRecords {
		resources = append(resources, Resource{
			Link:  resourceRecord.Link,
			Image: resourceRecord.Image,
			Title: resourceRecord.Title,
		})
	}

	return resources
}

func travelersRecordToTravelerType(travelersRecords []queries.GetTravelersByTravelIDRow) []Traveler {
	var travelers []Traveler
	for _, travelerRecord := range travelersRecords {
		travelers = append(travelers, Traveler{
			Link: travelerRecord.Link,
			Name: travelerRecord.Name,
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
