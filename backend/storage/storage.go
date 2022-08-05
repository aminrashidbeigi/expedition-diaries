package storage

import (
	"context"
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strings"

	"example.com/history-travelers/storage/queries"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func GetQueries() (*queries.Queries, error) {
	db, err := sql.Open("postgres", "user=postgres dbname=travels")
	if err != nil {
		return nil, err
	}
	queries_storage := queries.New(db)
	if err != nil {
		return nil, err
	}
	return queries_storage, nil
}

func Migrate() {
	m, err := migrate.New(
		"file://storage/migrations",
		"postgres://postgres@localhost:5432/travels?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func SeedCountries() {
	f, err := os.Open("storage/data/countries.csv")
	if err != nil {
		log.Fatal(err)
	}

	println("helllo")
	// remember to close the file at the end of the program
	defer f.Close()
	ctx := context.Background()
	defer ctx.Done()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	query_storage, err := GetQueries()
	if err != nil {
		log.Fatal(err)
	}
	for i, line := range data {
		if i > 0 { // omit header line
			var code string
			var name string

			for j, field := range line {
				if j == 0 {
					name = field
				} else if j == 1 {
					code = field
				}
			}
			query_storage.CreateCountry(ctx, queries.CreateCountryParams{
				Name: name,
				Code: strings.ToLower(code),
			})
		}
	}
}

func IsNoRowError(err error) bool {
	return err == sql.ErrNoRows
}
