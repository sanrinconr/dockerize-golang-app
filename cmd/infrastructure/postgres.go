package infrastructure

import (
	"database/sql"
	"errors"
	"log"

	// Postgres db driver.
	_ "github.com/lib/pq"
)

// Postgres an implementation must have an url of where is the database.
type Postgres struct {
	url string
}

// NewPosgres create a new instance of postgres db.
func NewPosgres(url string) (*Postgres, error) {
	if url == "" {
		return nil, errors.New("url cannot be empty")
	}

	return &Postgres{url: url}, nil
}

// Ping make a ping request to the database.
func (p *Postgres) Ping() error {
	db, err := sql.Open("postgres", p.url)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Panic("error closing connection db")
		}
	}(db)

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
