package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	db *gorm.DB
}

func (r Router) getCountryByCode(c *gin.Context) {
	code := c.Param("code")
	var country Country
	r.db.Where("code = ?", strings.ToUpper(code)).First(&country)
	c.IndentedJSON(http.StatusOK, country)
}

func (r Router) getCountries(c *gin.Context) {
	var countries []Country
	r.db.Find(&countries)

	c.IndentedJSON(http.StatusOK, countries)
}

func (r Router) addTravel(c *gin.Context) {
	travel := &Travel{
		Travelers: []Traveler{{Name: "Amin Rashir.db.igi"}},
		Resources: []Resource{},
		Countries: []Country{{Name: "Iran"}, {Name: "United Arab Emirates"}},
	}
	r.db.Create(travel)
	c.IndentedJSON(http.StatusCreated, travel)
}
