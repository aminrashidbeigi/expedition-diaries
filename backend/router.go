package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	db *gorm.DB
}

type AddTravelInput struct {
	Travelers []Traveler `json:"travelers" binding:"required"`
	Resources []Resource `json:"resources" binding:"required"`
	Countries []Country  `json:"countries" binding:"required"`
}

func (r Router) getCountryByCode(c *gin.Context) {
	code := c.Param("code")

	if len(code) != 2 {
		log.Println("Bad request")
		c.IndentedJSON(http.StatusBadRequest, "Country code not found. it should be 2 letters.")
		return
	}

	var country Country
	r.db.Where("code = ?", strings.ToUpper(code)).First(&country)

	if country.Name == "" {
		c.IndentedJSON(http.StatusNotFound, "Country code not found.")
		return
	}

	c.IndentedJSON(http.StatusOK, country)
}

func (r Router) getCountries(c *gin.Context) {
	var countries []Country
	r.db.Find(&countries)

	c.IndentedJSON(http.StatusOK, countries)
}

func (r Router) addTravel(c *gin.Context) {
	var input AddTravelInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Input is wrong")
		return
	}

	travel := &Travel{
		Travelers: input.Travelers,
		Resources: input.Resources,
		Countries: input.Countries,
	}
	r.db.Create(travel)
	c.IndentedJSON(http.StatusCreated, travel)
}
