// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package queries

import (
	"database/sql"
)

type Country struct {
	ID   int32
	Code string
	Name string
}

type Resource struct {
	ID       int32
	Title    string
	Link     string
	Image    string
	Language sql.NullString
	Type     sql.NullString
}

type Travel struct {
	ID          int32
	Title       string
	StartedAt   string
	EndedAt     string
	Route       sql.NullString
	Slug        sql.NullString
	Description sql.NullString
}

type TravelCountry struct {
	TravelID  int32
	CountryID int32
}

type TravelResource struct {
	TravelID   int32
	ResourceID int32
}

type TravelTraveler struct {
	TravelID   int32
	TravelerID int32
}

type Traveler struct {
	ID          int32
	Name        string
	Link        string
	Image       sql.NullString
	Nationality sql.NullString
	Slug        sql.NullString
}
