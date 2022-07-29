// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package queries

import (
	"context"
)

const createCountry = `-- name: CreateCountry :one
INSERT INTO countries (
  name, code
) VALUES (
  $1, $2
)
RETURNING id, code, name
`

type CreateCountryParams struct {
	Name string
	Code string
}

func (q *Queries) CreateCountry(ctx context.Context, arg CreateCountryParams) (Country, error) {
	row := q.db.QueryRowContext(ctx, createCountry, arg.Name, arg.Code)
	var i Country
	err := row.Scan(&i.ID, &i.Code, &i.Name)
	return i, err
}

const createResource = `-- name: CreateResource :one
INSERT INTO resources (
  title, link, image
) VALUES (
  $1, $2, $3
)
RETURNING id, title, link, image
`

type CreateResourceParams struct {
	Title string
	Link  string
	Image string
}

func (q *Queries) CreateResource(ctx context.Context, arg CreateResourceParams) (Resource, error) {
	row := q.db.QueryRowContext(ctx, createResource, arg.Title, arg.Link, arg.Image)
	var i Resource
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Link,
		&i.Image,
	)
	return i, err
}

const createTravel = `-- name: CreateTravel :one
INSERT INTO travels DEFAULT VALUES RETURNING id
`

func (q *Queries) CreateTravel(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, createTravel)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createTravelCountry = `-- name: CreateTravelCountry :one
INSERT INTO travel_countries (
  travel_id, country_id
) VALUES (
  $1, $2
)
RETURNING travel_id, country_id
`

type CreateTravelCountryParams struct {
	TravelID  int32
	CountryID int32
}

func (q *Queries) CreateTravelCountry(ctx context.Context, arg CreateTravelCountryParams) (TravelCountry, error) {
	row := q.db.QueryRowContext(ctx, createTravelCountry, arg.TravelID, arg.CountryID)
	var i TravelCountry
	err := row.Scan(&i.TravelID, &i.CountryID)
	return i, err
}

const createTravelResource = `-- name: CreateTravelResource :one
INSERT INTO travel_resources (
  travel_id, resource_id
) VALUES (
  $1, $2
)
RETURNING travel_id, resource_id
`

type CreateTravelResourceParams struct {
	TravelID   int32
	ResourceID int32
}

func (q *Queries) CreateTravelResource(ctx context.Context, arg CreateTravelResourceParams) (TravelResource, error) {
	row := q.db.QueryRowContext(ctx, createTravelResource, arg.TravelID, arg.ResourceID)
	var i TravelResource
	err := row.Scan(&i.TravelID, &i.ResourceID)
	return i, err
}

const createTravelTraveler = `-- name: CreateTravelTraveler :one
INSERT INTO travel_travelers (
  travel_id, traveler_id
) VALUES (
  $1, $2
)
RETURNING travel_id, traveler_id
`

type CreateTravelTravelerParams struct {
	TravelID   int32
	TravelerID int32
}

func (q *Queries) CreateTravelTraveler(ctx context.Context, arg CreateTravelTravelerParams) (TravelTraveler, error) {
	row := q.db.QueryRowContext(ctx, createTravelTraveler, arg.TravelID, arg.TravelerID)
	var i TravelTraveler
	err := row.Scan(&i.TravelID, &i.TravelerID)
	return i, err
}

const createTraveler = `-- name: CreateTraveler :one
INSERT INTO travelers (
  name, link
) VALUES (
  $1, $2
)
RETURNING id, name, link
`

type CreateTravelerParams struct {
	Name string
	Link string
}

func (q *Queries) CreateTraveler(ctx context.Context, arg CreateTravelerParams) (Traveler, error) {
	row := q.db.QueryRowContext(ctx, createTraveler, arg.Name, arg.Link)
	var i Traveler
	err := row.Scan(&i.ID, &i.Name, &i.Link)
	return i, err
}

const getCountries = `-- name: GetCountries :many
SELECT id, code, name FROM countries
`

