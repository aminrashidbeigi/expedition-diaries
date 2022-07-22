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