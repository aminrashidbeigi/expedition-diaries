package main

import (
	"gorm.io/gorm"
)

// Country table
type Country struct {
	gorm.Model
	Code    string `gorm:"uniqueIndex"`
	Name    string
	Travels []Travel `gorm:"many2many:country_travels;"`
}

// Traveler table
type Traveler struct {
	gorm.Model
	Name    string `gorm:"uniqueIndex"`
	Link    string
	Travels []Travel `gorm:"many2many:travel_travelers;"`
}

// Resource table
type Resource struct {
	gorm.Model
	Title string `gorm:"uniqueIndex"`
	Link  string
	Image string
}

// Travel table
type Travel struct {
	gorm.Model
	Resources []Resource `gorm:"ForeignKey:ID;"`
	Travelers []Traveler `gorm:"many2many:travel_travelers;"`
	Countries []Country  `gorm:"many2many:travel_countries;"`
}