func (q *Queries) GetCountries(ctx context.Context) ([]Country, error) {
	rows, err := q.db.QueryContext(ctx, getCountries)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Country
	for rows.Next() {
		var i Country
		if err := rows.Scan(&i.ID, &i.Code, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountriesByTravelID = `-- name: GetCountriesByTravelID :many
SELECT countries.id, code, name, travel_id, country_id, travels.id FROM countries
INNER JOIN travel_countries on countries.id = travel_countries.country_id
INNER JOIN travels on travel_countries.travel_id = travels.id
WHERE travels.id = $1
`

type GetCountriesByTravelIDRow struct {
	ID        int32
	Code      string
	Name      string
	TravelID  int32
	CountryID int32
	ID_2      int32
}

func (q *Queries) GetCountriesByTravelID(ctx context.Context, id int32) ([]GetCountriesByTravelIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getCountriesByTravelID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCountriesByTravelIDRow
	for rows.Next() {
		var i GetCountriesByTravelIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Name,
			&i.TravelID,
			&i.CountryID,
			&i.ID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCountryByCode = `-- name: GetCountryByCode :one
SELECT id, code, name FROM countries
WHERE code = $1 LIMIT 1
`

func (q *Queries) GetCountryByCode(ctx context.Context, code string) (Country, error) {
	row := q.db.QueryRowContext(ctx, getCountryByCode, code)
	var i Country
	err := row.Scan(&i.ID, &i.Code, &i.Name)
	return i, err
}

const getResourcesByTravelID = `-- name: GetResourcesByTravelID :many
SELECT resources.id, title, link, image, travel_id, resource_id, travels.id FROM resources
INNER JOIN travel_resources on travel_resources.resource_id = resources.id
INNER JOIN travels on travels.id = travel_resources.travel_id
WHERE travels.id = $1
`

type GetResourcesByTravelIDRow struct {
	ID         int32
	Title      string
	Link       string
	Image      string
	TravelID   int32
	ResourceID int32
	ID_2       int32
}

func (q *Queries) GetResourcesByTravelID(ctx context.Context, id int32) ([]GetResourcesByTravelIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getResourcesByTravelID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetResourcesByTravelIDRow
	for rows.Next() {
		var i GetResourcesByTravelIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Link,
			&i.Image,
			&i.TravelID,
			&i.ResourceID,
			&i.ID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTravelersByTravelID = `-- name: GetTravelersByTravelID :many
SELECT travelers.id, name, link, travel_id, traveler_id, travels.id FROM travelers
INNER JOIN travel_travelers on travelers.id = travel_travelers.traveler_id
INNER JOIN travels on travel_travelers.travel_id = travels.id
WHERE travels.id = $1
`

type GetTravelersByTravelIDRow struct {
	ID         int32
	Name       string
	Link       string
	TravelID   int32
	TravelerID int32
	ID_2       int32
}

func (q *Queries) GetTravelersByTravelID(ctx context.Context, id int32) ([]GetTravelersByTravelIDRow, error) {
	rows, err := q.db.QueryContext(ctx, getTravelersByTravelID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTravelersByTravelIDRow
	for rows.Next() {
		var i GetTravelersByTravelIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Link,
			&i.TravelID,
			&i.TravelerID,
			&i.ID_2,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTravlesByCountryCode = `-- name: GetTravlesByCountryCode :many
SELECT travels.id, travel_id, country_id, countries.id, code, name FROM travels
INNER JOIN travel_countries on travels.id = travel_countries.travel_id
INNER JOIN countries on travel_countries.country_id = countries.id
WHERE countries.code = $1
`

type GetTravlesByCountryCodeRow struct {
	ID        int32
	TravelID  int32
	CountryID int32
	ID_2      int32
	Code      string
	Name      string
}

func (q *Queries) GetTravlesByCountryCode(ctx context.Context, code string) ([]GetTravlesByCountryCodeRow, error) {
	rows, err := q.db.QueryContext(ctx, getTravlesByCountryCode, code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTravlesByCountryCodeRow
	for rows.Next() {
		var i GetTravlesByCountryCodeRow
		if err := rows.Scan(
			&i.ID,
			&i.TravelID,
			&i.CountryID,
			&i.ID_2,
			&i.Code,
			&i.Name,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
