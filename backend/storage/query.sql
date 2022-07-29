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





-- name: CreateResource :one
INSERT INTO resources (
  title, link, image
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateTraveler :one
INSERT INTO travelers (
  name, link
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateTravel :one
INSERT INTO travels DEFAULT VALUES RETURNING *;

-- name: CreateTravelCountry :one
INSERT INTO travel_countries (
  travel_id, country_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateTravelTraveler :one
INSERT INTO travel_travelers (
  travel_id, traveler_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateTravelResource :one
INSERT INTO travel_resources (
  travel_id, resource_id
) VALUES (
  $1, $2
)
RETURNING *;
