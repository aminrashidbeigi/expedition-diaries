-- name: GetCountries :many
SELECT * FROM countries;

-- name: GetCountryByCode :one
SELECT * FROM countries
WHERE code = $1 LIMIT 1;

-- name: CreateCountry :one
INSERT INTO countries (
  name, code
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetTravlesByCountryCode :many
SELECT * FROM travels
INNER JOIN travel_countries on travels.id = travel_countries.travel_id
INNER JOIN countries on travel_countries.country_id = countries.id
WHERE countries.code = $1
ORDER BY travels.id DESC;


-- name: GetCountriesByTravelID :many
SELECT * FROM countries
INNER JOIN travel_countries on countries.id = travel_countries.country_id
INNER JOIN travels on travel_countries.travel_id = travels.id
WHERE travels.id = $1;


-- name: CreateResource :one
INSERT INTO resources (
  title, link, image, language, type
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;


-- name: GetResourcesByTravelID :many
SELECT * FROM resources
INNER JOIN travel_resources on travel_resources.resource_id = resources.id
INNER JOIN travels on travels.id = travel_resources.travel_id
WHERE travels.id = $1;


-- name: CreateTraveler :one
INSERT INTO travelers (
  name, link, image, nationality, slug
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetTravelersByTravelID :many
SELECT * FROM travelers
INNER JOIN travel_travelers on travelers.id = travel_travelers.traveler_id
INNER JOIN travels on travel_travelers.travel_id = travels.id
WHERE travels.id = $1;


-- name: CreateTravel :one
INSERT INTO travels (
  title, description, started_at, ended_at, route, slug
) VALUES (
  $1, $2, $3, $4, $5, $6
) 
ON CONFLICT DO NOTHING RETURNING *;

-- name: GetTravelByTitle :one
SELECT * FROM travels
WHERE title = $1 LIMIT 1;

-- name: CreateTravelCountry :one
INSERT INTO travel_countries (
  travel_id, country_id
) VALUES (
  $1, $2
) 
ON CONFLICT DO NOTHING RETURNING *;

-- name: CreateTravelTraveler :one
INSERT INTO travel_travelers (
  travel_id, traveler_id
) VALUES (
  $1, $2
) 
ON CONFLICT DO NOTHING RETURNING *;

-- name: CreateTravelResource :one
INSERT INTO travel_resources (
  travel_id, resource_id
) VALUES (
  $1, $2
)
ON CONFLICT DO NOTHING RETURNING *;

-- name: GetTravels :many
SELECT * FROM travels ORDER BY id DESC LIMIT $1 OFFSET $2;

-- name: GetTravelBySlug :many
SELECT * FROM travels WHERE slug=$1 LIMIT 1;


-- name: GetTravelers :many
SELECT * FROM travelers ORDER BY id DESC;

-- name: GetTravelsByTravelerSlug :many
SELECT * FROM travels
INNER JOIN travel_travelers on travels.id = travel_travelers.travel_id
INNER JOIN travelers on travel_travelers.traveler_id = travelers.id
WHERE travelers.slug = @slug::text
ORDER BY travels.id DESC;

-- name: GetTravelerBySlug :one
SELECT * FROM travelers WHERE slug = @slug::text LIMIT 1;