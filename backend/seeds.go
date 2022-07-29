package main

import (
	"context"
	"encoding/csv"
	"log"
	"os"

	"example.com/history-travelers/storage/queries"
)

type Seeder struct {
	queries *queries.Queries
}

func (s Seeder) SeedCountries() {
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
			println("hello")
			s.queries.CreateCountry(ctx, queries.CreateCountryParams{
				Name: name,
				Code: code,
			})
		}
	}
}
