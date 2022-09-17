package storage

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aminrashidbeigi/history-travels/config"
	"github.com/aminrashidbeigi/history-travels/storage/queries"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Storage struct {
	Config config.DBConfig
}

func (s Storage) getDataSourceName() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		s.Config.Username, s.Config.Password, s.Config.DBName, s.Config.Host)
}

func (s Storage) GetQueries() (*queries.Queries, error) {
	db, err := sql.Open("postgres", s.getDataSourceName())
	if err != nil {
		return nil, err
	}
	queries_storage := queries.New(db)
	if err != nil {
		return nil, err
	}
	return queries_storage, nil
}

func (s Storage) getMigrateDBAddress() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		s.Config.Username, s.Config.Password, s.Config.Host, s.Config.Port, s.Config.DBName)
}

func (s Storage) Migrate() {
	m, err := migrate.New(
		"file://storage/migrations",
		s.getMigrateDBAddress())
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func (s Storage) SeedCountries() {
	f, err := os.Open("/server/storage/data/countries.csv")
	if err != nil {
		log.Fatal(err)
	}

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

	query_storage, err := s.GetQueries()
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
			_, err := query_storage.CreateCountry(ctx, queries.CreateCountryParams{
				Name: name,
				Code: strings.ToLower(code),
			})
			if err != nil {
				log.Printf("can not create country: %v", err)
			}
		}
	}
}

func IsNoRowError(err error) bool {
	return err == sql.ErrNoRows
}
